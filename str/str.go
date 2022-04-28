// Package str
// @Description:
package str

import (
	"html"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// E
// @Description: Encode HTML special characters in a string.
// @param value
// @return string
func E(value string) string {
	return html.EscapeString(value)
}

// RegularReplaceArray
// @Description: Replace a given pattern with each value in the array in sequentially.
// @param pattern
// @param replacements
// @param subject
// @return string
func RegularReplaceArray(pattern string, replacements []string, subject string) string {
	reg, _ := regexp.Compile(pattern)
	index := 0
	return reg.ReplaceAllStringFunc(subject, func(s string) string {
		if index < len(replacements) {
			response := replacements[index]
			index++
			return response
		}
		return s
	})
}

// After
// @Description:Return the remainder of a string after the first occurrence of a given value.
// @param subject
// @param search
// @return string
func After(subject string, search string) string {
	if search == "" {
		return subject
	}
	index := strings.Index(subject, search)
	if index == -1 {
		return subject
	}
	return strings.TrimLeft(subject[index+len(search):], " ")
}

// AfterLast
// @Description:Return the remainder of a string after the last occurrence of a given value.
// @param subject
// @param search
// @return string
func AfterLast(subject string, search string) string {
	if search == "" {
		return subject
	}
	index := strings.LastIndex(subject, search)
	if index == -1 {
		return subject
	}
	return strings.TrimLeft(subject[index+len(search):], " ")
}

// Before
// @Description:Get the portion of a string before the first occurrence of a given value.
// @param subject
// @param search
// @return string
func Before(subject string, search string) string {
	if search == "" {
		return subject
	}
	index := strings.Index(subject, search)
	if index == -1 {
		return subject
	}
	return subject[:index]
}

// BeforeLast
// @Description:Get the portion of a string before the last occurrence of a given value.
// @param subject
// @param search
// @return string
func BeforeLast(subject string, search string) string {
	if search == "" {
		return subject
	}
	index := strings.LastIndex(subject, search)
	if index == -1 {
		return subject
	}
	return subject[:index]
}

// Between
// @Description:Get the portion of a string between two given values.
// @param subject
// @param from
// @param to
// @return string
func Between(subject string, from string, to string) string {
	if from == "" || to == "" {
		return subject
	}
	return BeforeLast(After(subject, from), to)
}

// BetweenFirst
// @Description:Get the smallest possible portion of a string between two given values.
// @param subject
// @param from
// @param to
// @return string
func BetweenFirst(subject string, from string, to string) string {
	if from == "" || to == "" {
		return subject
	}
	return Before(After(subject, from), to)
}

// Replace
// @Description:the given value in the given string
// @param search
// @param replace
// @param subject
func Replace(search string, replace string, subject string) string {
	return strings.Replace(subject, search, replace, -1)
}

// ReplaceOfArraySearch
// @Description:replace  the given array value in the given string.
// @param search
// @param replace
// @param subject
// @return response
func ReplaceOfArraySearch(search []string, replace string, subject string) (response string) {
	response = subject
	for _, s := range search {
		response = strings.Replace(response, s, replace, -1)
	}
	return response
}

// ReplaceFirst
// @Description:Replace the first occurrence of a given value in the string.
// @param search
// @param replace
// @param subject
// @return string
func ReplaceFirst(search string, replace string, subject string) string {
	return strings.Replace(subject, search, replace, 1)
}

// ReplaceArray
// @Description: Replace a given value in the string sequentially with an array.
// @param search
// @param replace
// @param subject
// @return string
func ReplaceArray(search string, replace []string, subject string) string {
	for _, item := range replace {
		subject = ReplaceFirst(search, item, subject)
	}
	return subject
}

// ReplaceLast
// @Description:Replace the last occurrence of a given value in the string.
// @param search
// @param replace
// @param subject
// @return string
func ReplaceLast(search string, replace string, subject string) string {
	if search == "" {
		return subject
	}
	index := strings.LastIndex(subject, search)
	if index == -1 {
		return subject
	}
	return subject[:index] + replace + subject[index+len(search):]
}

// Substr
// @Description:Returns the portion of the string specified by the start and length parameters.
// @param subject
// @param start
// @param length
// @return string
func Substr(subject string, start int, length int) string {
	if start >= len(subject) {
		return ""
	}
	end := start + length
	if end >= len(subject) {
		end = len(subject)
	}
	return subject[start:end]
}

// SubstrCount
// @Description: Returns the number of substring occurrences.
// @param haystack
// @param needle
// @return int
func SubstrCount(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	return strings.Count(haystack, needle)
}

// SubstrReplace
// @Description:Replace text within a portion of a string.
// @param subject
// @param replace
// @param offset
// @param length
// @return string
func SubstrReplace(subject string, replace string, offset int, length int) string {
	l := len(subject)
	if offset < 0 {
		offset += l
	}
	if offset > l {
		offset = l
	} else if offset < 0 {
		offset = 0
	}
	start := subject[:offset]
	end := subject[offset:]
	l = len(end)
	if length < 0 {
		length += l
	}
	if length > l {
		length = l
	} else if length < 0 {
		length = 0
	}
	return start + replace + end[length:]
}

// Contains
// @Description:Determine if a given string contains a given substring.
// @param haystack
// @param needles
// @param ignoreCase
// @return bool
func Contains(haystack string, needles string, ignoreCase bool) bool {
	if ignoreCase {
		haystack = strings.ToLower(haystack)
		needles = strings.ToLower(needles)
	}
	if needles == "" {
		return false
	}
	index := strings.Index(haystack, needles)
	if index < 0 {
		return false
	}
	return true
}

// ContainsAny
// @Description:Determine if a given string contains any one in array values.
// @param haystack
// @param needles
// @param ignoreCase
// @return bool
func ContainsAny(haystack string, needles []string, ignoreCase bool) bool {
	for _, needle := range needles {
		if Contains(haystack, needle, ignoreCase) {
			return true
		}
	}
	return false
}

// ContainsAll
// @Description:Determine if a given string contains all array values.
// @param haystack
// @param needles
// @param ignoreCase
// @return bool
func ContainsAll(haystack string, needles []string, ignoreCase bool) bool {
	for _, needle := range needles {
		if !Contains(haystack, needle, ignoreCase) {
			return false
		}
	}
	return true
}

// EndsWith
// @Description: Determine if a given string ends with a given substring.
// @param haystack
// @param needles
// @return bool
func EndsWith(haystack string, needles string) bool {
	if haystack == "" || needles == "" {
		return false
	}
	index := strings.Index(haystack, needles)
	if index > 0 && index+len(needles) == len(haystack) {
		return true
	}
	return false
}

// Length
// @Description:
// @param value
// @return int
func Length(value string) int {
	return len(value)
}

// Is
// @Description:Determine if a given string matches a given pattern.
// @param pattern
// @param value
// @return bool
func Is(pattern string, value string) bool {
	if pattern == "" {
		return false
	}
	if pattern == value {
		return true
	}
	pattern = regexp.QuoteMeta(pattern)
	pattern = strings.ReplaceAll(pattern, "\\*", ".*")
	match, _ := regexp.MatchString("^"+pattern+"$", value)
	return match
}

// Finish
// @Description:Cap a string with a single instance of a given value.
// @param value
// @param cap
// @return string
func Finish(value string, cap string) string {
	quoted := regexp.QuoteMeta(cap)
	reg, _ := regexp.Compile("(?:" + quoted + ")+$")
	response := reg.ReplaceAllString(value, "") + cap
	return response
}

// Limit
// @Description: the number of characters in a string.
// @param value
// @param limit
// @param end
// @return string
func Limit(value string, limit int, end string) string {
	if len(value) <= limit {
		return value
	}
	return value[:limit] + end
}

// Lower
// @Description:the given string to lower-case.
// @param value
// @return string
func Lower(value string) string {
	return strings.ToLower(value)
}

// Upper
// @Description:Convert the given string to upper-case.
// @param value
// @return string
func Upper(value string) string {
	return strings.ToUpper(value)
}

// PadBoth
// @Description: Pad both sides of a string with another.
// @param value
// @param length
// @param pad
// @return response
func PadBoth(value string, length int, pad string) (response string) {
	if len(value) >= length {
		return value
	}
	flag := true
	for response = value; len(response) < length; {
		if flag == true {
			response += pad
		} else {
			response = pad + response
		}
		flag = !flag
	}
	if flag {
		response = response[len(response)-length:]
	} else {
		response = response[:length]
	}
	return response
}

// PadLeft
// @Description:Pad the left side of a string with another.
// @param value
// @param length
// @param pad
// @return string
func PadLeft(value string, length int, pad string) string {
	strLen := len(value)
	if strLen >= length {
		return value
	}
	return strings.Repeat(pad, int((length-strLen+1)/len(pad)))[:length-strLen] + value
}

// PadRight
// @Description:Pad the right side of a string with another.
// @param value
// @param length
// @param pad
// @return string
func PadRight(value string, length int, pad string) string {
	strLen := len(value)
	if strLen >= length {
		return value
	}
	return value + strings.Repeat(pad, int((length-strLen+1)/len(pad)))[:length-strLen]
}

// Start
// @Description:Begin a string with a single instance of a given value.
// @param value
// @param prefix
// @return string
func Start(value string, prefix string) string {
	quoted := regexp.QuoteMeta(prefix)
	reg, _ := regexp.Compile("^(?:" + quoted + ")+")
	response := prefix + reg.ReplaceAllString(value, "")
	return response
}

// StartsWith
// @Description:Determine if a given string starts with a given substring.
// @param haystack
// @param needles
// @return bool
func StartsWith(haystack string, needles string) bool {
	if haystack == "" || needles == "" {
		return false
	}
	index := strings.Index(haystack, needles)
	if index == 0 {
		return true
	}
	return false
}

// IsAscii
// @Description:Determine if a given string is 7 bit ASCII.
// @param value
// @return bool
func IsAscii(value string) bool {
	for _, c := range value {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Random
// @Description:Generate a more truly "random" alpha-numeric string.
// @param length
// @return string
func Random(length int) string {
	if length <= 0 {
		return ""
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// UcFirst
// @Description:Make a string's first character uppercase.
// @param value
// @return string
func UcFirst(value string) string {
	return Upper(Substr(value, 0, 1)) + Substr(value, 1, len(value))
}

// LcFirst
// @Description: Make a string's first character lowercase.
// @param value
// @return string
func LcFirst(value string) string {
	return Lower(Substr(value, 0, 1)) + Substr(value, 1, len(value))
}

// Reverse
// @Description:Reverse the given string.
// @param value
// @return string
func Reverse(value string) string {
	rns := []rune(value) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// Swap
// @Description: Swap multiple keywords in a string with other keywords.
// @param swapMap
// @param subject
// @return string
func Swap(swapMap map[string]string, subject string) string {
	for search, replace := range swapMap {
		subject = Replace(search, replace, subject)
	}
	return subject
}

// Remove
// @Description:Remove any occurrence of the given string in the subject.
// @param search
// @param subject
// @param caseSensitive
// @return string
func Remove(search string, subject string, caseSensitive bool) (response string) {
	if caseSensitive {
		return Replace(search, "", subject)
	}
	if search == "" || subject == "" {
		return subject
	}
	sub := Lower(subject)
	search = Lower(search)
	l := len(search)
	lastIndex := 0
	for {
		i := strings.Index(sub, search)
		if i < 0 {
			break
		}
		response += subject[lastIndex : lastIndex+i]

		lastIndex = i + l
		sub = sub[lastIndex:]
	}
	response += sub
	return response
}

// Studly
// @Description: Convert a value to studly caps case.
// @param value
// @return string
func Studly(value string) (response string) {
	if value == "" {
		return value
	}
	value = ReplaceOfArraySearch([]string{"-", "_"}, " ", value)
	for _, word := range strings.Split(value, " ") {
		response += UcFirst(word)
	}
	return response
}

// Camel
// @Description: Convert a value to camel case.
// @param value
// @return response
func Camel(value string) (response string) {
	if value == "" {
		return value
	}
	response = Studly(value)
	return LcFirst(response)
}

// Headline
// @Description: Convert the given string to title case for each word.
// @param value
// @return response
func Headline(value string) (response string) {
	if value == "" {
		return value
	}
	reg, _ := regexp.Compile("([A-Z]+)")
	response = reg.ReplaceAllStringFunc(value, func(s string) string {
		return "_" + s
	})
	response = ReplaceOfArraySearch([]string{"-", "_"}, " ", response)
	var responseArr []string
	for _, word := range strings.Split(response, " ") {
		responseArr = append(responseArr, UcFirst(word))
	}
	return strings.TrimLeft(strings.Join(responseArr, " "), " ")
}

// SnakeOfCustom
// @Description: Convert a string to kebab case.
// @param value
// @param delimiter
// @return response
func SnakeOfCustom(value string, delimiter string) (response string) {
	if value == "" {
		return value
	}

	strBuilder := strings.Builder{}
	for _, w := range value {
		// asset w is A~Z
		if w >= 65 && w <= 90 {
			strBuilder.WriteByte('_')
		}
		strBuilder.WriteRune(w)
	}
	response = strBuilder.String()

	// Revision does not use re
	//reg, _ := regexp.Compile("([A-Z]+)")
	//response = reg.ReplaceAllStringFunc(value, func(s string) string {
	//	return "_" + s
	//})

	response = ReplaceOfArraySearch([]string{"-", "_"}, " ", response)
	var responseArr []string
	for _, word := range strings.Split(response, " ") {
		responseArr = append(responseArr, LcFirst(word))
	}
	return strings.TrimLeft(strings.Join(responseArr, delimiter), " ")
}

// Snake
// @Description:Convert a string to kebab case.
// @param value
// @return response
func Snake(value string) (response string) {
	return SnakeOfCustom(value, "_")
}

// Kebab
// @Description: Convert a string to kebab case.
// @param value
// @return response
func Kebab(value string) (response string) {
	return SnakeOfCustom(value, "-")
}

// Mask
// @Description: Masks a portion of a string with a repeated character.
// @param subject
// @param character
// @param index
// @param length
// @return string
func Mask(subject string, character string, index int, length int) string {
	subLen := len(subject)
	if index < 0 {
		index += subLen
	}
	if index < 0 {
		index = 0
	}
	if index > subLen {
		index = subLen
	}

	var offset = 0
	if length > 0 {
		offset = length + index
	}
	if length < 0 {
		offset = length + subLen
	}
	if offset < index {
		offset = index
	}
	if offset > subLen {
		offset = subLen
	}
	return subject[:index] + PadBoth("", offset-index, "*") + subject[offset:]
}

// Excerpt
// @Description: Extracts an excerpt from text that matches the first instance of a phrase.
// @param text
// @param phrase
// @param options
// @return response
func Excerpt(text string, phrase string, options map[string]string) (response string) {
	index := strings.Index(text, phrase)
	if index < 0 {
		return ""
	}
	radius := 3
	omission := "..."
	if options != nil {
		if r, e := strconv.Atoi(options["radius"]); e == nil && r > 0 {
			radius = r
		}
		if options["omission"] != "" {
			omission = options["omission"]
		}
	}
	if index-radius > 0 {
		response += omission
	}
	if index < radius {
		response += text[:index+len(phrase)]
	} else {
		response += text[index-radius : index+len(phrase)]
	}
	if index+len(phrase)+radius > len(text) {
		return response + text[index+len(phrase):]
	} else {
		return response + text[index+len(phrase):index+len(phrase)+radius] + omission
	}
}
