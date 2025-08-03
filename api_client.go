package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func fetchExample() {
	resp, err := http.Get("placeholder") // update this!

	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("API Response:", string(body))
}
