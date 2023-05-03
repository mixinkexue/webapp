package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
	"web/ent"
	"web/pkg/resp"
	"web/service"
)

// ListPets 未被领养的动物
func ListPets(c *gin.Context) {
	resp.CommonResp(c, service.ListPet(context.Background()))
}

type CreatePetReq struct {
	AnimalBreed  string `json:"animal_breed,omitempty" binding:"required"`
	Introduction string `json:"introduction,omitempty" binding:"required"`
	Picture      string `json:"picture,omitempty" binding:"required"`
	Name         string `json:"name,omitempty" binding:"required"`
}

func CreatePet(c *gin.Context) {
	var req CreatePetReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		return
	}
	u := c.MustGet("user").(*ent.User)
	err = service.CreatePet(context.Background(), req.AnimalBreed, req.Introduction, req.Picture, req.Name, u)
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
