package pandas

import "sort"

type FloatSeries struct {
	data []float64
}

func NewFloatSeries(data ...float64) FloatSeries {
	return FloatSeries{data: data}
}

func (f FloatSeries) Append(elements ...float64) FloatSeries {
	changed := f.Clone()
	changed.data = append(changed.data, elements...)
	return f
}

func (f FloatSeries) Apply(oper func(float64) float64) FloatSeries {
	changed := FloatSeries{}
	for _, entry := range f.data {
		changed.data = append(changed.data, oper(entry))
	}
	return changed
}

func (f FloatSeries) Clone() FloatSeries {
	cloned := FloatSeries{}
	cloned.data = append(cloned.data, f.data...)
	return cloned
}

func (f FloatSeries) Sum() float64 {
	sum := float64(0)
	for _, entry := range f.data {
		sum += entry
	}
	return sum
}

func (f FloatSeries) Size() int {
	return len(f.data)
}

func (f FloatSeries) Avg() float64 {
	return f.Sum() / float64(f.Size())
}

func (f FloatSeries) Sort() FloatSeries {
	sorted := f.Clone()
	sort.Slice(sorted.data, func(i, j int) bool {
		return sorted.data[i] < sorted.data[j]
	})
	return sorted
}

func (f FloatSeries) Max() (pos int, max float64) {
	for x, entry := range f.data {
		if max < entry {
			max = entry
			pos = x
		}
	}
	return pos, max
}

func (f FloatSeries) Min() (pos int, min float64) {
	for x, entry := range f.data {
		if min > entry {
			min = entry
			pos = x
		}
	}
	return pos, min
}

func (f FloatSeries) Index(index int) float64 {
	return f.data[index]
}

func (f FloatSeries) Concat(x FloatSeries) FloatSeries {
	return NewFloatSeries(append(f.data, x.data...)...)
}

func (f FloatSeries) Subset(start int, end int) FloatSeries {
	return NewFloatSeries(f.data[start:end]...)
}

func (f FloatSeries) PassThrough(filter TruthFilter) FloatSeries {
	var data []float64
	for index, pass := range filter {
		if pass && index < f.Size() {
			data = append(data, f.Index(index))
		}
	}
	return NewFloatSeries(data...)
}

func (f FloatSeries) GreaterThan(value float64) (greater TruthFilter) {
	for _, entry := range f.data {
		greater = append(greater, entry > value)
	}
	return greater
}

func (f FloatSeries) SmallerThan(value float64) (smaller TruthFilter) {
	for _, entry := range f.data {
		smaller = append(smaller, entry < value)
	}
	return smaller
}

func (f FloatSeries) Find(val float64) int {
	for index, entry := range f.data {
		if entry == val {
			return index
		}
	}
	return -1
}

func (f FloatSeries) Filter(accept func(float64) bool) (filter TruthFilter) {
	for _, val := range f.data {
		filter = append(filter, accept(val))
	}
	return filter
}
