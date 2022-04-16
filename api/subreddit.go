package api

import (
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
		Name:        req.Name,
		Moderator:   req.Moderator,
		Description: req.Description,
	}

	subreddit, err := server.store.CreateSubreddit(ctx, arg)
	if err != nil {
		status, errRes := checkErr(err)
		ctx.JSON(status, errRes)
		return
	}

	ctx.JSON(http.StatusOK, subreddit)
}

type getSubredditRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (server *Server) getSubreddit(ctx *gin.Context) {
	var req getSubredditRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	subreddit, err := server.store.GetSubreddit(ctx, req.Name)
	if err != nil {
		status, errRes := checkErr(err)
		ctx.JSON(status, errRes)
		return
	}

	ctx.JSON(http.StatusOK, subreddit)
}

type deleteSubredditRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (server *Server) deleteSubreddit(ctx *gin.Context) {
	var req deleteSubredditRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := server.store.DeleteSubreddit(ctx, req.Name); err != nil {
		status, errRes := checkErr(err)
		ctx.JSON(status, errRes)
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
