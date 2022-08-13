package apihandler

import (
	"context"
	"time"
)

func SendSSEPingMessage(writer httpFlushWriter, requestContext context.Context) {
	go func() {
		for {
			select {
			case <-time.After(20 * time.Second):
				_, err := writer.Write([]byte(":" + time.Now().String()))
				if err != nil {
					return
				}
				writer.Flush()
			case <-requestContext.Done():
				return
			}
		}
	}()
}
