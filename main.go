package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s QUERY\n", os.Args[0])
		os.Exit(1)
	}

	resp, err := http.Get(fmt.Sprintf(
		"http://dictionary.goo.ne.jp/srch/jn/%s/m6u/",
		strings.Replace(url.QueryEscape(os.Args[1]), "+", "%20", -1),
	))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	html, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	candidates := html.Find(".list-search-a > li")
	if candidates.Length() >= 1 {
		candidates.Each(func(i int, s *goquery.Selection) {
			fmt.Println(s.Find(".title").Text())
			fmt.Printf(
				"\033[37m%s\033[0m\n",
				strings.TrimSpace(strings.Replace(s.Find(".mean").Text(), "\t", "\n", -1)),
			)
			fmt.Println()
		})
	} else {
		explanation := html.Find(".explanation").Text()
		fmt.Println(strings.TrimSpace(strings.Replace(explanation, "\t", "", -1)))
	}
}
