package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test_getReader(t *testing.T) {
	ass := assert.New(t)
	_, err := getReader("datasetTest.csv")
	ass.NoError(err)
}

func Test_getReader_badFile(t *testing.T) {
	ass := assert.New(t)
	_, err := getReader("nofile")
	ass.Error(err)
}

func Test_readCSV_length(t *testing.T) {
	ass := assert.New(t)
	reader, err := getReader("datasetTest.csv")
	ass.NoError(err)
	bilans, err := readBilan(reader)
	ass.NoError(err)
	var allBilans []Bilan
	for bilan := range bilans {
		allBilans = append(allBilans, bilan)
	}
	ass.Len(allBilans, 5)

	goodBilans := slices.DeleteFunc(allBilans, func(b Bilan) bool { return b.err != nil })
	ass.Len(goodBilans, 4)
}

func Test_writeParquet_hash(t *testing.T) {
	ass := assert.New(t)
	reader, err := getReader("datasetTest.csv")
	ass.NoError(err)
	bilans, err := readBilan(reader)
	ass.NoError(err)
	var parquet bytes.Buffer
	parquetWrite(bilans, &parquet, 10)
	parquetBytes := parquet.Bytes()
	fmt.Println(parquetBytes)
}
