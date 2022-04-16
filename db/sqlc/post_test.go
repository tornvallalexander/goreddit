package db

import (
	"context"
	"database/sql"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomPost(t *testing.T) Post {
	subreddit := createRandomSubreddit(t)
	require.NotEmpty(t, subreddit)

	user := createRandomUser(t)
	require.NotEmpty(t, user)

	arg := CreatePostParams{
		User:      user.Username,
		Title:     faker.Sentence(),
		Content:   faker.Paragraph(),
		Subreddit: subreddit.Name,
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.User, post.User)
	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.Content, post.Content)
	require.Equal(t, arg.Subreddit, post.Subreddit)

	require.NotEmpty(t, post.ID)

	return post
}

func createSubredditPost(t *testing.T, subreddit string) Post {
	user := createRandomUser(t)
	require.NotEmpty(t, user)

	arg := CreatePostParams{
		User:      user.Username,
		Title:     faker.Sentence(),
		Content:   faker.Paragraph(),
		Subreddit: subreddit,
	}

	post, err := testQueries.CreatePost(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.User, post.User)
	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.Content, post.Content)
	require.Equal(t, arg.Subreddit, post.Subreddit)

	require.NotEmpty(t, post.ID)

	return post
}

func TestCreatePost(t *testing.T) {
	createRandomPost(t)
}

func TestGetPost(t *testing.T) {
	post1 := createRandomPost(t)
	post2, err := testQueries.GetPost(context.Background(), post1.ID)
	require.NoError(t, err)

	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.User, post2.User)
	require.Equal(t, post1.Title, post2.Title)
	require.Equal(t, post1.Content, post2.Content)
	require.Equal(t, post1.CreatedAt, post2.CreatedAt)
	require.Equal(t, post1.Upvotes, post2.Upvotes)
}

func TestListPosts(t *testing.T) {
	subreddit := createRandomSubreddit(t)
	require.NotEmpty(t, subreddit)

	var lastPost Post
	for i := 0; i < 5; i++ {
		lastPost = createSubredditPost(t, subreddit.Name)
	}

	arg := ListPostsParams{
		Subreddit: lastPost.Subreddit,
		Limit:     5,
	}

	posts, err := testQueries.ListPosts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, posts)

	for _, post := range posts {
		require.NotEmpty(t, post)
		require.Equal(t, subreddit.Name, post.Subreddit)
	}
}

func TestDeletePost(t *testing.T) {
	post1 := createRandomPost(t)
	post2, err := testQueries.DeletePost(context.Background(), post1.ID)
	require.NoError(t, err)
	require.Equal(t, post1.ID, post2.ID)

	post3, err := testQueries.GetPost(context.Background(), post1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, post3)
}

func TestUpdatePost(t *testing.T) {
	post := createRandomPost(t)

	arg := UpdatePostParams{
		Title:   faker.Sentence(),
		Content: faker.Paragraph(),
		ID:      post.ID,
	}

	updatedPost, err := testQueries.UpdatePost(context.Background(), arg)
	require.NoError(t, err)

	post, err = testQueries.GetPost(context.Background(), post.ID)
	require.NoError(t, err)

	require.Equal(t, post.ID, updatedPost.ID)
	require.Equal(t, post.User, updatedPost.User)
	require.Equal(t, post.Title, updatedPost.Title)
	require.Equal(t, post.Content, updatedPost.Content)
	require.Equal(t, post.CreatedAt, updatedPost.CreatedAt)
	require.Equal(t, post.Upvotes, updatedPost.Upvotes)
}
