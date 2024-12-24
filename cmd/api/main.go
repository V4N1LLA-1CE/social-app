package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := &application{
		config: cfg,
	}

	log.Println("server listening on port :8080")
	log.Fatal(app.serve())
}
