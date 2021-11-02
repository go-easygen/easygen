
## Testing the templates on the fly


    echo 'Name: some-init-method' > /tmp/var.yaml

    $ EASYGEN_TS='{{.Name}}' easygen /tmp/var
    some-init-method

I.e., with the environment variable `EASYGEN_TS`, the `.tmpl` template file is not used.

	$ EASYGEN_TS='{{.Name}}' easygen -ts '{{clk2uc .Name}}' /tmp/var
	SomeInitMethod
	
	$ easygen -ts '{{substr (clk2uc .Name) 4}}' /tmp/var
	InitMethod
	
	$ easygen -ts '{{substr (clk2uc .Name) 8 3}}' /tmp/var
	Met
	
	$ easygen -ts '{{substr (clk2uc .Name) 4 -2}}' /tmp/var
	InitMeth

	$ easygen -ts '{{substr (clk2uc .Name) -6}}' /tmp/var
	Method

	$ easygen -ts '{{substr "abcde" -3}}' /tmp/var
	cde

I.e., command line value takes the highest priority, even overriding the environment variable `EASYGEN_TS`'s value.

As such, if you have a different naming convention than using `.tmpl` for template file and `.yaml` for data file, you can  override them in environment variables, `EASYGEN_ET` and `EASYGEN_EY`, so that you don't need to use `-et` and/or `-ey` to override them from command line each time. 

	echo 'Name: myConstantVariable' > /tmp/var.yml

	$ EASYGEN_EY=.yml easygen -ts '{{clc2ss .Name}}' /tmp/var
	MY_CONSTANT_VARIABLE


<a name="tips" />

## Tips

[ ](https://suntong.github.io/blogs/)

You can use `easygen` as an generic Go template testing tool with the `-ts` commandline option. For example,

```bash
echo "Age: 16" > /tmp/age.yaml

$ easygen -ts "{{.Age}}" /tmp/age
16

$ easygen -ts '{{printf "%x" .Age}}' /tmp/age
10

$ easygen -ts '{{date "Y4"}}' /tmp/age
2017
$ 

$ easygen -ts '{{date "I"}}' /tmp/age
2017-07-10

$ easygen -ts '{{date "."}}' /tmp/age
2017.07.10

$ easygen -ts '{{date ""}}' /tmp/age
20170710

$ easygen -ts '{{timestamp}}' /tmp/age
2017-07-10T08:53:25-04:00

$ easygen -ts '{{timestamp "unix"}}' /tmp/age
1499691221

echo '{FirstName: John, LastName: Doe}' > /tmp/name.yaml

$ easygen -ts '{{.FirstName}}'\''s full name is {{printf "%s%s" .FirstName .LastName | len}} letters long.' /tmp/name
John's full name is 7 letters long.

$ easygen -ts "{{.FirstName}} {{clk2ss .LastName}}'s full name is "'{{len (printf "%s%s" .FirstName .LastName)}} letters long.' /tmp/name
John DOE's full name is 7 letters long.

echo 'Name: some-init-method' > /tmp/var.yaml

$ easygen -ts '{{.Name}} {{6 | minus1}} {{minus1 6}} {{clk2lc .Name}} {{clk2uc .Name}}' /tmp/var
some-init-method 5 5 someInitMethod SomeInitMethod

# More built-in function examples. See https://github.com/go-easygen/easygen/issues/25
$ easygen -ts '{{stringsContains .Name "init"}}' /tmp/var
true

$ easygen -ts '{{stringsContains .Name "foobar"}}' /tmp/var
false

$ easygen -ts '{{if (stringsContains .Name "init")}}"{{.Name}}" contains "init"{{else}} there {{end}}!' /tmp/var
"some-init-method" contains "init"!

$ easygen -ts '{{if (stringsContains .Name "init")}}"{{.Name}}" splits into: {{range (stringsSplit .Name "-")}}{{.}} {{end}}{{else}} there {{end}}!' /tmp/var
"some-init-method" splits into: some init method !

$ echo 'Name: "%bob'"'"'s file"' | tee /tmp/var.yaml
Name: "%bob's file"

$ easygen -ts '{{quote4shell .Name}}' /tmp/var
'%bob'\''s file'

# Using environment variable
$ easygen -ts '{{ENV "SHELL"}}: {{.Age}}' /tmp/age
/bin/bash: 16

```
