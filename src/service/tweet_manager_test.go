package service_test

import (
	"testing"
	"github.com/alejandroagarcia/GO-Twitter/src/service"
	"github.com/alejandroagarcia/GO-Twitter/src/domain"
)

func isValidTweet(tweet1 *domain.Tweet, tweet2 *domain.Tweet) bool{
	if (tweet1.User != tweet2.User){
		return false
	}

	if (tweet1.Text != tweet2.Text){
		return false
	}

	return true
}

func TestPublishedTweetIsSaved(t *testing.T){

	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()

	if publishedTweet.User != user && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", user, text, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestWithoutUserIsNotPublished(t *testing.T){
	// Initilization
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T){
	// Initilization
	var tweet *domain.Tweet

	user := "Alejandro"
	var text string

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T){
	// Initilization
	var tweet *domain.Tweet

	user := "Alejandro"
	text := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	_, err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "text greater than 140 characters" {
		t.Error("Expected error is text greater than 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T){
	// Initialization

	service.InitializeService()
	var tweet, secondTweet *domain.Tweet // Fill de tweets with data

	user1 := "Alejandro"
	text1 := "Tweet 1"

	user2 := "Javier"
	text2 := "Tweet 2"

	tweet = domain.NewTweet(user1, text1)
	secondTweet = domain.NewTweet(user2, text2)

	// Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	// Validation
	publishedTweets := service.GetTweets()

	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(tweet, firstPublishedTweet){
		return
	}

	if !isValidTweet(secondTweet, secondPublishedTweet){
		return
	}
}

func TestCanRetrieveTweetById(t *testing.T){
	// Initialization

	service.InitializeService()
	var tweet *domain.Tweet // Fill de tweets with data
	var id int

	user := "Alejandro"
	text := "Tweet 1"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweetById(id)

	if !isValidTweet(tweet, publishedTweet){
		return
	}
}

func TestCanCountTheTweetSentByAnUser(t *testing.T){
	// Initialization

	service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet // Fill de tweets with data

	user := "Alejandro"
	anotherUser := "Otro Usuario"

	text := "Tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, text)
	thirdTweet = domain.NewTweet(anotherUser, text)

	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)
	service.PublishTweet(thirdTweet)

	// Operation
	count := service.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}