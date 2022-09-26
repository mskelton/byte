package util

import "time"

func Id() string {
	return time.Now().UTC().Format("20060102150405")
}
