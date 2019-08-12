package parser

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	md := `
	

---
title: "Front Matters"
description: "It really does"
---

This is some summary. This is some summary. This is some summary. This is some summary.

 <!--more-->

### Title


End value
`
	info, err := Parse(strings.NewReader(md))
	if err != nil {
		t.Fatalf("happp error: %v", err)
	}
	t.Logf("headers: %#v", info.Matters)
	t.Logf("content: %s", info.Content)
}
