package mfp

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
)

type diary struct {
	meals map[string][]food
	total food
}

type food struct {
	name        string
	calories    string
	carbs       string
	fat         string
	protein     string
	cholesterol string
	sodium      string
	sugars      string
	fiber       string
}

func newDiary(username string) (*diary, error) {
	d := diary{}
	doc, err := goquery.NewDocument("http://www.myfitnesspal.com/reports/printable_diary/" + username)
	if err != nil {
		return &d, err
	}
	d.meals = buildMeals(doc)
	d.total, err = buildTotal(doc)
	return &d, err
}

func buildMeals(doc *goquery.Document) map[string][]food {
	var meal = "Breakfast"
	var meals = make(map[string][]food)
	doc.Find("tbody").Find("tr").Each(func(i int, tr *goquery.Selection) {
		entry := entryData(tr)
		if len(entry) == 1 {
			meal = entry[0]
			return
		}
		food := newFood(entry)
		meals[meal] = append(meals[meal], food)
	})
	return meals
}

func buildTotal(doc *goquery.Document) (food, error) {
	totalRow := doc.Find("tfoot").Find("tr")
	entry := entryData(totalRow)
	if len(entry) == 0 {
		return food{}, errors.New("Gotta log, friendo")
	}
	return newFood(entry), nil
}

func entryData(entry *goquery.Selection) []string {
	var data []string
	entry.Find("td").Each(func(j int, td *goquery.Selection) {
		data = append(data, td.Text())
	})
	return data
}

func newFood(data []string) food {
	return food{
		name:        data[0],
		calories:    data[1],
		carbs:       data[2],
		fat:         data[3],
		protein:     data[4],
		cholesterol: data[5],
		sodium:      data[6],
		sugars:      data[7],
		fiber:       data[8],
	}
}
