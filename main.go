package main

import (
	"github.com/liuguangw/forumx/cmd"
	"log"
	"os"
)

func main() {
	if err := cmd.Execute(os.Args); err != nil {
		log.Fatalln(err)
	}
}
