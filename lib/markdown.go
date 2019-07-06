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


func Form2Array(form url.Values) (res [][]string) {

	log.Printf(">> %+v", form)

	re := regexp.MustCompile("[0-9]")

	//res = make([][]string, 50)

	for i, v := range form {

		match := re.FindString(i)

		if match == "" {
			continue
		}

		n,_ := strconv.Atoi(match)

		if !utils.KeyInSlice(n, res) {
			res = append(res, []string{})
		}

		log.Printf("%s %s", i,v[0])

		res[n] = append(res[n], v[0])
	}

	return res
}

func CreateMDTable(data [][]string ) string  {

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetRowLine(true)
	table.SetHeader([]string{"Name", "Price", "Desc"})
	table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	table.SetCenterSeparator("|")
	table.SetRowSeparator("-")
	table.AppendBulk(data) // Add Bulk Data
	table.Render()

	return tableString.String()
}