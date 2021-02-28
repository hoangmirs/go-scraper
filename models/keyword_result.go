package models

import (
	"github.com/beego/beego/v2/client/orm"
)

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
