package system

import (
	"ferry/models/system"
	"ferry/pkg/ldap"
	"ferry/pkg/logger"
	"ferry/tools"
	"ferry/tools/app"
	
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/*
  @Author : lanyulei
*/

// @Summary 列表数据
// @Description 获取JSON
// @Tags 用户
// @Param username query string false "username"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/sysUserList [get]
// @Security Bearer
func GetSysUserList(c *gin.Context) {
	var (
		pageIndex = 1
		pageSize  = 10
		err       error
		data      system.SysUser
	)

	size := c.Request.FormValue("pageSize")
	if size != "" {
		pageSize = tools.StrToInt(err, size)
	}

	index := c.Request.FormValue("pageIndex")
	if index != "" {
		pageIndex = tools.StrToInt(err, index)
	}

	data.Username = c.Request.FormValue("username")
	data.NickName = c.Request.FormValue("nickName")
	data.Status = c.Request.FormValue("status")
	data.Phone = c.Request.FormValue("phone")

	postId := c.Request.FormValue("postId")
	data.PostId, _ = tools.StringToInt(postId)

	deptId := c.Request.FormValue("deptId")
	data.DeptId, _ = tools.StringToInt(deptId)

	result, count, err := data.GetPage(pageSize, pageIndex)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

// @Summary 获取用户
// @Description 获取JSON
// @Tags 用户
// @Param userId path int true "用户编码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sysUser/{userId} [get]
// @Security
func GetSysUser(c *gin.Context) {
	var SysUser system.SysUser
	SysUser.UserId, _ = tools.StringToInt(c.Param("userId"))
	result, err := SysUser.Get()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	var SysRole system.SysRole
	var Post system.Post
	roles, _ := SysRole.GetList()
	posts, _ := Post.GetList()

	postIds := make([]int, 0)
	postIds = append(postIds, result.PostId)

	roleIds := make([]int, 0)
	roleIds = append(roleIds, result.RoleId)
	app.Custum(c, gin.H{
		"code":    200,
		"data":    result,
		"postIds": postIds,
		"roleIds": roleIds,
		"roles":   roles,
		"posts":   posts,
	})
}

// @Summary 获取当前登录用户
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/profile [get]
// @Security
func GetSysUserProfile(c *gin.Context) {
	var (
		Dept    system.Dept
		Post    system.Post
		SysRole system.SysRole
		SysUser system.SysUser
	)
	userId := tools.GetUserIdStr(c)
	SysUser.UserId, _ = tools.StringToInt(userId)
	result, err := SysUser.Get()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	//获取角色列表
	roles, err := SysRole.GetList()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	//获取职位列表
	posts, err := Post.GetList()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	//获取部门列表
	Dept.DeptId = result.DeptId
	dept, err := Dept.Get()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	postIds := make([]int, 0)
	postIds = append(postIds, result.PostId)

	roleIds := make([]int, 0)
	roleIds = append(roleIds, result.RoleId)

	app.Custum(c, gin.H{
		"code":    200,
		"data":    result,
		"postIds": postIds,
		"roleIds": roleIds,
		"roles":   roles,
		"posts":   posts,
		"dept":    dept,
	})
}

// @Summary 获取用户角色和职位
// @Description 获取JSON
// @Tags 用户
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sysUser [get]
// @Security
func GetSysUserInit(c *gin.Context) {
	var SysRole system.SysRole
	var Post system.Post
	roles, err := SysRole.GetList()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	posts, err := Post.GetList()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	mp := make(map[string]interface{}, 2)
	mp["roles"] = roles
	mp["posts"] = posts
	app.OK(c, mp, "")
}

// @Summary 创建用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body system.SysUser true "用户数据"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/sysUser [post]
func InsertSysUser(c *gin.Context) {
	var sysuser system.SysUser
	err := c.MustBindWith(&sysuser, binding.JSON)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	sysuser.CreateBy = tools.GetUserIdStr(c)
	id, err := sysuser.Insert()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, id, "添加成功")
}

// @Summary 修改用户数据
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body system.SysUser true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "修改失败"}"
// @Router /api/v1/sysuser/{userId} [put]
func UpdateSysUser(c *gin.Context) {
	var data system.SysUser
	err := c.Bind(&data)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	data.UpdateBy = tools.GetUserIdStr(c)
	result, err := data.Update(data.UserId)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, result, "修改成功")
}

// @Summary 删除用户数据
// @Description 删除数据
// @Tags 用户
// @Param userId path int true "userId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/sysuser/{userId} [delete]
func DeleteSysUser(c *gin.Context) {
	var data system.SysUser
	data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("userId", c)
	result, err := data.BatchDelete(IDS)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, result, "删除成功")
}

// @Summary 修改头像
// @Description 获取JSON
// @Tags 用户
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/user/profileAvatar [post]
func InsetSysUserAvatar(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	files := form.File["upload[]"]
	guid := uuid.New().String()
	filPath := "static/uploadfile/" + guid + ".jpg"
	for _, file := range files {
		logger.Info(file.Filename)
		// 上传文件至指定目录
		err = c.SaveUploadedFile(file, filPath)
		if err != nil {
			app.Error(c, -1, err, "")
			return
		}
	}
	sysuser := system.SysUser{}
	sysuser.UserId = tools.GetUserId(c)
	sysuser.Avatar = "/" + filPath
	sysuser.UpdateBy = tools.GetUserIdStr(c)
	_, _ = sysuser.Update(sysuser.UserId)
	app.OK(c, filPath, "修改成功")
}

func SysUserUpdatePwd(c *gin.Context) {
	var pwd system.SysUserPwd
	err := c.Bind(&pwd)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	if pwd.PasswordType == 0 {
		sysuser := system.SysUser{}
		sysuser.UserId = tools.GetUserId(c)
		_, err = sysuser.SetPwd(pwd)
		if err != nil {
			app.Error(c, -1, err, "")
			return
		}
	} else if pwd.PasswordType == 1 {
		// 修改ldap密码
		err = ldap.LdapUpdatePwd(tools.GetUserName(c), pwd.OldPassword, pwd.NewPassword)
		if err != nil {
			app.Error(c, -1, err, "")
			return
		}
	}

	app.OK(c, "", "密码修改成功")
}
