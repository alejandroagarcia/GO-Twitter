package service_test

import (
	"testing"
	"github.com/alejandroagarcia/GO-Twitter/src/service"
	"github.com/alejandroagarcia/GO-Twitter/src/domain"
)

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
	err = service.PublishTweet(tweet)

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
	err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T){
	// Initilization
	var tweet *domain.Tweet

	user := "Alejandro"
	text := "Prueba"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "text greater than 140 characters" {
		t.Error("Expected error is text greater than 140 characters")
	}
}