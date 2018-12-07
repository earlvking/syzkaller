// Copyright 2018 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package proggen

import (
	"github.com/google/syzkaller/prog"
	_ "github.com/google/syzkaller/sys"
)

var target = func() *prog.Target {
	target, err := prog.GetTarget("linux", "amd64")
	if err != nil {
		panic(err)
	}
	target.ConstMap = make(map[string]uint64)
	for _, c := range target.Consts {
		target.ConstMap[c.Name] = c.Value
	}
	return target
}()

func Fuzz(data []byte) int {
	progs, err := ParseData(data, target)
	if err != nil {
		return 0
	}
	return len(progs)
}