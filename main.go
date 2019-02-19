package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"net/http"
	"time"
)

var currentIP string
var lastUpdateTime time.Time


func SetIP(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["ip"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	key := keys[0]

	// Do this once for each unique policy, and use the policy for the life of the program
	// Policy creation/editing is not safe to use in multiple goroutines
	p := bluemonday.StrictPolicy()

	// Sanitize input to prevent cross site scripting attacks
	currentIP = p.Sanitize(string(key))
	log.Println("Url Param 'ip' is: " + currentIP)


	lastUpdateTime = time.Now()

	fmt.Fprintf(w,"Got new ip %s at time %s\n", currentIP, lastUpdateTime.String())


}


func GetIP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Current IP is %s\nLast update time is %s\n", currentIP, lastUpdateTime.String())
}

func main() {
	http.HandleFunc("/set", SetIP)
	http.HandleFunc("/", GetIP)
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
