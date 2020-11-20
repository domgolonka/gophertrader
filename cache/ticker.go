package cache

import "container/list"

type LRU struct {
	items map[interface{}]*list.Element
}
