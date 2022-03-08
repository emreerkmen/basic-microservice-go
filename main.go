package main

import (
	"basic-microservice/hello/hello"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
    fmt.Println("Hello, World!")
	hello.Hello()

	//We create a http handler to use in http listen and serve
	//It register a function to a path
	//You can read more when hover on HandleFunc 😀
	//If you do curl -v /bilibilipath it will run that handler because /bilibilipath most similar to /
	//when there is no HandleFunc for /bilibilipath 
	http.HandleFunc("/",func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello from first go server!!!")
		d, err := ioutil.ReadAll(r.Body)
		//log.Printf("Data %s\n", d);
		if err != nil {
			http.Error(rw, "Oooops!", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s\n", d)
	})

	http.HandleFunc("/test",func(http.ResponseWriter, *http.Request) {
		log.Println("Hello from test path!!!")
	})

	// It actually creates a server
	// we can bind an IP address instead of just port, 
	// but right now we are binding all of IP addresses because of absence of prefix port number
	http.ListenAndServe(":9090", nil)
}