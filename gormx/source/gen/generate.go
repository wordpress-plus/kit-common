package gen

import (
	"fmt"
	"github.com/alice52/jasypt-go"
	yamlgen "github.com/we7coreteam/gorm-gen-yaml"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

// https://zhuanlan.zhihu.com/p/653483236
// https://github.com/Alice52/go-tutorial/issues/5#issuecomment-1286325129

func G(dsn, outputDir, relationYaml string) {
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

	autoUpdateTimeField := gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
		return map[string][]string{"column": {"update_time"}, "type": {"datetime(3)"}, "autoUpdateTime": {}}
	})
	autoCreateTimeField := gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
		return map[string][]string{"column": {"create_time"}, "type": {"datetime(3)"}, "autoCreateTime": {}}
	})
	softDeleteField := gen.FieldType("delete_time", "gorm.DeletedAt")
	fieldOpts := []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField, softDeleteField}

	// allModel := g.GenerateAllTable(fieldOpts...)

	yamlgen.NewYamlGenerator(relationYaml).UseGormGenerator(g).Generate(fieldOpts...) // todo: bug fieldOpts is noi used
	// kg.ApplyInterface(func(UserInterface) {}, kg.GenerateModel("user"))
	g.Execute()
}
