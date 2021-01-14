package simple

import (
	"fmt"
	"strconv"
	"time"

	"github.com/etf1/kafka-message-scheduler/schedule"
)

// Schedule is a simple implementation of the schedule.Schedule interface, used for tests
type Schedule struct {
	id        string
	epoch     int64
	timestamp int64
}

func NewSchedule(id, epoch interface{}, timestamp ...time.Time) schedule.Schedule {
	var sid string

	switch v := id.(type) {
	case int:
		sid = strconv.Itoa(v)
	case int64:
		sid = strconv.FormatInt(v, 10)
	case string:
		sid = v
	default:
		sid = ""
	}

	var iepoch int64
	switch v := epoch.(type) {
	case int:
		iepoch = int64(v)
	case int64:
		iepoch = v
	case time.Time:
		iepoch = v.Unix()
	default:
		iepoch = time.Now().Unix()
	}

	ts := time.Now().Unix()
	if len(timestamp) == 1 {
		ts = timestamp[0].Unix()
	}

	return Schedule{
		id:        sid,
		epoch:     iepoch,
		timestamp: ts,
	}
}

func (s Schedule) ID() string {
	return s.id
}

func (s Schedule) Epoch() int64 {
	return s.epoch
}

func (s Schedule) Timestamp() int64 {
	return s.timestamp
}

func (s Schedule) String() string {
	return fmt.Sprintf("{id:%s epoch:%v date:%v timestamp:%v}", s.ID(), s.Epoch(), time.Unix(s.Epoch(), 0), s.Timestamp())
}
