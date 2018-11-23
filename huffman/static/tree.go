package static

type hItem struct {
	free     bool
	children [2]string
	weight   uint32
	val      string
}

func (i hItem) isEmpty() bool {
	return i.weight < 1
}

type hTree map[string]hItem

func (h hTree) getLastFree() hItem {
	for _, v := range h {
		if v.free {
			return v
		}
	}
	return hItem{}
}

func (h hTree) hasSingleFree() bool {
	r := false
	for _, v := range h {
		if v.free {
			if r {
				return false
			}
			r = true
		}
	}
	return r
}

func (h hTree) checkForTrivial() (map[byte]uint32, bool) {
	if len(h) <= 1 {
		r := make(map[byte]uint32)
		for k, _ := range h {
			r[[]byte(k)[0]] = 0
		}
		return r, true
	} else {
		return nil, false
	}
}

func (h hTree) getMin2() (hItem, hItem) {
	var (
		r1 hItem
		r2 hItem
	)

	for _, o := range h {
		if o.free {
			if r1.isEmpty() {
				r1 = o
			} else if r2.isEmpty() {
				if o.weight < r1.weight {
					r2 = r1
					r1 = o
				} else {
					r2 = o
				}
			} else if o.weight < r1.weight {
				r2 = r1
				r1 = o
			} else if o.weight < r2.weight {
				r2 = o
			}
		}
	}

	return r1, r2
}

func (h *hTree) buildTree() hItem {
	for !h.hasSingleFree() {
		i1, i2 := h.getMin2()
		n := hItem{weight: i1.weight + i2.weight, val: i1.val + i2.val, free: true}
		n.children[0] = i1.val
		n.children[1] = i2.val
		(*h)[n.val] = n
		i1.free = false
		i2.free = false
		(*h)[i1.val] = i1
		(*h)[i2.val] = i2
	}
	return h.getLastFree()
}

func (h hTree) fillCodes(i hItem, code uint32, m *map[string]uint32) {
	if len(i.val) == 1 {
		(*m)[i.val] = code
	} else {
		for i, v := range i.children {
			h.fillCodes(h[v], (code<<1)|uint32(i), m)
		}
	}
}

func BuildHuffmanTree(src *[]byte) hTree {
	r := make(map[string]hItem)
	for _, b := range *src {
		key := string([]byte{b})
		n, inMap := r[key]
		if inMap {
			n.weight++
			r[key] = n
		} else {
			r[key] = hItem{free: true, weight: 1, val: key}
		}
	}
	return r
}

func (h *hTree) get1code(r *map[string]uint32) {
	for k, _ := range *h {
		(*r)[k] = uint32(0)
	}
}

func (h *hTree) GetCodes() map[string]uint32 {
	r := make(map[string]uint32)
	switch len(*h) {
	case 0:
		return r
	case 1:
		h.get1code(&r)
	default:
		hi := h.buildTree()
		h.fillCodes(hi, uint32(0), &r)
	}
	return r
}
