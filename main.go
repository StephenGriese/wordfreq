package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value >= p[j].Value }

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m := make(map[string]int)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m[cleanup(scanner.Text())]++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	p := make(PairList, len(m))

	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	// p is now sorted

	for i, k := range p {
		fmt.Printf("%v\t%v\t%v\n", i+1, k.Value, k.Key)
	}
}

func cleanup(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	s = strings.TrimPrefix(s, "(")
	s = strings.TrimPrefix(s, "“")
	s = strings.TrimPrefix(s, "‘")
	s = strings.TrimPrefix(s, "\"")

	s = strings.TrimSuffix(s, ")")
	s = strings.TrimSuffix(s, "”")
	s = strings.TrimSuffix(s, "\"")
	s = strings.TrimSuffix(s, "’")
	s = strings.TrimSuffix(s, ".")
	s = strings.TrimSuffix(s, ",")
	s = strings.TrimSuffix(s, ",")
	s = strings.TrimSuffix(s, "!")
	s = strings.TrimSuffix(s, "?")
	s = strings.TrimSuffix(s, ":")
	s = strings.TrimSuffix(s, ";")
	s = strings.TrimSuffix(s, "'s")
	s = strings.TrimSuffix(s, "’")
	return s
}
