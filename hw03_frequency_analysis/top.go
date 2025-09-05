package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

// Структура для хранения слова и его частоты встречаемости.
type WordFrequency struct {
	Word      string // Слово
	Frequency int    // Частота встречаемости
}

// removePunctuation очищает текст от знаков препинания, оставляя только буквы, пробелы и разрешенные символы в словах.
func removePunctuation(text string) string {
	// Удаляем все, кроме букв, пробелов и разрешенных символов в словах
	re := regexp.MustCompile(`[^\p{L}\s'’-]`)
	text = re.ReplaceAllString(text, "")

	// Убираем одиночные разрешенные символы (которые не между буквами)
	re = regexp.MustCompile(`(^|\s)['’-]|['’-]($|\s)`)
	text = re.ReplaceAllString(text, "$1$2")

	return text
}

// preprocessing обрабатывает список слов, подсчитывает частоту встречаемости и сортирует по убыванию частоты.
// При одинаковой частоте слова сортируются по алфавиту.
func preprocessing(words []string) []*WordFrequency {
	// Создаем map для подсчета частоты слов
	freq := make(map[string]int, len(words)/2)
	for _, word := range words {
		freq[word]++
	}

	// Создаем слайс структур WordFrequency для хранения результатов
	wordFreqs := make([]*WordFrequency, 0, len(freq))
	for word, count := range freq {
		wordFreqs = append(wordFreqs, &WordFrequency{word, count})
	}

	// Сортируем слайс по убыванию частоты, а при равной частоте - по алфавиту
	sort.Slice(wordFreqs, func(i, j int) bool {
		if wordFreqs[i].Frequency == wordFreqs[j].Frequency {
			return wordFreqs[i].Word < wordFreqs[j].Word // сортировка по алфавиту при одинаковой частоте
		}
		return wordFreqs[i].Frequency > wordFreqs[j].Frequency
	})
	return wordFreqs
}

// Top10 возвращает топ-10 самых часто встречающихся слов в тексте.
func Top10(str string) []string {
	if str == "" {
		return nil
	}

	// Приводим текст к нижнему регистру
	str = strings.ToLower(str)
	// Очищаем текст от знаков препинания
	cleaned := removePunctuation(str)

	// Разбиваем текст на слова
	words := strings.Fields(cleaned)
	if len(words) == 0 {
		return nil
	}

	// Обрабатываем слова: подсчитываем частоту и сортируем
	wordFreqs := preprocessing(words)

	// Выбираем первые 10 слов (или меньше, если слов меньше 10)
	n := min(10, len(wordFreqs))
	result := make([]string, n)
	for i := range n {
		result[i] = wordFreqs[i].Word
	}
	return result
}
