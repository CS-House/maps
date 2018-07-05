package main

import (
        "net/http"
        "github.com/rs/cors"
)

func main() {

        c := cors.AllowAll()

        mux := http.NewServeMux()

        mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                http.ServeFile(w, r, "googlemaps.html")
        })

        handler := c.Handler(mux)

        http.ListenAndServe(":3000", handler)
}
