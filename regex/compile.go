package regex

type frag struct {
	s *state
	out []**state
}

func compile(p string) *state {
	stack := []frag{}

	push := func(f frag) {
		stack = append(stack, f)
	}
	pop := func() frag {
		r := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		return r
	}

	p = toPost(p)
	for _, c := range p {
		switch c {
		case '.':
			right := pop()
			left := pop()
			for _, o := range left.out {
				*o = right.s
			}
			push(frag {
				s: left.s,
				out: right.out,
			})
		case '?':
			f := pop()
			s := &state {
				t: split,
				ns1: f.s,
				ns2: nil,
			}
			push(frag {
				s: s,
				out: append(f.out, &s.ns2),
			})	
		case '+':
			f := pop()
			s := &state {
				t: split,
				ns1: f.s,
				ns2: nil,
			}
			for _, o := range f.out {
				*o = s
			}
			push(frag {
				s: f.s,
				out: []**state{&s.ns2},
			})
		case '*':
			f := pop()
			s := &state {
				t: split,
				ns1: f.s,
				ns2: nil,
			}
			for _, o := range f.out {
				*o = s
			}
			push(frag {
				s: s,
				out: []**state{&s.ns2},
			})	
		case '|':
			right := pop()
			left := pop()
			s := &state {
				t: split,
				ns1: left.s,
				ns2: right.s,
			}
			push(frag {
				s: s,
				out: append(left.out, right.out...),
			})
		default:
			s := &state{
				t: normal,
				c: c,
				ns1: nil,
				ns2: nil,
			}
			push(frag{
				s: s,
				out: []**state{&s.ns1},
			})
		}
	}

	start := &state {
		t: start,
		ns1: stack[0].s,
		ns2: nil,
	}
	match := &state{
		t: match,
	}
	for _, o := range stack[0].out {
		*o = match
	}
	return start
}