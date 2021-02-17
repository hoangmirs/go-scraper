package models

import (
	"database/sql/driver"
	"errors"

	"github.com/beego/beego/v2/client/orm"
)

type KeywordStatus string

const (
	KeywordType = "keyword"

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

type KeywordResult struct {
	Base

	KeyWord              string
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
	orm.RegisterModel(new(KeywordResult))
}

// CreateKeywordResult insert a new KeywordResult into database and returns last inserted Id on success.
func CreateKeywordResult(keywordResult *KeywordResult) (int64, error) {
	o := orm.NewOrm()

	return o.Insert(keywordResult)
}
