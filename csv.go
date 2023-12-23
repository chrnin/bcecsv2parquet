package main

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"time"
)

func readBilan(reader io.ReadCloser) (chan Bilan, error) {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'

	headers, err := csvReader.Read()
	if err != nil {
		return nil, nil
	}

	output := make(chan Bilan)
	go csvToBilan(output, csvReader, headers)
	return output, nil
}

func checkHeaders(headers []string) bool {
	return headers[0] == "Siren" &&
		headers[1] == "Date_cloture" &&
		headers[2] == "Type_bilan"
}

func csvToBilan(output chan Bilan, csvReader *csv.Reader, headers []string) {
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		bilan := Bilan{}
		bilan.Fields = make(map[string]int32)
		bilan.Siren = line[0]
		dateClotureExercice, err := time.Parse("20060102", line[1])
		if err != nil {
			bilan.err = err
			output <- bilan
			continue
		}
		bilan.DateClotureExercice = int32(dateClotureExercice.UnixNano() / int64(24*time.Hour))
		bilan.TypeBilan = line[2]
		for i := 3; i < len(headers); i++ {
			if line[i] != "" {
				val, err := strconv.ParseInt(line[i], 10, 32)
				bilan.Fields[headers[i]] = int32(val)
				if err != nil {
					continue
				}
			}
		}
		output <- bilan
	}
	close(output)
}
