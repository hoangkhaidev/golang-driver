package main

import (
	"fmt"
	"my-driver/db"
	"my-driver/handler"
	"my-driver/repository/repo_impl"
	"my-driver/router"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("=====>", os.Getenv("APP_NAME"))

	sql := &db.Sql{
		// Host: "host.docker.internal",
		Host: "localhost",
		Port: 5432,
		UserName: "postgres",
		Password: "1234",
		DbName: "golang",
	}
	sql.Connect()
	defer sql.Close()

	server := echo.New()

	UserHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}

	api := router.API{
		Echo:		 server,
		UserHandler: UserHandler,
	}
	api.SetupRouter()

	
	server.Logger.Fatal(server.Start(":4000"))
}


