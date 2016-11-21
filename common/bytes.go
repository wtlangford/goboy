// vim: noet:ts=3:sw=3:sts=3
package common

func ByteToInt(v byte) int {
	return int(int8(v))
}

func ByteToUint(v byte) uint {
	return uint(v)
}

func HasFlagSet(v byte, f byte) bool {
	return v&f == f
}

func HasFlagsSet(v byte, flags ...byte) bool {
	var f byte
	for _, flg := range flags {
		f |= flg
	}
	return v&f == f
}
