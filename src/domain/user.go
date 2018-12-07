package domain

type User struct {
	Name     string
	Email    string
	Nick     string
	Password string
}

// func NewUser(name string, email string, nick string, password string) *User {
// 	time := time.Now()
// 	newTweet := TextTweet{User: user, Text: text, Date: &time}
// 	return &newTweet
// }
