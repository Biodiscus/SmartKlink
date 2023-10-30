package utils

import (
	"fmt"
	"time"
)

func UnixTimestampString() string {
	return fmt.Sprint(time.Now().Unix())
}
