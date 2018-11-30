package services

import (
	"voting_system/app/helpers"
	"voting_system/app/services/models"
)

func VoterGetElectionsList(pagination helpers.Pagination) models.Elections {
	var elections models.Elections
	db := votingdb
	db.Limit(pagination.ItemPerPage).
		Offset(pagination.StartIndex).
		Order("id desc").
		Find(&elections)

	return elections
}

func VoterGetPossibleElectionsList(pagination helpers.Pagination) models.Elections {
	var posElections models.Elections

	db := votingdb
	db.Limit(pagination.ItemPerPage).
		Offset(pagination.StartIndex).
		Order("id desc").
		Where("state=?", "on").
		Find(&posElections)

	return posElections
}

func VoterGetElectionCandidatesList(electionid int) models.Candidates {
	var candidates models.Candidates

	db := votingdb
	db.Where("election_id=?", electionid).
		Find(&candidates)

	return candidates
}

func VoterGetCandidateInfo(candidateid int) models.Candidate {
	var candidate models.Candidate
	db := votingdb
	db.Where("id=?", candidateid).
		First(&candidate)

	return candidate
}

func VoterVoting(electionid, candidateid int) error {
	record := models.Voting{
		ElectionId:  electionid,
		CandidateId: candidateid,
		// Auto_hash: "",
	}

	db := votingdb
	err := db.Set("gorm:save_associations", false).
		Create(&record).Error

	return err
}

/*
// return 값: election_id, elected_candidate_id, all_vote, 해당 선거 후보자 결과 리스트
func VoterGetElectionResult(electionid int) (int, int, int, models.EndElectionResult) {
	var endElectionResult models.EndElectionResult
	var endElection models.Election
	var electedCandidate models.EndElectionCandidateInfo

	db := votingdb
	err := db.Model(&models.Election{}).
		Where("election_id=?", electionid).
		Find(&endElection).Error

	if err != nil {
		panic(err)
	}

	// 해당 선거 총 투표수
	var all_vote int
	db.Model(&models.Voting{}).
		Where("election_id=?", electionid).
		Count(&all_vote)

	// 해당 선거 선출자
	db.Model(&models.EndElectionCandidateInfo{}).
		Where("election_id=?", electionid).
		Order("desc poll").
		First(&electedCandidate)

	// 해당 선거 후보자 결과 리스트
	db.Model(&models.EndElectionCandidateInfo{}).
		Where("election_id=?", electionid).
		Order("desc poll").
		Find(&endElectionResult)

	return endElection.Id, electedCandidate.ElectionId, all_vote, endElectionResult
}
*/

func VoterGetElectionResult() models.Elections {
	var endElections models.Elections

	db := votingdb
	db.Where("state=?", 3).
		Find(&endElections)

	return endElections
}

func VoterGetElectionResultCandidate(election_id int) models.EndElectionResult {
	var candidates models.EndElectionResult

	db := votingdb
	db.Model(&models.EndElectionCandidateInfo{}).
		Where("election_id=?", electionid).
		Order("desc poll").
		Find(&candidates)

	return candidates
}
