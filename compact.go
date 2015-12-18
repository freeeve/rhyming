package main

import (
	"bytes"
	"fmt"
	"sort"
)

func AppendBytes(bs []byte, id int32) []byte {
	//fmt.Printf("encoded before: %x, %d\n", bs, id)
	ids := DecodeBytes(bs)
	ids = append(ids, id)
	SortInt32s(ids)
	br := bytes.Buffer{}
	for i, x := range ids {
		if i != 0 {
			x -= ids[i-1]
		}
		if x < 1<<7 {
			br.WriteByte(byte(x))
		} else if x < 1<<14 {
			br.WriteByte(byte(x>>8) | (2 << 6))
			br.WriteByte(byte(x))
		} else if x < 1<<22 {
			br.WriteByte(byte(x>>16) | (3 << 6))
			br.WriteByte(byte(x >> 8))
			br.WriteByte(byte(x))
		} else {
			panic("int out of range for compaction scheme")
		}
	}
	//fmt.Println("ids:", ids)
	//fmt.Printf("encoded after: %x\n", br.Bytes())
	return br.Bytes()
}

/*
* Need to store ~2^18 ids, unfortunately won't fit in 2 bytes for all cases... but we can pack it pretty tight.
* |1st byte|2nd byte|3rd byte|
* |GHIIIIII|JJJJJJJJ|KKKKKKKK|
* G = size indicator, 0 means 1-127
* H = size indicator or part of the 1-127 range for small values
*     if G is 0, H is part of 1-127, if G is 1, H indicates whether the value is 2 or 3 bytes.
*     H is 0 for 2 bytes (effectively able to store 1-16383), 1 for 3 (effectively able to store 1-4194303).
* I = 6 bits available to store a number in the first byte (combined with H in the smallest case)
* J = 8 bits available to store a number in the second byte (total of 14 bits combined with I)
* K = 8 bits available to store a number in the third byte (total of 22 bits combined with J and I)
 */
func DecodeBytes(bs []byte) []int32 {
	ids := []int32{}
	br := bytes.NewReader(bs)
	for br.Len() > 0 {
		b, err := br.ReadByte()
		if err != nil {
			fmt.Println(err)
			break
		}
		if b>>7 == 0 {
			if len(ids) > 0 {
				ids = append(ids, int32(b&0x7F)+ids[len(ids)-1])
			} else {
				ids = append(ids, int32(b&0x7f))
			}
		} else if b>>6 == 2 {
			b2, err := br.ReadByte()
			if err != nil {
				fmt.Println(err)
				break
			}
			b &= 0x3F
			if len(ids) > 0 {
				ids = append(ids, (int32(b)<<8|int32(b2))+ids[len(ids)-1])
			} else {
				ids = append(ids, int32(b)<<8|int32(b2))
			}
		} else if b>>6 == 3 {
			b2, err := br.ReadByte()
			if err != nil {
				fmt.Println(err)
				break
			}
			b3, err := br.ReadByte()
			if err != nil {
				fmt.Println(err)
				break
			}
			b &= 0x3F
			if len(ids) > 0 {
				ids = append(ids, (int32(b)<<16|int32(b2)<<8|int32(b3))+ids[len(ids)-1])
			} else {
				ids = append(ids, int32(b)<<16|int32(b2)<<8|int32(b3))
			}
		}
	}
	//fmt.Printf("encoded(hex): %x\n", bs)
	//	fmt.Println("decoded:", ids)
	return ids
}

type Int32Slice []int32

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func SortInt32s(a []int32)              { sort.Sort(Int32Slice(a)) }
