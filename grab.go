package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
	collystorage "github.com/gocolly/colly/storage"
)

func main() {
	fmt.Println("hello")

	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	cookieString := "..."

	var cookies []*http.Cookie
	for _, rawCookie := range strings.Split(cookieString, "; ") {
		parsedCookie := collystorage.UnstringifyCookies(rawCookie)
		cookies = append(cookies, parsedCookie[0])
	}

	c.SetCookies("https://www.facebook.com", cookies)

	c.Visit("https://www.facebook.com")
}

// func main() {
// 	c := colly.NewCollector()
//
// 	// Find and visit all links
// 	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 		e.Request.Visit(e.Attr("href"))
// 	})
//
// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL)
// 	})
//
// 	c.Visit("http://go-colly.org/")
// }
