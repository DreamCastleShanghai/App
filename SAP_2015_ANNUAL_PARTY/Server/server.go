package main

import (
	//"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	//"net/http"
	"fmt"
	"math/rand"
	//"net/url"
	"io"
	"os"
	//"io/ioutil"
	"github.com/bitly/go-simplejson"
	_ "github.com/go-sql-driver/mysql"
	"path/filepath"
	"strconv"
	"time"
	//"encoding/json"
	//"./MyDBStructs"
	"github.com/itsjamie/gin-cors"
	"github.com/virushuo/Go-Apns"
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

const (
	RootResDir       = "./res/"
	WebResDir        = "/res"
	VersionResDir    = "./versions/release/"
	WebVersionResDir = "/apk"
	HtmlResDir       = "./html/"
	WebHtmlResDir    = "/bs"

	IconFileName = "icon"

	TimeFormat = "2006-01-02 15:04:05"

	VOTE_NO_READY = 0
	VOTE_START    = 1
	VOTE_FINISHED = 2

	SAP_VOICE_CNT = 3
	DEMOJAM_CNT   = 8

	PRIZE_4_CNT = 30
	PRIZE_5_CNT = 30
	PRIZE_6_CNT = 30
)

var gEggHikingCnt int = 0
var gSapVoiceCnt [SAP_VOICE_CNT]int
var gDemoJamCnt [DEMOJAM_CNT]int

var gDB *gorm.DB
var gRelease bool = true
var gLocal bool = false
var gDemoJamVoteStatus int = VOTE_NO_READY
var gSAPVoiceStatus int = VOTE_NO_READY
var gEggHikingStatus int = VOTE_NO_READY
var gCanGetScores bool = true

var sustainbilityContext string = "1.    I take public transportation and/or cycle or walk to d-kom Shanghai venue.\n\n2.    I save paper by using electronic onsite guide in d-kom app.\n\n3.    I finish off my meals and have “clean plate” today.\n\n4.    I drink bottled water and recycle plastic bottles to recycle bins, and/or used my own cup to drink.\n\n5.    I do not smoke today.\n\n6.    At d-kom, I support to use old laptops and furniture that were moved from Labs China Shanghai Campus.\n\n7.    I share pictures about sustainability on the “Moments” of d-kom Shanghai App"

// **********************************************************************************************************************
// **********************************************************************************************************************
//
//			Database Structures
//
// **********************************************************************************************************************
// **********************************************************************************************************************
type DemoJamItem struct {
	DemoJamItemId int    `gorm:"column:DemoJamItemId;sql:"AUTO_INCREMENT"`
	TeamName      string `gorm:"column:TeamName"`
	Resource      string `gorm:"column:Resource"`
	Department    string `gorm:"column:Department"`
	Introduction  string `gorm:"column:Introduction"`
}

type DemoJamVote struct {
	//	DemoJamVoteId int `gorm:"column:DemoJamVoteId;sql:"AUTO_INCREMENT"`
	UserId        int `gorm:"column:UserId"`
	DemoJamItemId int `gorm:"column:DemoJamItemId"`
}

type DkomSurveyResult struct {
	//SurveyId 	int 	`gorm:"column:SurveyId;sql:"AUTO_INCREMENT"`
	UserId int `gorm:"column:UserId"`
	Q1     int `gorm:"column:Q1"`
	Q2     int `gorm:"column:Q2"`
	Q3     int `gorm:"column:Q3"`
	Q4     int `gorm:"column:Q4"`
}

type EggHikingItem struct {
	EggHikingTitle  string `gorm:"column:EggHikingTitle"`
	EggHikingDetail string `gorm:"column:EggHikingDetail"`
	Resource        string `gorm:"column:Resource"`
	EggHikingBG     string `gorm:"column:EggHikingBG"`
}

type HikingVote struct {
	UserId   int  `gorm:"column:UserId"`
	VoteFlag bool `gorm:"column:VoteFlag"`
}

type PictureWall struct {
	//	PictureWallId 	int 	`gorm:"column:PictureWallId;sql:"AUTO_INCREMENT"`
	UserId   int    `gorm:"column:UserId"`
	Picture  string `gorm:"column:Picture"`
	Category string `gorm:"column:Category"`
	Comment  string `gorm:"column:Comment"`
	//IsDelete		bool	`gorm:"column:IsDelete"`
	//PostTime 		int64 	`gorm:"column:PostTime"`
}

type Message struct {
	MessageType   int    `gorm:"column:MessageType"`
	MessageDetail string `gorm:"column:MessageDetail"`
	MessageTitle  string `gorm:"column:MessageTitle"`
	MessageTime   int64  `gorm:"column:MessageTime"`
}

type ScoreHistory struct {
	//	UserId      int    `gorm:"column:UserId"`
	ScoreType   int    `gorm:"column:ScoreType"`
	Score       int    `gorm:"column:Score"`
	ScoreDetail string `gorm:"column:ScoreDetail"`
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

type SessionSurveyResult struct {
	//SurveyId 	int 	`gorm:"column:SurveyId;sql:"AUTO_INCREMENT"`
	SessionId int `gorm:"column:SessionId"`
	UserId    int `gorm:"column:UserId"`
	A1        int `gorm:"column:A1"`
	A2        int `gorm:"column:A2"`
	A3        int `gorm:"column:A3"`
}

type SpeakerSessionRelation struct {
	SessionId int    `gorm:"column:SessionId"`
	SpeakerId int    `gorm:"column:SpeakerId"`
	Role      string `gorm:"column:Role"`
}

type StaticRes struct {
	Resource string `gorm:"column:Resource"`
	//	ResType  string `gorm:"column:ResType"`
	ResLable string `gorm:"column:ResLable"`
}

type SurveyInfo struct {
	//SurveyInfoId	int 	`gorm:"column:SurveyId;sql:"AUTO_INCREMENT"`
	SessionId int    `gorm:"column:SessionId"`
	QContent1 string `gorm:"column:QContent1"`
	Q11       string `gorm:"column:Q11"`
	Q12       string `gorm:"column:Q12"`
	Q13       string `gorm:"column:Q13"`
	Q14       string `gorm:"column:Q14"`
	QContent2 string `gorm:"column:QContent2"`
	Q21       string `gorm:"column:Q21"`
	Q22       string `gorm:"column:Q22"`
	Q23       string `gorm:"column:Q23"`
	Q24       string `gorm:"column:Q24"`
	//	Q3			string 	`gorm:"column:Q3"`
}

type SurveyAnswerInfo struct {
	Answer1 int `gorm:"column:Answer1"`
	Answer2 int `gorm:"column:Answer2"`
}

type Tests struct {
	IdTests int `gorm:"column:id_tests; primary_key:yes"`
	Temp    int `gorm:"column:temp"`
}

type User struct {
	UserId    int    `gorm:"column:UserId;sql:"AUTO_INCREMENT"`
	LoginName string `gorm:"column:LoginName"`
	PassWord  string `gorm:"column:PassWord"`
	FirstName string `gorm:"column:FirstName"`
	LastName  string `gorm:"column:LastName"`
	Icon      string `gorm:"column:Icon"`
	//	Score		int		`gorm:"column:Score"`
	//	Authority	int		`gorm:"column:Authority"`
}

type UserView struct {
	LoginName string `gorm:"column:LoginName"`
	FirstName string `gorm:"column:FirstName"`
	LastName  string `gorm:"column:LastName"`
	Icon      string `gorm:"column:Icon"`
	Score     int    `gorm:"column:Score"`
	//	Authority	int		`gorm:"column:Authority"`
	DemoJamId1   int   `gorm:"column:DemoJamId1"`
	DemoJamId2   int   `gorm:"column:DemoJamId2"`
	VoiceVoteId1 int   `gorm:"column:VoiceVoteId1"`
	VoiceVoteId2 int   `gorm:"column:VoiceVoteId2"`
	EggVoted     bool  `gorm:"column:EggVoted"`
	GreenAmb     bool  `gorm:"column:GreenAmb"`
	SubTime      int64 `gorm:"column:SubTime"`
}

type UserPictureRelation struct {
	//	RelationId 		int 	`gorm:"column:RelationId"`
	UserId        int  `gorm:"column:UserId"`
	PictureWallId int  `gorm:"column:PictureWallId"`
	LikeFlag      bool `gorm:"column:LikeFlag"`
}

type UserSessionRelation struct {
	//	RelationId	int 		`gorm:"column:relationid"; primary_key:yes; sql:"AUTO_INCREMENT"`
	UserId         int  `gorm:"column:UserId"`
	SessionId      int  `gorm:"column:SessionId"`
	LikeFlag       bool `gorm:"column:LikeFlag"`
	CollectionFlag bool `gorm:"column:CollectionFlag"`
}

type VoiceItem struct {
	VoiceItemId int    `gorm:"column:VoiceItemId;sql:"AUTO_INCREMENT"`
	VoicerName  string `gorm:"column:VoicerName"`
	SongName    string `gorm:"column:SongName"`
	VoicerPic   string `gorm:"column:VoicerPic"`
	VoicerDes   string `gorm:"column:VoicerDes"`
	VoicerBG    string `gorm:"column:VoicerBG"`
}

type VoiceVote struct {
	//	VoiceVoteId	int 	`gorm:"column:VoiceVoteId;sql:"AUTO_INCREMENT"`
	UserId      int `gorm:"column:UserId"`
	VoiceItemId int `gorm:"column:VoiceItemId"`
}

type AllSessionView struct {
	SessionId int    `gorm:"column:SessionId"`
	Title     string `gorm:"column:Title"`
	Format    string `gorm:"column:Format"`
	Track     string `gorm:"column:Track"`
	Location  string `gorm:"column:Location"`
	StartTime int64  `gorm:"column:StartTime"`
	EndTime   int64  `gorm:"column:EndTime"`
	//	Description	string	`gorm:"column:Description"`
	Point          int    `gorm:"column:Point"`
	Logo           string `gorm:"column:Logo"`
	LikeFlag       bool   `gorm:"column:LikeFlag"`
	LikeCnt        int    `gorm:"column:LikeCnt"`
	CollectionFlag bool   `gorm:"column:CollectionFlag"`
	Done           bool
}

type TempSession struct {
	SessionId int `gorm:"column:SessionId"`
}

type Speaker struct {
	//	UserId		int		`gorm:"column:UserId;sql:"AUTO_INCREMENT"`
	//	LoginName	string	`gorm:"column:LoginName"`
	//	PassWord	string	`gorm:"column:PassWord"`
	FirstName string `gorm:"column:FirstName"`
	LastName  string `gorm:"column:LastName"`
	Title     string `gorm:"column:Title"`
	Icon      string `gorm:"column:Icon"`
	Role      string `gorm:"column:Role"`
	//	Score		int		`gorm:"column:Score"`
	//	Authority	int		`gorm:"column:Authority"`
}

type PictureWallListView struct {
	PictureWallId int `gorm:"column:PictureWallId;sql:"AUTO_INCREMENT"`
	//	UserId			int 	`gorm:"column:UserId"`
	Icon        string `gorm:"column:Icon"`
	Picture     string `gorm:"column:Picture"`
	Category    string `gorm:"column:Category"`
	FirstName   string `gorm:"column:FirstName"`
	LastName    string `gorm:"column:LastName"`
	Title       string `gorm:"column:Title"`
	Comment     string `gorm:"column:Comment"`
	LikeFlagCnt int    `gorm:"column:LikeFlagCnt"`
	IsLiked     bool
	//IsDelete		bool	`gorm:"column:IsDelete"`
	//PostTime 		int64 	`gorm:"column:PostTime"`
}

type PictureWallView struct {
	//	PictureWallId 	int 	`gorm:"column:PictureWallId;sql:"AUTO_INCREMENT"`
	//	UserId			int 	`gorm:"column:UserId"`
	Picture  string `gorm:"column:Picture"`
	Category string `gorm:"column:Category"`
	Comment  string `gorm:"column:Comment"`
	//IsDelete		bool	`gorm:"column:IsDelete"`
	//PostTime 		int64 	`gorm:"column:PostTime"`
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
		addScore = 90
	case SpeakerOfOwnSessionID:
		addScore = 20
	}
	var canAdd bool = true
	if !gCanGetScores {
		addscore = 0
		detail = "Time out : " + detail
		canAdd = false
	}
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

// **********************************************************************************************************************
// **********************************************************************************************************************
//
//			router's selection logic function
//
// **********************************************************************************************************************
// **********************************************************************************************************************
func RouterGetSAP(c *gin.Context) {
	MyPrint("sap get start!")
	msgType := c.Query("tag")
	MyPrint("tag is : ", msgType)
	switch msgType {
	case "WV":
		RouterGetSapVoiceWinner(c)
	case "W":
		RouterGetWinner(c)
	case "CE":
		RouterGetCntOfEggHiking(c)
	case "CD":
		RouterGetCntOfDemoJam(c)
	case "CS":
		RouterGetCntOfSapVoice(c)
	case "scoreswitch":
		RouterGetGetScoresSwitch(c)
	case "hiking":
		RouterGetEggHikingSwitch(c)
	case "djstatus":
		RouterGetDemoJamSwitch(c)
	case "svstatus":
		RouterGetSAPVoiceSwitch(c)
	case "T0":
		RouterGetToken(c)
	case "M0":
		RouterGetMessage(c)
	case "L0":
		RouterGetLogin(c)
	case "U0":
		RouterGetUser(c)
	case "UI0":
		RouterGetUserIcon(c)
	case "SL0":
		RouterGetSessionList(c)
	case "VE0":
		RouterGetVoiceEnter(c)
	case "VV0":
		RouterGetVoiceVote(c)
	case "VL0":
		RouterGetVoiceList(c)
	case "DE0":
		RouterGetDemoJamEnter(c)
	case "DV0":
		RouterGetDemoJamVote(c)
	case "DL0":
		RouterGetDemoJamList(c)
	case "SV0":
		RouterGetSessionVote(c)
	case "SC0":
		RouterGetSessionCollect(c)
	case "R0":
		RouterGetRank(c)
	case "PS0":
		RouterGetPictureSubmit(c)
	case "PD0":
		RouterGetPictureDelete(c)
	case "PV0":
		RouterGetPictureVote(c)
	case "PL0":
		RouterGetPictureList(c)
	case "SSI0":
		RouterGetSessionSurveyInfo(c)
	case "SSS0":
		RouterGetSubmitSessionSurvey(c)
	case "DSS0":
		RouterGetSubmitDKOMSurvey(c)
	case "SD0":
		RouterGetSessionDetail(c)
	case "PML0":
		RouterGetPictureMyList(c)
	case "DVL0":
		RouterGetDemoJamVoiceList(c)
	case "MSL0":
		RouterGetMyScoreList(c)
	case "SI0":
		RouterGetSustainbilityInfo(c)
	case "SR0":
		RouterGetSustainabilitySubmit(c)
	case "MP0":
		RouterGetMap(c)
	case "SH0":
		RouterGetScoreHistory(c)
	case "EE0":
		RouterGetEggHikingEnter(c)
	case "EV0":
		RouterGetHiking(c)
	case "I0":
		RouterGetInformation(c)
	}
	MyPrint("sap get finished!")
}

func RouterPostSAP(c *gin.Context) {
	MyPrint("sap post start!")
	msgType := c.PostForm("tag")
	MyPrint("tag is : ", msgType)
	switch msgType {
	case "T0":
		RouterPostToken(c)
	case "M0":
		RouterPostMessage(c)
	case "L0":
		RouterPostLogin(c)
	case "U0":
		RouterPostUser(c)
	case "UI0":
		RouterPostUserIcon(c)
	case "SL0":
		RouterPostSessionList(c)
	case "VE0":
		RouterPostVoiceEnter(c)
	case "VV0":
		RouterPostVoiceVote(c)
	case "VL0":
		RouterPostVoiceList(c)
	case "DE0":
		RouterPostDemoJamEnter(c)
	case "DV0":
		RouterPostDemoJamVote(c)
	case "DL0":
		RouterPostDemoJamList(c)
	case "SV0":
		RouterPostSessionVote(c)
	case "SC0":
		RouterPostSessionCollect(c)
	case "R0":
		RouterPostRank(c)
	case "PS0":
		RouterPostPictureSubmit(c)
	case "PD0":
		RouterPostPictureDelete(c)
	case "PV0":
		RouterPostPictureVote(c)
	case "PL0":
		RouterPostPictureList(c)
	case "SSI0":
		RouterPostSessionSurveyInfo(c)
	case "SSS0":
		RouterPostSubmitSessionSurvey(c)
	case "DSS0":
		RouterPostSubmitDKOMSurvey(c)
	case "SD0":
		RouterPostSessionDetail(c)
	case "PML0":
		RouterPostPictureMyList(c)
	case "DVL0":
		RouterPostDemoJamVoiceList(c)
	case "MSL0":
		RouterPostMyScoreList(c)
	case "SI0":
		RouterPostSustainbilityInfo(c)
	case "SR0":
		RouterPostSustainabilitySubmit(c)
	case "MP0":
		RouterPostMap(c)
	case "SH0":
		RouterPostScoreHistory(c)
	case "EE0":
		RouterPostEggHikingEnter(c)
	case "EV0":
		RouterPostHiking(c)
	case "I0":
		RouterPostInformation(c)
	}
	MyPrint("sap post finished!")
}

// **********************************************************************************************************************
// **********************************************************************************************************************
//
//			Get Function
//
// **********************************************************************************************************************
// **********************************************************************************************************************
func RouterGetSapVoiceWinner(c *gin.Context) {
	MyPrint("Get : SAP Voice winner start!")
	gSapVoiceCnt[0] = 200
	gSapVoiceCnt[1] = 300
	gSapVoiceCnt[2] = 100
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	if gDB != nil {
		var sapVoiceWinnerList [SAP_VOICE_CNT]int
		sapVoiceWinnerList = getSapVoiceWinnerIdList()
		users := []User{}
		var totalcnt int = 0

		// 6 prize winner
		gDB.Raw("SELECT * FROM User WHERE (VoiceVoteId1 = ? OR VoiceVoteId2 = ?) AND IsPrized = false", sapVoiceWinnerList[2], sapVoiceWinnerList[2]).Scan(&users)
		totalcnt = len(users)
		MyPrint("6 prize users cnt : ", totalcnt)
		if totalcnt > PRIZE_6_CNT {
			totalcnt = PRIZE_6_CNT
		}
		for i := 0; i < totalcnt; i++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			id := r.Intn(totalcnt)
			gDB.Exec("INSERT INTO Winner (UserId, WinnerType) VALUES (?, 'sap voice 3')", users[id].UserId)
			gDB.Exec("UPDATE User Set IsPrized = true WHERE UserId = ?", users[id].UserId)
		}

		// 5 prize winner
		gDB.Raw("SELECT * FROM User WHERE (VoiceVoteId1 = ? OR VoiceVoteId2 = ?) AND IsPrized = false", sapVoiceWinnerList[1], sapVoiceWinnerList[1]).Scan(&users)
		totalcnt = len(users)
		MyPrint("5 prize users cnt : ", totalcnt)
		if totalcnt > PRIZE_5_CNT {
			totalcnt = PRIZE_5_CNT
		}
		for i := 0; i < totalcnt; i++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			id := r.Intn(totalcnt)
			gDB.Exec("INSERT INTO Winner (UserId, WinnerType) VALUES (?, 'sap voice 2')", users[id].UserId)
			gDB.Exec("UPDATE User Set IsPrized = true WHERE UserId = ?", users[id].UserId)
		}

		// 4 prize winner
		gDB.Raw("SELECT * FROM User WHERE (VoiceVoteId1 = ? OR VoiceVoteId2 = ?) AND IsPrized = false", sapVoiceWinnerList[0], sapVoiceWinnerList[0]).Scan(&users)
		totalcnt = len(users)
		MyPrint("4 prize users cnt : ", totalcnt)
		if totalcnt > PRIZE_4_CNT {
			totalcnt = PRIZE_4_CNT
		}
		for i := 0; i < totalcnt; i++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			id := r.Intn(totalcnt)
			gDB.Exec("INSERT INTO Winner (UserId, WinnerType) VALUES (?, 'sap voice 1')", users[id].UserId)
			gDB.Exec("UPDATE User Set IsPrized = true WHERE UserId = ?", users[id].UserId)
		}

		/*
			totalcnt := len(users)
			MyPrint("All users cnt : %d", totalcnt)
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			id := r.Intn(totalcnt)
			MyPrint("Winner is : %d", id)
			js.Set("u", users[id])
			js.Set("r", 1)
			gDB.Exec("INSERT INTO Winner (UserId, WinnerType) VALUES (?, 'sap voice')", users[id].UserId)
		*/
		js.Set("r", 1)
	} else {
		js.Set("r", 0)
	}
	c.JSON(200, js)
	MyPrint("Get : SAP Voice winner start!")
}

func RouterGetWinner(c *gin.Context) {
	MyPrint("Get : Lucky winner start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	if gDB != nil {
		users := []User{}
		gDB.Raw("SELECT * FROM User").Scan(&users)
		totalcnt := len(users)
		MyPrint("All users cnt : %d", totalcnt)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		id := r.Intn(totalcnt)
		MyPrint("Winner is : %d", id)
		js.Set("u", users[id])
		js.Set("r", 1)
		gDB.Exec("INSERT INTO Winner (UserId, WinnerType) VALUES (?, 'luckey')", users[id].UserId)
	} else {
		js.Set("r", 0)
	}
	c.JSON(200, js)
	MyPrint("Get : Lucky winner finished!")
}

func RouterGetCntOfEggHiking(c *gin.Context) {
	MyPrint("Get : Cnt of Egg Hiking start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("r", getEggHikingCnt())
	c.JSON(200, js)
	MyPrint("Get : Cnt of Egg Hiking finished!")
}

func RouterGetCntOfDemoJam(c *gin.Context) {
	MyPrint("Get : Cnt of Demo Jam start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("r", getDemoJamCnt())
	c.JSON(200, js)
	MyPrint("Get : Cnt of Demo Jam finished!")
}

func RouterGetCntOfSapVoice(c *gin.Context) {
	MyPrint("Get : Cnt of Sap Voice start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("r", getSapVoiceCnt())
	c.JSON(200, js)
	MyPrint("Get : Cnt of Sap Voice finished!")
}

func RouterGetGetScoresSwitch(c *gin.Context) {
	MyPrint("Get : Scores Switch start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	value := c.Query("v")
	valuebool, _ := strconv.ParseBool(value)
	MyPrint("scores switch status from ", gCanGetScores, "to ", valuebool)
	js.Set("old score switch", gCanGetScores)
	gCanGetScores = valuebool
	js.Set("new score switch", gCanGetScores)
	c.JSON(200, js)
	MyPrint("Get : Scores Switch finished!")
}

func RouterGetEggHikingSwitch(c *gin.Context) {
	MyPrint("Get : Hiking start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	value := c.Query("v")
	valueId, _ := strconv.Atoi(value)
	MyPrint("Hiking status from ", gEggHikingStatus, "to ", valueId)
	js.Set("old Hiking", gEggHikingStatus)
	gEggHikingStatus = valueId
	js.Set("new Hiking", gEggHikingStatus)
	c.JSON(200, js)
	MyPrint("Get : Hiking finished!")
}

func RouterGetDemoJamSwitch(c *gin.Context) {
	MyPrint("Get : Demo Jam Switch start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	value := c.Query("v")
	valueId, _ := strconv.Atoi(value)
	MyPrint("DemoJame status from ", gDemoJamVoteStatus, "to ", valueId)
	js.Set("old gDemoJamVoteStatus", gDemoJamVoteStatus)
	gDemoJamVoteStatus = valueId
	js.Set("new gDemoJamVoteStatus", gDemoJamVoteStatus)
	c.JSON(200, js)
	MyPrint("Get : Demo Jam Switch finished!")
}

func RouterGetSAPVoiceSwitch(c *gin.Context) {
	MyPrint("Get : SAP Voice Switch start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	value := c.Query("v")
	valueId, _ := strconv.Atoi(value)
	MyPrint("SAP voice status from ", gSAPVoiceStatus, "to ", valueId)
	js.Set("old gSAPVoiceStatus", gSAPVoiceStatus)
	gSAPVoiceStatus = valueId
	js.Set("new gSAPVoiceStatus", gSAPVoiceStatus)
	c.JSON(200, js)
	MyPrint("Get : SAP Voice Switch finished!")
}

func RouterGetToken(c *gin.Context) {
	MyPrint("Get : Token start!")
	uid := c.Query("uid")
	tk := c.Query("tk")
	gDB.Exec("UPDATE User SET DeviceToken = ? WHERE UserId = ?", tk, uid)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "T0")
	js.Set("r", 1)
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : Token finished!")
}

func RouterGetMessage(c *gin.Context) {
	MyPrint("Get : message start!")
	uid := c.Query("uid")
	messages := []Message{}
	if gDB != nil {
		timestamp := time.Now()
		MyPrint("time", timestamp.Unix())
		gDB.Raw("SELECT * FROM Message WHERE (UserId = ? OR UserId = 0) AND MessageTime <= ? AND MessageTime >= ?", uid, timestamp.Unix(), timestamp.Unix()-20*60).Scan(&messages)
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "M0")
	js.Set("mg", messages)
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : message finished!")
}

func RouterGetLogin(c *gin.Context) {
	MyPrint("Get : login start!")
	user := c.Query("usr")
	pwd := c.Query("pwd")
	MyPrint("user name : ", user)
	MyPrint("password : ", pwd)
	isLogin := false
	loginusers := []User{}
	if gDB != nil {
		gDB.Find(&loginusers, "LoginName = ? AND PassWord = ?", user, pwd)
		totalcount := len(loginusers)
		MyPrint("totalcount : ", totalcount)
		if totalcount == 1 {
			isLogin = true
		}
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "L0")
	MyPrint(js)
	if isLogin {
		js.Set("r", "1")
		js.Set("UserId", loginusers[0].UserId)
		MyPrint("login true!")
	} else {
		js.Set("r", "0")
		js.Set("UserId", -1)
		MyPrint("login false!")
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : login finished!")
}

func RouterGetUser(c *gin.Context) {
	MyPrint("Get : user start!")
	uid := c.Query("uid")
	MyPrint("user id : ", uid)
	users := []UserView{}
	findUser := false
	if gDB != nil {
		gDB.Raw("select * from User where UserId = ?", uid).Scan(&users)
		totalcount := len(users)
		MyPrint("totalcount : ", totalcount)
		MyPrint(users)
		if totalcount == 1 {
			findUser = true
		}
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	if findUser {
		js.Set("r", "1")
		js.Set("usr", users)
	} else {
		js.Set("r", "0")
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : user finished!")
}

func RouterGetUserIcon(c *gin.Context) {
	MyPrint("Get : user icon start!")
	uid := c.Query("uid")
	ptype := c.Query("ptype")
	MyPrint("user id : ", uid)
	MyPrint("pic type : ", ptype)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "UI0")
	MyPrint(js)
	js.Set("r", "0")
	MyPrint("create icon false!")
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : user icon finished!")
}

func RouterGetSessionList(c *gin.Context) {
	MyPrint("Get : all session start!")
	allSessionViews := []AllSessionView{}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	if gDB != nil {
		//gDB.Raw("select *, sum(aa.LikeFlag) as LikeCnt from (select a.SessionId, a.Speakerid, a.SessionTitle, a.Format, a.Track, a.StarTime, a.EndTime, a.SessionDescription, a.Point, b.FirstName, b.Lastname, b.SpeakerTitle, b.Company, b.Conuntry, b.Email, b.SpeakerIcon, b.SpeakerDescription, c.LikeFlag, c.CollectionFlag from Session a left join Speaker b on a.SpeakerId = b.SpeakerId left join User_Session_Relation c on a.SessionId = c.SessionId) as aa group by aa.SessionId").Scan(&allSessionViews)
		gDB.Raw("SELECT *, SUM(aa.LikeFlag) AS LikeCnt FROM (select a.SessionId, a.Title, a.Format, a.Location, a.Track, a.StartTime, a.EndTime, a.Description, a.Logo, a.Point, c.LikeFlag, c.CollectionFlag FROM Session a LEFT JOIN User_Session_Relation c ON a.SessionId = c.SessionId) AS aa GROUP BY aa.SessionId").Scan(&allSessionViews)
		totalcount := len(allSessionViews)

		uid := c.Query("uid")
		MyPrint("user id : ", uid)
		user := UserView{}
		gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&user)
		MyPrint(user)
		js.Set("usr", user)

		sidList := []UserSessionRelation{}
		gDB.Raw("SELECT * FROM User_Session_Relation WHERE UserId = ?", uid).Scan(&sidList)
		MyPrint(sidList)

		surveysidList := []SessionSurveyResult{}
		gDB.Raw("SELECT * FROM Session_Survey_Result WHERE UserId = ?", uid).Scan(&surveysidList)

		for id, _ := range allSessionViews {
			allSessionViews[id].CollectionFlag = false
			allSessionViews[id].LikeFlag = false
			for _, sid := range sidList {
				if allSessionViews[id].SessionId == sid.SessionId {
					allSessionViews[id].CollectionFlag = sid.CollectionFlag
					allSessionViews[id].LikeFlag = sid.LikeFlag
					break
				}
			}
			for _, ssid := range surveysidList {
				allSessionViews[id].Done = false
				if allSessionViews[id].SessionId == ssid.SessionId {
					allSessionViews[id].Done = true
					break
				}
			}
		}

		MyPrint("totalcount : ", totalcount)
		MyPrint(allSessionViews)
		js.Set("sel", allSessionViews)

		barRes := []StaticRes{}
		gDB.Raw("SELECT * FROM Static_Res WHERE ResType = 'bar'").Scan(&barRes)
		js.Set("bar", barRes)

		timestamp := time.Now()
		MyPrint("server time : ", timestamp)
		MyPrint("server unix time : ", timestamp.Unix())
		js.Set("stime", timestamp.Unix())
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : all session finished!")
}

func RouterGetVoiceEnter(c *gin.Context) {
	MyPrint("Get : Voice enter start!")
	uid := c.Query("uid")
	MyPrint("user id : ", uid)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "VE0")
	js.Set("r", gSAPVoiceStatus)
	if gDB != nil {
		users := []UserView{}
		gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&users)
		totalcount := len(users)
		MyPrint("totalcount : ", totalcount)
		MyPrint(users)
		if totalcount == 1 {
			js.Set("fv", users[0].VoiceVoteId1)
			js.Set("sv", users[0].VoiceVoteId2)
		} else {
			js.Set("fv", -1)
			js.Set("sv", -1)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : Voice enter finished!")
}

func RouterGetVoiceVote(c *gin.Context) {
	MyPrint("Get : Voice vote start!")
	uid := c.Query("uid")
	vid := c.Query("vid")
	MyPrint("user id : ", uid)
	MyPrint("vote id : ", vid)
	vote := VoiceVote{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	vote.UserId = uidInt
	vidInt, err := strconv.Atoi(vid)
	CheckErr(err)
	vote.VoiceItemId = vidInt
	MyPrint(vote)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "VV0")
	if gDB != nil {
		votes := []VoiceVote{}
		gDB.Raw("select * from Voice_Vote where UserId = ?", uid).Scan(&votes)
		totalcount := len(votes)
		MyPrint("totalcount : ", totalcount)
		MyPrint(votes)
		if totalcount == 0 {
			gDB.Create(&vote)
			js.Set("r", 1)
			gDB.Exec("UPDATE USER SET VoiceVoteId1 = ? WHERE UserId = ?", vid, uid)
			voteSapVoice(vidInt)
		} else if totalcount == 1 {
			if votes[0].VoiceItemId == vidInt {
				js.Set("r", 0)
			} else {
				gDB.Create(&vote)
				js.Set("r", 1)
				gDB.Exec("UPDATE USER SET VoiceVoteId2 = ? WHERE UserId = ?", vid, uid)
				voteSapVoice(vidInt)
			}
		} else if totalcount == 2 {
			js.Set("r", 0)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : Voice vote finished!")
}

func RouterGetVoiceList(c *gin.Context) {
	MyPrint("Get : Voice List start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "VL0")
	if gDB != nil {
		voteItems := []VoiceItem{}
		gDB.Find(&voteItems)
		totalcount := len(voteItems)
		MyPrint("totalcount : ", totalcount)
		MyPrint(voteItems)
		js.Set("vl", voteItems)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : Voice List finished!")
}

func RouterGetDemoJamEnter(c *gin.Context) {
	MyPrint("Get : DemoJam enter start!")
	uid := c.Query("uid")
	MyPrint("user id : ", uid)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "DE0")
	js.Set("r", gDemoJamVoteStatus)
	if gDB != nil {
		users := []UserView{}
		gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&users)
		totalcount := len(users)
		MyPrint("totalcount : ", totalcount)
		MyPrint(users)
		if totalcount == 1 {
			js.Set("fv", users[0].DemoJamId1)
			js.Set("sv", users[0].DemoJamId2)
		} else {
			js.Set("fv", -1)
			js.Set("sv", -1)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : DemoJam enter finished!")
}

func RouterGetDemoJamVote(c *gin.Context) {
	MyPrint("Get : DemoJam vote start!")
	uid := c.Query("uid")
	vid := c.Query("vid")
	MyPrint("user id : ", uid)
	MyPrint("vote id : ", vid)
	vote := DemoJamVote{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	vote.UserId = uidInt
	vidInt, err := strconv.Atoi(vid)
	CheckErr(err)
	vote.DemoJamItemId = vidInt
	MyPrint(vote)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "DV0")
	if gDB != nil {
		votes := []DemoJamVote{}
		gDB.Raw("select * from Demo_Jam_Vote where UserId = ?", uid).Scan(&votes)
		totalcount := len(votes)
		MyPrint("totalcount : ", totalcount)
		MyPrint(votes)
		if totalcount == 0 {
			gDB.Create(&vote)
			js.Set("r", 1)
			gDB.Exec("UPDATE USER SET DemoJamId1 = ? WHERE UserId = ?", vid, uid)
			voteDemoJam(vidInt)
		} else if totalcount == 1 {
			if votes[0].DemoJamItemId == vidInt {
				js.Set("r", 0)
			} else {
				gDB.Create(&vote)
				js.Set("r", 1)
				gDB.Exec("UPDATE USER SET DemoJamId2 = ? WHERE UserId = ?", vid, uid)
				voteDemoJam(vidInt)
			}
		} else if totalcount == 2 {
			js.Set("r", 0)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : DemoJam vote finished!")
}

func RouterGetDemoJamList(c *gin.Context) {
	MyPrint("Get : DemoJam List start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "DL0")
	if gDB != nil {
		djItems := []DemoJamItem{}
		gDB.Find(&djItems)
		totalcount := len(djItems)
		MyPrint("totalcount : ", totalcount)
		MyPrint(djItems)
		js.Set("dl", djItems)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : DemoJam List finished!")
}

func RouterGetSessionVote(c *gin.Context) {
	MyPrint("Get : vote session start!")
	uid := c.Query("uid")
	sid := c.Query("sid")
	value := c.Query("v")
	valueBool, err := strconv.ParseBool(value)
	CheckErr(err)
	MyPrint("user id : ", uid)
	MyPrint("session id : ", sid)
	MyPrint("value : ", valueBool)
	usrelation := UserSessionRelation{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	usrelation.UserId = uidInt
	sidInt, err := strconv.Atoi(sid)
	CheckErr(err)
	usrelation.SessionId = sidInt
	MyPrint(usrelation)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SV0")
	if gDB != nil {
		usrelations := []UserSessionRelation{}
		gDB.Raw("select * from User_Session_Relation where UserId = ? AND SessionId = ?", uid, sid).Scan(&usrelations)
		totalcount := len(usrelations)
		MyPrint("totalcount : ", totalcount)
		MyPrint(usrelations)
		if totalcount > 0 {
			gDB.Exec("UPDATE User_Session_Relation SET LikeFlag=? WHERE UserId = ? AND SessionId = ?", valueBool, uid, sid)
		} else {
			usrelation.LikeFlag = valueBool
			gDB.Create(&usrelation)
		}
		js.Set("r", valueBool)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : vote session finished!")
}

func RouterGetSessionCollect(c *gin.Context) {
	MyPrint("Get : collect session start!")
	uid := c.Query("uid")
	sid := c.Query("sid")
	value := c.Query("v")
	valueBool, err := strconv.ParseBool(value)
	CheckErr(err)
	MyPrint("user id : ", uid)
	MyPrint("session id : ", sid)
	MyPrint("value : ", valueBool)
	usrelation := UserSessionRelation{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	usrelation.UserId = uidInt
	sidInt, err := strconv.Atoi(sid)
	CheckErr(err)
	usrelation.SessionId = sidInt
	MyPrint(usrelation)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SC0")
	if gDB != nil {
		usrelations := []UserSessionRelation{}
		gDB.Raw("select * from User_Session_Relation where UserId = ? AND SessionId = ?", uid, sid).Scan(&usrelations)
		totalcount := len(usrelations)
		MyPrint("totalcount : ", totalcount)
		MyPrint(usrelations)
		if totalcount > 0 {
			gDB.Exec("UPDATE User_Session_Relation SET CollectionFlag=? WHERE UserId = ? AND SessionId = ?", valueBool, uid, sid)
		} else {
			usrelation.CollectionFlag = valueBool
			gDB.Create(&usrelation)
		}
		js.Set("r", valueBool)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : collect session finished!")
}

func RouterGetRank(c *gin.Context) {
	MyPrint("Get : user rank finished!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "R0")
	if gDB != nil {
		users := []UserView{}
		gDB.Raw("SELECT * FROM User ORDER BY Score DESC, SubTime limit 20").Scan(&users)
		totalcount := len(users)
		MyPrint("totalcount : ", totalcount)
		MyPrint(users)
		js.Set("rl", users)

		uid := c.Query("uid")
		MyPrint("user id : ", uid)
		user := UserView{}
		gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&user)
		MyPrint(user)
		js.Set("usr", user)

		var count int = 0
		gDB.Model(User{}).Where("Score > ?", user.Score).Count(&count)
		MyPrint("User now rank is : ", count)
		js.Set("urk", count)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : user rank finished!")
}

func RouterGetPictureSubmit(c *gin.Context) {
	MyPrint("Get : submit picture start!")
	MyPrint("Get : submit picture finished!")
}

func RouterGetPictureDelete(c *gin.Context) {
	MyPrint("Get : delete picture start!")
	uid := c.Query("uid")
	filepath := c.Query("filepath")
	MyPrint("user id : ", uid)
	MyPrint("filepath : ", filepath)
	if gDB != nil {
		gDB.Exec("UPDATE Picture_Wall SET IsDelete = '1' WHERE UserId = ? AND Picture = ? limit 1", uid, filepath)
	}
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "PD0")
	js.Set("r", "1")
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : delete picture finished!")
}

func RouterGetPictureVote(c *gin.Context) {
	MyPrint("Get : vote picture wall start!")
	uid := c.Query("uid")
	pwid := c.Query("pwid")
	value := c.Query("v")
	valueBool, err := strconv.ParseBool(value)
	CheckErr(err)
	MyPrint("user id : ", uid)
	MyPrint("picture wall id : ", pwid)
	MyPrint("value : ", valueBool)
	usrelation := UserPictureRelation{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	usrelation.UserId = uidInt
	pwidInt, err := strconv.Atoi(pwid)
	CheckErr(err)
	usrelation.PictureWallId = pwidInt
	MyPrint(usrelation)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "PV0")
	if gDB != nil {
		usrelations := []UserPictureRelation{}
		gDB.Raw("SELECT * FROM User_Picture_Relation WHERE UserId = ? AND PictureWallId = ?", uid, pwid).Scan(&usrelations)
		totalcount := len(usrelations)
		MyPrint("totalcount : ", totalcount)
		MyPrint(usrelations)
		if totalcount > 0 {
			gDB.Exec("UPDATE User_Picture_Relation SET LikeFlag=? WHERE UserId = ? AND PictureWallId = ?", valueBool, uid, pwid)
		} else {
			usrelation.LikeFlag = valueBool
			gDB.Create(&usrelation)
		}
		js.Set("r", valueBool)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : vote picture wall finished!")
}

func RouterGetPictureList(c *gin.Context) {
	MyPrint("Get : all picture start!")
	uid := c.Query("uid")
	category := c.Query("cat")
	psid := c.Query("psid")
	cnt := c.Query("cnt")
	//	sidInt, err := strconv.Atoi(psid)
	//	cntInt, err := strconv.Atoi(cnt)
	MyPrint("user id : ", uid)
	MyPrint("all pic category : ", category)
	MyPrint("all pic from : ", psid, ", cnt : ", cnt)
	PictureWalls := []PictureWallListView{}
	hasPic := false
	if gDB != nil {
		if category == "All" {
			//gDB.Raw("SELECT * FROM Picture_Wall order by SubTime limit ?, ?", sidInt, cntInt).Scan(&PictureWalls)
			//gDB.Raw("SELECT * FROM User a RIGHT JOIN (SELECT * FROM Picture_Wall ORDER BY SubTime LIMIT ?, ?) b on a.UserId = b.UserId", sidInt, cntInt).Scan(&PictureWalls)
			gDB.Raw("SELECT b.PictureWallId, a.Icon, a.FirstName, a.LastName, a.Title, b.Picture, b.Category, b.Comment, LikeFlagCnt FROM User a RIGHT JOIN (SELECT * FROM Picture_Wall WHERE IsDelete = 0 ORDER BY SubTime DESC LIMIT ?, ?) b on a.UserId = b.UserId left join (SELECT PictureWallId, count(*) as LikeFlagCnt FROM User_Picture_Relation WHERE LikeFlag = 1 group by PictureWallId) c on b.PictureWallId = c.PictureWallId", psid, cnt).Scan(&PictureWalls)
		} else {
			//gDB.Raw("SELECT * FROM Picture_Wall WHERE Category = ? order by SubTime limit ?, ?", catogory, sidInt, cntInt).Scan(&PictureWalls)
			//gDB.Raw("SELECT * FROM User a RIGHT JOIN (SELECT * FROM Picture_Wall WHERE Category = ? ORDER BY SubTime LIMIT ?, ?) b on a.UserId = b.UserId", category, sidInt, cntInt).Scan(&PictureWalls)
			gDB.Raw("SELECT b.PictureWallId, a.Icon, a.FirstName, a.LastName, a.Title, b.Picture, b.Category, b.Comment, LikeFlagCnt FROM User a RIGHT JOIN (SELECT * FROM Picture_Wall WHERE Category = ? AND IsDelete = 0 ORDER BY SubTime DESC LIMIT ?, ?) b on a.UserId = b.UserId left join (SELECT PictureWallId, count(*) as LikeFlagCnt FROM User_Picture_Relation WHERE LikeFlag = 1 group by PictureWallId) c on b.PictureWallId = c.PictureWallId", category, psid, cnt).Scan(&PictureWalls)
		}
		totalcount := len(PictureWalls)
		MyPrint("totalcount : ", totalcount)
		if totalcount > 0 {
			hasPic = true
		}
		upRelations := []UserPictureRelation{}
		gDB.Raw("SELECT * FROM User_Picture_Relation WHERE UserId = ?", uid).Scan(&upRelations)
		for id, _ := range PictureWalls {
			PictureWalls[id].IsLiked = false
			for _, sid := range upRelations {
				if PictureWalls[id].PictureWallId == sid.PictureWallId {
					PictureWalls[id].IsLiked = true
					MyPrint("is liked")
					break
				}
			}
		}
	}
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "PL0")
	if hasPic {
		js.Set("r", "1")
		js.Set("pl", PictureWalls)
	} else {
		js.Set("r", "0")
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : all picture finished!")

}

func RouterGetSessionSurveyInfo(c *gin.Context) {
	MyPrint("Get : session survey info start!")
	sid := c.Query("sid")
	MyPrint("session id : ", sid)
	surveyInfos := []SurveyInfo{}
	hasInfo := false
	if gDB != nil {
		gDB.Raw("SELECT * FROM Survey_Info WHERE SessionId = ?", sid).Scan(&surveyInfos)
		totalcount := len(surveyInfos)
		MyPrint("totalcount : ", totalcount)
		if totalcount == 1 {
			hasInfo = true
		}
	}
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SSI0")
	if hasInfo {
		js.Set("r", 1)
		js.Set("q", surveyInfos)
	} else {
		js.Set("r", 0)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : session survey info finished!")
}

func RouterGetSubmitSessionSurvey(c *gin.Context) {
	MyPrint("Get : submit session survey start!")
	uid := c.Query("uid")
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	sid := c.Query("sid")
	sidInt, err := strconv.Atoi(sid)
	CheckErr(err)
	a1 := c.Query("a1")
	a1Int, err := strconv.Atoi(a1)
	CheckErr(err)
	a2 := c.Query("a2")
	a2Int, err := strconv.Atoi(a2)
	CheckErr(err)
	a3 := c.Query("a3")
	a3Int, err := strconv.Atoi(a3)
	CheckErr(err)
	MyPrint("user id : ", uidInt)
	MyPrint("session id : ", sidInt)
	MyPrint("A1 : ", a1Int)
	MyPrint("A2 : ", a2Int)
	MyPrint("A3 : ", a3Int)
	surveyRes := SessionSurveyResult{}
	surveyRes.SessionId = sidInt
	surveyRes.UserId = uidInt
	surveyRes.A1 = a1Int
	surveyRes.A2 = a2Int
	surveyRes.A3 = a3Int
	ssRes := []SessionSurveyResult{}
	isSurvey := false
	if gDB != nil {
		gDB.Raw("SELECT * FROM Session_Survey_Result WHERE SessionId = ? AND UserId = ?", sid, uid).Scan(&ssRes)
		totalcount := len(ssRes)
		if totalcount > 0 {
			isSurvey = true
		}
		if !isSurvey {
			gDB.Create(&surveyRes)
		}
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SSS0")
	if isSurvey {
		js.Set("r", 0)
	} else {
		user := UserView{}
		answer := []SurveyAnswerInfo{}
		gDB.Raw("SELECT * FROM Survey_Info WHERE SessionId = ?", sid).Scan(&answer)
		totalcount := len(answer)
		if totalcount == 1 && answer[0].Answer1 == a1Int && answer[0].Answer2 == a2Int {
			addscore := AddUserScore(uidInt, SessionSurveyID, sid)
			js.Set("r", 1)
			js.Set("add", addscore)
		} else {
			js.Set("r", 2)
		}
		var rank int = 0
		gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&user)
		//	loc, _ := time.LoadLocation("Asia/Shanghai")
		//	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", user.SubTime.Unix(), loc)
		//gDB.Exec("SELECT COUNT(*) FROM User WHERE Score > ? && SubTime < ?", user.Score, user.SubTime).Count(&rank)
		gDB.Table("User").Where("Score > ?", user.Score).Count(&rank) //.Where("SubTime < ?", user.SubTime).Count(&rank)
		MyPrint("rank now is : ", rank)
		js.Set("rank", rank)
		js.Set("points", user.Score)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : submit session survey finished!")
}

func RouterGetSubmitDKOMSurvey(c *gin.Context) {
	MyPrint("Get : submit session survey start!")
	uid := c.Query("uid")
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	q1 := c.Query("q1")
	q1Int, err := strconv.Atoi(q1)
	CheckErr(err)
	q2 := c.Query("q2")
	q2Int, err := strconv.Atoi(q2)
	CheckErr(err)
	q3 := c.Query("q3")
	q3Int, err := strconv.Atoi(q3)
	CheckErr(err)
	q4 := c.Query("q4")
	q4Int, err := strconv.Atoi(q4)
	CheckErr(err)
	MyPrint("user id : ", uidInt)
	MyPrint("Q1 : ", q1Int)
	MyPrint("Q2 : ", q2Int)
	MyPrint("Q3 : ", q3Int)
	MyPrint("Q4 : ", q4Int)
	surveyRes := DkomSurveyResult{}
	surveyRes.UserId = uidInt
	surveyRes.Q1 = q1Int
	surveyRes.Q2 = q2Int
	surveyRes.Q3 = q3Int
	surveyRes.Q4 = q4Int
	if gDB != nil {
		gDB.Create(&surveyRes)
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "DSS0")
	js.Set("r", "1")
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : submit session survey finished!")
}

func RouterGetSessionDetail(c *gin.Context) {
	MyPrint("Get : submit detail start!")
	uid := c.Query("uid")
	sid := c.Query("sid")
	MyPrint("User id : ", uid)
	MyPrint("Session id : ", sid)
	sessions := []Session{}
	speakers := []Speaker{}
	ssRes := []SessionSurveyResult{}
	isFind := false
	isSurvey := false
	if gDB != nil {
		gDB.Raw("SELECT * FROM Session WHERE SessionId = ?", sid).Scan(&sessions)
		gDB.Raw("SELECT * FROM User a RIGHT JOIN (SELECT * FROM Speaker_Session_Relation WHERE SessionId = ?) AS b ON a.UserId = b.SpeakerId;", sid).Scan(&speakers)
		if len(sessions) == 1 {
			isFind = true
		}
		gDB.Raw("SELECT * FROM Session_Survey_Result WHERE SessionId = ? AND UserId = ?", sid, uid).Scan(&ssRes)
		totalcount := len(ssRes)
		if totalcount > 0 {
			isSurvey = true
		}
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SD0")
	if isFind {
		js.Set("r", "1")
		js.Set("s", sessions)
		js.Set("sp", speakers)
		js.Set("sv", isSurvey)
	} else {
		js.Set("r", "0")
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : submit detail finished!")
}

func RouterGetPictureMyList(c *gin.Context) {
	MyPrint("Get : my picture list start!")
	uid := c.Query("uid")
	myPictures := []PictureWallView{}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "PML0")
	if gDB != nil {
		gDB.Raw("SELECT * FROM Picture_Wall WHERE IsDelete = 0 AND UserId = ? order by SubTime", uid).Scan(&myPictures)
		totalcount := len(myPictures)
		MyPrint("totalcount : ", totalcount)
		if totalcount > 0 {
			js.Set("r", "1")
			js.Set("pl", myPictures)
		} else {
			js.Set("r", "0")
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : my picture list finished!")
}

func RouterGetDemoJamVoiceList(c *gin.Context) {
	MyPrint("Get : DemoJam Voice List start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "DVL0")
	if gDB != nil {
		djItems := []DemoJamItem{}
		gDB.Find(&djItems)
		totalcount := len(djItems)
		MyPrint("demo jam totalcount : ", totalcount)
		MyPrint(djItems)
		js.Set("dl", djItems)

		voteItems := []VoiceItem{}
		gDB.Find(&voteItems)
		totalcount = len(voteItems)
		MyPrint("sap voice totalcount : ", totalcount)
		MyPrint(voteItems)
		js.Set("vl", voteItems)

		ehItems := []EggHikingItem{}
		gDB.Find(&ehItems)
		totalcount = len(ehItems)
		MyPrint("egg hiking totalcount : ", totalcount)
		MyPrint(ehItems)
		js.Set("eh", ehItems)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : DemoJam Voice List finished!")
}

func RouterGetSustainbilityInfo(c *gin.Context) {
	MyPrint("Get : My Sustainbility Info start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SI0")
	js.Set("r", sustainbilityContext)
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : My Sustainbility Info finished!")
}

func RouterGetMyScoreList(c *gin.Context) {
	MyPrint("Get : My Score List start!")
	//uid := c.Query("uid")

	MyPrint("Get : My Score List finished!")
}

func RouterGetSustainabilitySubmit(c *gin.Context) {
	MyPrint("Get : Submit Sustainbility Survey start!")
	uid := c.Query("uid")
	uidInt, _ := strconv.Atoi(uid)
	MyPrint("User id : ", uid)
	user := UserView{}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SR0")
	if gDB != nil {
		gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&user)
		MyPrint(user)
		if user.GreenAmb {
			js.Set("r", 0)
		} else {
			js.Set("r", 1)
			gDB.Exec("UPDATE User SET GreenAmb = 1 WHERE UserId = ?", uid)
			AddUserScore(uidInt, SustainabilityCampaignID, "Sustainabiliity Campaign")
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : Submit Sustainbility Survey finished!")
}

func RouterGetMap(c *gin.Context) {
	MyPrint("Get : Map start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "MP0")
	mapRes := []StaticRes{}
	if gDB != nil {
		gDB.Raw("SELECT * FROM Static_Res WHERE ResType = 'map' ORDER BY ResLable").Scan(&mapRes)
	}
	if len(mapRes) > 0 {
		js.Set("r", 1)
		js.Set("map", mapRes)
	} else {
		js.Set("r", 0)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : Map finished!")
}

func RouterGetScoreHistory(c *gin.Context) {
	MyPrint("Get : Score History start!")
	uid := c.Query("uid")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SH0")
	if gDB != nil {
		sh := []ScoreHistory{}
		gDB.Raw("SELECT * FROM Score_History WHERE UserId = ?", uid).Scan(&sh)
		//totalcount := len(sh)
		js.Set("r", 1)
		js.Set("h", sh)
	} else {
		js.Set("r", 0)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : Score History finished!")
}

func RouterGetEggHikingEnter(c *gin.Context) {
	MyPrint("Get : Enter Egg Hiking start!")
	uid := c.Query("uid")
	MyPrint("User id : ", uid)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "EE0")
	if gDB != nil {
		js.Set("r", gEggHikingStatus)
		if gEggHikingStatus == VOTE_START {
			users := []UserView{}
			gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&users)
			if len(users) == 1 {
				if users[0].EggVoted {
					js.Set("e", 0)
				} else {
					js.Set("e", 1)
				}
			} else {
				js.Set("e", 0)
			}
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : Enter Egg Hiking finished!")
}

func RouterGetHiking(c *gin.Context) {
	MyPrint("Get : Egg Hiking start!")
	uid := c.Query("uid")
	MyPrint("User id : ", uid)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "EV0")
	if gDB != nil {
		rs := []HikingVote{}
		gDB.Raw("SELECT * FROM Hiking_Vote WHERE UserId = ?", uid).Scan(&rs)
		totalcount := len(rs)
		if totalcount == 0 {
			gDB.Exec("INSERT INTO Hiking_Vote (UserId, VoteFlag) VALUES (?, 1)", uid)
			gDB.Exec("UPDATE User SET EggVoted = true WHERE UserId = ?", uid)
			js.Set("r", 1)
			voteEggHiking()
		} else {
			js.Set("r", 0)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : Egg Hiking finished!")
}

func RouterGetInformation(c *gin.Context) {
	MyPrint("Get : Information start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "I0")
	meRes := []StaticRes{}
	if gDB != nil {
		gDB.Raw("SELECT * FROM Static_Res WHERE ResType = 'me'").Scan(&meRes)
	}
	if len(meRes) > 0 {
		js.Set("r", 1)
		js.Set("me", meRes)
	} else {
		js.Set("r", 0)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : Information finished!")
}

// **********************************************************************************************************************
// **********************************************************************************************************************
//
//			Post Function
//
// **********************************************************************************************************************
// **********************************************************************************************************************
func RouterPostToken(c *gin.Context) {
	MyPrint("Post : Token start!")
	uid := c.PostForm("uid")
	tk := c.PostForm("tk")
	if gDB != nil {
		gDB.Exec("UPDATE User SET DeviceToken = ? WHERE UserId = ?", tk, uid)
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "T0")
	js.Set("r", 1)
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : Token finished!")
}

func RouterPostMessage(c *gin.Context) {
	MyPrint("Post : message start!")
	uid := c.PostForm("uid")
	messages := []Message{}
	if gDB != nil {
		timestamp := time.Now()
		MyPrint("time", timestamp.Unix())
		gDB.Raw("SELECT * FROM Message WHERE (UserId = ? OR UserId = 0) AND MessageTime <= ? AND MessageTime >= ?", uid, timestamp.Unix(), timestamp.Unix()-20*60).Scan(&messages)
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "M0")
	js.Set("mg", messages)
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : message finished!")
}

func RouterPostLogin(c *gin.Context) {
	MyPrint("Post : login start!")
	user := c.PostForm("usr")
	pwd := c.PostForm("pwd")
	MyPrint("user name : ", user)
	MyPrint("password : ", pwd)
	isLogin := false
	loginusers := []User{}
	if gDB != nil {
		gDB.Find(&loginusers, "LoginName = ? AND PassWord = ?", user, pwd)
		totalcount := len(loginusers)
		MyPrint("totalcount : ", totalcount)
		if totalcount == 1 {
			isLogin = true
		}
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "L0")
	if isLogin {
		js.Set("r", "1")
		js.Set("UserId", loginusers[0].UserId)
		MyPrint("login true!")
	} else {
		js.Set("r", "0")
		js.Set("UserId", -1)
		MyPrint("login false!")
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : login finished!")
}

func RouterPostUser(c *gin.Context) {
	MyPrint("Post : user start!")
	uid := c.PostForm("uid")
	MyPrint("user id : ", uid)
	users := []UserView{}
	findUser := false
	if gDB != nil {
		gDB.Raw("select * from User where UserId = ?", uid).Scan(&users)
		totalcount := len(users)
		MyPrint("totalcount : ", totalcount)
		MyPrint(users)
		if totalcount == 1 {
			findUser = true
		}
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	if findUser {
		js.Set("r", "1")
		js.Set("usr", users)
	} else {
		js.Set("r", "0")
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : user finished!")
}

func RouterPostUserIcon(c *gin.Context) {
	MyPrint("Post : user icon start!")
	uid := c.PostForm("uid")
	uidInt, _ := strconv.Atoi(uid)
	ptype := c.PostForm("ptype")
	file, header, err := c.Request.FormFile("filepath")
	filename := header.Filename
	MyPrint("user id : ", uid)
	MyPrint("pic type : ", ptype)
	MyPrint("pic name : ", filename)
	timeStr := strconv.FormatInt(time.Now().Unix(), 10)
	serverfilename := uid + "/" + IconFileName + timeStr + "." + ptype
	MyPrint("icon file name : ", serverfilename)
	createIcon := true
	isFirstUpload := false
	filedir, _ := filepath.Abs(RootResDir + uid) // + "/" + IconFileName + "." + ptype)
	MyPrint("server dir : ", filedir)

	var f *os.File
	if !CheckDirIsExist(filedir) {
		os.MkdirAll(filedir, 0700)
	}
	filedir += "/" + IconFileName + timeStr + "." + ptype
	MyPrint("server dir : ", filedir)
	if CheckFileIsExist(filedir) {
		f, err = os.OpenFile(filedir, os.O_WRONLY, 0666)
		MyPrint("open user icon : ", serverfilename)
	} else {
		f, err = os.Create(filedir)
		isFirstUpload = true
		MyPrint("create user icon : ", serverfilename)
	}
	defer f.Close()
	//f, err := os.OpenFile(ResDir + filename, os.O_CREATE|os.O_WRONLY, 0666)
	if CheckErr(err) {
		MyPrint("upload user icon name failed!")
		createIcon = false
	}
	_, err = io.Copy(f, file)
	if CheckErr(err) {
		MyPrint("upload user icon failed!")
		createIcon = false
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	users := []User{}
	if gDB != nil {
		gDB.Find(&users, "UserId = ?", uid)
		totalcount := len(users)
		MyPrint("totalcount : ", totalcount)
		if totalcount == 1 {
			gDB.Exec("UPDATE User SET Icon = ? WHERE UserId = ?", serverfilename, uid)
		} else {
			createIcon = false
		}
	}

	js.Set("i", "UI0")
	MyPrint(js)
	if createIcon {
		js.Set("r", "1")
		MyPrint("create icon succeed!")
		if isFirstUpload {
			AddUserScore(uidInt, UploadAvatarID, "Upload Avater")
		}
	} else {
		js.Set("r", "0")
		MyPrint("create icon false!")
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : user icon finished!")
}

func RouterPostSessionList(c *gin.Context) {
	MyPrint("Post : all session start!")
	allSessionViews := []AllSessionView{}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	if gDB != nil {
		//gDB.Raw("select *, sum(aa.LikeFlag) as LikeCnt from (select a.SessionId, a.Speakerid, a.SessionTitle, a.Format, a.Track, a.StarTime, a.EndTime, a.SessionDescription, a.Point, b.FirstName, b.Lastname, b.SpeakerTitle, b.Company, b.Conuntry, b.Email, b.SpeakerIcon, b.SpeakerDescription, c.LikeFlag, c.CollectionFlag from Session a left join Speaker b on a.SpeakerId = b.SpeakerId left join User_Session_Relation c on a.SessionId = c.SessionId) as aa group by aa.SessionId").Scan(&allSessionViews)
		gDB.Raw("SELECT *, SUM(aa.LikeFlag) AS LikeCnt FROM (select a.SessionId, a.Title, a.Format, a.Location, a.Track, a.StartTime, a.EndTime, a.Description, a.Logo, a.Point, c.LikeFlag, c.CollectionFlag FROM Session a LEFT JOIN User_Session_Relation c ON a.SessionId = c.SessionId) AS aa GROUP BY aa.SessionId").Scan(&allSessionViews)
		totalcount := len(allSessionViews)

		uid := c.PostForm("uid")
		MyPrint("user id : ", uid)
		user := UserView{}
		gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&user)
		MyPrint(user)
		js.Set("usr", user)

		sidList := []UserSessionRelation{}
		gDB.Raw("SELECT * FROM User_Session_Relation WHERE UserId = ?", uid).Scan(&sidList)
		MyPrint(sidList)

		surveysidList := []SessionSurveyResult{}
		gDB.Raw("SELECT * FROM Session_Survey_Result WHERE UserId = ?", uid).Scan(&surveysidList)

		for id, _ := range allSessionViews {
			allSessionViews[id].CollectionFlag = false
			allSessionViews[id].LikeFlag = false
			for _, sid := range sidList {
				if allSessionViews[id].SessionId == sid.SessionId {
					allSessionViews[id].CollectionFlag = sid.CollectionFlag
					allSessionViews[id].LikeFlag = sid.LikeFlag
					break
				}
			}
			for _, ssid := range surveysidList {
				allSessionViews[id].Done = false
				if allSessionViews[id].SessionId == ssid.SessionId {
					allSessionViews[id].Done = true
					break
				}
			}
		}

		MyPrint("totalcount : ", totalcount)
		MyPrint(allSessionViews)
		js.Set("sel", allSessionViews)

		barRes := []StaticRes{}
		gDB.Raw("SELECT * FROM Static_Res WHERE ResType = 'bar'").Scan(&barRes)
		js.Set("bar", barRes)

		timestamp := time.Now()
		MyPrint("server time : ", timestamp)
		MyPrint("server unix time : ", timestamp.Unix())
		js.Set("stime", timestamp.Unix())
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : all session finished!")
}

func RouterPostVoiceEnter(c *gin.Context) {
	MyPrint("Post : Voice enter start!")
	uid := c.PostForm("uid")
	MyPrint("user id : ", uid)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "VE0")
	js.Set("r", gSAPVoiceStatus)
	if gDB != nil {
		users := []UserView{}
		gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&users)
		totalcount := len(users)
		MyPrint("totalcount : ", totalcount)
		MyPrint(users)
		if totalcount == 1 {
			js.Set("fv", users[0].VoiceVoteId1)
			js.Set("sv", users[0].VoiceVoteId2)
		} else {
			js.Set("fv", -1)
			js.Set("sv", -1)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : Voice enter finished!")
}

func RouterPostVoiceVote(c *gin.Context) {
	MyPrint("Post : Voice vote start!")
	uid := c.PostForm("uid")
	vid := c.PostForm("vid")
	MyPrint("user id : ", uid)
	MyPrint("vote id : ", vid)
	vote := VoiceVote{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	vote.UserId = uidInt
	vidInt, err := strconv.Atoi(vid)
	CheckErr(err)
	vote.VoiceItemId = vidInt
	MyPrint(vote)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "VV0")
	if gDB != nil {
		votes := []VoiceVote{}
		gDB.Raw("select * from Voice_Vote where UserId = ?", uid).Scan(&votes)
		totalcount := len(votes)
		MyPrint("totalcount : ", totalcount)
		MyPrint(votes)
		if totalcount == 0 {
			gDB.Create(&vote)
			js.Set("r", 1)
			gDB.Exec("UPDATE USER SET VoiceVoteId1 = ? WHERE UserId = ?", vid, uid)
			voteSapVoice(vidInt)
		} else if totalcount == 1 {
			if votes[0].VoiceItemId == vidInt {
				js.Set("r", 0)
			} else {
				gDB.Create(&vote)
				js.Set("r", 1)
				gDB.Exec("UPDATE USER SET VoiceVoteId2 = ? WHERE UserId = ?", vid, uid)
				voteSapVoice(vidInt)
			}
		} else if totalcount == 2 {
			js.Set("r", 0)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : Voice vote finished!")
}

func RouterPostVoiceList(c *gin.Context) {
	MyPrint("Post : Voice List start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "VL0")
	if gDB != nil {
		voteItems := []VoiceItem{}
		gDB.Find(&voteItems)
		totalcount := len(voteItems)
		MyPrint("totalcount : ", totalcount)
		MyPrint(voteItems)
		js.Set("vl", voteItems)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : Voice List finished!")
}

func RouterPostDemoJamEnter(c *gin.Context) {
	MyPrint("Post : DemoJam enter start!")
	uid := c.PostForm("uid")
	MyPrint("user id : ", uid)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "DE0")
	js.Set("r", gDemoJamVoteStatus)
	if gDB != nil {
		users := []UserView{}
		gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&users)
		totalcount := len(users)
		MyPrint("totalcount : ", totalcount)
		MyPrint(users)
		if totalcount == 1 {
			js.Set("fv", users[0].DemoJamId1)
			js.Set("sv", users[0].DemoJamId2)
		} else {
			js.Set("fv", -1)
			js.Set("sv", -1)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : DemoJam enter finished!")
}

func RouterPostDemoJamVote(c *gin.Context) {
	MyPrint("Post : DemoJam vote start!")
	uid := c.PostForm("uid")
	vid := c.PostForm("vid")
	MyPrint("user id : ", uid)
	MyPrint("vote id : ", vid)
	vote := DemoJamVote{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	vote.UserId = uidInt
	vidInt, err := strconv.Atoi(vid)
	CheckErr(err)
	vote.DemoJamItemId = vidInt
	MyPrint(vote)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "DV0")
	if gDB != nil {
		votes := []DemoJamVote{}
		gDB.Raw("select * from Demo_Jam_Vote where UserId = ?", uid).Scan(&votes)
		totalcount := len(votes)
		MyPrint("totalcount : ", totalcount)
		MyPrint(votes)
		if totalcount == 0 {
			gDB.Create(&vote)
			js.Set("r", 1)
			gDB.Exec("UPDATE USER SET DemoJamId1 = ? WHERE UserId = ?", vid, uid)
			voteDemoJam(vidInt)
		} else if totalcount == 1 {
			if votes[0].DemoJamItemId == vidInt {
				js.Set("r", 0)
			} else {
				gDB.Create(&vote)
				js.Set("r", 1)
				gDB.Exec("UPDATE USER SET DemoJamId2 = ? WHERE UserId = ?", vid, uid)
				voteDemoJam(vidInt)
			}
		} else if totalcount == 2 {
			js.Set("r", 0)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : DemoJam vote finished!")
}

func RouterPostDemoJamList(c *gin.Context) {
	MyPrint("Post : DemoJam List start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "DL0")
	if gDB != nil {
		djItems := []DemoJamItem{}
		gDB.Find(&djItems)
		totalcount := len(djItems)
		MyPrint("totalcount : ", totalcount)
		MyPrint(djItems)
		js.Set("dl", djItems)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : DemoJam List finished!")
}

func RouterPostSessionVote(c *gin.Context) {
	MyPrint("Post : vote session start!")
	uid := c.PostForm("uid")
	sid := c.PostForm("sid")
	value := c.PostForm("v")
	valueBool, err := strconv.ParseBool(value)
	CheckErr(err)
	MyPrint("user id : ", uid)
	MyPrint("session id : ", sid)
	MyPrint("value : ", valueBool)
	usrelation := UserSessionRelation{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	usrelation.UserId = uidInt
	sidInt, err := strconv.Atoi(sid)
	CheckErr(err)
	usrelation.SessionId = sidInt
	MyPrint(usrelation)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SV0")
	if gDB != nil {
		usrelations := []UserSessionRelation{}
		gDB.Raw("select * from User_Session_Relation where UserId = ? AND SessionId = ?", uid, sid).Scan(&usrelations)
		totalcount := len(usrelations)
		MyPrint("totalcount : ", totalcount)
		MyPrint(usrelations)
		if totalcount > 0 {
			gDB.Exec("UPDATE User_Session_Relation SET LikeFlag=? WHERE UserId = ? AND SessionId = ?", valueBool, uid, sid)
		} else {
			usrelation.LikeFlag = valueBool
			gDB.Create(&usrelation)
		}
		js.Set("r", valueBool)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : vote session finished!")
}

func RouterPostSessionCollect(c *gin.Context) {
	MyPrint("Post : collect session start!")
	uid := c.PostForm("uid")
	sid := c.PostForm("sid")
	value := c.PostForm("v")
	valueBool, err := strconv.ParseBool(value)
	CheckErr(err)
	MyPrint("user id : ", uid)
	MyPrint("session id : ", sid)
	MyPrint("value : ", valueBool)
	usrelation := UserSessionRelation{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	usrelation.UserId = uidInt
	sidInt, err := strconv.Atoi(sid)
	CheckErr(err)
	usrelation.SessionId = sidInt
	MyPrint(usrelation)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SC0")
	if gDB != nil {
		usrelations := []UserSessionRelation{}
		gDB.Raw("select * from User_Session_Relation where UserId = ? AND SessionId = ?", uid, sid).Scan(&usrelations)
		totalcount := len(usrelations)
		MyPrint("totalcount : ", totalcount)
		MyPrint(usrelations)
		if totalcount > 0 {
			gDB.Exec("UPDATE User_Session_Relation SET CollectionFlag=? WHERE UserId = ? AND SessionId = ?", valueBool, uid, sid)
		} else {
			usrelation.CollectionFlag = valueBool
			gDB.Create(&usrelation)
		}
		js.Set("r", valueBool)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : collect session finished!")
}

func RouterPostRank(c *gin.Context) {
	MyPrint("Post : user rank finished!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "R0")
	if gDB != nil {
		users := []UserView{}
		gDB.Raw("SELECT * FROM User ORDER BY Score DESC, SubTime limit 20").Scan(&users)
		totalcount := len(users)
		MyPrint("totalcount : ", totalcount)
		MyPrint(users)
		js.Set("rl", users)

		uid := c.PostForm("uid")
		MyPrint("user id : ", uid)
		user := UserView{}
		gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&user)
		MyPrint(user)
		js.Set("usr", user)

		var count int = 0
		gDB.Model(User{}).Where("Score > ?", user.Score).Count(&count)
		MyPrint("User now score is : ", count)
		js.Set("urk", count)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : user rank finished!")
}

func RouterPostPictureSubmit(c *gin.Context) {
	MyPrint("Post : submit picture start!")
	uid := c.PostForm("uid")
	ptype := c.PostForm("ptype")
	cat := c.PostForm("cat")
	file, header, err := c.Request.FormFile("filepath")
	filename := header.Filename
	MyPrint("user id : ", uid)
	MyPrint("catogory : ", cat)
	MyPrint("pic type : ", ptype)
	MyPrint("pic name : ", filename)
	serverfilename := strconv.FormatInt(time.Now().Unix(), 10)
	serverfilename += "." + ptype //.Format(TimeFormat)
	MyPrint("file name : ", serverfilename)
	subSucceed := true
	filedir, _ := filepath.Abs(RootResDir + uid) // + "/" + IconFileName + "." + ptype)
	MyPrint("server dir : ", filedir)

	var f *os.File
	if !CheckDirIsExist(filedir) {
		os.MkdirAll(filedir, 0700)
		MyPrint("create dir : ", filedir)
	}

	filedir += "/" + serverfilename
	MyPrint("server dir : ", filedir)
	if CheckFileIsExist(filedir) {
		f, err = os.OpenFile(filedir, os.O_WRONLY, 0666)
		MyPrint("open picture : ", serverfilename)
	} else {
		f, err = os.Create(filedir)
		MyPrint("create picture : ", serverfilename)
	}
	defer f.Close()
	if CheckErr(err) {
		MyPrint("upload picture failed!")
		subSucceed = false
	}
	_, err = io.Copy(f, file)
	if CheckErr(err) {
		MyPrint("upload picture failed!")
		subSucceed = false
	}
	uidInt, err := strconv.Atoi(uid)
	if (gDB != nil) && subSucceed {
		CheckErr(err)
		pWall := PictureWall{}
		pWall.UserId = uidInt
		pWall.Category = cat
		pWall.Picture = uid + "/" + serverfilename
		//pWall.PostTime = time.Now().Format(TimeFormatf)
		gDB.Create(pWall)
		MyPrint("create picture in database!")
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "PS0")
	MyPrint(js)
	if subSucceed {
		js.Set("r", "1")
		//AddUserScore(uidInt, UploadPhotoID, "Upload Photo")
		MyPrint("submit picture succeed!")
	} else {
		js.Set("r", "0")
		MyPrint("submit picture false!")
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : submit picture finished!")
}

func RouterPostPictureDelete(c *gin.Context) {
	MyPrint("Post : delete picture start!")
	uid := c.PostForm("uid")
	filepath := c.PostForm("filepath")
	MyPrint("user id : ", uid)
	MyPrint("filepath : ", filepath)
	if gDB != nil {
		gDB.Exec("UPDATE Picture_Wall SET IsDelete = '1' WHERE UserId = ? AND Picture = ? limit 1", uid, filepath)
	}
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "PD0")
	js.Set("r", "1")
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : delete picture finished!")
}

func RouterPostPictureVote(c *gin.Context) {
	MyPrint("Post : vote picture wall start!")
	uid := c.PostForm("uid")
	pwid := c.PostForm("pwid")
	value := c.PostForm("v")
	valueBool, err := strconv.ParseBool(value)
	CheckErr(err)
	MyPrint("user id : ", uid)
	MyPrint("picture wall id : ", pwid)
	MyPrint("value : ", valueBool)
	usrelation := UserPictureRelation{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	usrelation.UserId = uidInt
	pwidInt, err := strconv.Atoi(pwid)
	CheckErr(err)
	usrelation.PictureWallId = pwidInt
	MyPrint(usrelation)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "PV0")
	if gDB != nil {
		usrelations := []UserPictureRelation{}
		gDB.Raw("SELECT * FROM User_Picture_Relation WHERE UserId = ? AND PictureWallId = ?", uid, pwid).Scan(&usrelations)
		totalcount := len(usrelations)
		MyPrint("totalcount : ", totalcount)
		MyPrint(usrelations)
		if totalcount > 0 {
			gDB.Exec("UPDATE User_Picture_Relation SET LikeFlag=? WHERE UserId = ? AND PictureWallId = ?", valueBool, uid, pwid)
		} else {
			usrelation.LikeFlag = valueBool
			gDB.Create(&usrelation)
		}
		js.Set("r", valueBool)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : vote picture wall finished!")
}

func RouterPostPictureList(c *gin.Context) {
	MyPrint("Post : all picture start!")
	uid := c.PostForm("uid")
	category := c.PostForm("cat")
	psid := c.PostForm("psid")
	cnt := c.PostForm("cnt")
	//	sidInt, err := strconv.Atoi(psid)
	//	cntInt, err := strconv.Atoi(cnt)
	MyPrint("user id : ", uid)
	MyPrint("all pic category : ", category)
	MyPrint("all pic from : ", psid, ", cnt : ", cnt)
	PictureWalls := []PictureWallListView{}
	hasPic := false
	if gDB != nil {
		if category == "All" {
			//gDB.Raw("SELECT * FROM Picture_Wall order by SubTime limit ?, ?", sidInt, cntInt).Scan(&PictureWalls)
			//gDB.Raw("SELECT * FROM User a RIGHT JOIN (SELECT * FROM Picture_Wall ORDER BY SubTime LIMIT ?, ?) b on a.UserId = b.UserId", sidInt, cntInt).Scan(&PictureWalls)
			gDB.Raw("SELECT b.PictureWallId, a.Icon, a.FirstName, a.LastName, a.Title, b.Picture, b.Category, b.Comment, LikeFlagCnt FROM User a RIGHT JOIN (SELECT * FROM Picture_Wall WHERE IsDelete = 0 ORDER BY SubTime DESC LIMIT ?, ?) b on a.UserId = b.UserId left join (SELECT PictureWallId, count(*) as LikeFlagCnt FROM User_Picture_Relation WHERE LikeFlag = 1 group by PictureWallId) c on b.PictureWallId = c.PictureWallId", psid, cnt).Scan(&PictureWalls)
		} else {
			//gDB.Raw("SELECT * FROM Picture_Wall WHERE Category = ? order by SubTime limit ?, ?", catogory, sidInt, cntInt).Scan(&PictureWalls)
			//gDB.Raw("SELECT * FROM User a RIGHT JOIN (SELECT * FROM Picture_Wall WHERE Category = ? ORDER BY SubTime LIMIT ?, ?) b on a.UserId = b.UserId", category, sidInt, cntInt).Scan(&PictureWalls)
			gDB.Raw("SELECT b.PictureWallId, a.Icon, a.FirstName, a.LastName, a.Title, b.Picture, b.Category, b.Comment, LikeFlagCnt FROM User a RIGHT JOIN (SELECT * FROM Picture_Wall WHERE Category = ? AND IsDelete = 0 ORDER BY SubTime DESC LIMIT ?, ?) b on a.UserId = b.UserId left join (SELECT PictureWallId, count(*) as LikeFlagCnt FROM User_Picture_Relation WHERE LikeFlag = 1 group by PictureWallId) c on b.PictureWallId = c.PictureWallId", category, psid, cnt).Scan(&PictureWalls)
		}
		totalcount := len(PictureWalls)
		MyPrint("totalcount : ", totalcount)
		if totalcount > 0 {
			hasPic = true
		}
		upRelations := []UserPictureRelation{}
		gDB.Raw("SELECT * FROM User_Picture_Relation WHERE UserId = ?", uid).Scan(&upRelations)
		for id, _ := range PictureWalls {
			PictureWalls[id].IsLiked = false
			for _, sid := range upRelations {
				if PictureWalls[id].PictureWallId == sid.PictureWallId {
					PictureWalls[id].IsLiked = sid.LikeFlag
					MyPrint("is liked")
					break
				}
			}
		}
	}
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "PL0")
	if hasPic {
		js.Set("r", "1")
		js.Set("pl", PictureWalls)
	} else {
		js.Set("r", "0")
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : all picture finished!")

}

func RouterPostSessionSurveyInfo(c *gin.Context) {
	MyPrint("Post : session survey info start!")
	sid := c.PostForm("sid")
	MyPrint("session id : ", sid)
	surveyInfos := []SurveyInfo{}
	hasInfo := false
	if gDB != nil {
		gDB.Raw("SELECT * FROM Survey_Info WHERE SessionId = ?", sid).Scan(&surveyInfos)
		totalcount := len(surveyInfos)
		MyPrint("totalcount : ", totalcount)
		if totalcount == 1 {
			hasInfo = true
		}
	}
	//db.Where("name LIKE ?", "%jin%").Find(&users)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SSI0")
	if hasInfo {
		js.Set("r", 1)
		js.Set("q", surveyInfos)
		MyPrint(surveyInfos)
	} else {
		js.Set("r", 0)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : session survey info finished!")
}

func RouterPostSubmitSessionSurvey(c *gin.Context) {
	MyPrint("Post : submit session survey start!")
	uid := c.PostForm("uid")
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	sid := c.PostForm("sid")
	sidInt, err := strconv.Atoi(sid)
	CheckErr(err)
	a1 := c.PostForm("a1")
	a1Int, err := strconv.Atoi(a1)
	CheckErr(err)
	a2 := c.PostForm("a2")
	a2Int, err := strconv.Atoi(a2)
	CheckErr(err)
	a3 := c.PostForm("a3")
	a3Int, err := strconv.Atoi(a3)
	CheckErr(err)
	MyPrint("user id : ", uidInt)
	MyPrint("session id : ", sidInt)
	MyPrint("A1 : ", a1Int)
	MyPrint("A2 : ", a2Int)
	MyPrint("A3 : ", a3Int)
	surveyRes := SessionSurveyResult{}
	surveyRes.SessionId = sidInt
	surveyRes.UserId = uidInt
	surveyRes.A1 = a1Int
	surveyRes.A2 = a2Int
	surveyRes.A3 = a3Int
	ssRes := []SessionSurveyResult{}
	isSurvey := false
	if gDB != nil {
		gDB.Raw("SELECT * FROM Session_Survey_Result WHERE SessionId = ? AND UserId = ?", sid, uid).Scan(&ssRes)
		totalcount := len(ssRes)
		if totalcount > 0 {
			isSurvey = true
		}
		if !isSurvey {
			gDB.Create(&surveyRes)
		}
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SSS0")
	if isSurvey {
		js.Set("r", 0)
	} else {
		user := UserView{}
		answer := []SurveyAnswerInfo{}
		gDB.Raw("SELECT * FROM Survey_Info WHERE SessionId = ?", sid).Scan(&answer)
		totalcount := len(answer)
		if totalcount == 1 && answer[0].Answer1 == a1Int && answer[0].Answer2 == a2Int {
			addscore := AddUserScore(uidInt, SessionSurveyID, sid)
			js.Set("r", 1)
			js.Set("add", addscore)
		} else {
			js.Set("r", 2)
		}
		var rank int = 0
		gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&user)
		//	loc, _ := time.LoadLocation("Asia/Shanghai")
		//	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", user.SubTime.Unix(), loc)
		//gDB.Exec("SELECT COUNT(*) FROM User WHERE Score > ? && SubTime < ?", user.Score, user.SubTime).Count(&rank)
		gDB.Table("User").Where("Score > ?", user.Score).Count(&rank) //.Where("SubTime < ?", user.SubTime).Count(&rank)
		MyPrint("rank now is : ", rank)
		js.Set("rank", rank)
		js.Set("points", user.Score)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : submit session survey finished!")
}

func RouterPostSubmitDKOMSurvey(c *gin.Context) {
	MyPrint("Get : submit session survey start!")
	uid := c.PostForm("uid")
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	q1 := c.PostForm("q1")
	q1Int, err := strconv.Atoi(q1)
	CheckErr(err)
	q2 := c.PostForm("q2")
	q2Int, err := strconv.Atoi(q2)
	CheckErr(err)
	q3 := c.PostForm("q3")
	q3Int, err := strconv.Atoi(q3)
	CheckErr(err)
	q4 := c.PostForm("q4")
	q4Int, err := strconv.Atoi(q4)
	CheckErr(err)
	MyPrint("user id : ", uidInt)
	MyPrint("Q1 : ", q1Int)
	MyPrint("Q2 : ", q2Int)
	MyPrint("Q3 : ", q3Int)
	MyPrint("Q4 : ", q4Int)
	surveyRes := DkomSurveyResult{}
	surveyRes.UserId = uidInt
	surveyRes.Q1 = q1Int
	surveyRes.Q2 = q2Int
	surveyRes.Q3 = q3Int
	surveyRes.Q4 = q4Int
	if gDB != nil {
		gDB.Create(&surveyRes)
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "DSS0")
	js.Set("r", "1")
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Get : submit session survey finished!")
}

func RouterPostSessionDetail(c *gin.Context) {
	MyPrint("Post : submit detail start!")
	uid := c.PostForm("uid")
	sid := c.PostForm("sid")
	MyPrint("User id : ", uid)
	MyPrint("Session id : ", sid)
	sessions := []Session{}
	speakers := []Speaker{}
	ssRes := []SessionSurveyResult{}
	isFind := false
	isSurvey := false
	if gDB != nil {
		gDB.Raw("SELECT * FROM Session WHERE SessionId = ?", sid).Scan(&sessions)
		gDB.Raw("SELECT * FROM User a RIGHT JOIN (SELECT * FROM Speaker_Session_Relation WHERE SessionId = ?) AS b ON a.UserId = b.SpeakerId;", sid).Scan(&speakers)
		if len(sessions) == 1 {
			isFind = true
		}
		gDB.Raw("SELECT * FROM Session_Survey_Result WHERE SessionId = ? AND UserId = ?", sid, uid).Scan(&ssRes)
		totalcount := len(ssRes)
		if totalcount > 0 {
			isSurvey = true
		}
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SD0")
	if isFind {
		js.Set("r", "1")
		js.Set("s", sessions)
		js.Set("sp", speakers)
		js.Set("sv", isSurvey)
	} else {
		js.Set("r", "0")
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : submit detail finished!")
}

func RouterPostPictureMyList(c *gin.Context) {
	MyPrint("Post : my picture list start!")
	uid := c.PostForm("uid")
	myPictures := []PictureWallView{}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "PML0")
	if gDB != nil {
		gDB.Raw("SELECT * FROM Picture_Wall WHERE IsDelete = 0 AND UserId = ? order by SubTime", uid).Scan(&myPictures)
		totalcount := len(myPictures)
		MyPrint("totalcount : ", totalcount)
		if totalcount > 0 {
			js.Set("r", "1")
			js.Set("pl", myPictures)
		} else {
			js.Set("r", "0")
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : my picture list finished!")
}

func RouterPostDemoJamVoiceList(c *gin.Context) {
	MyPrint("Post : DemoJam Voice List start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "DVL0")
	if gDB != nil {
		djItems := []DemoJamItem{}
		gDB.Find(&djItems)
		totalcount := len(djItems)
		MyPrint("demo jam totalcount : ", totalcount)
		MyPrint(djItems)
		js.Set("dl", djItems)

		voteItems := []VoiceItem{}
		gDB.Find(&voteItems)
		totalcount = len(voteItems)
		MyPrint("sap voice totalcount : ", totalcount)
		MyPrint(voteItems)
		js.Set("vl", voteItems)

		ehItems := []EggHikingItem{}
		gDB.Find(&ehItems)
		totalcount = len(ehItems)
		MyPrint("egg hiking totalcount : ", totalcount)
		MyPrint(ehItems)
		js.Set("eh", ehItems)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : DemoJam Voice List finished!")
}

func RouterPostMyScoreList(c *gin.Context) {
	MyPrint("Post : My Score List start!")
	//uid := c.Query("uid")

	MyPrint("Post : My Score List finished!")
}

func RouterPostSustainbilityInfo(c *gin.Context) {
	MyPrint("Post : My Sustainbility Info start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SI0")
	js.Set("r", sustainbilityContext)
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : My Sustainbility Info finished!")
}

func RouterPostSustainabilitySubmit(c *gin.Context) {
	MyPrint("Post : Sustainbility Info Submit start!")
	uid := c.PostForm("uid")
	uidInt, _ := strconv.Atoi(uid)
	MyPrint("User id : ", uid)
	user := UserView{}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SR0")
	if gDB != nil {
		gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&user)
		MyPrint(user)
		if user.GreenAmb {
			js.Set("r", 0)
		} else {
			js.Set("r", 1)
			gDB.Exec("UPDATE User SET GreenAmb = 1 WHERE UserId = ?", uid)
			AddUserScore(uidInt, SustainabilityCampaignID, "Sustainabiliity Campaign")
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : Sustainbility Info Submit finished!")
}

func RouterPostMap(c *gin.Context) {
	MyPrint("Post : Map start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "MP0")
	mapRes := []StaticRes{}
	if gDB != nil {
		gDB.Raw("SELECT * FROM Static_Res WHERE ResType = 'map' ORDER BY ResLable").Scan(&mapRes)
	}
	if len(mapRes) > 0 {
		js.Set("r", 1)
		js.Set("map", mapRes)
	} else {
		js.Set("r", 0)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : Map finished!")
}

func RouterPostScoreHistory(c *gin.Context) {
	MyPrint("Post : Score History start!")
	uid := c.PostForm("uid")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "SH0")
	if gDB != nil {
		sh := []ScoreHistory{}
		gDB.Raw("SELECT * FROM Score_History WHERE UserId = ?", uid).Scan(&sh)
		//		totalcount := len(sh)
		js.Set("r", 1)
		js.Set("h", sh)
	} else {
		js.Set("r", 0)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : Score History finished!")
}

func RouterPostEggHikingEnter(c *gin.Context) {
	MyPrint("Post : Enter Egg Hiking start!")
	uid := c.PostForm("uid")
	MyPrint("User id : ", uid)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "EE0")
	if gDB != nil {
		js.Set("r", gEggHikingStatus)
		if gEggHikingStatus == VOTE_START {
			users := []UserView{}
			gDB.Raw("SELECT * FROM User WHERE UserId = ?", uid).Scan(&users)
			if len(users) == 1 {
				if users[0].EggVoted {
					js.Set("e", 0)
				} else {
					js.Set("e", 1)
				}
			} else {
				js.Set("e", 0)
			}
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : Enter Egg Hiking finished!")
}

func RouterPostHiking(c *gin.Context) {
	MyPrint("Post : Egg Hiking start!")
	uid := c.PostForm("uid")
	MyPrint("User id : ", uid)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "EV0")
	if gDB != nil {
		rs := []HikingVote{}
		gDB.Raw("SELECT * FROM Hiking_Vote WHERE UserId = ?", uid).Scan(&rs)
		totalcount := len(rs)
		if totalcount == 0 {
			voteEggHiking()
			gDB.Exec("INSERT INTO Hiking_Vote (UserId, VoteFlag) VALUES (?, 1)", uid)
			gDB.Exec("UPDATE User SET EggVoted = true WHERE UserId = ?", uid)
			js.Set("r", 1)
		} else {
			js.Set("r", 0)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : Egg Hiking finished!")
}

func RouterPostInformation(c *gin.Context) {
	MyPrint("Post : Information start!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "I0")
	meRes := []StaticRes{}
	if gDB != nil {
		gDB.Raw("SELECT * FROM Static_Res WHERE ResType = 'me'").Scan(&meRes)
	}
	if len(meRes) > 0 {
		js.Set("r", 1)
		js.Set("me", meRes)
	} else {
		js.Set("r", 0)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	MyPrint(jss)
	MyPrint(js)
	c.JSON(200, jss)
	MyPrint("Post : Information finished!")
}

// **********************************************************************************************************************
// **********************************************************************************************************************
//
//			main function
//
// **********************************************************************************************************************
// **********************************************************************************************************************
func main() {
	argCnt := len(os.Args)

	for i := 1; i < argCnt; i++ {
		if os.Args[i] == "debug" {
			gRelease = false
		} else if os.Args[i] == "local" {
			gLocal = true
		}
	}

	fmt.Println("Release Mode : ", gRelease)

	if gRelease {
		gin.SetMode(gin.ReleaseMode)
	}

	gDB = ConnectDB(gRelease)

	//TestFunc()

	MyPrint("start server!")

	//router := gin.Default()

	router := gin.New()

	//authorized := router.Group("/")
	router.Use(cors.Middleware((cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	})))

	router.GET("/sap", RouterGetSAP)
	router.POST("/sap", RouterPostSAP)

	router.Static(WebResDir, RootResDir)
	router.Static(WebVersionResDir, VersionResDir)

	router.Static(WebHtmlResDir, HtmlResDir)

	router.Run(":8080")

	gDB.Close()
}

// **********************************************************************************************************************
// **********************************************************************************************************************
//
//			common function
//
// **********************************************************************************************************************
// **********************************************************************************************************************
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

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CheckDirIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func ConnectDB(isRelease bool) *gorm.DB {
	fmt.Println("start to connecting db!")
	var connectStr string
	if gLocal {
		connectStr = "root@tcp(127.0.0.1:3306)/SAP?charset=utf8&parseTime=True"
	} else {
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

func TestFunc() {
	MyPrint("start to test db!")
	MyPrint(gDB)
	//var user User
	user := User{}
	gDB.First(&user)
	MyPrint(user)

	users := []User{}
	query := gDB.Find(&users)
	CheckErr(query.Error)
	MyPrint(users)

	for _, v := range users {
		MyPrint("uid : ", v.UserId)
	}

	mytest := Tests{}
	mytests := []Tests{}
	gDB.First(&mytest)
	MyPrint(mytest)
	//tx := db.Begin()
	//db.Model(&mytest).Update("Temp", 50)
	//mytest.IdTests = 899
	//db.Save(&mytest)
	//tx.Commit()
	//MyPrint(mytest)

	mytest.Temp = 120
	//db.Find(&mytest).Update("IdTests", 100)
	gDB.Save(&mytest)
	MyPrint(mytest)

	gDB.Find(&mytests)
	MyPrint(mytests)
	MyPrint("start to test db finished!")
}

func voteEggHiking() {
	if gEggHikingStatus == VOTE_START {
		gEggHikingCnt++
		MyPrint("now egg hiking is : ", gEggHikingCnt)
	}
}

func voteDemoJam(id int) {
	if gDemoJamVoteStatus == VOTE_START {
		if id >= 0 && id < DEMOJAM_CNT {
			gDemoJamCnt[id]++
			MyPrint("now demojam %d is : %d", id, gDemoJamCnt[id])
		}
	}
}

func voteSapVoice(id int) {
	if gSAPVoiceStatus == VOTE_START {
		if id >= 0 && id < SAP_VOICE_CNT {
			gSapVoiceCnt[id]++
			MyPrint("now sap voice %d is : %d", id, gSapVoiceCnt[id])
		}
	}
}

func getEggHikingCnt() int {
	return gEggHikingCnt
}

func getDemoJamCnt() [DEMOJAM_CNT]int {
	return gDemoJamCnt
}

func getSapVoiceCnt() [SAP_VOICE_CNT]int {
	return gSapVoiceCnt
}

func getDemoJamWinnerId() int {
	var winnerId = 0
	var winnerCntNumber = 0
	for i := 0; i < DEMOJAM_CNT; i++ {
		if gDemoJamCnt[i] > winnerCntNumber {
			winnerCntNumber = gDemoJamCnt[i]
			winnerId = i
		}
	}
	MyPrint("DemoJam winner %d is : %d", winnerId, winnerCntNumber)
	return winnerId
}

func getSapVoiceWinnerIdList() [SAP_VOICE_CNT]int {
	var winnerIdList [SAP_VOICE_CNT]int
	var tempGroup [SAP_VOICE_CNT]int
	for i := 0; i < SAP_VOICE_CNT; i++ {
		tempGroup[i] = gSapVoiceCnt[i]
		//MyPrint("tempGroup[i] : ", tempGroup[i])
		//MyPrint("gSapVoiceCnt[i] : ", gSapVoiceCnt[i])
	}
	// 4 prize
	winnerIdList[0] = getSapVoiceWinner(tempGroup)
	tempGroup[winnerIdList[0]] = -1
	// 5 prize
	winnerIdList[1] = getSapVoiceWinner(tempGroup)
	tempGroup[winnerIdList[1]] = -1
	// 6 prize
	winnerIdList[2] = getSapVoiceWinner(tempGroup)
	tempGroup[winnerIdList[2]] = -1
	for i := 0; i < SAP_VOICE_CNT; i++ {
		MyPrint("SAP Voice winner %d is : %d, %d", i, winnerIdList[i], gSapVoiceCnt[winnerIdList[i]])
	}
	return winnerIdList
}

func getSapVoiceWinner(group [SAP_VOICE_CNT]int) int {
	var winnerCntNumber int = 0
	var winnerId int = 0
	for i := 0; i < SAP_VOICE_CNT; i++ {
		//MyPrint("group[i] : ", group[i])
		if group[i] > winnerCntNumber {
			winnerId = i
			winnerCntNumber = group[i]
		}
	}
	//MyPrint("getSapVoiceWinner : ", winnerId)
	return winnerId
}

func notification(tp int, id int, body string, token string) {
	apn, err := apns.New("prod.pem", "key-noenc.pem", "gateway.push.apple.com:2195", 1*time.Second)
	//	apn, err := apns.New("prod.pem", "key-noenc.pem", "gateway.sandbox.push.apple.com:2195", 1*time.Second)
	if err != nil {
		fmt.Printf("connect error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("connect successed!")
	go readError(apn.ErrorChan)
	//token := "a1e909eb31f244fccafe4bcb252ed5e3d1d87d2e0a4d962f9e8946046a8d354e"

	payload := apns.Payload{}
	payload.Aps.Alert.Body = body //"Congratulations!\nYou won a sport camera in the raffle!\nPlease go to the right side of the stage after the party to claim your prize or contact Ms. Karen Zhao at 18800349005."
	payload.Aps.Sound = "bingbong.aiff"
	payload.SetCustom("id", id) //time.Now().Unix())
	payload.SetCustom("tp", tp) //0)

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
	err = apn.Send(&notification)
	MyPrint("send id(%x): %s\n", notification.Identifier, err)

	apn.Close()
}

func readError(errorChan <-chan error) {
	for {
		apnerror := <-errorChan
		fmt.Println(apnerror.Error())
	}
}
