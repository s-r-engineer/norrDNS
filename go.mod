module github.com/s-r-engineer/norrDNS

go 1.24.3

replace github.com/getsentry/sentry-go v0.32.0 => github.com/s-r-engineer/sentry-go v0.0.0-20250510204624-fc2a804d7a6f

require (
	github.com/miekg/dns v1.1.66
	github.com/s-r-engineer/library v0.2.14
	gorm.io/driver/sqlite v1.5.7
	gorm.io/gorm v1.26.1
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/getsentry/sentry-go v0.32.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	golang.org/x/tools v0.32.0 // indirect
)
