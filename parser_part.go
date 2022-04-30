package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type game_info struct {
	gname  string
	gid    string
	cprice int
	hcard  bool
	lcard  bool
	ocard  bool
	market bool
	err    bool
}

//Выгрузка всех доступных в базе игр
func GameInfo(game_id string) *game_info {
	var sceinfo game_info
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("steamcardexchange.net", "www.steamcardexchange.net"),
		colly.CacheDir("./cache"),
	)
	//	sceinfo.gid = game_id
	//Есть доступные карты <span class="card-amount green">
	//Карт нет или последняя - gray & red
	//Оверсток - orange
	//dom > first child > Attr
	c.OnHTML(".price-container", func(e *colly.HTMLElement) {
		if e.Text != "" {
			zzz := e.ChildAttrs("span", "class")
			fmt.Println(zzz[0])
			switch zzz[0] {
			case "card-amount green":
				sceinfo.hcard = true
			case "card-amount red", "card-amount gray":
				sceinfo.lcard = true
			case "card-amount orange":
				sceinfo.ocard = true
			default:
				break
			}
		}
	})

	//Название игры и ее статус <span class="game-title"><h2 class="empty">
	c.OnHTML(".game-title", func(e *colly.HTMLElement) {
		sceinfo.gname = strings.Replace(e.Text, " (Non-marketable - Trade-in disabled)", "", 1)
		if len(sceinfo.gname) == len(e.Text) {
			sceinfo.market = true
		}
	})

	//Цена <span class="game-price">Cards: X / Worth: Yc</span>
	c.OnHTML(".game-price", func(e *colly.HTMLElement) {
		//Разбитие текста на две половины
		splstr := strings.SplitAfterN(e.Text, " / Worth: ", -1)
		//Получение цены
		i, err := strconv.Atoi(strings.Replace(splstr[1], "c", "", 1))
		if err != nil {
			sceinfo.err = true
		} else {
			sceinfo.cprice = i
		}

		//		sceinfo.gname = strings.Replace(e.Text, " (Non-marketable - Trade-in disabled)", "", 1)
		//		if len(sceinfo.gname) == len(e.Text) {
		//			sceinfo.market = true
		//		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	//Переход на страницу с игрой
	c.Visit("https://www.steamcardexchange.net/index.php?inventorygame-appid-" + game_id)

	//Если не получилось получить имя игры или цену, то возвращать с ошибкой
	if len(sceinfo.gname) == 0 || sceinfo.cprice == 0 {
		sceinfo.err = true
	}
	return &sceinfo
}

//Выгрузка всех доступных в базе игр
func ScrapAll() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("steamcardexchange.net", "www.steamcardexchange.net"),
		colly.CacheDir("./cache"),
	)
	sce_info := [][]string{}
	sce_infon := [][]string{}
	c.OnHTML("option", func(e *colly.HTMLElement) {
		if e.Text != "" {
			sce_infon = [][]string{{e.Text, strings.Replace(e.Attr("value"), "index.php?gamepage-appid-", "", 1)}}
			sce_info = append(sce_info, sce_infon...)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.steamcardexchange.net/index.php?gamepage-appid-1273750")

	for _, p := range sce_info {
		fmt.Printf("%s with ID - %s\n", p[0], p[1])
	}
}
