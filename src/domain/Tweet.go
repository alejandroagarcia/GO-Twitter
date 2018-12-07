package domain

import "time"

type Tweet interface {
	String() string
	PrintableTweet() string
	GetUser() string
	GetText() string
	SetId(int)
	GetId() int
	GetDate() *time.Time
}
