package main

import (
	"fmt"

	"./data"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func mains() {
	db := gormConnect2()
	defer db.Close()
	db.LogMode(true)

	//db.CreateTable(&data.Member{})
	db.CreateTable(&data.Article{})
	//db.CreateTable(&data.Race{})
	//db.CreateTable(&data.RaceCard{})
	// db.CreateTable(&data.Administrator{})
	//db.CreateTable(&data.Point{})
	race := data.Race{}
	db.First(&race, 2)
	fmt.Println(race.Title)
	var racecards []data.RaceCard
	db.Model(&race).Related(&racecards, "RaceRefer")
	for _, b := range racecards {
		member := data.Member{}
		if db.Where("no = ?", b.No+12).First(&member).RecordNotFound() {
			fmt.Println("no member")
		} else {
			fmt.Println(member.Name)
		}
	}

}

func gormConnect2() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=hinatazaka password=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	return db
}
