package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/go-resty/resty/v2"
)

func main() {
	log.Println("Starting web-api...")
	realMain()
}

func realMain() {
	http.HandleFunc("/me", func(w http.ResponseWriter, req *http.Request) {
		client := resty.New()
		resp, err := client.R().
			EnableTrace().
			Get(fmt.Sprintf("%s/foobar", os.Getenv("FOOBAR_API_HOST")))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var foobarMsg struct {
			Msg string `json:"msg"`
		}
		err = json.Unmarshal(resp.Body(), &foobarMsg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		b, err := json.Marshal(map[string]interface{}{
			"msg": foobarMsg.Msg,
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
