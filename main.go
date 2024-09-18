package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", RandomString())
	})

	c.OnHTML("main#main", func(main *colly.HTMLElement) {
		main.ForEach("article[id^='post']", func(i int, article *colly.HTMLElement) {
			/*
				获取图片信息
			*/
			figure := article.DOM.Find("figure.post-thumbnail")
			a := figure.Find("a[href]")
			imgPath, exist := a.Find("img[data-lazy-src]").Attr("data-lazy-src")
			if exist {
				//如果存在图片路径，则认为抓取正常
				fmt.Println(imgPath)
			}

			/*
				获取文字信息
			*/
			header := article.DOM.Find("header.entry-header")
			nameA := header.Find("h2.entry-title").Find("a[href]")
			name := nameA.Text()
			detailUrl, exist := nameA.Attr("href")
			fmt.Println("name:", name, ", detailUrl:", detailUrl)

			div := header.Find("div.entry-meta")
			div.Find("time.entry-date.published")

		})
	})

	c.Visit("https://bgirlbboy.com/category/bboys/")
}
