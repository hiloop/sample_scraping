package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"./data"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

type Images struct {
	Images []string `json:"images"`
}

func main() {
	db := connect()
	defer db.Close()
	// 19984
	//
	var start int = 36240
	for i := start; i < 36250; i++ {
		var no string = strconv.Itoa(i)
		t := time.Now()
		fmt.Println(no + ":" + t.Format("2006/01/02 15:04:05"))
		url := "https://www.hinatazaka46.com/s/official/diary/detail/" + no + "?ima=0000&cd=member"

		// Getリクエスト
		res, _ := http.Get(url)
		defer res.Body.Close()

		// 読み取り
		buf, _ := ioutil.ReadAll(res.Body)

		// 文字コード判定
		det := chardet.NewTextDetector()
		detRslt, _ := det.DetectBest(buf)

		// 文字コード変換
		bReader := bytes.NewReader(buf)
		reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

		// HTMLパース
		doc, _ := goquery.NewDocumentFromReader(reader)

		title := doc.Find("div.c-blog-article__title").Text()
		posting_date := doc.Find("div.c-blog-article__date > time").Text()
		name := doc.Find("div.c-blog-article__name > a").Text()
		body := doc.Find("div.c-blog-article__text").Text()
		var slice []string
		imgs := doc.Find("div.c-blog-article__text img")
		imgs.Each(func(i int, s *goquery.Selection) {
			src, _ := s.Attr("src")
			slice = append(slice, src)
		})
		res1D := &Images{Images: slice}
		res1B, _ := json.Marshal(res1D)
		articleEx := data.Article{}
		articleEx.No = no
		articleEx.PostingDate = posting_date
		articleEx.Title = strings.TrimSpace(title)
		articleEx.Member = name
		articleEx.Body = strings.TrimSpace(body)
		articleEx.Images = string(res1B)
		articleEx.ImageCount = float64(imgs.Length())
		db.Create(&articleEx)
		time.Sleep(time.Millisecond * 200)
	}
}

func connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=hinatazaka password=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	return db
}
