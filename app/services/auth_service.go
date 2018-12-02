package services

import (
	"log"
	"voting_system/app/services/models"

	"golang.org/x/crypto/bcrypt"
)

// 입력 받은 password를 가지고 해싱해서 hashedPassword를 생성한다
func GenerateHashedPassword(userPassword string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

// 로그인 시 비밀번호 맞는지 확인하는 함수
func CheckPassword(id, userPassword string) bool {
	var account models.Voter

	db := votingdb
	db.Where("id=?", id).
		Find(&account)

	hashedPassword := []byte(account.Password)
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(userPassword))

	if err != nil {
		return false
	} else {
		return true
	}
}

func Exists(id string) bool {
	var voter models.Voter
	db := votingdb
	err := db.Where("id=?", id).
		Find(&voter).Error

	if err != nil {
		return false
	} else {
		return true
	}
}

func CreateVoter(account models.Voter) (models.Voter, error) {
	hashedPassword := GenerateHashedPassword(account.Password)

	record := models.Voter{
		StudentId: account.StudentId,
		Password:  hashedPassword,
		Name:      account.Name,
		Major:     account.Major,
		College:   account.College,
		Mobile:    account.Mobile,
		Address:   account.Address,
		Email:     account.Email,
		Sex:       account.Sex,
		Birth:     account.Birth,
	}

	db := votingdb
	err := db.Set("gorm:save_associations", false).
		Create(&record).Error

	return account, err
}

func Login() {

}

func Logout() {

}
