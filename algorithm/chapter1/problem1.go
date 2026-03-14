package chapter1

import (
	"fmt"
)

func IsUniqueCheck(input string) bool {
	fmt.Printf("チェック開始: %q === \n", input)
	seen := make(map[rune]struct{})

	for i, r := range input {
		fmt.Printf("[%d文字目] 現在の文字(r): '%c', 記録(seen): %v\n", i+1, r, seen)
		if _, ok := seen[r]; ok {
			fmt.Printf("重複が発見されました。 '%c'はすでに記録に登録されています。処理を終了します。\n\n", r)
			return false
		}

		seen[r] = struct{}{}
		fmt.Printf(" %c を記録に登録しました", r)
	}
	fmt.Printf("全文字のチェック済みです。重複はありません。\n")
	return true
}

func main() {
	IsUniqueCheck("dog")
	IsUniqueCheck("aba")
}
