package main

import (
	"github.com/liuguangw/forumx/app/cmd"
	"log"
	"os"
)

func main() {
	if err := cmd.Execute(os.Args); err != nil {
		log.Fatalln(err)
	}
}
