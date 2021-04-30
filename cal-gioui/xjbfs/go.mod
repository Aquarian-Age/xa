module zr

go 1.15

require (
	gioui.org v0.0.0-20210209100840-ebfb17ec6c56
	gioui.org/cmd v0.0.0-20210209100840-ebfb17ec6c56 // indirect
	github.com/gonoto/notosans v0.0.0-20200703162533-d78fef05ce80
	liangzi.local/cal v0.0.0
	liangzi.local/xjbfs v0.0.0
)

replace (
	liangzi.local/cal => ../../../../cal
	liangzi.local/xjbfs => ../../
)
