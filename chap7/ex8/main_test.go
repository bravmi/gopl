package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTracksSortByArtist(t *testing.T) {
	tracks := []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	want := []*Track{
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
	}
	sort.Sort(byArtist(tracks))
	for i, track := range tracks {
		assert.Equal(t, want[i], track)
	}
}

func TestTracksSortByYear(t *testing.T) {
	tracks := []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	want := []*Track{
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	}
	sort.Sort(byYear(tracks))
	for i, track := range tracks {
		assert.Equal(t, want[i], track)
	}
}

func TestTracksSortByCustom(t *testing.T) {
	tracks := []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	want := []*Track{
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
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
		assert.Equal(t, want[i], track)
	}
}

func TestTracksSortByColumns(t *testing.T) {
	tracks := []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	want := []*Track{
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	sort.Sort(byColumns{tracks, []string{"Title", "Year", "Artist"}})
	for i, track := range tracks {
		assert.Equal(t, want[i], track)
	}
}

func TestTracksSortByColumnsStable(t *testing.T) {
	tracks := []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
	want := []*Track{
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
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
		assert.Equal(t, want[i], track)
	}
}
