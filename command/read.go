package command

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const compulsory = "一右雨円王音下火花貝学気九休玉金空月犬見五口校左三山子四糸字耳七車手十出女小上森人水正生青夕石赤千川先早草足村大男竹中虫町天田土二日入年白八百文木本名目立力林六引羽雲園遠何科夏家歌画回会海絵外角楽活間丸岩顔汽記帰弓牛魚京強教近兄形計元言原戸古午後語工公広交光考行高黄合谷国黒今才細作算止市矢姉思紙寺自時室社弱首秋週春書少場色食心新親図数西声星晴切雪船線前組走多太体台地池知茶昼長鳥朝直通弟店点電刀冬当東答頭同道読内南肉馬売買麦半番父風分聞米歩母方北毎妹万明鳴毛門夜野友用曜来里理話悪安暗医委意育員院飲運泳駅央横屋温化荷界開階寒感漢館岸起期客究急級宮球去橋業曲局銀区苦具君係軽血決研県庫湖向幸港号根祭皿仕死使始指歯詩次事持式実写者主守取酒受州拾終習集住重宿所暑助昭消商章勝乗植申身神真深進世整昔全相送想息速族他打対待代第題炭短談着注柱丁帳調追定庭笛鉄転都度投豆島湯登等動童農波配倍箱畑発反坂板皮悲美鼻筆氷表秒病品負部服福物平返勉放味命面問役薬由油有遊予羊洋葉陽様落流旅両緑礼列練路和愛案以衣位茨印英栄媛塩岡億加果貨課芽賀改械害街各覚潟完官管関観願岐希季旗器機議求泣給挙漁共協鏡競極熊訓軍郡群径景芸欠結建健験固功好香候康佐差菜最埼材崎昨札刷察参産散残氏司試児治滋辞鹿失借種周祝順初松笑唱焼照城縄臣信井成省清静席積折節説浅戦選然争倉巣束側続卒孫帯隊達単置仲沖兆低底的典伝徒努灯働特徳栃奈梨熱念敗梅博阪飯飛必票標不夫付府阜富副兵別辺変便包法望牧末満未民無約勇要養浴利陸良料量輪類令冷例連老労録圧囲移因永営衛易益液演応往桜可仮価河過快解格確額刊幹慣眼紀基寄規喜技義逆久旧救居許境均禁句型経潔件険検限現減故個護効厚耕航鉱構興講告混査再災妻採際在財罪殺雑酸賛士支史志枝師資飼示似識質舎謝授修述術準序招証象賞条状常情織職制性政勢精製税責績接設絶祖素総造像増則測属率損貸態団断築貯張停提程適統堂銅導得毒独任燃能破犯判版比肥非費備評貧布婦武復複仏粉編弁保墓報豊防貿暴脈務夢迷綿輸余容略留領歴胃異遺域宇映延沿恩我灰拡革閣割株干巻看簡危机揮貴疑吸供胸郷勤筋系敬警劇激穴券絹権憲源厳己呼誤后孝皇紅降鋼刻穀骨困砂座済裁策冊蚕至私姿視詞誌磁射捨尺若樹収宗就衆従縦縮熟純処署諸除承将傷障蒸針仁垂推寸盛聖誠舌宣専泉洗染銭善奏窓創装層操蔵臓存尊退宅担探誕段暖値宙忠著庁頂腸潮賃痛敵展討党糖届難乳認納脳派拝背肺俳班晩否批秘俵腹奮並陛閉片補暮宝訪亡忘棒枚幕密盟模訳郵優預幼欲翌乱卵覧裏律臨朗論"

func isKanji(r rune) bool {
	reg := regexp.MustCompile(`[\x{2E80}-\x{2FDF}々〇〻\x{3400}-\x{4DBF}\x{4E00}-\x{9FFF}\x{F900}-\x{FAFF}\x{20000}-\x{2FFFF}]`)
	return reg.MatchString(string(r))
}

func isCompulsory(r rune) bool {
	return strings.Contains(compulsory, string(r))
}

func isHiraKata(r rune) bool {
	reg := regexp.MustCompile(`[\x{3041}-\x{3096}\x{30A1}-\x{30FA}]`)
	return reg.MatchString(string(r))
}

func isOK(r rune) bool {
	return isHiraKata(r) || isCompulsory(r)
}

func isAllOK(k string) bool {
	j := false
	for _, r := range k {
		if isOK(r) {
			j = true
		}
	}
	return j
}

func yorn(q string) bool {
	result := true
	fmt.Printf("%s", q)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i := scanner.Text()

		if i == "" || i == "Y" || i == "y" {
			break
		} else if i == "N" || i == "n" {
			result = false
			break
		} else {
			fmt.Println("YかNかでお願いしますっ！")
			fmt.Print(q)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("\033[1F\033[2K")
	return result
}

func colorize(l string) string {
	colored := ""
	for _, r := range l {
		if isKanji(r) {
			if isCompulsory(r) {
				colored = colored + "\033[30m\033[47m" + string(r) + "\033[0m"
			} else {
				colored = colored + "\033[41m" + string(r) + "\033[0m"
			}
		} else {
			colored = colored + string(r)
		}
	}
	return colored
}

func separate(l string) []string {
	result := []string{}
	b := isKanji([]rune(l)[0])
	var s string
	for _, r := range l {
		if b != isKanji(r) {
			result = append(result, s)
			b = isKanji(r)
			s = string(r)
		} else {
			s = s + string(r)
		}
	}
	result = append(result, s)
	return result
}

func filterKanji(s []string) []string {
	result := []string{}
	if !isKanji([]rune(s[0])[0]) {
		s = s[1:]
	}
	for i, t := range s {
		if i%2==0 {
			for _, r := range t {
				if !isCompulsory(r) {
					result = append(result, t)
					break
				}
			}
		}
	}
	return result
}

func question(k string) string {
	a := ""
	fmt.Printf("読み方、教えてください！ :%s => ", k)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i := scanner.Text()

		if isAllOK(i) {
			a = i
		} else {
			fmt.Println("あたしの読める字でお願いします！")
			fmt.Printf("読み方、教えてください！ :%s => ", k)
			continue
		}

		if yorn(a + " で大丈夫ですか？[y/n]:") {
			break
		} else {
			fmt.Printf("もういっかいお願いします！ :%s => ", k)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return a
}

func CmdRead(c *cli.Context) error {
	// open file
	f, err := os.Open(c.Args().Get(0))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer f.Close()

	// read all text
	b, _ := ioutil.ReadAll(f)

	// red := color.New(color.BgRed, color.Bold)
	// wb := color.New(color.FgBlack, color.BgWhite)

	// TODO: kanji count
	// var total int
	// var kanji int
	// var unknown int

	slice := strings.Split(string(b), "\n")

	var draft string

	for r, l := range slice {
		colored := colorize(l)
		if strings.Contains(colored, "\033[41m") {
			fmt.Printf("%d:%s\n", r+1, colored)
			// fmt.Printf("%s\n", strings.Join(separate(l), ","))
			for _, k := range filterKanji(separate(l)) {
				a := question(k)
				l = strings.Replace(l, k, a, 1)
				fmt.Printf("%d:%s\n", r+1, colorize(l))
			}
			
			draft = draft + l + "\n"
		} else {
			draft = draft + l + "\n"
		}
	}
	fmt.Println("\n--------------------------------------------------\n")
	fmt.Println(draft)
	fmt.Println("\n--------------------------------------------------\n")
	fmt.Printf("全部読めましたーっ！\n教えてくれて、ありがとうございます！")


	// for _, r := range string(b) {
	// 	total++
	// 	if isKanji(r) {
	// 		kanji++
	// 		if isCompulsory(r) {
	// 			fmt.Printf("%s", string(r))
	// 		} else {
	// 			unknown++
	// 			red.Printf("%s", string(r))
	// 		}
	// 	} else {
	// 		fmt.Printf("%s", string(r))
	// 	}
	// }

	// TODO: ここ関数化しよう
	// 果穂さんの使った漢字追加用
	// slice := strings.Split(string(b), "\n")

	// for _, l := range slice {
	// 	if strings.Contains(l, "果穂「") {
	// 		for _, r := range l {
	// 			if isKanji(r) {
	// 				if isCompulsory(r) {
	// 					// fmt.Printf("%s", string(r))
	// 				} else {
	// 					// red.Printf("%s", string(r))
	// 					fmt.Printf("%s\n", string(r))
	// 				}
	// 			} else {
	// 				// fmt.Printf("%s", string(r))
	// 			}
	// 		}
	// 		// fmt.Printf("\n")
	// 	}
	// }

	// fmt.Printf("文字数は全部で%d文字で、そのうち漢字は%d文字でした！\n", total, kanji)

	// if unknown != 0 {
	// 	fmt.Printf("読めない漢字が%d文字ありました……\n", unknown)
	// } else {
	// 	fmt.Println("全部読めましたーっ！")
	// }
	return nil
}
