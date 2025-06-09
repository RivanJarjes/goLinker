package core

import "time"

type Link struct {
	ID       int64
	Code     string
	Original string
	Created  time.Time
}
