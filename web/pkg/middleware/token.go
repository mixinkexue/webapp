package middleware

import (
	"context"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"web/db"
	"web/ent"
	"web/ent/token"
	"web/ent/user"
)

var ErrLogin = errors.New("未登录")

func Auth(c *gin.Context) {
	session := sessions.Default(c)
	val, ok := session.Get("token").(string)
	if !ok || val == "" {
		c.Error(ErrLogin)
		c.Abort()
		return
	}
	u, err := QueryToken(val)
	if err != nil {
		log.Println(err)
		c.Error(ErrLogin)
		c.Abort()
		return
	}
	c.Set("user", u)
}

func QueryToken(val string) (*ent.User, error) {
	return db.Client.User.Query().Where(user.HasTokenWith(token.Val(val))).Only(context.Background())
}
