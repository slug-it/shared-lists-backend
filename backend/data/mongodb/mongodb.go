package mongodb

import (
    "log"
    "labix.org/v2/mgo"
)

type mongodb struct {
}

func New() *mongodb {
    return &mongodb{}
}

func (self *mongodb) Start(url string) {
    log.Println("Connecting to db at", url)

    session, err := mgo.Dial(url)
    if err != nil {
        log.Fatal("Connection error: ", url)
    }

    db := session.DB("listespesa")
    log.Println(db.Name)
}
