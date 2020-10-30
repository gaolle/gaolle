---
title: validator
date: 2020-10-29 10:29:13
tags: validator
---

```go
package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
)

type Time struct {
	StartTime int64 `json:"start_time" validate:"required,textFunction"`    // 时间戳
	EndTime   int64 `json:"end_time" validate:"required,gtfield=StartTime"` // 时间戳
}

func textFunction(fl validator.FieldLevel) bool {
	today := fl.Field().Int()
	now := time.Now().Unix()
	if today > now {
		return false
	}
	return true
}

func main() {
	validate := validator.New()
	validate.RegisterValidation("textFunction", textFunction)

	t := &Time{
		StartTime: 111,
		EndTime:   2222,
	}

	if err := validate.Struct(t); err != nil {
		fmt.Println(err)
	}
}

```

