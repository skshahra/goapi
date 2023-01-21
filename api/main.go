package main

import (
	// "fmt"

	"fmt"
	"net/http"
	"os"
	"time"

	"bitbucket.org/skshahriarahmed/sh_ra/config"
	"bitbucket.org/skshahriarahmed/sh_ra/logs"
	"bitbucket.org/skshahriarahmed/sh_ra/route"

	"github.com/gorilla/mux"
	"github.com/pytimer/mux-logrus"
)


func init (){
	// LoadEnvVar()
	config.LoadEnvVar()
}
func main(){
	
	r := mux.NewRouter()
	route.Router(r)

	HttpServer := &http.Server{
		Handler: r,
		Addr:  os.Getenv("SERVER_IP"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	r.Use(muxlogrus.NewLogger().Middleware)
	fmt.Println("✨ Server is running on ",os.Getenv("SERVER_IP"))
	if err := HttpServer.ListenAndServe(); err != nil {
		logs.ERROR("Error in HttpServer.ListenAndServe() ",err)
		panic(err)
	}

}


