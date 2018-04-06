package mfp

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
)

//Diary is the day's food diary for the user
type Diary struct {
	Meals map[string][]Food
	Total Food
}

//Food is a name and its nutrition
type Food struct {
	Name        string
	Calories    string
	Carbs       string
	Fat         string
	Protein     string
	Cholesterol string
	Sodium      string
	Sugars      string
	Fiber       string
}

//NewDiary fetches and builds the Diary for the given username
func NewDiary(username string) (*Diary, error) {
	d := Diary{}
	doc, err := goquery.NewDocument("http://www.myfitnesspal.com/reports/printable_diary/" + username)
	if err != nil {
		return &d, err
	}
	d.Meals = buildMeals(doc)
	d.Total, err = buildTotal(doc)
	return &d, err
}

func buildMeals(doc *goquery.Document) map[string][]Food {
	var meal = "Breakfast"
	var meals = make(map[string][]Food)
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

func buildTotal(doc *goquery.Document) (Food, error) {
	totalRow := doc.Find("tfoot").Find("tr")
	entry := entryData(totalRow)
	if len(entry) == 0 {
		return Food{}, errors.New("Gotta log, friendo")
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

func newFood(data []string) Food {
	return Food{
		Name:        data[0],
		Calories:    data[1],
		Carbs:       data[2],
		Fat:         data[3],
		Protein:     data[4],
		Cholesterol: data[5],
		Sodium:      data[6],
		Sugars:      data[7],
		Fiber:       data[8],
	}
}
