package myProject

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    site, err := ioutil.ReadFile("build/index.html")

    if err != nil {
        panic(err)
    }

    fmt.Fprint(w, string(site))
}