package main

import (
	"my-driver/db"
	"my-driver/handler"
	"my-driver/repository/repo_impl"
	"my-driver/router"

	"github.com/labstack/echo/v4"
)

func main() {
	sql := &db.Sql{
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


