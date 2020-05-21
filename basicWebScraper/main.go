package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	result, err  := GetLatestBlogTitles("https://golangcode.com")

	if(err != nil) {
		panic(err)
	}

	fmt.Println("Titles : ");
	fmt.Printf(result);
}

func GetLatestBlogTitles(url string) (string, error){
	resp, err := http.Get(url);

	if(err != nil) {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if(err != nil) {
		return "", err
	}

	titles := ""
	doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
		j := fmt.Sprintf("%d", i)
		titles += j + " - " + s.Text() + "\n" 
	})

	return titles, nil	
}
