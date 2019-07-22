package lib

import (
	"expenditure/setting"
	"github.com/adlio/trello"
	"log"
	"net/http"
)

var client *trello.Client
var board *trello.Board
var list *trello.List

var err error

func Setup() {
	client = trello.NewClient(setting.TrelloSetting.AppKey, setting.TrelloSetting.Token)

	board, err = client.GetBoard(setting.TrelloSetting.BoardID, trello.Defaults())

	if err != nil {
		log.Fatalln("Error: Selecting Board:", err.Error())
	}

	lists, err := board.GetLists(trello.Defaults())

	if err != nil {
		log.Fatalln("Error: Selecting Lists of Board:", err.Error())
	}

	list = lists[setting.TrelloSetting.ListNumber]

}

func CreateCard(r *http.Request, filename string) (card *trello.Card) {

	form := r.PostForm

	m := Form2Map(form)

	arr := Map2Array(m)

	table := CreateMDTable(arr)

	log.Println(table)

	desc := form.Get("notes") + "\n" +
		form.Get("email") + "\n" +
		table

	card = &trello.Card{
		IDList: list.ID,
		Name:   form.Get("title"),
		Desc:   desc,
	}

	err = client.CreateCard(card, trello.Defaults())

	if err != nil {
		log.Fatalln("Error on creating card", err.Error())
	}

	if filename != "" {
		attachUrl := "http://" + r.Host + "/attach/exp_" + filename
		addAttach(card, attachUrl)
	}

	return
}

func addAttach(card *trello.Card, url string) {

	attach := trello.Attachment{URL: url, Name: "Invoice"}

	err = card.AddURLAttachment(&attach)

	if err != nil {
		log.Fatalf("Add attachment error : %v", err)
	}

}

/*
func AddCard(card *trello.Card) {

	err = list.AddCard(card, trello.Defaults())

	if err != nil {
		log.Fatalln("Error on adding card to list", err.Error())
	}

}
*/
