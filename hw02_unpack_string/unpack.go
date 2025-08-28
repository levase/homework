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

	runes := []rune(str)
	var result strings.Builder

	i := 0
	for i < len(runes) {
		current := runes[i]
		// Если текущий символ - цифра (это ошибка, цифра должна следовать только после символа)
		if unicode.IsDigit(current) {
			return "", ErrInvalidString
		}
		// Если текущий символ - обратный слеш (экранирование)
		if current == '\\' {
			// Проверка: слеш не может быть последним
			if i+1 >= len(runes) {
				return "", ErrInvalidString
			}
			// Экранированный символ добавляем как есть
			escapedChar := runes[i+1]
			// Проверяем, есть ли после экранированного символа цифра
			if i+2 < len(runes) && unicode.IsDigit(runes[i+2]) {
				count, err := strconv.Atoi(string(runes[i+2]))
				if err != nil {
					return "", ErrInvalidString
				}
				if count > 0 {
					result.WriteString(strings.Repeat(string(escapedChar), count))
				}
				i += 3
			} else {
				result.WriteRune(escapedChar)
				i += 2
			}
			continue
		}
		// Обычный символ (буква, \n, \t, руна и т.д.)
		// Проверяем, есть ли следующая цифра (только одна цифра!)
		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			// Разрешены только однозначные цифры (не числа)
			count, err := strconv.Atoi(string(runes[i+1]))
			if err != nil {
				return "", ErrInvalidString
			}
			if count > 0 {
				result.WriteString(strings.Repeat(string(current), count))
			}
			// Если count == 0, символ не добавляется
			i += 2
		} else {
			result.WriteRune(current)
			i++
		}
	}
	return result.String(), nil
}
