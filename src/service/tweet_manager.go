package service


import (
	"github.com/alejandroagarcia/GO-Twitter/src/domain"
	"fmt"
)

var tweet *domain.Tweet
var tweets []*domain.Tweet 


func InitializeService(){
	tweets = make([]*domain.Tweet, 0)
}

func PublishTweet(t *domain.Tweet) (int, error) {
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

	t.Id = len(tweets) + 1

	tweets = append(tweets, t)

	return t.Id, nil

}

func GetTweet() *domain.Tweet {
	return tweets[len(tweets)-1];
}

func GetTweetById(id int) *domain.Tweet {
	if id > 0 && id <= len(tweets){
		tweet := tweets[id - 1]
		return tweet
	}
	
	return nil
}

func GetTweets() []*domain.Tweet {
	return tweets;
}

func CountTweetsByUser(user string) int {
	var count int

	for _, value := range tweets {
		if value.User == user {
			count ++
		}
	}

	return count
}