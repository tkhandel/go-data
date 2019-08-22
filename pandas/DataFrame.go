package pandas

import (
	"github.com/pkg/errors"
	"github.com/tkhandel/go-data/element"
)

type DataFrame struct {
	columns       map[string]Column
	stringColumns map[string]StringSeries
	intColumns    map[string]IntSeries
	floatColumns  map[string]FloatSeries
}

type Column struct {
	name  string
	dType element.Dtype
}

func NewStringColumn(name string) Column {
	return Column{
		name:  name,
		dType: element.StringType,
	}
}

func NewIntColumn(name string) Column {
	return Column{
		name:  name,
		dType: element.IntType,
	}
}

func NewFloatColumn(name string) Column {
	return Column{
		name:  name,
		dType: element.FloatType,
	}
}

func NewDataFrame(columns ...Column) DataFrame {
	df := DataFrame{}
	for _, col := range columns {
		if _, ok := df.columns[col.name]; ok {
			panic(errors.Errorf("duplicate column name: %s", col.name))
		}
		df.columns[col.name] = col

		switch col.dType {
		case element.StringType:
			df.stringColumns[col.name] = NewStringSeries()
		case element.IntType:
			df.intColumns[col.name] = NewIntSeries()
		case element.FloatType:
			df.floatColumns[col.name] = NewFloatSeries()
		default:
			panic(errors.Errorf("unknown column type: %s for column: %s", col.dType.String(), col.name))
		}
	}
	return df
}

func (df DataFrame) Columns() (columns []Column) {
	for _, col := range df.columns {
		columns = append(columns, col)
	}
	return columns
}

func (df DataFrame) StringColumn(colName string) StringSeries {
	col, ok := df.stringColumns[colName]
	if !ok {
		panic(errors.Errorf("string column named %s not found", colName))
	}
	return col
}

func (df DataFrame) FloatColumn(colName string) FloatSeries {
	col, ok := df.floatColumns[colName]
	if !ok {
		panic(errors.Errorf("float column named %s not found", colName))
	}
	return col
}

func (df DataFrame) IntColumn(colName string) IntSeries {
	col, ok := df.intColumns[colName]
	if !ok {
		panic(errors.Errorf("int column named %s not found", colName))
	}
	return col
}

func (df DataFrame) DropColumn(name string) DataFrame {
	changed := df.Clone()

	delete(changed.columns, name)

	// The column name cannot be duplicated, so the delete will actually work only on one of them
	delete(changed.stringColumns, name)
	delete(changed.intColumns, name)
	delete(changed.floatColumns, name)

	return changed
}

func (df DataFrame) Clone() (cloned DataFrame) {
	for _, col := range df.columns {
		cloned.columns[col.name] = col

		switch col.dType {
		case element.StringType:
			cloned.stringColumns[col.name] = df.StringColumn(col.name)
		case element.IntType:
			cloned.intColumns[col.name] = df.IntColumn(col.name)
		case element.FloatType:
			cloned.floatColumns[col.name] = df.FloatColumn(col.name)
		}
	}
	return cloned
}

func (df DataFrame) SetStringColumn(colName string, value StringSeries) DataFrame {
	changed := df.Clone()
	changed.stringColumns[colName] = value.Clone()
	return df
}

func (df DataFrame) SetIntColumn(colName string, value IntSeries) DataFrame {
	changed := df.Clone()
	changed.intColumns[colName] = value.Clone()
	return df
}

func (df DataFrame) SetFloatColumn(colName string, value FloatSeries) DataFrame {
	changed := df.Clone()
	changed.floatColumns[colName] = value.Clone()
	return df
}
