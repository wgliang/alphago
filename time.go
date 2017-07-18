package alphago

import (
	"fmt"
	"strings"
	"time"
)

var (
	patterns = []string{
		"Y", "2006",
		"y", "06",

		"m", "01",
		"n", "1",
		"M", "Jan",
		"F", "January",

		"d", "02",
		"j", "2",

		"D", "Mon",
		"l", "Monday",

		"g", "3",
		"G", "15",
		"h", "03",
		"H", "15",

		"a", "pm",
		"A", "PM",

		"i", "04",
		"s", "05",
	}
	layouts = []string{
		"2006-01-02 15:04:05 -0700 MST",
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05 -0700 MST",
		"2006/01/02 15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02 -0700 MST",
		"2006-01-02 -0700",
		"2006-01-02",
		"2006/01/02 -0700 MST",
		"2006/01/02 -0700",
		"2006/01/02",
		"2006-01-02 15:04:05 -0700 -0700",
		"2006/01/02 15:04:05 -0700 -0700",
		"2006-01-02 -0700 -0700",
		"2006/01/02 -0700 -0700",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}
)

// TimeNow is a function that return string of Now.
func TimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// TimeFormat is a function just as time.Format but doesn't need time.Time if
// you want the current time string.
func TimeFormat(format string, ts ...time.Time) string {
	replacer := strings.NewReplacer(patterns...)
	format = replacer.Replace(format)

	t := time.Now()
	if len(ts) > 0 {
		t = ts[0]
	}
	return t.Format(format)
}

// Str2LocalTime is a function that will transfer a string into the Time via
// local zone based on Str2Time().
func Str2LocalTime(value string) time.Time {
	if value == "" {
		return time.Time{}
	}
	zoneName, offset := time.Now().Zone()

	zoneValue := offset / 3600 * 100
	if zoneValue > 0 {
		value += fmt.Sprintf(" +%04d", zoneValue)
	} else {
		value += fmt.Sprintf(" -%04d", zoneValue)
	}

	if zoneName != "" {
		value = fmt.Sprintf("%s %s", value, zoneName)
	}
	return Str2Time(value)
}

// Str2Time is a function that will transfer a string into the Time,
// and support more formate than offical function.
func Str2Time(value string) time.Time {
	if value == "" {
		return time.Time{}
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, value)
		if err == nil {
			return t
		}
	}

	return time.Time{}
}

// IsWeekend is a function that will return if the time is Sunday or Saturday.
func IsWeekend(t time.Time) bool {
	wd := t.Weekday()
	if wd == time.Sunday || wd == time.Saturday {
		return true
	}
	return false
}
