module github.com/kivra/krakend-shadowproxy

go 1.24.0

toolchain go1.24.5

require (
	github.com/luraproject/lura/v2 v2.12.0
	golang.org/x/text v0.29.0
)

require (
	github.com/krakend/flatmap v1.2.0 // indirect
	github.com/valyala/fastrand v1.1.0 // indirect
)

replace github.com/luraproject/lura/v2 v2.2.4 => github.com/kivra/lura/v2 v2.2.4-0.20230228151500-8e359d392b5d
