package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	endpoint := "127.0.0.1:9000"
	key := "GnFz1opVMB2awtzfp66y"
	secret := "0SV6TCcUg2CqeMKe304BsURNnH0KeTjcmymAZAS3"
	bucket := "ttt"
	object := "Sample-MP4-Video-File-Download.mp4"
	expiry := 85400
	url := presignedGetObject_(endpoint, key, secret, bucket, object, expiry)
	t.Log(url)
}
