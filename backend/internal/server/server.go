package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var server http.Server
var router = mux.NewRouter()

func SetValues(host string, port string) {
	server.Addr = fmt.Sprintf("%s:%s", host, port)
}

func StopServer(x int8) {
	if x == 0 {
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Println("Error in shutting down server!")
		}
		fmt.Println("")
		log.Println("------------------Shutting down server-----------------------\n")
		return
	}
	timer := time.NewTimer(time.Duration(x) * time.Minute)
	fmt.Println("----------------------Shutting Down server in", x, "min---------------------")
	<-timer.C
	err := server.Shutdown(context.Background())
	if err != nil {
		log.Fatal("Error shutting down server!")
	}
}

func StartServer() {
	log.Println("---------------------Starting server---------------------")

	server.Handler = router

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Fatal(server.ListenAndServe())

	}()

	<-stop

	StopServer(0)
}

func init() {
	Router()
}
