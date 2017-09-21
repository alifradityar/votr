package main

import (
	"fmt"
	"net/http"

	"github.com/alifradityar/votr/votr-api/config"
	"github.com/alifradityar/votr/votr-api/handler"
	"github.com/alifradityar/votr/votr-api/router"
	"github.com/facebookgo/inject"
	redis "github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

func main() {
	conf := config.Get()

	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", conf.MysqlUsername, conf.MysqlPassword, conf.MysqlHost, conf.MysqlDatabase))
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxOpenConns(conf.MysqlConnectionLimit)
	redistClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.RedisHost, conf.RedisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// Setup dependency injection
	var rh handler.Root
	err = inject.Populate(db, redistClient, &rh)
	if err != nil {
		fmt.Println(err)
	}

	// Setup router
	r := router.CreateRouter(rh)

	// Serve
	fmt.Println("Votr started in Port: " + conf.Port)
	err = http.ListenAndServe(":"+conf.Port, r)
	if err != nil {
		fmt.Println(err)
	}
}
