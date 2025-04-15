package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "¡Hola desde Go! 🚀")
    })

    fmt.Println("Servidor ejecutándose en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}