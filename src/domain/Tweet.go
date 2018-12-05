package domain

import "time"

type Tweet struct {
	User string;
	Text string;
	Date *time.Time;
}

func NewTweet (user string, text string) *Tweet{
	time := time.Now()
	newTweet := Tweet{user, text, &time}
	return &newTweet
}