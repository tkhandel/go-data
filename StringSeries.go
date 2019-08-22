package pandas

type StringSeries struct {
	data []string
}

func NewStringSeries(data ...string) StringSeries {
	return StringSeries{data: data}
}

func (s StringSeries) Append(elements ...string) StringSeries {
	changed := s.Clone()
	changed.data = append(changed.data, elements...)
	return changed
}

func (s StringSeries) Apply(oper func(string) string) StringSeries {
	changed := StringSeries{}
	for _, entry := range s.data {
		changed.data = append(changed.data, oper(entry))
	}
	return changed
}

func (s StringSeries) Clone() StringSeries {
	cloned := StringSeries{}
	cloned.data = append(cloned.data, s.data...)
	return cloned
}

func (s StringSeries) Size() int {
	return len(s.data)
}

func (s StringSeries) Index(pos int) string {
	return s.data[pos]
}

func (s StringSeries) Concat(x StringSeries) StringSeries {
	return NewStringSeries(append(s.data, x.data...)...)
}

func (s StringSeries) Subset(start int, end int) StringSeries {
	return NewStringSeries(s.data[start:end]...)
}

func (s StringSeries) PassThrough(filter TruthFilter) StringSeries {
	var data []string
	for index, pass := range filter {
		if pass && index < s.Size() {
			data = append(data, s.Index(index))
		}
	}
	return NewStringSeries(data...)
}

func (s StringSeries) Equal(str string) (notEqual TruthFilter) {
	for _, val := range s.data {
		notEqual = append(notEqual, val == str)
	}
	return notEqual
}

func (s StringSeries) NotEqual(str string) (notEqual TruthFilter) {
	for _, val := range s.data {
		notEqual = append(notEqual, val != str)
	}
	return notEqual
}

func (s StringSeries) Filter(accept func(string) bool) (filter TruthFilter) {
	for _, val := range s.data {
		filter = append(filter, accept(val))
	}
	return filter
}
