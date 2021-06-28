package inmemory

import (
	"context"
	"github.com/dihmuzikien/smallurl/domain"
	"testing"
	"time"
)

func TestFullDb(t *testing.T) {
	t.Run("put->get->delete behaves properly", func(t *testing.T) {
		testcase := struct {
			id      string
			dest    string
			created time.Time
		}{
			id:      "testID",
			dest:    "dest.com",
			created: time.Now(),
		}
		db := New()
		putErr := db.Put(context.TODO(), domain.Url{
			ID:          testcase.id,
			Destination: testcase.dest,
			Created:     testcase.created,
		})
		if putErr != nil {
			t.Error("unexpected error", putErr)
		}
		getRes, getErr := db.Get(context.TODO(), testcase.id)
		if getErr != nil {
			t.Error("unexpected error", getErr)
		}
		if getRes.Destination != testcase.dest {
			t.Errorf("want %v got %v", testcase.dest, getRes.Destination)
		}
		if getRes.Created != testcase.created {
			t.Errorf("want %v got %v", testcase.created, getRes.Created)
		}

		delErr := db.Delete(context.TODO(), testcase.id)
		if delErr != nil {
			t.Error("unexpected error", delErr)
		}
		_, endErr := db.Get(context.TODO(), testcase.id)
		if endErr != domain.RepoGetNotFoundError {
			t.Errorf("want %v got %v", domain.RepoGetNotFoundError, delErr)
		}
	})

	t.Run("multiple puts and list", func(t *testing.T) {
		testcases := []struct {
			id      string
			dest    string
			created time.Time
		}{
			{
				id:      "first",
				dest:    "first.com",
				created: time.Now(),
			},
			{
				id:      "second",
				dest:    "second.com",
				created: time.Now().Add(time.Minute * 300),
			},
		}
		testcaseContains := func(url domain.Url) bool{
			for _, v := range testcases {
				if url.ID == v.id && url.Destination == v.dest && url.Created == v.created {
					return true
				}
			}
			return false
		}
		db := New()
		for _, v := range testcases {
			putErr := db.Put(context.TODO(), domain.Url{
				ID:          v.id,
				Destination: v.dest,
				Created:     v.created,
			})
			if putErr != nil {
				t.Error("unexpected error", putErr)
			}
		}
		listRes, listErr := db.List(context.TODO())
		if listErr != nil {
			t.Error("unexpected error", listErr)
		}
		if len(listRes) != len(testcases) {
			t.Errorf("want %v got %v", len(testcases), len(listRes))
		}
		for _, v := range listRes {
			if !testcaseContains(v) {
				t.Errorf("expect %v in the response but it's not there", v)
			}
		}
	})

}

func TestDb_Get(t *testing.T) {
	t.Run("Get not found returns ErrNotfound", func(t *testing.T){
		db := New()
		_, err := db.Get(context.TODO(), "testID")
		if err != domain.RepoGetNotFoundError {
			t.Errorf("Expected ErrNotFound but got %v", err)
		}
	})

	t.Run("not found returns normally", func(t *testing.T){
		db := New()
		current := time.Now()
		dest := "destination.com"
		db.storage["testID"] = data{
			dest: dest,
			created: current,
		}
		res, err := db.Get(context.TODO(), "testID")
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		if res.Created != current {
			t.Errorf("want %v got %v", current, res.Created)
		}
		if res.Destination != dest {
			t.Errorf("want %v got %v", dest, res.Destination)
		}
	})

}
