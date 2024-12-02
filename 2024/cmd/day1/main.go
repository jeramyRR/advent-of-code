package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := flag.String("i", "", "input file")
	flag.Parse()

	list1, list2 := ReadToLists(*input)
	distanceTotal := CalcDistance(list1, list2)
	fmt.Println("Total distance: ", distanceTotal)

	similarity := CalcSimilarity(list1, list2)
	fmt.Println("Similarity: ", similarity)
}

func ReadToLists(input string) ([]int, []int) {
	file, err := os.Open(input)
	if err != nil {
		log.Fatalf("failed to open file: %s, error: %s", input, err)
	}

	defer file.Close()

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		new_line := strings.Replace(line, "   ", " ", -1)
		items := strings.Split(new_line, " ")

		num1, err := strconv.Atoi(strings.TrimSpace(items[0]))
		if err != nil {
			log.Fatalf("failed to parse %s, error: %s", items[0], err)
		}
		list1 = append(list1, num1)

		num2, err := strconv.Atoi(strings.TrimSpace(items[1]))
		if err != nil {
			log.Fatalf("failed to parse %s, error: %s", items[1], err)
		}
		list2 = append(list2, num2)
	}

	return list1, list2
}

func CalcDistance(list1 []int, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)

	distanceTotal := 0

	for i := 0; i < len(list1); i++ {
		num1 := list1[i]
		num2 := list2[i]

		distanceTotal += Abs(num2 - num1)
	}

	return distanceTotal
}

func CalcSimilarity(list1 []int, list2 []int) int {
	sim := 0

	for _, n := range list1 {
		count := CountOccurences(n, list2)
		sim += count * n
	}

	return sim
}

func CountOccurences(v int, list []int) int {
	count := 0

	for _, n := range list {
		if v == n {
			count++
		}
	}

	return count
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
