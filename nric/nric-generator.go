package nric

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func firstChar(val int) string {
	switch val {
	case 0:
		return "S"
	case 1:
		return "T"
	case 2:
		return "F"
	case 3:
		return "G"
	default:
		return ""
	}
}

func lastChar1(val int) string {
	switch val {
	case 0:
		return "J"
	case 1:
		return "Z"
	case 2:
		return "I"
	case 3:
		return "H"
	case 4:
		return "G"
	case 5:
		return "F"
	case 6:
		return "E"
	case 7:
		return "D"
	case 8:
		return "C"
	case 9:
		return "B"
	case 10:
		return "A"
	default:
		return ""
	}
}

func lastChar2(val int) string {
	switch val {
	case 0:
		return "X"
	case 1:
		return "W"
	case 2:
		return "U"
	case 3:
		return "T"
	case 4:
		return "R"
	case 5:
		return "Q"
	case 6:
		return "P"
	case 7:
		return "N"
	case 8:
		return "M"
	case 9:
		return "L"
	case 10:
		return "K"
	default:
		return ""
	}
}

func GenerateNricMultiple(count int) string {
	nrics := []string{}
	for i := 0; i < count; i++ {
		nrics = append(nrics, GenerateNric())

	}
	nricsString := strings.Join(nrics, "\n")
	return nricsString
}

func GenerateNric() string {
	rand.Seed(time.Now().UnixNano())
	nric := []string{}

	// generate a random 1st letter
	a := firstChar(rand.Intn(4))
	nric = append(nric, a)

	for i := 0; i < 7; i++ {
		nric = append(nric, fmt.Sprintf("%d", rand.Intn(10)))
	}

	x := (2 * getInt(nric[1])) + (7 * getInt(nric[2])) + (6 * getInt(nric[3])) + (5 * getInt(nric[4])) + (4 * getInt(nric[5])) + (3 * getInt(nric[6])) + (2 * getInt(nric[7]))

	if nric[0] == "T" || nric[0] == "G" {
		x += 4
	}

	y := x % 11

	if nric[0] == "S" || nric[0] == "T" {
		nric = append(nric, lastChar1(y))
	} else if nric[0] == "F" || nric[0] == "G" {
		nric = append(nric, lastChar2(y))
	}
	nricString := strings.Join(nric, "")
	return nricString
}

func getInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
