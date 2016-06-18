package main

import (
    "fmt"
    "net/http"
    "time"
    "strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi, the time is %s", time.Now())
    fmt.Println("Received connection.")
}

func getPort() string {
    for {
        fmt.Print("Enter port number to create server (1024 to 65535): ")
        var resp string
        _, err := fmt.Scanln(&resp)
        if err != nil {
            continue
        }
        num, err := strconv.ParseInt(resp, 10, 32)
        if err != nil {
            continue
        }
        if num >= 1024 && num <= 65535 {
            return ":" + resp
        }
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(getPort(), nil)
}

