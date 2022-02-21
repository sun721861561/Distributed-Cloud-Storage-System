package common

// Storage type ( denotes where the file is stored )
type StoreType int

const (
	_ StoreType = iota
	// StoreLocal : Node Local
	StoreLocal
	// StoreCeph : Ceph Cluster
	StoreCeph
	// StoreOSS : Ali OSS
	StoreOSS
	// StoreMix : Mixed (Ceph and OSS)
	StoreMix
	// StoreAll : All types of storage store one copy of the data
	StoreAll
)