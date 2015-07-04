
# easygen – Easy to use universal code/text generator

## Install

	go get github.com/suntong001/easygen
	ls -l $GOPATH/bin

## Test

	export PATH=$PATH:$GOPATH/bin

	$ easygen $GOPATH/src/github.com/suntong001/easygen/Test/list0
	The colors are: red, blue, white, .

	cd $GOPATH/src/github.com/suntong001/easygen

	$ easygen Test/list1 
	The quoted colors are: "red", "blue", "white", .

	$ easygen -tf test/listfunc1 test/list0
	red, blue, white.

```
$ go test -v 
=== RUN TestList0
--- PASS: TestList0 (0.00s)
        easygen_test.go:17: First and plainest list test
=== RUN TestList1Text
--- PASS: TestList1Text (0.00s)
        easygen_test.go:28: Second test, with text template
=== RUN TestList1HTML
--- PASS: TestList1HTML (0.00s)
        easygen_test.go:36: Second test, with html template
=== RUN TestListFunc1
--- PASS: TestListFunc1 (0.00s)
        easygen_test.go:47: Test custom template function - minus1
=== RUN: ExampleFunc1
--- PASS: ExampleFunc1 (0.00s)
=== RUN: ExampleList0Func1
--- PASS: ExampleList0Func1 (0.00s)
=== RUN: ExampleList0StrTemplate
--- PASS: ExampleList0StrTemplate (0.00s)
=== RUN: ExampleFunc2
--- PASS: ExampleFunc2 (0.00s)
=== RUN: ExampleTestExample
--- PASS: ExampleTestExample (0.00s)
=== RUN: ExampleCommandLineCobraViper
--- PASS: ExampleCommandLineCobraViper (0.00s)
=== RUN: ExampleCommandLineOptInitFull
--- PASS: ExampleCommandLineOptInitFull (0.00s)
PASS
ok      github.com/suntong001/easygen   0.014s
```

## Help

```
$ easygen

Usage:
 easygen [flags] YamlFileName

  -debug=0: debugging level
  -html=false: treat the template file as html instead of text
  -tf="": .tmpl template file name (default: same as .yaml file)
  -ts="": template string (in text)

YamlFileName: The name for the .yaml data and .tmpl template file
        Only the name part, without extension. Can include the path as well.

```

## Details

My (updated) blog about it is at [here](https://github.com/suntong001/blog/blob/master/GoOptP7-easygen.md), and [here](https://sfxpt.wordpress.com/2015/07/04/easygen-is-now-coding-itself/).

<a name="tips" />
## Tips

You can use `easygen` as an generic Google template testing tool with the `-ts` commandline option. For example,

```
echo "Age: 16" > /tmp/age.yaml

$ easygen -ts "{{.Age}}" /tmp/age
16

$ easygen -ts '{{printf "%x" .Age}}' /tmp/age
10

echo '{FirstName: John, LastName: Doe}' > /tmp/name.yaml

$ easygen -ts '{{.FirstName}}'\''s full name is {{printf "%s%s" .FirstName .LastName | len}} letters long.' /tmp/name
John's full name is 7 letters long.

$ easygen -ts "{{.FirstName}} {{ck2ss .LastName}}'s full name is "'{{len (printf "%s%s" .FirstName .LastName)}} letters long.' /tmp/name
John DOE's full name is 7 letters long.

echo 'Name: some-init-method' > /tmp/var.yaml

$ easygen -ts '{{.Name}} {{6 | minus1}} {{minus1 6}} {{ck2lc .Name}} {{ck2uc .Name}}' /tmp/var
some-init-method 5 5 someInitMethod SomeInitMethod

```

## Author(s)

Tong SUN, suntong at/from cpan.org.

All patches welcome. 
