package main

import (
	_ "github.com/lib/pq"

	"github.com/nadavbm/nogobk/api/dat"
	"github.com/nadavbm/nogobk/api/server"
	"github.com/nadavbm/nogobk/pkg/logger"
)

func main() {
	l := logger.SetLogger()

	dat.InitDB()

	l.Info("starting server on port 8081")
	s := server.NewServer(*l, ":8081")

	err := s.Run()
	if err != nil {
		panic(err)
	}

}
