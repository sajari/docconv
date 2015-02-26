package jsconfig

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

type JSConfig map[string]interface{}

func LoadConfig(path string) (JSConfig, error) {
	bytes, err := ioutil.ReadFile(path)
	if err == nil {
		var js JSConfig
		err = json.Unmarshal(bytes, &js)
		if err != nil {
			panic(err)
		}
		return js, nil
	}
	return nil, err
}

func (c JSConfig) loop(path string) interface{} {
	conf := c
	items := strings.Split(path, ".")
	var ivalue interface{}

	l := len(items)
	for i, item := range items {
		if i < l-1 {
			conf = conf[item].(map[string]interface{})
		}
		ivalue = conf[item]
	}
	return ivalue
}

func (c JSConfig) String(path string, def string) (value string) {
	defer func() {
		if r := recover(); r != nil {
			value = def
		}
	}()
	ivalue := c.loop(path)
	value = ivalue.(string)
	return
}

func (c JSConfig) Int(path string, def int) (value int) {
	defer func() {
		if r := recover(); r != nil {
			value = def
		}
	}()
	ivalue := c.loop(path)
	value = int(ivalue.(float64))
	return
}

func (c JSConfig) Bool(path string, def bool) (value bool) {
	defer func() {
		if r := recover(); r != nil {
			value = def
		}
	}()
	ivalue := c.loop(path)
	value = ivalue.(bool)
	return
}

func (c JSConfig) Float(path string, def float64) (value float64) {
	defer func() {
		if r := recover(); r != nil {
			value = def
		}
	}()
	ivalue := c.loop(path)
	value = ivalue.(float64)
	return
}

func (c JSConfig) Array(path string) []interface{} {
	ivalue := c.loop(path)
	value := ivalue.([]interface{})
	return value
}
