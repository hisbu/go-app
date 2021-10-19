package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"syscall"
	"time"
)

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/health", health)
	http.HandleFunc("/print", handlePrint)
	http.HandleFunc("/date", handleDate)
	http.HandleFunc("/shell", handleShell)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func int8ToStr(arr []int8) string {
	b := make([]byte, 0, len(arr))
	for _, v := range arr {
		if v == 0x00 {
			break
		}
		b = append(b, byte(v))
	}
	return string(b)
}

func handleShell(w http.ResponseWriter, r *http.Request) {
	var uname syscall.Utsname
	if err := syscall.Uname(&uname); err == nil {
		var ver = int8ToStr(uname.Release[:])
		w.Write([]byte(ver))

	}

}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "welcome"
	w.Write([]byte(message))
}

func health(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "server is currently healthy\n")
}

func handlePrint(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		payload := struct {
			Name   string `json:"name"`
			Age    int    `json:"age"`
			Gender string `json:"gender"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResp, err := json.Marshal(payload)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}

	http.Error(w, "Only accept POST request", http.StatusBadRequest)
}

func handleDate(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	fmt.Fprintf(w, "Current time is :", currentTime.Format("2006-01-02 15:04:05"), "\n")
	fmt.Println("hallo", currentTime.String())
}
