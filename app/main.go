package main
 
import (
    "log"
	"net/http"
	"github.com/Efrat19/go-boilerplate-server/app/router"
)
 
func main() {
 
    router := router.NewRouter()
 
    log.Fatal(http.ListenAndServe(":8080", router))
}