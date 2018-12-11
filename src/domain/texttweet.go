package domain

import (
	"fmt"
	"time"
)

type TextTweet struct {
	Id   int
	User string `json: "user"`
	Text string `json: "text"`
	Date *time.Time
}

func NewTextTweet(user string, text string) *TextTweet {
	time := time.Now()
	newTweet := TextTweet{User: user, Text: text, Date: &time}
	return &newTweet
}

func (t *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", t.User, t.Text)
}

func (t *TextTweet) String() string {
	return t.PrintableTweet()
}

func (t *TextTweet) GetUser() string {
	return t.User
}

func (t *TextTweet) GetText() string {
	return t.Text
}

func (t *TextTweet) SetId(id int) {
	t.Id = id
}

func (t *TextTweet) GetId() int {
	return t.Id
}

func (t *TextTweet) GetDate() *time.Time {
	return t.Date
}
