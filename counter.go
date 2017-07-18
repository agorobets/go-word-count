package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

type Counter struct {
	re *regexp.Regexp
}

func NewCounter(re *regexp.Regexp) *Counter {
	return &Counter{
		re: re,
	}
}

func (c *Counter) LoadAndCount(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if res := c.re.FindAll(body, -1); res != nil {
		return len(res), nil
	}

	return 0, nil
}
