package main

import "C"
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//export Get
func Get(i int) string {
	URL := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i)
	fmt.Printf("Request to %s\n\n", URL)
	resp, err := http.Get(URL)
	if err != nil {
		return "failure"
	}
	data := map[string]interface{}{}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "failure"
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return "failure"
	}
	fmt.Println("Response:\n")
	for k, v := range data {
		fmt.Printf("%s => %v\n", k, v)
	}
	return "success"
}

func main() {
}
