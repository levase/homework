package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})
}

func TestTop10Additional(t *testing.T) {
	t.Run("single word", func(t *testing.T) {
		result := Top10("hello")
		expected := []string{"hello"}
		require.Equal(t, expected, result)
	})

	t.Run("multiple same words", func(t *testing.T) {
		result := Top10("word word word word")
		expected := []string{"word"}
		require.Equal(t, expected, result)
	})

	t.Run("alphabetical order for same frequency", func(t *testing.T) {
		text := "zebra apple banana apple banana zebra"
		result := Top10(text)
		expected := []string{"apple", "banana", "zebra"}
		require.Equal(t, expected, result)
	})

	t.Run("text with punctuation", func(t *testing.T) {
		text := "Hello, world! World? Hello... World: hello;"
		result := Top10(text)
		expected := []string{"hello", "world"}
		require.Equal(t, expected, result)
	})

	t.Run("text with mixed case", func(t *testing.T) {
		text := "Hello hello HELLO World world"
		result := Top10(text)
		expected := []string{"hello", "world"}
		require.Equal(t, expected, result)
	})

	t.Run("text with numbers", func(t *testing.T) {
		text := "test 123 123 123 test 456 456"
		result := Top10(text)
		expected := []string{"123", "456", "test"}
		require.Equal(t, expected, result)
	})

	t.Run("text with apostrophes", func(t *testing.T) {
		text := "don't can't won't don't can't"
		result := Top10(text)
		expected := []string{"can't", "don't", "won't"}
		require.Equal(t, expected, result)
	})

	t.Run("text with hyphens", func(t *testing.T) {
		text := "well-known state-of-the-art well-known"
		result := Top10(text)
		expected := []string{"well-known", "state-of-the-art"}
		require.Equal(t, expected, result)
	})

	t.Run("unicode characters", func(t *testing.T) {
		text := "привет мир привет test 测试 测试"
		result := Top10(text)
		expected := []string{"привет", "测试", "test", "мир"}
		require.Equal(t, expected, result)
	})

	t.Run("multiple spaces", func(t *testing.T) {
		text := "word    word   another     test"
		result := Top10(text)
		expected := []string{"word", "another", "test"}
		require.Equal(t, expected, result)
	})

	t.Run("leading and trailing spaces", func(t *testing.T) {
		text := "   hello world   "
		result := Top10(text)
		expected := []string{"hello", "world"}
		require.Equal(t, expected, result)
	})

	t.Run("empty string", func(t *testing.T) {
		result := Top10("")
		require.Nil(t, result)
	})

	t.Run("only whitespace", func(t *testing.T) {
		result := Top10("   \t\n\r   ")
		require.Nil(t, result)
	})

	t.Run("single character words", func(t *testing.T) {
		text := "a b c a b a"
		result := Top10(text)
		expected := []string{"a", "b", "c"}
		require.Equal(t, expected, result)
	})
}
