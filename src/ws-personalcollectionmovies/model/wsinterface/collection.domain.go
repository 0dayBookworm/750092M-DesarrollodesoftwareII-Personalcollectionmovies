package wsinterface

type AddRequest struct {
    ID string
}

type CollectionResponse struct {
    Status     string
	Message    string
}

type MyMovieShort struct {
    Username      string
    ID            string  
	Title         string  
	OriginalTitle string  
	ReleaseDate   string   
	VoteAverage   float64 
	PosterPath    string  
}