package old_main

import (
	"basic-microservice/hello/hello"
	"basic-microservice/hello/product-api/handlers"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func old_main() {
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

	/*--------Video 2--------*/

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)
	productsHandler := handlers.NewProducts(l)

	serveMux := http.NewServeMux()
	serveMux.Handle("/hello", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)
	serveMux.Handle("/", productsHandler)

	// in video idle timeout info is important. Until that timeount is finished, the connection remains open
	// and do not need to hand shake again
	// we can tune that values for requirements
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
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
