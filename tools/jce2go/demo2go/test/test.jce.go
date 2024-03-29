// DO NOT EDIT IT.
// code generated by jce2go v1.0.
// source: test.jce
package test

import (
	"fmt"
	"io"

	"github.com/edte/erpc/codec/jce"
	"github.com/edte/jce2go/demo2go/base"
)

// 占位使用，避免导入的这些包没有被使用
var _ = fmt.Errorf
var _ = io.ReadFull
var _ = jce.BYTE

// RequestPacket struct implement
type RequestPacket struct {
	B       int8                    `json:"b" tag:"1"`
	S       int16                   `json:"s" tag:"2"`
	I       int32                   `json:"i" tag:"3"`
	L       int64                   `json:"l" tag:"4"`
	F       float32                 `json:"f" tag:"5"`
	D       float64                 `json:"d" tag:"6"`
	S1      string                  `json:"s1" tag:"7"`
	S2      string                  `json:"s2" tag:"8"`
	I2      int32                   `json:"i2" tag:"9"`
	Buffer1 []int8                  `json:"buffer1" tag:"10"`
	Buffer2 []uint8                 `json:"buffer2" tag:"11"`
	Arr1    []string                `json:"arr1" tag:"12"`
	Arr2    [][]int32               `json:"arr2" tag:"13"`
	M1      map[string]string       `json:"m1" tag:"14"`
	Arr4    []map[int32]string      `json:"arr4" tag:"15"`
	Arr3    []base.Request          `json:"arr3" tag:"16"`
	M2      map[string]base.Request `json:"m2" tag:"17"`
}

func (st *RequestPacket) resetDefault() {
	st.S2 = "test"
}

// ReadFrom reads from io.Reader and put into struct.
func (st *RequestPacket) ReadFrom(r io.Reader) (n int64, err error) {
	var (
		have bool
		ty   jce.JceEncodeType
	)

	decoder := jce.NewDecoder(r)
	st.resetDefault()

	// [step 1] read B
	if err = decoder.ReadInt8(&st.B, 1, true); err != nil {
		return
	}
	// [step 2] read S
	if err = decoder.ReadInt16(&st.S, 2, true); err != nil {
		return
	}
	// [step 3] read I
	if err = decoder.ReadInt32(&st.I, 3, true); err != nil {
		return
	}
	// [step 4] read L
	if err = decoder.ReadInt64(&st.L, 4, true); err != nil {
		return
	}
	// [step 5] read F
	if err = decoder.ReadFloat32(&st.F, 5, true); err != nil {
		return
	}
	// [step 6] read D
	if err = decoder.ReadFloat64(&st.D, 6, true); err != nil {
		return
	}
	// [step 7] read S1
	if err = decoder.ReadString(&st.S1, 7, true); err != nil {
		return
	}
	// [step 8] read S2
	if err = decoder.ReadString(&st.S2, 8, false); err != nil {
		return
	}
	// [step 9] read I2
	if err = decoder.ReadInt32(&st.I2, 9, false); err != nil {
		return
	}
	// [step 10] read Buffer1
	if err = decoder.ReadSliceInt8(&st.Buffer1, 10, true); err != nil {
		return
	}
	// [step 11] read Buffer2
	if err = decoder.ReadSliceUint8(&st.Buffer2, 11, true); err != nil {
		return
	}
	// [step 12] read Arr1
	var length2 uint32

	// [step 12.1] read type、tag
	if ty, have, err = decoder.ReadHead(12, false); err != nil || !have {
		return
	}
	// [step 12.2] read list length
	if length2, err = decoder.ReadLength(); err != nil {
		return
	}
	// [step 12.3] read data
	st.Arr1 = make([]string, length2)
	for i2 := uint32(0); i2 < length2; i2++ {
		// [step 0] read Arr1[i2]
		if err = decoder.ReadString(&st.Arr1[i2], 0, false); err != nil {
			return
		}

	}
	// [step 13] read Arr2
	var length3 uint32

	// [step 13.1] read type、tag
	if ty, have, err = decoder.ReadHead(13, false); err != nil || !have {
		return
	}
	// [step 13.2] read list length
	if length3, err = decoder.ReadLength(); err != nil {
		return
	}
	// [step 13.3] read data
	st.Arr2 = make([][]int32, length3)
	for i3 := uint32(0); i3 < length3; i3++ {
		// [step 0] read Arr2[i3]
		var length4 uint32

		// [step 0.1] read type、tag
		if ty, have, err = decoder.ReadHead(0, false); err != nil || !have {
			return
		}
		// [step 0.2] read list length
		if length4, err = decoder.ReadLength(); err != nil {
			return
		}
		// [step 0.3] read data
		st.Arr2[i3] = make([]int32, length4)
		for i4 := uint32(0); i4 < length4; i4++ {
			// [step 0] read Arr2[i3][i4]
			if err = decoder.ReadInt32(&st.Arr2[i3][i4], 0, false); err != nil {
				return
			}

		}

	}
	// [step 14] read M1
	var length5 uint32

	// [step 14.1] read type、tag
	if ty, have, err = decoder.ReadHead(14, true); err != nil {
		return
	}
	// [step 14.2] read length
	if length5, err = decoder.ReadLength(); err != nil {
		return
	}
	// [step 14.3] read data
	st.M1 = make(map[string]string, 0)
	var k5 string
	var v5 string
	for i := uint32(0); i < length5; i++ {
		// [step 0] read k5
		if err = decoder.ReadString(&k5, 0, false); err != nil {
			return
		}
		// [step 1] read v5
		if err = decoder.ReadString(&v5, 1, false); err != nil {
			return
		}

		st.M1[k5] = v5
	}
	// [step 15] read Arr4
	var length6 uint32

	// [step 15.1] read type、tag
	if ty, have, err = decoder.ReadHead(15, true); err != nil || !have {
		return
	}
	// [step 15.2] read list length
	if length6, err = decoder.ReadLength(); err != nil {
		return
	}
	// [step 15.3] read data
	st.Arr4 = make([]map[int32]string, length6)
	for i6 := uint32(0); i6 < length6; i6++ {
		// [step 0] read Arr4[i6]
		var length7 uint32

		// [step 0.1] read type、tag
		if ty, have, err = decoder.ReadHead(0, false); err != nil {
			return
		}
		// [step 0.2] read length
		if length7, err = decoder.ReadLength(); err != nil {
			return
		}
		// [step 0.3] read data
		st.Arr4[i6] = make(map[int32]string, 0)
		var k7 int32
		var v7 string
		for i := uint32(0); i < length7; i++ {
			// [step 0] read k7
			if err = decoder.ReadInt32(&k7, 0, false); err != nil {
				return
			}
			// [step 1] read v7
			if err = decoder.ReadString(&v7, 1, false); err != nil {
				return
			}

			st.Arr4[i6][k7] = v7
		}

	}
	// [step 16] read Arr3
	var length8 uint32

	// [step 16.1] read type、tag
	if ty, have, err = decoder.ReadHead(16, true); err != nil || !have {
		return
	}
	// [step 16.2] read list length
	if length8, err = decoder.ReadLength(); err != nil {
		return
	}
	// [step 16.3] read data
	st.Arr3 = make([]base.Request, length8)
	for i8 := uint32(0); i8 < length8; i8++ {
		// [step 0] read Arr3[i8]
		if _, err = st.Arr3[i8].ReadFrom(decoder.Reader()); err != nil {
			return
		}

	}
	// [step 17] read M2
	var length9 uint32

	// [step 17.1] read type、tag
	if ty, have, err = decoder.ReadHead(17, true); err != nil {
		return
	}
	// [step 17.2] read length
	if length9, err = decoder.ReadLength(); err != nil {
		return
	}
	// [step 17.3] read data
	st.M2 = make(map[string]base.Request, 0)
	var k9 string
	var v9 base.Request
	for i := uint32(0); i < length9; i++ {
		// [step 0] read k9
		if err = decoder.ReadString(&k9, 0, false); err != nil {
			return
		}
		// [step 1] read v9
		if _, err = v9.ReadFrom(decoder.Reader()); err != nil {
			return
		}

		st.M2[k9] = v9
	}

	_ = err
	_ = have
	_ = ty
	return
}

// WriteTo encode struct to io.Writer
func (st *RequestPacket) WriteTo(w io.Writer) (n int64, err error) {
	encoder := jce.NewEncoder(w)
	st.resetDefault()

	// [step 1] write B
	if err = encoder.WriteInt8(st.B, 1); err != nil {
		return
	}
	// [step 2] write S
	if err = encoder.WriteInt16(st.S, 2); err != nil {
		return
	}
	// [step 3] write I
	if err = encoder.WriteInt32(st.I, 3); err != nil {
		return
	}
	// [step 4] write L
	if err = encoder.WriteInt64(st.L, 4); err != nil {
		return
	}
	// [step 5] write F
	if err = encoder.WriteFloat32(st.F, 5); err != nil {
		return
	}
	// [step 6] write D
	if err = encoder.WriteFloat64(st.D, 6); err != nil {
		return
	}
	// [step 7] write S1
	if err = encoder.WriteString(st.S1, 7); err != nil {
		return
	}
	// [step 8] write S2
	if err = encoder.WriteString(st.S2, 8); err != nil {
		return
	}
	// [step 9] write I2
	if err = encoder.WriteInt32(st.I2, 9); err != nil {
		return
	}
	// [step 10] write Buffer1
	if err = encoder.WriteSliceInt8(st.Buffer1, 10); err != nil {
		return
	}
	// [step 11] write Buffer2
	if err = encoder.WriteSliceUint8(st.Buffer2, 11); err != nil {
		return
	}
	// [step 12] write Arr1
	// [step 12.1] write type、tag
	if err = encoder.WriteHead(jce.LIST, 12); err != nil {
		return
	}
	// [step 12.2] write list length
	if err = encoder.WriteLength(uint32(len(st.Arr1))); err != nil {
		return
	}
	// [step 12.3] write data
	for _, v12 := range st.Arr1 {
		// [step 0] write v12
		if err = encoder.WriteString(v12, 0); err != nil {
			return
		}
	}
	// [step 13] write Arr2
	// [step 13.1] write type、tag
	if err = encoder.WriteHead(jce.LIST, 13); err != nil {
		return
	}
	// [step 13.2] write list length
	if err = encoder.WriteLength(uint32(len(st.Arr2))); err != nil {
		return
	}
	// [step 13.3] write data
	for _, v13 := range st.Arr2 {
		// [step 0] write v13
		// [step 0.1] write type、tag
		if err = encoder.WriteHead(jce.LIST, 0); err != nil {
			return
		}
		// [step 0.2] write list length
		if err = encoder.WriteLength(uint32(len(v13))); err != nil {
			return
		}
		// [step 0.3] write data
		for _, v14 := range v13 {
			// [step 0] write v14
			if err = encoder.WriteInt32(v14, 0); err != nil {
				return
			}
		}
	}
	// [step 14] write M1
	// [step 14.1] write type、tag
	if err = encoder.WriteHead(jce.MAP, 14); err != nil {
		return
	}
	// [step 14.2] write length
	if err = encoder.WriteLength(uint32(len(st.M1))); err != nil {
		return
	}
	// [step 14.3] write data
	for k15, v15 := range st.M1 {
		// [step 0] write k15
		if err = encoder.WriteString(k15, 0); err != nil {
			return
		}
		// [step 1] write v15
		if err = encoder.WriteString(v15, 1); err != nil {
			return
		}
	}
	// [step 15] write Arr4
	// [step 15.1] write type、tag
	if err = encoder.WriteHead(jce.LIST, 15); err != nil {
		return
	}
	// [step 15.2] write list length
	if err = encoder.WriteLength(uint32(len(st.Arr4))); err != nil {
		return
	}
	// [step 15.3] write data
	for _, v16 := range st.Arr4 {
		// [step 0] write v16
		// [step 0.1] write type、tag
		if err = encoder.WriteHead(jce.MAP, 0); err != nil {
			return
		}
		// [step 0.2] write length
		if err = encoder.WriteLength(uint32(len(v16))); err != nil {
			return
		}
		// [step 0.3] write data
		for k17, v17 := range v16 {
			// [step 0] write k17
			if err = encoder.WriteInt32(k17, 0); err != nil {
				return
			}
			// [step 1] write v17
			if err = encoder.WriteString(v17, 1); err != nil {
				return
			}
		}
	}
	// [step 16] write Arr3
	// [step 16.1] write type、tag
	if err = encoder.WriteHead(jce.LIST, 16); err != nil {
		return
	}
	// [step 16.2] write list length
	if err = encoder.WriteLength(uint32(len(st.Arr3))); err != nil {
		return
	}
	// [step 16.3] write data
	for _, v18 := range st.Arr3 {
		// [step 0] write v18
		if _, err = v18.WriteTo(encoder.Writer()); err != nil {
			return
		}
	}
	// [step 17] write M2
	// [step 17.1] write type、tag
	if err = encoder.WriteHead(jce.MAP, 17); err != nil {
		return
	}
	// [step 17.2] write length
	if err = encoder.WriteLength(uint32(len(st.M2))); err != nil {
		return
	}
	// [step 17.3] write data
	for k19, v19 := range st.M2 {
		// [step 0] write k19
		if err = encoder.WriteString(k19, 0); err != nil {
			return
		}
		// [step 1] write v19
		if _, err = v19.WriteTo(encoder.Writer()); err != nil {
			return
		}
	}

	// flush to io.Writer
	err = encoder.Flush()
	return
}
