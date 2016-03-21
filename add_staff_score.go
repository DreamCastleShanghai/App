package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
	"time"
)

const (
	SessionSurveyID          = 0
	DemoJamVoteID            = 1
	UploadPhotoID            = 2
	UploadAvatarID           = 3
	SustainabilityCampaignID = 4
	StafforAmbassadorID      = 5
	SpeakerOfOwnSessionID    = 6
)

type SpeakerSessionRelation struct {
	SessionId int    `gorm:"column:SessionId"`
	SpeakerId int    `gorm:"column:SpeakerId"`
	Role      string `gorm:"column:Role"`
}

type ScoreHistory struct {
	//	UserId      int    `gorm:"column:UserId"`
	ScoreType   int    `gorm:"column:ScoreType"`
	Score       int    `gorm:"column:Score"`
	ScoreDetail string `gorm:"column:ScoreDetail"`
}

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

type TempUser struct {
	UserId    int    `gorm:"column:UserId;sql:"AUTO_INCREMENT"`
	LoginName string `gorm:"column:LoginName"`
	PassWord  string `gorm:"column:PassWord"`
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

	users := []TempUser{}

	if gDB != nil {
		gDB.Raw("select * from sap.tempuser a left join sap.User b on a.password = b.password").Scan(&users)
		for _, user := range users {
			//sidInt := strconv.Itoa(relation.SessionId)
			AddUserScore(user.UserId, StafforAmbassadorID, "Staff Credits")
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

// **********************************************************************************************************************
// **********************************************************************************************************************
//
//			add user score
//
// **********************************************************************************************************************
// **********************************************************************************************************************

func AddUserScore(userid int, scoretype int, detail string) (addscore int) {
	var addScore int = 0
	switch scoretype {
	case SessionSurveyID:
		addScore = 20
	case DemoJamVoteID:
		addScore = 35
	case UploadPhotoID:
		addScore = 2
	case UploadAvatarID:
		addScore = 5
	case SustainabilityCampaignID:
		addScore = 5
	case StafforAmbassadorID:
		addScore = 80
	case SpeakerOfOwnSessionID:
		addScore = 20
	}
	var canAdd bool = true
	if gDB != nil {
		scoreHistory := []ScoreHistory{}
		if scoretype == UploadAvatarID {
			gDB.Raw("SELECT * FROM Score_History WHERE ScoreType = ? AND UserId = ?", scoretype, userid).Scan(&scoreHistory)
			if len(scoreHistory) > 0 {
				canAdd = false
			}
		} else if scoretype == UploadPhotoID {
			gDB.Raw("SELECT * FROM Score_History WHERE ScoreType = ? AND UserId = ?", scoretype, userid).Scan(&scoreHistory)
			if len(scoreHistory) >= 6 {
				canAdd = false
			}
		}
		/* else if scoretype == DemoJamVoteID {
			gDB.Raw("SELECT * FROM Score_History WHERE ScoreType = ? AND UserId = ?", scoretype, userid).Scan(&scoreHistory)
			if len(scoreHistory) >= 6 {
				canAdd = false
			}
		} else if scoretype == SustainabilityCampaignID {
			gDB.Raw("SELECT * FROM Score_History WHERE ScoreType = ? AND UserId = ?", scoretype, userid).Scan(&scoreHistory)
			if len(scoreHistory) >= 6 {
				canAdd = false
			}
		}
		*/
		if canAdd {
			gDB.Exec("UPDATE User SET Score = Score + ?, SubTime = ? WHERE UserId = ?", addScore, time.Now().Unix(), userid)
			gDB.Exec("INSERT INTO Score_History (UserId, ScoreType, Score, ScoreDetail) VALUES (?, ?, ?, ?)", userid, scoretype, addScore, detail)
			MyPrint("Add score succeed !")
		} else {
			MyPrint("Add score failed !")
		}
	}
	return addScore
}
