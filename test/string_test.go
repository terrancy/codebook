package test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	strings "terrancy/awesome/string"
)

var kmpCases = []struct {
	name     string
	target   string
	patten   string
	expected int
}{
	{
		name:     "test_case_1",
		target:   "abacabab",
		patten:   "abab",
		expected: 1,
	},
	{
		name:     "test_case_2",
		target:   "aaaaa",
		patten:   "aa",
		expected: 4,
	},
	{
		name:     "test_case_3",
		target:   "abc",
		patten:   "d",
		expected: 0,
	},
}

// TestKMP NC149.kmp算法
func TestKMP(t *testing.T) {
	for _, tt := range kmpCases {
		t.Run(tt.name, func(t *testing.T) {
			res := strings.KMP(tt.target, tt.patten)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var permutationCases = []struct {
	name     string
	str      string
	expected []string
}{
	{
		name:     "test_case_1",
		str:      "aab",
		expected: []string{"aab", "aba", "baa"},
	},
	{
		name:     "test_case_2",
		str:      "abc",
		expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
	},
}

// TestPermutation NC121.字符串的排列
func TestPermutation(t *testing.T) {
	for _, tt := range permutationCases {
		t.Run(tt.name, func(t *testing.T) {
			res := strings.Permutation(tt.str)
			sort.Strings(res)
			sort.Strings(tt.expected)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var longestPalindromeCases = []struct {
	name     string
	str      string
	expected int
}{
	{
		name:     "test_case_1",
		str:      "ccbcabaabba",
		expected: 4,
	},
	{
		name:     "test_case_2",
		str:      "babad",
		expected: 3,
	},
	{
		name:     "test_case_3",
		str:      "cbbd",
		expected: 2,
	},
	{
		name:     "test_case_4",
		str:      "a",
		expected: 1,
	},
}

// TestLongestPalindrome NC17.最长回文子串
func TestLongestPalindrome(t *testing.T) {
	for _, tt := range longestPalindromeCases {
		t.Run(tt.name, func(t *testing.T) {
			res := strings.LongestPalindrome(tt.str)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var longRepeatedSubstringCases = []struct {
	name     string
	str      string
	expected int
}{
	{
		name:     "test_case_1",
		str:      "abcab",
		expected: 0,
	},
	{
		name:     "test_case_2",
		str:      "aaaa",
		expected: 4,
	},
	{
		name:     "test_case_3",
		str:      "abab",
		expected: 4,
	},
	{
		name:     "test_case_4",
		str:      "a",
		expected: 1,
	},
}

// TestLongRepeatedSubstring NC142.最长重复子串
func TestLongRepeatedSubstring(t *testing.T) {
	for _, tt := range longRepeatedSubstringCases {
		t.Run(tt.name, func(t *testing.T) {
			res := strings.LongRepeatedSubstring(tt.str)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var firstNotRepeatingCharCases = []struct {
	name     string
	str      string
	expected int
}{
	{
		name:     "test_case_1",
		str:      "google",
		expected: 4,
	},
	{
		name:     "test_case_2",
		str:      "abaccdeff",
		expected: 1,
	},
	{
		name:     "test_case_3",
		str:      "aabb",
		expected: -1,
	},
}

// TestFirstNotRepeatingChar NC31.第一个只出现一次的字符
func TestFirstNotRepeatingChar(t *testing.T) {
	for _, tt := range firstNotRepeatingCharCases {
		t.Run(tt.name, func(t *testing.T) {
			res := strings.FirstNotRepeatingChar(tt.str)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var checkIpAddressCases = []struct {
	name     string
	ip       string
	expected string
}{
	{
		name:     "test_case_1",
		ip:       "127.0.0.1",
		expected: "IPv4",
	},
	{
		name:     "test_case_2",
		ip:       "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
		expected: "IPv6",
	},
	{
		name:     "test_case_3",
		ip:       "256.256.256.256",
		expected: "Neither",
	},
}

// TestCheckIpAddress NC113.验证IP地址
func TestCheckIpAddress(t *testing.T) {
	for _, tt := range checkIpAddressCases {
		t.Run(tt.name, func(t *testing.T) {
			res := strings.CheckIpAddress(tt.ip)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var reverseSentenceCases = []struct {
	name     string
	str      string
	expected string
}{
	{
		name:     "test_case_1",
		str:      "nowcoder. a am I",
		expected: "I am a nowcoder.",
	},
	{
		name:     "test_case_2",
		str:      "",
		expected: "",
	},
	{
		name:     "test_case_3",
		str:      "hello world",
		expected: "world hello",
	},
}

// TestReverseSentence JZ73.翻转单词序列
func TestReverseSentence(t *testing.T) {
	for _, tt := range reverseSentenceCases {
		t.Run(tt.name, func(t *testing.T) {
			res := strings.ReverseSentence(tt.str)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var generateParenthesisCases = []struct {
	name     string
	n        int
	expected []string
}{
	{
		name:     "test_case_1",
		n:        3,
		expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
	},
	{
		name:     "test_case_2",
		n:        1,
		expected: []string{"()"},
	},
}

// TestGenerateParenthesis NC26.括号生成
func TestGenerateParenthesis(t *testing.T) {
	for _, tt := range generateParenthesisCases {
		t.Run(tt.name, func(t *testing.T) {
			res := strings.GenerateParenthesis(tt.n)
			sort.Strings(res)
			sort.Strings(tt.expected)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var checkBracketValidCases = []struct {
	name     string
	str      string
	expected bool
}{
	{
		name:     "test_case_1",
		str:      "()[]{}",
		expected: true,
	},
	{
		name:     "test_case_2",
		str:      "(]",
		expected: false,
	},
	{
		name:     "test_case_3",
		str:      "([)]",
		expected: false,
	},
	{
		name:     "test_case_4",
		str:      "{[]}",
		expected: true,
	},
}

// TestCheckBracketValid LC20.有效的括号
func TestCheckBracketValid(t *testing.T) {
	for _, tt := range checkBracketValidCases {
		t.Run(tt.name, func(t *testing.T) {
			res := strings.CheckBracketValid(tt.str)
			assert.Equal(t, tt.expected, res)
		})
	}
}

var fullJustifyCases = []struct {
	name     string
	words    []string
	maxWidth int
	expected []string
}{
	{
		name:     "test_case_1",
		words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
		maxWidth: 16,
		expected: []string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		},
	},
	{
		name:     "test_case_2",
		words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
		maxWidth: 16,
		expected: []string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		},
	},
	{
		name:     "test_case_3",
		words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
		maxWidth: 20,
		expected: []string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		},
	},
}

// TestFullJustify LC68.文本左右对齐
func TestFullJustify(t *testing.T) {
	for _, tt := range fullJustifyCases {
		t.Run(tt.name, func(t *testing.T) {
			res := strings.FullJustify(tt.words, tt.maxWidth)
			assert.Equal(t, tt.expected, res)
		})
	}
}

// TestFullJustifyV2 LC68.文本左右对齐(仿照 C++ 写法)
func TestFullJustifyV2(t *testing.T) {
	for _, tt := range fullJustifyCases {
		t.Run(tt.name, func(t *testing.T) {
			res := strings.FullJustifyV2(tt.words, tt.maxWidth)
			assert.Equal(t, tt.expected, res)
		})
	}
}