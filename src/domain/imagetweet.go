package domain

import (
	"fmt"
	"time"
)

type ImageTweet struct {
	TextTweet
	Url string `json "url"`
}

func NewImageTweet(user string, text string, url string) *ImageTweet {
	time := time.Now()
	newTweet := ImageTweet{TextTweet{User: user, Text: text, Date: &time}, url}
	return &newTweet
}

func (t *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s, Link: %s", t.User, t.Text, t.Url)
}

func (t *ImageTweet) String() string {
	return t.PrintableTweet()
}
