package g

import (
	"fmt"
	"github.com/alice52/jasypt-go"
	ggy "github.com/we7coreteam/gorm-gen-yaml"
	"github.com/wordpress-plus/kit-common/kg"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

var (
	autoUpdateTimeField = gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
		return map[string][]string{"column": {"update_time"}, "type": {"datetime(3)"}, "autoUpdateTime": {}}
	})
	autoCreateTimeField = gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
		return map[string][]string{"column": {"create_time"}, "type": {"datetime(3)"}, "autoCreateTime": {}}
	})
	softDeleteField = gen.FieldType("delete_time", "gorm.DeletedAt")
)
var FieldOpts = []gen.ModelOpt{autoCreateTimeField, autoUpdateTimeField, softDeleteField}

// https://zhuanlan.zhihu.com/p/653483236
// https://github.com/Alice52/go-tutorial/issues/5#issuecomment-1286325129

// Deprecated: use G2
func G(dbTpe, dsn, outputDir, relationYaml string) (*gen.Generator, *gorm.DB) {
	var DSN string
	if v, err := jasypt.New().Decrypt(dsn); err != nil {
		DSN = dsn
	} else {
		DSN = v
	}
	var dialector gorm.Dialector
	switch dbTpe {
	case kg.DbMysql:
		dialector = mysql.Open(DSN)
	case kg.DbPgsql:
		dialector = postgres.Open(DSN)
	default:
		panic("unknown db type")
	}

	return genCore(dialector, outputDir, relationYaml)
}

func G2(outputDir, relationYaml string) (*gen.Generator, *gorm.DB) {
	var dialector gorm.Dialector
	switch kg.C.System.DbType {
	case kg.DbMysql:
		dialector = mysql.Open(kg.C.Mysql.Dsn())
	case kg.DbPgsql:
		dialector = postgres.Open(kg.C.Pgsql.Dsn())
	default:
		panic("unknown db type")
	}

	return genCore(dialector, outputDir, relationYaml)
}

func genCore(dialector gorm.Dialector, outputDir string, relationYaml string) (*gen.Generator, *gorm.DB) {
	// 连接数据库
	db, err := gorm.Open(dialector)
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
	g.WithOpts(FieldOpts...)

	ggy.NewYamlGenerator(relationYaml).UseGormGenerator(g).Generate(FieldOpts...) // fieldOpts is not used
	//g.ApplyBasic(g.GenerateAllTable()...) // will lose relation, so donot use it after NewYamlGenerator
	//g.ApplyInterface(func(upsTagInterface) {}, g.GenerateModel("archived_ups_tag", FieldOpts...))

	return g, db
}
