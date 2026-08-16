package unescape

import "strings"

// Unescape 把常见的转义序列还原成字面字符：
//
//	\n \t \r \\ \" \' \a \b \f \v
//	\xNN（两位十六进制）和 \0NNN（最多三位八进制）
//
// 不认识的转义原样保留反斜杠。
func Unescape(s string) string {
	var b strings.Builder
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] != '\\' || i+1 >= len(runes) {
			b.WriteRune(runes[i])
			continue
		}
		nxt := runes[i+1]
		switch nxt {
		case 'n':
			b.WriteRune('\n')
			i++
		case 't':
			b.WriteRune('\t')
			i++
		case 'r':
			b.WriteRune('\r')
			i++
		case '\\':
			b.WriteRune('\\')
			i++
		case '"':
			b.WriteRune('"')
			i++
		case '\'':
			b.WriteRune('\'')
			i++
		case 'a':
			b.WriteRune('\a')
			i++
		case 'b':
			b.WriteRune('\b')
			i++
		case 'f':
			b.WriteRune('\f')
			i++
		case 'v':
			b.WriteRune('\v')
			i++
		case 'x':
			// \xNN
			if i+3 < len(runes) {
				if v, ok := parseHex(runes[i+2], runes[i+3]); ok {
					b.WriteByte(byte(v))
					i += 3
					continue
				}
			}
			b.WriteRune('\\')
		case '0', '1', '2', '3', '4', '5', '6', '7':
			// \NNN / \0NNN 最多三位八进制，和 Go 字符串转义一致（\101 即 'A'）
			j := i + 1
			val := 0
			cnt := 0
			for cnt < 3 && j < len(runes) && runes[j] >= '0' && runes[j] <= '7' {
				val = val*8 + int(runes[j]-'0')
				j++
				cnt++
			}
			if cnt > 0 {
				b.WriteByte(byte(val))
				i = j - 1
				continue
			}
			b.WriteRune('\\')
		default:
			b.WriteRune('\\')
		}
	}
	return b.String()
}

func parseHex(a, b rune) (byte, bool) {
	hi, ok1 := hexVal(a)
	lo, ok2 := hexVal(b)
	if !ok1 || !ok2 {
		return 0, false
	}
	return byte(hi*16 + lo), true
}

func hexVal(r rune) (byte, bool) {
	switch {
	case r >= '0' && r <= '9':
		return byte(r - '0'), true
	case r >= 'a' && r <= 'f':
		return byte(r - 'a' + 10), true
	case r >= 'A' && r <= 'F':
		return byte(r - 'A' + 10), true
	}
	return 0, false
}
