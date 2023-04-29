// gorm-gen 根据数据库表结构自动生成 model 和对应的 crud 方法
package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
	"strings"
)

// 自定义数据类型映射(也就是 mysql 中的数据类型与 Go 语言中的数据类型的映射)
/*var dataMap = map[string]func(detailType string) (dataType string){
	// int 数据类型映射
	"int": func(detailType string) (dataType string) { return "int32" },

	// bool 数据类型映射
	"tinyint": func(detailType string) (dataType string) {
		if strings.HasPrefix(detailType, "tinyint(1)") {
			return "bool"
		}
		return "byte"
	},

	// decimal 带小数点的金额数据类型映射,
	"decimal": func(detailType string) (dataType string) {
		return "decimal.Decimal"
	},
}
*/

var dataMap = map[string]func(columnType gorm.ColumnType) (dataType string){
	"int": func(columnType gorm.ColumnType) (dataType string) { return "int32" },
	"tinyint": func(columnType gorm.ColumnType) (dataType string) {
		if strings.HasPrefix(columnType.Name(), "tinyint(1)") {
			return "bool"
		}
		return "byte"
	},
	"decimal": func(columnType gorm.ColumnType) (dataType string) {
		return "decimal.Decimal"
	},
}

func main() {
	fmt.Println(os.Getwd())
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./gen/query",  // crud 方法生成路径
		ModelPkgPath: "./gen/model", // 数据库模型生成路径
		Mode:         gen.WithoutContext | gen.WithDefaultQuery,
		// if you want the nullable field generation property to be pointer type, set FieldNullable true
		FieldNullable: true,
		// if you want to assign field which has default value in `Create` API, set FieldCoverable true, reference: https://gorm.io/docs/create.html#Default-Values
		FieldCoverable: true,
		// if you want to generate field with unsigned integer type, set FieldSignable true
		/* FieldSignable: true,*/
		// if you want to generate index tags from database, set FieldWithIndexTag true
		FieldWithIndexTag: true,
		// if you want to generate type tags from database, set FieldWithTypeTag true
		FieldWithTypeTag: true, // 生成结构体对象的时候, 结构体标签包含类型信息
		// if you need unit tests for query code, set WithUnitTest true
		/* WithUnitTest: true, */
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	db, _ := gorm.Open(mysql.Open("root:wang0909@(127.0.0.1:3306)/blog_project?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(db)

	// 使用自定义的数据映射
	g.WithDataTypeMap(dataMap)

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.
	// 想对已有的model生成crud等基础方法可以直接指定model struct ，例如model.User{}
	// 如果是想直接生成表的model和crud方法，则可以指定表的名称，例如g.GenerateModel("company")
	// 想自定义某个表生成特性，比如struct的名称/字段类型/tag等，可以指定opt，例如g.GenerateModel("company",gen.FieldIgnore("address")), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address"))
	//g.ApplyBasic(model.User{}, g.GenerateModel("company"), g.GenerateModelAs("people", "Person", gen.FieldIgnore("address")))
	//g.ApplyBasic(model.User{})
	g.ApplyBasic(g.GenerateAllTable()...) // 生成所有表的 crud 方法
	// apply diy interfaces on structs or table models
	// 如果想给某些表或者model生成自定义方法，可以用ApplyInterface，第一个参数是方法接口，可以参考DIY部分文档定义
	//g.ApplyInterface(func(method model.Method) {}, model.User{}, g.GenerateModel("company"))

	g.GenerateAllTable() // 生成所有表的模型

	// execute the action of code generation
	g.Execute()
}
