package controller

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"webapp/ent"
	"webapp/pkg/resp"
	"webapp/service"
)

// ListPets 未被领养的动物
func ListPets(c *gin.Context) {
	resp.CommonResp(c, service.ListPet(context.Background()))
}

func CreatePet(c *gin.Context) {
	AnimalBreed := c.PostForm("animal_breed")
	Introduction := c.PostForm("introduction")
	Name := c.PostForm("name")
	if Name == "" || Introduction == "" || AnimalBreed == "" {
		c.Error(errors.New("参数不全"))
		return
	}
	//上传
	file, err := c.FormFile("picture")
	if err != nil {
		c.Error(err)
		return
	}
	open, err := file.Open()
	if err != nil {
		c.Error(err)
		return
	}
	addr, err := service.UploadFile(strconv.Itoa(rand.Int()), file.Filename, open)
	if err != nil {
		c.Error(err)
		return
	}
	u := c.MustGet("user").(*ent.User)

	err = service.CreatePet(context.Background(), AnimalBreed, Introduction, addr, Name, u)
	if err != nil {
		c.Error(err)
		return
	}
	resp.CommonResp(c, nil)
}

func AdoptAPet(c *gin.Context) {
	t := c.PostForm("pet_id")
	pid, err := strconv.Atoi(t)
	if err != nil {
		c.Error(err)
		return
	}
	err = service.AdoptAPet(context.Background(), pid, c.MustGet("user").(*ent.User))
	if err != nil {
		c.Error(err)
		return
	}
	resp.CommonResp(c, nil)
}

func GetMine(c *gin.Context) {
	res, err := service.GetMine(context.Background(), c.MustGet("user").(*ent.User))
	if err != nil {
		c.Error(err)
		return
	}
	resp.CommonResp(c, res)
}

func UpdatePet(c *gin.Context) {
	AnimalBreed := c.PostForm("animalBreed")
	Introduction := c.PostForm("introduction")
	Name := c.PostForm("name")
	Id := c.PostForm("id")
	if Name == "" || Introduction == "" || AnimalBreed == "" {
		c.Error(errors.New("参数不全"))
		return
	}
	//上传
	file, _ := c.FormFile("picture")
	var addr string
	if file != nil {
		open, err := file.Open()
		if err != nil {
			c.Error(err)
			return
		}
		addr, err = service.UploadFile(strconv.Itoa(rand.Int()), file.Filename, open)
		if err != nil {
			c.Error(err)
			return
		}
	}

	id, _ := strconv.Atoi(Id)
	pet, err := service.UpdatePet(context.Background(), id, AnimalBreed, Introduction, addr, Name)
	if err != nil {
		c.Error(err)
		return
	}
	resp.CommonResp(c, pet)
}

func GetPet(ctx *gin.Context) {
	id := ctx.Query("id")
	atoi, _ := strconv.Atoi(id)
	u, err := service.SelectPet(ctx, atoi)
	if err != nil {
		ctx.Error(err)
	}
	resp.CommonResp(ctx, u)
}
