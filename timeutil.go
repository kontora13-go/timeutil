// Copyright 2024 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

package timeutil

import "time"

var (
	defaultLocation *time.Location = time.FixedZone("Europe/Moscow", 3*3600)
)

const (
	RegularAccountDtEnd             = 4102444800
	RequestTimeFormat               = "2006-01-02 15:04:05"
	MsServerQueryTimeFormat         = "2006-01-02 15:04:05.000"
	MsServerFileGroupNameTimeFormat = "2006-01-02T15-04-05.000"
	MsServerFileRangeTimeFormat     = "2006-01-02 15:04:05.000"
	LogsFormat                      = "2006-01-02T15:04:05.000"
)

// GetDefaultLocation - location по-умолчанию
func GetDefaultLocation() *time.Location {
	return defaultLocation
}

// GetNowTime - получение текущего времени в соответствии с location
func GetNowTime() time.Time {
	return time.Now().In(GetDefaultLocation())
}

// GetNowTimeWithTrimmedTZ - получение текущего времени без тайм-зоны
func GetNowTimeWithTrimmedTZ() time.Time {
	return TrimTimeZone(time.Now().In(GetDefaultLocation()))
}

// Get2100Time - получение даты 2100 год как условная конечная дата
func Get2100Time() time.Time {
	return time.Unix(RegularAccountDtEnd, 0)
}

// TrimTimeZone - обрезка тайм-зоны даты
func TrimTimeZone(now time.Time) time.Time {
	return time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Nanosecond(),
		time.UTC,
	)
}

// ChangeTimeZone - изменение тайм-зоны без изменения времени
// - example - шаблонное время с тайм-зоной
func ChangeTimeZone(dt time.Time, example time.Time) time.Time {
	_, offset1 := example.Zone()
	_, offset2 := dt.Zone()
	if offset2 != offset1 {
		offset2 -= offset1
		dt = dt.Add(time.Second * time.Duration(offset2))
		dt = dt.In(example.Location())
	}
	return dt
}
