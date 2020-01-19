// +build !windows

package vfs

import "syscall"

import "os"

var (
	defaultUID, defaultGID int
)

func init() {
	defaultUID = os.Getuid()
	defaultGID = os.Getuid()
	if defaultUID < 0 {
		defaultUID = 65534
	}
	if defaultGID < 0 {
		defaultGID = 65534
	}
}

func (fi S3FileInfo) getFileInfoSys() interface{} {
	nlink := uint64(1)
	if fi.IsDir() {
		nlink = 0
	}
	return &syscall.Stat_t{Nlink: nlink,
		Uid: uint32(defaultUID),
		Gid: uint32(defaultGID)}
}
