package util

const (
	// DELTA int32 = -0x61C88647
	// SUM   int32 = -0x3910C8E0
	DELTA int32 = -1640531527
	SUM   int32 = -957401312
	ZERO  int32 = 0
)

func Encrypt32(info []byte, key []int32) []byte {
	n := 8 - len(info)%8
	encryptStr := make([]byte, len(info)+n)
	encryptStr[0] = byte(n)
	copy(encryptStr[n:], info[0:])
	result := make([]byte, len(encryptStr))
	for i := 0; i < len(result); i += 8 {
		tempEncrypt := Encrypt(encryptStr, i, key, 32)
		copy(result[i:], tempEncrypt[0:8])
	}
	return result
}

func Decrypt32(secretInfo []byte, key []int32) []byte {
	var decryptStr []byte
	length := len(secretInfo)
	tempDecrypt := make([]byte, length)
	for i := 0; i < length; i += 8 {
		decryptStr = Decrypt(secretInfo, i, key, 32)
		copy(tempDecrypt[i:], decryptStr[0:8])
	}
	if len(decryptStr) > 0 {
		n := tempDecrypt[0]
		return tempDecrypt[n:len(decryptStr)]
	} else {
		return nil
	}

}

func Encrypt(content []byte, offset int, key []int32, times int) []byte {
	tempInt := byteToInt(content, offset)
	y, z, sum, delta := tempInt[0], tempInt[1], ZERO, DELTA
	a, b, c, d := key[0], key[1], key[2], key[3]
	for i := 0; i < times; i++ {
		sum += delta
		y += ((z << 4) + a) ^ (z + sum) ^ ((z >> 5) + b)
		z += ((y << 4) + c) ^ (y + sum) ^ ((y >> 5) + d)
	}
	tempInt[0] = y
	tempInt[1] = z
	return intToByte(tempInt, 0)
}

func Decrypt(encryptContent []byte, offset int, key []int32, times int) []byte {
	tempInt := byteToInt(encryptContent, offset)
	y, z, sum, delta := tempInt[0], tempInt[1], SUM, DELTA
	a, b, c, d := key[0], key[1], key[2], key[3]

	for i := 0; i < times; i++ {
		z -= ((y << 4) + c) ^ (y + sum) ^ ((y >> 5) + d)
		y -= ((z << 4) + a) ^ (z + sum) ^ ((z >> 5) + b)
		sum -= delta
	}
	tempInt[0] = y
	tempInt[1] = z

	return intToByte(tempInt, 0)
}

func byteToInt(content []byte, offset int) []int32 {
	result := make([]int32, len(content)>>2)
	length := len(content)
	for i, j := 0, offset; j < length; i, j = i+1, j+4 {
		result[i] = transform(content[j+3]) | transform(content[j+2])<<8 |
			transform(content[j+1])<<16 | int32(content[j])<<24
	}
	return result
}

func intToByte(content []int32, offset int) []byte {
	result := make([]byte, len(content)<<2)
	length := len(result)
	for i, j := 0, offset; j < length; i, j = i+1, j+4 {
		result[j+3] = byte(content[i] & 0xff)
		result[j+2] = byte((content[i] >> 8) & 0xff)
		result[j+1] = byte(((content[i] >> 16) & 0xff))
		result[j] = byte(((content[i] >> 24) & 0xff))
	}
	return result
}

func transform(data byte) int32 {
	temp := int32(data)
	if temp < 0 {
		temp += 256
	}
	return temp
}
