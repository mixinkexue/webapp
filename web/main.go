package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"web/controller"
	"web/pkg/middleware"
)

func main() {
	r := gin.Default()
	r.Static("/static", "font")
	store := cookie.NewStore([]byte("store"))
	r.Use(middleware.Cors())
	r.Use(middleware.HandleError)
	r.Use(sessions.Sessions("token", store))

	{
		user := r.Group("user")
		user.POST("register", controller.Register)
		user.POST("login", controller.Login)
		user.Use(middleware.Auth)
		user.POST("logout", controller.Logout)
		user.GET("info", controller.UserInfo)
		user.POST("update", controller.Update)
		user.POST("updateHead", controller.UpdateHead)
	}
	{
		pet := r.Group("pet")
		pet.GET("list", controller.ListPets)
		pet.Use(middleware.Auth)
		pet.POST("create", controller.CreatePet)
		pet.POST("adopt", controller.AdoptAPet)
		pet.GET("mine", controller.GetMine)
	}

	r.Run()
}
