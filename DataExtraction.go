package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"path"
	"strconv"
)

func main() {
	// set data directory that files will be exported to
	dataDir := "C:\\Users\\katjad\\Desktop\\Random\\ebirdData\\eBirdDataGo"
	// creating layers of Readers to extract the data from the compressed file without decompressing it
	filepath := "C:\\Users\\katjad\\Desktop\\Random\\ebirdData\\ebird_reference_dataset_v2016_western_hemisphere.tar\\ebird_reference_dataset_v2016_western_hemisphere\\ERD2016SS\\2003\\checklists.csv.gz"
	file, _ := os.Open(filepath)
	gzipReader, _ := gzip.NewReader(file)
	csvReader := csv.NewReader(gzipReader)
	// read the first line of the file containing just the names
	csvReader.Read()
	// initialize an empty map to hold a list of counts for total and species
	counts := map[[2]float32][2]int{}

	i := 0
	// run a for loop over the entire file
	for {
		//progress indicator after every 10000 checklists
		i++
		if i%10000 == 0 {
			fmt.Println(i)
		}
		l, err := csvReader.Read()
		// exit the loop when the end of file is reached
		if err == io.EOF {
			break
		}
		// only include primary checklists
		if l[18] == "1" {
			latfloat, errlat := strconv.ParseFloat(l[2], 32)
			longfloat, errlong := strconv.ParseFloat(l[3], 32)
			// skip locations with erroneous coordinates
			if errlat != nil || errlong != nil {
				continue
			}
			// round the latitude and longitude to the next 1-decimal to group locations in grid
			latitude := float32(math.Floor(latfloat*1) / 1)
			longitude := float32(math.Floor(longfloat*1) / 1)
			location := [2]float32{latitude, longitude}
			// add one to the total of observations in the square
			temp := counts[location]
			temp[0]++
			// This adds the count for a particular bird to the total at the location
			speciesint, errspecies := strconv.ParseInt(l[3915], 10, 32)
			if errspecies != nil {
				continue
			}
			// add 1 for every non-zero sighting
			if speciesint > 0 {
				temp[1]++
			}
			// update the map for the given location
			counts[location] = temp
		}
	}

	// check that the results are as expected
	fmt.Println(counts)
	fmt.Println(len(counts))

	// Create a csv file to export to
	resultFile, _ := os.Create(path.Join(dataDir, "red_breasted_2003.csv"))
	resultWriter := csv.NewWriter(resultFile)
	// iterate through every location and write the coordinates and values
	for location, value := range counts {
		lat := strconv.FormatFloat(float64(location[0]), 'E', -1, 32)
		long := strconv.FormatFloat(float64(location[1]), 'E', -1, 32)
		resultWriter.Write([]string{lat, long, strconv.Itoa(value[0]), strconv.Itoa(value[1])})
	}
	resultWriter.Flush()
	resultFile.Close()

}

/*
delete duplicates from shared checklists
assign each observation to nearest hotspot
output: for every month, number of observation at that hotspot, number of observations of species x
*/
