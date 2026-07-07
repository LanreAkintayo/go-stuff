package main

import (
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)


type errors map[string][]string

var EmailRX = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	if err, ok := e[field]; ok {
		return err[0]
	}
	return ""
}

type Form struct {
	url.Values
	Errors errors
}


func NewForm(data url.Values) *Form {
	return &Form{
		Values: data,
		Errors: make(errors),
	}
}

func (f *Form) Required(fields ...string) *Form{
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}

	return f
}

func (f *Form) Valid() bool{
	return len(f.Errors) == 0
}

func (f *Form) MaxLength(field string, count int) *Form {
	value := f.Get(field)
	if value == ""{
		return f
	}
	if utf8.RuneCountInString(value) > count {
		f.Errors.Add(field, "This field is too long")
	}

	return f
}

func (f *Form) MinLength(field string, count int) *Form {
	value := f.Get(field)
	if value == ""{
		return f
	}
	if utf8.RuneCountInString(value) < count {
		f.Errors.Add(field, "This field is too short")
	}

	return f
}

func (f *Form) Matches(field string, pattern *regexp.Regexp) *Form {
	value := f.Get(field)
	if value == ""{
		return f
	}
	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field does not match the pattern")
	}

	return f
}