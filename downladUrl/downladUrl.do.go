package downladUrl

import (
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func Generator() []string {
	//https://wallpaperstock.net/wallpapers_p2.html
	var s string = "https://wallpaperstock.net/wallpapers.html"
	var allUrl []string
	var workers int = 5
	var mutex sync.Mutex
	ch := make(chan []string)
	for i := 2; i < workers+2; i++ {
		mutex.Lock()
		var sTemp = s[:37] + "_p" + strconv.Itoa(i) + s[37:]
		go findLinks(sTemp, ch)
		mutex.Unlock()
		allUrl = append(allUrl, <-ch...)
	}
	return allUrl
}

func findLinks(url string, c chan []string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	c <- visit(nil, doc)
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" && strings.Contains(a.Val, "wallpapers/thumbs") {
				links = append(links, "https:"+a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}
