package allnum

import (
	"github.com/gocolly/colly"
	"strconv"
)
var AllCsdn int
var IsNew int

func All(url string){
	c:=colly.NewCollector()
	c.OnHTML("div.container", func(e *colly.HTMLElement) {//老版本
		allcsdn,_:=strconv.Atoi(e.DOM.Find("div.data-info").Find("dl").First().Find("dt").Text())
		AllCsdn=allcsdn
		IsNew=2
	})
	c.OnHTML("div.user-profile-head-info-b", func(e *colly.HTMLElement) {//新版本
		AllCsdnString:=e.DOM.Find("li").First().Next().Find("div.user-profile-statistics-num").Text()
		AllCsdn,_=strconv.Atoi(AllCsdnString)
		IsNew=1
	})
	c.Visit(url)
}
