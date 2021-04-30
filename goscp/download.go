package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"net"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
)

//遠程服務器目錄下載到本地指定位置
func downloadRemoteDirToLocal(conf Conf) {
	u, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	dir := u.HomeDir
	var localDir string //指定本地位置
	switch runtime.GOOS {
	case "windows":
		localDir = dir + "\\DownloadRemoteFiles\\"
	case "linux":
		localDir = dir + "/DownloadRemoteFiles/"
	}

	//創建本地下載目錄 同時指定目錄權限0755(drwxr-xr-x)
	err = os.MkdirAll(localDir, os.ModeDir|os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating DownloadRemoteFiles directory: %s", err.Error())
		panic(err)
	}

	user := conf.User
	pass := conf.Pass
	host := conf.Host
	remoteDir := conf.RemotePath

	if conf.LocalPath != "" {
		localDir = conf.LocalPath
	}

	//ssh
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: func(host string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	conn, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//sftp
	sftpClient, err := sftp.NewClient(conn, sftp.MaxPacket(20480))
	if err != nil {
		log.Fatal(err)
	}
	defer sftpClient.Close()

	files, err := sftpClient.ReadDir(remoteDir)
	if err != nil {
		log.Fatal(err)
	}
	//
	for _, file := range files {
		srcInvPath := path.Join(remoteDir, file.Name())
		dstInvPath := filepath.Join(localDir, file.Name())

		dstFile, err := os.Create(dstInvPath)
		if err != nil {
			log.Fatal(err)
		}
		defer dstFile.Close()

		srcFile, err := sftpClient.OpenFile(srcInvPath, os.O_RDONLY)
		if err != nil {
			log.Fatal(err)
		}

		bytes, err := io.Copy(dstFile, srcFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("copied %s to %s transport:%d bytes\n", file.Name(), localDir, bytes)

		err = dstFile.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}
}
