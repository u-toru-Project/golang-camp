package chapter16

import (
	"fmt"
	"strings"
)

var xmlEscapes = []struct {
	Raw    rune
	Entity string
}{
	{'&', "&amp;"},
	{'"', "&quot;"},
	{'\'', "&apos;"},
	{'<', "&lt;"},
	{'>', "&gt;"},
}

func EncodeXml(text string) string {
	var builder strings.Builder
	for _, ch := range text {
		encoded := false
		for _, pair := range xmlEscapes {
			if ch == pair.Raw {
				builder.WriteString(pair.Entity)
				encoded = true
				break
			}
		}
		if !encoded {
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}

func DecodeXml(text string) string {
	decode := map[string]rune{}
	for _, pair := range xmlEscapes {
		decode[pair.Entity] = pair.Raw
	}
	var builder strings.Builder
	for i := 0; i < len(text); i++ {
		if text[i] != '&' {
			builder.WriteByte(text[i])
			continue
		}
		end := strings.IndexByte(text[i:], ';')
		if end < 0 {
			panic("Invalid XML entity.")
		}
		end += i
		entity := text[i : end+1]
		raw, ok := decode[entity]
		if !ok {
			panic("Unknown entity '" + entity + "'.")
		}
		builder.WriteRune(raw)
		i = end
	}
	return builder.String()
}

func RunQ1612() {
	raw := "<tag attr=\"x\">"
	encoded := EncodeXml(raw)
	fmt.Printf("%s -> %s -> %s\n", raw, encoded, DecodeXml(encoded))
}
