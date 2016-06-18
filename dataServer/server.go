/*
 * File: server.go
 * ------------------------------
 * Implements a simple server. 
 */

package main

import (
    "fmt"
    "net/http"
    "time"
    "strconv"
    "math/rand"
    // "io/ioutil"
    "os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // simple time server that responds to request with 
    // the current time, prints message to server console 
    // to acknowledge request
    fmt.Println("Received connection.")
    fmt.Fprintf(w, "Hi, the time is %s\n", time.Now())
    fmt.Fprintf(w, "  Here's your random integer: " + strconv.Itoa(rand.Intn(100)) + "\n")
    fortune, _ := exec.Command("fortune", "").Output()
    fmt.Fprintf(w, "  Here's your fortune: \n" + string(fortune))

}

func getPort() string {
    // Prompts user to specify a port to host server on
    // from 1024 (lowest port that doesn't require root 
    // access) and 65535 (highest port number)
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
    // installs handler function for all requests 
    // listens on user-defined port and serves requests 
    // using the installed handler
    http.HandleFunc("/", handler)
    http.ListenAndServe(getPort(), nil)
}

