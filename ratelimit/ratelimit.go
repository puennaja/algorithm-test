package ratelimit

var (
	success string = "{status: 200, message: OK}"
	fail    string = "{status: 429, message: Too many requests}"
)

type RateLimiter struct {
	datas  []string
	limit  int
	lenght int
}

func (rl *RateLimiter) appendData(data string) {
	rl.datas = append(rl.datas, data)
}

func (rl *RateLimiter) pop() {
	if len(rl.datas) > rl.lenght {
		rl.datas = rl.datas[1:]
	}
}

func (rl *RateLimiter) isBreakLimit(data string) bool {
	count := 0
	rl.pop()
	for _, v := range rl.datas {
		if v == data {
			count++
		}
		if count > rl.limit {
			rl.appendData("fail")
			return true
		}
	}
	rl.appendData(data)
	return false
}

func NewRateLimiter(limit, lenght int) *RateLimiter {
	return &RateLimiter{
		datas:  []string{},
		limit:  limit - 1,
		lenght: lenght - 1,
	}
}

func RequestLimiter(requests []string) []string {
	rateLimit5 := NewRateLimiter(2, 5)
	rateLimit30 := NewRateLimiter(5, 30)
	out := []string{}
	for _, req := range requests {
		is5Fail := rateLimit5.isBreakLimit(req)
		is30Fail := rateLimit30.isBreakLimit(req)
		if is5Fail || is30Fail {
			out = append(out, fail)
		} else {
			out = append(out, success)
		}
	}
	return out
}
