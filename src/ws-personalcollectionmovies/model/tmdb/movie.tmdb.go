package tmdb

import (
	tmdb "github.com/ryanbradynd05/go-tmdb"
)


var (
	ApiKey = ""
	TMDb = &tmdb.TMDb{}
	language = "es"
	)

func Init(pApiKey, pLanguage string) {
	ApiKey = pApiKey
    TMDb = tmdb.Init(ApiKey)
    language = pLanguage
}

func SearchMovie(pTitle, pYear string) (*tmdb.MovieSearchResults, error){
	var options = make(map[string]string)
	options["language"] = language
	options["year"] = pYear
	return TMDb.SearchMovie(pTitle, options)
}

func GetMovieInfo(pID int) (*tmdb.Movie, error){
	var options = make(map[string]string)
	options["language"] = language
	options["append_to_response"] = "videos"
	return TMDb.GetMovieInfo(pID, options)
}

func GetUpcomingMovies(pPage string) (*tmdb.MovieDatedResults, error) {
	var options = make(map[string]string)
	options["page"]=pPage
	options["language"]=language
	return TMDb.GetMovieUpcoming(options)
}

func GetMovieNowPlaying(pPage string) (*tmdb.MovieDatedResults, error) {
	var options = make(map[string]string)
	options["page"] = pPage
	options["language"]=language
	return TMDb.GetMovieNowPlaying(options)
}