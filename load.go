package godata

import (
	"encoding/csv"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tkhandel/go-data/log"
	"io"
)

type CSV struct {
	HeadersPresent bool
}

func (c CSV) LoadCSV(rdr io.Reader) (DataFrame, error) {
	csvRdr := csv.NewReader(rdr)
	rows, err := csvRdr.ReadAll()
	if err != nil {
		readErr := ProcessingError{Err: errors.Wrap(err, "reading data rows")}
		log.Get().Error(readErr.Error())
		return DataFrame{}, readErr
	}

	if len(rows) == 0 {
		return NewDataFrame()
	}

	var columns []Column
	if c.HeadersPresent {
		for _, col := range rows[0] {
			columns = append(columns, NewStringColumn(col))
		}
		rows = rows[1:]
	} else {
		for i := range rows[0] {
			columns = append(columns, NewStringColumn(fmt.Sprintf("Column %d", i)))
		}
	}
	df, err := NewDataFrame(columns...)
	if err != nil {
		readErr := ProcessingError{Err: errors.Wrap(err, "creating data frame")}
		log.Get().Error(readErr.Error())
		return DataFrame{}, readErr
	}

	for j := range columns {
		var colVal []string
		for i := range rows {
			colVal = append(colVal, rows[i][j])
		}
		df = df.SetStringColumn(columns[j].name, NewStringSeries(colVal...))
	}
	return df, nil
}
