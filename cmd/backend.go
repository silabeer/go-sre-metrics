package main

import (
	"context"
	"restapi/internal/handler"
	"restapi/internal/server"

	"github.com/sirupsen/logrus"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetLevel(logrus.DebugLevel)
	srv := new(server.Server)
	logrus.Info("PORT:", "8082")
	if err := srv.Run("8082", handler.InitRoutes()); err != nil {
		logrus.Fatalf("error occurred while running http server: %s\n", err.Error())
	}
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s\n", err.Error())
	}
}
