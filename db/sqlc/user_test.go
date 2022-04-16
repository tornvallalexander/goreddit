package db

import (
	"context"
	"database/sql"
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       faker.Username(),
		HashedPassword: faker.Password(),
		FullName:       faker.Name(),
		Email:          faker.Email(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	require.NotEmpty(t, user1)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.Equal(t, user1.PasswordChangedAt, user2.PasswordChangedAt)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	require.NotEmpty(t, user1)

	user2, err := testQueries.DeleteUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.Equal(t, user1.Username, user2.Username)

	deletedUser, err := testQueries.GetUser(context.Background(), user1.Username)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedUser)
}
