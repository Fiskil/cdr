### Applying patches 

`patch_*.go` file contains patched go types which support unmashalling different JSON variants from AER energy plans endpoints. When generating `energy.gen.go`, you need to apply JSON-patch to swagger file first, and generate only afterwards. 

You can apply JSON-patch with [json-patch](https://github.com/evanphx/json-patch) or any tool that supports RFC6902 JSON patches. 

```
$ cat cdr_energy.swagger.json | json-patch -p cdr_energy.swagger.json.patch > patched_cdr_energy.swagger.json

$ oapi-codegen --old-config-style -generate client,types -package energy patched_cdr_energy.swagger.json > energy.gen.go
```
