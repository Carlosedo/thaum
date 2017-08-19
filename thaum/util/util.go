package util

import modifiers "github.com/Flaque/thaum/thaum/modifiers"

func Keys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func AddStringsToSet(strs []string, set map[string]string) map[string]string {
	for _, s := range strs {
        // only add variable names, without the modifier
		set[modifiers.GetKey(s)] = ""
	}
	return set
}
