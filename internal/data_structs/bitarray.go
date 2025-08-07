package data_structs

import "encoding/binary"

// Realizes operations with bits using binary operations.
// Needed for processing values with dynamic binary size
type BitArray struct {
	counter uint8
	data    uint64
}

// Add value to the right of bitarray, pass 0 to add 0, other values will be interpreted as 1
func (bs BitArray) PushRight(val int) BitArray {
	var add_val uint64
	if val == 0 {
		add_val = uint64(0)
	} else {
		add_val = uint64(1)
	}
	return BitArray{
		counter: bs.counter + 1,
		data:    (bs.data << 1) | add_val,
	}
}

// Add value to the left of bitarray, pass 0 to add 0, other values will be interpreted as 1
func (bs BitArray) PushLeft(val int) BitArray {
	var add_val uint64
	if val == 0 {
		add_val = 0
	} else {
		add_val = 1
	}
	var pointer uint64 = add_val << uint64(bs.counter)
	return BitArray{
		counter: bs.counter + 1,
		data:    bs.data | pointer,
	}
}

// merges 2 bitstrings together, argument bitstring will be added to the right
func (bs BitArray) Merge(r_bs BitArray) BitArray {
	new_data := bs.data << uint64(r_bs.counter)
	new_data |= r_bs.data

	return BitArray{
		counter: bs.counter + r_bs.counter,
		data:    new_data,
	}

}

// Get leftmost bytes.
// Will retun zerolength array if not enough data is provedided
func (bs BitArray) PopLeftBytes() (BitArray, []byte) {
	byte_len := bs.counter / 8
	if byte_len == 0 {
		return bs, []byte{}
	}
	// truncate rigtmost bits not included in return
	trunc_data := bs.data >> uint64(bs.counter-byte_len*8)
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], trunc_data)
	return BitArray{
		counter: byte_len * 8,
		data:    trunc_data,
	}, buf[8-byte_len:]
}
