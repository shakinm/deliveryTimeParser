package parser

import (
	"regexp"
	"strconv"
	"strings"
)

var PeriodLitOrdering = []uint{HOUR, WORK_DAY, DAY, WEEK, MONTH}
var PeriodLit = map[uint][]string{
	HOUR:     []string{`час`, `[\\b|\d|\s](ч)[\.|\s]`},
	WORK_DAY: []string{`[\\b|\d|\s](р.*д)[\S+|\s]`},
	DAY:      []string{`ден`, `[\\b|\d|\s](дн)`, `[\\b|\d|\s](д)[\.|\s]`},
	WEEK:     []string{`нед`, `нд`, `[\\b|\d|\s](н)[\.|\s]`},
	MONTH:    []string{`мес`, `[\\b|\d|\s](м)[\.|\s]`},
}

var LitNumMatchOrdering = []uint{20, 30, 40, 50, 60, 70, 80, 90, 11, 12, 13, 14, 15, 16, 17, 18, 19, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var LitNumMatch = map[uint][]string{
	20: []string{`[\\b|\d|\s](двад)[\S+|\s]`},
	30: []string{`[\\b|\d|\s](трид)[\S+|\s]`},
	40: []string{`[\\b|\d|\s](сор)[\S+|\s]`},
	50: []string{`[\\b|\d|\s](пят(ь|и)д)[\S+|\s]`},
	60: []string{`[\\b|\d|\s](шест(ь|и)д)[\S+|\s]`},
	70: []string{`[\\b|\d|\s](сем(ь|и)д)[\S+|\s]`},
	80: []string{`[\\b|\d|\s](вос(емь|ми)д)[\S+|\s]`},
	90: []string{`[\\b|\d|\s](девян)[\S+|\s]`},

	11: []string{`[\\b|\d|\s](один)[\S+|\s]`},
	12: []string{`[\\b|\d|\s](две)[\S+|\s]`},
	13: []string{`[\\b|\d|\s](трин)[\S+|\s]`},
	14: []string{`[\\b|\d|\s](четырн)[\S+|\s]`},
	15: []string{`[\\b|\d|\s](пятна)[\S+|\s]`},
	16: []string{`[\\b|\d|\s](шестна)[\S+|\s]`},
	17: []string{`[\\b|\d|\s](семн)[\S+|\s]`},
	18: []string{`[\\b|\d|\s](восемн)[\S+|\s]`},
	19: []string{`[\\b|\d|\s](девятн)[\S+|\s]`},

	1:  []string{`[\\b|\\d|\s](од)[\S+|\s]`},
	2:  []string{`[\\b|\\d|\s](дв)[\S+|\s]`, `[\\b|\\d|\s](пар(а|ы))[\S+|\s]`},
	3:  []string{`[\\b|\\d|\s](тр)[\S+|\s]`},
	4:  []string{`[\\b|\\d|\s](чет)[\S+|\s]`},
	5:  []string{`[\\b|\\d|\s](пят)[\S+|\s]`},
	6:  []string{`[\\b|\\d|\s](шест)[\S+|\s]`},
	7:  []string{`[\\b|\\d|\s](сем)[\S+|\s]`},
	8:  []string{`[\\b|\\d|\s](вос(е|ь))[\S+|\s]`},
	9:  []string{`[\\b|\\d|\s](дев)[\S+|\s]`},
	10: []string{`[\\b|\\d|\s](дес)[\S+|\s]`},
}

func Do(deliveryString string) (delivery Delivery) {

	var deliveryParts []string

	deliveryString = strings.TrimSpace(deliveryString)

	rc := regexp.MustCompile(`\s*(х|ух|ех|ёх)`)

	deliveryString = rc.ReplaceAllLiteralString(deliveryString, "")

	rc = regexp.MustCompile(`(до)|(\-)|((^не)мен)|(не.*бол)|(\()`)

	if rc.MatchString(deliveryString) {
		deliveryParts = rc.Split(deliveryString, 2)

		if len(deliveryParts) > 1 {
			delivery.From.PeriodVal, delivery.From.PeriodType = getPart(deliveryParts[0])

			if delivery.From.PeriodVal != 0 {
				delivery.From.Valid = true
			}

			delivery.To.PeriodVal, delivery.To.PeriodType = getPart(deliveryParts[1])

			if delivery.From.PeriodType == 0 && delivery.From.PeriodVal != 0 {
				delivery.From.PeriodType = delivery.To.PeriodType
			}

			if delivery.From.PeriodType == delivery.To.PeriodType {
				if delivery.To.PeriodVal == 0 && delivery.From.PeriodVal > delivery.To.PeriodVal {
					delivery.To.PeriodType = 0
					delivery.To.Valid = false
				}
			} else if delivery.To.PeriodType != 0 && delivery.To.PeriodVal == 0 {
				delivery.To.PeriodVal = 1
				delivery.To.Valid = true
			}

			if delivery.To.PeriodVal != 0 {
				delivery.To.Valid = true
			}

		} else if len(deliveryParts) == 1 {
			delivery.To.PeriodVal, delivery.To.PeriodType = getPart(deliveryParts[0])
		}
		return
	}

	rc = regexp.MustCompile(`(от)|(бол)|(не.*мен)`)
	if rc.MatchString(deliveryString) {
		delivery.From.PeriodVal, delivery.From.PeriodType = getPart(deliveryString)
		delivery.From.Valid = true
		return
	}
	delivery.To.PeriodVal, delivery.To.PeriodType = getPart(deliveryString)
	delivery.To.Valid = true
	if delivery.To.PeriodVal!=0 && delivery.To.PeriodType==0{
		delivery.To.PeriodType=DAY
	}
	return
}

func getPart(deliveryString string) (periodVal uint, periodType uint) {

	for _, num := range LitNumMatchOrdering {
		patterns := LitNumMatch[num]
		for _, pattern := range patterns {
			rc := regexp.MustCompile(pattern)
			if rc.MatchString(deliveryString) {
				periodVal = num
				goto FindType
			}
		}
	}
	if periodVal == 0 {
		rc := regexp.MustCompile(`\d+`)
		matches := rc.FindStringSubmatch(deliveryString)
		if len(matches) > 0 {
			val, _ := strconv.ParseInt(matches[0], 10, 64)
			periodVal = uint(val)
		}
	}

FindType:

	for _, num := range PeriodLitOrdering {
		patterns := PeriodLit[num]
		for _, pattern := range patterns {
			rc := regexp.MustCompile(pattern)
			if rc.MatchString(deliveryString) {
				periodType = num
				goto Next
			}
		}
	}

Next:

	return
}
