package controller

import (
	"context"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"web/ent"
	"web/pkg/resp"
	"web/service"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var req LoginReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	val, err := service.Login(context.Background(), req.Username, req.Password)
	if err != nil {
		ctx.Error(err)
		return
	}
	session := sessions.Default(ctx)
	session.Set("token", val)
	session.Save()
	resp.CommonResp(ctx, nil)
}
func Register(ctx *gin.Context) {
	var req LoginReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		return
	}
	val, err := service.Register(context.Background(), req.Username, req.Password)
	if err != nil {
		ctx.Error(err)
		return
	}
	session := sessions.Default(ctx)
	session.Set("token", val)
	session.Save()
	resp.CommonResp(ctx, nil)
}

func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	u := ctx.MustGet("user").(*ent.User)
	if err := service.Logout(context.Background(), u); err != nil {
		ctx.Error(err)
		return
	}
	resp.CommonResp(ctx, nil)
}

func UserInfo(ctx *gin.Context) {
	resp.CommonResp(ctx, ctx.MustGet("user").(*ent.User))
}

type uReq struct {
	Age    string `json:"age" binding:"required"`
	Phone  string `json:"phone" binding:"required"`
	Email  string `json:"email" binding:"required"`
	Addr   string `json:"addr" binding:"required"`
	Career string `json:"career" binding:"required"`
}

func Update(ctx *gin.Context) {
	var r uReq
	err := ctx.ShouldBind(&r)
	if err != nil {
		ctx.Error(err)
		return
	}
	u := ctx.MustGet("user").(*ent.User)
	u.Age, _ = strconv.Atoi(r.Age)
	u.Phone = r.Phone
	u.Email = r.Email
	u.Addr = r.Addr
	u.Career = r.Career
	err = service.UpdateUser(context.Background(), u)
	if err != nil {
		ctx.Error(err)
	}
	resp.CommonResp(ctx, nil)
}

func UpdateHead(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.Error(err)
		return
	}
	open, err := file.Open()
	if err != nil {
		c.Error(err)
		return
	}
	url, err := service.UploadFile(strconv.Itoa(rand.Int()), file.Filename, open)
	if err != nil {
		c.Error(err)
		return
	}
	u := c.MustGet("user").(*ent.User)
	u.Avatar = url
	err = service.UpdateUser(context.Background(), u)
	if err != nil {
		c.Error(err)
	}
	resp.CommonResp(c, nil)
}
