package main


import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"fmt"
//	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// my structure
type User struct {
	UserId		int		`gorm:"column:UserId;"`
	LoginName	string	`gorm:"column:LoginName"`
	PassWord	string	`gorm:"column:PassWord"`
	FirstName	string	`gorm:"column:FirstName"`
	LastName	string	`gorm:"column:LastName"`
	Icon 		string	`gorm:"column:Icon"`
	Rank		int		`gorm:"column:Rank"`
	Authority	int		`gorm:"column:Authority"`
}

type Session struct {
	SessionId	int
	SpeakerId	int
	Title		string
	Format		string
	Track		string
	StarTime	string
	EndTime		string
	Description	string
	Point		int
}

type Speaker struct {
	SpeakerId	int
	FirstName	string
	LastName	string
	Title		string
	Company		string
	Country		string
	Email		string
	Icon 		string
	Description	string
}

type Survey struct {
	SurveyId	int
	UserId		int
	SpeakerId	int
	SessionId	int
	SpeakerRank	int
	SessionRank	int
}

type UserSessionRelation struct {
	UserId		int
	SessionId	int
	LikeFlag	string
	CollectionFlag	string
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
func RouterPostLogin(c *gin.Context) {
	fmt.Println("login post start!")
	user := c.PostForm("usr")
	pwd  := c.PostForm("pwd")
	fmt.Println("user name : ", user)
	fmt.Println("password : ", pwd)
	isLogin := false
	loginusers := []User{}
	if gDB != nil {
		var totalcount int = 0
		gDB.Find(&loginusers, "LoginName = ? AND PassWord = ?", user, pwd)
		totalcount = len(loginusers)
		fmt.Println("totalcount : ", totalcount)
		if totalcount == 1 {
			isLogin = true
		}
	}
	if isLogin {
		c.JSON(200, gin.H{"i": "L0", "r": 1})
		fmt.Println("login true!")
	} else {
		c.JSON(200, gin.H{"i": "L0", "r": 0})
		fmt.Println("login false!")
	}
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
		var totalcount int = 0
		gDB.Find(&loginusers, "LoginName = ? AND PassWord = ?", user, pwd)
		totalcount = len(loginusers)
		fmt.Println("totalcount : ", totalcount)
		if totalcount == 1 {
			isLogin = true
		}
	}
	if isLogin {
		c.JSON(200, gin.H{"i": "L0", "r": 1})
		fmt.Println("login true!")
	} else {
		c.JSON(200, gin.H{"i": "L0", "r": 0})
		fmt.Println("login false!")
	}
	fmt.Println("login get finished!")
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

	//TestFunc()

	fmt.Println("start server!")
	router := gin.Default()
	router.GET("/sina", RouterSina)
	router.GET("/baidu", RouterBaidu)
	router.POST("/login", RouterPostLogin)
	router.GET("/login", RouterGetLogin)
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
	db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/EMC_Annual_Party?charset=utf8&parseTime=True")
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
