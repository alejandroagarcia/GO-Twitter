package domain

import "time"

type Tweet struct {
	Id int;
	User string;
	Text string;
	Date *time.Time;
}

func NewTweet (user string, text string) *Tweet{
	time := time.Now()
	newTweet := Tweet{User: user, Text: text, Date: &time}
	return &newTweet
}