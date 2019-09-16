package parser

import (
	"fmt"
	"strings"
)

type Delivery struct {
	From DeliveryFrom
	To   DeliveryTo
}

func (d *Delivery) Human() string {
	if d.From.Valid && d.To.Valid && d.From.PeriodType == d.To.PeriodType {
		return fmt.Sprintf("от %d %s", d.From.PeriodVal, d.To.Human())
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", d.From.Human(), d.To.Human()))

}

type DeliveryFrom struct {
	PeriodVal  uint
	PeriodType uint
	Valid      bool
}

func (d *DeliveryFrom) Human() string {
	if d.Valid {
		return fmt.Sprintf("от %d %s", d.PeriodVal, periodTypeToString(d.PeriodVal, d.PeriodType))
	}
	return ""
}

type DeliveryTo struct {
	PeriodVal  uint
	PeriodType uint
	Valid      bool
}

func (d *DeliveryTo) Human() string {
	if d.Valid {
		return fmt.Sprintf("до %d %s", d.PeriodVal, periodTypeToString(d.PeriodVal, d.PeriodType))
	}
	return ""
}

var DeliveryTimeToWords = map[uint]map[int]string{
	HOUR:     map[int]string{1: "часа", 2: "часов"},
	WORK_DAY: map[int]string{1: "рабочего дня", 2: "рабочих дней"},
	DAY:      map[int]string{1: "дня", 2: "дней"},
	WEEK:     map[int]string{1: "недели", 2: "недель"},
	MONTH:    map[int]string{1: "месяца", 2: "месяцев"},
}

func periodTypeToString(periodVal uint, periodType uint) (periodTypeWords string) {

	if periodVal%10 == 1 && periodVal != 11 {
		periodTypeWords = DeliveryTimeToWords[periodType][1]
	} else {
		periodTypeWords = DeliveryTimeToWords[periodType][2]
	}

	return periodTypeWords
}

const HOUR = 1
const WORK_DAY = 2
const DAY = 3
const WEEK = 4
const MONTH = 5
