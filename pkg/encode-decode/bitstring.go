package encodedecode_ssaf2

import "encoding/binary"

// Realizes operations with bits using binary operations.
// Needed for processing values with dynamic binary size
type BitString struct {
	counter uint8
	data    uint64
}

// Create empty bitsting
func NewBitString() BitString {
	/*
		Create empty bitstring
	*/
	return BitString{
		counter: 0,
		data:    0,
	}
}

// Add value to the right of bitstring
func (bs BitString) PushRight(val bool) BitString {
	var add_val uint64
	if val {
		add_val = uint64(1)
	} else {
		add_val = uint64(0)
	}
	return BitString{
		counter: bs.counter + 1,
		data:    (bs.data << 1) | add_val,
	}
}

// Add value to the left of bitstring
func (bs BitString) PushLeft(val bool) BitString {
	var add_val uint64
	if val {
		add_val = 1
	} else {
		add_val = 0
	}
	var pointer uint64 = add_val << uint64(bs.counter)
	return BitString{
		counter: bs.counter + 1,
		data:    bs.data | pointer,
	}
}

// merges 2 bitstrings together, argument bitstring will be added to the right
func (bs BitString) Merge(r_bs BitString) BitString {
	new_data := bs.data << uint64(r_bs.counter)
	new_data |= r_bs.data

	return BitString{
		counter: bs.counter + r_bs.counter,
		data:    new_data,
	}

}

// Get leftmost bytes.
// Will retun zerolength array if not enough data is provedided
func (bs BitString) PopLeftBytes() (BitString, []byte) {
	byte_len := bs.counter / 8
	if byte_len == 0 {
		return bs, []byte{}
	}
	// truncate rigtmost bits not included in return
	trunc_data := bs.data >> uint64(bs.counter-byte_len*8)
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], trunc_data)
	return BitString{
		counter: byte_len * 8,
		data:    trunc_data,
	}, buf[8-byte_len:]
}
