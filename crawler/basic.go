package main

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gocolly/colly"
	"os"
	"regexp"
	"strings"
)

func main() {
	url := "https://books.studygolang.com/gopl-zh/"
	c := colly.NewCollector(
		colly.MaxDepth(2),
	)

	converter := md.NewConverter("", true, nil)

	/*
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})
	 */

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	c.OnResponse(func (r *colly.Response) {
		fmt.Println("Visited: ", r.Request.URL)
		if r.Request.Depth == 2 {
			sub := strings.Split(r.Request.URL.Path, "/")
			localPath := "./" + sub[1] + "/" + sub[2]
			fileName := strings.Split(sub[3], ".")[0]

			if _, err := os.Stat(localPath); os.IsNotExist(err) {
				os.MkdirAll(localPath, os.ModePerm)
			}

			f, err := os.OpenFile(localPath + "/" + fileName + ".md", os.O_CREATE|os.O_WRONLY, 0644)
			//f, err := os.OpenFile("." + r.Request.URL.Path, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			markdown, err := converter.ConvertBytes(r.Body)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			//_, err = f.Write(r.Body)
			_, err = f.Write(markdown)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			f.Close()
		}
	})
	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		match, _ := regexp.MatchString("ch[0-9]*/ch[0-9]*-[0-9]*.html", e.Attr("href"))
		if match {
			e.Request.Visit(e.Attr("href"))
		}
	})

	c.Visit(url)
}

