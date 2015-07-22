/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Config
 *
 * @file config.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Oct 10 15:24:00 2014
 *
 **/

package ggconf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//===================================================================
// Public APIs
//===================================================================

type Config struct {
	ok bool
	m  map[string]string
}

func (this *Config) Load(path string) error {
	this.init()
	if err := this.parse(path); err != nil {
		return err
	}
	return nil
}

func (this *Config) GetString(key string) (string, error) {
	this.init()
	if value, found := this.m[key]; found {
		return value, nil
	} else {
		return "", ErrKeyNotFound
	}
}

func (this *Config) GetInt(key string) (int64, error) {
	this.init()
	if value, found := this.m[key]; found {
		return strconv.ParseInt(value, 10, 64)
	} else {
		return 0, ErrKeyNotFound
	}
}

func (this *Config) GetFloat(key string) (float64, error) {
	this.init()
	if value, found := this.m[key]; found {
		return strconv.ParseFloat(value, 64)
	} else {
		return 0, ErrKeyNotFound
	}
}

//===================================================================
// Private
//===================================================================

const (
	SEP = ":"
)

func (this *Config) init() {
	if this.ok {
		return
	}
	this.m = make(map[string]string)
	this.ok = true
}

func (this *Config) parse(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	br := bufio.NewReader(f)
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		this.parseLine(string(line))
	}

	return nil
}

func (this *Config) parseLine(line string) {
	tokens := strings.Split(line, SEP)
	if len(tokens) != 2 {
		return
	}
	key := strings.Trim(tokens[0], " \t")
	value := strings.Trim(tokens[1], " \t")
	if _, found := this.m[key]; found {
		fmt.Fprintln(os.Stderr, "duplicate key found")
	}
	this.m[key] = value
}
