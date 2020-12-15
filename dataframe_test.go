package godata

import (
	"github.com/stretchr/testify/require"
	"github.com/tkhandel/go-data/element"
	"testing"
)

func TestNewDataFrame(t *testing.T) {
	df := testDF()

	val1, err := df.StringColumn(col1)
	require.NoError(t, err)
	require.Equal(t, col1Val, val1)

	val2, err := df.StringColumn(col2)
	require.NoError(t, err)
	require.Equal(t, col2Val, val2)

	val3, err := df.IntColumn(col3)
	require.NoError(t, err)
	require.Equal(t, col3Val, val3)

	val4, err := df.FloatColumn(col4)
	require.NoError(t, err)
	require.Equal(t, col4Val, val4)
}

func TestNewDataFrame_DuplicateColumns(t *testing.T) {
	df, err := NewDataFrame(NewStringColumn("col1"), NewFloatColumn("col1"))
	require.IsType(t, Duplicate{}, err)
	require.Equal(t, "col1", err.(Duplicate).Value)
	require.Empty(t, df)
}

func TestNewDataFrame_InvalidColumnType(t *testing.T) {
	df, err := NewDataFrame(NewStringColumn("col1"), Column{
		name:  "foo",
		dType: element.Dtype(10),
	})
	require.IsType(t, Unknown{}, err)
	require.Equal(t, element.Dtype(10).String(), err.(Unknown).Value)
	require.Empty(t, df)
}

func TestDataFrame_Columns(t *testing.T) {
	df := testDF()
	exp := []Column{
		NewStringColumn(col1),
		NewStringColumn(col2),
		NewIntColumn(col3),
		NewFloatColumn(col4)}
	require.ElementsMatch(t, exp, df.Columns())
}

func TestDataFrame_StringColumn(t *testing.T) {
	df := testDF()
	val, err := df.StringColumn(col2)
	require.NoError(t, err)
	require.Equal(t, col2Val, val)

	val, err = df.StringColumn(col3)
	require.Error(t, err)
	require.IsType(t, Unknown{}, err)
	require.Empty(t, val)
}

func TestDataFrame_IntColumn(t *testing.T) {
	df := testDF()
	val, err := df.IntColumn(col3)
	require.NoError(t, err)
	require.Equal(t, col3Val, val)

	val, err = df.IntColumn(col1)
	require.Error(t, err)
	require.IsType(t, Unknown{}, err)
	require.Empty(t, val)
}

func TestDataFrame_FloatColumn(t *testing.T) {
	df := testDF()
	val, err := df.FloatColumn(col4)
	require.NoError(t, err)
	require.Equal(t, col4Val, val)

	val, err = df.FloatColumn(col1)
	require.Error(t, err)
	require.IsType(t, Unknown{}, err)
	require.Empty(t, val)
}

func TestDataFrame_DropColumn_String(t *testing.T) {
	df := testDF()
	changed := df.DropColumn(col2)

	val, err := changed.StringColumn(col2)
	require.Error(t, err)
	require.IsType(t, Unknown{}, err)
	require.Empty(t, val)

	val, err = df.StringColumn(col2)
	require.NoError(t, err)
	require.Equal(t, col2Val, val)
}

func TestDataFrame_DropColumn_Int(t *testing.T) {
	df := testDF()
	changed := df.DropColumn(col3)

	val, err := changed.IntColumn(col3)
	require.Error(t, err)
	require.IsType(t, Unknown{}, err)
	require.Empty(t, val)

	val, err = df.IntColumn(col3)
	require.NoError(t, err)
	require.Equal(t, col3Val, val)
}

func TestDataFrame_DropColumn_Float(t *testing.T) {
	df := testDF()
	changed := df.DropColumn(col4)

	val, err := changed.FloatColumn(col4)
	require.Error(t, err)
	require.IsType(t, Unknown{}, err)
	require.Empty(t, val)

	val, err = df.FloatColumn(col4)
	require.NoError(t, err)
	require.Equal(t, col4Val, val)
}

func TestDataFrame_SetIntColumn_ReplaceColumnValue(t *testing.T) {
	df := testDF()
	newVal := NewIntSeries(11, 12, 13)
	changed, err := df.SetIntColumn(col3, newVal)
	require.NoError(t, err)

	val, err := changed.IntColumn(col3)
	require.NoError(t, err)
	require.Equal(t, newVal, val)

	val, err = df.IntColumn(col3)
	require.NoError(t, err)
	require.Equal(t, col3Val, val)
}

func TestDataFrame_SetStringColumn_ReplaceColumnValue(t *testing.T) {
	df := testDF()
	newVal := NewStringSeries("five", "six")
	changed, err := df.SetStringColumn(col2, newVal)
	require.NoError(t, err)

	val, err := changed.StringColumn(col2)
	require.NoError(t, err)
	require.Equal(t, newVal, val)

	val, err = df.StringColumn(col2)
	require.NoError(t, err)
	require.Equal(t, col2Val, val)
}

func TestDataFrame_SetFloatColumn_ReplaceColumnValue(t *testing.T) {
	df := testDF()
	newVal := NewFloatSeries(11, 12, 13)
	changed, err := df.SetFloatColumn(col4, newVal)
	require.NoError(t, err)

	val, err := changed.FloatColumn(col4)
	require.NoError(t, err)
	require.Equal(t, newVal, val)

	val, err = df.FloatColumn(col4)
	require.NoError(t, err)
	require.Equal(t, col4Val, val)
}

func TestDataFrame_SetFloatColumn_DuplicateColumn(t *testing.T) {
	df := testDF()
	newVal := NewFloatSeries(11, 12, 13)
	_, err := df.SetFloatColumn(col3, newVal)
	require.Error(t, err)
	require.IsType(t, Duplicate{}, err)

	val, err := df.FloatColumn(col3)
	require.NoError(t, err)
	require.Equal(t, col3Val, val)
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
	df, _ := NewDataFrame(
		NewStringColumn(col1),
		NewStringColumn(col2),
		NewIntColumn(col3),
		NewFloatColumn(col4))

	df, _ = df.SetStringColumn(col1, col1Val)
	df, _ = df.SetStringColumn(col2, col2Val)
	df, _ = df.SetIntColumn(col3, col3Val)
	df, _ = df.SetFloatColumn(col4, col4Val)
	return df
}
