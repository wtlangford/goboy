// vim: noet:ts=3:sw=3:sts=3
package common

func ByteToInt(v byte) int {
	return int(int8(v))
}

func ByteToUint(v byte) uint {
	return uint(v)
}
