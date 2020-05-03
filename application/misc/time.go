package misc

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Month struct {
	JanuaryStart   time.Time
	JanuaryEnd     time.Time
	FebruaryStart  time.Time
	FebruaryEnd    time.Time
	MarchStart     time.Time
	MarchEnd       time.Time
	AprilStart     time.Time
	AprilEnd       time.Time
	MayStart       time.Time
	MayEnd         time.Time
	JuneStart      time.Time
	JuneEnd        time.Time
	JulyStart      time.Time
	JulyEnd        time.Time
	AugustStart    time.Time
	AugustEnd      time.Time
	SeptemberStart time.Time
	SeptemberEnd   time.Time
	OctoberStart   time.Time
	OctoberEnd     time.Time
	NovemberStart  time.Time
	NovemberEnd    time.Time
	DecemberStart  time.Time
	DecemberEnd    time.Time
}

type Day struct {
	Day1     string
	Day2     string
	Day3     string
	Day4     string
	Day5     string
	Day6     string
	Day7     string
	Day8     string
	Day9     string
	Day10    string
	Day11    string
	Day12    string
	Day13    string
	Day14    string
	Day15    string
	Day16    string
	Day17    string
	Day18    string
	Day19    string
	Day20    string
	Day21    string
	Day22    string
	Day23    string
	Day24    string
	Day25    string
	Day26    string
	Day27    string
	Day28    string
	Day29    string
	Day30    string
	Day31    string
	DayEnd1  string
	DayEnd2  string
	DayEnd3  string
	DayEnd4  string
	DayEnd5  string
	DayEnd6  string
	DayEnd7  string
	DayEnd8  string
	DayEnd9  string
	DayEnd10 string
	DayEnd11 string
	DayEnd12 string
	DayEnd13 string
	DayEnd14 string
	DayEnd15 string
	DayEnd16 string
	DayEnd17 string
	DayEnd18 string
	DayEnd19 string
	DayEnd20 string
	DayEnd21 string
	DayEnd22 string
	DayEnd23 string
	DayEnd24 string
	DayEnd25 string
	DayEnd26 string
	DayEnd27 string
	DayEnd28 string
	DayEnd29 string
	DayEnd30 string
	DayEnd31 string
}

func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func isLeap(year int) bool {
	return year%400 == 0 || year%4 == 0 && year%100 != 0
}

func MonthHelper(year int) Month {
	layoutISO := "2006-01-02"
	yearData := strconv.Itoa(year)
	month := Month{}
	var februaryEnd time.Time

	if isLeap(year) {
		februaryEnd, _ = time.Parse(layoutISO, fmt.Sprintf("%s-02-29", yearData))
	} else {
		februaryEnd, _ = time.Parse(layoutISO, fmt.Sprintf("%s-02-28", yearData))
	}

	januaryStart, _ := time.Parse(layoutISO, fmt.Sprintf("%s-01-01", yearData))
	januaryEnd, _ := time.Parse(layoutISO, fmt.Sprintf("%s-01-31", yearData))
	februaryStart, _ := time.Parse(layoutISO, fmt.Sprintf("%s-02-01", yearData))
	marchStart, _ := time.Parse(layoutISO, fmt.Sprintf("%s-03-01", yearData))
	marchEnd, _ := time.Parse(layoutISO, fmt.Sprintf("%s-03-31", yearData))
	aprilStart, _ := time.Parse(layoutISO, fmt.Sprintf("%s-04-01", yearData))
	aprilEnd, _ := time.Parse(layoutISO, fmt.Sprintf("%s-04-30", yearData))
	mayStart, _ := time.Parse(layoutISO, fmt.Sprintf("%s-05-01", yearData))
	mayEnd, _ := time.Parse(layoutISO, fmt.Sprintf("%s-05-31", yearData))
	juneStart, _ := time.Parse(layoutISO, fmt.Sprintf("%s-06-01", yearData))
	juneEnd, _ := time.Parse(layoutISO, fmt.Sprintf("%s-06-30", yearData))
	julyStart, _ := time.Parse(layoutISO, fmt.Sprintf("%s-07-01", yearData))
	julyEnd, _ := time.Parse(layoutISO, fmt.Sprintf("%s-07-31", yearData))
	augustStart, _ := time.Parse(layoutISO, fmt.Sprintf("%s-08-01", yearData))
	augustEnd, _ := time.Parse(layoutISO, fmt.Sprintf("%s-08-31", yearData))
	semStart, _ := time.Parse(layoutISO, fmt.Sprintf("%s-09-01", yearData))
	semEnd, _ := time.Parse(layoutISO, fmt.Sprintf("%s-09-30", yearData))
	okStart, _ := time.Parse(layoutISO, fmt.Sprintf("%s-10-01", yearData))
	okEnd, _ := time.Parse(layoutISO, fmt.Sprintf("%s-10-31", yearData))
	novStart, _ := time.Parse(layoutISO, fmt.Sprintf("%s-11-01", yearData))
	novEnd, _ := time.Parse(layoutISO, fmt.Sprintf("%s-11-30", yearData))
	desStart, _ := time.Parse(layoutISO, fmt.Sprintf("%s-12-01", yearData))
	desEnd, _ := time.Parse(layoutISO, fmt.Sprintf("%s-12-31", yearData))

	month.JanuaryStart = januaryStart
	month.JanuaryEnd = januaryEnd
	month.FebruaryStart = februaryStart
	month.FebruaryEnd = februaryEnd
	month.MarchStart = marchStart
	month.MarchEnd = marchEnd
	month.AprilStart = aprilStart
	month.AprilEnd = aprilEnd
	month.MayStart = mayStart
	month.MayEnd = mayEnd
	month.JuneStart = juneStart
	month.JuneEnd = juneEnd
	month.JulyStart = julyStart
	month.JulyEnd = julyEnd
	month.AugustStart = augustStart
	month.AugustEnd = augustEnd
	month.SeptemberStart = semStart
	month.SeptemberEnd = semEnd
	month.OctoberStart = okStart
	month.OctoberEnd = okEnd
	month.NovemberStart = novStart
	month.NovemberEnd = novEnd
	month.DecemberStart = desStart
	month.DecemberEnd = desEnd

	return month
}

func TimeHelper(date string, dateType string) int {
	arrayDate := strings.Split(date, "-")
	var dateTime int

	if dateType == "year" {
		yearDate := arrayDate[0]
		yearInt, _ := strconv.ParseInt(yearDate, 10, 32)
		dateTime = int(yearInt)
	} else if dateType == "month" {
		monthDate := arrayDate[1]
		monthInt, _ := strconv.ParseInt(monthDate, 10, 32)
		dateTime = int(monthInt)
	}

	return dateTime

}

func DayHelper(date string) Day {
	var (
		day          Day
		arrayDayData int
		dayData      string
		dayEndData   string
	)

	arrayDate := strings.Split(date, "-")

	yearDate := arrayDate[0]
	monthDate := arrayDate[1]

	yearInt, _ := strconv.ParseInt(yearDate, 10, 32)

	switch monthDate {
	case "01":
		arrayDayData = 31
	case "02":
		if isLeap(int(yearInt)) {
			arrayDayData = 29
		} else {
			arrayDayData = 28
		}
	case "03":
		arrayDayData = 31
	case "04":
		arrayDayData = 30
	case "05":
		arrayDayData = 31
	case "06":
		arrayDayData = 30
	case "07":
		arrayDayData = 31
	case "08":
		arrayDayData = 31
	case "09":
		arrayDayData = 30
	case "10":
		arrayDayData = 31
	case "11":
		arrayDayData = 30
	case "12":
		arrayDayData = 31
	}

	for i := 1; i < arrayDayData+1; i++ {
		index := strconv.FormatInt(int64(i), 10)
		if i == 1 || i == 2 || i == 3 || i == 4 || i == 5 || i == 6 || i == 7 || i == 8 || i == 9 {
			dayData = fmt.Sprintf("%s-%s-0%s 00:00:00", yearDate, monthDate, index)
			dayEndData = fmt.Sprintf("%s-%s-0%s 23:59:59", yearDate, monthDate, index)
		} else {
			dayData = fmt.Sprintf("%s-%s-%s 00:00:00", yearDate, monthDate, index)
			dayEndData = fmt.Sprintf("%s-%s-%s 23:59:59", yearDate, monthDate, index)
		}

		switch i {
		case 1:
			day.Day1 = dayData
			day.DayEnd1 = dayData
		case 2:
			day.Day2 = dayData
			day.DayEnd2 = dayEndData
		case 3:
			day.Day3 = dayData
			day.DayEnd3 = dayEndData
		case 4:
			day.Day4 = dayData
			day.DayEnd4 = dayEndData
		case 5:
			day.Day5 = dayData
			day.DayEnd5 = dayEndData
		case 6:
			day.Day6 = dayData
			day.DayEnd6 = dayEndData
		case 7:
			day.Day7 = dayData
			day.DayEnd7 = dayEndData
		case 8:
			day.Day8 = dayData
			day.DayEnd8 = dayEndData
		case 9:
			day.Day9 = dayData
			day.DayEnd9 = dayEndData
		case 10:
			day.Day10 = dayData
			day.DayEnd10 = dayEndData
		case 11:
			day.Day11 = dayData
			day.DayEnd11 = dayEndData
		case 12:
			day.Day12 = dayData
			day.DayEnd12 = dayEndData
		case 13:
			day.Day13 = dayData
			day.DayEnd13 = dayEndData
		case 14:
			day.Day14 = dayData
			day.DayEnd14 = dayEndData
		case 15:
			day.Day15 = dayData
			day.DayEnd15 = dayEndData
		case 16:
			day.Day16 = dayData
			day.DayEnd16 = dayEndData
		case 17:
			day.Day17 = dayData
			day.DayEnd17 = dayEndData
		case 18:
			day.Day18 = dayData
			day.DayEnd18 = dayEndData
		case 19:
			day.Day19 = dayData
			day.DayEnd19 = dayEndData
		case 20:
			day.Day20 = dayData
			day.DayEnd20 = dayEndData
		case 21:
			day.Day21 = dayData
			day.DayEnd21 = dayEndData
		case 22:
			day.Day22 = dayData
			day.DayEnd22 = dayEndData
		case 23:
			day.Day23 = dayData
			day.DayEnd23 = dayEndData
		case 24:
			day.Day24 = dayData
			day.DayEnd24 = dayEndData
		case 25:
			day.Day25 = dayData
			day.DayEnd25 = dayEndData
		case 26:
			day.Day26 = dayData
			day.DayEnd26 = dayEndData
		case 27:
			day.Day27 = dayData
			day.DayEnd27 = dayEndData
		case 28:
			day.Day28 = dayData
			day.DayEnd28 = dayEndData
		case 29:
			day.Day29 = dayData
			day.DayEnd29 = dayEndData
		case 30:
			day.Day30 = dayData
			day.DayEnd30 = dayEndData
		case 31:
			day.Day31 = dayData
			day.DayEnd31 = dayEndData
		}
	}

	return day
}

func DateHelper(startDate string, endDate string, typeFilter string) (string, string) {
	var (
		yearStartData  string
		yearEndData    string
		monthStartData string
		monthEndData   string
	)

	arrayDateStart := strings.Split(startDate, "-")
	arrayDateEnd := strings.Split(endDate, "-")

	yearDateStart := arrayDateStart[0]
	yearDateEnd := arrayDateEnd[0]
	monthDateStart := arrayDateStart[1]
	monthDateEnd := arrayDateEnd[1]
	dayStart := splitDay(arrayDateStart[2])
	dayEnd := splitDay(arrayDateEnd[2])

	if typeFilter == "year" {
		yeStart, _ := strconv.ParseInt(yearDateStart, 10, 32)
		yeEnd, _ := strconv.ParseInt(yearDateEnd, 10, 32)

		yearStart := yeStart - 1
		yearEnd := yeEnd - 1

		yearStartData = strconv.FormatInt(yearStart, 10)
		yearEndData = strconv.FormatInt(yearEnd, 10)
		monthStartData = monthDateStart
		monthEndData = monthDateEnd
	} else if typeFilter == "month" {
		moStart, _ := strconv.ParseInt(monthDateStart, 10, 32)
		moEnd, _ := strconv.ParseInt(monthDateEnd, 10, 32)

		monthStart := moStart - 1
		monthEnd := moEnd - 1

		if monthStart == 11 || monthStart == 12 {
			monthStartData = strconv.FormatInt(monthStart, 10)
		} else {
			monthStartData = "0" + strconv.FormatInt(monthStart, 10)
		}
		if monthEnd == 11 || monthEnd == 12 {
			monthEndData = strconv.FormatInt(monthEnd, 10)
		} else {
			monthEndData = "0" + strconv.FormatInt(monthEnd, 10)
		}

		yearStartData = yearDateStart
		yearEndData = yearDateEnd
	}

	dayData := fmt.Sprintf("%s-%s-%s 00:00:00", yearStartData, monthStartData, dayStart)
	dayEndData := fmt.Sprintf("%s-%s-%s 23:59:59", yearEndData, monthEndData, dayEnd)

	return dayData, dayEndData
}

func splitDay(day string) string {
	var splitData []string
	var mergeData string

	splitted := strings.Split(day, "")

	for _, v := range splitted {
		splitData = append(splitData, v)
	}

	for _, v := range splitData[:2] {
		mergeData += v
	}

	return mergeData
}
