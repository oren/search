package main

import (
	"sort"
)

// 1:3, 2:2, 4:9 -> 4:9, 1:3, 2:2
func RankByValue(score map[int]int) PairList {
	pl := make(PairList, len(score))
	i := 0
	for k, v := range score {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
