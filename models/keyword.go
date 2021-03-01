package models

import (
	"database/sql/driver"
	"errors"

	"github.com/beego/beego/v2/client/orm"
)

type KeywordStatus string

const (
	Failed     KeywordStatus = "failed"
	Pending    KeywordStatus = "pending"
	Processed  KeywordStatus = "processed"
	Processing KeywordStatus = "processing"

	InvalidKeywordStatusErr = "invalid keyword status"
)

func (k KeywordStatus) Value() (driver.Value, error) {
	switch k {
	case Pending, Processing, Processed, Failed:
		return string(k), nil
	}
	return nil, errors.New(InvalidKeywordStatusErr)
}

type Keyword struct {
	Base

	Keyword string
	Status  KeywordStatus `orm:"type(KeywordStatus);default(pending)"`

	LinksCount           int
	NonAdwordLinks       string `orm:"type(json)"`
	NonAdwordLinksCount  int
	AdwordLinks          string `orm:"type(json)"`
	AdwordLinksCount     int
	ShopAdwordLinks      string `orm:"type(json)"`
	ShopAdwordLinksCount int
	HtmlCode             string

	User *User `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Keyword))
}

// CreateKeyword insert a new Keyword into database and returns last inserted Id on success.
func CreateKeyword(keyword *Keyword) (int64, error) {
	o := orm.NewOrm()
	keyword.Status = Pending

	return o.Insert(keyword)
}

// UpdateKeyword update a Keyword and returns the error if any
func UpdateKeyword(keyword *Keyword) error {
	o := orm.NewOrm()

	_, err := o.Update(keyword)
	return err
}
