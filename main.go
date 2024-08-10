package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel{
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}


func main()  {
	// Find .env file
 err := godotenv.Load(".env")
 if err != nil{
  log.Fatalf("Error loading .env file: %s", err)
 }

 bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

 go printCommandEvents(bot.CommandEvents())

 bot.Command("my year-of-birth is <year>", &slacker.CommandDefinition{
	Description: "year-of-birth calc",
	Examples: []string{"my year-of-birth is 1998"},
	Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter)  {
		year := request.Param("year")
		yob, err := strconv.Atoi(year)
		if err != nil {	
			fmt.Println("error")
		}
		age := 2024-yob
		r :=fmt.Sprintf("Age is %d", age)
		response.Reply(r) 
	},
 })

 ctx, cancel := context.WithCancel(context.Background())
 defer cancel()

 err2 := bot.Listen(ctx)
 if err2 != nil{
	log.Fatal(err2)
 }
}