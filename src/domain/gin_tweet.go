package domain

type GinTweet struct {
	User string `json "user"`
	Text string `json "text"`
	Id   int    `json "id"`
}
