package backend

import (
    "log"
    "runtime"
    "./server"
    "./data/mongodb"
)

func Boot(listening_url string, db_url string) {
    log.Println("Booting backend...")

    cpus := runtime.NumCPU()
    if runtime.GOMAXPROCS(cpus) >= 1 {
        log.Println("Using", cpus, "core/s")
    } else {
        log.Println("WARNING: i cannot use", cpus, "core/s")
    }

    log.Println("Server startup...")

    db := mongodb.New()
    s := server.New(db)

    db.Start(db_url)
    s.Start(listening_url)
}
