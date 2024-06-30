package main

import (
	"context"
	"fmt"

	"github.com/VassilevEmil/goMicro/application"
)

func main() {

	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println(err)
	}

}
