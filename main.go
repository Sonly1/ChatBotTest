package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Message struct {
    Text string `json:"text"`
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var msg Message
        err := json.NewDecoder(r.Body).Decode(&msg)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        response := Message{Text: "You said: " + msg.Text}
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    } else {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/chat", chatHandler)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })
    fmt.Println("Starting server at port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}