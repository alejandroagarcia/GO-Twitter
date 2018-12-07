package main

import (
	"strconv"
	"time"

	"github.com/abiosoft/ishell"
	"github.com/alejandroagarcia/GO-Twitter/src/domain"
	"github.com/alejandroagarcia/GO-Twitter/src/service"
)

func main() {
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewFileTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {
			var newTweet domain.Tweet

			defer c.ShowPrompt(true)

			choice := c.MultiChoice([]string{
				"TextTweet",
				"ImageTweet",
				"QuoteTweet",
				"Salir",
			}, "Qué tipo de Tweet quieres crear?")

			if choice == 0 {
				c.Print("Write your username: ")

				user := c.ReadLine()

				c.Print("Write your tweet: ")

				text := c.ReadLine()

				newTweet = domain.NewTextTweet(user, text)

			} else if choice == 1 {
				c.Print("Write your username: ")

				user := c.ReadLine()

				c.Print("Write your tweet: ")

				text := c.ReadLine()

				c.Print("Write your link image: ")

				link := c.ReadLine()

				newTweet = domain.NewImageTweet(user, text, link)

			} else if choice == 2 {
				c.Print("Escribe el ID del Tweet a citar: ")

				id, _ := strconv.Atoi(c.ReadLine())

				quotedTweet := tweetManager.GetTweetById(id)

				if quotedTweet == nil {
					c.Println("ID inexistente.")
					time.Sleep(2000)
					return
				}

				c.Print("Write your username: ")

				user := c.ReadLine()

				c.Print("Write your tweet: ")

				text := c.ReadLine()

				newTweet = domain.NewQuoteTweet(user, text, quotedTweet)

			} else {
				return
			}

			_, err := tweetManager.PublishTweet(newTweet)

			if err != nil {
				c.Print(err.Error())
			} else {
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
			} else {
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
			} else {
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
			} else {
				c.Println("ID inexistente.")
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "getTweetsContainText",
		Help: "Obtiene los tweets que contienen la palabra indicada",
		Func: func(c *ishell.Context) {
			var searchResult = make(chan domain.Tweet)
			var quit = make(chan string)

			defer c.ShowPrompt(true)

			c.Print("Ingrese la palabra: ")
			query := c.ReadLine()

			tweetManager.SearchTweetsContaining(query, searchResult, quit)

			for {
				select {
				case tweet := <-searchResult:
					c.Println(tweet)
				case results := <-quit:
					if results == "finishedEmpty" {
						c.Println("No se encontraron coincidencias.")
					}

					return

				}
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
			} else {
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
			} else {
				c.Println("Usuario inexistente o no contiene tweets.")
			}

			return
		},
	})

	shell.Run()

}
