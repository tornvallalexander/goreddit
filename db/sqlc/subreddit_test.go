package db

import (
	"context"
	"database/sql"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomSubreddit(t *testing.T) Subreddit {
	user := createRandomUser(t)
	require.NotEmpty(t, user)

	arg := CreateSubredditParams{
		Name:        faker.Username(),
		Moderator:   user.Username,
		Description: faker.Paragraph(),
	}

	subreddit, err := testQueries.CreateSubreddit(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, subreddit.Name)

	require.Equal(t, arg.Name, subreddit.Name)
	require.Equal(t, arg.Moderator, subreddit.Moderator)
	require.Equal(t, arg.Description, subreddit.Description)

	return subreddit
}

func TestCreateSubreddit(t *testing.T) {
	createRandomSubreddit(t)
}

func TestGetSubreddit(t *testing.T) {
	subreddit1 := createRandomSubreddit(t)
	require.NotEmpty(t, subreddit1)

	subreddit2, err := testQueries.GetSubreddit(context.Background(), subreddit1.Name)
	require.NoError(t, err)

	require.Equal(t, subreddit1.Name, subreddit2.Name)
	require.Equal(t, subreddit1.Moderator, subreddit2.Moderator)
	require.Equal(t, subreddit1.Description, subreddit2.Description)
	require.Equal(t, subreddit1.Followers, subreddit2.Followers)
	require.Equal(t, subreddit1.CreatedAt, subreddit2.CreatedAt)
}

func TestDeleteSubreddit(t *testing.T) {
	subreddit1 := createRandomSubreddit(t)
	require.NotEmpty(t, subreddit1)

	err := testQueries.DeleteSubreddit(context.Background(), subreddit1.Name)
	require.NoError(t, err)

	subreddit2, err := testQueries.GetSubreddit(context.Background(), subreddit1.Name)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, subreddit2)
}
