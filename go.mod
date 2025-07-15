module github.com/kivra/krakend-shadowproxy

go 1.23.0

toolchain go1.24.5

require (
	github.com/luraproject/lura/v2 v2.10.2
	golang.org/x/text v0.27.0
)

require (
	github.com/krakendio/flatmap v1.1.1 // indirect
	github.com/valyala/fastrand v1.1.0 // indirect
)

replace github.com/luraproject/lura/v2 v2.2.4 => github.com/kivra/lura/v2 v2.2.4-0.20230228151500-8e359d392b5d
