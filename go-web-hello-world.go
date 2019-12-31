package main

import (
    "fmt"
    "net/http"
    "flag"
    "strconv"
    "log"
)

func main() {
    var p int
    flag.IntVar(&p, "port", 80, "Need a postive number less than 65535")
    flag.Parse()

    if p <= 0 || p >= 65535 {
        p = 80
    }
    ps := strconv.Itoa(p)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Go Web Hello World!")
    })
    log.Print("Listening at :" + ps)
    log.Fatal(http.ListenAndServe(":" + ps, nil))
}
