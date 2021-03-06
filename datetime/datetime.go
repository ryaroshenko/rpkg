// datetime - пакет для работы с датой в "паскалевском" формате
package datetime

import (
	"time"
)

const (
	DateTimeFormat = "02.01.2006 15:04:05.000"
)

// DateTime - тип данных, хранящий дату+время ввиде строки
// формата "02.01.2006 15:04:05.000" ("dd.mm.yyyy hh:mm:ss.zzz")
type DateTime string

// String - поддержка для функций форматирования типа fmt.Printf()
func (dt DateTime) String() string {
	return string(dt)
}

// Decode - функция декодирования строки дата+время во внутрений тип time.Time
func (dt DateTime) Decode() (time.Time, error) {
	return time.Parse(DateTimeFormat, string(dt))
}

// EncodeTime - функция кодирования значения time.Time в тип datetime.DateTime
func EncodeTime(t time.Time) DateTime {
	return DateTime(t.Format(DateTimeFormat))
}
