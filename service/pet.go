package service

import (
	"context"
	"errors"
	"log"
	"webapp/ent"
	"webapp/ent/pet"
	"webapp/ent/user"
)

func ListPet(ctx context.Context) []*ent.Pet {
	all, err := client.Pet.Query().Where(pet.Not(pet.HasAdopter())).All(ctx)
	if err != nil {
		log.Println(err)
	}
	return all
}

func CreatePet(ctx context.Context, AnimalBreed, Introduction, Picture, Name string, u *ent.User) error {
	return client.Pet.Create().SetOwner(u).SetAnimalBreed(AnimalBreed).SetIntroduction(Introduction).SetPicture(Picture).
		SetName(Name).Exec(ctx)
}

func AdoptAPet(ctx context.Context, pid int, u *ent.User) error {
	count, err := client.Pet.Query().Where(pet.ID(pid)).Where(pet.Not(pet.HasAdopter())).Count(ctx)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("不存在被想要领养的宠物")
	}
	//todo:不能领养自己的
	return client.Pet.Update().Where(pet.ID(pid)).SetAdopter(u).Exec(ctx)
}

func GetMine(ctx context.Context, u *ent.User) (res [3][]*ent.Pet, err error) {
	res[0], err = client.Pet.Query().Where(pet.HasAdopterWith(user.ID(u.ID))).All(ctx)
	if err != nil {
		return
	}

	res[1], err = client.Pet.Query().Where(pet.HasOwnerWith(user.ID(u.ID))).All(ctx)
	if err != nil {
		return
	}
	//没有被领养的
	res[2] = make([]*ent.Pet, 0)
	var count = 0
	for i := len(res[1]) - 1; i >= 0; i-- {
		count, err = client.Pet.Query().Where(pet.ID(res[1][i].ID), pet.HasAdopter()).Count(ctx)
		if err != nil {
			return
		}
		if count == 0 {
			res[2] = append(res[2], res[1][i])
			res[1] = append(res[1][:i], res[1][i+1:]...)
		}
	}
	return
}
