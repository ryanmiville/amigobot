package htmlparse

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
	"github.com/ryanmiville/amigobot/mfp"
)

//Fetcher fetches an mfp Diary by parsing HTML
type Fetcher struct{}

//Fetch parses the HTML from a public mfp account's diary page
func (f Fetcher) Fetch(username string) (*mfp.Diary, error) {
	d := mfp.Diary{}
	doc, err := goquery.NewDocument("http://www.myfitnesspal.com/reports/printable_diary/" + username)
	if err != nil {
		return &d, err
	}
	d.Meals = buildMeals(doc)
	d.Total, err = buildTotal(doc)
	return &d, err
}

func buildMeals(doc *goquery.Document) map[string][]mfp.Food {
	var meal = "Breakfast"
	var meals = make(map[string][]mfp.Food)
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

func buildTotal(doc *goquery.Document) (mfp.Food, error) {
	totalRow := doc.Find("tfoot").Find("tr")
	entry := entryData(totalRow)
	if len(entry) == 0 {
		return mfp.Food{}, errors.New("Gotta log, friendo")
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

func newFood(data []string) mfp.Food {
	return mfp.Food{
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
