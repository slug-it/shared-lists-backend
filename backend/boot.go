package backend

import (
    "log"
    "./server"
    "./data"
)

func Boot(listening_url string, db_url string) {
    log.Println("Booting backend...")

    db := data.New()
    s := server.New()

    db.Start(db_url)
    s.Start(listening_url)
}
