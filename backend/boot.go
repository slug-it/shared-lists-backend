package backend

import (
    "log"
    "./server"
    "./data/mongodb"
)

func Boot(listening_url string, db_url string) {
    log.Println("Booting backend...")

    db := mongodb.New()
    s := server.New()

    db.Start(db_url)
    s.Start(listening_url)
}
