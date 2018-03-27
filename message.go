package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type macroPercentages struct {
	carbs   int
	protein int
	fat     int
}

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
	table, buffer := newTable([]string{"Macros", "Grams", "Percent"}, 10)

	m, err := newMacroPercentages(diary)
	if err != nil {
		return "Error parsing macros"
	}

	table.Append([]string{"Carbs", diary.Total.Carbs, fmt.Sprintf("%d%%", m.carbs)})
	table.Append([]string{"Protein", diary.Total.Protein, fmt.Sprintf("%d%%", m.protein)})
	table.Append([]string{"Fat", diary.Total.Fat, fmt.Sprintf("%d%%", m.fat)})
	table.Render()
	return "```" + buffer.String() + "```"
}

func calsMessage(diary *Diary) string {
	table, buffer := newTable([]string{"Food", "Calories"}, 17)

	for _, v := range formatTableData(diary) {
		table.Append(v)
	}

	table.Render()
	return "```" + buffer.String() + "```"
}

func newTable(headers []string, colWidth int) (*tablewriter.Table, *bytes.Buffer) {
	buffer := new(bytes.Buffer)
	table := tablewriter.NewWriter(buffer)
	table.SetColWidth(colWidth)
	table.SetHeader(headers)
	table.SetRowLine(true)
	return table, buffer
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

func newMacroPercentages(diary *Diary) (macroPercentages, error) {
	m := macroPercentages{}
	carbs, cErr := parseMacro(diary.Total.Carbs)
	protein, pErr := parseMacro(diary.Total.Protein)
	fat, fErr := parseMacro(diary.Total.Fat)
	if cErr != nil || pErr != nil || fErr != nil {
		return m, errors.New("Error parsing macros")
	}

	total := carbs + protein + fat
	m = macroPercentages{
		carbs:   (100 * carbs) / total,
		protein: (100 * protein) / total,
		fat:     (100 * fat) / total,
	}

	return m, nil
}
func parseMacro(macro string) (int, error) {
	return strconv.Atoi(macro[:len(macro)-1])
}
