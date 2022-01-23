package hw03frequencyanalysis

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

var englishText = "cat and dog, one dog,two cats and one man"
var japanText = "飛翔(はばた)いたら 戻らないと言って1 目指したのは 蒼い " +
	"蒼い あの空 “悲しみ”はまだ覚えられず ”切なさ”は今つかみはじめた " +
	"あなたへと抱く この感情も 今”言葉”に変わっていく 未知なる世界の " +
	"遊迷(ゆめ)から目覚めて この羽根を広げ 飛び立つ 飛翔(はばた)いたら " +
	"戻らないと言って 目指したのは 白い 白い あの雲 突き抜けたら " +
	"みつかると知って 振り切るほど 蒼い 蒼い あの空 蒼い 蒼い あの空 " +
	"蒼い 蒼い あの空"
var dashText = "- - - -- -- --- -- -- - - "
var dashWordText = "какой-то очень какойто какойто какой то очень очень-очень очень-очень"

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		assert.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"}
			assert.Subset(t, expected, Top10(text))
		} else {
			expected := []string{"он", "и", "а", "что", "ты", "не", "если", "-", "то", "Кристофер"}
			assert.ElementsMatch(t, expected, Top10(text))
		}
	})

	t.Run("english words test", func(t *testing.T) {
		expected := []string{"cat", "and", "dog", "one", "two", "cats", "man"}
		assert.Subset(t, expected, Top10(englishText))
	})
	t.Run("japan words test", func(t *testing.T) {
		expected := []string{"蒼い", "あの空", "白い", "目指したのは", "飛翔(はばた)いたら", "“悲しみ”はまだ覚えられず", "”切なさ”は今つかみはじめた", "あなたへと抱く", "あの雲", "この感情も"}
		assert.Subset(t, expected, Top10(japanText))
	})
	t.Run("dash test", func(t *testing.T) {
		assert.Len(t, Top10(dashText), 0)
	})
	t.Run("dash word test", func(t *testing.T) {
		expected := []string{"какой-то", "очень", "очень-очень", "какойто", "какой", "то"}
		assert.Subset(t, expected, Top10(dashWordText))
	})
}
