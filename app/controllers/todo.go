package controllers
 
import (
    "encoding/json"
    "fmt"
    "net/http"

	"github.com/gorilla/mux"
	
	"github.com/Efrat19/go-boilerplate-server/app/models"
)
 
func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}
 
func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := models.Todos{
        models.Todo{Name: "Write presentation"},
        models.Todo{Name: "Host meetup"},
    }
 
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}
 
func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}