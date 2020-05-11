package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "1 + 2, example of golang.org"

	x := url.QueryEscape(urlString)

	fmt.Println("Before : ", urlString)
	fmt.Println("After : ", x)

	//Program to make longer query

	params := url.Values{}

	params.Add("q", "1 + 2")
	params.Add("p", "example of golang.org")

	output := params.Encode()

	//fmt.Println("Before : ", params)
	fmt.Println("After : ", output)
}