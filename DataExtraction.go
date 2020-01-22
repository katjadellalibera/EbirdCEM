package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	filepath := "C:\\Users\\katjad\\Desktop\\Random\\ebirdData\\ebird_reference_dataset_v2016_western_hemisphere.tar\\ebird_reference_dataset_v2016_western_hemisphere\\ERD2016SS\\2002\\checklists.csv.gz"
	file, _ := os.Open(filepath)
	gzipReader, _ := gzip.NewReader(file)
	csvReader := csv.NewReader(gzipReader)

	for i := 0; i < 10; i++ {
		l, _ := csvReader.Read()
		fmt.Println(l[0:2])
	}
}
