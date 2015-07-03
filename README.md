
# EasyGen – Universal code/text generator that is easy to use

## Install

	go get github.com/suntong001/EasyGen
	ls -l $GOPATH/bin

## Test

	export PATH=$PATH:$GOPATH/bin

	$ EasyGen $GOPATH/src/github.com/suntong001/EasyGen/Test/list0
	The colors are: red, blue, white, .

	cd $GOPATH/src/github.com/suntong001/EasyGen

	$ EasyGen Test/list1 
	The quoted colors are: "red", "blue", "white", .

```
$ go test -v 
=== RUN TestList0
--- PASS: TestList0 (0.00s)
        EasyGen_test.go:15: First and plainest list test
=== RUN TestList1Text
--- PASS: TestList1Text (0.00s)
        EasyGen_test.go:26: Second test, with text template
=== RUN TestList1HTML
--- PASS: TestList1HTML (0.00s)
        EasyGen_test.go:34: Second test, with html template
=== RUN TestListFunc1
--- PASS: TestListFunc1 (0.00s)
        EasyGen_test.go:45: Test custom template function - minus1
=== RUN: ExampleFunc1
--- PASS: ExampleFunc1 (0.00s)
=== RUN: ExampleTestExample
--- PASS: ExampleTestExample (0.00s)
=== RUN: ExampleCommandLineCobraViper
--- PASS: ExampleCommandLineCobraViper (0.00s)
=== RUN: ExampleCommandLineOptInitFull
--- PASS: ExampleCommandLineOptInitFull (0.00s)
PASS
ok      github.com/suntong001/EasyGen   0.011s
```

## Help

```
$ EasyGen

Usage:
 EasyGen [flags] YamlFileName

Flags:

  -html=false: treat the template file as html instead of text
  -tf="": .tmpl template file name (default: same as .yaml file)
  -ts="": template string (in text)

YamlFileName: The name for the .yaml data and .tmpl template file
        Only the name part, without extension. Can include the path as well.
```

## Details

My (updated) blog about it is at [here](https://github.com/suntong001/blog/blob/master/GoOptP7-EasyGen.md).

## Tips

You can use `EasyGen` as an generic Google template testing tool with the `-ts` commandline option. For example,

```
echo "Age: 16" > /tmp/age.yaml

$ EasyGen -ts "{{.Age}}" /tmp/age
16

$ EasyGen -ts '{{printf "%x" .Age}}' /tmp/age
10

echo '{FirstName: John, LastName: Doe}' > /tmp/name.yaml

$ EasyGen -ts '{{.FirstName}}'\''s full name is {{printf "%s%s" .FirstName .LastName | len}} letters long.' /tmp/name
John's full name is 7 letters long.

$ EasyGen -ts '{{.FirstName}}'\''s full name is {{len (printf "%s%s" .FirstName .LastName)}} letters long.' /tmp/name
John's full name is 7 letters long.

```
