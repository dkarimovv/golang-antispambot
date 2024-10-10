# EN
# Telegram Anti-Spam Bot

This project is a Telegram bot designed to remove spam messages from group chats. The project is written in Go and uses webhooks to process incoming messages and commands. It is easily deployable on Yandex Cloud using serverless functions.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Setup](#setup)

## Features

- **Spam Removal**: The bot filters spam messages based on a predefined list of forbidden words and automatically removes them from the chat.
- **User Management**: The bot can issue warnings, delete messages, and ban users for violating rules.
- **Webhook Support**: The bot interacts with the Telegram API using webhooks, reducing message processing delays.
- **Logging**: Detailed logs capture all bot actions, including message filtering and user management.
- **Yandex Cloud Deployment**: Easily deploy the bot using serverless functions in Yandex Cloud, simplifying support and scaling.

## Technologies

- **Go 1.20+**: The primary programming language used for bot development.
- **Telegram Bot API**: Used for integrating with Telegram.
- **Yandex Cloud Functions**: For deploying the bot in the cloud.
- **SQLite**: For storing user data and bot statistics.
- **Webhook**: The bot uses webhooks to communicate with Telegram for better performance.

## Setup

### Requirements

1. **Go**: Ensure you have Go 1.20 or newer installed.
2. **Telegram Bot**: Obtain your bot token from [BotFather](https://t.me/botfather) on Telegram.
3. **Yandex Cloud**: Set up a Yandex Cloud account and create a function to host the bot.

## Configuration
The bot uses configuration files for setup:

**spam_words.txt**: A file containing the list of forbidden words, one per line.
**config.yaml**: A file with settings, including webhook parameters, bot token, and other options.


# RU
# Telegram Anti-Spam Bot

Этот проект представляет собой бота для Telegram, который предназначен для удаления спам-сообщений из групповых чатов. Проект написан на языке Go и использует вебхуки для обработки входящих сообщений и команд. Проект легко разворачивается на Yandex Cloud с использованием безсерверных функций.

## Оглавление

- [Особенности](#особенности)
- [Технологии](#технологии)
- [Установка](#установка)

## Особенности

- **Удаление спама**: Бот фильтрует спам-сообщения по заранее заданному списку запрещённых слов и автоматически удаляет их из чата.
- **Управление пользователями**: Бот может выдавать предупреждения, удалять сообщения и исключать пользователей за нарушение правил.
- **Поддержка вебхуков**: Для взаимодействия с Telegram API бот использует вебхуки, что снижает задержку обработки сообщений.
- **Логирование**: Подробные лог-файлы, отображающие все действия бота, включая фильтрацию сообщений и управление пользователями.
- **Разворачивание на Yandex Cloud**: Бот легко развернуть с помощью безсерверных функций Yandex Cloud, что упрощает поддержку и масштабирование.

## Технологии

- **Go 1.20+**: Основной язык программирования для разработки бота.
- **Telegram Bot API**: Используется для интеграции с Telegram.
- **Yandex Cloud Functions**: Для развертывания бота в облаке.
- **SQLite**: Для хранения данных о пользователях и статистики бота.
- **Webhook**: Взаимодействие с Telegram через вебхуки для повышения производительности.

## Установка

### Требования

1. **Go**: Убедитесь, что у вас установлена версия Go 1.20 или новее.
2. **Telegram Bot**: Получите токен для своего бота у [BotFather](https://t.me/botfather) в Telegram.
3. **Yandex Cloud**: Настройте аккаунт в Yandex Cloud и создайте функцию для размещения бота.

## Конфигурация
Бот использует конфигурационные файлы для настройки:

**spam_words.txt**: Файл, содержащий список запрещённых слов, каждое на новой строке.
**config.yaml**: Файл с настройками, где указаны параметры вебхука, токен и другие параметры.
