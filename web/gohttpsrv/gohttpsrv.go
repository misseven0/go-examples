package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}
	// io.WriteString(w, "hello, world!\n")
	j, _ := json.Marshal(&data)
	w.Write(j)
}
func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
