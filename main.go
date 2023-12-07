package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type conRange struct {
	start  int
	end    int
	factor int
}
type valRange struct {
	start int
	end   int
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

func getNewRanges(ranges []valRange, converter []conRange) []valRange {
	var newRanges []valRange
	sort.Slice(converter, func(i, j int) bool { return converter[i].start < converter[j].start })
	for _, r := range ranges {
		cur := r.start
		for _, c := range converter {
			if c.end < r.start {
				continue
			}
			if r.end < c.start {
				break
			}
			if c.start <= cur {
				if r.end <= c.end {
					newRanges = append(newRanges, valRange{start: cur + c.factor, end: r.end + c.factor})
					cur = r.end
					break
				} else {
					newRanges = append(newRanges, valRange{start: cur + c.factor, end: c.end + c.factor})
					cur = c.end + 1
				}
			} else {
				if r.end <= c.end {
					newRanges = append(newRanges, valRange{start: cur, end: c.start - 1}, valRange{start: c.start + c.factor, end: r.end + c.factor})
					cur = r.end
					break
				} else {
					newRanges = append(newRanges, valRange{start: cur, end: c.start - 1}, valRange{start: c.start + c.factor, end: c.end + c.factor})
					cur = c.end + 1
				}
			}

		}
		if cur < r.end {
			newRanges = append(newRanges, valRange{start: cur, end: r.end})
		}
	}
	return newRanges
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	toConvert := convertList(strings.Split(scanner.Text(), " ")[1:])
	var valRanges []valRange
	for i := 0; i < len(toConvert); i += 2 {
		valRanges = append(valRanges, valRange{start: toConvert[i], end: toConvert[i+1] + toConvert[i] - 1})
	}
	var converter []conRange
	if !scanner.Scan() {
		return
	}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			valRanges = getNewRanges(valRanges, converter)
			converter = converter[:0]
		}
		ranges := convertList(strings.Split(text, " "))
		if len(ranges) != 3 {
			continue
		}
		converter = append(converter, conRange{start: ranges[1], end: ranges[1] + ranges[2] - 1, factor: ranges[0] - ranges[1]})

	}
	valRanges = getNewRanges(valRanges, converter)
	if len(valRanges) == 0 {
		return
	}
	minValue := valRanges[0].start
	for _, v := range valRanges {
		minValue = min(minValue, v.start)
	}
	fmt.Println(minValue)
}
