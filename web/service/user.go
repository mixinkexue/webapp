package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"web/ent"
	"web/ent/token"
	"web/ent/user"
)

func Register(ctx context.Context, username, passwd string) (string, error) {
	//查询是否存在一个同名用户
	count, err := client.User.Query().Where(user.Username(username)).Count(ctx)
	if err != nil {
		return "", err
	}
	if count != 0 {
		return "", errors.New("存在一个用户了")
	}
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	tx, err := client.Tx(ctx)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	u, err := tx.User.Create().SetUsername(username).SetPassword(string(hashPwd)).Save(ctx)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	//生成token
	newUuid, err := uuid.NewUUID()
	if err != nil {
		tx.Rollback()
		return "", err
	}
	token := newUuid.String()
	//token与uid对应
	_, err = tx.Token.Create().SetVal(token).SetOwner(u).Save(ctx)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return "", err
	}
	return token, nil
}

func Login(ctx context.Context, username, passwd string) (string, error) {
	u, err := client.User.Query().Where(user.Username(username)).Only(ctx)
	if err != nil {
		return "", err
	}

	newUuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	val := newUuid.String()
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passwd))
	if err != nil {
		return "", err
	}
	_, err = client.Token.Create().SetVal(val).SetOwner(u).Save(ctx)
	if err != nil {
		err = client.Token.Update().SetVal(val).Where(token.HasOwnerWith(user.ID(u.ID))).Exec(ctx)
		if err != nil {
			return "", err
		}
	}
	return val, nil
}

func Logout(ctx context.Context, u *ent.User) error {
	_, err := client.Token.Delete().Where(token.HasOwnerWith(user.ID(u.ID))).Exec(ctx)
	return err
}

func UpdateUser(ctx context.Context, u *ent.User) error {
	_, err := client.User.UpdateOneID(u.ID).
		SetAvatar(u.Avatar).
		SetAge(u.Age).SetPhone(u.Phone).SetAddr(u.Addr).SetEmail(u.Email).
		SetAvatar(u.Avatar).SetCareer(u.Career).
		Save(ctx)
	return err
}
