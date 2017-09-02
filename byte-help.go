package Figo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"strconv"
)

var Bh  = ByteHelp{}

type ByteHelp struct {

}

func (p *ByteHelp) I162B(n int16) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, &n)
	return bytesBuffer.Bytes()
}

func (p *ByteHelp) B2I16(b []byte) int16 {
	bytesBuffer := bytes.NewBuffer(b)
	log.Println(b)
	var v int16
	binary.Read(bytesBuffer, binary.BigEndian, &v)
	return int16(v)
}

func (p *ByteHelp) BStr(bs []byte)string{
	v:=""
	for  _,b :=range bs  {
		s:=fmt.Sprint(uint8(b))
		if len(s)<2{
			s=fmt.Sprint("00",s)
		}else if len(s)<3{
			s=fmt.Sprint("0",s)
		}
		v = fmt.Sprint(v," ",s)
	}
	return v
}

func (p *ByteHelp) Append(bss ...[]byte)[]byte {
	out := []byte{}
	for _,bs := range bss {
		for _,b:=range bs{
			out = append(out,b)
		}
	}
	return out
}

func (p *ByteHelp) ToHex(bs []byte)string {
	buffer := new(bytes.Buffer)
	for _, b := range bs {
		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}
	return buffer.String()
}


func (p *ByteHelp) FromHex(hex string)[]byte {
	length := len(hex) / 2
	slice := make([]byte, length)
	rs := []rune(hex)
	for i := 0; i < length; i++ {
		s := string(rs[i*2 : i*2+2])
		value, _ := strconv.ParseInt(s, 16, 10)
		slice[i] = byte(value & 0xFF)
	}
	return slice
}