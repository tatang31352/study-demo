package main

import (
	"crypto/tls"
	"embed"
	"fmt"
	"github.com/cloudlena/adapters/logging"
	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"io/fs"
	"log"
	"net/http"
	"os"
	"s3manager/internal/app/s3manager"
	"time"
)

//go:embed web/template
var templateFS embed.FS

//go:embed web/static
var staticFS embed.FS

type configuration struct {
	Endpoint      string
	UseIam        bool
	IamEndpoint   string
	AccessKey     string
	SecretKey     string
	Region        string
	AllowDelete   bool
	ForceDownload bool
	UseSSL        bool
	SkipSSLVerify bool
	SignatureType string
	ListRecursive bool
	Port          string
	Timeout       int64
	SseType       string
	SseKey        string
}

func parseConfiguration() configuration {
	var accessKey, secretKey, iamEndpoint string
	viper.AutomaticEnv()
	viper.SetDefault("ENDPOINT", "andy_test@gmail.com")
	endpoint := viper.GetString("ENDPOINT")

	fmt.Println(endpoint)
	useIam := viper.GetBool("USE_IAM")

	if useIam {
		iamEndpoint = viper.GetString("IAM_ENDPOINT")
		if len(iamEndpoint) == 0 {
			log.Fatal("IAM_ENDPOINT not set")
		}
	} else {
		accessKey = viper.GetString("ACCESS_KEY")
		if len(accessKey) == 0 {
			log.Fatal("Access Key not set")
		}

		secretKey = viper.GetString("SECRET_KEY")
		if len(secretKey) == 0 {
			log.Fatal("Secret Key not set")
		}
	}

	region := viper.GetString("REGION")
	viper.SetDefault("ALLOW_DELETE", true)
	allowDelete := viper.GetBool("ALLOW_DELETE")

	viper.SetDefault("FORCE_DOWNLOAD", true)
	forceDownload := viper.GetBool("FORCE_DOWNLOAD")

	viper.SetDefault("USE_SSL", true)
	useSSL := viper.GetBool("USE_SSL")

	viper.SetDefault("SKIP_SSL_VERIFICATION", false)
	skipSSLVerification := viper.GetBool("SKIP_SSL_VERIFICATION")

	viper.SetDefault("SIGNATURE_TYPE", "V4")
	signatureType := viper.GetString("SIGNATURE_TYPE")

	listRecursive := viper.GetBool("LIST_RECURSIVE")

	viper.SetDefault("PORT", "8080")
	port := viper.GetString("PORT")

	viper.SetDefault("TIMEOUT", 600)
	timeout := viper.GetInt64("TIMEOUT")

	viper.SetDefault("SSE_TYPE", "")
	sseType := viper.GetString("SSE_TYPE")

	viper.SetDefault("SSE_KEY", "")
	sseKey := viper.GetString("SSE_KEY")
	return configuration{
		Endpoint:      endpoint,
		UseIam:        useIam,
		IamEndpoint:   iamEndpoint,
		AccessKey:     accessKey,
		SecretKey:     secretKey,
		Region:        region,
		AllowDelete:   allowDelete,
		ForceDownload: forceDownload,
		UseSSL:        useSSL,
		ListRecursive: listRecursive,
		SignatureType: signatureType,
		SkipSSLVerify: skipSSLVerification,
		Port:          port,
		Timeout:       timeout,
		SseType:       sseType,
		SseKey:        sseKey,
	}

}

func main() {
	configuration := parseConfiguration()
	sseType := s3manager.SSEType{Type: configuration.SseType, Key: configuration.SseKey}
	serverTimeout := time.Duration(configuration.Timeout) * time.Second

	// Set up templates
	templates, err := fs.Sub(templateFS, "web/template")
	if err != nil {
		log.Fatal(err)
	}

	// Set up statics
	statics, err := fs.Sub(staticFS, "web/static")
	if err != nil {
		log.Fatal(err)
	}

	// Set up S3 client
	opts := &minio.Options{
		Secure: configuration.UseSSL,
	}

	if configuration.UseIam {
		opts.Creds = credentials.NewIAM(configuration.IamEndpoint)
	} else {
		var signatureType credentials.SignatureType
		switch configuration.SignatureType {
		case "V2":
			signatureType = credentials.SignatureV2
		case "V4":
			signatureType = credentials.SignatureV4
		case "V4Streaming":
			signatureType = credentials.SignatureV4Streaming
		case "Anonymous":
			signatureType = credentials.SignatureAnonymous
		default:
			log.Fatalf("Invalid SIGNATURE_TYPE: %s", configuration.SignatureType)
		}

		opts.Creds = credentials.NewStatic(configuration.AccessKey, configuration.SecretKey, "", signatureType)
	}

	if configuration.Region != "" {
		opts.Region = configuration.Region
	}
	if configuration.UseSSL && configuration.SkipSSLVerify {
		opts.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} //nolint:gosec
	}
	s3, err := minio.New(configuration.Endpoint, opts)
	if err != nil {
		log.Fatalln(fmt.Errorf("error creating s3 client: %w", err))
	}

	// Set up router
	r := mux.NewRouter()
	r.Handle("/", http.RedirectHandler("/buckets", http.StatusPermanentRedirect)).Methods(http.MethodGet)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(statics)))).Methods(http.MethodGet)
	r.Handle("/buckets", s3manager.HandleBucketsView(s3, templates, configuration.AllowDelete)).Methods(http.MethodGet)
	r.PathPrefix("/buckets/").Handler(s3manager.HandleBucketView(s3, templates, configuration.AllowDelete, configuration.ListRecursive)).Methods(http.MethodGet)
	r.Handle("/api/buckets", s3manager.HandleCreateBucket(s3)).Methods(http.MethodPost)
	if configuration.AllowDelete {
		r.Handle("/api/buckets/{bucketName}", s3manager.HandleDeleteBucket(s3)).Methods(http.MethodDelete)
	}

	r.Handle("/api/buckets/{bucketName}/objects", s3manager.HandleCreateObject(s3, sseType)).Methods(http.MethodPost)
	r.Handle("/api/buckets/{bucketName}/objects/{objectName:.*}/url", s3manager.HandleGenerateUrl(s3)).Methods(http.MethodGet)
	r.Handle("/api/buckets/{bucketName}/objects/{objectName:.*}", s3manager.HandleGetObject(s3, configuration.ForceDownload)).Methods(http.MethodGet)
	if configuration.AllowDelete {
		r.Handle("/api/buckets/{bucketName}/objects/{objectName:.*}", s3manager.HandleDeleteObject(s3)).Methods(http.MethodDelete)
	}

	lr := logging.Handler(os.Stdout)(r)
	srv := &http.Server{
		Addr:         ":" + configuration.Port,
		Handler:      lr,
		ReadTimeout:  serverTimeout,
		WriteTimeout: serverTimeout,
	}
	log.Fatal(srv.ListenAndServe())
}
