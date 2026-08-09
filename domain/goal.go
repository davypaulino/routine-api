
package domain

import "time"

type Goal struct {
	Title		string
	Description	string
	TargetDate	time.Time
	StartedAt	time.Time
}

func NewGoal(title, description string, targetDate time.Time) Goal {
	return Goal{
		Title: title,
		Description: description,
		TargetDate: targetDate,
		StartedAt: time.Now(),
	}
}
