package domain

import (
	"fmt"
	"time"
)

func NewQuoteTweet(user string, text string, quotedTweet Tweet) *QuoteTweet {
	time := time.Now()
	newTweet := QuoteTweet{TextTweet{User: user, Text: text, Date: &time}, quotedTweet}
	return &newTweet
}

type QuoteTweet struct {
	TextTweet
	QuotedText Tweet
}

func (t *QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf("%s \"%s\"", t.TextTweet.String(), t.QuotedText.String())
}

func (t *QuoteTweet) String() string {
	return t.PrintableTweet()
}
