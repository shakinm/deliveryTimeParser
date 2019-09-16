## Библиотека для разбора строк с сроками доставки/поставки

#### Описание 
На вход принимаются строки типа:

- "от трех дней до месяца"
- "до 2 недель"

Если в разбираемом тесте не указана еденица имзерения времени (час, день и т.п.) подставляется **день**

Результат структура:
```go

const HOUR = 1
const WORK_DAY = 2
const DAY = 3
const WEEK = 4
const MONTH = 5

type Delivery struct {
	From DeliveryFrom
	To   DeliveryTo
}

type DeliveryFrom struct {
	PeriodVal  uint
	PeriodType uint
	Valid      bool
}

type DeliveryTo struct {
	PeriodVal  uint
	PeriodType uint
	Valid      bool
}

```
#### Пример
```.go
delivery := Do("от трех дней до месяца")

fmt.Println(delivery.Human)
// от 3 дней до 1 месяца

```
 