package lib

import (
	"expenditure/setting"
	"github.com/adlio/trello"
	"log"
	"net/url"
)

var client *trello.Client
var board *trello.Board
var list *trello.List

var err  error

func Setup() {
	client = trello.NewClient(setting.TrelloSetting.AppKey, setting.TrelloSetting.Token)

	board, err = client.GetBoard("KE4wqorD", trello.Defaults())

	if err != nil {
		log.Fatalln("Error: Selecting Board:", err.Error())
	}

	lists, err := board.GetLists(trello.Defaults())

	if err != nil {
		log.Fatalln("Error: Selecting Lists of Board:", err.Error())
	}

	list = lists[setting.TrelloSetting.ListNumber]

}

func CreateCard(form  url.Values) (card *trello.Card) {

	arr := Form2Array(form)

	table := CreateMDTable(arr)


	log.Println(table)

	desc :=  form.Get("notes") + "\n" +
		     form.Get("email") + "\n" +
			  table

	card = &trello.Card{
		IDList: list.ID,
		Name: form.Get("title"),
		Desc: desc,
	}

	err = client.CreateCard(card, trello.Defaults())

	if err != nil {
		log.Fatalln("Error on creating card", err.Error())
	}

	return
}
/*
func AddCard(card *trello.Card) {

	err = list.AddCard(card, trello.Defaults())

	if err != nil {
		log.Fatalln("Error on adding card to list", err.Error())
	}

}
*/
