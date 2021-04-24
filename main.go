package main

import (
    "os"
    "log"
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("views")))
    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}
