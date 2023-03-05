# Telega Bot - Integrating OpenAI APIs in Golang
This project aims to develop a Telegram bot that incorporates OpenAI’s APIs. The goal of this project is to automate tasks and increase human-bot interaction in a natural and intuitive way.

## Features

* Easy integration with OpenAI APIs to enrich the Telegram Bot user experience.

* Automation of tasks specific to the needs of a particular organization by creating custom bots.

## Requirements

* [Golang](https://go.dev/) 1.18 or higher
* [Telegram Bot API](https://core.telegram.org/bots/api)
* [OpenAI API](https://openai.com/api)

## Getting Started

To install this project, simply clone the repository.

Create a bot token by going to the [Telegram web interface](https://core.telegram.org/bots#botfather).

Then, create a openAI API token by going to the [OpenAI web interface](https://platform.openai.com/).

Set the config file (/config/application.yml) paras TELEGRAM_BOT_TOKEN & OPENAI_API_KEY value from above tokens.

Finally, start the bot.
```
go run main.go
```

## Docker
This project has a docker-compose.yml file, which will start the mkdocs application on your local machine and help you see changes instantly.
Set the environment paras TELEGRAM_BOT_TOKEN & OPENAI_API_KEY value and just run by docker compose. 
```
docker compose up
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
