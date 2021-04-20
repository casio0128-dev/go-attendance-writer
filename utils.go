package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func parse(str string) (string, error) {
	if len(str) <= 0 {
		return "", fmt.Errorf("args is none.")
	}

	reg, err := regexp.Compile(`({.{1,4}})`)
	if err != nil {
		return "", err
	}

	for _, match := range reg.FindAllString(str, -1) {
		if reg.MatchString(str) {
			str = parseDate(str, match, time.Now())
		}
	}

	return str, nil
}

func parseDate(src, format string, tim time.Time) (result string) {
	switch format {
	case "{YYYY}":
		result = strings.ReplaceAll(src, "{YYYY}", tim.Format("2006"))
	case "{MM}":
		result = strings.ReplaceAll(src, "{MM}", tim.Format("01"))
	case "{M}":
		result = strings.ReplaceAll(src, "{M}", tim.Format("1"))
	case "{DD}":
		result = strings.ReplaceAll(src, "{DD}", tim.Format("02"))
	case "{D}":
		result = strings.ReplaceAll(src, "{D}", tim.Format("2"))
	case "{hh}":
		result = strings.ReplaceAll(src, "{hh}", tim.Format("15"))
	case "{0h}":
		result = strings.ReplaceAll(src, "{0h}", tim.Format("03"))
	case "{h}":
		result = strings.ReplaceAll(src, "{h}", tim.Format("3"))
	case "{mm}":
		result = strings.ReplaceAll(src, "{mm}", tim.Format("04"))
	case "{m}":
		result = strings.ReplaceAll(src, "{m}", tim.Format("4"))
	case "{dd}":
		result = strings.ReplaceAll(src, "{dd}", tim.Format("05"))
	case "{d}":
		result = strings.ReplaceAll(src, "{d}", tim.Format("5"))
	case "{WW}":
		result = strings.ReplaceAll(src, "{WW}", tim.Format("Monday"))
	case "{W}":
		result = strings.ReplaceAll(src, "{W}", tim.Format("Mon"))
	case "{JW}":
		result = strings.ReplaceAll(src, "{JW}", convertWeek(tim.Format("Mon")))
	}
	return
}

func convertWeek(week string) string {
	switch week {
	case "Mon":
		return "月"
	case "Tue":
		return "火"
	case "Wed":
		return "水"
	case "Thu":
		return "木"
	case "Fri":
		return "金"
	case "Sat":
		return "土"
	case "Sun":
		return "日"
	}
	return ""
}
