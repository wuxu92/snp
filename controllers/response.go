package controllers

import (

)

type ResObj struct {
	code int
	message string
	data interface{}
}

func (this *ResObj) Json() map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = this.code
	res["message"] = this.message
	res["data"] = this.data
	return res
}

const (
	CODE_OK	= 0
	CODE_ERR = 1
	CODE_ILLEGAL_PARAM = 2
	CODE_NULL_EXCEPTION = 3
	CODE_NO_SUCH_PKG = 4
	CODE_NO_SUCH_GRP = 5
	CODE_NO_SUCH_SITE = 6

	CODE_MGO_CONN_ERR = 21
	CODE_MGO_QUERY_ERR = 22
	CODE_MGO_NO_RESULT = 23
  CODE_MGO_BAD_ID = 24

  CODE_PASSWORD_ERR = 101

  CODE_CODE_ERR = 1000
)

var MESSAGES = map[int]string{
  CODE_OK: "OK",
  CODE_ERR: "Unknown error",
  CODE_ILLEGAL_PARAM: "Illegal input parameter(s)",
  CODE_NULL_EXCEPTION: "Null Exception, no such obj",
  CODE_NO_SUCH_PKG: "No such package",
  CODE_NO_SUCH_GRP: "Pkg/System Has No such Group",
  CODE_NO_SUCH_SITE: "No such Site",

  CODE_MGO_CONN_ERR: "mgo connect error",
  CODE_MGO_QUERY_ERR: "mgo query error",
  CODE_MGO_NO_RESULT: "no reuslt for this query",
  CODE_MGO_BAD_ID: "mongodb id is illegal",

  CODE_PASSWORD_ERR: "wrong password",
}

// set response code and message
// use this in controllers @see controllers/grp.go line 49
func (this *ResObj) SetCode(code int) *ResObj{
  if val, ok := MESSAGES[code]; ok {
    this.code = code
    this.message = val
  } else {
    this.code = CODE_CODE_ERR
    this.message = "err code not exist"
  }
  return this
}