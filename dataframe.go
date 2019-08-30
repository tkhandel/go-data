package godata

import (
	"github.com/tkhandel/go-data/element"
	"github.com/tkhandel/go-data/log"
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

func NewDataFrame(columns ...Column) (DataFrame, error) {
	df := DataFrame{
		columns:       make(map[string]Column),
		stringColumns: make(map[string]StringSeries),
		intColumns:    make(map[string]IntSeries),
		floatColumns:  make(map[string]FloatSeries),
	}

	for _, col := range columns {
		if _, ok := df.columns[col.name]; ok {
			err := Duplicate{What: "column", Value: col.name}
			log.Get().Errorf(err.Error())
			return DataFrame{}, err
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
			err := Unknown{What: "column type", Value: col.dType.String()}
			log.Get().Error(err.Error())
			return DataFrame{}, err
		}
	}
	return df, nil
}

func (df DataFrame) Columns() (columns []Column) {
	for _, col := range df.columns {
		columns = append(columns, col)
	}
	return columns
}

func (df DataFrame) StringColumn(colName string) (StringSeries, error) {
	col, ok := df.stringColumns[colName]
	if !ok {
		err := Unknown{What: "column", Value: colName}
		log.Get().Errorf(err.Error())
		return StringSeries{}, err
	}
	return col.Clone(), nil
}

func (df DataFrame) FloatColumn(colName string) (FloatSeries, error) {
	col, ok := df.floatColumns[colName]
	if !ok {
		err := Unknown{What: "column", Value: colName}
		log.Get().Errorf(err.Error())
		return FloatSeries{}, err
	}
	return col.Clone(), nil
}

func (df DataFrame) IntColumn(colName string) (IntSeries, error) {
	col, ok := df.intColumns[colName]
	if !ok {
		err := Unknown{What: "column", Value: colName}
		log.Get().Errorf(err.Error())
		return IntSeries{}, err
	}
	return col.Clone(), nil
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

func (df DataFrame) Clone() DataFrame {
	cloned, _ := NewDataFrame(df.Columns()...)

	for _, col := range df.columns {
		switch col.dType {
		case element.StringType:
			cloned.stringColumns[col.name] = df.stringColumns[col.name]
		case element.IntType:
			cloned.intColumns[col.name] = df.intColumns[col.name]
		case element.FloatType:
			cloned.floatColumns[col.name] = df.floatColumns[col.name]
		}
	}
	return cloned
}

func (df DataFrame) SetStringColumn(colName string, value StringSeries) DataFrame {
	changed := df.Clone()
	if _, ok := changed.columns[colName]; !ok {
		changed.columns[colName] = NewStringColumn(colName)
	} else if _, ok := changed.stringColumns[colName]; !ok {
		log.Get().Warnf("column name %s already exists and is not of type string", colName)
		return changed
	}
	changed.stringColumns[colName] = value.Clone()
	return changed
}

func (df DataFrame) SetIntColumn(colName string, value IntSeries) DataFrame {
	changed := df.Clone()
	if _, ok := changed.columns[colName]; !ok {
		changed.columns[colName] = NewIntColumn(colName)
	} else if _, ok := changed.intColumns[colName]; !ok {
		log.Get().Warnf("column name %s already exists and is not of type int", colName)
		return changed
	}
	changed.intColumns[colName] = value.Clone()
	return changed
}

func (df DataFrame) SetFloatColumn(colName string, value FloatSeries) DataFrame {
	changed := df.Clone()
	if _, ok := changed.columns[colName]; !ok {
		changed.columns[colName] = NewFloatColumn(colName)
	} else if _, ok := changed.floatColumns[colName]; !ok {
		log.Get().Warnf("column name %s already exists and is not of type float", colName)
		return changed
	}
	changed.floatColumns[colName] = value.Clone()
	return changed
}
