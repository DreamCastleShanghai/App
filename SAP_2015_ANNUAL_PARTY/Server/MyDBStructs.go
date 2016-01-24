package MyDBStructs

// my structure
type User struct {
	UserId		int		`gorm:"column:UserId;sql:"AUTO_INCREMENT"`
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
	SessionId	int 	`gorm:"column:SessionId;sql:"AUTO_INCREMENT"`
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
	SpeakerId	int 	`gorm:"column:SpeakerId;sql:"AUTO_INCREMENT"`
	FirstName	string	`gorm:"column:FirstName"`
	LastName	string	`gorm:"column:LastName"`
	SpeakerTitle string	`gorm:"column:SpeakerTitle"`
	Company		string	`gorm:"column:Company"`
	Country		string	`gorm:"column:Country"`
	Email		string	`gorm:"column:Email"`
	SpeakerIcon string	`gorm:"column:SpeakerIcon"`
	SpeakerDescription	string	`gorm:"column:SpeakerDescription"`
}

type SurveyInfo struct {
	//SurveyInfoId	int 	`gorm:"column:SurveyId;sql:"AUTO_INCREMENT"`
	Q11			string	`gorm:"column:Q11"`
	Q12			string	`gorm:"column:Q12"`
	Q13			string	`gorm:"column:Q13"`
	Q14			string	`gorm:"column:Q14"`
	Q21			string	`gorm:"column:Q21"`
	Q22			string	`gorm:"column:Q22"`
	Q23			string	`gorm:"column:Q23"`
	Q24			string	`gorm:"column:Q24"`
	Q3			string 	`gorm:"column:Q3"`
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
	LikeCnt		int 	`gorm:"column:LikeCnt"`
	CollectionFlag bool	`gorm:"column:CollectionFlag"`
}

type TempSession struct {
	SessionId	int 	`gorm:"column:SessionId"`	
}

type VoiceItem struct {
	VoiceItemId			int 	`gorm:"column:VoiceItemId;sql:"AUTO_INCREMENT"`
	VoicerName			string	`gorm:"column:VoicerName"`
	SongName			string	`gorm:"column:SongName"`
}

type VoiceVote struct {
	VoiceVoteId	int 	`gorm:"column:VoiceVoteId;sql:"AUTO_INCREMENT"`
	UserId		int 	`gorm:"column:UserId"`
	VoiceItemId int 	`gorm:"column:VoiceItemId"`
	VoicerPic	string	`gorm:"column:VoicerPic"`
}

type DemoJamItem struct {
	DemoJamItemId	int 	`gorm:"column:DemoJamItemId;sql:"AUTO_INCREMENT"`
	TeamName		string	`gorm:"column:TeamName"`
	Department		string	`gorm:"column:Department"`
	Introduction	string	`gorm:"column:Introduction"`
}

type DemoJamVote struct {
	DemoJamVoteId	int 	`gorm:"column:DemoJamVoteId;sql:"AUTO_INCREMENT"`
	UserId			int 	`gorm:"column:UserId"`
	DemoJamItemId 	int 	`gorm:"column:DemoJamItemId"`
}

type PictureWall struct {
	PictureWallId 	int 	`gorm:"column:PictureWallId;sql:"AUTO_INCREMENT"`
	UserId			int 	`gorm:"column:UserId"`
	Picture 		string	`gorm:"column:Picture"`
	Category 		string	`gorm:"column:Category"`
	Comment			string	`gorm:"column:Comment"`
	//IsDelete		bool	`gorm:"column:IsDelete"`
	//PostTime 		int64 	`gorm:"column:PostTime"`
}

type SessionSurveyResult struct {
	//SurveyId 	int 	`gorm:"column:SurveyId;sql:"AUTO_INCREMENT"`
	SessionId 	int 	`gorm:"column:SessionId"`
	UserId 		int 	`gorm:"column:UserId"`
	A1			int 	`gorm:"column:A1"`
	A2			int 	`gorm:"column:A2"`
	A3			int 	`gorm:"column:A3"`
}

type DkomSurveyResult struct {
	//SurveyId 	int 	`gorm:"column:SurveyId;sql:"AUTO_INCREMENT"`
	UserId		int 	`gorm:"column:UserId"`
	Q1 			int 	`gorm:"column:Q1"`
	Q2 			int 	`gorm:"column:Q2"`
	Q3 			int 	`gorm:"column:Q3"`
	Q4 			int 	`gorm:"column:Q4"`
}
/*
type Vote struct {
	VoteId	int 	`gorm:"column:VoteId;sql:"AUTO_INCREMENT"`
	UserId	int 	`gorm:"column:UserId"`
	VoteItemId int 	`gorm:"column:VoteItemId"`
}

type Message struct {
	Name, Text string
}

type test_struct struct {
	Test string
}
*/

type Tests struct {
	IdTests	int `gorm:"column:id_tests; primary_key:yes"`
	Temp	int `gorm:"column:temp"`
}
