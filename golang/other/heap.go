package other

type Heap struct {
	queue []int
	size  int
}

func (h *Heap) Push(num int) {
	if h.size <= len(h.queue) {
		h.queue = append(h.queue, num)
		h.size++
	}
	for i := h.size - 1; i > 0; {
		p := (i - 1) / 2
		if h.queue[p] > h.queue[i] {
			h.queue[p], h.queue[i] = h.queue[i], h.queue[p]
			i = p
		} else {
			break
		}
	}
}

func (h *Heap) Pop() (int, bool) {
	if h.size <= 0 {
		return 0, false
	}
	if h.size == 1 {
		h.size--
		return h.queue[0], true
	}
	h.size--
	num := h.queue[0]
	h.queue[0] = h.queue[h.size]
	for i := 0; i < h.size; {
		l, r := 2*i+1, 2*i+2
		if l >= h.size {
			break
		}
		if r >= h.size {
			if h.queue[i] > h.queue[l] {
				h.queue[i], h.queue[l] = h.queue[l], h.queue[i]
			}
			break
		}
		if h.queue[i] < h.queue[l] && h.queue[i] < h.queue[r] {
			break
		}
		if h.queue[l] < h.queue[r] {
			h.queue[i], h.queue[l] = h.queue[l], h.queue[i]
			i = l
		} else {
			h.queue[i], h.queue[r] = h.queue[r], h.queue[i]
			i = r
		}
	}

	return num, true
}

func (h *Heap) Top() (int, bool) {
	if h.size <= 0 {
		return 0, false
	}
	return h.queue[0], true

}

func (h *Heap) Size() int {
	return h.size
}
