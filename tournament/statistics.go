package tournament

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"text/tabwriter"
)

const (
	// WinPoints is the number of points a player gains for a win.
	WinPoints = 3
	// TiePoints is the number of points a player gains for a tie.
	TiePoints = 1
)

// PlayerStatistics represents the outcome of a player from a tournament.
// Points are handed out based on common soccer rules, i.e. 3 points for a win,
// 1 point for a tie, and 0 points for a loss.
type PlayerStatistics struct {
	PlayerName string
	Played     int
	Won        int
	Lost       int
	Tied       int
	Points     int
}

// Apply cumulates the delta statistics to the receiver statistics.
func (p *PlayerStatistics) Apply(deltaStatistics *PlayerStatistics) {
	p.Played += deltaStatistics.Played
	p.Won += deltaStatistics.Won
	p.Lost += deltaStatistics.Lost
	p.Tied += deltaStatistics.Tied
	p.Points += deltaStatistics.Points
}

// Result is the outcome of a tournament.
type Result []PlayerStatistics

func (t Result) Len() int      { return len(t) }
func (t Result) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t Result) Less(i, j int) bool {
	if t[i].Points == t[j].Points {
		if t[i].Won == t[j].Won {
			return t[i].Tied < t[j].Tied
		}
		return t[i].Won < t[j].Won
	}
	return t[i].Points < t[j].Points
}

func (t Result) String() string {
	const headFormat = "%8s\t%-16s\t%8s\t%8s\t%8s\t%8s\t%8s\n"
	const rowFormat = "%8d\t%-16s\t%8d\t%8d\t%8d\t%8d\t%8d\n"
	var sep16 = strings.Repeat("-", 16)
	var sep8 = strings.Repeat("-", 8)
	sort.Sort(sort.Reverse(t))
	buf := bytes.NewBufferString("")
	tw := new(tabwriter.Writer).Init(buf, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, headFormat, "Rank", "Player", "Points", "Games", "Won", "Lost", "Tied")
	fmt.Fprintf(tw, headFormat, sep8, sep16, sep8, sep8, sep8, sep8, sep8)
	for rank, stats := range t {
		fmt.Fprintf(tw, rowFormat, rank+1, stats.PlayerName, stats.Points, stats.Played, stats.Won,
			stats.Lost, stats.Tied)
	}
	tw.Flush()
	return buf.String()
}
