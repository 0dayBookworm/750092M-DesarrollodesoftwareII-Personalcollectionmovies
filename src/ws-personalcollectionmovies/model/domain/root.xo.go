// Package domain contains the types for schema 'public'.
package domain

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// Root represents a row from public.root.
type Root struct {
	Username string // username
	Pass     string // pass

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Root exists in the database.
func (r *Root) Exists() bool {
	return r._exists
}

// Deleted provides information if the Root has been deleted from the database.
func (r *Root) Deleted() bool {
	return r._deleted
}

// Insert inserts the Root to the database.
func (r *Root) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.root (` +
		`username, pass` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING username`

	// run query
	XOLog(sqlstr, r.Username, r.Pass)
	err = db.QueryRow(sqlstr, r.Username, r.Pass).Scan(&r.Username)
	if err != nil {
		return err
	}

	// set existence
	r._exists = true

	return nil
}

// Update updates the Root in the database.
func (r *Root) Update(db XODB) error {
	var err error

	// if deleted, bail
	if r._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.root SET (` +
		`pass` +
		`) = ( ` +
		`$1` +
		`) WHERE username = $2`

	// run query
	XOLog(sqlstr, r.Pass, r.Username)
	_, err = db.Exec(sqlstr, r.Pass, r.Username)
	return err
}

// Save saves the Root to the database.
func (r *Root) Save(db XODB) error {
	if r.Exists() {
		return r.Update(db)
	}

	return r.Insert(db)
}

// Upsert performs an upsert for Root.
//
// NOTE: PostgreSQL 9.5+ only
func (r *Root) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.root (` +
		`username, pass` +
		`) VALUES (` +
		`$1, $2` +
		`) ON CONFLICT (username) DO UPDATE SET (` +
		`username, pass` +
		`) = (` +
		`EXCLUDED.username, EXCLUDED.pass` +
		`)`

	// run query
	XOLog(sqlstr, r.Username, r.Pass)
	_, err = db.Exec(sqlstr, r.Username, r.Pass)
	if err != nil {
		return err
	}

	// set existence
	r._exists = true

	return nil
}

// Delete deletes the Root from the database.
func (r *Root) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return nil
	}

	// if deleted, bail
	if r._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.root WHERE username = $1`

	// run query
	XOLog(sqlstr, r.Username)
	_, err = db.Exec(sqlstr, r.Username)
	if err != nil {
		return err
	}

	// set deleted
	r._deleted = true

	return nil
}

// RootByUsername retrieves a row from 'public.root' as a Root.
//
// Generated from index 'pk_root'.
func RootByUsername(db XODB, username string) (*Root, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`username, pass ` +
		`FROM public.root ` +
		`WHERE username = $1`

	// run query
	XOLog(sqlstr, username)
	r := Root{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, username).Scan(&r.Username, &r.Pass)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
