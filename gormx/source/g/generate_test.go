package g

import (
	"github.com/wordpress-plus/kit-common/kg"
	"testing"
)

func TestG(t *testing.T) {

	//FieldOpts = append(FieldOpts, gen.FieldType("delete_time", "gorm.DeletedAt"))
	g, db := G(kg.C.System.DbType, "UsMJS4zUTEE+9DDmZGY4u2diHz7GleHU0eIOOzN8mTS8hexx46hQBlctWvcBgm9jNOkl5+YEzh4d+9DaqMKFuc8DvSV6AgN9escAHlUN/KWU/KuBww44NHHoONOMox68uuMxK9gMH8n4El9COR4/hE9EIT52z9GwM22tV9VUj3I=", "./dal", "./relation.yaml")
	info := g.Data[db.NamingStrategy.SchemaName("archived_ups_tag")]
	g.ApplyInterface(func(upsTagInterface) {}, info.QueryStructMeta /*g.GenerateModel("archived_ups_tag")*/) /* relation will lose*/
	g.Execute()
}
