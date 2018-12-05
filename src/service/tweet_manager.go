package service

import (
	"github.com/alejandroagarcia/GO-Twitter/src/domain"
	"fmt"
)

var tweet *domain.Tweet

func PublishTweet(t *domain.Tweet) error {
	var err error

	if t.User == "" {
		err = fmt.Errorf("user is required")
		return err
	}

	if t.Text == "" {
		err = fmt.Errorf("text is required")
		return err
	}

	if len(t.Text) > 140 {
		err = fmt.Errorf("text greater than 140 characters")
		return err
	}

	tweet = t;
	return nil;
}

func GetTweet() *domain.Tweet {
	return tweet;
}