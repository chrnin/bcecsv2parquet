package main

import (
	"bytes"
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
	ass.Len(allBilans, 4)

	goodBilans := slices.DeleteFunc(allBilans, func(b Bilan) bool { return b.err != nil })
	ass.Len(goodBilans, 2)
}

func Test_writeNotEmptyParquet(t *testing.T) {
	ass := assert.New(t)
	reader, err := getReader("datasetTest.csv")
	ass.NoError(err)
	bilans, err := readBilan(reader)
	ass.NoError(err)
	var parquet bytes.Buffer
	err = parquetWrite(bilans, &parquet, 10)
	ass.NoError(err)

	ass.True(len(parquet.Bytes()) > 100)
}
