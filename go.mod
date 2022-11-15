module github.com/kivra/krakend-shadowproxy

go 1.19

require (
	github.com/luraproject/lura/v2 v2.0.5
	golang.org/x/text v0.4.0
)

require (
	github.com/krakendio/flatmap v1.1.1 // indirect
	github.com/valyala/fastrand v1.1.0 // indirect
)

replace github.com/luraproject/lura/v2 v2.0.5 => github.com/kivra/lura/v2 v2.0.6-0.20220705080224-3d4d9afe2806
