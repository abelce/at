package jsonapi

import (
	"reflect"
	"testing"
)

func TestWords2Condition(t *testing.T) {
	//最普通条件
	filter := `name eq 'hahah'`
	except := Condition{
		Name:    "name",
		Operand: "eq",
		Value:   Value{S: "hahah"},
	}
	f, e := NewFilter(filter)
	if e != nil {
		t.Error(e)
	}
	if len(f.Conditions) != 1 {
		t.Errorf("except 1 got %d", len(f.Conditions))
	}
	cond := f.Conditions[0]
	if !reflect.DeepEqual(except, cond) {
		t.Errorf("jsonapi: except %v got %v", except, cond)
	}
	//document path问题
	filter = `x.y.z eq '123'`
	except = Condition{
		Name:    "x.y.z",
		Operand: "eq",
		Value:   Value{S: "123"},
	}
	f, e = NewFilter(filter)
	if e != nil {
		t.Error(e)
	}
	if len(f.Conditions) != 1 {
		t.Errorf("except 1 got %d", len(f.Conditions))
	}
	cond = f.Conditions[0]

	if !reflect.DeepEqual(except, cond) {
		t.Errorf("jsonapi: except %v got %v", except, cond)
	}
	//错误的key 有空格的key
	filter = `x.y z eq '123'`
	f, e = NewFilter(filter)
	if e == nil {
		t.Errorf("except a error[wrong filter format] got nil")
	}
	//错误的 不支持的 operand
	filter = `x.y.z ax '123'`
	f, e = NewFilter(filter)
	if e == nil {
		t.Errorf("except a error[not supported operand ...] got nil")
	}

	//值内有特殊空格
	filter = `name eq 'hi yo'`
	except = Condition{
		Name:    "name",
		Operand: "eq",
		Value:   Value{S: "hi yo"},
	}
	f, e = NewFilter(filter)
	if e != nil {
		t.Error(e)
	}
	if len(f.Conditions) != 1 {
		t.Errorf("except 1 got %d", len(f.Conditions))
	}
	cond = f.Conditions[0]
	if !reflect.DeepEqual(except, cond) {
		t.Errorf("jsonapi: except %v got %v", except, cond)
	}

	//内部混合了and & 操作符
	filter = `name eq 'hahah and h1 & d2' and address eq '1234'`
	except1 := Condition{
		Name:    "name",
		Operand: "eq",
		Value:   Value{S: "hahah and h1 & d2"},
	}
	except2 := Condition{
		Name:    "address",
		Operand: "eq",
		Value:   Value{S: "1234"},
	}
	f, e = NewFilter(filter)
	if e != nil {
		t.Error(e)
	}
	if len(f.Conditions) != 2 {
		t.Errorf("except 2 got %d", len(f.Conditions))
	}
	cond1 := f.Conditions[0]
	cond2 := f.Conditions[1]
	if !reflect.DeepEqual(except1, cond1) {
		t.Errorf("jsonapi: except %v got %v", except1, cond1)
	}
	if !reflect.DeepEqual(except2, cond2) {
		t.Errorf("jsonapi: except %v got %v", except2, cond2)
	}

	//内部混合中文
	filter = `name like '%Shangh中文ai Disneyland Park（Adult Ticket）%'`
	except = Condition{
		Name:    "name",
		Operand: "like",
		Value:   Value{S: "%Shangh中文ai Disneyland Park（Adult Ticket）%"},
	}
	f, e = NewFilter(filter)
	if e != nil {
		t.Error(e)
	}
	if len(f.Conditions) != 1 {
		t.Errorf("except 1 got %d", len(f.Conditions))
	}
	cond = f.Conditions[0]
	if !reflect.DeepEqual(except, cond) {
		t.Errorf("jsonapi: except %v got %v", except, cond)
	}
	//内部混合中文2
	filter = `name like '%Ancient City Wall, Shaanxi History Museum, Bell Tower and Drum Tower\'s and Muslims Quarter Bus Tour%'`
	except = Condition{
		Name:    "name",
		Operand: "like",
		Value:   Value{S: `%Ancient City Wall, Shaanxi History Museum, Bell Tower and Drum Tower\'s and Muslims Quarter Bus Tour%`},
	}
	f, e = NewFilter(filter)
	if e != nil {
		t.Error(e)
	}
	if len(f.Conditions) != 1 {
		t.Errorf("except 1 got %d", len(f.Conditions))
	}
	cond = f.Conditions[0]
	if !reflect.DeepEqual(except, cond) {
		t.Errorf("jsonapi: except %v got %v", except, cond)
	}

	//操作符情况-全是字符串
	filter = `name in ('n1','n 2', 'n3')`
	except = Condition{
		Name:    "name",
		Operand: "in",
		Value:   Value{SS: []string{"n1", "n 2", "n3"}},
	}
	f, e = NewFilter(filter)
	if e != nil {
		t.Error(e)
	}
	if len(f.Conditions) != 1 {
		t.Errorf("except 1 got %d", len(f.Conditions))
	}
	cond = f.Conditions[0]
	if !reflect.DeepEqual(except, cond) {
		t.Errorf("jsonapi: except %v got %v", except, cond)
	}
	//in-全是数字情况
	filter = `id in (1,123, 3.44)`
	except = Condition{
		Name:    "id",
		Operand: "in",
		Value:   Value{NS: []string{"1", "123", "3.44"}},
	}
	f, e = NewFilter(filter)
	if e != nil {
		t.Error(e)
	}
	if len(f.Conditions) != 1 {
		t.Errorf("except 1 got %d", len(f.Conditions))
	}
	cond = f.Conditions[0]
	if !reflect.DeepEqual(except, cond) {
		t.Errorf("jsonapi: except %v got %v", except, cond)
	}
	//in-全是bool情况 这种情况 现实中不应该存在
	filter = `status in (true,false, true)`
	f, e = NewFilter(filter)
	if e == nil {
		t.Error("except error[wrong filter format] got nil")
	}
	//in-混合类型 这种情况 也不应存在
	filter = `status in ('asd',123, '233')`
	f, e = NewFilter(filter)
	if e == nil {
		t.Error("except error[wrong filter format] got nil")
	}

	//有其他的条件
	filter = `name eq 'hahah' and isDeleted eq true and k1 in ('123','456') and k2 in (12,34,56) and k3 ne '123' and k4 ge 4 and k41 gt 41 and k5 le '5' and k51 lt '51' and k6 like '%34z%'`
	excepts := []Condition{
		Condition{
			Name:    "name",
			Operand: "eq",
			Value:   Value{S: "hahah"},
		},
		Condition{
			Name:    "isDeleted",
			Operand: "eq",
			Value:   Value{B: true},
		},
		Condition{
			Name:    "k1",
			Operand: "in",
			Value:   Value{SS: []string{"123", "456"}},
		},
		Condition{
			Name:    "k2",
			Operand: "in",
			Value:   Value{NS: []string{"12", "34", "56"}},
		},
		Condition{
			Name:    "k3",
			Operand: "ne",
			Value:   Value{S: "123"},
		},
		Condition{
			Name:    "k4",
			Operand: "ge",
			Value:   Value{N: "4"},
		},
		Condition{
			Name:    "k41",
			Operand: "gt",
			Value:   Value{N: "41"},
		},
		Condition{
			Name:    "k5",
			Operand: "le",
			Value:   Value{S: "5"},
		},
		Condition{
			Name:    "k51",
			Operand: "lt",
			Value:   Value{S: "51"},
		},
		Condition{
			Name:    "k6",
			Operand: "like",
			Value:   Value{S: "%34z%"},
		},
	}
	f, e = NewFilter(filter)
	if e != nil {
		t.Error(e)
	}
	if len(f.Conditions) != 10 {
		t.Errorf("except 10 got %d", len(f.Conditions))
	}
	for i := 0; i < len(f.Conditions); i++ {
		if !reflect.DeepEqual(excepts[i], f.Conditions[i]) {
			t.Errorf("jsonapi: except %v got %v", excepts[i], f.Conditions[i])
		}
	}

}
