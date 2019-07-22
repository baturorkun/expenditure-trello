package lib

import (
	"expenditure/utils"
	"github.com/olekukonko/tablewriter"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func Form2Map(form url.Values) (res []map[string]string) {

	log.Printf(">> %+v", form)

	re := regexp.MustCompile("[0-9]")

	//res = make([][]string, 50)

	for i, v := range form {

		match := re.FindString(i)

		if match == "" {
			continue
		}

		n, _ := strconv.Atoi(match)

		field := strings.Replace(i, match, "", -1)

		if !utils.KeyInMap(n, res) {
			m := make(map[string]string)
			res = append(res, m)
		}

		log.Printf("%s %s", i, v[0])

		//res[n] = append(res[n], v[0])
		res[n][field] = v[0]
	}

	return res
}

func Map2Array(val []map[string]string) (res [][]string) {

	for _, m := range val {

		res = append(res, []string{m["date"], m["expense"], m["currency"], m["notes"]})
	}

	return
}

func CreateMDTable(data [][]string) string {

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetRowLine(true)
	table.SetHeader([]string{"Date", "Expense", "Amount", "Currency", "Notes"})
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.SetCenterSeparator("|")
	table.SetRowSeparator("-")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()

	return tableString.String()
}
