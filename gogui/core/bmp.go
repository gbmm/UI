package core

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
)

type BITMAPINFOHEADER struct {
	BiSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

func ReaaBmp(path string) (*BITMAPINFOHEADER, *[]byte) {
	header := BITMAPINFOHEADER{}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	//type拆成两个byte来读
	var headA, headB byte
	//Read第二个参数字节序一般windows/linux大部分都是LittleEndian,苹果系统用BigEndian
	binary.Read(file, binary.LittleEndian, &headA)
	binary.Read(file, binary.LittleEndian, &headB)

	//文件大小
	var size uint32
	binary.Read(file, binary.LittleEndian, &size)

	//预留字节
	var reservedA, reservedB uint16
	binary.Read(file, binary.LittleEndian, &reservedA)
	binary.Read(file, binary.LittleEndian, &reservedB)

	//偏移字节
	var offbits uint32
	binary.Read(file, binary.LittleEndian, &offbits)

	//fmt.Println(headA, headB, size, reservedA, reservedB, offbits)

	infoHeader := &header
	binary.Read(file, binary.LittleEndian, infoHeader)
	data, err := ioutil.ReadAll(file)
	return infoHeader, &data
}
