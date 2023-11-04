package data

import (
	"time"

	up "github.com/upper/db/v4"
)

type User struct {
	ID        int       `db:"id,omitempty"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Active    int       `db:"user_active"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdateAt  time.Time `db:"updated_at"`
	Token     Token     `db:"-"`
}

func (u *User) Table() string {
	return "users"
}

func (u *User) GetAll() ([]*User, error) {
	collection := upper.Collection(u.Table())
	var all []*User

	res := collection.Find()
	err := res.All(&all)
	if err != nil {
		return nil, err
	}

	return all, nil
}

func (u *User) GetByEmail(email string) (*User, error) {
	collection := upper.Collection(u.Table())
	var theUser User

	res := collection.Find(up.Cond{"email =": email})
	err := res.One(&theUser)
	if err != nil {
		return nil, err
	}

	var token Token
	collection = upper.Collection(token.Table())
	res = collection.Find(up.Cond{"user_id =": theUser.ID, "expiry <": time.Now()}).OrderBy("created_at desc")
	err = res.One(&token)
	if err != nil {
		if err != up.ErrNilRecord && err != up.ErrNoMoreRows {
			return nil, err
		}
	}

	theUser.Token = token

	return &theUser, nil
}
