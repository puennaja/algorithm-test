package timeconversion

import (
	"fmt"
	"strconv"
	"strings"
)

// time format 12:00:00AM
// change to 24 format

const (
	am string = "AM"
	pm string = "PM"
)

type Time12Form struct {
	hour   string
	minute string
	sec    string
	half   string
}

func NewTime(s string) *Time12Form {
	upper := strings.ToUpper(s)
	upper = strings.ReplaceAll(upper, am, ":"+am)
	upper = strings.ReplaceAll(upper, pm, ":"+pm)
	timeStr := strings.Split(upper, ":")
	return &Time12Form{
		hour:   timeStr[0],
		minute: timeStr[1],
		sec:    timeStr[2],
		half:   timeStr[3],
	}
}

func (t *Time12Form) ToString24Form() string {
	hourInt, _ := strconv.Atoi(t.hour)
	if t.half == am {
		hourInt = hourInt % 12
	} else {
		hourInt = (hourInt % 12) + 12
	}
	return fmt.Sprintf("%02d:%s:%s", hourInt, t.minute, t.sec)
}

func TimeConversion(s string) string {
	t := NewTime(s)
	return t.ToString24Form()
}
