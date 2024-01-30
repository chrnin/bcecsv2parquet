package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"
)

func readBilan(reader io.ReadCloser) (chan Bilan, error) {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = ';'

	headers, err := csvReader.Read()
	if err != nil {
		return nil, err
	}
	if !checkHeaders(headers) {
		return nil, fmt.Errorf("l'entête du fichier source ne correspond pas au format attendu")
	}
	output := make(chan Bilan)
	go csvToBilan(output, csvReader, headers)
	return output, nil
}

func checkHeaders(headers []string) bool {
	return headers[0] == "siren" &&
		headers[1] == "date_cloture_exercice" &&
		headers[2] == "type_bilan" &&
		headers[3] == "confidentiality"
}

func csvToBilan(output chan Bilan, csvReader *csv.Reader, headers []string) {
	for {
		line, err := csvReader.Read()
		bilan := Bilan{}
		if err == io.EOF {
			break
		} else if err != nil {
			bilan.err = err
			output <- bilan
			continue
		}
		bilan.Liasse = make(map[string]int32)
		bilan.Siren = line[0]
		dateClotureExercice, err := time.Parse("20060102", line[1])
		if err != nil {
			bilan.err = err
			output <- bilan
			continue
		}
		bilan.DateClotureExercice = int32(dateClotureExercice.UnixNano() / int64(24*time.Hour))
		bilan.TypeBilan = line[2]
		for i := 4; i < len(headers); i++ {
			if line[i] != "" {
				val, err := strconv.ParseInt(line[i], 10, 32)
				bilan.Liasse[headers[i]] = int32(val)
				if err != nil {
					continue
				}
			}
		}
		output <- bilan
	}
	close(output)
}
