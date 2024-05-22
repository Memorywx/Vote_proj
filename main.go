package main

import (
	"WxVote/app"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app.Start()
}
