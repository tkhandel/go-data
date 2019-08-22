package element

type Iterator struct {
	err  error
	pos  int
	data []Element
}

func NewIterator(data []interface{}) *Iterator {
	var elements []Element
	for _, d := range data {
		elements = append(elements, New(d))
	}
	return &Iterator{data: elements, pos: -1}
}

func (i *Iterator) HasNext() bool {
	return i.err != nil && i.pos+1 < len(i.data)
}

func (i *Iterator) NextInt() int {
	i.pos++
	val, err := i.data[i.pos].Int()
	if err != nil {
		i.err = err
	}
	return val
}

func (i *Iterator) NextString() string {
	i.pos++
	val, err := i.data[i.pos].String()
	if err != nil {
		i.err = err
	}
	return val
}

func (i *Iterator) NextFloat() float64 {
	i.pos++
	val, err := i.data[i.pos].Float()
	if err != nil {
		i.err = err
	}
	return val
}

func (i *Iterator) Error() error {
	return i.err
}
