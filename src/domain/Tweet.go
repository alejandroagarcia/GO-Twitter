package domain

import (
	"time"
	"fmt"
)

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

func (t *Tweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", t.User, t.Text)
}

func (t *Tweet) String() string {
	return t.PrintableTweet()
}