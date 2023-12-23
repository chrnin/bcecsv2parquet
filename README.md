# bcecsv2parquet
conversion de l'export complet des bilans du format CSV vers Parquet

## installation/exécution
### build
`go install github.com/chrnin/bcecsv2parquet@install`
### run
`$GOROOT/bin/bcecsv2parquet --input bce.csv --output bce.parquet`

## options
#### -groupSize int
taille d'un groupe de ligne, 10 par défaut (default 10)
#### -input string
fichier d'entrée, Stdin par défaut
#### -output string
fichier de sortie, Stdout par défaut
