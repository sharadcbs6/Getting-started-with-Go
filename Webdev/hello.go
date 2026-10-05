package main
import (
	"fmt"
	"net/http"
)
func main(){
	http.HandleFunc("/ad", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Hello World</h1>");
    })
	http.Handle("/",http.StripPrefix("/home",http.FileServer(http.Dir("static/"))))
    http.ListenAndServe(":8080", nil)
}