func myAtoi(s string) int {
    res := strings.TrimSpace(s)
	if len(res) == 0 {
		return 0
	}

	var flag bool

	if res[0] == '-' {
		flag = true
		res = res[1:]
	} else if res[0] == '+' {
		res = res[1:]
	}

	var rees int
	for _, tmp := range res {
		if tmp < '0' || tmp > '9' {
			break
		}
		rees = rees*10 + int(tmp-'0')
		if rees > 2147483647 && !flag {
			return 2147483647
		}

		if -rees < -2147483648 && flag {
			return -2147483648
		}
	}
	if flag {
		rees = -rees
	}
	return rees
}
