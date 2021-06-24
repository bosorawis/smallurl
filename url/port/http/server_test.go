package http

import (
	"encoding/json"
	"github.com/dihmuzikien/smallurl/domain"
	"github.com/dihmuzikien/smallurl/domain/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandleListUrl(t *testing.T){
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	retList := []domain.Url {
		{
			ID: "test1",
			Destination: "https://google.com",
			Created: time.Now().Add(-100 * time.Second),
		},
		{
			ID: "test2",
			Destination: "https://yahoo.com",
			Created: time.Now().Add(-200 * time.Second),
		},
		{
			ID: "test3",
			Destination: "https://github.com",
			Created: time.Now().Add(-300 * time.Second),
		},
	}

	m := mocks.NewMockUrlUseCase(ctrl)
	m.EXPECT().List(gomock.Any()).Return(retList, nil)
	type expectedResponse struct {
		ID string `json:"id"`
		Destination string `json:"destination"`
	}
	sut, _ := New(m)
	req := httptest.NewRequest("GET", "/v1", nil)
	w := httptest.NewRecorder()
	handler := sut.handleListUrl()
	handler(w, req)
	resp := w.Result()
	if resp.StatusCode != http.StatusOK{
		t.Errorf("want %v got %v", http.StatusOK, resp.StatusCode)
	}
	var body []expectedResponse
	err := json.NewDecoder(resp.Body).Decode(&body)

	if err != nil {
		t.Errorf("unexpected error decoding response %v", err)
	}
	if len(body) != len(retList){
		t.Errorf("response length not match expected result")
	}
}
