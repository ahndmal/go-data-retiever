package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
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
		fmt.Println(err)
		log.Println(err)
	}
	body := response.Body
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Println("Error when creating document from reader", err)
	}
	defer body.Close()
	chanels := make([]MnfChannel, 0)
	doc.Find("table.rating-table").
		Find("tr").Find("td").Each(func(i int, td *goquery.Selection) {
		chn := MnfChannel{}
		switch i {
		case 1:
			chn.Title = td.Text()
		case 3:
			atoi, err := strconv.Atoi(td.Text())
			if err != nil {
				log.Println(err)
			}
			chn.Subscribers = atoi
		case 4:
			chn.Views = i
		case 5:
			chn.Videos = i
		case 6:
			chn.VideosPlusViews = i
		}
		chanels = append(chanels, chn)
	})
	//doc.Find("a.rating-table__channel-name").Each(func(i int, linkA *goquery.Selection) {
	//	println(linkA.Text())
	//})
	log.Println(chanels)

}
