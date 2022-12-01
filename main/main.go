package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)


func main() {
// open data.txt file and access contents
	bytes, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
// split string by new line
	stringArr := strings.Split(string(bytes), "\n")

	// create blank array which will be filled with data from file
	var intArr []float64
	// loop through data string and convert string into float number
		for _, value := range stringArr {
			if value != "" {
			result, err := strconv.ParseFloat(value, 64)
				if err != nil {
					panic(err) 	
		}
			// append array with float numbers of data
		intArr = append(intArr, result)
			}
	}

	// Prints calculations 
	fmt.Println("Average:", math.Round(Average(intArr)))
	fmt.Println("Median:", math.Round(Median(intArr)))
	fmt.Println("Variance:", int(math.Round(Variance(intArr, Average(intArr)))))
	fmt.Println("Standard Deviation", int(math.Round(StandardDiv(Variance(intArr, Average(intArr))))))

}

// calculates the mean by adding 
func Average(numbers []float64) float64 {
	var sum float64
	var result float64
	for _, num := range numbers {
		sum += num
		result = sum / float64(len(numbers))
	}
	return result
}
// Calculates the median by sorting numbers and calculating the middle
func Median(numbers []float64) float64 {
	// blank variable for result
	var median float64
	// make a blank array which will house the sorted array
	sortArr := make([]float64, len(numbers))
	// copy the contents of the old array into the nw
	copy(sortArr, numbers)
	// sort in ascending order
	sort.Float64s(sortArr)
	length := len(numbers)
	if length == 0 {
		return 0
	} else if length % 2 == 0 {
		median = (sortArr[length/2 - 1] + sortArr[length/2]) / 2
	} else {
		median = sortArr[length / 2]
	}
	return median
}
// Calculate the Variance by working out average 
func Variance(numbers []float64, average float64) float64 {
	var variance float64
	// loop through numbers and subtracct number by average/Then times number by itself/add to variable
	for _, value := range numbers {
	variance += (value - average) * (value - average)
	}
//divide number by the length of the array to calculate the Variance
result := variance / float64(len(numbers))

return result
}

func StandardDiv(variance float64) float64 {
// calculate standard deviation by finding the square root of the variance
	standardDiv := math.Sqrt(variance)
	return standardDiv
}