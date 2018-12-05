package service

var Tweet string

func PublishTweet(tweet string){
	Tweet = tweet;
}

func GetTweet() string {
	return Tweet;
}