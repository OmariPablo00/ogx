module github.com/niconical/ogx/example/migrate

go 1.18

replace github.com/niconical/ogx => ../..

replace github.com/niconical/ogx/dbfixture => ../../dbfixture

replace github.com/niconical/ogx/extra/ogxdebug => ../../extra/ogxdebug

replace github.com/niconical/ogx/dialect/ogdialect => ../../dialect/ogdialect

require (
	gitee.com/opengauss/openGauss-connector-go-pq v1.0.4
	github.com/niconical/ogx v1.1.7
	github.com/niconical/ogx/dialect/ogdialect v0.0.0-20220911050444-60ec07fd608c
	github.com/niconical/ogx/extra/ogxdebug v0.0.0-20220911050444-60ec07fd608c
	github.com/urfave/cli/v2 v2.16.3
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
	golang.org/x/sys v0.0.0-20220913175220-63ea55921009 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
)
