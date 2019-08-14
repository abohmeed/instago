package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("meta", func(e *colly.HTMLElement) {
		tag := e.Attr("property")
		if tag == "og:image" {
			open.Run(e.Attr("content"))
		}
	})
	if len(os.Args) == 1 {
		fmt.Println("Please enter the Instagram image URL")
		os.Exit(1)
	}
	url := os.Args[1]
	c.Visit(url)
}
