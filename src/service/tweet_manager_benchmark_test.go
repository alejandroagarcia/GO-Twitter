package service_test

import (
	"testing"

	"github.com/alejandroagarcia/GO-Twitter/src/domain"
	"github.com/alejandroagarcia/GO-Twitter/src/service"
)

func BenchmarkPublishTweetWithFileTweetWriter(b *testing.B) {
	// Initialization

	fileTweetWriter := service.NewFileTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(fileTweetWriter)

	var tweet domain.Tweet // Fill the tweet with data

	user := "Alejandro"
	text := "first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation

	for index := 0; index < b.N; index++ {
		tweetManager.PublishTweet(tweet)
	}
}
