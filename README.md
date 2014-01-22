json.tool
=========
A small program to validate and pretty-print from the command line

The tool is similar to python json.tool except it can handle multiple
json objects and would print all of them.

Installation
------------

```bash
go get github.com/ghais/json.tool
````

Examples
--------

```bash
$ echo '{"json":"obj"}' | json.tool
{

	"json": "obj"

}

$ echo '{1.2:3.4}' | json.tool
Einvalid character '1' looking for beginning of object key string
```

The tool was developed to format the output produced by mongoexport

```bash
$ mongoexport -d test_db -c test_col | json.tool > export.formatted.txt
```
