package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting Application ... ")
	log.Println("I love Niorukomi")

	var url = "https://www.google.com"
	
	client := http.Client{}
	request := http.NewRequest("GET", url, )


	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(fmt.Errorf("error making get request %v", err))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("error reading the body of the response %v", err))
	}

	err = os.WriteFile("index.html", body, os.ModePerm)
	if err != nil {
		panic(err) 
	}
}
