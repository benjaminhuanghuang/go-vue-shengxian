package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

func Page(Limit, page int64) (limit, offset int64) {
	if Limit > 0 {
		limit = Limit
	} else {
		limit = 0
	}

	if page > 0 {
		offset = (page - 1) * limit
	} else {
		offset = -1
	}

	return limit, offset
}

func Sort(Sort string) (sort string) {
	if Sort != "" {
		sort = Sort
	} else {
		sort = "create_at desc"
	}

	return sort
}

const TimeLayout = "2006-10-01 15:04:04"

var (
	Local = time.FixedZone("CST", 8*3600)
)

func GetNow() string {
	now := time.Now().In(Local).Format(TimeLayout)
	return now
}

func TimeFormat(s string) string {
	result, err := time.ParseInLocation(TimeLayout, s, time.Local)
	if err != nil {
		panic(err)
	}
	return result.In(Local).Format(TimeLayout)
}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5Str := fmt.Sprintf("%x", w.Sum(nil))
	return md5Str
}
