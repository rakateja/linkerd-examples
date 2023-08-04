package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"net/http"
)

func main() {
	log.Println("Starting foobar-api...")
	realMain()
}

func realMain() {
	http.HandleFunc("/foobar", func(w http.ResponseWriter, req *http.Request) {
		b, err := json.Marshal(map[string]interface{}{
			"msg": "FOOBAR",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})
	log.Printf("foobar-api is started. Listening to %s", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
