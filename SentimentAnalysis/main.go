package main

import (
	"fmt"

	"github.com/cdipaolo/sentiment"
)

var (
	analysis *sentiment.Analysis
	models   sentiment.Models
	err      error
)

func main() {
	models, err = sentiment.Restore()

	if err != nil {
		panic(err)
	}

	PrintSentimentExpression("Your mother is an awful lady")

	PrintSentimentExpression("Holy place")

}

func PrintSentimentExpression(text string) {
	analysis = models.SentimentAnalysis(text, sentiment.English)

	if analysis.Score == 1 {
		fmt.Println("Positive Statement")
	} else {
		fmt.Println("Negative Statement")
	}
}
