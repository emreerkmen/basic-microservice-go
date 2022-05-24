package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

// When write test, we don't have to think about how we can inject mock log object.
// We can use that beautiful constructır for this.
// We try to do kind of Dependency Injection.
func NewGoodbye(l *log.Logger) *Goodbye{
	return &Goodbye{l}
}

// ServeHTTP func. signature is exactly same with handler func. So we implement ServeHTTP interface
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Hello from first go server!!!")
	d, err := ioutil.ReadAll(r.Body)
	//log.Printf("Data %s\n", d);
	if err != nil {
		http.Error(rw, "Oooops!", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Goodbye %s\n", d)
}