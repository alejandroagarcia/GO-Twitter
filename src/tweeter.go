package main

import (
	"github.com/abiosoft/ishell"
	"github.com/alejandroagarcia/GO-Twitter/src/service"
	"github.com/alejandroagarcia/GO-Twitter/src/domain"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {
			var err error;
			defer c.ShowPrompt(true)

			c.Print("Write your username: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			newTweet := domain.NewTweet(user, text)

			
			err = service.PublishTweet(newTweet)

			if err != nil{
				c.Print(err.Error())
			} else{
				c.Print("Tweet sent\n")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweet := service.GetTweet()

			if tweet != nil {
				c.Println(tweet.User, tweet.Text, tweet.Date)
			} else{
				c.Println("No hay tweets creados.")
			}

			return
		},
	})

	shell.Run()

}
