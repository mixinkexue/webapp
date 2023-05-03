package db

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"webapp/ent"
)

var Client *ent.Client

func init() {
	var err error
	Client, err = ent.Open("mysql", "root:root@tcp(localhost:3306)/adopt?parseTime=True")

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	Client = Client.Debug()
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
