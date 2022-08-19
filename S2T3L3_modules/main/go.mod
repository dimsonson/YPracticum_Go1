module example/main

go 1.19

replace example/somemodule/somepackage => ../somemodule/somepackage

require example/somemodule/somepackage v0.0.0-00010101000000-000000000000

require (
	github.com/yuin/goldmark v1.4.13 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220812174116-3211cb980234 // indirect
	golang.org/x/sys v0.0.0-20220818161305-2296e01440c6 // indirect
	golang.org/x/tools v0.1.12 // indirect
)
