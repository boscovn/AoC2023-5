package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type conRange struct {
	start  int
	end    int
	factor int
}

func convertList(list []string) []int {
	var intList []int
	for _, v := range list {
		num, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		intList = append(intList, num)
	}
	return intList
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	toConvert := convertList(strings.Split(scanner.Text(), " ")[1:])
	var converter []conRange
	if !scanner.Scan() {
		return
	}

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			for k, v := range toConvert {
				for _, r := range converter {
					if r.start <= v && v < r.end {
						toConvert[k] = v + r.factor
					}
				}
			}
			converter = converter[:0]
		}
		ranges := convertList(strings.Split(text, " "))
		if len(ranges) != 3 {
			continue
		}
		converter = append(converter, conRange{start: ranges[1], end: ranges[1] + ranges[2], factor: ranges[0] - ranges[1]})

	}
	for k, v := range toConvert {
		for _, r := range converter {
			if r.start <= v && v < r.end {
				toConvert[k] = v + r.factor
			}
		}
	}

	fmt.Println(slices.Min(toConvert))
}
