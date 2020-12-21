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

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Printf("hello world!")
	d, err := ioutil.ReadAll(r.Body)
	// log.Printf("data is %s\n", d)
	if err != nil {
		http.Error(rw, "Oooops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}
