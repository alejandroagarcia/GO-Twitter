package service


import (
	"github.com/alejandroagarcia/GO-Twitter/src/domain"
	"fmt"
)

type TweetManager struct {
	tweet *domain.Tweet
	tweets []*domain.Tweet
	mapTweetsByUser map[string][]*domain.Tweet
}

func NewTweetManager() *TweetManager{
	var tm TweetManager
	tm.tweets = make([]*domain.Tweet, 0)
	tm.mapTweetsByUser = make(map[string][]*domain.Tweet)

	return &tm
}

func (tm *TweetManager) PublishTweet(t *domain.Tweet) (int, error) {
	var err error

	if t.User == "" {
		err = fmt.Errorf("user is required")
		return -1, err
	}

	if t.Text == "" {
		err = fmt.Errorf("text is required")
		return -1, err
	}

	if len(t.Text) > 140 {
		err = fmt.Errorf("text greater than 140 characters")
		return -1, err
	}

	t.Id = len(tm.tweets) + 1

	_, existe := tm.mapTweetsByUser[t.User]

	if existe {
		tm.mapTweetsByUser[t.User] = append(tm.mapTweetsByUser[t.User], t)
	} else {
		tm.mapTweetsByUser[t.User] = make([]*domain.Tweet, 0)
		tm.mapTweetsByUser[t.User] = append(tm.mapTweetsByUser[t.User], t)
	}

	tm.tweets = append(tm.tweets, t)

	return t.Id, nil

}

func (tm *TweetManager) GetTweet() *domain.Tweet {
	return tm.tweets[len(tm.tweets)-1];
}

func (tm *TweetManager) GetTweetById(id int) *domain.Tweet {
	if id > 0 && id <= len(tm.tweets){
		tweet := tm.tweets[id - 1]
		return tweet
	}
	
	return nil
}

func (tm *TweetManager) GetTweets() []*domain.Tweet {
	return tm.tweets;
}

func (tm *TweetManager) CountTweetsByUser(user string) int {
	var count int

	for _, value := range tm.tweets {
		if value.User == user {
			count ++
		}
	}

	return count
}

func (tm *TweetManager) GetTweetsByUser(user string) []*domain.Tweet {
	return tm.mapTweetsByUser[user]
}