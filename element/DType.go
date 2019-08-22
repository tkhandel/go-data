package element

type Dtype int

const (
	_ Dtype = iota
	IntType
	StringType
	FloatType
)

func (d Dtype) String() string {
	switch d {
	case IntType:
		return "Integer"
	case StringType:
		return "String"
	case FloatType:
		return "Float"
	}
	return ""
}
