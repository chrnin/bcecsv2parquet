package main

import (
	"fmt"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
	"io"
)

func parquetWrite(bilans chan Bilan, output io.WriteCloser, rowGroupSize int64) error {
	pw, err := writer.NewParquetWriterFromWriter(output, new(Bilan), 2)
	if err != nil {
		return err
	}
	pw.RowGroupSize = rowGroupSize * 1024 * 1024 //10M
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	for bilan := range bilans {
		if err := pw.Write(bilan); err != nil {
			fmt.Println(err)
			continue
		}
	}
	err = pw.WriteStop()
	if err != nil {
		return err
	}

	return output.Close()
}
