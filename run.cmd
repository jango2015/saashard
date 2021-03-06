@echo off

go get github.com/go-yaml/yaml

go tool yacc -o %~dp0sqlparser\yacc.go -v %~dp0sqlparser\yacc.output %~dp0sqlparser\yacc.y && gofmt -w %~dp0sqlparser\yacc.go

go install github.com/berkaroad/saashard/utils/golog
go install github.com/berkaroad/saashard/config
go install github.com/berkaroad/saashard/errors
go install github.com/berkaroad/saashard/sqlparser/sqltypes
go install github.com/berkaroad/saashard/sqlparser
go install github.com/berkaroad/saashard/backend/mysql
go install github.com/berkaroad/saashard/backend
go install github.com/berkaroad/saashard/net/mysql
go install github.com/berkaroad/saashard/route
go install github.com/berkaroad/saashard/proxy
go install github.com/berkaroad/saashard/admin
go install github.com/berkaroad/saashard/statistic
go install github.com/berkaroad/saashard/server

go build -o %~dp0bin\saashard.exe github.com/berkaroad/saashard/cmd/saashard
%~dp0bin\saashard.exe -config %~dp0conf\ss.yaml

@pause