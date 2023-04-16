package cinema_test

import (
	"testing"

	"github.com/luisnquin/event-glance/pkg/cinema"
)

func TestSmoke(t *testing.T) {
	movies, err := cinema.NextReleases()
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	for _, movie := range movies {
		t.Log(movie)
	}
}
