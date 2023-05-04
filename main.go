package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"webapp/controller"
	"webapp/pkg/middleware"
)

func main() {
	r := gin.Default()
	r.Static("/static", "font")
	store := cookie.NewStore([]byte("store"))
	webapp := r.Group("webapp")
	webapp.Use(middleware.Cors())
	webapp.Use(middleware.HandleError)
	webapp.Use(sessions.Sessions("token", store))

	{
		user := webapp.Group("user")
		user.POST("register", controller.Register)
		user.POST("login", controller.Login)
		user.Use(middleware.Auth)
		user.POST("logout", controller.Logout)
		user.GET("info", controller.UserInfo)
		user.POST("update", controller.Update)
		user.POST("updateHead", controller.UpdateHead)
	}
	{
		pet := webapp.Group("pet")
		pet.GET("list", controller.ListPets)
		pet.Use(middleware.Auth)
		pet.POST("create", controller.CreatePet)
		pet.POST("update", controller.UpdatePet)
		pet.GET("getPet", controller.GetPet)
		pet.POST("adopt", controller.AdoptAPet)
		pet.GET("mine", controller.GetMine)
	}

	r.Run()
}
