package data

import (
    "log"
    "labix.org/v2/mgo"
)

type db_connection struct {
}

func New() *db_connection {
    return &db_connection{}
}

func (self *db_connection) Start(url string) {
    log.Println("Connecting to db at", url)

    session, err := mgo.Dial(url)
    if err != nil {
        log.Fatal("Connection error: ", url)
    }

    db := session.DB("listespesa")
    log.Println(db.Name)
}
