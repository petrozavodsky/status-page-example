package main

import (
	"context"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"status_page/internal/out"
	"status_page/internal/server"
	"syscall"
	"time"
)

func main() {

	out.GetResultData()
	router := mux.NewRouter()
	handler := server.NewHandler()
	handler.Register(router)

	app := &http.Server{Addr: "localhost:8585", Handler: router}

	go func() {

		log.Info("Старт приложения")

		if err := app.ListenAndServe(); err != nil {
			log.Fatal("не удалось запустить сервер: ", err)
		}

	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := app.Shutdown(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
