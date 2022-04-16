package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/tornvallalexander/goreddit/db/sqlc"
	"net/http"
)

type createSubredditRequest struct {
	Name        string `json:"name"`
	Moderator   string `json:"moderator"`
	Description string `json:"description"`
}

func (server *Server) createSubreddit(ctx *gin.Context) {
	var req createSubredditRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSubredditParams{
		Name:      req.Name,
		Moderator: req.Moderator,
		Description: sql.NullString{
			String: req.Description,
			Valid:  true,
		},
	}

	subreddit, err := server.store.CreateSubreddit(ctx, arg)
	if err != nil {
		status, errRes := checkErr(err)
		ctx.JSON(status, errRes)
		return
	}

	ctx.JSON(http.StatusOK, subreddit)
}
