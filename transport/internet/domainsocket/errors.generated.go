package domainsocket

import "v2ray.com/core/common/errors"

type errPathObjHolder struct{}

/*
@question 这是什么写法
*/
func newError(values ...interface{}) *errors.Error {
	return errors.New(values...).WithPathObj(errPathObjHolder{})
}
