package pandas

type TruthFilter []bool

func (t TruthFilter) Not() (not TruthFilter) {
	for _, val := range t {
		not = append(not, !val)
	}
	return not
}

func (t TruthFilter) And(addFilter TruthFilter) (and TruthFilter) {
	for i := range t {
		and = append(and, t[i] && addFilter[i])
	}
	return and
}

func (t TruthFilter) Or(addFilter TruthFilter) (or TruthFilter) {
	for i := range t {
		or = append(or, t[i] || addFilter[i])
	}
	return or
}
