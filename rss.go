package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urltofeed(url string) (RSSFeed, error) {
	httpclient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpclient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, err
	}
	rssfeed := RSSFeed{}
	err = xml.Unmarshal(dat, &rssfeed)
	if err != nil {
		return RSSFeed{}, err
	}
	return rssfeed, nil
}
