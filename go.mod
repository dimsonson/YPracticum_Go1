module Stepik

go 1.19

// директивой replace указываем положение корня
// модуля ypmodule относительно main/go.mod
replace toppackage/middlepackage/bottompackage/mathxxx => /Users/dbo/Yandex.Disk.localized/go/src/github.com/dimsonson/Stepik/toppackage/middlepackage/bottompackage/mathxxx

require toppackage/middlepackage/bottompackage/mathxxx v0.0.0-00010101000000-000000000000 // indirect
