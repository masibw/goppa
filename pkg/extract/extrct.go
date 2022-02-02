package extract

import "strings"

func NameOfTest(line string) (name string) {
	splitLine := strings.Fields(line)
	return splitLine[2]
}

func Elapsed(line string) (value string) {
	splitLine := strings.Fields(line)
	elapsedStr := splitLine[len(splitLine)-1]
	value = elapsedStr[1 : len(elapsedStr)-2]
	return
}
