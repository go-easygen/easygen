package egVar_test

import (
	"fmt"
	"os"

	"github.com/go-easygen/easygen"
	"github.com/go-easygen/easygen/egVar"
)

func ExampleEgVar_output() {
	// EgVar variable names converting full coverage test

	tmpl0 := egVar.NewTemplate().Customize()
	tmpl := tmpl0.Funcs(easygen.FuncDefs()).Funcs(egVar.FuncDefs())
	fmt.Println("\n## From lk - KebabCase => CamelCase/SnakeCase")
	easygen.Process0(tmpl, os.Stdout,
		"{{.Name}}: {{clk2lc .Name}} {{clk2uc .Name}}\n", "../test/listfunc2")
	easygen.Process0(tmpl, os.Stdout,
		"{{.Name}}: {{clk2ls .Name}} {{clk2ss .Name}}\n", "../test/listfunc2")

	fmt.Println("\n## From ls/ss - LowerSnakeCase/ScreamingSnakeCase")
	easygen.Process0(tmpl, os.Stdout,
		"{{clk2ls .Name}} {{clk2ss .Name}} =>\n", "../test/listfunc2")
	fmt.Println("### From ls")
	easygen.Process0(tmpl, os.Stdout,
		" {{clk2ls .Name | cls2lc}} {{clk2ls .Name | cls2uc}}\n", "../test/listfunc2")
	easygen.Process0(tmpl, os.Stdout,
		" {{clk2ls .Name | cls2ss}} {{clk2ls .Name | cls2lk}} {{clk2ls .Name | cls2hh}}\n", "../test/listfunc2")
	fmt.Println("### From ss")
	easygen.Process0(tmpl, os.Stdout,
		" {{clk2ss .Name | css2lc}} {{clk2ss .Name | css2uc}}\n", "../test/listfunc2")
	easygen.Process0(tmpl, os.Stdout,
		" {{clk2ss .Name | css2ls}} {{clk2ss .Name | css2lk}} {{clk2ss .Name | css2hh}}\n", "../test/listfunc2")

	fmt.Println("\n## From lc/uc - LowerCamelCase/UpperCamelCaseKeepCaps")
	easygen.Process0(tmpl, os.Stdout,
		"{{clk2lc .Name}} {{clk2uc .Name}} =>\n", "../test/listfunc2")
	easygen.Process0(tmpl, os.Stdout,
		"{{clk2lc .Name}}: {{clk2lc .Name | clk2uc}} {{clk2lc .Name | clc2ls}} {{clk2lc .Name | clc2ss}}\n", "../test/listfunc2")
	easygen.Process0(tmpl, os.Stdout,
		"{{clk2uc .Name}}: {{clk2uc .Name | cuc2lc}} {{clk2uc .Name | cuc2ls}} {{clk2uc .Name | cuc2ss}}\n", "../test/listfunc2")

	// Output:
	//
	// ## From lk - KebabCase => CamelCase/SnakeCase
	// some-init-method: someInitMethod SomeInitMethod
	// some-init-method: some_init_method SOME_INIT_METHOD
	//
	// ## From ls/ss - LowerSnakeCase/ScreamingSnakeCase
	// some_init_method SOME_INIT_METHOD =>
	// ### From ls
	//  someInitMethod SomeInitMethod
	//  SOME_INIT_METHOD some-init-method Some-Init-Method
	// ### From ss
	//  someINITMETHOD SOMEINITMETHOD
	//  some_init_method some-init-method Some-Init-Method
	//
	// ## From lc/uc - LowerCamelCase/UpperCamelCaseKeepCaps
	// someInitMethod SomeInitMethod =>
	// someInitMethod: SomeInitMethod some_init_method SOME_INIT_METHOD
	// SomeInitMethod: someInitMethod some_init_method SOME_INIT_METHOD

}
