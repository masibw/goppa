package usecase

import (
	"fmt"
	"github.com/masibw/goppa/domain/loader"
	"log"
	"strconv"
)

func CompareWithPrev(prevFileName string, currentFileName string, l loader.Loader, border float64) (diff []string) {
	prevTestData, err := l.Load(prevFileName)
	if err != nil {
		log.Fatal(err)
	}

	currentTestData, err := l.Load(currentFileName)
	if err != nil {
		log.Fatal(err)
	}

	prevTestMap := make(map[string]float64, len(prevTestData))
	for _, prevTest := range prevTestData {
		prevTestMap[prevTest.Name] = prevTest.Elapsed
	}

	for _, testData := range currentTestData {
		if prevElapsed, exist := prevTestMap[testData.Name]; exist {
			if testData.IsSlowerThan(prevElapsed, border) {
				diff = append(diff, fmt.Sprintf("'%s', prev: %s, current: %s", testData.Name, strconv.FormatFloat(prevElapsed, 'f', -1, 64)+"s", strconv.FormatFloat(testData.Elapsed, 'f', -1, 64)+"s"))
			}
		}
	}
	return diff
}
