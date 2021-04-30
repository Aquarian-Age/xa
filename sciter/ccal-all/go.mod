module liangzi.local/ccal-sciter

go 1.15

require (
	github.com/lxn/win v0.0.0-20201111105847-2a20daff6a55 // indirect
	github.com/sciter-sdk/go-sciter v0.5.0
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	golang.org/x/sys v0.0.0-20201223074533-0d417f636930 // indirect
	liangzi.local/nongli v0.0.0
	liangzi.local/qx v0.0.0
	liangzi.local/sjqm v0.0.0
	liangzi.local/ts v0.0.0
)

replace (
	liangzi.local/nongli => /home/xuan/src/ccal-cli/
	liangzi.local/qx => /home/xuan/src/qxqm/v2
	liangzi.local/sjqm => /home/xuan/src/sjqm/
	liangzi.local/ts => /home/xuan/src/ts/v0.6.9/
)
