package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// Проверка: строка не должна быть пустой
	if str == "" {
		return "", nil
	}

	var result strings.Builder
	runes := []rune(str)
	i := 0

	for i < len(runes) {
		// Обрабатываем escape-последовательности
		if runes[i] == '\\' {
			if i+1 >= len(runes) {
				return "", ErrInvalidString
			}
			// Проверяем, что экранируется только допустимый символ
			if runes[i+1] != '\\' && !unicode.IsDigit(runes[i+1]) {
				return "", ErrInvalidString
			}
			// Экранированный символ добавляем как есть
			char := runes[i+1]
			i += 2

			// Проверяем, есть ли цифра после экранированного символа
			if i < len(runes) && unicode.IsDigit(runes[i]) {
				count, _ := strconv.Atoi(string(runes[i]))
				result.WriteString(strings.Repeat(string(char), count))
				i++
			} else {
				result.WriteRune(char)
			}
			continue
		}

		// Если текущий символ - цифра (это ошибка, цифра должна следовать только после символа)
		if unicode.IsDigit(runes[i]) {
			return "", ErrInvalidString
		}

		char := runes[i]
		i++

		// Проверяем, есть ли цифра после символа
		if i < len(runes) && unicode.IsDigit(runes[i]) {
			count, _ := strconv.Atoi(string(runes[i]))
			result.WriteString(strings.Repeat(string(char), count))
			i++
		} else {
			result.WriteRune(char)
		}
	}

	return result.String(), nil
}
