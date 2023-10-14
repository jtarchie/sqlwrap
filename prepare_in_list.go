// nolint
//
//line prepare_in_list.rl:1
package sqlwrap

import (
	"fmt"
	"github.com/nasa9084/go-builderpool"
	"strconv"
)

//line prepare_in_list.go:14
const prepare_in_list_start int = 15
const prepare_in_list_first_final int = 15
const prepare_in_list_error int = -1

const prepare_in_list_en_main int = 15

//line prepare_in_list.rl:13

func prepareInList(data string, params Params) (string, error) {
	builder, count := builderpool.Get(), 0
	defer builderpool.Release(builder)

	var start, end int
	cs, p, pe := 0, 0, len(data)

//line prepare_in_list.go:33
	{
		cs = prepare_in_list_start
	}

//line prepare_in_list.go:38
	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 15:
			goto st_case_15
		case 0:
			goto st_case_0
		case 1:
			goto st_case_1
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 16:
			goto st_case_16
		case 17:
			goto st_case_17
		case 18:
			goto st_case_18
		case 19:
			goto st_case_19
		case 20:
			goto st_case_20
		case 21:
			goto st_case_21
		case 22:
			goto st_case_22
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		case 12:
			goto st_case_12
		case 13:
			goto st_case_13
		case 14:
			goto st_case_14
		}
		goto st_out
	st_case_15:
		if data[p] == 32 {
			goto st1
		}
		goto st0
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
		if data[p] == 32 {
			goto st1
		}
		goto st0
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 32:
			goto st1
		case 73:
			goto st2
		case 105:
			goto st2
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 32:
			goto st1
		case 78:
			goto st3
		case 110:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 32 {
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 32:
			goto st4
		case 40:
			goto st5
		case 73:
			goto st2
		case 105:
			goto st2
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 32:
			goto st6
		case 58:
			goto tr7
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 32:
			goto st6
		case 58:
			goto tr7
		case 73:
			goto st2
		case 105:
			goto st2
		}
		goto st0
	tr7:
//line prepare_in_list.rl:23
		start = p
		goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line prepare_in_list.go:196
		if data[p] == 32 {
			goto st9
		}
		goto st8
	tr17:
//line prepare_in_list.rl:23
		start = p
		goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line prepare_in_list.go:210
		switch data[p] {
		case 32:
			goto tr10
		case 41:
			goto tr11
		}
		goto st8
	tr10:
//line prepare_in_list.rl:24

		{
			end = p

			paramName := data[start+1 : end]
			if _, ok := params[paramName]; !ok {
				return "", fmt.Errorf("could not find param for IN list %q", paramName)
			}

			builder.WriteString(data[count:start])

			values, ok := params[paramName].(Values)
			if !ok {
				return "", fmt.Errorf("could not read param %q as array", paramName)
			}

			for index, value := range values {
				indexParamName := paramName + strconv.Itoa(index)

				builder.WriteByte(data[start])
				builder.WriteString(indexParamName)

				if index < len(values)-1 {
					builder.WriteByte(',')
				}

				params[indexParamName] = value
			}

			delete(params, paramName)

			count = end
		}

		goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line prepare_in_list.go:260
		switch data[p] {
		case 32:
			goto tr10
		case 41:
			goto tr11
		case 73:
			goto st10
		case 105:
			goto st10
		}
		goto st8
	tr25:
//line prepare_in_list.rl:23
		start = p
		goto st16
	tr11:
//line prepare_in_list.rl:24

		{
			end = p

			paramName := data[start+1 : end]
			if _, ok := params[paramName]; !ok {
				return "", fmt.Errorf("could not find param for IN list %q", paramName)
			}

			builder.WriteString(data[count:start])

			values, ok := params[paramName].(Values)
			if !ok {
				return "", fmt.Errorf("could not read param %q as array", paramName)
			}

			for index, value := range values {
				indexParamName := paramName + strconv.Itoa(index)

				builder.WriteByte(data[start])
				builder.WriteString(indexParamName)

				if index < len(values)-1 {
					builder.WriteByte(',')
				}

				params[indexParamName] = value
			}

			delete(params, paramName)

			count = end
		}

		goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line prepare_in_list.go:318
		switch data[p] {
		case 32:
			goto tr19
		case 41:
			goto tr11
		}
		goto st16
	tr19:
//line prepare_in_list.rl:24

		{
			end = p

			paramName := data[start+1 : end]
			if _, ok := params[paramName]; !ok {
				return "", fmt.Errorf("could not find param for IN list %q", paramName)
			}

			builder.WriteString(data[count:start])

			values, ok := params[paramName].(Values)
			if !ok {
				return "", fmt.Errorf("could not read param %q as array", paramName)
			}

			for index, value := range values {
				indexParamName := paramName + strconv.Itoa(index)

				builder.WriteByte(data[start])
				builder.WriteString(indexParamName)

				if index < len(values)-1 {
					builder.WriteByte(',')
				}

				params[indexParamName] = value
			}

			delete(params, paramName)

			count = end
		}

		goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line prepare_in_list.go:368
		switch data[p] {
		case 32:
			goto tr19
		case 41:
			goto tr11
		case 73:
			goto st18
		case 105:
			goto st18
		}
		goto st16
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch data[p] {
		case 32:
			goto tr19
		case 41:
			goto tr11
		case 78:
			goto st19
		case 110:
			goto st19
		}
		goto st16
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		switch data[p] {
		case 32:
			goto tr22
		case 41:
			goto tr11
		}
		goto st16
	tr22:
//line prepare_in_list.rl:24

		{
			end = p

			paramName := data[start+1 : end]
			if _, ok := params[paramName]; !ok {
				return "", fmt.Errorf("could not find param for IN list %q", paramName)
			}

			builder.WriteString(data[count:start])

			values, ok := params[paramName].(Values)
			if !ok {
				return "", fmt.Errorf("could not read param %q as array", paramName)
			}

			for index, value := range values {
				indexParamName := paramName + strconv.Itoa(index)

				builder.WriteByte(data[start])
				builder.WriteString(indexParamName)

				if index < len(values)-1 {
					builder.WriteByte(',')
				}

				params[indexParamName] = value
			}

			delete(params, paramName)

			count = end
		}

		goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line prepare_in_list.go:450
		switch data[p] {
		case 32:
			goto tr22
		case 40:
			goto st21
		case 41:
			goto tr11
		case 73:
			goto st18
		case 105:
			goto st18
		}
		goto st16
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		switch data[p] {
		case 32:
			goto tr24
		case 41:
			goto tr11
		case 58:
			goto tr25
		}
		goto st16
	tr24:
//line prepare_in_list.rl:24

		{
			end = p

			paramName := data[start+1 : end]
			if _, ok := params[paramName]; !ok {
				return "", fmt.Errorf("could not find param for IN list %q", paramName)
			}

			builder.WriteString(data[count:start])

			values, ok := params[paramName].(Values)
			if !ok {
				return "", fmt.Errorf("could not read param %q as array", paramName)
			}

			for index, value := range values {
				indexParamName := paramName + strconv.Itoa(index)

				builder.WriteByte(data[start])
				builder.WriteString(indexParamName)

				if index < len(values)-1 {
					builder.WriteByte(',')
				}

				params[indexParamName] = value
			}

			delete(params, paramName)

			count = end
		}

		goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line prepare_in_list.go:520
		switch data[p] {
		case 32:
			goto tr24
		case 41:
			goto tr11
		case 58:
			goto tr25
		case 73:
			goto st18
		case 105:
			goto st18
		}
		goto st16
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		switch data[p] {
		case 32:
			goto tr10
		case 41:
			goto tr11
		case 78:
			goto st11
		case 110:
			goto st11
		}
		goto st8
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 32:
			goto tr14
		case 41:
			goto tr11
		}
		goto st8
	tr14:
//line prepare_in_list.rl:24

		{
			end = p

			paramName := data[start+1 : end]
			if _, ok := params[paramName]; !ok {
				return "", fmt.Errorf("could not find param for IN list %q", paramName)
			}

			builder.WriteString(data[count:start])

			values, ok := params[paramName].(Values)
			if !ok {
				return "", fmt.Errorf("could not read param %q as array", paramName)
			}

			for index, value := range values {
				indexParamName := paramName + strconv.Itoa(index)

				builder.WriteByte(data[start])
				builder.WriteString(indexParamName)

				if index < len(values)-1 {
					builder.WriteByte(',')
				}

				params[indexParamName] = value
			}

			delete(params, paramName)

			count = end
		}

		goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line prepare_in_list.go:604
		switch data[p] {
		case 32:
			goto tr14
		case 40:
			goto st13
		case 41:
			goto tr11
		case 73:
			goto st10
		case 105:
			goto st10
		}
		goto st8
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		switch data[p] {
		case 32:
			goto tr16
		case 41:
			goto tr11
		case 58:
			goto tr17
		}
		goto st8
	tr16:
//line prepare_in_list.rl:24

		{
			end = p

			paramName := data[start+1 : end]
			if _, ok := params[paramName]; !ok {
				return "", fmt.Errorf("could not find param for IN list %q", paramName)
			}

			builder.WriteString(data[count:start])

			values, ok := params[paramName].(Values)
			if !ok {
				return "", fmt.Errorf("could not read param %q as array", paramName)
			}

			for index, value := range values {
				indexParamName := paramName + strconv.Itoa(index)

				builder.WriteByte(data[start])
				builder.WriteString(indexParamName)

				if index < len(values)-1 {
					builder.WriteByte(',')
				}

				params[indexParamName] = value
			}

			delete(params, paramName)

			count = end
		}

		goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line prepare_in_list.go:674
		switch data[p] {
		case 32:
			goto tr16
		case 41:
			goto tr11
		case 58:
			goto tr17
		case 73:
			goto st10
		case 105:
			goto st10
		}
		goto st8
	st_out:
	_test_eof0:
		cs = 0
		goto _test_eof
	_test_eof1:
		cs = 1
		goto _test_eof
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof16:
		cs = 16
		goto _test_eof
	_test_eof17:
		cs = 17
		goto _test_eof
	_test_eof18:
		cs = 18
		goto _test_eof
	_test_eof19:
		cs = 19
		goto _test_eof
	_test_eof20:
		cs = 20
		goto _test_eof
	_test_eof21:
		cs = 21
		goto _test_eof
	_test_eof22:
		cs = 22
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
		goto _test_eof
	_test_eof12:
		cs = 12
		goto _test_eof
	_test_eof13:
		cs = 13
		goto _test_eof
	_test_eof14:
		cs = 14
		goto _test_eof

	_test_eof:
		{
		}
	}

//line prepare_in_list.rl:66

	builder.WriteString(data[count:])

	return builder.String(), nil
}
