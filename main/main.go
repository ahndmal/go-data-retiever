package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	response, err := http.Get("https://manifest.in.ua/rt/it/")
	if err != nil {
		fmt.Println(err)
		log.Println(err)
	}
	body := response.Body
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Println("Error when creating document from reader", err)
	}
	defer body.Close()
	doc.Find("table.rating-table").Children().Find("tr").Each(func(i int, tr *goquery.Selection) {
		println(tr.Text())
	})

}
