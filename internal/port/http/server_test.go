package http

import (
	"encoding/json"
	"github.com/dihmuzikien/smallurl/domain"
	"github.com/dihmuzikien/smallurl/domain/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func performRequestWithBody(r http.Handler, method, path, body string) *httptest.ResponseRecorder {
	reader := strings.NewReader(body)
	req, _ := http.NewRequest(method, path, reader)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHandleListUrl(t *testing.T) {
	t.Run("test list handler", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		retList := []domain.Url{
			{
				ID:          "test1",
				Destination: "https://google.com",
				Created:     time.Now().Add(-100 * time.Second),
			},
			{
				ID:          "test2",
				Destination: "https://yahoo.com",
				Created:     time.Now().Add(-200 * time.Second),
			},
			{
				ID:          "test3",
				Destination: "https://github.com",
				Created:     time.Now().Add(-300 * time.Second),
			},
		}

		m := mocks.NewMockUrlUseCase(ctrl)
		m.EXPECT().List(gomock.Any()).Return(retList, nil)
		type expectedResponse struct {
			ID          string `json:"id"`
			Destination string `json:"destination"`
		}
		sut, _ := New(m)
		w := performRequest(sut, http.MethodGet, "/v1")
		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("want %v got %v", http.StatusOK, resp.StatusCode)
		}
		var body []expectedResponse
		err := json.NewDecoder(resp.Body).Decode(&body)

		if err != nil {
			t.Errorf("unexpected error decoding response %v", err)
		}
		if len(body) != len(retList) {
			t.Errorf("response length not match expected result")
		}
	})
}
func TestHandleCreateUrl(t *testing.T) {
	t.Run("successful create", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		testcases := []struct {
			body        string
			generatedId string
			destination string
		}{
			{
				body:        `{"destination": "https://google.com"}`,
				generatedId: "generated-id",
				destination: "https://google.com",
			},
			{
				body:        `{"destination": "https://youtube.com"}`,
				generatedId: "another-generated-id",
				destination: "https://youtube.com",
			},
		}
		type expectedResponse struct {
			ID string `json:"id"`
		}

		for _, tc := range testcases {
			m := mocks.NewMockUrlUseCase(ctrl)
			m.EXPECT().
				Create(gomock.Any(), gomock.Eq(tc.destination)).
				Return(domain.Url{
					ID:          tc.generatedId,
					Destination: tc.destination,
					Created:     time.Now(),
				}, nil)

			sut, _ := New(m)
			w := performRequestWithBody(sut, http.MethodPost, "/v1", tc.body)
			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Errorf("want %v got %v", http.StatusOK, resp.StatusCode)
			}
			var respBody expectedResponse
			if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
				t.Errorf("unexpected error decoding response %v", err)
			}
			if respBody.ID != tc.generatedId {
				t.Errorf("want %v but got %v", tc.generatedId, respBody.ID)
			}
		}
	})
	t.Run("CreateUrlWithAlias invalid alias returns 400", func(t *testing.T) {
		testcases := []struct {
			body string
		}{
			{
				body: `{"alias":"","destination": "https://google.com"}`,
			},
			{
				body: `{"alias":"aa","destination": "https://google.com"}`,
			},
		}
		for _, tc := range testcases {
			ctrl := gomock.NewController(t)
			t.Cleanup(ctrl.Finish)
			m := mocks.NewMockUrlUseCase(ctrl)
			m.EXPECT().CreateWithId(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			sut, _ := New(m)
			w := performRequestWithBody(sut, http.MethodPost, "/v1/alias", tc.body)
			resp := w.Result()
			if resp.StatusCode != http.StatusBadRequest {
				t.Errorf("want %v got %v", http.StatusBadRequest, resp.StatusCode)
			}
		}
	})
}

func TestHandleCreateUrlWithAlias(t *testing.T) {
	t.Run("create URL with alias succeed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		t.Cleanup(ctrl.Finish)
		testcases := []struct {
			body        string
			alias       string
			destination string
		}{
			{
				body:        `{"alias":"test1","destination": "https://google.com"}`,
				alias:       "test1",
				destination: "https://google.com",
			},
			{
				body:        `{"alias":"test2","destination": "https://youtube.com"}`,
				alias:       "test2",
				destination: "https://youtube.com",
			},
		}
		type expectedResponse struct {
			ID string `json:"id"`
		}

		for _, tc := range testcases {
			m := mocks.NewMockUrlUseCase(ctrl)
			m.EXPECT().
				CreateWithId(gomock.Any(), gomock.Eq(tc.alias), gomock.Eq(tc.destination)).
				Return(domain.Url{
					ID:          tc.alias,
					Destination: tc.destination,
					Created:     time.Now(),
				}, nil)

			sut, _ := New(m)
			w := performRequestWithBody(sut, http.MethodPost, "/v1/alias", tc.body)
			resp := w.Result()
			if resp.StatusCode != http.StatusOK {
				t.Errorf("want %v got %v", http.StatusOK, resp.StatusCode)
			}
			var respBody expectedResponse
			if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
				t.Errorf("unexpected error decoding response %v", err)
			}
			if respBody.ID != tc.alias {
				t.Errorf("want %v got %v", tc.alias, respBody.ID)
			}
		}

	})

	t.Run("CreateUrlWithAlias invalid alias returns 400", func(t *testing.T) {
		testcases := []struct {
			body string
		}{
			{
				body: `{"alias":"","destination": "https://google.com"}`,
			},
			{
				body: `{"alias":"aa","destination": "https://google.com"}`,
			},
		}
		for _, tc := range testcases {
			ctrl := gomock.NewController(t)
			t.Cleanup(ctrl.Finish)
			m := mocks.NewMockUrlUseCase(ctrl)
			m.EXPECT().CreateWithId(gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
			sut, _ := New(m)
			w := performRequestWithBody(sut, http.MethodPost, "/v1/alias", tc.body)
			resp := w.Result()
			if resp.StatusCode != http.StatusBadRequest {
				t.Errorf("want %v got %v", http.StatusBadRequest, resp.StatusCode)
			}
		}
	})
}
