package domain

type ViewList struct {
    Username      string
    Dateandtime   string
    Place		  string
    ID            string  // id
	Title         string  // title
	OriginalTitle string  // original_title
	ReleaseDate   string    // release_date
	VoteAverage   float64 // vote_average
	PosterPath    string  // poster_path
	
	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ViewList exists in the database.
func (wl *ViewList) Exists() bool {
	return wl._exists
}

// Deleted provides information if the ViewList has been deleted from the database.
func (wl *ViewList) Deleted() bool {
	return wl._deleted
}

func ViewListByUsername(db XODB, username string) (*[]ViewList, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`username, dateandtime, place, id, title, original_title, release_date, vote_average, poster_path ` +
		`FROM public.viewlist ` +
		`WHERE username = $1`

	// run query
	XOLog(sqlstr, username)
	result := []ViewList{}

	rows, err := db.Query(sqlstr, username)
	if err != nil {
		return nil, err
	}
    for rows.Next() {
            wl := ViewList{
        		_exists: true,
        	}
        	err := rows.Scan(&wl.Username, &wl.Dateandtime, &wl.Place, &wl.ID, &wl.Title, &wl.OriginalTitle, &wl.ReleaseDate, &wl.VoteAverage, &wl.PosterPath)
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

func ViewListInLast30Days(db XODB) (*[]MovieMapping, error) {
	var err error
	
	// sql query
	const sqlstr = `SELECT ` +
		`id, title, original_title, release_date, vote_average, poster_path ` +
		`FROM public.movie_mapping ` +
		`WHERE id IN (SELECT ` +
		`id ` +
		`FROM public.viewlist ` +
		`WHERE viewlist.dateandtime > (CURRENT_DATE - INTERVAL '30 days'))`
	// run query
	XOLog(sqlstr)
	result := []MovieMapping{}

	rows, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
    for rows.Next() {
            wl := MovieMapping{
        		_exists: true,
        	}
        	err := rows.Scan(&wl.ID, &wl.Title, &wl.OriginalTitle, &wl.ReleaseDate, &wl.VoteAverage, &wl.PosterPath)
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

func ViewListOrderByCinema(db XODB) (*[]ViewList, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`username, dateandtime, place, id, title, original_title, release_date, vote_average, poster_path ` +
		`FROM public.viewlist ` +
		`ORDER BY place`

	// run query
	XOLog(sqlstr)
	result := []ViewList{}

	rows, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
    for rows.Next() {
            wl := ViewList{
        		_exists: true,
        	}
        	err := rows.Scan(&wl.Username, &wl.Dateandtime, &wl.Place, &wl.ID, &wl.Title, &wl.OriginalTitle, &wl.ReleaseDate, &wl.VoteAverage, &wl.PosterPath)
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