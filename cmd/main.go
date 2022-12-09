package main

import (
	"github.com/sirupsen/logrus"
	"mk/internal/repositories"
	"mk/internal/services"
	"mk/internal/web"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	db, err := repositories.NewDb()
	if err != nil {
		logrus.Fatal(err)
	}

	repo := repositories.NewRepository(db)
	service := services.NewService(repo)
	router := web.NewHandler(service)
	serv := new(Server)

	if err := serv.Run("8080", router.InitRoutes()); err != nil {
		logrus.Fatal(err.Error())
	}
}
