package goapp

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestPut(t *testing.T){
	t.Run("single put", func(t *testing.T){
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

	})

}
