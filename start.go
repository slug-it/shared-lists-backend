package main

import "./backend"

func main() {
    backend.Boot("localhost:8080", "localhost:9080")
}
