package copier

import (
	"github.com/jinzhu/copier"
	"github.com/micro-services-roadmap/kit-common/util"
	"time"
)

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

func CopyWithTime(toValue interface{}, fromValue interface{}) error {

	option := copier.Option{
		Converters: []copier.TypeConverter{
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
		},
	}

	return copier.CopyWithOption(toValue, fromValue, option)
}
