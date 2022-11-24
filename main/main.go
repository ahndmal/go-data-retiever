package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type MnfChannel struct {
	Title           string
	Subscribers     int
	Views           int
	Videos          int
	VideosPlusViews int
}

func main() {
	response, err := http.Get("https://manifest.in.ua/rt/it/")
	if err != nil {
		log.Println(err)
	}

	body := response.Body
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Println("Error when creating document from reader", err)
	}
	defer body.Close()

	chanels := make([]MnfChannel, 0)
	trs := doc.Find("table.rating-table").
		Find("tr")
	for _, tr := range trs.Nodes {
		fmt.Println(tr.FirstChild.Data)
	}
	//doc.Find("a.rating-table__channel-name").Each(func(i int, linkA *goquery.Selection) {
	//	println(linkA.Text())
	//})
	log.Println(chanels)

}
