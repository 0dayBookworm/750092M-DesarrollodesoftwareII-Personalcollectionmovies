package domain

type WatchList struct {
    Username      string
    ID            string  // id
	Title         string  // title
	OriginalTitle string  // original_title
	ReleaseDate   string    // release_date
	VoteAverage   float64 // vote_average
	PosterPath    string  // poster_path
	
	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WatchList exists in the database.
func (wl *WatchList) Exists() bool {
	return wl._exists
}

// Deleted provides information if the WatchList has been deleted from the database.
func (wl *WatchList) Deleted() bool {
	return wl._deleted
}

func WatchListByUsername(db XODB, username string) (*[]WatchList, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`username, id, title, original_title, release_date, vote_average, poster_path ` +
		`FROM public.watchlist ` +
		`WHERE username = $1`

	// run query
	XOLog(sqlstr, username)
	result := []WatchList{}

	rows, err := db.Query(sqlstr, username)
	if err != nil {
		return nil, err
	}
    for rows.Next() {
            wl := WatchList{
        		_exists: true,
        	}
        	err := rows.Scan(&wl.Username, &wl.ID, &wl.Title, &wl.OriginalTitle, &wl.ReleaseDate, &wl.VoteAverage, &wl.PosterPath)
            if err != nil {
                return nil, err
            }
            result = append(result, wl)
    }
    err = rows.Err()
    if err != nil {
        return nil, err
    }
	return &result, nil
}
