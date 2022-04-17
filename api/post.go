package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/tornvallalexander/goreddit/db/sqlc"
	"net/http"
)

type createPostRequest struct {
	User      string `json:"user" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Content   string `json:"content" binding:"required"`
	Subreddit string `json:"subreddit" binding:"required"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePostParams{
		User:      req.User,
		Title:     req.Title,
		Content:   req.Content,
		Subreddit: req.Subreddit,
	}

	post, err := server.store.CreatePost(ctx, arg)
	if err != nil {
		status, errRes := checkErr(err)
		ctx.JSON(status, errRes)
		return
	}

	ctx.JSON(http.StatusOK, post)
}

type getPostRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getPost(ctx *gin.Context) {
	var req getPostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	subreddit, err := server.store.GetPost(ctx, req.ID)
	if err != nil {
		status, errRes := checkErr(err)
		ctx.JSON(status, errRes)
		return
	}

	ctx.JSON(http.StatusOK, subreddit)
}

type deletePostRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) deletePost(ctx *gin.Context) {
	var req deletePostRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	post, err := server.store.DeletePost(ctx, req.ID)
	if err != nil {
		status, errRes := checkErr(err)
		ctx.JSON(status, errRes)
		return
	}

	ctx.JSON(http.StatusOK, post)
}

type listPostsRequest struct {
	Subreddit string `form:"subreddit" binding:"required"`
	Limit     int32  `form:"limit" binding:"required,min=1,max=10"`
}

func (server *Server) listPosts(ctx *gin.Context) {
	var req listPostsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPostsParams{
		Subreddit: req.Subreddit,
		Limit:     req.Limit,
	}

	posts, err := server.store.ListPosts(ctx, arg)
	if err != nil {
		status, errRes := checkErr(err)
		ctx.JSON(status, errRes)
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
