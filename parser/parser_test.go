package parser

import (
	"testing"
)

func TestDo(t *testing.T) {

	type testCaseStructure struct {
		sourceString string
		deliveryFrom DeliveryFrom
		deliveryTo   DeliveryTo
		human        string
	}

	tCase := new(testCaseStructure)

	tCase.sourceString = "от трех дней до месяца"
	tCase.deliveryFrom.PeriodVal = 3
	tCase.deliveryFrom.PeriodType = 3
	tCase.deliveryTo.PeriodVal = 1
	tCase.deliveryTo.PeriodType = 5
	tCase.deliveryFrom.Valid=true
	tCase.deliveryTo.Valid=true
	tCase.human = "от 3 дней до 1 месяца"

	delivery := Do(tCase.sourceString)
	if tCase.human != delivery.Human() {
		t.Fatalf("Ожидалось: `%s`, получили: `%s`", tCase.human, delivery.Human())
	}
	if tCase.deliveryTo!=delivery.To {
		t.Fatalf("Ошибка разбора deliveryTo для  %s", tCase.sourceString)
	}
	if tCase.deliveryFrom!=delivery.From {
		t.Fatalf("Ошибка разбора deliveryFrom для  %s", tCase.sourceString)
	}

	tCase = new(testCaseStructure)
	tCase.sourceString = "до пары недель"
	tCase.deliveryFrom.PeriodVal = 0
	tCase.deliveryFrom.PeriodType = 0
	tCase.deliveryFrom.Valid=false
	tCase.deliveryTo.Valid=true
	tCase.deliveryTo.PeriodVal = 2
	tCase.deliveryTo.PeriodType = 4
	tCase.human = "до 2 недель"
	delivery = Do(tCase.sourceString)
	if tCase.human != delivery.Human() {
	t.Fatalf("Ожидалось: `%s`, получили: `%s`", tCase.human, delivery.Human())
	}
	if tCase.deliveryTo!=delivery.To {
		t.Fatalf("Ошибка разбора deliveryTo для  %s", tCase.sourceString)
	}
	if tCase.deliveryFrom!=delivery.From {
		t.Fatalf("Ошибка разбора deliveryFrom для  %s", tCase.sourceString)
	}

	tCase = new(testCaseStructure)
	tCase.sourceString = "1-3 часа"
	tCase.deliveryFrom.PeriodVal = 1
	tCase.deliveryFrom.PeriodType = 1
	tCase.deliveryFrom.Valid=true
	tCase.deliveryTo.Valid=true
	tCase.deliveryTo.PeriodVal = 3
	tCase.deliveryTo.PeriodType = 1
	tCase.human = "от 1 до 3 часов"
	delivery = Do(tCase.sourceString)
	if tCase.human != delivery.Human() {
		t.Fatalf("Ожидалось: `%s`, получили: `%s`", tCase.human, delivery.Human())
	}
	if tCase.deliveryTo!=delivery.To {
		t.Fatalf("Ошибка разбора deliveryTo для  %s", tCase.sourceString)
	}
	if tCase.deliveryFrom!=delivery.From {
		t.Fatalf("Ошибка разбора deliveryFrom для  %s", tCase.sourceString)
	}

	tCase = new(testCaseStructure)
	tCase.sourceString = "от четырнадцати рабочих дней"
	tCase.deliveryFrom.PeriodVal = 14
	tCase.deliveryFrom.PeriodType = 2
	tCase.deliveryFrom.Valid=true
	tCase.deliveryTo.Valid=false
	tCase.deliveryTo.PeriodVal = 0
	tCase.deliveryTo.PeriodType = 0
	tCase.human = "от 14 рабочих дней"
	delivery = Do(tCase.sourceString)
	if tCase.human != delivery.Human() {
		t.Fatalf("Ожидалось: `%s`, получили: `%s`", tCase.human, delivery.Human())
	}
	if tCase.deliveryTo!=delivery.To {
		t.Fatalf("Ошибка разбора deliveryTo для  %s", tCase.sourceString)
	}
	if tCase.deliveryFrom!=delivery.From {
		t.Fatalf("Ошибка разбора deliveryFrom для  %s", tCase.sourceString)
	}

	tCase = new(testCaseStructure)
	tCase.sourceString = "от 1 до семи недель"
	tCase.deliveryFrom.PeriodVal = 1
	tCase.deliveryFrom.PeriodType = 4
	tCase.deliveryFrom.Valid=true
	tCase.deliveryTo.Valid=true
	tCase.deliveryTo.PeriodVal = 7
	tCase.deliveryTo.PeriodType = 4
	tCase.human = "от 1 до 7 недель"
	delivery = Do(tCase.sourceString)
	if tCase.human != delivery.Human() {
		t.Fatalf("Ожидалось: `%s`, получили: `%s`", tCase.human, delivery.Human())
	}
	if tCase.deliveryTo!=delivery.To {
		t.Fatalf("Ошибка разбора deliveryTo для  %s", tCase.sourceString)
	}
	if tCase.deliveryFrom!=delivery.From {
		t.Fatalf("Ошибка разбора deliveryFrom для  %s", tCase.sourceString)
	}

	tCase = new(testCaseStructure)
	tCase.sourceString = "не менее 2-х часов"
	tCase.deliveryFrom.PeriodVal = 2
	tCase.deliveryFrom.PeriodType = 1
	tCase.deliveryFrom.Valid=true
	tCase.deliveryTo.Valid=false
	tCase.deliveryTo.PeriodVal = 0
	tCase.deliveryTo.PeriodType = 0
	tCase.human = "от 2 часов"
	delivery = Do(tCase.sourceString)
	if tCase.human != delivery.Human() {
		t.Fatalf("Ожидалось: `%s`, получили: `%s`", tCase.human, delivery.Human())
	}
	if tCase.deliveryTo!=delivery.To {
		t.Fatalf("Ошибка разбора deliveryTo для  %s", tCase.sourceString)
	}
	if tCase.deliveryFrom!=delivery.From {
		t.Fatalf("Ошибка разбора deliveryFrom для  %s", tCase.sourceString)
	}

	tCase = new(testCaseStructure)
	tCase.sourceString = "22 рабочих дня"
	tCase.deliveryFrom.PeriodVal = 0
	tCase.deliveryFrom.PeriodType = 0
	tCase.deliveryFrom.Valid=false
	tCase.deliveryTo.Valid=true
	tCase.deliveryTo.PeriodVal = 22
	tCase.deliveryTo.PeriodType = 2
	tCase.human = "до 22 рабочих дней"
	delivery = Do(tCase.sourceString)
	if tCase.human != delivery.Human() {
		t.Fatalf("Ожидалось: `%s`, получили: `%s`", tCase.human, delivery.Human())
	}
	if tCase.deliveryTo!=delivery.To {
		t.Fatalf("Ошибка разбора deliveryTo для  %s", tCase.sourceString)
	}
	if tCase.deliveryFrom!=delivery.From {
		t.Fatalf("Ошибка разбора deliveryFrom для  %s", tCase.sourceString)
	}

	tCase = new(testCaseStructure)
	tCase.sourceString = "6(8) месяцев"
	tCase.deliveryFrom.PeriodVal = 6
	tCase.deliveryFrom.PeriodType = 5
	tCase.deliveryFrom.Valid=true
	tCase.deliveryTo.Valid=true
	tCase.deliveryTo.PeriodVal = 8
	tCase.deliveryTo.PeriodType = 5
	tCase.human = "от 6 до 8 месяцев"
	delivery = Do(tCase.sourceString)
	if tCase.human != delivery.Human() {
		t.Fatalf("Ожидалось: `%s`, получили: `%s`", tCase.human, delivery.Human())
	}
	if tCase.deliveryTo!=delivery.To {
		t.Fatalf("Ошибка разбора deliveryTo для  %s", tCase.sourceString)
	}
	if tCase.deliveryFrom!=delivery.From {
		t.Fatalf("Ошибка разбора deliveryFrom для  %s", tCase.sourceString)
	}

	tCase = new(testCaseStructure)
	tCase.sourceString = "около 3 часов"
	tCase.deliveryFrom.PeriodVal = 0
	tCase.deliveryFrom.PeriodType = 0
	tCase.deliveryFrom.Valid=false
	tCase.deliveryTo.Valid=true
	tCase.deliveryTo.PeriodVal = 3
	tCase.deliveryTo.PeriodType = 1
	tCase.human = "до 3 часов"
	delivery = Do(tCase.sourceString)
	if tCase.human != delivery.Human() {
		t.Fatalf("Ожидалось: `%s`, получили: `%s`", tCase.human, delivery.Human())
	}
	if tCase.deliveryTo!=delivery.To {
		t.Fatalf("Ошибка разбора deliveryTo для  %s", tCase.sourceString)
	}
	if tCase.deliveryFrom!=delivery.From {
		t.Fatalf("Ошибка разбора deliveryFrom для  %s", tCase.sourceString)
	}

	tCase = new(testCaseStructure)
	tCase.sourceString = "10"
	tCase.deliveryFrom.PeriodVal = 0
	tCase.deliveryFrom.PeriodType = 0
	tCase.deliveryTo.PeriodVal = 10
	tCase.deliveryTo.PeriodType = 14
	tCase.human = "до 10 дней"
	delivery = Do(tCase.sourceString)
	if tCase.human != delivery.Human() {
		t.Fatalf("Ожидалось: `%s`, получили: `%s`", tCase.human, delivery.Human())
	}

}
