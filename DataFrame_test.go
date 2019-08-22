package pandas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDataFrame(t *testing.T) {
	df := testDF()

	assert.Equal(t, NewStringSeries(), df.StringColumn(col1))
	assert.Equal(t, NewStringSeries(), df.StringColumn(col2))
	assert.Equal(t, NewIntSeries(), df.IntColumn(col3))
	assert.Equal(t, NewFloatSeries(), df.FloatColumn(col4))
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
	assert.Equal(t, NewStringSeries(), df.StringColumn(col2))
	assert.Panics(t, func() {
		df.StringColumn(col3)
	})
}

func TestDataFrame_IntColumn(t *testing.T) {
	df := testDF()
	assert.Equal(t, NewIntSeries(), df.IntColumn(col3))
	assert.Panics(t, func() {
		df.IntColumn(col1)
	})
}

func TestDataFrame_FloatColumn(t *testing.T) {
	df := testDF()
	assert.Equal(t, NewFloatSeries(), df.FloatColumn(col4))
	assert.Panics(t, func() {
		df.FloatColumn(col1)
	})
}

const (
	col1 = "col1"
	col2 = "col2"
	col3 = "col3"
	col4 = "col4"
)

func testDF() DataFrame {
	return NewDataFrame(
		NewStringColumn(col1),
		NewStringColumn(col2),
		NewIntColumn(col3),
		NewFloatColumn(col4))
}
