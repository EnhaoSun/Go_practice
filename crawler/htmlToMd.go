package main

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gocolly/colly"
	"os"
)

func main() {
	url := os.Args[1]
	fileName := os.Args[2]
	c := colly.NewCollector()

	converter := md.NewConverter("", true, nil)

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited: ", r.Request.URL)
		markdown, err := converter.ConvertBytes(r.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
		_, err = f.Write(markdown)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		f.Close()
	})
	c.Visit(url)
}