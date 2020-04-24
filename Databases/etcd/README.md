#go-micro使用etcd

    packge main

    import(  
        "github.com/micro/go-micro"  
        "github.com/micro/go-plugins/registry/etcdv3"  
        "log"  
        "time"  
    )

    func main(){  
    
        resistre := etcdv3.NewRegistry(func(options*registry.Options){
            //etcd地址
           options.Adds = []string{"127.0.0.1:2379"}
           //etcd用户名密码,如果设置的话
           etcdv3.Auth("root","password")(options) 
        })  
    
        service := micro.NewService(
            micro.Registry(registre),
            micro.Name("greeter")
            micro.RegisterTTL(time.Second*30),
            micro,RegisterInterval(time.Second*15)
        )
        
        service.Init()
        if err := service.Run();err != nil{
            log.Fatal(err)
        }
    }