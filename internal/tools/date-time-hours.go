package tools

import (
	"fmt"
	"strings"
	"time"
)

func DateTimeHoursFormatter(date string, start string, end string) (time.Time, time.Time, time.Time, float64, error)  {

	fmt.Printf("\ndate: %v | start: %v | end: %v\n",date,start,end)

	var total_hours float64 = 0
	newDate, dateParseErr := time.Parse("2006-01-02",date)
	startTime, timeParseErr1 := time.Parse("15:04:05",start)
	endTime, timeParseErr2 := time.Parse("15:04:05",end)

	if dateParseErr != nil {
		return newDate, startTime, endTime, total_hours, dateParseErr
	}
	if timeParseErr1 != nil {
		return newDate, startTime, endTime, total_hours, timeParseErr1
	}
	if timeParseErr2 != nil {
		return newDate, startTime, endTime, total_hours, timeParseErr2
	}

	p_total_hours := &total_hours
	*p_total_hours = endTime.Sub(startTime).Hours()-0.5

	return newDate, startTime, endTime, *p_total_hours, nil
}

func DateFormatter(date *string) ()  {

	before, _, _ := strings.Cut(*date,"T")
	*date = before
}
