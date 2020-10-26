package dat

import (
	"database/sql"
	"encoding/base32"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/nadavbm/nogobk/pkg/logger"
	"github.com/pkg/errors"
)

type DbCookieStore struct {
	Codecs  []securecookie.Codec
	Options *sessions.Options
	Path    string
	DbPool  *sql.DB
}

func NewCookieStore(l *logger.Logger, keyPairs ...[]byte) (*DbCookieStore, error) {
	conn := GetDBConnString()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		l.Info("could not open connection to database")
		return nil, err
	}

	return NewCookieStoreFromPool(db, keyPairs...)
}

func NewCookieStoreFromPool(db *sql.DB, keyPairs ...[]byte) (*DbCookieStore, error) {
	dbStore := &DbCookieStore{
		Codecs: securecookie.CodecsFromPairs(keyPairs...),
		Options: &sessions.Options{
			Path:   "/",
			MaxAge: 86400 * 30,
		},
		DbPool: db,
	}

	return dbStore, nil
}

// close db connection to cookie store
func (cs *DbCookieStore) Close() {
	cs.DbPool.Close()
}

// Get Fetches a session for a given name after it has been added to the registry
func (cs *DbCookieStore) Get(r *http.Request, name string) (*sessions.Session, error) {
	return sessions.GetRegistry(r).Get(cs, name)
}

// New returns a new session for the given name without adding it to the registry.
func (cs *DbCookieStore) New(r *http.Request, name string) (*sessions.Session, error) {
	session := sessions.NewSession(cs, name)
	if session == nil {
		return nil, nil
	}

	opts := *cs.Options
	session.Options = &(opts)
	session.IsNew = true

	var err error
	if c, errCookie := r.Cookie(name); errCookie == nil {
		err = securecookie.DecodeMulti(name, c.Value, &session.ID, cs.Codecs...)
		if err == nil {
			err = cs.load(session)
			if err == nil {
				session.IsNew = false
			} else if errors.Cause(err) == sql.ErrNoRows {
				err = nil
			}
		}
	}

	cs.MaxAge(cs.Options.MaxAge)

	return session, err
}

// Save saves the given session into the database and deletes cookies if needed
func (cs *DbCookieStore) Save(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	// Set delete if max-age is < 0
	if session.Options.MaxAge < 0 {
		if err := cs.destroy(session); err != nil {
			return err
		}
		http.SetCookie(w, sessions.NewCookie(session.Name(), "", session.Options))
		return nil
	}

	if session.ID == "" {
		// Generate a random session ID key suitable for storage in the DB
		session.ID = strings.TrimRight(
			base32.StdEncoding.EncodeToString(
				securecookie.GenerateRandomKey(32),
			), "=")
	}

	if err := cs.save(session); err != nil {
		return err
	}

	// Keep the session ID key in a cookie so it can be looked up in DB later.
	encoded, err := securecookie.EncodeMulti(session.Name(), session.ID, cs.Codecs...)
	if err != nil {
		return err
	}

	http.SetCookie(w, sessions.NewCookie(session.Name(), encoded, session.Options))
	return nil
}

func (cs *DbCookieStore) MaxLength(l int) {
	for _, c := range cs.Codecs {
		if codec, ok := c.(*securecookie.SecureCookie); ok {
			codec.MaxLength(l)
		}
	}
}

func (cs *DbCookieStore) MaxAge(age int) {
	cs.Options.MaxAge = age

	// Set the maxAge for each securecookie instance.
	for _, codec := range cs.Codecs {
		if sc, ok := codec.(*securecookie.SecureCookie); ok {
			sc.MaxAge(age)
		}
	}
}

func (cs *DbCookieStore) load(session *sessions.Session) error {
	var s Session

	err := cs.selectOne(&s, session.ID)
	if err != nil {
		return err
	}

	return securecookie.DecodeMulti(session.Name(), string(s.Csrf), &session.Values, cs.Codecs...)
}

// save writes encoded session.Values to a database record.
// writes to http_sessions table by default.
func (cs *DbCookieStore) save(session *sessions.Session) error {
	encoded, err := securecookie.EncodeMulti(session.Name(), session.Values, cs.Codecs...)
	if err != nil {
		return err
	}

	crOn := session.Values["created"]
	exOn := session.Values["expires"]

	var expiresOn time.Time

	createdOn, ok := crOn.(time.Time)
	if !ok {
		createdOn = time.Now()
	}

	if exOn == nil {
		expiresOn = time.Now().Add(time.Second * time.Duration(session.Options.MaxAge))
	} else {
		expiresOn = exOn.(time.Time)
		if expiresOn.Sub(time.Now().Add(time.Second*time.Duration(session.Options.MaxAge))) < 0 {
			expiresOn = time.Now().Add(time.Second * time.Duration(session.Options.MaxAge))
		}
	}

	s := Session{
		Token:   session.ID,
		Csrf:    encoded,
		Created: createdOn,
		Expires: expiresOn,
	}

	if session.IsNew {
		return cs.insert(&s)
	}

	return cs.update(&s)
}

// Delete session
func (cs *DbCookieStore) destroy(session *sessions.Session) error {
	_, err := cs.DbPool.Exec("DELETE FROM sessions WHERE key = $1", session.ID)
	return err
}

func (cs *DbCookieStore) selectOne(s *Session, key string) error {
	stmt := "SELECT userid, token, csrf, created, expires FROM sessions WHERE key = $1"
	err := cs.DbPool.QueryRow(stmt, key).Scan(&s.UserId, &s.Token, &s.Csrf, &s.Created, &s.Expires)
	if err != nil {
		return errors.Wrapf(err, "Unable to find session in the database")
	}

	return nil
}

func (cs *DbCookieStore) insert(s *Session) error {
	stmt := `INSERT INTO sessions (userid, token, csrf, created, expires)
           VALUES ($1, $2, $3, $4, $5)`
	_, err := cs.DbPool.Exec(stmt, s.UserId, s.Token, s.Csrf, s.Created, s.Expires)

	return err
}

func (cs *DbCookieStore) update(s *Session) error {
	stmt := `UPDATE sessions SET csrf=$1, expires=$2 WHERE token=$3`
	_, err := cs.DbPool.Exec(stmt, s.Csrf, s.Expires, s.Token)

	return err
}
