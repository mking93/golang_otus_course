package main //hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

func Top10(str string) []string {
	words := strings.Split(str, " ")

	var m map[string]int
	m = getCountUniqueWords(words)

	res := getTop10WordsAlphabetically(m)
	fmt.Println(res)
	return res
}

func getCountUniqueWords(words []string) map[string]int {
	var m map[string]int
	m = make(map[string]int)
	for i := 0; i < len(words); i++ {
		if val, ok := m[words[i]]; ok {
			m[words[i]] = val + 1
			m[words[i]]++
		} else {
			m[words[i]] = 1
		}
	}
	return m
}

func getTop10WordsAlphabetically(m map[string]int) []string {
	type kv struct {
		k string
		v int
	}

	topByValues := make([]kv, 0, len(m))
	for k, v := range m {
		topByValues = append(topByValues, kv{k, v})
	}

	sort.Slice(topByValues, func(i, j int) bool {
		return topByValues[i].v > topByValues[j].v
	})

	count := 0
	if len(topByValues) > 10 {
		count = 10
	} else {
		count = len(topByValues) - 1
	}

	top10 := make([]kv, 0, count)
	for k := 0; k < count; k++ {
		top10 = append(top10, topByValues[k])
	}

	sort.Slice(top10, func(i, j int) bool {
		return top10[i].k < top10[j].k
	})

	top10ByKeys := make([]string, 0, 10)
	for r := 0; r < len(top10); r++ {
		top10ByKeys = append(top10ByKeys, top10[r].k)
	}

	return top10ByKeys
}

func main() {
	//Top10("cat and dog, one dog,two cats and one man")
	//Top10("")
	//Top10("Вчера я приехал в Пятигорск, нанял квартиру на краю города, на самом высоком месте, у подошвы Машука: во время грозы облака будут спускаться до моей кровли. Нынче в пять часов утра, когда я открыл окно, моя комната наполнилась запахом цветов, растуших в скромном палисаднике. Ветки цветущих черешен смотрят мне в окна, и ветер иногда усыпает мой письменный стол их белыми лепестками. Вид с трех сторон у меня чудесный. На запад пятиглавый Бешту синеет, как «последняя туча рассеянной бури»; на север поднимается Машук, как мохнатая персидская шапка, и закрывает всю эту часть небосклона; на восток смотреть веселее: внизу передо мною пестреет чистенький, новенький городок, шумят целебные ключи, шумит разноязычная толпа, — а там, дальше, амфитеатром громоздятся горы все синее и туманнее, а на краю горизонта тянется серебряная цепь снеговых вершин, начинаясь Казбеком и оканчиваясь двуглавым Эльборусом... Весело жить в такой земле! Какое-то отрадное чувство разлито во всех моих жилах. Воздух чист и свеж, как поцелуй ребенка; солнце ярко, небо сине — чего бы, кажется, больше? — зачем тут страсти, желания, сожаления?.. Однако пора. Пойду к Елизаветинскому источнику: там, говорят, утром собирается все водяное общество.Спустясь в середину города, я пошел бульваром, где встретил несколько печальных групп, медленно подымающихся в гору; то были большею частию семейства степных помещиков; об этом можно было тотчас догадаться по истертым, старомодным сюртукам мужей и по изысканным нарядам жен и дочерей; видно, у них вся водяная молодежь была уже на перечете, потому что они на меня посмотрели с нежным любопытством: петербургский покрой сюртука ввел их в заблуждение, но, скоро узнав армейские эполеты, они с негодованием отвернулись.\nЖены местных властей, так сказать хозяйки вод, были благосклоннее; у них есть лорнеты, они менее обращают внимания на мундир, они привыкли на Кавказе встречать под нумерованной пуговицей пылкое сердце и под белой фуражкой образованный ум. Эти дамы очень милы; и долго милы! Всякий год их обожатели сменяются новыми, и в этом-то, может быть, секрет их неутомимой любезности. Подымаясь по узкой тропинке к Елизаветинскому источнику, я обогнал толпу мужчин, штатских и военных, которые, как я узнал после, составляют особенный класс людей между чающими движения воды. Они пьют — однако не воду, гуляют мало, волочатся только мимоходом; они играют и жалуются на скуку. Они франты: опуская свой оплетенный стакан в колодец кислосерной воды, они принимают академические позы: штатские носят светло-голубые галстуки, военные выпускают из-за воротника брыжи. Они исповедывают глубокое презрение к провинциальным домам и вздыхают о столичных аристократических гостиных, куда их не пускают.\nНаконец вот и колодец... На площадке близ него построен домик с красной кровлею над ванной, а подальше галерея, где гуляют во время дождя. Несколько раненых офицеров сидели на лавке, подобрав костыли, — бледные, грустные. Несколько дам скорыми шагами ходили взад и вперед по площадке, ожидая действия вод. Между ними были два-три хорошеньких личика. Под виноградными аллеями, покрывающими скат Машука, мелькали порою пестрые шляпки любительниц уединения вдвоем, потому что всегда возле такой шляпки я замечал или военную фуражку или безобразную круглую шляпу. На крутой скале, где построен павильон, называемый Эоловой Арфой, торчали любители видов и наводили телескоп на Эльборус; между ними было два гувернера с своими воспитанниками, приехавшими лечиться от золотухи.\nЯ остановился, запыхавшись, на краю горы и, прислонясь к углу домика, стал рассматривать окрестность, как вдруг слышу за собой знакомый голос:\n— Печорин! давно ли здесь?\nОборачиваюсь: Грушницкий! Мы обнялись. Я познакомился с ним в действующем отряде. Он был ранен пулей в ногу и поехал на воды с неделю прежде меня. Грушницкий — юнкер. Он только год в службе, носит, по особенному роду франтовства, толстую солдатскую шинель. У него георгиевский солдатский крестик. Он хорошо сложен, смугл и черноволос; ему на вид можно дать двадцать пять лет, хотя ему едва ли двадцать один год. Он закидывает голову назад, когда говорит, и поминутно крутит усы левой рукой, ибо правою опирается на костыль. Говорит он скоро и вычурно: он из тех людей, которые на все случаи жизни имеют готовые пышные фразы, которых просто прекрасное не трогает и которые важно драпируются в необыкновенные чувства, возвышенные страсти и исключительные страдания. Производить эффект — их наслаждение; они нравятся романтическим провинциалкам до безумия. Под старость они делаются либо мирными помещиками, либо пьяницами — иногда тем и другим. В их душе часто много добрых свойств, но ни на грош поэзии. Грушницкого страсть была декламировать: он закидывал вас словами, как скоро разговор выходил из круга обыкновенных понятий; спорить с ним я никогда не мог. Он не отвечает на ваши возражения, он вас не слушает. Только что вы остановитесь, он начинает длинную тираду, по-видимому имеющую какую-то связь с тем, что вы сказали, но которая в самом деле есть только продолжение его собственной речи.\nОн довольно остер: эпиграммы его часто забавны, но никогда не бывают метки и злы: он никого не убьет одним словом; он не знает людей и их слабых струн, потому что занимался целую жизнь одним собою. Его цель — сделаться героем романа. Он так часто старался уверить других в том, что он существо, не созданное для мира, обреченное каким-то тайным страданиям, что он сам почти в этом уверился. Оттого-то он так гордо носит свою толстую солдатскую шинель. Я его понял, и он за это меня не любит, хотя мы наружно в самых дружеских отношениях. Грушницкий слывет отличным храбрецом; я его видел в деле; он махает шашкой, кричит и бросается вперед, зажмуря глаза. Это что-то не русская храбрость!..\nЯ его также не люблю: я чувствую, что мы когда-нибудь с ним столкнемся на узкой дороге, и одному из нас несдобровать.\nПриезд его на Кавказ — также следствие его романтического фанатизма: я уверен, что накануне отъезда из отцовской деревни он говорил с мрачным видом какой-нибудь хорошенькой соседке, что он едет не так, просто, служить, но что ищет смерти, потому что... тут, он, верно, закрыл глаза рукою и продолжал так: «Нет, вы (или ты) этого не должны знать! Ваша чистая душа содрогнется! Да и к чему? Что я для вас! Поймете ли вы меня?» — и так далее.\nОн мне сам говорил, что причина, побудившая его вступить в К. полк, останется вечною тайной между им и небесами.\nВпрочем, в те минуты, когда сбрасывает трагическую мантию, Грушницкий довольно мил и забавен. Мне любопытно видеть его с женщинами: тут-то он, я думаю, старается!\nМы встретились старыми приятелями. Я начал его расспрашивать об образе жизни на водах и о примечательных лицах.\n— Мы ведем жизнь довольно прозаическую, — сказал он, вздохнув, — пьющие утром воду — вялы, как все больные, а пьющие вино повечеру — несносны, как все здоровые. Женские общества есть; только от них небольшое утешение: они играют в вист, одеваются дурно и ужасно говорят по-французски. Нынешний год из Москвы одна только княгиня Лиговская с дочерью; но я с ними незнаком. Моя солдатская шинель — как печать отвержения. Участие, которое она возбуждает, тяжело, как милостыня.\nВ эту минуту прошли к колодцу мимо нас две дамы: одна пожилая, другая молоденькая, стройная. Их лиц за шляпками я не разглядел, но они одеты были по строгим правилам лучшего вкуса: ничего лишнего! На второй было закрытое платье gris de perles 1, легкая шелковая косынка вилась вокруг ее гибкой шеи. Ботинки couleur puce 2 стягивали у щиколотки ее сухощавую ножку так мило, что даже не посвященный в таинства красоты непременно бы ахнул, хотя от удивления. Ее легкая, но благородная походка имела в себе что-то девственное, ускользающее от определения, но понятное взору. Когда она прошла мимо нас, от нее повеяло тем неизъяснимым ароматом, которым дышит иногда записка милой женщины.")
}
