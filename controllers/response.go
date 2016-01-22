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
)