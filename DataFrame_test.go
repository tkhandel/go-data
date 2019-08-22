package pandas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDataFrame(t *testing.T) {
	df := testDF()

	assert.Equal(t, col1Val, df.StringColumn(col1))
	assert.Equal(t, col2Val, df.StringColumn(col2))
	assert.Equal(t, col3Val, df.IntColumn(col3))
	assert.Equal(t, col4Val, df.FloatColumn(col4))
}

func TestNewDataFrame_DuplicateColumns(t *testing.T) {
	assert.Panics(t, func() {
		NewDataFrame(NewStringColumn("col1"), NewFloatColumn("col1"))
	})
}

func TestDataFrame_Columns(t *testing.T) {
	df := testDF()
	exp := []Column{
		NewStringColumn(col1),
		NewStringColumn(col2),
		NewIntColumn(col3),
		NewFloatColumn(col4)}
	assert.ElementsMatch(t, exp, df.Columns())
}

func TestDataFrame_StringColumn(t *testing.T) {
	df := testDF()
	assert.Equal(t, col2Val, df.StringColumn(col2))
	assert.Panics(t, func() {
		df.StringColumn(col3)
	})
}

func TestDataFrame_IntColumn(t *testing.T) {
	df := testDF()
	assert.Equal(t, col3Val, df.IntColumn(col3))
	assert.Panics(t, func() {
		df.IntColumn(col1)
	})
}

func TestDataFrame_FloatColumn(t *testing.T) {
	df := testDF()
	assert.Equal(t, col4Val, df.FloatColumn(col4))
	assert.Panics(t, func() {
		df.FloatColumn(col1)
	})
}

func TestDataFrame_DropColumn_String(t *testing.T) {
	df := testDF()
	changed := df.DropColumn(col2)

	assert.Panics(t, func() {
		changed.StringColumn(col2)
	})
	assert.Equal(t, col2Val, df.StringColumn(col2))
}

func TestDataFrame_DropColumn_Int(t *testing.T) {
	df := testDF()
	changed := df.DropColumn(col3)

	assert.Panics(t, func() {
		changed.IntColumn(col3)
	})
	assert.Equal(t, col3Val, df.IntColumn(col3))
}

func TestDataFrame_DropColumn_Float(t *testing.T) {
	df := testDF()
	changed := df.DropColumn(col4)

	assert.Panics(t, func() {
		changed.FloatColumn(col4)
	})
	assert.Equal(t, col4Val, df.FloatColumn(col4))
}

const (
	col1 = "col1"
	col2 = "col2"
	col3 = "col3"
	col4 = "col4"
)

var (
	col1Val = NewStringSeries("one", "two")
	col2Val = NewStringSeries("three", "four")
	col3Val = NewIntSeries(5, 6, 7)
	col4Val = NewFloatSeries(8, 9, 10)
)

func testDF() DataFrame {
	df := NewDataFrame(
		NewStringColumn(col1),
		NewStringColumn(col2),
		NewIntColumn(col3),
		NewFloatColumn(col4))

	return df.
		SetStringColumn(col1, col1Val).
		SetStringColumn(col2, col2Val).
		SetIntColumn(col3, col3Val).
		SetFloatColumn(col4, col4Val)
}
