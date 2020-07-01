// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64 amd64p32 386

package runtime

import (
	"runtime/internal/sys"
	"unsafe"
)

// adjust Gobuf as if it executed a call to fn with context ctxt
// and then did an immediate gosave.
func gostartcall(buf *gobuf, fn, ctxt unsafe.Pointer) {
	sp := buf.sp
	if sys.RegSize > sys.PtrSize {
		sp -= sys.PtrSize
		*(*uintptr)(unsafe.Pointer(sp)) = 0
	}
	sp -= sys.PtrSize                        //在栈上空出一个返回位置
	*(*uintptr)(unsafe.Pointer(sp)) = buf.pc //此处的pc 指向 goexit + 1 的位置
	buf.sp = sp                              //重新设置sp
	//将newg的pc指向runtime.main
	buf.pc = uintptr(fn)
	buf.ctxt = ctxt
}
