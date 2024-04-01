package g

import (
	"fmt"
	"github.com/alice52/jasypt-go"
	ggy "github.com/we7coreteam/gorm-gen-yaml"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

var autoUpdateTimeField = gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
	return map[string][]string{"column": {"update_time"}, "type": {"datetime(3)"}, "autoUpdateTime": {}}
})
var autoCreateTimeField = gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
	return map[string][]string{"column": {"create_time"}, "type": {"datetime(3)"}, "autoCreateTime": {}}
})
var softDeleteField = gen.FieldType("delete_time", "gorm.DeletedAt")

var FieldOpts = []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField, softDeleteField}

// https://zhuanlan.zhihu.com/p/653483236
// https://github.com/Alice52/go-tutorial/issues/5#issuecomment-1286325129

func G(dsn, outputDir, relationYaml string) *gen.Generator {
	var MySQLDSN string
	if v, err := jasypt.New().Decrypt(dsn); err != nil {
		MySQLDSN = dsn
	} else {
		MySQLDSN = v
	}

	// 连接数据库
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           outputDir, //"./source/gen/dal",
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldSignable:     false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  true,
	})

	g.UseDB(db)

	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int64" },
	}
	g.WithDataTypeMap(dataMap)

	ggy.NewYamlGenerator(relationYaml).UseGormGenerator(g).Generate(FieldOpts...) // fieldOpts is not used
	g.ApplyBasic(g.GenerateAllTable(FieldOpts...)...)
	//g.ApplyInterface(func(upsTagInterface) {}, g.GenerateModel("archived_ups_tag", FieldOpts...))

	return g
}
