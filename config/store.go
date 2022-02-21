package config

import (
	cmn "filestore-server/common"
)

const (
	// TempLocalRootDir : Local temporary storage path
	TempLocalRootDir = "/data/fileserver/"
	// TempPartRootDir : Local temporary storage path for chunked files
	TempPartRootDir = "/data/fileserver_part/"
	// CephRootDir : Ceph's path prefix
	CephRootDir = "/ceph"
	// OSSRootDir : OSS storage path prefix
	OSSRootDir = "oss/"
	// CurrentStoreType : Set the storage type of the current file
	CurrentStoreType = cmn.StoreLocal
)
