package utils

import (
    "fmt"
    "testing"
    "time"
)

func TestTimestamp(t *testing.T) {
    timestamp := time.Now().Unix()
    str := time.Unix(timestamp, 0).Format("2006-01-02")
    fmt.Println(str)
    fmt.Println(DateFormat(timestamp))
}

func DateFormat(timestamp int64) string {
    // tm := time.Unix(timestamp, 0)
    return time.Unix(timestamp, 0).Format("2006-01-02")
}

func TestDateTimezone(t *testing.T) {
    var timestamp int64 = 1655830471
    datetime := timeZoneTrans(timestamp, 4)
    fmt.Println(datetime)
}

func timeZoneTrans(timestamp int64, timezone float32) string {
    timezoneLocal := time.FixedZone("local", int(timezone*60*60))
    timestampLocal := time.Unix(timestamp, 0).In(timezoneLocal)
    fmt.Println(timestamp, timestampLocal)
    return timestampLocal.Format("2006-01-02 15:04:05")
}
