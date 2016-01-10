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
//			logic function
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
	case "V0":
		RouterGetVote(c)
	case "VS0":
		RouterGetVoteSession(c)
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
	case "V0":
		RouterPostVote(c)
	case "VS0":
		RouterPostVoteSession(c)
	}
	fmt.Println("sap post finished!")
}

func RouterPostLogin(c *gin.Context) {
	fmt.Println("login post start!")
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
	fmt.Println("login post finished!")
}

func RouterGetLogin(c *gin.Context) {
	fmt.Println("login get start!")
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
	fmt.Println("login get finished!")
}

func RouterPostAllSession(c *gin.Context) {
	fmt.Println("all session get start!")
	allSessionViews := []AllSessionView{}
	if gDB != nil {
		gDB.Raw("select * from Session natural join Speaker").Scan(&allSessionViews)
		totalcount := len(allSessionViews)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(allSessionViews)
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("sel", allSessionViews)
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("all session get finished!")
}

func RouterGetAllSession(c *gin.Context) {
	fmt.Println("all session get start!")
	allSessionViews := []AllSessionView{}
	if gDB != nil {
		gDB.Raw("select * from Session natural join Speaker").Scan(&allSessionViews)
		totalcount := len(allSessionViews)
		fmt.Println("totalcount : ", totalcount)
		fmt.Println(allSessionViews)
	}
	js, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	js.Set("sel", allSessionViews)
	jss, err := simplejson.NewJson([]byte(`{}`))
	CheckErr(err)
	jss.Set("result", js)
	fmt.Println(jss)
	fmt.Println(js)
	c.JSON(200, jss)
	fmt.Println("all session get finished!")
}

func RouterPostUser(c *gin.Context) {
	fmt.Println("user post start!")
	uid := c.PostForm("uid")
	fmt.Println("user id : ", uid)
	users := []User{}
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
	fmt.Println("user post finished!")
}

func RouterGetUser(c *gin.Context) {
	fmt.Println("user get start!")
	uid := c.Query("uid")
	fmt.Println("user id : ", uid)
	users := []User{}
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
	fmt.Println("user get finished!")
}

func RouterPostVote(c *gin.Context) {
	fmt.Println("vote post start!")
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
	fmt.Println(js)
	c.JSON(200, js)
	fmt.Println("vote post finished!")
}

func RouterGetVote(c *gin.Context) {
	fmt.Println("vote get start!")
	uid := c.Query("uid")
	vid := c.Query("vid")
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
		gDB.Raw("select * from Vote where UserId = ? AND VoteItemId = ?", uid, vid).Scan(&votes)
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
	fmt.Println(js)
	c.JSON(200, js)
	fmt.Println("vote get finished!")
}

func RouterPostVoteSession(c *gin.Context) {
}

func RouterGetVoteSession(c *gin.Context) {
	fmt.Println("vote session get start!")
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
			gDB.Exec("UPDATE User_Session_Relation SET CollectionFlag=? WHERE UserId = ? AND SessionId = ?", !usrelations[0].CollectionFlag, uid, sid)
			js.Set("r", 0)
		} else {
			gDB.Create(&usrelation)
			js.Set("r", 1)
		}
	}
	fmt.Println(js)
	c.JSON(200, js)
	fmt.Println("vote session get finished!")
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
