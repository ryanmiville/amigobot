package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

//NewCaloriesMessage builds the table of foods and calories for the diary
func NewCaloriesMessage(diary *Diary) string {
	message := diaryMessage(diary)
	if len(message) > 2000 {
		totalStart := len(message) - 102
		return "```" + message[totalStart:]
	}
	return message
}

//NewMacrosMessage fdsa
func NewMacrosMessage(diary *Diary) string {
	buffer := new(bytes.Buffer)
	table := tablewriter.NewWriter(buffer)
	table.SetColWidth(10)
	table.SetHeader([]string{"Macros", "Grams", "Percent"})
	table.SetRowLine(true)

	carbs := asInt(diary.Total.Carbs)
	protein := asInt(diary.Total.Protein)
	fat := asInt(diary.Total.Fat)
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

func diaryMessage(diary *Diary) string {
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

func asInt(macro string) int {
	m, err := strconv.Atoi(macro[:len(macro)-1])
	if err != nil {
		log.Fatal(err)
	}
	return m
}

/*
+-------------------+----------+
|       FOOD        | CALORIES |
+-------------------+----------+
| BREAKFAST         |          |
+-------------------+----------+
| Unsweetened       |       45 |
| Vanilla Almond    |          |
| Mil...            |          |
+-------------------+----------+
| Whey Protein      |      110 |
| Vanilla, 31 grams |          |
+-------------------+----------+
| Thick Sliced      |      120 |
| Bacon (80 Cal     |          |
| Per...            |          |
+-------------------+----------+
| 1 Large Egg, 3    |      210 |
| egg (50g)         |          |
+-------------------+----------+
| LUNCH             |          |
+-------------------+----------+
| Butter, 1 T.      |      100 |
+-------------------+----------+
| Chicken Breast,   |      660 |
| 1.5 lb(s)         |          |
+-------------------+----------+
| DINNER            |          |
+-------------------+----------+
| shredded sharp    |      167 |
| cheddar , 1.5     |          |
| o...              |          |
+-------------------+----------+
| Beef Smoked       |      280 |
| Sausage, 5 oz.    |          |
+-------------------+----------+
| Total             | 1,692    |
+-------------------+----------+
*/
