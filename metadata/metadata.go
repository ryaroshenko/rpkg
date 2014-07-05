package metadata

import (
	"time"
)

const (
	DateTimeFormat = "02.01.2006 15:04:05.000"
)

// DateTime - тип данных, хранящий дату+время ввиде строки
// формата "02.01.2006 15:04:05.000" ("dd.mm.yyyy hh:mm:ss.zzz")
type DateTime string

func (dt DateTime) String() string {
	return string(dt)
}

// Decode - функция декодирования строки дата+время во внутрений тип time.Time
func (dt DateTime) Decode() (time.Time, error) {
	return time.Parse(DateTimeFormat, string(dt))
}

// EncodeTime - функция кодирования значения time.Time в тип dshop.metadata.DateTime
func EncodeTime(t time.Time) DateTime {
	return DateTime(t.Format(DateTimeFormat))
}

type Tran struct {
	ShopName   string
	TranTime   DateTime
	PackageUID string
}

type Shop struct {
	Name        string
	Description string
}

type Goods struct {
	UID       string
	Article   string
	ShortName string
	Name      string
	TranInfo  Tran
}
