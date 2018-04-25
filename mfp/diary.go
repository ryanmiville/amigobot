package mfp

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

//Fetcher is the interface for fetching a user's Diary
type Fetcher interface {
	//Fetch fetches and builds the Diary for the given username
	Fetch(username string) (*Diary, error)
}
