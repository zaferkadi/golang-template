package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/template-go-server/db/mock"
	db "github.com/template-go-server/db/sqlc"
	"github.com/template-go-server/util"
)

func TestGetGenreAPI(t *testing.T) {
	genre := randomGenre()

	testCases := []struct {
		name          string
		genreID       int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			genreID: genre.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetGenre(gomock.Any(), gomock.Eq(genre.ID)).
					Times(1).
					Return(genre, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchGenre(t, recorder.Body, genre)
			},
		},
		{
			name:    "NotFound",
			genreID: genre.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetGenre(gomock.Any(), gomock.Eq(genre.ID)).
					Times(1).
					Return(db.Genre{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "InvalidID",
			genreID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetGenre(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:    "InternalError",
			genreID: genre.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetGenre(gomock.Any(), gomock.Eq(genre.ID)).
					Times(1).
					Return(db.Genre{}, sql.ErrConnDone)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/genres/%d", tc.genreID)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, req)
			tc.checkResponse(t, recorder)
		})
	}

}

func TestCreateGenreAPI(t *testing.T) {
	genre := randomGenre()

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":        genre.Name,
				"description": genre.Description,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateGenreParams{
					Name:        genre.Name,
					Description: genre.Description,
				}

				store.EXPECT().CreateGenre(gomock.Any(), gomock.Eq(arg)).Times(1).Return(genre, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchGenre(t, recorder.Body, genre)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{},
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().
					CreateGenre(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/genres"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})
	}
}

func randomGenre() db.Genre {
	return db.Genre{
		ID:          int32(util.RandomInt(1, 1000)),
		Name:        util.RandomString(4),
		Description: util.RandomString(10),
	}
}

func requireBodyMatchGenre(t *testing.T, body *bytes.Buffer, genre db.Genre) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotGenre db.Genre
	err = json.Unmarshal(data, &gotGenre)
	require.NoError(t, err)
	require.Equal(t, genre, gotGenre)
}
