package events

import (
	"errors"
	"strconv"
)

type Event struct {
	Title   string   `json:"title"`
	Members []Member `json:"members"`
	Limit   int      `json:"limit"`
}

type eventInfo struct {
	id  string
	key string
}

var client *AtndClient

var events = []eventInfo{
	eventInfo{"89931", "ce7eaf8216f555dec0471990c88750b3"},
}

func init() {
	client = newAtndClient()
}

type Member struct {
	Name string `json:"name"`
}

func QueryEventIds() []string {
	ids := []string{}
	for i := range events {
		ids = append(ids, strconv.Itoa(i))
	}
	return ids
}

func QueryEvent(id string) (*Event, error) {
	index, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid id:" + id)
	}
	e := events[index]
	return client.queryEvent(e.id, e.key)
}
