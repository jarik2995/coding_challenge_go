package dict

// Example of map implementation
type Dict struct {
	items [50][][]interface{}
}

func (dict *Dict) hash(key string) (adr int) {
	for i := 0; i < len(key); i++ {
		adr = (adr + int(key[i])*i) % len(dict.items)
	}
	return
}

func (dict *Dict) Set(key string, value interface{}) {
	dict.items[dict.hash(key)] = append(dict.items[dict.hash(key)], []interface{}{key, value})
}

func (dict *Dict) Get(key string) (value interface{}) {
	for _, i := range dict.items[dict.hash(key)] {
		if i[0] == key {
			value = i[1]
		}
	}
	return
}

func (dict *Dict) Keys() (keys []interface{}) {
	for _, bucks := range dict.items {
		for _, buck := range bucks {
			keys = append(keys, buck[0])
		}
	}
	return
}

func (dict *Dict) Values() (keys []interface{}) {
	for _, bucks := range dict.items {
		for _, buck := range bucks {
			keys = append(keys, buck[1])
		}
	}
	return
}
