package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

//NewCaloriesMessage builds the table of foods and calories for the diary
func NewCaloriesMessage(diary *Diary) string {
	message := calsMessage(diary)
	if len(message) > 2000 {
		totalStart := len(message) - 102
		return "```" + message[totalStart:]
	}
	return message
}

//NewMacrosMessage builds a table of macros for the diary
func NewMacrosMessage(diary *Diary) string {
	buffer := new(bytes.Buffer)
	table := tablewriter.NewWriter(buffer)
	table.SetColWidth(10)
	table.SetHeader([]string{"Macros", "Grams", "Percent"})
	table.SetRowLine(true)

	carbs, cErr := parseMacro(diary.Total.Carbs)
	protein, pErr := parseMacro(diary.Total.Protein)
	fat, fErr := parseMacro(diary.Total.Fat)
	if cErr != nil || pErr != nil || fErr != nil {
		return "Error parsing macros"
	}

	total := carbs + protein + fat
	cp := (100 * carbs) / total
	pp := (100 * protein) / total
	fp := (100 * fat) / total

	table.Append([]string{"Carbs", diary.Total.Carbs, fmt.Sprintf("%d%%", cp)})
	table.Append([]string{"Protein", diary.Total.Protein, fmt.Sprintf("%d%%", pp)})
	table.Append([]string{"Fat", diary.Total.Fat, fmt.Sprintf("%d%%", fp)})
	table.Render()
	return "```" + buffer.String() + "```"
}

func calsMessage(diary *Diary) string {
	buffer := new(bytes.Buffer)
	table := tablewriter.NewWriter(buffer)
	table.SetColWidth(17)
	table.SetHeader([]string{"Food", "Calories"})
	table.SetRowLine(true)
	for _, v := range formatTableData(diary) {
		table.Append(v)
	}
	table.Render()
	return "```" + buffer.String() + "```"
}

func formatTableData(diary *Diary) [][]string {
	var data [][]string
	meals := []string{"Breakfast", "Lunch", "Dinner", "Snacks"}
	for _, m := range meals {
		if f, ok := diary.Meals[m]; ok {
			data = append(data, []string{strings.ToUpper(m), ""})
			data = addFoods(data, f)
		}
	}
	data = append(data, []string{"Total", diary.Total.Calories})
	return data
}

func addFoods(data [][]string, foods []Food) [][]string {
	for _, food := range foods {
		name := formatFoodName(food.Name)
		data = append(data, []string{name, food.Calories})
	}
	return data
}

func formatFoodName(name string) string {
	strippedBrandSlice := strings.SplitN(name, "- ", 2)
	stripped := strippedBrandSlice[len(strippedBrandSlice)-1]
	if len(stripped) > 32 {
		return stripped[:30] + "..."
	}
	return stripped
}

func parseMacro(macro string) (int, error) {
	return strconv.Atoi(macro[:len(macro)-1])
}
