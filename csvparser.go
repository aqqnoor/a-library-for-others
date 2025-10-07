package main

import (
	"errors"
	"io"
)

type CSVParser interface {
	ReadLine(r io.Reader) (string, error)
	GetField(n int) (string, error)
	GetNumberOfFields() int
}

var (
	ErrQuote      = errors.New("excess or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")
)

type Qwe struct {
	Line   string
	LineF  []string
	ExpNum int
}

func (q *Qwe) ReadLine(r io.Reader) (string, error) {
	buf := make([]byte, 1)
	var bit []byte
	inqu := false

	for {
		n, err := r.Read(buf)

		if n > 0 {

			char := buf[0]

			if char == '\n' && !inqu {
				break
			}

			if char != '\r' {
				bit = append(bit, char)
			}

			if char == '"' {
				inqu = !inqu
			}

		}

		if err != nil {

			if err == io.EOF && len(bit) > 0 {
				break
			}

			if err == io.EOF {
				return "", io.EOF
			}

			return "", err
		}
	}

	fields, perr := parseF(string(bit))
	if perr != nil {
		return "", perr
	}

	if q.ExpNum != len(fields) && q.ExpNum != 0 {
		return "", ErrFieldCount
	} else {
		q.ExpNum = len(fields)
	}

	q.LineF = fields

	var withoutqu []byte

	for i, field := range fields {
		if i > 0 {
			withoutqu = append(withoutqu, ',')
		}
		withoutqu = append(withoutqu, []byte(field)...)
	}

	q.Line = string(withoutqu)

	return q.Line, nil
}

func (q *Qwe) GetField(n int) (string, error) {
	if n < 0 || n >= len(q.LineF) {
		return "", ErrFieldCount
	}
	return q.LineF[n], nil
}

func (q *Qwe) GetNumberOfFields() int {
	return len(q.LineF)
}

func parseF(l string) ([]string, error) {
	var fields []string
	var tempF []rune
	inqu := false

	for i := range l {
		char := rune(l[i])

		if char == ',' && !inqu {
			fields = append(fields, removeQuotes(string(tempF)))
			tempF = nil
			continue
		}
		if char == '"' {
			if inqu && i+1 < len(l) && l[i+1] == '"' {
				tempF = append(tempF, char)
				i++
			} else {
				inqu = !inqu
			}
			continue
		}

		tempF = append(tempF, char)

	}
	if l[len(l)-1] != ',' {
		fields = append(fields, removeQuotes(string(tempF)))
	}

	if inqu {
		return nil, ErrQuote
	}
	return fields, nil
}

func removeQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}

	return s
}
