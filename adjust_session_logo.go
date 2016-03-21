package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
	//"time"
)

type UserView struct {
	UserId    int    `gorm:"column:UserId;sql:"AUTO_INCREMENT"`
	LoginName string `gorm:"column:LoginName"`
	FirstName string `gorm:"column:FirstName"`
	LastName  string `gorm:"column:LastName"`
	Icon      string `gorm:"column:Icon"`
	Score     int    `gorm:"column:Score"`
	//	Authority	int		`gorm:"column:Authority"`
	DemoJamId1   int    `gorm:"column:DemoJamId1"`
	DemoJamId2   int    `gorm:"column:DemoJamId2"`
	VoiceVoteId1 int    `gorm:"column:VoiceVoteId1"`
	VoiceVoteId2 int    `gorm:"column:VoiceVoteId2"`
	EggVoted     bool   `gorm:"column:EggVoted"`
	GreenAmb     bool   `gorm:"column:GreenAmb"`
	SubTime      int64  `gorm:"column:SubTime"`
	DeviceToken  string `gorm:"column:DeviceToken"`
}

type Session struct {
	SessionId int `gorm:"column:SessionId;sql:"AUTO_INCREMENT"`
	//	SpeakerId	int 	`gorm:"column:SpeakerId"`
	Title       string `gorm:"column:Title"`
	Format      string `gorm:"column:Format"`
	Track       string `gorm:"column:Track"`
	Location    string `gorm:"column:Location"`
	StartTime   int64  `gorm:"column:StartTime"`
	EndTime     int64  `gorm:"column:EndTime"`
	Description string `gorm:"column:Description"`
	Point       int    `gorm:"column:Point"`
	Logo        string `gorm:"column:Logo"`
}

var gRelease bool = true
var gLocal bool = false
var gDB *gorm.DB

func main() {
	argCnt := len(os.Args)

	var messageId int = 0

	for i := 1; i < argCnt; i++ {
		if os.Args[i] == "debug" {
			gRelease = false
		} else if os.Args[i] == "local" {
			gLocal = true
		} else {
			messageId, _ = strconv.Atoi(os.Args[i])
			MyPrint("Message id : ", messageId)
		}
	}

	fmt.Println("Release Mode : ", gRelease)

	gDB = ConnectDB(gRelease)

	//sessions := []Session{}
	if gDB != nil {
		gDB.Exec("UPDATE Session SET Logo = 'session/s4h.png' WHERE track = 'S/4HANA'")
		gDB.Exec("UPDATE Session SET Logo = 'session/han.png' WHERE track = 'HANA, HCP & Cloud Infrastructure'")
		gDB.Exec("UPDATE Session SET Logo = 'session/sme.png' WHERE track = 'Small & Medium Enterprises'")
		gDB.Exec("UPDATE Session SET Logo = 'session/lob.png' WHERE track = 'LoB Applications'")
		gDB.Exec("UPDATE Session SET Logo = 'session/iot.png' WHERE track = 'Internet of Things'")
		gDB.Exec("UPDATE Session SET Logo = 'session/ue.png' WHERE track = 'User Experience'")
		gDB.Exec("UPDATE Session SET Logo = 'session/ana.png' WHERE track = 'Analytics'")
		gDB.Exec("UPDATE Session SET Logo = 'session/see.png' WHERE track = 'Security & Engineering Excellence'")
		gDB.Exec("UPDATE Session SET Logo = 'session/ia.png' WHERE track = 'Industry Applications'")
		//gDB.Exec("UPDATE Session SET Logo = 'session/ana.png' WHERE track = 'Analytics'")
		/*
			gDB.Raw("select * from Session").Scan(&sessions)

			for _, session := range sessions {
				if session.Track == "S/4HANA"
			}
		*/
	}

	gDB.Close()
}

func ConnectDB(isRelease bool) *gorm.DB {
	MyPrint("start to connecting db!")
	var connectStr string
	if gLocal {
		MyPrint("Local : ")
		connectStr = "root@tcp(127.0.0.1:3306)/SAP?charset=utf8&parseTime=True"
	} else {
		MyPrint("Remote : ")
		connectStr = "root:1011@/SAP?charset=utf8&parseTime=True"
	}
	db, err := gorm.Open("mysql", connectStr)
	if CheckErr(err) {
		return nil
	}
	fmt.Println("start to connecting db finished!")

	fmt.Println("start to init db!")
	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	if isRelease {
		db.LogMode(false)
	} else {
		db.LogMode(true)
	}
	db.SingularTable(true)
	//db.AutoMigrate(&User{}, &Tests{})
	fmt.Println("start to init db finished!")

	return &db
}

func MyPrint(a ...interface{}) {
	if gRelease {
		return
	} else {
		fmt.Println(a)
	}
}

func CheckErr(err error) bool {
	if err != nil {
		panic(err)
		return true
	}
	return false
}
