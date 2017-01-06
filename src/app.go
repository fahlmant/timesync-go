package main


//Import needed packages
import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    //Create http handler
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
    //Bind to 8080
    log.Fatal(http.ListenAndServe(":8080", nil))
}
