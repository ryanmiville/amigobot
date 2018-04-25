package mock

import "github.com/ryanmiville/amigobot/mfp"

//Fetcher mocks an mfp.Fetcher
type Fetcher struct {
	FetchFn func(username string) (*mfp.Diary, error)
}

//Fetch calls the mocked implementation
func (f *Fetcher) Fetch(username string) (*mfp.Diary, error) {
	return f.FetchFn(username)
}
