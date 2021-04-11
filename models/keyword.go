package models

import (
	"database/sql/driver"
	"errors"

	"github.com/beego/beego/v2/client/orm"
)

// KeywordStatus is an implementation of a string for the SQL type
type KeywordStatus string

const (
	Failed     KeywordStatus = "failed"
	Pending    KeywordStatus = "pending"
	Processed  KeywordStatus = "processed"
	Processing KeywordStatus = "processing"

	InvalidKeywordStatusErr = "invalid keyword status"
)

// Value implements the driver.Valuer interface,
// and turns the KeywordStatus into a string for SQL storage
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
	NonAdwordLinks       string `orm:"type(json);null"`
	NonAdwordLinksCount  int
	AdwordLinks          string `orm:"type(json);null"`
	AdwordLinksCount     int
	ShopAdwordLinks      string `orm:"type(json);null"`
	ShopAdwordLinksCount int
	HtmlCode             string

	User *User `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Keyword))
}

// CreateKeyword inserts a new Keyword into database and returns last inserted Id on success.
func CreateKeyword(keyword *Keyword) (int64, error) {
	if len(keyword.Keyword) == 0 {
		return 0, errors.New("Keyword required")
	}

	if keyword.User == nil {
		return 0, errors.New("User required")
	}

	o := orm.NewOrm()
	keyword.Status = Pending

	return o.Insert(keyword)
}

// UpdateKeyword updates a Keyword and returns the error if any
func UpdateKeyword(keyword *Keyword) error {
	o := orm.NewOrm()

	_, err := o.Update(keyword)
	return err
}

// GetKeywordByID returns the Keyword by passing keywordID
func GetKeywordByID(keywordID int64) (*Keyword, error) {
	keyword := &Keyword{Base: Base{Id: uint(keywordID)}}

	o := orm.NewOrm()
	err := o.Read(keyword)

	return keyword, err
}

// GetKeywords returns keywords by query
func GetKeywords(query map[string]interface{}) ([]*Keyword, error) {
	keywords := []*Keyword{}

	_, err := keywordsQuerySeter(query).All(&keywords)

	return keywords, err
}

// GetKeywordsCount returns the number of current user's uploaded keywords
func GetKeywordsCount(query map[string]interface{}) (int64, error) {
	return keywordsQuerySeter(query).Count()
}

// GetKeyword gets a Keyword by query
func GetKeywordByQuery(query map[string]interface{}) (*Keyword, error) {
	keyword := &Keyword{}
	err := keywordsQuerySeter(query).RelatedSel().One(keyword)
	if err == orm.ErrNoRows {
		return nil, errors.New("Keyword not found")
	} else if err != nil {
		return nil, err
	}

	return keyword, nil
}

func keywordsQuerySeter(query map[string]interface{}) orm.QuerySeter {
	o := orm.NewOrm()
	querySetter := o.QueryTable(Keyword{})

	for key, value := range query {
		switch key {
		case "order":
			querySetter = querySetter.OrderBy(value.(string))
		case "limit":
			querySetter = querySetter.Limit(value)
		case "offset":
			querySetter = querySetter.Offset(value)
		default:
			querySetter = querySetter.Filter(key, value)
		}
	}

	return querySetter
}
