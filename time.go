package helper_GO

import (
	"fmt"
	"time"
)

func TimeSleep(d time.Duration) {
	time.Sleep(d)
}

func Time_UnixNow() string {
	return fmt.Sprint(time.Now().Unix())
}
