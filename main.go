package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)
type Desc struct {
	ID int ` gorm:"primary_key;auto_increment"`
	Title string
	ViewNum string
}
func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	)
	//colly.DisallowedDomains("https://greasyfork.org/zh-CN","https://goproxy.cn/")
	c.AllowURLRevisit=true
	// On every a element which has href attribute call callback
	c.OnHTML("div.navList-box", func(e *colly.HTMLElement) {
		//fmt.Println(e.DOM.Find("span.read-count").Text())
		e.ForEach("article.blog-list-box", func(i int, element *colly.HTMLElement) {
			http_articleid:=element.ChildAttr("a","href")
			c.Visit(http_articleid)
			time.Sleep(time.Second)
		})
	})
	c.OnHTML("div.article-header-box", func(e *colly.HTMLElement) {
		dom:=e.DOM
		title:=dom.Find("h1.title-article").Text()
		view_num:=dom.Find("span.read-count").Text()
		fmt.Println("访问成功","标题：",title,"阅读量：",view_num)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("错误",err,response)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	i:=1
	for {
		fmt.Println("第",i,"次刷博客")
		fmt.Println("等待三十秒钟自动开启...")
		time.Sleep(time.Second*30)
		err:=c.Visit("https://blog.csdn.net/Xiang_lhh/")
		if err !=nil{
			fmt.Println("出现错误",err)
		}
		i++
		if i==100{
			break
		}
	}
}
