package routers

import (
	"voting_system/app/controllers"

	"github.com/gin-gonic/gin"
)

func AddRoutesVoter(r *gin.RouterGroup) {
	r.GET("/voter/elections", controllers.VoterGetElectionsList)
	r.GET("/voter/elections/possible_List", controllers.VoterGetPossibleElectionsList)
	r.GET("/voter/elections/:electionid/candidates", controllers.VoterGetElectionCandidatesList)
	r.GET("/voter/candidates/:candidateid", controllers.VoterGetCandidateInfo)
	r.POST("/voter/elections/voting", controllers.VoterVoting)
	//r.GET("/voter/elections/result/:electionid", controllers.VoterGetElectionResult)
}
