package service_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alejandroagarcia/GO-Twitter/src/domain"
	"github.com/alejandroagarcia/GO-Twitter/src/service"
)

func isValidTweet(tweet1 domain.Tweet, tweet2 domain.Tweet) bool {
	if tweet1.GetUser() != tweet2.GetUser() {
		return false
	}

	if tweet1.GetText() != tweet2.GetText() {
		return false
	}

	return true
}

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	// Operation
	tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweet()

	if publishedTweet.GetUser() != user && publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s", user, text, publishedTweet.GetUser(), publishedTweet.GetText())
	}

	if publishedTweet.GetDate() == nil {
		t.Error("Expected date can't be nil")
	}
}

func TestWithoutUserIsNotPublished(t *testing.T) {
	// Initilization

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {
	// Initilization

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet

	user := "Alejandro"
	var text string

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	// Initilization

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet

	user := "Alejandro"
	text := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "text greater than 140 characters" {
		t.Error("Expected error is text greater than 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	// Initialization

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet domain.Tweet // Fill de tweets with data

	user1 := "Alejandro"
	text1 := "Tweet 1"

	user2 := "Javier"
	text2 := "Tweet 2"

	tweet = domain.NewTextTweet(user1, text1)
	secondTweet = domain.NewTextTweet(user2, text2)

	// Operation
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()

	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(tweet, firstPublishedTweet) {
		return
	}

	if !isValidTweet(secondTweet, secondPublishedTweet) {
		return
	}
}

func TestCanRetrieveTweetById(t *testing.T) {
	// Initialization

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet // Fill de tweets with data
	var id int

	user := "Alejandro"
	text := "Tweet 1"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	if !isValidTweet(tweet, publishedTweet) {
		return
	}
}

func TestCanCountTheTweetSentByAnUser(t *testing.T) {
	// Initialization

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet, thirdTweet domain.Tweet // Fill de tweets with data

	user := "Alejandro"
	anotherUser := "Otro Usuario"

	text := "Tweet"

	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, text)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	count := tweetManager.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {
	// Initialization

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet, secondTweet, thirdTweet domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	// publish the 3 tweets

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 { /* handle error */
	}
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(tweet, firstPublishedTweet) {
		return
	}

	if !isValidTweet(secondTweet, secondPublishedTweet) {
		return
	}

	// check if isValidTweet for firstPublishedTweet and secondPublishedTweet
}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {

	// Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is my image",
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png")
	// Operation
	text := tweet.PrintableTweet()
	// Validation
	expectedText := "@grupoesfera: This is my image, Link: http://www.grupoesfera.com.ar/common/img/grupoesfera.png"

	// return fmt.Sprintf("@%s: %s, Link: %s", t.User, t.Text, t.Url)
	if text != expectedText {
		t.Errorf("Expected text is: %s", expectedText)
	}

}

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
	// Initialization
	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)

	// Operation
	text := tweet.PrintableTweet()

	// Validation
	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	if text != expectedText {
		fmt.Println(text)
		t.Errorf("Expected text is: %s", expectedText)
	}
}

func TestPublishedTweetIsSavedToExternalResource(t *testing.T) {
	// Initialization

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet // Fill the tweet with data

	user := "Alejandro"
	text := "Prueba"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	id, _ := tweetManager.PublishTweet(tweet)

	// Validation
	memoryWritter := (tweetWriter).(*service.MemoryTweetWriter)
	savedTweet := memoryWritter.GetLastSavedTweet()

	if savedTweet == nil {
		t.Errorf("Expected one tweet")
	}

	if savedTweet.GetId() != id {
		t.Errorf("Expected %d but was %d", id, savedTweet.GetId())
	}
}

func TestCanSearchForTweetContainingText(t *testing.T) {
	// Initialization

	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter() // Mock implementation
	tweetManager := service.NewTweetManager(tweetWriter)

	var tweet domain.Tweet // Fill the tweet with data

	user := "Alejandro"
	text := "first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	tweetManager.PublishTweet(tweet)

	searchResult := make(chan domain.Tweet)
	quit := make(chan string)
	query := "first"

	tweetManager.SearchTweetsContaining(query, searchResult, quit)

	// Validation
	foundTweet := <-searchResult

	if foundTweet == nil {
		t.Errorf("Expected one tweet")
	}

	if !strings.Contains(foundTweet.GetText(), query) {
		t.Errorf("Expected %s contain %s", foundTweet.GetText(), query)
	}
}
