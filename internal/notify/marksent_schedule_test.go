package notify

import (
	"testing"
	"time"
)

func TestMarkSentRejectsBeforeScheduleAt(t *testing.T) {
	s := New()
	now := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	future := now.Add(24 * time.Hour)

	s.Create(CreateInput{
		ID: "SCH1", Recipient: "u", Content: "c",
		ScheduleAt: &future,
	}, now)

	sentTime := now.Add(1 * time.Hour)
	_, err := s.MarkSent("SCH1", sentTime)
	if err == nil {
		t.Errorf("MarkSent should reject send before ScheduleAt: sentTime=%v, scheduleAt=%v",
			sentTime, future)
	}
}

func TestMarkSentAllowsAfterScheduleAt(t *testing.T) {
	s := New()
	now := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	future := now.Add(1 * time.Hour)

	s.Create(CreateInput{
		ID: "SCH2", Recipient: "u", Content: "c",
		ScheduleAt: &future,
	}, now)

	sentTime := now.Add(2 * time.Hour)
	_, err := s.MarkSent("SCH2", sentTime)
	if err != nil {
		t.Errorf("MarkSent should allow send after ScheduleAt: %v", err)
	}
}
