package wsinterface

type AddRequest struct {
    ID string
    Place string
}

type CollectionResponse struct {
    Status     string
	Message    string
}

type MyMovieShort struct {
    ID            string  
	Title         string  
	OriginalTitle string  
	ReleaseDate   string   
	VoteAverage   float64 
	PosterPath    string  
}