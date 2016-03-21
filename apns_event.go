package main

import (
	"fmt"
	//"github.com/bitly/go-simplejson"
	//"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/virushuo/Go-Apns"
	"os"
	"strconv"
	"time"
)

type UserView struct {
	UserId    string `gorm:"column:UserId"`
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

var gRelease bool = true
var gLocal bool = false
var gDB *gorm.DB

const (
	RootResDir       = "./res/"
	WebResDir        = "/res"
	VersionResDir    = "./versions/release/"
	WebVersionResDir = "/apk"
	IconFileName     = "icon"
	TimeFormat       = "2006-01-02 15:04:05"
	VOTE_NO_READY    = 0
	VOTE_START       = 1
	VOTE_FINISHED    = 2
	NOTICE_FAVORITE  = 0
	NOTICE_EVENT     = 1
	NOTICE_PRIZE     = 2
)

var notificationTitle = []string{
	"Build your agenda",
	"Check-in",
	"Sustainability campaign",
	"Keynote",
	"d-kom survey",
	"DemoJam",
	"Evening party",
	"Important",
}

var notificationContent = []string{
	"Build your agenda to get the most out of our sessions & earn as many credits as possible!",
	"The early bird get the worm… Check in before 9:15 AM tomorrow and you will have a health band!",
	"Read & Agree “I am a Sustainability Guardian”, and get 5 credits.",
	"The keynote will start in 15 min.  Please join us at Grand Ballroom II&III on 7th Floor",
	"Take the d-kom Survey & Win a Kindle!",
	"The Demo Jam will start in 15 min.  Please join us at Pearl Hall on 7th Floor",
	"The evening party will start in 15 min.  Please join us at Grand Ballroom II&III on 7th Floor",
	"All testing data (including credits) have been cleared on 06:30 2nd March.\nThanks for your kindly support and understanding.\nPS: Here recommend you to download latest version (IOS from apple store and  android from QR code).\nBest Regards,",
}

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

	users := []UserView{}
	if gDB != nil {
		gDB.Raw("select * from User").Scan(&users)

		apn, err := apns.New("prod.pem", "key-noenc.pem", "gateway.push.apple.com:2195", 1*time.Second)
		//	apn, err := apns.New("prod.pem", "key-noenc.pem", "gateway.sandbox.push.apple.com:2195", 1*time.Second)
		if err != nil {
			fmt.Printf("connect error: %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Println("connect successed!")
		go readError(apn.ErrorChan)

		for _, user := range users {
			if user.DeviceToken != "" {
				notification(apn, user.DeviceToken, NOTICE_EVENT, time.Now().Unix(), notificationTitle[messageId], notificationContent[messageId])
			}
			gDB.Exec("INSERT INTO Message (UserId, MessageTitle, MessageDetail, MessageTime, MessageType) VALUES (?, ?, ?, ?, 2)", user.UserId, notificationTitle[messageId], notificationContent[messageId], time.Now().Unix())
		}

		apn.Close()
	}

	gDB.Close()
}

func notification(apn *apns.Apn, token string, tp int, id int64, title string, body string) {
	//token := "a1e909eb31f244fccafe4bcb252ed5e3d1d87d2e0a4d962f9e8946046a8d354e"
	MyPrint("%d, %s, %d, %d, %s, %s", apn, token, tp, id, title, body)
	payload := apns.Payload{}
	payload.Aps.Alert.Body = body //"Congratulations!\nYou won a sport camera in the raffle!\nPlease go to the right side of the stage after the party to claim your prize or contact Ms. Karen Zhao at 18800349005."
	payload.Aps.Sound = "bingbong.aiff"
	payload.SetCustom("id", id) //time.Now().Unix())
	payload.SetCustom("tp", tp) //0)
	payload.SetCustom("title", title)

	//{"id":"12345678","tp":0,"aps":{"alert":{"body":"Message content"}}}

	//js, err := simplejson.NewJson([]byte(`{}`))
	//js.Set("aps", "alert")
	//	js.Set("aps", "badge")
	//	js.Set("badge", 2)
	//	js.Set("alert", "body")
	//	js.Set("alert", "action-loc-key")
	//body, _ := js.String()
	//fmt.Println(string(js))

	//body, _ := js.String()
	//payload.Aps.Alert.Body = body

	notification := apns.Notification{}
	notification.DeviceToken = token
	notification.Identifier = 1
	notification.Payload = &payload
	err := apn.Send(&notification)
	MyPrint("send id(%x): %s\n", notification.Identifier, err)
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

/*
func main() {
	apn, err := apns.New("prod.pem", "key-noenc.pem", "gateway.push.apple.com:2195", 1*time.Second)
	//	apn, err := apns.New("prod.pem", "key-noenc.pem", "gateway.sandbox.push.apple.com:2195", 1*time.Second)
	if err != nil {
		fmt.Printf("connect error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("connect successed!")
	go readError(apn.ErrorChan)
	token := "a1e909eb31f244fccafe4bcb252ed5e3d1d87d2e0a4d962f9e8946046a8d354e"

	payload := apns.Payload{}
	payload.Aps.Alert.Body = "Congratulations!\nYou won a sport camera in the raffle!\nPlease go to the right side of the stage after the party to claim your prize or contact Ms. Karen Zhao at 18800349005."
	payload.Aps.Sound = "bingbong.aiff"
	payload.SetCustom("id", time.Now().Unix())
	payload.SetCustom("tp", 2)
	payload.SetCustom("title", "test")

	//{"id":"12345678","tp":0,"aps":{"alert":{"body":"Message content"}}}

	//js, err := simplejson.NewJson([]byte(`{}`))
	//js.Set("aps", "alert")
	//	js.Set("aps", "badge")
	//	js.Set("badge", 2)
	//	js.Set("alert", "body")
	//	js.Set("alert", "action-loc-key")
	//body, _ := js.String()
	//fmt.Println(string(js))

	//body, _ := js.String()
	//payload.Aps.Alert.Body = body

	notification := apns.Notification{}
	notification.DeviceToken = token
	notification.Identifier = 0
	notification.Payload = &payload
	err = apn.Send(&notification)
	fmt.Printf("send id(%x): %s\n", notification.Identifier, err)

	apn.Close()
}
*/

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

func readError(errorChan <-chan error) {
	for {
		apnerror := <-errorChan
		fmt.Println(apnerror.Error())
	}
}
