/*
 * The code is automatically generated by the Goland.
 * Copyright © Aquarian-Age. All Rights Reserved.
 * Licensed under MIT
 */

package cmd

import (
	"fmt"
	"os"
)

var (
	Version   = ""
	GoVersion = ""
	GitCommit = ""
	BuildTime = ""
)

//输出版本信息
func PrintVersion() {
	fmt.Printf("Version: %s\n", Version)
	fmt.Printf("Go Version: %s\n", GoVersion)
	fmt.Printf("Git Commit: %s\n", GitCommit)
	fmt.Printf("Build Time: %s\n", BuildTime)
	os.Exit(0)
}
