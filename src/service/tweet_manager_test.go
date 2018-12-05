package service_test

import (
	"testing"
	"github.com/alejandroagarcia/GO-Twitter/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T){
	var tweet string = "this is my first tweet"

	service.PublishTweet(tweet)

	if service.Tweet != tweet {
		t.Error("Expected tweet is", tweet)
	}
}