package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"./data"
	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func main1() {
	db := gormConnect()
	defer db.Close()
	var start int = 1
	var qty int = 25
	for i := start; i < start+qty; i++ {
		var no string = strconv.Itoa(i)
		t := time.Now()
		fmt.Println(no + ":" + t.Format("2006/01/02 15:04:05"))
		url := "https://www.hinatazaka46.com/s/official/artist/" + no + "?ima=0000"

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
		name := doc.Find("div.c-member__name--info").Text()
		kana := doc.Find("div.c-member__kana").Text()
		if name == "" {
			continue
		}
		member := data.Member{}
		member.No = uint(i)
		member.Name = strings.TrimSpace(name)
		member.Kana = strings.TrimSpace(kana)
		doc.Find("table.p-member__info-table > tbody > tr> td.c-member__info-td__text").Each(func(index int, s *goquery.Selection) {
			switch index {
			case 0:
				member.Birthday = strings.TrimSpace(s.Text())
			case 1:
				member.Sign = strings.TrimSpace(s.Text())
			case 2:
				var str string = strings.TrimSpace(s.Text())
				str = strings.Replace(str, "cm", "", -1)
				converted, _ := strconv.ParseFloat(str, 64)
				member.Height = converted
			case 3:
				member.Birthplace = strings.TrimSpace(s.Text())
			case 4:
				member.Blood = strings.TrimSpace(s.Text())
			}
		})
		db.Create(&member)
		time.Sleep(time.Second)
	}
}

func gormConnect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=hinatazaka password=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	return db
}
