package services

import (
	"voting_system/app/helpers"
	"voting_system/app/services/models"
)

func CountRowsCandidate() int {
	var count int
	votingdb.Model(&models.Candidate{}).
		Count(&count)
	return count
}

func CountRowsElection() int {
	var count int
	votingdb.Model(&models.Election{}).
		Count(&count)
	return count
}

// 선거 관리
func AdminGetElectionsList(pagination helpers.Pagination) models.Elections {
	var elections models.Elections
	db := votingdb
	db.Limit(pagination.ItemPerPage).
		Offset(pagination.StartIndex).
		Find(&elections)

	return elections
}

func AdminGetElectionInfo(id int) models.Election {
	var election models.Election
	db := votingdb

	db.Where("id=?", id).
		First(&election)

	return election
}

func AdminCreateElection(election models.Election) (models.Election, error) {
	record := models.Election{
		Title:             election.Title,
		Major:             election.Major,
		College:           election.College,
		Content:           election.Content,
		ElectionStartTime: election.ElectionStartTime,
		ElectionEndTime:   election.ElectionEndTime,
		State:             election.State, // defualt로 1, 선거 전 으로 해야 하나
		Id:                election.Id,    // 이것도 자동으로 auto_increment
		AdminId:           election.AdminId,
	}
	db := votingdb
	err := db.Set("gorm:save_associations", false).
		Create(&record).Error

	return record, err
}

func AdminStartElection(electionid int) (models.Election, error) {
	db := votingdb
	err := db.Set("gorm:save_associations", false).
		Model(&models.Election{}).
		Where("id=?", electionid).
		Update("state", 2).Error

	if err != nil {
		panic(err)
	}

	var election models.Election
	err = db.Where("id=?", electionid).
		Find(&election).Error

	if err != nil {
		panic(err)
	}

	return election, err
}

func AdminEndElection(electionid int) (models.Election, error) {
	db := votingdb
	err := db.Set("gorm:save_associations", false).
		Model(&models.Election{}).
		Where("id=?", electionid).
		Update("state", 3).Error

	if err != nil {
		panic(err)
	}

	var election models.Election
	err = db.Where("id=?", electionid).
		Find(&election).Error

	var count int
	db.Model(&models.Votings{}).
		Where("election_id=?", electionid).
		Count(&count)

	if err != nil {
		panic(err)
	}

	var candidatelist models.Candidates
	db.Model(&models.Candidate{}).
		Where("election_id=?", election.Id).
		Find(&candidatelist)

	for i := 0; i < len(candidatelist); i++ {
		var poll int

		db.Model(&models.Voting{}).
			Where("candidate_id=?", candidatelist[i].Id).
			Count(&poll)

		candidateInfo := AdminGetCandidateInfo(candidatelist[i].Id)
		endElectionCandidateInfo := models.EndElectionCandidateInfo{
			ElectionId:  election.Id,
			All_vote:    count,
			CandidateId: candidateInfo.Id,
			Poll:        poll,
			StudentId:   candidateInfo.StudentId,
			Name:        candidateInfo.Name,
			Major:       candidateInfo.Major,
			College:     candidateInfo.College,
			Thumbnail:   candidateInfo.Thumbnail,
			Resume:      candidateInfo.Resume,
		}

		err := db.Set("gorm:save_associations", false).
			Create(&endElectionCandidateInfo).Error

		if err != nil {
			panic(err)
		}
	}

	return election, err
}

func AdminEditElection(electionid int, election models.Election) (models.Election, error) {
	record := models.Election{
		Title:             election.Title,
		Major:             election.Major,
		College:           election.College,
		Content:           election.Content,
		ElectionStartTime: election.ElectionStartTime,
		ElectionEndTime:   election.ElectionEndTime,
		State:             election.State,
		Id:                election.Id,
		AdminId:           election.AdminId,
	}

	db := votingdb
	err := db.Set("gorm:save_associations", false).
		Model(&models.Election{}).
		Where("elction_id=?", electionid).
		Updates(&record).Error

	if err != nil {
		panic(err)
	}

	return record, err
}

func AdminElectionResult(pagination helpers.Pagination) models.Elections {
	var endElections models.Elections

	db := votingdb
	db.Limit(pagination.ItemPerPage).
		Offset(pagination.StartIndex).
		Where("state=?", 3).
		Find(&endElections)

	return endElections
}

func AdminElectionResultCandidates(electionid int) models.EndElectionResult {
	var candidates models.EndElectionResult

	db := votingdb
	db.Model(&models.EndElectionCandidateInfo{}).
		Where("election_id=?", electionid).
		Order("desc poll").
		Find(&candidates)

	return candidates
}

// 후보자 관리
func AdminGetCandidatesList(pagination helpers.Pagination) models.Candidates {
	var candidates models.Candidates
	db := votingdb
	db.Limit(pagination.ItemPerPage).
		Offset(pagination.StartIndex).
		//Order("id desc").
		Find(&candidates)

	return candidates
}

func AdminGetCandidateInfo(id int) models.Candidate {
	var candidate models.Candidate
	db := votingdb
	db.Where("id=?", id).
		First(&candidate)

	return candidate
}

func AdminCreateCandidate(candidate models.Candidate) (models.Candidate, error) {
	record := models.Candidate{
		Name:       candidate.Name,
		Major:      candidate.Major,
		College:    candidate.College,
		Thumbnail:  candidate.Thumbnail,
		Resume:     candidate.Resume,
		ElectionId: candidate.ElectionId,
		//Id:         candidate.Id, //: ID는 auto_increment 설정한다
	}
	db := votingdb
	err := db.Set("gorm:save_associations", false).
		Create(&record).Error

	return record, err
}

func AdminEditCandidate(candidateid int, candidate models.Candidate) (models.Candidate, error) {
	record := models.Candidate{
		Name:       candidate.Name,
		Major:      candidate.Major,
		College:    candidate.College,
		Thumbnail:  candidate.Thumbnail,
		Resume:     candidate.Resume,
		ElectionId: candidate.ElectionId,
	}

	db := votingdb
	err := db.Set("gorm:save_associations", false).
		Model(&models.Candidate{}).
		Where("id=?", candidateid).
		Updates(&record).Error

	if err != nil {
		panic(err)
	}

	return record, err
}
