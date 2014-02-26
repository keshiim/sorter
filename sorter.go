package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

import (
	"sorter/algorithms/bubblesort"
	"sorter/algorithms/qsort"
)

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Read values", values)
		t1 := time.Now()

		switch *algorithm {
		case "qsort":
			{
				qsort.QuickSort(values)
			}
		case "bubblesort":
			{
				bubblesort.BubbleSort(values)
			}
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown orunsupported.")
		}

		t2 := time.Now()
		fmt.Println("The sorting process takes", t2.Sub(t1), "to complete.")
		err := writeValues(values, *outfile)
		fmt.Println("Write values", values)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Print("Failed to open the input file", infile)
		return nil, err
	}
	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}

	return
}

func writeValues(values []int, outfileName string) error {
	outfile, err := os.Create(outfileName)
	if err != nil {
		fmt.Print("Failed to create the outfile", outfileName)
		return err
	}
	defer outfile.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		outfile.WriteString(str + "\n")
	}
	return nil
}
