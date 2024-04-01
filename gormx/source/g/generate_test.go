package g

import (
	"testing"
)

func TestG(t *testing.T) {

	g := G("UsMJS4zUTEE+9DDmZGY4u2diHz7GleHU0eIOOzN8mTS8hexx46hQBlctWvcBgm9jNOkl5+YEzh4d+9DaqMKFuc8DvSV6AgN9escAHlUN/KWU/KuBww44NHHoONOMox68uuMxK9gMH8n4El9COR4/hE9EIT52z9GwM22tV9VUj3I=", "./dal", "./relation.yaml")
	g.ApplyInterface(func(upsTagInterface) {}, g.GenerateModel("archived_ups_tag", FieldOpts...))
	g.Execute()
}
