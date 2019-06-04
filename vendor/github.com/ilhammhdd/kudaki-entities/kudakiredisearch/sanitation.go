package kudakiredisearch

import "strings"

type RedisearchText string

func (rt RedisearchText) Sanitize() string {
	replacer := strings.NewReplacer(`,`, `\,`, `.`, `\.`, `<`, `\<`, `>`, `\>`, `{`, `\{`, `}`, `\}`, `[`, `\[`, `]`, `\]`, `"`, `\"`, `'`, `\'`, `:`, `\:`, `;`, `\;`, `!`, `\!`, `@`, `\@`, `#`, `\#`, `$`, `\$`, `%`, `\%`, `^`, `\^`, `&`, `\&`, `*`, `\*`, `(`, `\(`, `)`, `\)`, `-`, `\-`, `+`, `\+`, `=`, `\=`, `~`, `\~`)
	return replacer.Replace(string(rt))
}
