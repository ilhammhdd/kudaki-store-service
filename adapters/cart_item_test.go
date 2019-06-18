package adapters_test

import (
	"strings"
	"testing"
)

type RedisearchText string

func (rt *RedisearchText) Set(s string) {
	*rt = RedisearchText(s)
}

func (rt *RedisearchText) Sanitize() string {
	replacer := strings.NewReplacer(`,`, `\,`, `.`, `\.`, `<`, `\<`, `>`, `\>`, `{`, `\{`, `}`, `\}`, `[`, `\[`, `]`, `\]`, `"`, `\"`, `'`, `\'`, `:`, `\:`, `;`, `\;`, `!`, `\!`, `@`, `\@`, `#`, `\#`, `$`, `\$`, `%`, `\%`, `^`, `\^`, `&`, `\&`, `*`, `\*`, `(`, `\(`, `)`, `\)`, `-`, `\-`, `+`, `\+`, `=`, `\=`, `~`, `\~`)
	return replacer.Replace(string(*rt))
}

func (rt *RedisearchText) UnSanitize() string {
	replacer := strings.NewReplacer(`\,`, `,`, `\.`, `.`, `\<`, `<`, `\>`, `>`, `\{`, `{`, `\}`, `}`, `\[`, `[`, `\]`, `]`, `\"`, `"`, `\'`, `'`, `\:`, `:`, `\;`, `;`, `\!`, `!`, `\@`, `@`, `\#`, `#`, `\$`, `$`, `\%`, `%`, `\^`, `^`, `\&`, `&`, `\*`, `*`, `\(`, `(`, `\)`, `)`, `\-`, `-`, `\+`, `+`, `\=`, `=`, `\~`, `~`)
	return replacer.Replace(string(*rt))
}

type Sanitizer interface {
	Set(string)
	Sanitize() string
	UnSanitize() string
}

func TestSanitizer(t *testing.T) {
	doSanitize(new(RedisearchText), t)
}

func doSanitize(sn Sanitizer, t *testing.T) {
	sn.Set("a3d8e979-bc68-4060-82e3-e74d2bb9ccc6")
	t.Log(sn.UnSanitize())
}
