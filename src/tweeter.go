package main

import (
	"github.com/abiosoft/ishell"
	"github.com/alejandroagarcia/GO-Twitter/src/service"
	"github.com/alejandroagarcia/GO-Twitter/src/domain"
	"strconv"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)

			c.Print("Write your username: ")

			user := c.ReadLine()

			c.Print("Write your tweet: ")

			text := c.ReadLine()

			newTweet := domain.NewTweet(user, text)
		
			_, err := service.PublishTweet(newTweet)

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
				c.Println(tweet.Id, tweet.User, tweet.Text, tweet.Date)
			} else{
				c.Println("No hay tweets creados.")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweets := service.GetTweets()

			if tweets != nil {
				for _, valor := range tweets {
					c.Println(valor.Id, valor.User, valor.Text, valor.Date)
				}
				// for index := 0; index < len(tweets); index++ {
				// 	c.Println(tweets[index].Id, tweets[index].User, tweets[index].Text, tweets[index].Date)
				// }
			} else{
				c.Println("No hay tweets creados.")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetById",
		Help: "Obtiene el Tweet por el ID",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Ingrese el ID: ")

			id, _ := strconv.Atoi(c.ReadLine())

			tweet := service.GetTweetById(id)

			if tweet != nil {
				c.Println(tweet.Id, tweet.User, tweet.Text, tweet.Date)
			} else{
				c.Println("ID inexistente.")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getCountTweetsByUserName",
		Help: "Obtiene la cantidad de tweets por usuario",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Ingrese el usuario: ")

			user := c.ReadLine()

			count := service.CountTweetsByUser(user)

			if count > 0 {
				c.Println(count)
			} else{
				c.Println("El usuario no posee tweets.")
			}

			return
		},
	})

	shell.Run()

}
