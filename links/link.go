package links

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string `json:"href,omitempty"`
	Text string `json:"text,omitempty"`
}

var links []Link

func Parse(htmlFile io.Reader) ([]Link, error) {

	doc, err := html.Parse(htmlFile)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
		}
		break
	}
	if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
		s := strings.TrimSpace(n.FirstChild.Data)
		ret.Text = s
	}
	return ret
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var retNodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		retNodes = append(retNodes, linkNodes(c)...)
	}
	return retNodes
}
