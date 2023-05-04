package GeeCache

// 本地需实现peerpicker
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// 远程节点需实现peerGetter
type PeerGetter interface {
	Get(Group string, key string) ([]byte, error)
}