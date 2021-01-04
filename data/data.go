package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Member struct {
	gorm.Model
	No          uint
	Name        string  `gorm:"type:varchar(100);unique_index"`
	Kana        string  `gorm:"type:varchar(100)"`
	Birthday    string  `gorm:"type:varchar(20)"`
	Sign        string  `gorm:"type:varchar(10)"`
	Height      float64 `gorm:"type:numeric"`
	Birthplace  string  `gorm:"type:varchar(20)"`
	Blood       string  `gorm:"type:varchar(10)"`
	GraduatedAt string  `gorm:"type:varchar(20)"`
}

type Article struct {
	gorm.Model
	No          string  `gorm:"type:varchar(10)"`
	PostingDate string  `gorm:"type:varchar(20)"`
	Member      string  `gorm:"type:varchar(100);index"`
	Title       string  `gorm:"type:text"`
	Body        string  `gorm:"type:text"`
	Images      string  `gorm:"type:text"`
	ImageCount  float64 `gorm:"type:numeric"`
}

type Race struct {
	gorm.Model
	Title       string     `gorm:"type:varchar(100)"`
	Description string     `gorm:"type:text"`
	CountFrom   string     `gorm:"type:varchar(20)"`
	CountTo     string     `gorm:"type:varchar(20)"`
	VoteFrom    string     `gorm:"type:varchar(20)"`
	VoteTo      string     `gorm:"type:varchar(20)"`
	RaceCards   []RaceCard `gorm:"foreignkey:UserRefer"`
}

type RaceCard struct {
	gorm.Model
	RaceRefer uint
	No        int `gorm:"type:integer"`
	MemberNo  int
	Odds      int
}
type Administrator struct {
	gorm.Model
	Code     string `json:"code" gorm:"varchar(20);unique;not null"`
	Password string `json:"password" gorm:"type:text;not null"`
	Note     string `json:"note" gorm:"type:text"`
}
type Account struct {
	gorm.Model
	Code     string `json:"code" gorm:"varchar(100);unique;not null"`
	Password string `json:"password" gorm:"type:text"`
	Note     string `json:"note" gorm:"type:text"`
}
type Point struct {
	gorm.Model
	Quantity     string `json:"quantity" gorm:"type:numeric"`
	AccountRefer string `json:"accountRefer" gorm:"type:integer"`
	ExpiryDate   string `json:"note" gorm:"type:text"`
	RaceRefer    uint   `json:"raceRefer" gorm:"type:integer"`
	Action       string `json:"action" gorm:"type:text"`
	Type         string `json:"type" gorm:"type:text"`
}
