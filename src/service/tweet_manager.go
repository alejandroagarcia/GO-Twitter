package service

import (
	"fmt"
	"strings"

	"github.com/alejandroagarcia/GO-Twitter/src/domain"
)

type TweetManager struct {
	tweet           domain.Tweet
	tweets          []domain.Tweet
	mapTweetsByUser map[string][]domain.Tweet
	tweetWriter     TweetWriter
}

func NewTweetManager(tw TweetWriter) *TweetManager {
	var tm TweetManager
	tm.tweets = make([]domain.Tweet, 0)
	tm.mapTweetsByUser = make(map[string][]domain.Tweet)
	tm.tweetWriter = tw

	return &tm
}

func (tm *TweetManager) PublishTweet(t domain.Tweet) (int, error) {
	var err error

	if t.GetUser() == "" {
		err = fmt.Errorf("user is required")
		return -1, err
	}

	if t.GetText() == "" {
		err = fmt.Errorf("text is required")
		return -1, err
	}

	if len(t.GetText()) > 140 {
		err = fmt.Errorf("text greater than 140 characters")
		return -1, err
	}

	t.SetId(len(tm.tweets) + 1)

	_, existe := tm.mapTweetsByUser[t.GetUser()]

	if existe {
		tm.mapTweetsByUser[t.GetUser()] = append(tm.mapTweetsByUser[t.GetUser()], t)
	} else {
		tm.mapTweetsByUser[t.GetUser()] = make([]domain.Tweet, 0)
		tm.mapTweetsByUser[t.GetUser()] = append(tm.mapTweetsByUser[t.GetUser()], t)
	}

	tm.tweets = append(tm.tweets, t)
	tm.tweetWriter.Write(t)

	return t.GetId(), nil

}

func (tm *TweetManager) GetTweet() domain.Tweet {
	if len(tm.tweets) > 0 {
		return tm.tweets[len(tm.tweets)-1]
	}

	return nil
}

func (tm *TweetManager) GetTweetById(id int) domain.Tweet {
	if id > 0 && id <= len(tm.tweets) {
		tweet := tm.tweets[id-1]
		return tweet
	}

	return nil
}

func (tm *TweetManager) GetTweets() []domain.Tweet {
	if len(tm.tweets) > 0 {
		return tm.tweets
	}

	return nil
}

func (tm *TweetManager) CountTweetsByUser(user string) int {
	var count int

	for _, value := range tm.tweets {
		if value.GetUser() == user {
			count++
		}
	}

	return count
}

func (tm *TweetManager) GetTweetsByUser(user string) []domain.Tweet {
	return tm.mapTweetsByUser[user]
}

func (tm *TweetManager) SearchTweetsContaining(query string, c chan domain.Tweet, quit chan string) {
	var count int
	go func() {
		for _, v := range tm.tweets {
			if strings.Contains(v.GetText(), query) {
				c <- v
				count++
			}
		}

		if count > 0 {
			quit <- "finishedWithResults"
		} else {
			quit <- "finishedEmpty"
		}

	}()
}
