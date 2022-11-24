package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
)

const FilePathPrefix = "/Users/czy0538/Documents/0A学习资料/GO/gopl.io/ch8/exec8.7/data/"
const HostURL = "https://czy0538.github.io/"

func saveFile(buf *io.Reader, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	io.Copy(f, *buf)
	return nil
}

func exactURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	// 解析resp.Body为[]bytes类型
	htmlBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read html body error:%v", err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				if linkBytesRaw := []byte(link.String()); bytes.HasPrefix(linkBytesRaw, []byte(HostURL)) {
					links = append(links, link.String())
					linkLocal := bytes.ReplaceAll(linkBytesRaw, []byte(HostURL), []byte(FilePathPrefix))
					fmt.Println(linkLocal)
					htmlBody = bytes.ReplaceAll(htmlBody, linkBytesRaw, linkLocal)
				}

			}
		}
	}
	forEachNode(doc, visitNode, nil)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
func main() {
	err := exactURL(HostURL)
	if err != nil {
		log.Fatalln(err)
	}
}
