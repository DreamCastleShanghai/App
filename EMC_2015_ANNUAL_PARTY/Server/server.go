package main


import (
	//"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"fmt"
	//"time"
	"strconv"
	"github.com/bitly/go-simplejson"
	_ "github.com/go-sql-driver/mysql"
)

// my structure
type User struct {
	UserId		int		`gorm:"column:UserId"`
	LoginName	string	`gorm:"column:LoginName"`
	PassWord	string	`gorm:"column:PassWord"`
	FirstName	string	`gorm:"column:FirstName"`
	LastName	string	`gorm:"column:LastName"`
	Icon 		string	`gorm:"column:Icon"`
	Rank		int		`gorm:"column:Rank"`
	Authority	int		`gorm:"column:Authority"`
}

type UserView struct {
	LoginName	string	`gorm:"column:LoginName"`
	FirstName	string	`gorm:"column:FirstName"`
	LastName	string	`gorm:"column:LastName"`
	Icon 		string	`gorm:"column:Icon"`
	Rank		int		`gorm:"column:Rank"`
	Authority	int		`gorm:"column:Authority"`
	DemoJamId	int 	`gorm:"column:DemoJamId"`
	VoiceVoteId	int 	`gorm:"column:VoiceVoteId"`
}

type Session struct {
	SessionId	int 	`gorm:"column:SessionId"`
	SpeakerId	int 	`gorm:"column:SpeakerId"`
	SessionTitle string	`gorm:"column:SessionTitle"`
	Format		string	`gorm:"column:Format"`
	Track		string	`gorm:"column:Track"`
	Location	string	`gorm:"column:Location"`
	StarTime	int64	`gorm:"column:StarTime"`
	EndTime		int64	`gorm:"column:EndTime"`
	SessionDescription	string	`gorm:"column:SessionDescription"`
	Point		int 	`gorm:"column:Point"`
}

type Speaker struct {
	SpeakerId	int 	`gorm:"column:SpeakerId"`
	FirstName	string	`gorm:"column:FirstName"`
	LastName	string	`gorm:"column:LastName"`
	SpeakerTitle string	`gorm:"column:SpeakerTitle"`
	Company		string	`gorm:"column:Company"`
	Country		string	`gorm:"column:Country"`
	Email		string	`gorm:"column:Email"`
	SpeakerIcon string	`gorm:"column:SpeakerIcon"`
	SpeakerDescription	string	`gorm:"column:SpeakerDescription"`
}

type Survey struct {
	SurveyId	int `gorm:"column:SurveyId"`
	UserId		int `gorm:"column:UserId"`
	SpeakerId	int `gorm:"column:SpeakerId"`
	SessionId	int `gorm:"column:SessionID"`
	SpeakerRank	int `gorm:"column:SpeakerRank"`
	SessionRank	int `gorm:"column:SessionRank"`
}

type UserSessionRelation struct {
	RelationId	int 		`gorm:"column:relationid"; primary_key:yes; sql:"AUTO_INCREMENT"`
	UserId		int 		`gorm:"column:UserId"`
	SessionId	int 		`gorm:"column:SessionId"`
	LikeFlag	bool		`gorm:"column:LikeFlag"`
	CollectionFlag	bool	`gorm:"column:CollectionFlag"`
}

type AllSessionView struct {
	SessionId	int 	`gorm:"column:SessionId"`
	SessionTitle string	`gorm:"column:SessionTitle"`
	Format		string	`gorm:"column:Format"`
	Track		string	`gorm:"column:Track"`
	Location	string	`gorm:"column:Location"`
	StarTime	int64	`gorm:"column:StarTime"`
	EndTime		int64	`gorm:"column:EndTime"`
	SessionDescription	string	`gorm:"column:SessionDescription"`
	Point		int 	`gorm:"column:Point"`
	FirstName	string	`gorm:"column:FirstName"`
	LastName	string	`gorm:"column:LastName"`
	SpeakerTitle string	`gorm:"column:SpeakerTitle"`
	Company		string	`gorm:"column:Company"`
	Country		string	`gorm:"column:Country"`
	Email		string	`gorm:"column:Email"`
	SpeakerIcon string	`gorm:"column:SpeakerIcon"`
	SpeakerDescription	string	`gorm:"column:SpeakerDescription"`
	LikeFlag	bool 	`gorm:"column:LikeFlag"`
	CollectionFlag bool	`gorm:"column:CollectionFlag"`
	CollectedCnt	int `gorm:"column:CollectedCnt"`
}

type TempSession struct {
	SessionId	int 	`gorm:"column:SessionId"`	
}

type VoiceVote struct {
	VoiceVoteId	int 	`gorm:"column:VoiceVoteId;sql:"AUTO_INCREMENT"`
	UserId		int 	`gorm:"column:UserId"`
	VoiceItemId int 	`gorm:"column:VoiceItemId"`
}

type DemoJamVote struct {
	DemoJamVoteId	int 	`gorm:"column:DemoJamVoteId;sql:"AUTO_INCREMENT"`
	UserId			int 	`gorm:"column:UserId"`
	DemoJamItemId 	int 	`gorm:"column:DemoJamItemId"`
}

type Vote struct {
	VoteId	int 	`gorm:"column:VoteId;sql:"AUTO_INCREMENT"`
	UserId	int 	`gorm:"column:UserId"`
	VoteItemId int 	`gorm:"column:VoteItemId"`
}

type Tests struct {
	IdTests	int `gorm:"column:id_tests; primary_key:yes"`
	Temp	int `gorm:"column:temp"`
}

var gDB *gorm.DB

// ***********************************************************
//
//			router's selection logic function
//
// ***********************************************************
func RouterGetSAP(c *gin.Context) {
	fmt.Println("sap get start!")
	msgType := c.Query("tag")
	fmt.Println("tag is : ", msgType)
	switch msgType {
	case "L0":
		RouterGetLogin(c)
	case "S0":
		RouterGetAllSession(c)
	case "U0":
		RouterGetUser(c)
	case "VV0":
		RouterGetVoteVoice(c)
	case "VD0":
		RouterGetVoteDemoJam(c)
	case "VS0":
		RouterGetVoteSession(c)
	case "R0":
		RouterGetRank(c)
	}
	fmt.Println("sap get finished!")
}

func RouterPostSAP(c *gin.Context) {
	fmt.Println("sap post start!")
	msgType := c.PostForm("tag")
	fmt.Println("tag is : ", msgType)
	switch msgType {
	case "L0":
		RouterPostLogin(c)
	case "S0":
		RouterPostAllSession(c)
	case "U0":
		RouterPostUser(c)
	case "VV0":
		RouterGetVoteVoice(c)
	case "VD0":
		RouterGetVoteDemoJam(c)
	case "VS0":
		RouterPostVoteSession(c)
	case "R0":
		RouterPostRank(c)
	}
	fmt.Println("sap post finished!")
}





// ***********************************************************
//
//			Get Function
//
// ***********************************************************
func RouterGetLogin(c *gin.Context) {
	fmt.Println("Get : login start!")
	user := c.Query("usr")
	pwd  := c.Query("pwd")
	fmt.Println("user name : ", user)
	fmt.Println("password : ", pwd)
	isLogin := false
	loginusers := []User{}
	if gDB != nil {
		gDB.Find(&loginusers, "LoginName = ? AND PassWord = ?", user, pwd)
		totalcount := len(loginusers)
		fmt.Println("totalcount : ", totalcount)
		if totalcount == 1 {
			isLogin = true
		}
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "L0")
	fmt.Println(js)
	if isLogin {
		js.Set("r", "1")
		js.Set("UserId", loginusers[0].UserId)
		fmt.Println("login true!")
	} else {
		js.Set("r", "0")
		js.Set("UserId", -1)
		fmt.Println("login false!")
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("Get : login finished!")
}

func RouterGetAllSession(c *gin.Context) {
	fmt.Println("Get : all session start!")
	allSessionViews := []AllSessionView{}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	if gDB != nil {
		gDB.Raw("select *, sum(aa.LikeFlag) as CollectedCnt from (select a.SessionId, a.Speakerid, a.SessionTitle, a.Format, a.Track, a.StarTime, a.EndTime, a.SessionDescription, a.Point, b.FirstName, b.Lastname, b.SpeakerTitle, b.Company, b.Conuntry, b.Email, b.SpeakerIcon, b.SpeakerDescription, c.LikeFlag, c.CollectionFlag from Session a left join Speaker b on a.SpeakerId = b.SpeakerId left join User_Session_Relation c on a.SessionId = c.SessionId) as aa group by aa.SessionId").Scan(&allSessionViews)
		totalcount := len(allSessionViews)

		uid := c.Query("uid")
		fmt.Println("user id : ", uid)
		user := UserView{}
		gDB.Raw("select * from User where UserId = ?", uid).Scan(&user)
		fmt.Println(user)
		js.Set("usr", user)

		sidList := []TempSession{}
		gDB.Raw("select SessionId from User_Session_Relation where CollectionFlag = true AND UserId = ?", uid).Scan(&sidList)
		fmt.Println(sidList)

		for id, _ := range allSessionViews {
			allSessionViews[id].CollectionFlag = false
			fmt.Println("session : ", allSessionViews[id])
			for _, sid := range sidList {
				fmt.Println("sid : ", sid)
				if allSessionViews[id].SessionId == sid.SessionId {
					allSessionViews[id].CollectionFlag = true
					fmt.Println("changed")
					break
				}
			}
		}

		fmt.Println("totalcount : ", totalcount)
		fmt.Println(allSessionViews)
		js.Set("sel", allSessionViews)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("Get : all session finished!")
}

func RouterGetUser(c *gin.Context) {
	fmt.Println("Get : user start!")
	uid := c.Query("uid")
	fmt.Println("user id : ", uid)
	users := []UserView{}
	if gDB != nil {
		gDB.Raw("select * from User where UserId = ?", uid).Scan(&users)
		totalcount := len(users)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(users)
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("result", users)
	fmt.Println(js)
	c.JSON(200, js)
	fmt.Println("Get : user finished!")
}

func RouterGetVoteVoice(c *gin.Context) {
	fmt.Println("Get : DemoJam vote start!")
	uid := c.Query("uid")
	vid := c.Query("vid")
	fmt.Println("user id : ", uid)
	fmt.Println("vote id : ", vid)
	vote := VoiceVote{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	vote.UserId = uidInt
	vidInt, err := strconv.Atoi(vid)
	CheckErr(err)
	vote.VoiceItemId = vidInt
	fmt.Println(vote)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "VV0")
	if gDB != nil {
		votes := []VoiceVote{}
		gDB.Raw("select * from Voice_Vote where UserId = ? AND VoiceItemId = ?", uid, vid).Scan(&votes)
		totalcount := len(votes)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(votes)
		if  totalcount > 0 {
			js.Set("r", 0)
		} else {
			gDB.Create(&vote)
			js.Set("r", 1)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("Get : DemoJam vote finished!")
}

func RouterGetVoteDemoJam(c *gin.Context) {
	fmt.Println("Get : DemoJam vote start!")
	uid := c.Query("uid")
	vid := c.Query("vid")
	fmt.Println("user id : ", uid)
	fmt.Println("vote id : ", vid)
	vote := DemoJamVote{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	vote.UserId = uidInt
	vidInt, err := strconv.Atoi(vid)
	CheckErr(err)
	vote.DemoJamItemId = vidInt
	fmt.Println(vote)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "VD0")
	if gDB != nil {
		votes := []DemoJamVote{}
		gDB.Raw("select * from Demo_Jam_Vote where UserId = ? AND DemoJamItemId = ?", uid, vid).Scan(&votes)
		totalcount := len(votes)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(votes)
		if  totalcount > 0 {
			js.Set("r", 0)
		} else {
			gDB.Create(&vote)
			js.Set("r", 1)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("Get : DemoJam vote finished!")
}

func RouterGetVoteSession(c *gin.Context) {
	fmt.Println("Get : vote session start!")
	uid := c.Query("uid")
	sid := c.Query("sid")
	fmt.Println("user id : ", uid)
	fmt.Println("session id : ", sid)
	usrelation := UserSessionRelation{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	usrelation.UserId = uidInt
	sidInt, err := strconv.Atoi(sid)
	CheckErr(err)
	usrelation.SessionId = sidInt
	fmt.Println(usrelation)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "VS0")
	if gDB != nil {
		usrelations := []UserSessionRelation{}
		gDB.Raw("select * from User_Session_Relation where UserId = ? AND SessionId = ?", uid, sid).Scan(&usrelations)
		totalcount := len(usrelations)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(usrelations)
		if  totalcount > 0 {
			gDB.Exec("UPDATE User_Session_Relation SET LikeFlag=? WHERE UserId = ? AND SessionId = ?", !usrelations[0].LikeFlag, uid, sid)
			js.Set("r", 0)
		} else {
			usrelation.LikeFlag = true
			gDB.Create(&usrelation)
			js.Set("r", 1)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("Get : vote session finished!")
}

func RouterGetCollectSession(c *gin.Context) {
	fmt.Println("Get : collect session start!")
	uid := c.Query("uid")
	sid := c.Query("sid")
	fmt.Println("user id : ", uid)
	fmt.Println("session id : ", sid)
	usrelation := UserSessionRelation{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	usrelation.UserId = uidInt
	sidInt, err := strconv.Atoi(sid)
	CheckErr(err)
	usrelation.SessionId = sidInt
	fmt.Println(usrelation)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "CS0")
	if gDB != nil {
		usrelations := []UserSessionRelation{}
		gDB.Raw("select * from User_Session_Relation where UserId = ? AND SessionId = ?", uid, sid).Scan(&usrelations)
		totalcount := len(usrelations)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(usrelations)
		if  totalcount > 0 {
			gDB.Exec("UPDATE User_Session_Relation SET CollectionFlag=? WHERE UserId = ? AND SessionId = ?", !usrelations[0].CollectionFlag, uid, sid)
			js.Set("r", 0)
		} else {
			usrelation.CollectionFlag = true
			gDB.Create(&usrelation)
			js.Set("r", 1)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("Get : collect session finished!")
}

func RouterGetRank(c *gin.Context) {
	fmt.Println("Get : user rank finished!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "R0")
	if gDB != nil {
		users := []UserView{}
		gDB.Raw("SELECT * FROM User order by Rank desc limit 10").Scan(&users)
		totalcount := len(users)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(users)
		js.Set("rl", users)

		uid := c.Query("uid")
		fmt.Println("user id : ", uid)
		user := UserView{}
		gDB.Raw("select * from User where UserId = ?", uid).Scan(&user)
		fmt.Println(user)
		js.Set("usr", user)

		var count int = 0
		gDB.Model(User{}).Where("rank > ?", user.Rank).Count(&count)
		fmt.Println("User now rank is : ", count)
		js.Set("urk", count)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("Get : user rank finished!")
}





// ***********************************************************
//
//			Post Function
//
// ***********************************************************
func RouterPostLogin(c *gin.Context) {
	fmt.Println("Post : login start!")
	user := c.PostForm("usr")
	pwd  := c.PostForm("pwd")
	fmt.Println("user name : ", user)
	fmt.Println("password : ", pwd)
	isLogin := false
	loginusers := []User{}
	if gDB != nil {
		gDB.Find(&loginusers, "LoginName = ? AND PassWord = ?", user, pwd)
		totalcount := len(loginusers)
		fmt.Println("totalcount : ", totalcount)
		if totalcount == 1 {
			isLogin = true
		}
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("result", "")
	js.Set("i", "L0")
	if isLogin {
		js.Set("r", "1")
		js.Set("UserId", loginusers[0].UserId)
		fmt.Println("login true!")
	} else {
		js.Set("r", "0")
		js.Set("UserId", -1)
		fmt.Println("login false!")
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("Post : login finished!")
}


func RouterPostAllSession(c *gin.Context) {
	fmt.Println("Post : all session start!")
	allSessionViews := []AllSessionView{}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	if gDB != nil {
		gDB.Raw("select *, sum(aa.LikeFlag) as CollectedCnt from (select a.SessionId, a.Speakerid, a.SessionTitle, a.Format, a.Track, a.StarTime, a.EndTime, a.SessionDescription, a.Point, b.FirstName, b.Lastname, b.SpeakerTitle, b.Company, b.Conuntry, b.Email, b.SpeakerIcon, b.SpeakerDescription, c.LikeFlag, c.CollectionFlag from Session a left join Speaker b on a.SpeakerId = b.SpeakerId left join User_Session_Relation c on a.SessionId = c.SessionId) as aa group by aa.SessionId").Scan(&allSessionViews)
		totalcount := len(allSessionViews)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(allSessionViews)
		js.Set("sel", allSessionViews)

		uid := c.PostForm("uid")
		fmt.Println("user id : ", uid)
		user := UserView{}
		gDB.Raw("select * from User where UserId = ?", uid).Scan(&user)
		fmt.Println(user)
		js.Set("usr", user)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("Post : all session finished!")
}

func RouterPostUser(c *gin.Context) {
	fmt.Println("Post : user start!")
	uid := c.PostForm("uid")
	fmt.Println("user id : ", uid)
	users := []UserView{}
	if gDB != nil {
		gDB.Raw("select * from User where UserId = ?", uid).Scan(&users)
		totalcount := len(users)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(users)
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("result", users)
	fmt.Println(js)
	c.JSON(200, js)
	fmt.Println("Post : user finished!")
}


func RouterPostVote(c *gin.Context) {
	fmt.Println("Post : vote object start!")
	uid := c.PostForm("uid")
	vid := c.PostForm("vid")
	fmt.Println("user id : ", uid)
	fmt.Println("vote id : ", vid)
	vote := Vote{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	vote.UserId = uidInt
	vidInt, err := strconv.Atoi(vid)
	CheckErr(err)
	vote.VoteItemId = vidInt
	fmt.Println(vote)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "V0")
	if gDB != nil {
		votes := []Vote{}
		gDB.Raw("select * from Vote where UserId = ? AND VoteId = ?", uid, vid).Scan(&votes)
		totalcount := len(votes)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(votes)
		if  totalcount > 0 {
			js.Set("r", 0)
		} else {
			gDB.Create(&vote)
			js.Set("r", 1)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("Post : vote ojbect finished!")
}


func RouterPostVoteSession(c *gin.Context) {
	fmt.Println("Post : vote session start!")
	uid := c.PostForm("uid")
	sid := c.PostForm("sid")
	fmt.Println("user id : ", uid)
	fmt.Println("session id : ", sid)
	usrelation := UserSessionRelation{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	usrelation.UserId = uidInt
	sidInt, err := strconv.Atoi(sid)
	CheckErr(err)
	usrelation.SessionId = sidInt
	fmt.Println(usrelation)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "VS0")
	if gDB != nil {
		usrelations := []UserSessionRelation{}
		gDB.Raw("select * from User_Session_Relation where UserId = ? AND SessionId = ?", uid, sid).Scan(&usrelations)
		totalcount := len(usrelations)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(usrelations)
		if  totalcount > 0 {
			gDB.Exec("UPDATE User_Session_Relation SET LikeFlag=? WHERE UserId = ? AND SessionId = ?", !usrelations[0].LikeFlag, uid, sid)
			js.Set("r", 0)
		} else {
			usrelation.LikeFlag = true
			gDB.Create(&usrelation)
			js.Set("r", 1)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("Post : vote session finished!")
}

func RouterPostCollectSession(c *gin.Context) {
	fmt.Println("Post : collect session start!")
	uid := c.PostForm("uid")
	sid := c.PostForm("sid")
	fmt.Println("user id : ", uid)
	fmt.Println("session id : ", sid)
	usrelation := UserSessionRelation{}
	uidInt, err := strconv.Atoi(uid)
	CheckErr(err)
	usrelation.UserId = uidInt
	sidInt, err := strconv.Atoi(sid)
	CheckErr(err)
	usrelation.SessionId = sidInt
	fmt.Println(usrelation)
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "CS0")
	if gDB != nil {
		usrelations := []UserSessionRelation{}
		gDB.Raw("select * from User_Session_Relation where UserId = ? AND SessionId = ?", uid, sid).Scan(&usrelations)
		totalcount := len(usrelations)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(usrelations)
		if  totalcount > 0 {
			gDB.Exec("UPDATE User_Session_Relation SET CollectionFlag=? WHERE UserId = ? AND SessionId = ?", !usrelations[0].CollectionFlag, uid, sid)
			js.Set("r", 0)
		} else {
			usrelation.CollectionFlag = true
			gDB.Create(&usrelation)
			js.Set("r", 1)
		}
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("Post : collect session finished!")
}

func RouterPostRank(c *gin.Context) {
	fmt.Println("Post : user rank finished!")
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("i", "R0")
	if gDB != nil {
		users := []UserView{}
		gDB.Raw("SELECT * FROM User order by Rank desc limit 10").Scan(&users)
		totalcount := len(users)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(users)
		js.Set("rl", users)

		uid := c.PostForm("uid")
		fmt.Println("user id : ", uid)
		user := UserView{}
		gDB.Raw("select * from User where UserId = ?", uid).Scan(&user)
		fmt.Println(user)
		js.Set("usr", user)

		var count int = 0
		gDB.Model(User{}).Where("rank > ?", user.Rank).Count(&count)
		fmt.Println("User now rank is : ", count)
		js.Set("urk", count)
	}
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("Post : user rank finished!")
}

func RouterBaidu(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
}

func RouterSina(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.sina.com.cn")
}








// ***********************************************************
//
//			main function
//
// ***********************************************************
func main() {

	gin.SetMode(gin.ReleaseMode)

	gDB = ConnectDB()

	TestFunc()

	fmt.Println("start server!")
	router := gin.Default()

	router.POST("/login", RouterPostLogin)
	router.GET("/login", RouterGetLogin)

	router.POST("/allsession", RouterPostAllSession)
	router.GET("/allsession", RouterGetAllSession)

	router.GET("/sap", RouterGetSAP)
	router.POST("/sap", RouterPostSAP)

	router.GET("/sina", RouterSina)
	router.GET("/baidu", RouterBaidu)

	router.Run(":8080")

	gDB.Close()
}





// ***********************************************************
//
//			common function
//
// ***********************************************************
func CheckErr(err error) bool {
	if err != nil {
		panic(err)
		return true
	}
	return false
}

func ConnectDB() *gorm.DB {
	fmt.Println("start to connecting db!")
	db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/SAP?charset=utf8&parseTime=True")
	if CheckErr(err) {
		return nil
	}
	fmt.Println("start to connecting db finished!")

	fmt.Println("start to init db!")	
	db.DB()
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
    db.LogMode(true)
	db.SingularTable(true)
	db.AutoMigrate(&User{}, &Tests{})
	fmt.Println("start to init db finished!")

	return &db
}

func TestFunc() {
	fmt.Println("start to test db!")
	fmt.Println(gDB)
	//var user User
	user := User{}
	gDB.First(&user)
	fmt.Println(user)

	users := []User{}
	query := gDB.Find(&users)
	CheckErr(query.Error)
	fmt.Println(users)

	for _, v := range users {
		fmt.Println("uid : ", v.UserId)
	}


	mytest := Tests{}
	mytests := []Tests{}
	gDB.First(&mytest)
	fmt.Println(mytest)
	//tx := db.Begin()
	//db.Model(&mytest).Update("Temp", 50)
	//mytest.IdTests = 899
	//db.Save(&mytest)
	//tx.Commit()
	//fmt.Println(mytest)

	mytest.Temp = 120
	//db.Find(&mytest).Update("IdTests", 100)
	gDB.Save(&mytest)
	fmt.Println(mytest)

	gDB.Find(&mytests)
	fmt.Println(mytests)	
	fmt.Println("start to test db finished!")
}
