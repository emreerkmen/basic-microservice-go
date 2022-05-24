package main

import (
	"basic-microservice/hello/hello"
	"basic-microservice/hello/product-api/handlers_gorilla"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/gorilla/mux"
)

//var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {
	fmt.Println("Hello, World!")
	hello.Hello()

	//We create a http handler to use in http listen and serve
	//It register a function to a path
	//You can read more when hover on HandleFunc 😀
	//If you do curl -v /bilibilipath it will run that handler because /bilibilipath most similar to /
	//when there is no HandleFunc for /bilibilipath
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello from first go server!!!")
		d, err := ioutil.ReadAll(r.Body)
		//log.Printf("Data %s\n", d);
		if err != nil {
			http.Error(rw, "Oooops!", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s\n", d)
	})

	http.HandleFunc("/test", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello from test path!!!")
	})

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	productsHandler := handlers.NewProducts(l)

	router := mux.NewRouter();
	subGetRouter := router.Methods(http.MethodGet).Subrouter();
	subGetRouter.HandleFunc("/", productsHandler.GetProducts)

	subPutRouter := router.Methods(http.MethodPut).Subrouter();
	//Gorilla automaticly understand to use regex when see curly brackets
	subPutRouter.HandleFunc("/{id:[0-9]+}", productsHandler.UpdateProducts)

	// in video idle timeout info is important. Until that timeount is finished, the connection remains open
	// and do not need to hand shake again
	// we can tune that values for requirements
	// create a new server
	server := http.Server{
		Addr:         ":9090",      // configure the bind address
		Handler:      router,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// It actually creates a server
	// we can bind an IP address instead of just port,
	// but right now we are binding all of IP addresses because of absence of prefix port number
	//http.ListenAndServe(":9090", nil)
	//http.ListenAndServe(":9090", serveMux)
	//server.ListenAndServe()

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	signals := <-signalChannel
	l.Println("Recieved terminate, graceful shutdown", signals)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// Before shutdown It close open connections bıla bıla....
	server.Shutdown(timeoutContext)
}
