package domain

import "time"

type EventKind string

const (
	TicketIssued   EventKind = "ticket_issued"
	BenefitClaimed EventKind = "benefit_claimed"
	BenefitUsed    EventKind = "benefit_used"
	TicketRefunded EventKind = "ticket_refunded"
	TripChanged    EventKind = "trip_changed"
)

type JourneyEvent struct {
	EventID    string
	JourneyID  string
	OperatorID string
	Version    int64
	Kind       EventKind
	OccurredAt time.Time
}
