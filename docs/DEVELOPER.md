
### Developer Notes

```shell
/opt/homebrew/bin/swagger-codegen generate -i etc/openapi.yaml -l go-server
go mod init github.com/cunnie/k-v.io
go mod tidy
 # fix main.go's broken import to
 # `sw "github.com/cunnie/k-v.io/kvio"`
go mod tidy # yes, a second time
```
