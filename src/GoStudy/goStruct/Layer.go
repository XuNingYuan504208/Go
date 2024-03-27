package goStruct

import (
	"time"
)

type Field struct {
	settingsBId string
}

type List[Field any] []Field

type Layer struct {
	id         int
	businessId string
	title      string

	createdAt time.Time

	publicStatus bool
	fields       []Field
}
