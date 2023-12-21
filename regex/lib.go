package regex

type typ int

const (
	normal typ = iota
	start
	split 
	match 
)

type state struct {
	t typ
	c rune
	ns1 *state
	ns2 *state
}

func (s *state) advance(c rune) []*state {
	switch s.t {
	case normal:
		if s.c == c {
			return []*state{s.ns1}
		}
		return nil
	case split:
		return append(s.ns1.advance(c), s.ns2.advance(c)...)
	case start:
		return s.ns1.advance(c)
	default:
		return nil
	}
}

func (s *state) match() bool {
	switch s.t {
	case match:
		return true
	case normal:
		return false
	default:
		if s.ns1 != nil && s.ns1.match() {
			return true
		}
		if s.ns2 != nil && s.ns2.match() {
			return true
		}
	}
	return false
}

func Visualize(pattern string, path string) error {
	return visualize(compile(pattern), path)
}

func Match(s string, p string) bool {
	start := compile(p)
	currs := make(map[*state]struct{})
	currs[start] = struct{}{}

	for _, c := range s {
		ncurrs := make(map[*state]struct{})
		for curr := range currs {
			for _, n := range curr.advance(c) {
				ncurrs[n] = struct{}{}
			}
		}
		if len(ncurrs) == 0 {
			return false
		}
		currs = ncurrs
	}
	for curr := range currs {
		if curr.match() {
			return true
		}
	}
	return false
}
