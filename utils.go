package bgls

import "encoding/hex"

func Bytestohex(x []byte) string {
	return hex.EncodeToString(x)
}
