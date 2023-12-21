package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var output string
	var input string
	var groupSize int64
	flag.StringVar(&output, "output", "", "fichier de sortie, Stdout par défaut")
	flag.StringVar(&input, "input", "", "fichier d'entrée, Stdin par défaut")
	flag.Int64Var(&groupSize, "groupSize", 10, "taille d'un groupe de ligne, 10 par défaut")

	flag.Parse()

	reader, err := getReader(input)
	if err != nil {
		panic(err.Error())
	}

	writer, err := getWriter(output)
	if err != nil {
		panic(err.Error())
	}

	bilanChannel, err := readBilan(reader)
	if err != nil {
		panic(err.Error())
	}

	err = parquetWrite(bilanChannel, writer, groupSize)
	fmt.Println(err)
}

func getReader(filename string) (io.ReadCloser, error) {
	if filename == "" {
		return os.Stdin, nil
	}
	return os.Open(filename)
}

func getWriter(filename string) (io.WriteCloser, error) {
	if filename == "" {
		return os.Stdout, nil
	}
	return os.Create(filename)
}
