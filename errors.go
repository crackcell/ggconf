/***************************************************************
 *
 * Copyright (c) 2014, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Error definitions.
 *
 * @file errors.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Fri Oct 10 15:46:47 2014
 *
 **/

package ggconf

import (
	"errors"
)

//===================================================================
// Public APIs
//===================================================================

var (
	ErrInvalidContent = errors.New("invalid content")
	ErrKeyNotFound    = errors.New("key not found")
)

//===================================================================
// Private
//===================================================================
