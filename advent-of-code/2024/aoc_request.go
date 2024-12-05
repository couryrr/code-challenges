package main

import (
	"net/http"
	"fmt"
	"io"
	"os"
)

var cookie = http.Cookie{
	Name: "session", 
	Value: os.Getenv("AOC_SESSION"),
}

func GetDataForDay(url, file_name string) []byte{
	input := CheckForInputFile(file_name)

	if input == nil {
		input = RequestInput(url, file_name)
	}
	return input
}

func CheckForInputFile(file_name string) []byte{
	file, _ := os.ReadFile(file_name)
	return file	
}

func RequestInput(url string, file_name string) []byte{
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		panic("failed to create rquest")
	}
	
	req.AddCookie(&cookie)

	resp, err := client.Do(req)	
	if err != nil {
		fmt.Printf("%v\n", err)
		panic("failed to access day 1 input")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	
	if err != nil {
		fmt.Printf("%v\n", err)
		panic("network connection failed")
	}
	
	os.WriteFile(file_name, body, 0644)
	return body
}
