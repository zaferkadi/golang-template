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

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/template-go-server/db/mock"
	db "github.com/template-go-server/db/sqlc"
	"github.com/template-go-server/util"
)

func TestGetAuthorAPI(t *testing.T) {
	author := randomAuthor()

	testCases := []struct {
		name          string
		authorID      int32
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:     "OK",
			authorID: author.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAuthor(gomock.Any(), gomock.Eq(author.ID)).
					Times(1).
					Return(author, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAuthor(t, recorder.Body, author)
			},
		},
		{
			name:     "NotFound",
			authorID: author.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAuthor(gomock.Any(), gomock.Eq(author.ID)).
					Times(1).
					Return(db.Author{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:     "InternalError",
			authorID: author.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAuthor(gomock.Any(), gomock.Eq(author.ID)).
					Times(1).
					Return(db.Author{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
		{
			name:     "InvalidID",
			authorID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAuthor(gomock.Any(), gomock.Any()).
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

			// start test server and send request
			server := NewTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/authors/%d", tc.authorID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

// func TestCreateAuthorAPI(t *testing.T) {
// 	author := randomAuthor()
// 	testCases := []struct {
// 		name          string
// 		body          gin.H
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				arg := db.CreateAuthorParams{
// 					Name: author.Name,
// 					Bio:  author.Bio,
// 				}
// 				store.EXPECT().
// 					CreateAuthor(gomock.Any(), gomock.Eq(arg)).
// 					Times(1).
// 					Return(author, nil)

// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 				requireBodyMatchAuthor(t, recorder.Body, author)
// 			},
// 		},
// 	}
// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {

// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			store := mockdb.NewMockStore(ctrl)
// 			tc.buildStubs(store)

// 			// start test server and send request
// 			server, _ := NewServer(store)
// 			recorder := httptest.NewRecorder()

// 			// Marshal body data to JSON
// 			data, err := json.Marshal(tc.body)
// 			require.NoError(t, err)

// 			url := "/authors"
// 			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
// 			require.NoError(t, err)

// 			server.router.ServeHTTP(recorder, request)
// 			tc.checkResponse(recorder)
// 		})
// 	}
// }

func randomAuthor() db.Author {
	return db.Author{
		ID:    int32(util.RandomInt(1, 1000)),
		Owner: util.RandomOwner(),
		Bio:   util.RandomText(),
	}
}

func requireBodyMatchAuthor(t *testing.T, body *bytes.Buffer, author db.Author) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAuthor db.Author
	err = json.Unmarshal(data, &gotAuthor)
	require.NoError(t, err)
	require.Equal(t, author, gotAuthor)
}
