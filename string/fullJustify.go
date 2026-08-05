package strings

import "strings"

// FullJustify
// @Title: LC68.文本左右对齐
// @Description: 给定一个单词数组和一个长度 maxWidth，重新排版单词，使其成为每行恰好有 maxWidth 个字符、且左右两端对齐的文本。
// @Description: 你需要使用贪心算法来放置一个给定背包内的单词，也就是说，在一行中顺序填充尽可能多的单词。
// @Description: 单词的数量可能为 0 个，并且每个单词的长度之和总是小于或等于 maxWidth。
// @Description: 若单词之间的空格数不够均匀，左侧的空格会多于右侧。
// @Description: 最后一行应当是左对齐的，并且单词之间不应插入额外的空格。
// @Link: https://leetcode.cn/problems/text-justification/
// @param words
// @param maxWidth
// @return []string
func FullJustify(words []string, maxWidth int) []string {
	if len(words) == 0 {
		return []string{}
	}

	var (
		result  = make([]string, 0, 4)
		line    = make([]string, 0, 4)
		lineLen = 0
	)

	// 贪心逐行填充: lineLen 为当前行字符总长度, len(line) 为间隔数
	// 当 lineLen + len(line) + wordLen > maxWidth 时, 说明放入当前单词后每行间隔数 = 单词数-1,
	// 且每个间隔至少 1 个空格, 故溢出判定为字符总长 + 间隔数 > maxWidth
	for i := 0; i < len(words); i++ {
		wordLen := len(words[i])
		if lineLen+len(line)+wordLen > maxWidth {
			result = append(result, formatLine(line, lineLen, maxWidth))
			line = line[:0]
			lineLen = 0
		}
		line = append(line, words[i])
		lineLen += wordLen
	}

	// 最后一行: 左对齐, 单词间单空格, 尾部补空格至 maxWidth
	lastLine := ""
	for i, w := range line {
		if i > 0 {
			lastLine += " "
		}
		lastLine += w
	}
	for len(lastLine) < maxWidth {
		lastLine += " "
	}
	result = append(result, lastLine)

	return result
}

// formatLine 将一行单词两端对齐
// 例: line=["This","is","an"], lineLen=9, maxWidth=16
// totalSpaces=7, gaps=2, base=3, extra=1
// 间隔分配: 第1个间隔 3+1=4 空格, 第2个间隔 3 空格
// 结果: "This    is    an" (4+3=7 个空格, 总长 9+7=16)
func formatLine(line []string, lineLen, maxWidth int) string {
	// 单单词行: 左对齐, 尾部补空格
	if len(line) == 1 {
		res := line[0]
		for len(res) < maxWidth {
			res += " "
		}
		return res
	}

	totalSpaces := maxWidth - lineLen
	gaps := len(line) - 1
	base := totalSpaces / gaps
	extra := totalSpaces % gaps

	// 前 extra 个间隔各多分 1 个空格, 保证左多右少
	res := ""
	for i, w := range line {
		res += w
		if i < len(line)-1 {
			spaces := base
			if i < extra {
				spaces++
			}
			for j := 0; j < spaces; j++ {
				res += " "
			}
		}
	}
	return res
}

// FullJustifyV2
// @Title: LC68.文本左右对齐
// @Description: 仿照 C++ 写法: while 循环 + begin 索引, 单函数集中处理所有行的排版
// @Link: https://leetcode.cn/problems/text-justification/
// @solution: https://leetcode.cn/problems/text-justification/solutions/3618368/luo-ji-jian-dan-yi-jie-fa-yong-shi-ji-ba-3g1j/
// @param words
// @param maxWidth
// @return []string
func FullJustifyV2(words []string, maxWidth int) []string {
	var (
		n   = len(words)
		i   = 0
		res = make([]string, 0, 4)
	)

	for i < n {
		begin := i
		totalChars := 0

		// 当前行可容纳的单词数: 字符总长 + 间隔数(=单词数-1) <= maxWidth
		// 例: words=["This","is","an"], maxWidth=16
		// i=0: totalChars=4, 4+0+4=8<=16
		// i=1: totalChars=6, 6+1+2=9<=16
		// i=2: totalChars=8, 8+2+2=12<=16
		// i=3: totalChars=13, 13+3+4=20>16 -> 停止
		for i < n && totalChars+len(words[i])+(i-begin) <= maxWidth {
			totalChars += len(words[i])
			i++
		}

		// 最后一行 或 单单词行: 左对齐, 尾部补空格
		if i == n || i-begin == 1 {
			line := words[begin]
			for j := begin + 1; j < i; j++ {
				line += " " + words[j]
			}
			line += strings.Repeat(" ", maxWidth-len(line))
			res = append(res, line)
			continue
		}

		// 两端对齐: (i-begin) 个单词, (i-begin-1) 个间隔, (maxWidth-totalChars) 个空格
		// base = 空格总数 / 间隔数, extra = 空格总数 % 间隔数
		// 前 extra 个间隔各多 1 个空格 (左多右少)
		gaps := i - begin - 1
		spaces := maxWidth - totalChars
		base := spaces / gaps
		extra := spaces % gaps

		// 例: line=["This","is","an"], totalChars=9, maxWidth=16
		// spaces=7, gaps=2, base=3, extra=1
		// j=1: 3 + (1 <= 1 ? 1 : 0) = 4
		// j=2: 3 + (2 <= 1 ? 1 : 0) = 3
		// 结果: "This    is    an" (4+3=7 空格, 总长 9+7=16)
		line := words[begin]
		for j := begin + 1; j < i; j++ {
			need := base
			if j-begin <= extra {
				need++
			}
			line += strings.Repeat(" ", need)
			line += words[j]
		}
		res = append(res, line)
	}

	return res
}
