package copier

import (
	"github.com/gookit/goutil/jsonutil"
	"github.com/jinzhu/copier"
	"github.com/micro-services-roadmap/kit-common/util"
	"time"
)

var mapOfIntersConverts = []copier.TypeConverter{
	{
		SrcType: map[string]interface{}{},
		DstType: copier.String,
		Fn: func(src interface{}) (interface{}, error) {
			return jsonutil.EncodeString(src)
		},
	},
	{
		SrcType: copier.String,
		DstType: map[string]interface{}{},
		Fn: func(src interface{}) (interface{}, error) {
			str := src.(string)
			if len(str) == 0 {
				return map[string]interface{}{}, nil
			}

			var kmap = map[string]interface{}{}
			err := jsonutil.DecodeString(str, &kmap)
			return kmap, err
		},
	},
}

var mapOfStringConverts = []copier.TypeConverter{
	{
		SrcType: map[string]string{},
		DstType: copier.String,
		Fn: func(src interface{}) (interface{}, error) {
			return jsonutil.EncodeString(src)
		},
	},
	{
		SrcType: copier.String,
		DstType: map[string]string{},
		Fn: func(src interface{}) (interface{}, error) {
			str := src.(string)
			if len(str) == 0 {
				return map[string]string{}, nil
			}

			var kmap = map[string]string{}
			return kmap, jsonutil.DecodeString(str, &kmap)
		},
	},
}

var timeConverts = []copier.TypeConverter{
	{
		SrcType: time.Time{},
		DstType: copier.String,
		Fn: func(src interface{}) (interface{}, error) {
			return util.Format(src.(time.Time)), nil
		},
	},
	//{
	//	SrcType: &time.Time{},
	//	DstType: &str,
	//	Fn: func(src interface{}) (interface{}, error) {
	//		if src == nil {
	//			return nil, nil
	//		}
	//		vl := util.Format(src.(*time.Time))
	//		return &vl, nil
	//	},
	//},
	{
		SrcType: copier.String,
		DstType: time.Time{},
		Fn: func(src interface{}) (interface{}, error) {
			return util.TryParse(src.(string)), nil
		},
	},
	{
		SrcType: copier.String,
		DstType: &time.Time{},
		Fn: func(src interface{}) (interface{}, error) {
			time := util.TryParse(src.(string))
			if time.IsZero() {
				return nil, nil
			} else {
				return &time, nil
			}
		},
	},
}

//	{
//		SrcType: time.Time{},
//		DstType: copier.String,
//		Fn: func(src interface{}) (interface{}, error) {
//			s, ok := src.(time.Time)
//
//			if !ok {
//				return nil, errors.New("src type not matching")
//			}
//
//			return s.Format(time.RFC3339), nil
//		},
//	},

var str string = ""

func CopyWithTime(toValue interface{}, fromValue interface{}, ops ...*copier.Option) error {

	var usedOps *copier.Option
	if len(ops) > 0 {
		usedOps = ops[0]
	} else {
		usedOps = &copier.Option{}
	}

	converters := append(usedOps.Converters, mapOfIntersConverts...)
	converters = append(converters, mapOfStringConverts...)
	converters = append(converters, timeConverts...)
	usedOps.Converters = converters

	return copier.CopyWithOption(toValue, fromValue, *usedOps)
}
