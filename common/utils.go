package common

import (
	"time"

	"bytes"
	"strconv"

	"github.com/lib/pq"
)

const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	Stamp       = "Jan _2 15:04:05"
	StampMilli  = "Jan _2 15:04:05.000"
	StampMicro  = "Jan _2 15:04:05.000000"
	StampNano   = "Jan _2 15:04:05.000000000"
)

func StringfyNullTimeToRFC3339(OldDate pq.NullTime) string {
	if OldDate.Valid {
		return OldDate.Time.Format(RFC3339)
	}
	return ""
}

func StringfyDateToRFC3339(OldDate time.Time) string {
	if OldDate.IsZero() {
		return ""
	}
	return OldDate.Format(RFC3339)

}

func IsValidString(s string) bool {
	if s != "" {
		return true
	}
	return false
}

func IsValidInt64(i int64) bool {
	if i != 0 {
		return true
	}
	return false
}

func ParseStringToDateRFC3339(stringfyDate string) pq.NullTime {
	t, err := time.Parse(time.RFC3339, stringfyDate)
	if err != nil {
		return pq.NullTime{}
	}
	return pq.NullTime{Time: t, Valid: true}
}

func IdsToString(ids []int64, delim string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(ids); i++ {
		buffer.WriteString(strconv.FormatInt(ids[i], 10))
		if i != len(ids)-1 {
			buffer.WriteString(delim)
		}
	}

	return buffer.String()
}

func GetDiffs(org []int64, curr []int64) (Deleted []int64, Added []int64) {
	orgMap := make(map[int64]bool)
	currMap := make(map[int64]bool)
	for _, v := range curr {
		currMap[v] = true
	}
	for _, v := range org {
		orgMap[v] = true
	}
	for _, v := range curr {
		if !orgMap[v] {
			Added = append(Added, v)
		}
	}
	for _, v := range org {
		if !currMap[v] {
			Deleted = append(Deleted, v)
		}
	}
	return
}

// n is the lenght of the slice, f(int) bool return true if the item with index i is the wanted one
// for example
// a := []string{`sdf`, `adf`, `fff`, `a`}
// idx := SearchInSlice(len(a), func(i int) bool {return a[i] == `a`})
// if idx == -1 {// not found} else {fmt.Println(a[idx] == `a`)}
func SearchInSlice(n int, f func(int) bool) int {
	for i := 0; i < n; i++ {
		if f(i) {
			return i
		}
	}
	return -1
}
