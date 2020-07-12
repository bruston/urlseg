package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		u := s.Text()
		urls := paths(u)
		for _, v := range urls {
			fmt.Println(v)
		}
	}
}

func paths(u string) []string {
	purl, err := url.Parse(u)
	if err != nil {
		return nil
	}
	u = strings.TrimPrefix(u, purl.Scheme+"://")
	split := strings.Split(purl.RequestURI(), "/")
	if len(split) < 2 {
		return []string{u}
	}
	urls := make([]string, 0, len(split)+1)
	urls = append(urls, purl.Scheme+"://"+u)
	newURL := purl.Scheme + "://" + purl.Host
	for i := range split {
		if i == len(split)-1 {
			continue
		}
		if strings.HasSuffix(newURL, "/") {
			newURL += split[i]
		} else {
			newURL += "/" + split[i]
		}
		urls = append(urls, newURL)
	}
	return urls
}
