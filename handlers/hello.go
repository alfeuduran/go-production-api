package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServerHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello Alfeu")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadGateway)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)

}
