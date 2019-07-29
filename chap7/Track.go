package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(t string) time.Duration {
	duration, e := time.ParseDuration(t)
	if e != nil {
		panic(e)
	}
	return duration
}

func printTrack(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	_ = tw.Flush()
}

type byArtist []*Track

func (s byArtist) Len() int {
	return len(s)
}

func (s byArtist) Less(i, j int) bool {
	return s[i].Artist < s[j].Artist
}

func (s byArtist) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	//sort.Sort(byArtist(tracks))
	//printTrack(tracks)
	//fmt.Println()
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTrack(tracks)
}
