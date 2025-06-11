package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go-gin-example/pkg/app"
	"go-gin-example/pkg/e"
	"go-gin-example/pkg/util"
	"go-gin-example/service/auth_service"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Authenticate user
// @Description Authenticates a user by validating the provided username and password, returning a response indicating success or failure.
// @Tags authentication
// @Accept application/x-www-form-urlencoded
// @Produces application/json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} app.Response "Successful authentication response"
// @Failure 400 {object} app.Response "Invalid request parameters"
// @Failure 401 {object} app.Response "Unauthorized - Invalid credentials"
// @Failure 500 {object} app.Response "Internal server error"
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username := c.PostForm("username")
	password := c.PostForm("password")

	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
