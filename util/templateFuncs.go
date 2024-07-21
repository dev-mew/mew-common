package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const layout = "2006-01-02"

func Plus(x, y int) int {
	return x + y
}

func ContainsAny(str string, substrings []string) bool {
	lowerStr := strings.ToLower(str)
	for _, substring := range substrings {
		if strings.Contains(lowerStr, strings.ToLower(substring)) {
			return true
		}
	}
	return false
}

func FormatDateMMMxx(ds time.Time) string {
	return ds.Format("02 Jan 2006")
}

func YearsList() []int {
	today := time.Now()
	currentYear := today.Year()
	startYear := currentYear - 18
	endYear := currentYear - 100
	years := make([]int, 0, startYear-endYear+1)
	for i := startYear; i >= endYear; i-- {
		years = append(years, i)
	}
	return years
}

func IsValidDateOfBirth(day, month, year string) (string, bool) {
	_day, err := strconv.Atoi(day)
	if err != nil {
		return "", false
	}
	_year, err := strconv.Atoi(year)
	if err != nil {
		return "", false
	}

	monthMap := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}

	monthInt, ok := monthMap[month]
	if !ok {
		return "", false
	}

	daysInMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	isLeapYear := _year%4 == 0 && (_year%100 != 0 || _year%400 == 0)
	if isLeapYear && monthInt == 2 {
		daysInMonth[1] = 29
	}

	if _day < 1 || _day > daysInMonth[monthInt-1] {
		return "", false
	}

	birthDate := fmt.Sprintf("%d-%02d-%02d", _year, monthInt, _day)

	_, err = time.Parse(layout, birthDate)
	if err != nil {
		return "", false
	}

	return birthDate, true
}
