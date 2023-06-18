
This is a demo clib for minio. Maked by Golang

### Environment 
MacOs 10.15.7
go version go1.20.3 darwin/amd64

### How to Use?

Run Shell Script
```
cd minio-clib
chmod +x build.sh
./build.sh mac
./build.sh win
./build.sh linux
```

### How to Test Locally?
```
./build.sh mac
cd test && ./exec.sh
```
Then this below will show in the std ouput of commandline
```
Presigned URL: http://127.0.0.1:9000/ttt/Sample-MP4-Video-File-Download.mp4?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=GnFz1opVMB2awtzfp66y%2F20230618%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20230618T031815Z&X-Amz-Expires=85400&X-Amz-SignedHeaders=host&X-Amz-Signature=b62d8907fc61dae6997a20e5e48974c857a61e67cb3582bbc78840a1b3bcfa9b
```

### How To build linux clib from mac
``` 
brew tap messense/macos-cross-toolchains
brew install x86_64-unknown-linux-gnu
brew install aarch64-unknown-linux-gnu
```

