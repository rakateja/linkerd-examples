package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"net/http"
)

func main() {
	log.Println("Starting web-api...")
	realMain()
}

func realMain() {
	http.HandleFunc("/me", func(w http.ResponseWriter, req *http.Request) {
		b, err := json.Marshal(map[string]interface{}{
			"msg": "FOOBAR",
			"src": "foobar-api",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})
	log.Printf("web-api is started. Listening to %s", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil)
}
