package main

func rulesFromInt(ruleInt int) []bool {
	ruleset := make([]bool, 8)
	for i := 0; i < 8; i++ {
		ruleset[i] = false
	}

	for i := range ruleset {
		if ruleInt == 0 {
			return ruleset
		}
		rem := ruleInt % 2

		if rem == 1 {
			ruleset[i] = true
			ruleInt--
		}
		ruleInt = ruleInt / 2
	}

	return ruleset
}
