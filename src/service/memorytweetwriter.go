package service

import "github.com/alejandroagarcia/GO-Twitter/src/domain"

type MemoryTweetWriter struct {
	Tweets []domain.Tweet
}

func NewMemoryTweetWriter() *MemoryTweetWriter {
	var mtw MemoryTweetWriter
	mtw.Tweets = make([]domain.Tweet, 0)
	return &mtw
}

func (mtw *MemoryTweetWriter) Write(tweet domain.Tweet) {
	mtw.Tweets = append(mtw.Tweets, tweet)
}

func (mtw *MemoryTweetWriter) GetLastSavedTweet() domain.Tweet {
	return mtw.Tweets[len(mtw.Tweets)-1]
}
