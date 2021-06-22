package server

import (
	"bytes"
	"encoding/json"
	"github.com/dihmuzikien/smallurl/domain/url/mock"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HttpPost(t *testing.T){
	t.Run("Successful put", func(t *testing.T){
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		db := mock.NewMockRepository(ctrl)
		id := "test-id"
		url := "http://localhost"
		db.EXPECT().Put(gomock.Any(), gomock.Eq(id), gomock.Eq(url)).Return(nil)
		sut, _ := NewServer(db)
		server := httptest.NewServer(sut)
		defer server.Close()
		v := struct {
			ID string  `json:"id"`
			Destination string `json:"destination"`
		} {
			ID: id,
			Destination: url,
		}
		data, _ := json.Marshal(v)
		resp, err := http.Post(server.URL + "/v1", "application/json", bytes.NewBuffer(data))
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}
		if resp.StatusCode != http.StatusCreated {
			t.Errorf("want 200 got %d", resp.StatusCode)
		}

	})

}
