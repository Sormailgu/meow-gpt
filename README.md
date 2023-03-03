# Telega Bot - Integrating OpenAI APIs
This project aims to develop a Telegram bot that incorporates OpenAI’s APIs. The goal of this project is to automate tasks and increase human-bot interaction in a natural and intuitive way.

## Features

* Easy integration with OpenAI APIs to enrich the user experience.

* Ability for users to perform natural language processing (NLP) tasks, such as sentiment analysis and automated text summarization without the need for coding knowledge.

* Natural language understanding (NLU) capabilities for enhanced interaction.

* Automation of tasks specific to the needs of a particular organization by creating custom bots.

## Requirements

* [Golang](https://go.dev/) 1.18 or higher
* [Telegram Bot API](https://core.telegram.org/bots/api)
* [OpenAI API](https://openai.com/api)

## Getting Started

To install this project, simply clone the repository.

Create a bot token by going to the [Telegram web interface](https://core.telegram.org/bots#botfather).

Then, create a openAI API token by going to the [OpenAI web interface](https://platform.openai.com/).

Set the environment paras TELEGRAM_TOKEN & OPENAI_TOKEN value from above tokens.

Finally, start the bot with environment paras TELEGRAM_TOKEN & OPENAI_TOKEN.

## Docker
This project has a docker-compose.yml file, which will start the mkdocs application on your local machine and help you see changes instantly.
Set the environment paras TELEGRAM_TOKEN & OPENAI_TOKEN value and just run by docker compose. 
```
docker compose up
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
