package util 

import 
(
	"net/http"
	"log"
)

/*
  List of http response status can be found at:

  https://golang.org/src/net/http/status.go
*/

func SendBadReqStatus(w http.ResponseWriter, err error) {
    log.Print(err)
    //Http status code: 400
    http.Error(w, err.Error(), http.StatusBadRequest) 
}

func SendUnauthorizedStatus(w http.ResponseWriter) {
    //Http status code: 401
    http.Error(w, "Unauthorized Access", http.StatusUnauthorized) 
}

func SendStatusNotFound(w http.ResponseWriter, err error) {
    log.Print(err)
    //Http status code: 404
    http.Error(w, err.Error(), http.StatusNotFound) 
}

func SendInternalServerErrorStatus(w http.ResponseWriter, err error) {
    log.Print(err)
    //Http status code: 500
    http.Error(w, err.Error(), http.StatusInternalServerError) 
}