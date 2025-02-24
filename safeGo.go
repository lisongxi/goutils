package goutils

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
)

func SafeGo(ctx context.Context, fn func()) {
	if fn == nil {
		log.Println("SafeGo: nil function provided")
		return
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				err := fmt.Errorf("panic: %v", r)
				stack := debug.Stack()

				log.Printf(
					"[SafeGo] error=%s | stack:\n%s",
					err.Error(),
					string(stack),
				)
				// TODO collect error info
			}
		}()

		fn()
	}()
}
