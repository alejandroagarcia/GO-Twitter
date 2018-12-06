package main

import (
	"github.com/abiosoft/ishell"
	"github.com/alejandroagarcia/GO-Twitter/src/service"
	"github.com/alejandroagarcia/GO-Twitter/src/domain"
	"strconv"
)

func main() {
	tweetManager := service.NewTweetManager()
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
		
			_, err := tweetManager.PublishTweet(newTweet)

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

			tweet := tweetManager.GetTweet()

			if tweet != nil {
				c.Println(tweet)
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

			tweets := tweetManager.GetTweets()

			if tweets != nil {
				for _, valor := range tweets {
					c.Println(valor)
				}
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

			tweet := tweetManager.GetTweetById(id)

			if tweet != nil {
				c.Println(tweet)
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

			count := tweetManager.CountTweetsByUser(user)

			if count > 0 {
				c.Println(count)
			} else{
				c.Println("El usuario no posee tweets.")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetsByUserName",
		Help: "Muestra todos los tweets de un usuario",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Write your username: ")

			user := c.ReadLine()

			tweets := tweetManager.GetTweetsByUser(user)

			if tweets != nil {
				for _, valor := range tweets {
					c.Println(valor)
				}
			} else{
				c.Println("Usuario inexistente o no contiene tweets.")
			}

			return
		},
	})

	shell.Run()

}
