package loader

import (
	"bufio"
	"github.com/masibw/goppa/domain/valueobject"
	"github.com/masibw/goppa/pkg/extract"
	"os"
	"strconv"
	"strings"
)

type VerboseLoader struct {
}

func NewVerboseLoader() *VerboseLoader {
	return &VerboseLoader{}
}

func (l *VerboseLoader) Load(fileName string) (testData []valueobject.TestEvent, err error) {
	var file *os.File
	file, err = os.Open(fileName)
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "---") {
			var elapsedTime float64
			elapsedTime, err = strconv.ParseFloat(extract.Elapsed(line), 64)
			if err != nil {
				return
			}
			testName := extract.NameOfTest(line)

			testData = append(testData, valueobject.TestEvent{
				Name:    testName,
				Elapsed: elapsedTime,
			})
		}
	}
	return
}
