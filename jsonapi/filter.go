package jsonapi

import (
	"errors"
	"regexp"
	"strings"
)

const (
	FilterEq   = "eq"
	FilterNe   = "ne"
	FilterLt   = "lt"
	FilterGt   = "gt"
	FilterLe   = "le"
	FilterGe   = "ge"
	FilterLike = "like"
	FilterIn   = "in"
)

type Value struct {
	S  string
	SS []string
	N  string
	NS []string
	B  bool
}

type Condition struct {
	Name    string
	Operand string
	Value   Value
}
type Filter struct {
	Conditions []Condition
}

func getAllowOps() []string {
	return []string{
		FilterEq, FilterLe, FilterLt, FilterGt, FilterGe, FilterLike, FilterIn, FilterNe,
	}
}
func isOperand(op string) bool {
	allowOps := getAllowOps()
	for _, cop := range allowOps {
		if op == cop {
			return true
		}
	}

	return false
}

//读到第一个空作为key
//读到第二空格作为Operand
//读到第三个空格作为value
//对于value
//后面没有其他的condition了 name eq ‘dsad and dsaf and adsaf ’
//后面还有其他的condition   name eq  'xxx and ddsa' and  id in ('12', '3214')
func words2condition(words []string, start int) (*Condition, int, error) {
	ws := words[start:]
	if len(ws) < 3 {
		return nil, 0, errors.New("jsonapi: wrong filter format")
	}
	key := ws[0]
	op := ws[1]
	if !isOperand(op) {
		allowOps := getAllowOps()
		return nil, 0, errors.New("jsonapi: " + op + " is a not supported operand only [" + strings.Join(allowOps, ",") + "] ")
	}
	val := ws[2]
	end := 3
	if len(ws) > 3 {
		for j := 3; j < len(ws); j++ {
			nws := ws[j:]
			if len(nws) > 3 {
				if !isOperand(nws[2]) {
					val = val + " " + ws[j]
					end = end + 1
				} else {
					if nws[0] != "and" {
						val = val + " " + ws[j]
						end = end + 1
					} else {
						end = end + 1
						break
					}

				}
			} else {
				//把剩下的接起来
				val = val + " " + strings.Join(ws[j:], " ")
				end = end + len(ws[j:])
				break
			}
		}
	}
	condition := new(Condition)
	condition.Name = key
	condition.Operand = op

	value := Value{}

	if op == FilterIn { //in ('a','b','c') or in (12,23,45,67) or in ('asd') or in (123)
		isN := true
		isS := true

		val = strings.TrimSpace(val)
		val = strings.TrimLeft(val, "(")
		val = strings.TrimRight(val, ")")
		tmp := strings.Split(val, ",")

		for _, t := range tmp {
			t = strings.TrimSpace(t)
			re := regexp.MustCompile("^\\d+(|.\\d+)$")
			if n := re.FindString(t); n != "" {
				value.NS = append(value.NS, n)
			} else {
				isN = false
				break
			}

		}

		for _, t := range tmp {
			t = strings.TrimSpace(t)
			if -1 == strings.Index(t, "'") {
				isS = false
				break
			} else {
				value.SS = append(value.SS, strings.Trim(t, "'"))
			}
		}

		if !isN && !isS {
			return nil, 0, errors.New("jsonapi: wrong filter format")
		}
	} else {
		re := regexp.MustCompile("^\\d+(|.\\d+)$")
		reb := regexp.MustCompile("^true|false$")
		if n := re.FindString(val); n != "" {
			value.N = n
		} else if n := reb.FindString(val); n != "" {
			if n == "true" {
				value.B = true
			} else {
				value.B = false
			}

		} else {
			if val == "''" {
				value.S = ""
			} else {
				value.S = strings.Trim(val, "'")
			}
		}
	}

	condition.Value = value

	return condition, end + start, nil
}

//@todo 去掉”号的情况
func NewFilter(filter string) (*Filter, error) {
	words := strings.Split(filter, " ")
	start := 0
	var f Filter
	for {
		cond, end, e := words2condition(words, start)
		if e != nil {
			return nil, e
		}
		f.Conditions = append(f.Conditions, *cond)
		if end >= len(words) {
			break
		}
		start = end
	}

	return &f, nil
}
