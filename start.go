package main

import "./backend"

func main() {
    backend.Boot("localhost:8080", "http://localhost:9080")
}
