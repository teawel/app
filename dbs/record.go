package dbs

import "time"

type Record struct {
	Time  time.Time
	Key   []byte
	Value map[string]string
}
