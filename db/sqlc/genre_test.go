package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/template-go-server/util"
)

func createRandomGenre(t *testing.T) Genre {

	arg := CreateGenreParams{
		Name:        util.RandomString(4),
		Description: util.RandomText(),
	}

	genre, err := testQueries.CreateGenre(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, genre)

	require.Equal(t, arg.Name, genre.Name)
	require.Equal(t, arg.Description, genre.Description)

	require.NotZero(t, genre.ID)

	return genre
}

func TestCreateGenre(t *testing.T) {
	createRandomGenre(t)
}

func TestGetGenre(t *testing.T) {
	genre1 := createRandomGenre(t)
	genre2, err := testQueries.GetGenre(context.Background(), genre1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, genre2)

	require.Equal(t, genre1.ID, genre2.ID)
	require.Equal(t, genre1.Name, genre2.Name)
	require.Equal(t, genre1.Description, genre2.Description)

}

func TestUpdateGenre(t *testing.T) {
	genre1 := createRandomGenre(t)

	arg := UpdateGenreParams{
		ID:          genre1.ID,
		Name:        util.RandomString(4),
		Description: util.RandomText(),
	}

	genre2, err := testQueries.UpdateGenre(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, genre2)

	require.Equal(t, genre1.ID, genre2.ID)
	require.Equal(t, arg.Name, genre2.Name)
	require.Equal(t, arg.Description, genre2.Description)

}

func TestDeleteGenre(t *testing.T) {
	genre1 := createRandomGenre(t)

	err := testQueries.DeleteGenre(context.Background(), genre1.ID)
	require.NoError(t, err)

	genre2, err := testQueries.GetGenre(context.Background(), genre1.ID)

	require.Error(t, err)

	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, genre2)

}

func TestListGenres(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomGenre(t)
	}

	arg := ListGenresParams{
		Limit:  5,
		Offset: 5,
	}

	genres, err := testQueries.ListGenres(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, genres, 5)

	for _, genre := range genres {
		require.NotEmpty(t, genre)
	}
}
