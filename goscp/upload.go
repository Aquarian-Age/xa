package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path"
	"path/filepath"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

//scp 上传 本地目录到远程(仅限目录 无法指定单个文件)
func uploadLocalDirToRemote(conf Conf) {
	user := conf.User
	password := conf.Pass
	localHome := conf.LocalPath
	remoteHome := conf.RemotePath
	host := conf.Host

	if exists, err := PathExists(localHome); err != nil {
		log.Fatal(err)
	} else if exists == false {
		log.Fatal(localHome + " not exists")
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: func(host string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	c, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal(err)
	}

	UploadDir(c, localHome, remoteHome)

}

//
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//
func UploadDir(c *ssh.Client, localHome string, remoteHome string) error {

	//
	sfc, err := sftp.NewClient(c, sftp.MaxPacketUnchecked(32768))
	if err != nil {
		return err
	}
	defer sfc.Close()
	/////////////////////////////
	files, err := os.ReadDir(localHome)
	if err != nil {
		fmt.Println("read local dir:", err)
	}
	for _, file := range files {
		srcInvPath := path.Join(localHome, file.Name())
		dstInvPath := filepath.Join(remoteHome, file.Name())

		dest, err := sfc.Create(dstInvPath)
		if err != nil {
			fmt.Println("遠程文件創建失敗:", err)
		}
		defer dest.Close()

		src, err := os.Open(srcInvPath)
		if err != nil {
			fmt.Println("本地文件打開失敗:", err)
		}
		defer src.Close()

		//只能傳輸當前目錄下的文件 不能傳輸當前目錄下的子目錄
		//這裏如果碰到有子目錄無法傳輸子目錄
		bytes, err := io.Copy(dest, src)
		if err != nil {
			fmt.Printf("copy:%v\n", err.Error())
		}
		log.Printf("sending file %s to %s transport:%d bytes\n", file.Name(), remoteHome, bytes)
	}

	return err
}
