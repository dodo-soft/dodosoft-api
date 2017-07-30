package events

import (
	"errors"
	"fmt"
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
var eventCache *cache

var events = []eventInfo{
	eventInfo{"89931", "ce7eaf8216f555dec0471990c88750b3"},
}

func init() {
	client = newAtndClient()
	eventCache = newCache()
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

	// get event from cache
	var event *Event
	eventFromCache := eventCache.get(id)
	if eventFromCache != nil {
		e := (*eventFromCache).(Event)
		event = &e
	}
	if event != nil {
		fmt.Println(event)
		return event, nil
	}
	// query to the atnd server
	event, err = client.queryEvent(e.id, e.key)
	if err != nil {
		return nil, err
	}
	// store to the cache
	eventCache.Put(id, *event)

	return event, nil
}
