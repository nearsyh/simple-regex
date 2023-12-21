package regex

import "strings"

type node struct {
	v string
	l *node
	r *node
}

func toAST(p string) *node {
	i := 0
	var r *node
	for i < len(p) {
		r, i = toASTInternal(p, i, r)
	}
	return r
}

func toASTInternal(p string, i int, prev *node) (*node, int) {
	if i == len(p) {
		return prev, len(p)
	}
	if p[i] == '+' || p[i] == '*' {
		// Highest priority
		if prev.r != nil {
			return &node{
				v: prev.v,
				l: prev.l,
				r: &node{
					v: p[i : i+1],
					l: prev.r,
				},
			}, i + 1
		} else {
			return &node{
				v: p[i : i+1],
				l: prev.l,
				r: nil,
			}, i + 1
		}
	}

	if p[i] == '?' {
		if prev.r != nil {
			return &node {
				v: prev.v,
				l: prev.l,
				r: &node {
					v: "?",
					l: prev.r,
				},
			}, i + 1
		}
		return &node {
			v: "?",
			l: prev,
		}, i + 1
	}

	if p[i] == '|' {
		l := prev
		r, ni := toASTInternal(p, i+1, nil)
		return &node {
			v: "|",
			l: l,
			r: r,
		}, ni
	}

	if p[i] == '(' {
		l := prev
		var (
			r *node
			ni int = i + 1
		)
		for {
			r, ni = toASTInternal(p, ni, r)
			if p[ni] == ')' {
				break
			}
		}
		r = &node {
			v: "",
			l: r,
		}
		if prev == nil {
			return r, ni + 1
		}
		return &node {
			v: ".",
			l: l,
			r: r,
		}, ni + 1
	}

	if prev == nil {
		return &node {
			v: p[i: i+1],
		}, i + 1
	} else {
		if prev.v == "|" {
			return &node{
				v: "|",
				l: prev.l,
				r: &node{
					v: ".",
					l: prev.r,
					r: &node{p[i : i+1], nil, nil},
				},
			}, i + 1
		}
		return &node{
			v: ".",
			l: prev,
			r: &node{p[i : i+1], nil, nil},
		}, i + 1
	}
}

func treeToPost(r *node) string {
	if r.l == nil && r.r == nil {
		return r.v
	}
	var sb strings.Builder
	if r.l != nil {
		sb.WriteString(treeToPost(r.l))
	}
	if r.r != nil {
		sb.WriteString(treeToPost(r.r))
	}
	sb.WriteString(r.v)
	return sb.String()
}

func toPost(p string) string {
	return treeToPost(toAST(p))
}
