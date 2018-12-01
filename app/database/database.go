package database

import (
	"voting_system/app/services/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/votingDB?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	db.LogMode(true)

	DropTablesIfExists(db)
	AutoMigrate(db)
	AutoPopulate(db)
	AddForeignKeys(db)

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.Administrator{},
		&models.Candidate{},
		&models.Election{},
		&models.EndElectionCandidateInfo{},
		&models.Voter{},
		&models.Voting{},
	)
}

func DropTablesIfExists(db *gorm.DB) {
	//db.Exec("DROP TABLE Administrators, Candidates, Elections, EndElectionCandidateInfos, Voters, Votings CASCADE;")
	db.DropTable(&models.Administrator{})
	db.DropTable(&models.Candidate{})
	db.DropTable(&models.Election{})
	db.DropTable(&models.EndElectionCandidateInfo{})
	db.DropTable(&models.Voter{})
	db.DropTable(&models.Voting{})
}

func AddForeignKeys(db *gorm.DB) {

}

func AutoPopulate(db *gorm.DB) {
	PopulateAdmins(db)
	PopulateVoters(db)
	PopulateElection(db)
	PopulateCandidates(db)
}

func PopulateAdmins(db *gorm.DB) {
	db.Create(&models.Administrator{
		Id:        "201202274",
		Password:  "000000",
		Name:      "윤인배",
		Mobile:    "010-2786-2455",
		Ssn:       "",
		Address:   "서초구 내곡동",
		Email:     "iby2455@gmail.com",
		Sex:       "남",
		Birth:     "10/19/1993",
		Authority: 1,
	})
	db.Create(&models.Administrator{
		Id:        "201200000",
		Password:  "000001",
		Name:      "김철수",
		Mobile:    "010-1423-2455",
		Ssn:       "",
		Address:   "서초구 잠실동",
		Email:     "chulsu@gmail.com",
		Sex:       "남",
		Birth:     "10/10/1993",
		Authority: 2,
	})
	db.Create(&models.Administrator{
		Id:        "201211111",
		Password:  "000002",
		Name:      "안현주",
		Mobile:    "010-1423-3942",
		Ssn:       "",
		Address:   "강남구 역삼동",
		Email:     "djdj@gmail.com",
		Sex:       "여",
		Birth:     "12/11/1993",
		Authority: 3,
	})
}

func PopulateVoters(db *gorm.DB) {
	db.Create(&models.Voter{
		StudentId: "201202274",
		Password:  "1111",
		Name:      "윤인배",
		Major:     "컴퓨터공학과",
		College:   "공과대학",
		Mobile:    "010-2786-2455",
		Address:   "서초구 내곡동",
		Email:     "iby2455@naver.com",
		Sex:       "M",
		Birth:     "1993-10-19",
	})
}

func PopulateElection(db *gorm.DB) {
	db.Create(&models.Election{
		Title:             "제 101호 부학생회장",
		Major:             "",
		College:           "종합대학",
		Content:           "부학생회장선거",
		ElectionStartTime: "1541989543",
		ElectionEndTime:   "1542690743",
		State:             1,
		Id:                98,
		AdminId:           "201202274",
	})
	db.Create(&models.Election{
		Title:             "제 101호 학생회장",
		Major:             "",
		College:           "종합대학",
		Content:           "종합대학회장선거",
		ElectionStartTime: "1541989543",
		ElectionEndTime:   "1542690743",
		State:             2,
		Id:                99,
		AdminId:           "201202274",
	})
	db.Create(&models.Election{
		Title:             "제 101호 컴퓨터공학과 회장 선거",
		Major:             "컴퓨터공학과",
		College:           "공과대학",
		Content:           "컴퓨터공학과회장선거",
		ElectionStartTime: "1541989543",
		ElectionEndTime:   "1542690743",
		State:             3,
		Id:                100,
		AdminId:           "201202274",
	})
}

func PopulateCandidates(db *gorm.DB) {
	db.Create(&models.Candidate{
		StudentId:  "201301923",
		Name:       "도로링",
		Major:      "고양이학과",
		College:    "동물대학",
		Thumbnail:  "aaa",
		Resume:     "ddd",
		Id:         13,
		ElectionId: 99,
	})
}

/*
func PopulateCandidates(db *gorm.DB) {
	db.Create(&models.Voter{})
}
*/
