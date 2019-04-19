package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/logger"
)

var db database
var log = logger.Init("logger", true, true, ioutil.Discard)

func main() {
	log.Infoln("Starting program")
	log.Infoln("Initializing database")
	db = newDB()
	db.init()
	//db.test()
	fmt.Println(db.userlist())
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/list", userlistHandler)
	http.ListenAndServe("0.0.0.0:8080", nil)
	db.close()
}
