package main

import (
	"liangzi.local/qx"
	"liangzi.local/qx/pkg/xjbfs"
)

func get(year, month, day, hour int) *qx.YanQin {
	return qx.NewYanQin(year, month, day, hour)
}

func getXjbf(ygz, mgz, dgz, hgz string) *xjbfs.XJBF {
	return xjbfs.NewXJBF(ygz, mgz, dgz, hgz)
}
