package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var spamWords []string
var userWarnings = make(map[int64]int)

const maxWarnings = 5

func main() {
	// Устанавливаем токен
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
	}

	// Логируем начало работы бота
	log.Println("Бот успешно запущен и готов к приему сообщений.")

	// Загрузка спам-слов из файла
	loadSpamWords("spam.txt")

	// Настройка long polling
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	log.Println("Ожидание сообщений...")

	for update := range updates {
		if update.Message != nil {
			log.Printf("Получено сообщение от пользователя %s: %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.IsCommand() {
				log.Printf("Получена команда: %s", update.Message.Command())
				handleCommand(bot, update.Message)
			} else {
				if isSpam(update.Message.Text) {
					log.Printf("Сообщение содержит спам от пользователя %s. Увеличиваем счетчик предупреждений.", update.Message.From.UserName)
					handleSpam(bot, update.Message)
				} else {
					log.Printf("Сообщение от пользователя %s не содержит спама.", update.Message.From.UserName)
				}
			}
		}
	}
}

// Функция для загрузки спам-слов из файла
func loadSpamWords(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Ошибка открытия файла со спам-словами: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text()) // Удаляем лишние пробелы
		if word != "" {                           // Игнорируем пустые строки
			spamWords = append(spamWords, word)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка чтения файла со спам-словами: %v", err)
	}

	log.Printf("Загружены спам-слова: %v", spamWords) // Выводим загруженные спам-слова
}

// Функция для замены латинских символов на кириллические
func transliterateToRussian(input string) string {
	replacements := map[rune]rune{
		'A': 'А', 'a': 'а', 'B': 'В', 'E': 'Е', 'e': 'е', 'K': 'К', 'M': 'М', 'H': 'Н', 'O': 'О', 'o': 'о',
		'P': 'Р', 'C': 'С', 'c': 'с', 'T': 'Т', 'y': 'у', 'X': 'Х',
	}

	var result strings.Builder
	for _, char := range input {
		if replacement, found := replacements[char]; found {
			result.WriteRune(replacement)
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

// Функция для проверки, содержит ли сообщение спам
func isSpam(message string) bool {
	transliteratedMessage := transliterateToRussian(strings.ToLower(message)) // Преобразуем сообщение
	log.Printf("Транслитерированное сообщение: %s", transliteratedMessage)    // Для отладки
	for _, spamWord := range spamWords {
		if strings.Contains(transliteratedMessage, strings.ToLower(spamWord)) {
			return true
		}
	}
	return false
}

// Функция обработки спам-сообщений
func handleSpam(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	userID := message.From.ID
	userWarnings[userID]++
	log.Printf("Пользователь %s получил предупреждение %d/%d", message.From.UserName, userWarnings[userID], maxWarnings)

	if userWarnings[userID] >= maxWarnings {
		log.Printf("Пользователь %s часто спамит! Обратите внимание!.", message.From.UserName)
		delete(userWarnings, userID)
	} else {
		log.Printf("Удаление сообщения пользователя %s.", message.From.UserName)
		deleteMessage(bot, message.Chat.ID, message.MessageID)
	}
}

// Функция для удаления сообщения
func deleteMessage(bot *tgbotapi.BotAPI, chatID int64, messageID int) {
	msg := tgbotapi.NewDeleteMessage(chatID, messageID)
	if _, err := bot.Request(msg); err != nil {
		log.Printf("Ошибка удаления сообщения: %v", err)
	} else {
		log.Println("Сообщение успешно удалено.")
	}
}

// Функция обработки команд
func handleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Привет! Я бот, который борется со спамом.")
		bot.Send(msg)
		log.Println("Отправлено приветственное сообщение.")
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда.")
		bot.Send(msg)
		log.Printf("Получена неизвестная команда от пользователя %s.", message.From.UserName)
	}
}
