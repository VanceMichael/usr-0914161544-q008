package domain

import "testing"

func TestJourneyEventHasMonotonicVersion(t *testing.T) {
	event := JourneyEvent{JourneyID: "j-1", Version: 3, Kind: TicketIssued}
	if event.Version < 1 {
		t.Fatal("旅程事件版本必须为正数")
	}
}
