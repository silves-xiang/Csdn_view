package main

import (
	allnum "collytest/AllArticle"
	"flag"
	"fmt"
	"github.com/gocolly/colly"
	"math"
	"strconv"
	"time"
)

type Desc struct {
	ID int ` gorm:"primary_key;auto_increment"`
	Title string
	ViewNum string
}
func main() {
	var num int
	var url string
	flag.IntVar(&num,"num",100,"输入你每篇博客增加多少访问量")
	flag.StringVar(&url,"url","https://blog.csdn.net/Xiang_lhh","输入你要刷的博客列表，例如https://blog.csdn.net/Xiang_lhh")
	flag.Parse()
	Csdn_views(num,url)
}
func Csdn_views(num int , url string){
	// Instantiate default collector
	c := colly.NewCollector(//初始化colly
	)
	c.AllowURLRevisit=true//允许重复访问链接
	c.OnHTML("div.navList-box", func(e *colly.HTMLElement) {//回调函数，查找每篇文章的子链接,新版本
		//fmt.Println(*e)
		e.ForEach("article.blog-list-box", func(i int, element *colly.HTMLElement) {
			//遍历每个article标签
			http_articleid:=element.ChildAttr("a","href")//得到标签属性
			err:=c.Visit(http_articleid)//递归访问子链接
			if err!=nil{
				fmt.Println("访问子链接出现错误",err)
			}
			time.Sleep(time.Second)//间隔一秒
		})
	})
	c.OnHTML("main.content-mian-box", func(e *colly.HTMLElement) {//回调函数，查找每篇文章的子链接，旧版本
		e.ForEach("div.article-item-box", func(i int, element *colly.HTMLElement) {
			//遍历每个article标签
			http_articleid:=element.ChildAttr("a","href")//得到标签属性
			err:=c.Visit(http_articleid)//递归访问子链接
			if err!=nil{
				fmt.Println("访问子链接出现错误",err)
			}
			time.Sleep(time.Second)//间隔一秒
		})
	})
	c.OnHTML("div.article-header-box", func(e *colly.HTMLElement) {//自动匹配每篇文章的html
		dom:=e.DOM//返回DOM对象
		title:=dom.Find("h1.title-article").Text()//找到文章标题
		view_num:=dom.Find("span.read-count").Text()//找到每篇文章的访问量
		//注意colly为递归调用，不会重复刷新文章列表的页面，如果从文章列表中获取访问量，则访问量不会改变
		fmt.Println("访问成功","标题：",title,"阅读量：",view_num)
	})
	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("错误",err,response)//如果出错，进行输出
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {//访问之前
		fmt.Println("Visiting", r.URL.String())
	})
	total:=1//访问次数
	allnum.All(url)
	pages:=math.Ceil(float64(allnum.AllCsdn)/40)
	for total<=num{
		fmt.Println("第",total,"次刷博客")

		var err error
		if allnum.IsNew==1 {
			err = c.Visit(url)
		}
		if allnum.IsNew==2{
			for i:=1;i<=int(pages);i++{
				err=c.Visit(url+"/article/list/"+strconv.Itoa(i))//老版本分页
			}
		}
		if err !=nil{
			fmt.Println("出现错误",err)
		}
		fmt.Println("等待三十秒钟自动开启...")
		time.Sleep(time.Second*30)//每刷全部博客一次，自动间隔三十秒
		total++
	}
}