package main

import "log"

func main() {
	cfg := config{
		addr: ":8080",
	}

	app := &application{
		config: cfg,
	}

	r := app.routes()

	log.Fatal(app.serve(r))
}
