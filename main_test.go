package main

import (
	"github.com/stretchr/testify/assert"
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
