package server

import (
    "log"
    "net/http"
    "../data"
)

type server struct {
    db data.Data
}

func New(db data.Data) *server {
    if g_server != nil {
        log.Fatal("Trying to create another server!")
    }

    g_server = &server{db}
    return g_server
}

func handler(writer http.ResponseWriter, request *http.Request) {
    log.Println(request.RemoteAddr, " is asking for ", request.RequestURI)
}

func (self *server) Start(listening_url string) {
    log.Println("Listening on: ", listening_url)

    http.HandleFunc("/", handler)
    err := http.ListenAndServe(listening_url, nil)
    if err != nil {
        log.Fatal("Listening error: ", err)
    }
}

var g_server *server
