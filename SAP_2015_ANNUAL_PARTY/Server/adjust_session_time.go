package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
	"time"
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
	RealTime    string `gorm:"column:RealTime"`
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

	MyPrint("Release Mode : ", gRelease)

	gDB = ConnectDB(gRelease)

	sessions := []Session{}

	if gDB != nil {

		gDB.Raw("select * from Session").Scan(&sessions)

		loc, _ := time.LoadLocation("Asia/Shanghai")
		var realtime string
		var starttime string
		var endtime string
		var startunix int64
		var endunix int64

		for _, session := range sessions {
			realtime = session.RealTime
			MyPrint("loc : ", loc)
			MyPrint("real time : " + realtime)
			starttime = Substr(realtime, 0, 5)
			endtime = Substr(realtime, 8, 5)
			MyPrint("start time : " + starttime)
			MyPrint("end time : " + endtime)

			starttime = "2016-02-29 " + Substr(realtime, 0, 5) + ":00"
			endtime = "2016-02-29 " + Substr(realtime, 8, 5) + ":00"
			MyPrint("start time : " + starttime)
			MyPrint("end time : " + endtime)

			startParse, _ := time.ParseInLocation("2006-01-02 15:04:05", starttime, loc)
			endParse, _ := time.ParseInLocation("2006-01-02 15:04:05", endtime, loc)
			MyPrint("startParse : ", startParse)

			startunix = startParse.Unix()
			endunix = endParse.Unix()

			//MyPrint(realtime)
			//MyPrint(starttime)
			MyPrint(startunix)
			//MyPrint(endtime)
			MyPrint(endunix)
			gDB.Exec("UPDATE Session SET StartTime = ?, EndTime = ? WHERE SessionId = ?", startunix, endunix, session.SessionId)
		}
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

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
