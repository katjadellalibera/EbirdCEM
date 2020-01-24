package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	// creating layers of Readers to extract the data from the compressed file without decompressing it
	filepath := "C:\\Users\\katjad\\Desktop\\Random\\ebirdData\\ebird_reference_dataset_v2016_western_hemisphere.tar\\ebird_reference_dataset_v2016_western_hemisphere\\ERD2016SS\\2002\\checklists.csv.gz"
	file, _ := os.Open(filepath)
	gzipReader, _ := gzip.NewReader(file)
	csvReader := csv.NewReader(gzipReader)
	//
	csvReader.Read()
	counts := map[[2]float32]int{}

	for i := 0; i < 1000; i++ {
		l, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if l[18] == "1" {
			latfloat, errlat := strconv.ParseFloat(l[2], 32)
			longfloat, errlong := strconv.ParseFloat(l[3], 32)
			if errlat != nil || errlong != nil {
				continue
			}
			latitude := float32(math.Floor(latfloat*10) / 10)
			longitude := float32(math.Floor(longfloat*10) / 10)
			counts[[2]float32{latitude, longitude}] += 1
		}
	}

	fmt.Println(counts)
	fmt.Println(len(counts))
	/*for i := 0; i < 1; i++ {
		l, _ := csvReader.Read()
		fmt.Println(l[0:20])
	}
	*/
}

/*
delete duplicates from shared checklists
assign each observation to nearest hotspot
output: for every month, number of observation at that hotspot, number of observations of species x
*/
