// Read CSV to a map list with "encoding/csv"
package csv_map

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

type CSVMapReader struct {
	Reader *csv.Reader
	Heads []string    // head line
}

// new a CSVMapReader
func NewCSVMapReader(r io.Reader) *CSVMapReader {
	reader := csv.NewReader(r)
	heads, err := reader.Read()
	if err != nil {
		log.Fatal("read heads failed, err: %v", err)
	}
	return &CSVMapReader{
		Reader: reader,
		Heads: heads,
	}
}

// read one line to a map
func (r *CSVMapReader) Read() (recordMap map[string]string, err error) {
	recordList, err := r.Reader.Read()
	if err != nil {
		return map[string]string{}, err
	}

	recordMap = make(map[string]string)
	for index, head := range r.Heads {
		if _, exist := recordMap[head]; exist {
			return nil, fmt.Errorf("the same head: %v", head)
		}
		recordMap[head] = recordList[index]
	}

	return recordMap, nil
}

// read all data line to a map list
func (r *CSVMapReader) ReadAll() (records []map[string]string, err error) {
	for {
		record, err := r.Read()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
}