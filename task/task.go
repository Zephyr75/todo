package task

import (
	"time"
)


const (
	Chill  int = 0
	Mid        = 1
	Urgent     = 2
)

type Task struct {
	Title string
	Done  bool
	Importance int
	Project string
	Date time.Time
}


