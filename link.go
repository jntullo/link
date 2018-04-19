package link

import (
	"golang.org/x/net/html"
	"fmt"
	"errors"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func Traverse(n *html.Node) []Link {
	var links []Link
	links = traverse(n, links)
	return links
}

func traverse(n *html.Node, l []Link) []Link {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		link := parseLink(c)
		if link.Href != "" {
			l = append(l, link)
		}
		l = traverse(c, l)
	}
	return l
}

func parseLink(n *html.Node) Link{
	var l Link
	if isAnchorElement(n) {
		href, err := getAttribute(n, "href")
		if err != nil {
			fmt.Println(err)
		}
		text := strings.TrimSpace(n.FirstChild.Data)
		if href != "" {
			l = Link{href, text}
			return l
		}
	}
	return l
}

func isAnchorElement(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func getAttribute(n *html.Node, key string) (string, error) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, nil
		}
	}
	return "", errors.New(key + " does not exist in attribute!")
}