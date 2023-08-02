package task

import (
	"time"
)

type Task struct {
	Title string
	Done  bool
	Importance int
	Project string
	Date time.Time
}


