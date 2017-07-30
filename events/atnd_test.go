package events

import (
	"fmt"
	"testing"
)

func TestQueryEvent(t *testing.T) {
	client := newAtndClient()
	event, err := client.queryEvent("89931", "ce7eaf8216f555dec0471990c88750b3")
	if err != nil {
		t.Fail()
	}
	fmt.Println(event.Limit)
}
