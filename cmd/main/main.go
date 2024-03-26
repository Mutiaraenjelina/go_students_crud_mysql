/* NIM : 11422008
   Nama: Mutiara Enjelina
   Prodi: STr TRPL
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Mutiaraenjelina/go_students_crud_mysql/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
    r := mux.NewRouter()
    routes.RegisterStudentsRoutes(r)
    http.Handle("/", r)
    fmt.Print("Starting Server on localhost:9010")
    log.Fatal(http.ListenAndServe("localhost:9010", r))
}
