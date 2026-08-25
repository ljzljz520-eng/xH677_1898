package audit

import "strings"

func NormalizeActor(actor string) string   { return strings.TrimSpace(strings.ToLower(actor)) }
func EventKey(batch, action string) string { return batch + ":" + action }
func Actions(actions []string) map[string]int {
	m := map[string]int{}
	for _, a := range actions {
		m[a]++
	}
	return m
}
