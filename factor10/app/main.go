package main

import (
    "fmt"
    "log"
    "net/http"
    "math/rand"
	"time"
)

func main() {

	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	randValue := r1.Intn(100)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%d\n", randValue)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}