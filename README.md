Resource Package Generator for goland
=====================================

Very simple generator used to embed some file into executable. Usefull to store template files.

usage :
```
Usage of res2go:
res2go [-package name] files...
  -package string
        package name (default "resources")

examples:
res2go ./resources

res2go ./templates/index.html ./templates/style.js

res2go -p myresources index.html
```
