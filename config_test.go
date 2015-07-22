/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Test
 *
 * @file config_test.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Oct 11 17:14:58 2014
 *
 **/

package ggconf

import (
	_ "fmt"
	"testing"
)

//===================================================================
// Public APIs
//===================================================================

var c Config

func TestLoad(t *testing.T) {
	if err := c.Load("./test.conf"); err != nil {
		t.Error(err)
	}
}

func TestString(t *testing.T) {
	if v, err := c.GetString("key1"); err != nil || v != "value1" {
		t.Error()
	}
}

func TestInt(t *testing.T) {
	if v, err := c.GetInt("keyint"); err != nil || v != 1 {
		t.Error()
	}
}

//===================================================================
// Private
//===================================================================
