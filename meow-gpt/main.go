package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"strings"
)

func main() {
	InitConfig()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	tgToken := viper.GetString("TELEGRAM_BOT_TOKEN")
	b, err := bot.New(tgToken, opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	var gptToken = viper.GetString("OPENAI_API_KEY")
	var gptModel = viper.GetString("OPENAI_MODEL")
	var gptTemperature = float32(viper.GetFloat64("OPENAI_TEMPERATURE"))
	var gptMaxTokens = viper.GetInt("OPENAI_MAX_TOKENS")
	var gptTopP = float32(viper.GetFloat64("OPENAI_TOP_P"))
	var gptFrequencyPenalty = float32(viper.GetFloat64("OPENAI_FREQUENCY_PENALTY"))
	var gptPresencePenalty = float32(viper.GetFloat64("OPENAI_PRESENCE_PENALTY"))

	client := openai.NewClient(gptToken)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:            gptModel,
			MaxTokens:        gptMaxTokens,
			Temperature:      gptTemperature,
			FrequencyPenalty: gptFrequencyPenalty,
			PresencePenalty:  gptPresencePenalty,
			TopP:             gptTopP,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: update.Message.Text,
				},
			},
		},
	)
	var respTxt = update.Message.Text
	if err != nil {
		respTxt = "GPT return error! Error message:\n" + err.Error()
	} else {
		respTxt = resp.Choices[0].Message.Content
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

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	fmt.Printf("%+v\n", viper.AllSettings())
}
