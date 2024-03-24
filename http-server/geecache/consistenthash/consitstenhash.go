package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

/*
replicas 是虚拟节点倍数
keys是哈希环
hashMap是虚拟节点与真实节点的映射表
*/
type Map struct {
	hash     Hash
	replicas int
	keys     []int
	hashMap  map[int]string
}

// 允许自定义虚拟节点倍数和Hash函数
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// 把一些值加入到哈希中去
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

/*
选择节点
首先计算key的哈希值
然后顺时针找到第一个匹配的虚拟节点的下表idx，从m.keys中取到对应的哈希值
如果idx == len(m.keys) 则说明需要取到0，所以直接用一个取模来处理这样的情况
*/
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
