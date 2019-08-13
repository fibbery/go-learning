package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"net/http"
	"os"
	"strings"
)

func ExampleScrape() {
	host := "https://www.75txt.org"
	catalog := "/5/5524/"
	res, err := http.Get(host + catalog)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".col-md-3").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		item := s.Find("a")
		title, _ := decodeToGBK(item.Text())
		href, exist := item.Attr("href")
		if exist {
			fmt.Printf("fetch Chapter[%s] : %s \n", title, href)
			content, _ := http.Get(host + href)
			contentdoc, _ := goquery.NewDocumentFromReader(content.Body)
			txt, _ := decodeToGBK(contentdoc.Find("#htmlContent").Text())
			file, _ := os.OpenFile(
				"/Users/fibbery/Desktop/book/"+title+".txt",
				os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
				0666,
			)
			file.Write([]byte(strings.Replace(txt, "聽聽聽聽", "", -1)))
			file.Close()
		}
	})
}

func decodeToGBK(text string) (string, error) {
	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewDecoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return text, err
	}

	return string(dst[:nDst]), nil
}

func main() {
	//ExampleScrape()
	// force colors on for TextFormatter
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, FullTimestamp: true,})
	//// then wrap the log output with it
	logrus.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	logrus.WithField("name", "jiangnenghua").Infof("hello this is %s", "world")
}
