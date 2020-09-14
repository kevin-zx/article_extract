package article_extract

import (
	"fmt"
	http_util "github.com/kevin-zx/http-util"
	"testing"
)

func TestCalculate(t *testing.T) {
	html, err := http_util.GetWebConFromUrl("https://news.163.com/20/0821/04/FKHDIVVD0001899O.html")
	if err != nil {
		panic(err)
	}
	calculate(html)
}

func TestExtractArticle(t *testing.T) {
	html, err := http_util.GetWebConFromUrl("http://www.cdcxsj.com/xinwendongtai/pinpaixuanchuan/209.html")
	//html, err := http_util.GetWebConFromUrl("http://www.cdcxsj.com/xinwendongtai/gongsixinwen/246.html")
	//html,err := http_util.GetWebConFromUrl("https://news.163.com/20/0821/04/FKHDIVVD0001899O.html")
	if err != nil {
		panic(err)
	}
	article, err := ExtractArticle(html)
	if err != nil {
		panic(err)
	}
	fmt.Printf("title: %s\n", article.Title)
	fmt.Printf("score: %.4f\n", article.Score)
	fmt.Printf("summary: %s\n", article.Summary)
	fmt.Printf("content: %s\n", article.ContentText)
	fmt.Printf("contentHtml: %s\n", article.ContentHTML)
}

func Test_removeSuccessiveLink(t *testing.T) {
	html, err := http_util.GetWebConFromUrl("http://www.cdcxsj.com/xinwendongtai/gongsixinwen/246.html")
	//html,err := http_util.GetWebConFromUrl("https://news.163.com/20/0821/04/FKHDIVVD0001899O.html")
	if err != nil {
		panic(err)
	}
	//doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	//if err != nil {
	//	panic(err)
	//}
	doc, err := removeScriptAndStyle(html)
	if err != nil {
		panic(err)
	}
	content := doc.Find(".icon1")
	content = removeSuccessiveLink(content)
	fmt.Println(content.Text())
}
