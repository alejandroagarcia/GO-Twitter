package service

import "github.com/alejandroagarcia/GO-Twitter/src/domain"
import (
	"os"
)

type FileTweetWriter struct {
}

func NewFileTweetWriter() *FileTweetWriter {
	var ftw FileTweetWriter
	return &ftw
}

func (ftw *FileTweetWriter) Write(tweet domain.Tweet) {
	f, _ := os.Create("./tweets")
	go func() {
		defer f.Close()

		f.WriteString(tweet.String() + "\n")

		f.Sync()
	}()

}
