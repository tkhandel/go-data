package pandas

import "sort"

type IntSeries struct {
	data []int64
}

func NewIntSeries(data ...int64) IntSeries {
	return IntSeries{data: data}
}

func (i IntSeries) Append(elements ...int64) IntSeries {
	changed := i.Clone()
	changed.data = append(changed.data, elements...)
	return i
}

func (i IntSeries) Apply(oper func(int64) int64) IntSeries {
	changed := IntSeries{}
	for _, entry := range i.data {
		changed.data = append(changed.data, oper(entry))
	}
	return changed
}

func (i IntSeries) Clone() IntSeries {
	cloned := IntSeries{}
	cloned.data = append(cloned.data, i.data...)
	return cloned
}

func (i IntSeries) Sum() int64 {
	sum := int64(0)
	for _, entry := range i.data {
		sum += entry
	}
	return sum
}

func (i IntSeries) Size() int {
	return len(i.data)
}

func (i IntSeries) Avg() float64 {
	return float64(i.Sum()) / float64(i.Size())
}

func (i IntSeries) Sort() IntSeries {
	sorted := i.Clone()
	sort.Slice(sorted.data, func(i, j int) bool {
		return sorted.data[i] < sorted.data[j]
	})
	return sorted
}

func (i IntSeries) Max() (pos int, max int64) {
	for x, entry := range i.data {
		if max < entry {
			max = entry
			pos = x
		}
	}
	return pos, max
}

func (i IntSeries) Min() (pos int, min int64) {
	for x, entry := range i.data {
		if min > entry {
			min = entry
			pos = x
		}
	}
	return pos, min
}

func (i IntSeries) Index(index int) int64 {
	return i.data[index]
}

func (i IntSeries) Concat(x IntSeries) IntSeries {
	return NewIntSeries(append(i.data, x.data...)...)
}

func (i IntSeries) Subset(start int, end int) IntSeries {
	return NewIntSeries(i.data[start:end]...)
}

func (i IntSeries) PassThrough(filter TruthFilter) IntSeries {
	var data []int64
	for index, pass := range filter {
		if pass && index < i.Size() {
			data = append(data, i.Index(index))
		}
	}
	return NewIntSeries(data...)
}

func (i IntSeries) GreaterThan(value int64) (greater TruthFilter) {
	for _, entry := range i.data {
		greater = append(greater, entry > value)
	}
	return greater
}

func (i IntSeries) SmallerThan(value int64) (smaller TruthFilter) {
	for _, entry := range i.data {
		smaller = append(smaller, entry < value)
	}
	return smaller
}

func (i IntSeries) Find(val int64) int {
	for index, entry := range i.data {
		if entry == val {
			return index
		}
	}
	return -1
}

func (i IntSeries) Filter(accept func(int64) bool) (filter TruthFilter) {
	for _, val := range i.data {
		filter = append(filter, accept(val))
	}
	return filter
}
