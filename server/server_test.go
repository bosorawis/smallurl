package server

import (
	"github.com/dihmuzikien/smallurl/domain/url/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestServer_Put(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db := mock.NewMockRepository(ctrl)
	db.EXPECT().Get(gomock.Any(), gomock.Eq)
}
