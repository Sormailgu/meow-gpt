package main

import (
	"context"
	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/spf13/viper"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	InitConfig()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	//tgToken := viper.GetString("telegram.token")
	tgToken := os.Getenv("TELEGRAM_TOKEN")
	b, err := bot.New(tgToken, opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	//gptToken := viper.GetString("openAI.token")
	gptToken := os.Getenv("OPENAI_TOKEN")
	gptClient := gogpt.NewClient(gptToken)
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 1024,
		Prompt:    update.Message.Text,
		//Temperature: 0,
		//TopP:        0,
	}
	resp, err := gptClient.CreateCompletion(ctx, req)

	var respTxt = update.Message.Text
	if err != nil {
		respTxt = "GPT return error! " + err.Error()
	} else {
		respTxt = resp.Choices[0].Text
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   respTxt,
	})
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("init error")
	}
}
