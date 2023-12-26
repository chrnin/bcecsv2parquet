package main

type Bilan struct {
	Siren               string           `parquet:"name=siren, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	DateClotureExercice int32            `parquet:"name=date_cloture_exercice, type=INT32, convertedtype=DATE"`
	TypeBilan           string           `parquet:"name=type_bilan, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Liasse              map[string]int32 `parquet:"name=liasse, type=MAP, convertedtype=MAP, keytype=BYTE_ARRAY, keyconvertedtype=UTF8, valuetype=INT32"`
	err                 error
}
