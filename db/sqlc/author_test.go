package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/template-go-server/util"
)

func createRandomAuthor(t *testing.T) Author {
	arg := CreateAuthorParams{
		Name: util.RandomOwner(),
		Bio:  util.RandomText(),
	}

	author, err := testQueries.CreateAuthor(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, author)

	require.Equal(t, arg.Name, author.Name)
	require.Equal(t, arg.Bio, author.Bio)

	require.NotZero(t, author.ID)
	require.NotZero(t, author.CreatedAt)

	return author
}

func TestCreateAuthor(t *testing.T) {
	createRandomAuthor(t)
}

func TestGetAuthor(t *testing.T) {
	author1 := createRandomAuthor(t)
	author2, err := testQueries.GetAuthor(context.Background(), author1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, author2)

	require.Equal(t, author1.ID, author2.ID)
	require.Equal(t, author1.Name, author2.Name)
	require.Equal(t, author1.Bio, author2.Bio)

	require.WithinDuration(t, author1.CreatedAt.Time, author2.CreatedAt.Time, time.Second)
}

func TestUpdateAuthor(t *testing.T) {
	author1 := createRandomAuthor(t)

	arg := UpdateAuthorParams{
		ID:   author1.ID,
		Name: util.RandomOwner(),
	}

	author2, err := testQueries.UpdateAuthor(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, author2)

	require.Equal(t, author1.ID, author2.ID)
	require.Equal(t, arg.Name, author2.Name)
	require.Equal(t, author1.Bio, author2.Bio)

	require.WithinDuration(t, author1.CreatedAt.Time, author2.CreatedAt.Time, time.Second)
}

func TestDeleteAuthor(t *testing.T) {
	author1 := createRandomAuthor(t)

	err := testQueries.DeleteAuthor(context.Background(), author1.ID)
	require.NoError(t, err)

	author2, err := testQueries.GetAuthor(context.Background(), author1.ID)

	require.Error(t, err)

	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, author2)

}

func TestListAuthors(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAuthor(t)
	}

	arg := ListAuthorsParams{
		Limit:  5,
		Offset: 5,
	}

	authors, err := testQueries.ListAuthors(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, authors, 5)

	for _, author := range authors {
		require.NotEmpty(t, author)
	}
}
