package public

const (
	SET = 1
	GET = 2
)

type GetReq struct {
	Key string
}

type GetResp struct {
	Value string
}

type SetReq struct {
	Key   string
	Value string
}

// uint32转bytes
func Uint32ToBytes(n uint32) []byte {
	return []byte{byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)}
}

// bytes转uint32
func BytesToUint32(b []byte) uint32 {
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
}

// 编码数据
func Encode(tag uint32, data []byte) []byte {
	resp := make([]byte, 8+len(data))
	l := Uint32ToBytes(uint32(len(data)))
	t := Uint32ToBytes(tag)
	copy(resp, l)
	copy(resp[4:], t)
	copy(resp[8:], data)
	return resp
}

// 解码数据
func Uncode(data []byte) (uint32, uint32) {
	l := BytesToUint32(data[:4])
	t := BytesToUint32(data[4:8])
	return l, t
}
