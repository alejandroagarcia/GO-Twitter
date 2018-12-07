package service

import "github.com/alejandroagarcia/GO-Twitter/src/domain"

type TweetWriter interface {
	Write(domain.Tweet)
}
