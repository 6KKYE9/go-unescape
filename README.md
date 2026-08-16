# go-unescape

把字符串里的转义序列还原成真字符。常见那套 `\n \t \r \\ \" \' \a \b \f \v` 都管，再顺手支持 `\xNN`（十六进制字节）和 `\0NNN`（八进制，最多三位）。不认识的转义（比如 `\q`）就老老实实保留反斜杠，不瞎猜。

## 装

```bash
go build -o unescape ./cmd/unescape
```

## 用

```bash
echo 'line1\nline2\tend' | ./unescape
# line1
# line2    end

echo '\x41\x42' | ./unescape     # AB
```

## 当库用

```go
import "unescape"

unescape.Unescape(`a\nb`)   // "a\nb" (含真实换行)
unescape.Unescape(`\x41`)   // "A"
```

注意：八进制最多吃三位，多出来的数字会留在原串里，不会吞掉。

## License

MIT
