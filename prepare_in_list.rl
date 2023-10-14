//nolint
package sqlwrap

import (
  "github.com/nasa9084/go-builderpool"
  "fmt"
  "strconv"
)

%%{
  machine prepare_in_list;
  write data;
}%%

func prepareInList(data string, params Params) (string, error) {
	builder, count := builderpool.Get(), 0
	defer builderpool.Release(builder)

  var ts, te, act, start, end int
  cs, p, pe, eof := 0, 0, len(data), len(data)

  %%{
    action start { start = p }
    action name {
      {
        end = p

        paramName := data[start+1: end]
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
    }
    whitespace = " ";
    name       = ([:] any+) >start %name;
    in_list = whitespace+ /in/i whitespace+ "(" whitespace* name whitespace* ")";
    
    main := (any* in_list any*)*;

    write init;
    write exec;
  }%%

  builder.WriteString(data[count:])

  _, _, _, _ = ts, te, act, eof

  return builder.String(), nil
}