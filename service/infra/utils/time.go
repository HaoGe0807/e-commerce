package utils

import (
	"e-commerce/service/infra/errors"
	"time"
)

var loc *time.Location

const (
	DEFAULT_FORMAT_LAYOUT_DATE      = "2006-01-02"
	DEFAULT_FORMAT_LAYOUT_DATETIME  = "2006-01-02 15:04:05"
	DEFAULT_FORMAT_LAYOUT_SHORTDATE = "20060102"
	DEFAULT_FORMAT_LAYOUT_CST       = "2006-01-02 15:04:05 +0800 CST"
)

func GetTimestamp() int64 {
	return time.Now().Unix()
}

func GetOffSetTimestamp(years, months, days int) int64 {
	return time.Now().AddDate(years, months, days).Unix()
}

func GetDate() string {

	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05")
}

func GetDayTime() int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")

	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, loc)
	return t.Unix()
}

func ParseInLocation(value string, loc *time.Location, layouts ...string) (time.Time, error) {
	layout := DEFAULT_FORMAT_LAYOUT_DATE
	if len(layouts) > 0 {
		layout = layouts[0]
	}
	return time.ParseInLocation(layout, value, loc)
}

func GetNowTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	timeStr := time.Now().Format(DEFAULT_FORMAT_LAYOUT_DATETIME)
	t, _ := time.ParseInLocation(DEFAULT_FORMAT_LAYOUT_DATETIME, timeStr, loc)
	return t
}

func GetStartAndEndTimeOfDate(t time.Time) (startTime, endTime time.Time) {
	startTime = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	endTime = startTime.Add(time.Hour * 24)
	return startTime, endTime
}

func GetStartTimestampOfDay() int64 {
	currentTime := time.Now()
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
}

func GetEndTimestampOfDay() int64 {
	currentTime := time.Now()
	return time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location()).Unix()
}

func GetBoundaryTimestampOfMonth(baseTime int64) int64 {
	currentTime := time.Unix(baseTime, 0)
	return time.Date(currentTime.Year(), currentTime.Month(), 26, 23, 59, 59, 0, currentTime.Location()).Unix()
}

func GetTimestampOfYear(baseTime int64) int32 {
	DateTime := time.Unix(baseTime, 0)
	return int32(DateTime.Year())
}

func GetTimestampOfMonth(baseTime int64) int32 {
	DateTime := time.Unix(baseTime, 0)
	return int32(DateTime.Month())
}

func GetTimeZone(timeZone int64) (*time.Location, int64, error) {

	switch timeZone {
	case 0:
		secondsEastOfUTC := int((0 * time.Hour).Seconds())
		return time.FixedZone("UTC", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 1:
		secondsEastOfUTC := int((1 * time.Hour).Seconds())
		return time.FixedZone("UTC+1", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 2:
		secondsEastOfUTC := int((2 * time.Hour).Seconds())
		return time.FixedZone("UTC+2", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 3:
		secondsEastOfUTC := int((3 * time.Hour).Seconds())
		return time.FixedZone("UTC+3", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 4:
		secondsEastOfUTC := int((4 * time.Hour).Seconds())
		return time.FixedZone("UTC+4", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 5:
		secondsEastOfUTC := int((5 * time.Hour).Seconds())
		return time.FixedZone("UTC+5", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 6:
		secondsEastOfUTC := int((6 * time.Hour).Seconds())
		return time.FixedZone("UTC+6", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 7:
		secondsEastOfUTC := int((7 * time.Hour).Seconds())
		return time.FixedZone("UTC+7", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 8:
		secondsEastOfUTC := int((8 * time.Hour).Seconds())
		return time.FixedZone("UTC+8", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 9:
		secondsEastOfUTC := int((9 * time.Hour).Seconds())
		return time.FixedZone("UTC+9", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 10:
		secondsEastOfUTC := int((10 * time.Hour).Seconds())
		return time.FixedZone("UTC+10", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 11:
		secondsEastOfUTC := int((11 * time.Hour).Seconds())
		return time.FixedZone("UTC+11", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 12:
		secondsEastOfUTC := int((12 * time.Hour).Seconds())
		return time.FixedZone("UTC+12", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 13:
		secondsEastOfUTC := int((-11 * time.Hour).Seconds())
		return time.FixedZone("UTC-11", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 14:
		secondsEastOfUTC := int((-10 * time.Hour).Seconds())
		return time.FixedZone("UTC-10", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 15:
		secondsEastOfUTC := int((-9 * time.Hour).Seconds())
		return time.FixedZone("UTC-9", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 16:
		secondsEastOfUTC := int((-8 * time.Hour).Seconds())
		return time.FixedZone("UTC-8", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 17:
		secondsEastOfUTC := int((-7 * time.Hour).Seconds())
		return time.FixedZone("UTC-7", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 18:
		secondsEastOfUTC := int((-6 * time.Hour).Seconds())
		return time.FixedZone("UTC-6", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 19:
		secondsEastOfUTC := int((-5 * time.Hour).Seconds())
		return time.FixedZone("UTC-5", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 20:
		secondsEastOfUTC := int((-4 * time.Hour).Seconds())
		return time.FixedZone("UTC-4", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 21:
		secondsEastOfUTC := int((-3 * time.Hour).Seconds())
		return time.FixedZone("UTC-3", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 22:
		secondsEastOfUTC := int((-2 * time.Hour).Seconds())
		return time.FixedZone("UTC-2", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	case 23:
		secondsEastOfUTC := int((-1 * time.Hour).Seconds())
		return time.FixedZone("UTC-1", secondsEastOfUTC), int64(secondsEastOfUTC), nil
	default:
		return nil, 0, errors.ErrorEnum(errors.ERR_INVALID_TIMEZONE, "invalid time Zone")
	}
}

func GetStartAndEndTimeOfYesterdayDate(t time.Time) (startTime, endTime time.Time) {
	startTime = time.Date(t.Year(), t.Month(), t.Day()-1, 0, 0, 0, 0, t.Location())
	endTime = startTime.Add(time.Hour * 24)
	return startTime, endTime
}

func GetScreenStatisticStartAndEndTimeOfDate(t time.Time) (startTime, endTime time.Time) {
	startTime = time.Date(t.Year(), t.Month(), t.Day(), 6, 0, 0, 0, t.Location())
	endTime = startTime.Add(time.Hour * (24 - 6))
	now := time.Now()
	if now.Year() == startTime.Year() && now.Month() == startTime.Month() && now.Day() == startTime.Day() {
		endTime = now
		if endTime.Before(startTime) {
			endTime = startTime
		}
	}
	return startTime, endTime
}

func GetTimeLocation() *time.Location {
	if loc == nil {
		loc, _ = time.LoadLocation("Asia/Shanghai")
	}
	return loc
}

func GetTimeByDate(date string) (time.Time, error) {
	dateTime, err := time.Parse(DEFAULT_FORMAT_LAYOUT_DATE, date)
	if err != nil {
		return dateTime, err
	}
	return time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, GetTimeLocation()), nil
}

func GetLast7DayTime() time.Time {
	now := time.Now()
	last7Day := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, GetTimeLocation()).
		AddDate(0, 0, -7)
	return last7Day
}

func GetQueryDateRangeForStatistic(startDateTime, endDateTime time.Time) ([]time.Time, []time.Time) {
	last7Daytime := GetLast7DayTime()
	t1Range := make([]time.Time, 0)
	t7Range := make([]time.Time, 0)

	if startDateTime.After(endDateTime) {
		return nil, nil
	}

	//note: startDateTime, endDateTime是自然日0点的时间, 返回的时候需要对结束时间加上24小时
	if endDateTime.Before(last7Daytime) || endDateTime.Equal(last7Daytime) {
		t7Range = append(t7Range, startDateTime)
		t7Range = append(t7Range, endDateTime.Add(24*time.Hour))
	} else if startDateTime.After(last7Daytime) {
		t1Range = append(t1Range, startDateTime)
		t1Range = append(t1Range, endDateTime.Add(24*time.Hour))
	} else {
		t1Range = append(t1Range, last7Daytime.Add(24*time.Hour)) //这里的开始时间是6天前的0点(7天前的24点)
		t1Range = append(t1Range, endDateTime.Add(24*time.Hour))
		t7Range = append(t7Range, startDateTime)
		t7Range = append(t7Range, last7Daytime.Add(24*time.Hour))
	}

	return t1Range, t7Range
}

func GetTodayTime() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, GetTimeLocation())
}

func GetBizDateTime(ts int64) time.Time {
	t := time.Unix(ts, 0)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, GetTimeLocation())
}
