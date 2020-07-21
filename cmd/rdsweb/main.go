package main

import "github.com/gookit/gcli/v2"

func main() {
	app := gcli.NewApp()

	app.AddCommand(&gcli.Command{
		Name: "serve",
		UseFor: "start the application http serve",
		Func: func(c *gcli.Command, args []string) error {


			return nil
		},
	})

	app.Run()
}
