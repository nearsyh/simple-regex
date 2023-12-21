package regex

import (
	"fmt"
	"os"
	"strconv"

	"github.com/awalterschulze/gographviz"
)

type vnode struct {
	name string
}

func dfs(s *state, f func (*state) bool) {
	if f(s) {
		if s.ns1 != nil {
			dfs(s.ns1, f)
		}
		if s.ns2 != nil {
			dfs(s.ns2, f)
		}
	}
}

func (s *state) name() string {
	switch s.t {
	case normal:
		return "Normal"
	case split:
		return "Split"
	case match:
		return "Match"
	case start:
		return "Start"
	}
	return "Unknown"
}

func (s *state) nextStates() []*state {
	var r []*state
	if s.ns1 != nil {
		r = append(r, s.ns1)
	}
	if s.ns2 != nil {
		r = append(r, s.ns2)
	}
	return r
}

func (s *state) attrs() map[string]string {
	if s.t == normal {
		return map[string]string{"label" : string(s.c)}
	}
	return nil
}

func visualize(s *state, path string) error {
	g := gographviz.NewGraph()
	g.Name = strconv.Quote("Regex SM")
	g.Directed = true
	g.AddAttr(g.Name, string(gographviz.RankDir), "LR")

	nodes := make(map[*state]string)
	id := 0
	dfs(s, func(s *state) bool {
		if _, ok := nodes[s]; ok {
			return false;
		}
		id += 1
		nodes[s] = strconv.Quote(fmt.Sprintf("%d %s", id, s.name()))
		return true
	})
	for _, n := range nodes {
		g.AddNode(g.Name, n, nil)
	}
	for s, n := range nodes {
		for _, ns := range s.nextStates() {
			g.AddEdge(n, nodes[ns], true, ns.attrs())
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("regex: fail to visualize the graph: %w", err)
	}
	_, err = f.WriteString(g.String())
	if err != nil {
		return fmt.Errorf("regex: fail to dump the graph: %w", err)
	}
	return nil
}