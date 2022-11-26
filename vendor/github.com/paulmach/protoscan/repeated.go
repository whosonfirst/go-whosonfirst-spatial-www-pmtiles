// Code generated by internal/gen_repeated.go. DO NOT EDIT.

package protoscan

// RepeatedFloat will append the repeated value(s) to the buffer.
// This method supports packed or unpacked encoding.
func (m *Message) RepeatedFloat(buf []float32) ([]float32, error) {
	if m.wireType == WireType32bit {
		v, err := m.Float()
		if err != nil {
			return nil, err
		}

		return append(buf, v), nil
	}

	l, err := m.packedLength()
	if err != nil {
		return nil, err
	}

	// if provided we append.
	if buf == nil {
		buf = make([]float32, 0, l/4)
	}

	postIndex := m.Index + l
	for m.Index < postIndex {
		v, err := m.Float()
		if err != nil {
			return nil, err
		}

		buf = append(buf, v)
	}

	return buf, nil
}

// RepeatedDouble will append the repeated value(s) to the buffer.
// This method supports packed or unpacked encoding.
func (m *Message) RepeatedDouble(buf []float64) ([]float64, error) {
	if m.wireType == WireType64bit {
		v, err := m.Double()
		if err != nil {
			return nil, err
		}

		return append(buf, v), nil
	}

	l, err := m.packedLength()
	if err != nil {
		return nil, err
	}

	// if provided we append.
	if buf == nil {
		buf = make([]float64, 0, l/8)
	}

	postIndex := m.Index + l
	for m.Index < postIndex {
		v, err := m.Double()
		if err != nil {
			return nil, err
		}

		buf = append(buf, v)
	}

	return buf, nil
}

// RepeatedInt32 will append the repeated value(s) to the buffer.
// This method supports packed or unpacked encoding.
func (m *Message) RepeatedInt32(buf []int32) ([]int32, error) {
	if m.wireType == WireTypeVarint {
		v, err := m.Int32()
		if err != nil {
			return nil, err
		}

		return append(buf, v), nil
	}

	l, err := m.packedLength()
	if err != nil {
		return nil, err
	}

	// if provided we append.
	if buf == nil {
		buf = make([]int32, 0, m.count(l))
	}

	postIndex := m.Index + l
	for m.Index < postIndex {
		v, err := m.Int32()
		if err != nil {
			return nil, err
		}

		buf = append(buf, v)
	}

	return buf, nil
}

// RepeatedInt64 will append the repeated value(s) to the buffer.
// This method supports packed or unpacked encoding.
func (m *Message) RepeatedInt64(buf []int64) ([]int64, error) {
	if m.wireType == WireTypeVarint {
		v, err := m.Int64()
		if err != nil {
			return nil, err
		}

		return append(buf, v), nil
	}

	l, err := m.packedLength()
	if err != nil {
		return nil, err
	}

	// if provided we append.
	if buf == nil {
		buf = make([]int64, 0, m.count(l))
	}

	postIndex := m.Index + l
	for m.Index < postIndex {
		v, err := m.Int64()
		if err != nil {
			return nil, err
		}

		buf = append(buf, v)
	}

	return buf, nil
}

// RepeatedUint32 will append the repeated value(s) to the buffer.
// This method supports packed or unpacked encoding.
func (m *Message) RepeatedUint32(buf []uint32) ([]uint32, error) {
	if m.wireType == WireTypeVarint {
		v, err := m.Uint32()
		if err != nil {
			return nil, err
		}

		return append(buf, v), nil
	}

	l, err := m.packedLength()
	if err != nil {
		return nil, err
	}

	// if provided we append.
	if buf == nil {
		buf = make([]uint32, 0, m.count(l))
	}

	postIndex := m.Index + l
	for m.Index < postIndex {
		v, err := m.Uint32()
		if err != nil {
			return nil, err
		}

		buf = append(buf, v)
	}

	return buf, nil
}

// RepeatedUint64 will append the repeated value(s) to the buffer.
// This method supports packed or unpacked encoding.
func (m *Message) RepeatedUint64(buf []uint64) ([]uint64, error) {
	if m.wireType == WireTypeVarint {
		v, err := m.Uint64()
		if err != nil {
			return nil, err
		}

		return append(buf, v), nil
	}

	l, err := m.packedLength()
	if err != nil {
		return nil, err
	}

	// if provided we append.
	if buf == nil {
		buf = make([]uint64, 0, m.count(l))
	}

	postIndex := m.Index + l
	for m.Index < postIndex {
		v, err := m.Uint64()
		if err != nil {
			return nil, err
		}

		buf = append(buf, v)
	}

	return buf, nil
}

// RepeatedSint32 will append the repeated value(s) to the buffer.
// This method supports packed or unpacked encoding.
func (m *Message) RepeatedSint32(buf []int32) ([]int32, error) {
	if m.wireType == WireTypeVarint {
		v, err := m.Sint32()
		if err != nil {
			return nil, err
		}

		return append(buf, v), nil
	}

	l, err := m.packedLength()
	if err != nil {
		return nil, err
	}

	// if provided we append.
	if buf == nil {
		buf = make([]int32, 0, m.count(l))
	}

	postIndex := m.Index + l
	for m.Index < postIndex {
		v, err := m.Sint32()
		if err != nil {
			return nil, err
		}

		buf = append(buf, v)
	}

	return buf, nil
}

// RepeatedSint64 will append the repeated value(s) to the buffer.
// This method supports packed or unpacked encoding.
func (m *Message) RepeatedSint64(buf []int64) ([]int64, error) {
	if m.wireType == WireTypeVarint {
		v, err := m.Sint64()
		if err != nil {
			return nil, err
		}

		return append(buf, v), nil
	}

	l, err := m.packedLength()
	if err != nil {
		return nil, err
	}

	// if provided we append.
	if buf == nil {
		buf = make([]int64, 0, m.count(l))
	}

	postIndex := m.Index + l
	for m.Index < postIndex {
		v, err := m.Sint64()
		if err != nil {
			return nil, err
		}

		buf = append(buf, v)
	}

	return buf, nil
}

// RepeatedFixed32 will append the repeated value(s) to the buffer.
// This method supports packed or unpacked encoding.
func (m *Message) RepeatedFixed32(buf []uint32) ([]uint32, error) {
	if m.wireType == WireType32bit {
		v, err := m.Fixed32()
		if err != nil {
			return nil, err
		}

		return append(buf, v), nil
	}

	l, err := m.packedLength()
	if err != nil {
		return nil, err
	}

	// if provided we append.
	if buf == nil {
		buf = make([]uint32, 0, l/4)
	}

	postIndex := m.Index + l
	for m.Index < postIndex {
		v, err := m.Fixed32()
		if err != nil {
			return nil, err
		}

		buf = append(buf, v)
	}

	return buf, nil
}

// RepeatedFixed64 will append the repeated value(s) to the buffer.
// This method supports packed or unpacked encoding.
func (m *Message) RepeatedFixed64(buf []uint64) ([]uint64, error) {
	if m.wireType == WireType64bit {
		v, err := m.Fixed64()
		if err != nil {
			return nil, err
		}

		return append(buf, v), nil
	}

	l, err := m.packedLength()
	if err != nil {
		return nil, err
	}

	// if provided we append.
	if buf == nil {
		buf = make([]uint64, 0, l/8)
	}

	postIndex := m.Index + l
	for m.Index < postIndex {
		v, err := m.Fixed64()
		if err != nil {
			return nil, err
		}

		buf = append(buf, v)
	}

	return buf, nil
}

// RepeatedSfixed32 will append the repeated value(s) to the buffer.
// This method supports packed or unpacked encoding.
func (m *Message) RepeatedSfixed32(buf []int32) ([]int32, error) {
	if m.wireType == WireType32bit {
		v, err := m.Sfixed32()
		if err != nil {
			return nil, err
		}

		return append(buf, v), nil
	}

	l, err := m.packedLength()
	if err != nil {
		return nil, err
	}

	// if provided we append.
	if buf == nil {
		buf = make([]int32, 0, l/4)
	}

	postIndex := m.Index + l
	for m.Index < postIndex {
		v, err := m.Sfixed32()
		if err != nil {
			return nil, err
		}

		buf = append(buf, v)
	}

	return buf, nil
}

// RepeatedSfixed64 will append the repeated value(s) to the buffer.
// This method supports packed or unpacked encoding.
func (m *Message) RepeatedSfixed64(buf []int64) ([]int64, error) {
	if m.wireType == WireType64bit {
		v, err := m.Sfixed64()
		if err != nil {
			return nil, err
		}

		return append(buf, v), nil
	}

	l, err := m.packedLength()
	if err != nil {
		return nil, err
	}

	// if provided we append.
	if buf == nil {
		buf = make([]int64, 0, l/8)
	}

	postIndex := m.Index + l
	for m.Index < postIndex {
		v, err := m.Sfixed64()
		if err != nil {
			return nil, err
		}

		buf = append(buf, v)
	}

	return buf, nil
}

// RepeatedBool will append the repeated value(s) to the buffer.
// This method supports packed or unpacked encoding.
func (m *Message) RepeatedBool(buf []bool) ([]bool, error) {
	if m.wireType == WireTypeVarint {
		v, err := m.Bool()
		if err != nil {
			return nil, err
		}

		return append(buf, v), nil
	}

	l, err := m.packedLength()
	if err != nil {
		return nil, err
	}

	// if provided we append.
	if buf == nil {
		buf = make([]bool, 0, l)
	}

	postIndex := m.Index + l
	for m.Index < postIndex {
		v, err := m.Bool()
		if err != nil {
			return nil, err
		}

		buf = append(buf, v)
	}

	return buf, nil
}
