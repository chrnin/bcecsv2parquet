package main

import (
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/writer"
	"io"
)

func parquetWrite(bilans chan Bilan, output io.Writer, rowGroupSize int64) error {
	pw, err := writer.NewParquetWriterFromWriter(output, new(Bilan), 2)
	if err != nil {
		return err
	}
	pw.RowGroupSize = rowGroupSize * 1024 * 1024
	pw.CompressionType = parquet.CompressionCodec_SNAPPY

	for bilan := range bilans {
		if bilan.err == nil {
			if err := pw.Write(bilan); err != nil {
				continue
			}
		}
	}
	err = pw.WriteStop()
	return err
}
