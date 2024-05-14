package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTracksSortByArtist(t *testing.T) {
	expectedArtistOrder := []string{"Alicia Keys", "Delilah", "Martin Solveig", "Moby"}
	sort.Sort(byArtist(tracks))
	for i, track := range tracks {
		assert.Equal(t, expectedArtistOrder[i], track.Artist)
	}
}

func TestTracksSortByYear(t *testing.T) {
	expectedYearOrder := []int{1992, 2007, 2011, 2012}
	sort.Sort(byYear(tracks))
	for i, track := range tracks {
		assert.Equal(t, expectedYearOrder[i], track.Year)
	}
}

func TestTracksSortByCustom(t *testing.T) {
	expectedTitleOrder := []string{"Go", "Go", "Go Ahead", "Ready 2 Go"}
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	for i, track := range tracks {
		assert.Equal(t, expectedTitleOrder[i], track.Title)
	}
}

func TestTracksSortByColumns(t *testing.T) {
	expectedTitleOrder := []string{"Go", "Go", "Go Ahead", "Ready 2 Go"}
	sort.Sort(byColumns{tracks, []string{"Title", "Year", "Artist"}})
	for i, track := range tracks {
		assert.Equal(t, expectedTitleOrder[i], track.Title)
	}
}

func TestTracksSortByColumnsStable(t *testing.T) {
	expectedTitleOrder := []string{"Go", "Go", "Go Ahead", "Ready 2 Go"}
	sort.Stable(customSort{tracks, func(x, y *Track) bool {
		return x.Artist < y.Artist
	}})
	sort.Stable(customSort{tracks, func(x, y *Track) bool {
		return x.Year < y.Year
	}})
	sort.Stable(customSort{tracks, func(x, y *Track) bool {
		return x.Title < y.Title
	}})
	for i, track := range tracks {
		assert.Equal(t, expectedTitleOrder[i], track.Title)
	}
}
