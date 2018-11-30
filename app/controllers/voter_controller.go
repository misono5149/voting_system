package controllers

import (
	"strconv"
	"voting_system/app/helpers"
	"voting_system/app/services"
	"voting_system/app/services/models"

	"github.com/gin-gonic/gin"
)

func VoterGetElectionsList(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	amountElections := services.CountRowsElection()

	currentPage, _ := strconv.Atoi(queryParams.Get("page"))
	itemPerPage, _ := strconv.Atoi(queryParams.Get("limit"))
	pagination := helpers.MakePagination(amountElections, currentPage, itemPerPage)

	var list models.Elections
	list = services.VoterGetElectionsList(pagination)

	if len(list) <= 0 {
		c.JSON(200, gin.H{
			"status": 200,
			"error":  "생성된 선거가 없습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status":       200,
			"current_page": currentPage,
			"list":         list,
		})
	}
}

func VoterGetPossibleElectionsList(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	amountElections := services.CountRowsElection()

	currentPage, _ := strconv.Atoi(queryParams.Get("page"))
	itemPerPage, _ := strconv.Atoi(queryParams.Get("limit"))
	pagination := helpers.MakePagination(amountElections, currentPage, itemPerPage)

	var list models.Elections
	list = services.VoterGetPossibleElectionsList(pagination)

	if len(list) <= 0 {
		c.JSON(200, gin.H{
			"status": 200,
			"error":  "생성된 선거가 없습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status":       200,
			"current_page": currentPage,
			"list":         list,
		})
	}
}

func VoterGetElectionCandidatesList(c *gin.Context) {
	electionId, err := strconv.Atoi(c.Params.ByName("electionid"))
	if err != nil {
		panic(err)
	}
	list := services.VoterGetElectionCandidatesList(electionId)

	if len(list) <= 0 {
		c.JSON(200, gin.H{
			"status": 200,
			"error":  "해당 선거엔 후보자가 없습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status":    200,
			"candidate": list,
		})
	}
}

func VoterGetCandidateInfo(c *gin.Context) {
	candidateId, err := strconv.Atoi(c.Params.ByName("candidateid"))
	if err != nil {
		panic(err)
	}
	candidate := services.VoterGetCandidateInfo(candidateId)

	if candidate == (models.Candidate{}) {
		c.JSON(404, gin.H{
			"status": 404,
			"error":  "등록된 후보자가 없습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status":   200,
			"election": candidate,
		})
	}
}

func VoterVoting(c *gin.Context) {
	var request models.Voting
	c.ShouldBindJSON(&request)

	err := services.VoterVoting(request.ElectionId, request.CandidateId)

	if err != nil {
		c.JSON(404, gin.H{
			"status": 404,
			"error":  "투표 과정에 오류가 발생했습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status": 200,
		})
	}
}

func VoterGetElectionResult(c *gin.Context) {
	electionId, err := strconv.Atoi(c.Params.ByName("electionid"))
	if err != nil {
		panic(err)
	}

	electionId, electedCandidateId, all_vote, candidateList := services.VoterGetElectionResult(electionId)
	if err != nil {
		c.JSON(404, gin.H{
			"status": 404,
			"error":  "해당 투표 결과를 가져오지 못했습니다",
		})
	} else {
		c.JSON(200, gin.H{
			"status":               200,
			"election_id":          electionId,
			"elected_candidate_id": electedCandidateId,
			"all_vote":             all_vote,
			"candidate":            candidateList,
		})
	}
}
