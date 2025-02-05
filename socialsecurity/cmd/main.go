package main

import (
	"database/sql"
	"log"
	"socialsecurity/internal/app/application"
	"socialsecurity/internal/app/user"
	"socialsecurity/internal/config"
	v1 "socialsecurity/internal/handlers/v1"
	mssqlApplication "socialsecurity/internal/mssql/application"
	mssqlUser "socialsecurity/internal/mssql/user"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {

	cfg, err := config.Load("../config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mssql", cfg.GetDSN())
	if err != nil {
		log.Fatal(err)
	}

	// ctx, cancel := context.WithCancel(context.Background())

	// defer cancel()

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := mssqlUser.NewUserRepository(db)

	userService := user.NewService(userRepo)

	applicationRepo := mssqlApplication.NewApplicationRepository(db)

	applicationService := application.NewService(applicationRepo)

	svc := v1.NewAllSercivces(userService, applicationService)

	handler := v1.NewHandler(svc)

	r := handler.InitRoutes()

	r.Run(":3000")
	// user, err := userService.AddUser(ctx, types.CreateUserRequest{
	// 	Name:  "John Doe",
	// 	Email: "john.doe@example.com",
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(user)
}
