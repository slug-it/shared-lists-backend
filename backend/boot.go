package backend

import (
    "log"
    "./server"
)

func Boot(listening_url string, db_url string) {
    log.Println("Booting backend...")

    s := server.New()
    s.Start(listening_url)
}
