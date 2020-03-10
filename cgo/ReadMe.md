## 使用

#### 编译C库
```sh
gcc -Wall -c hello.c
ar -rv libhello.a hello.o
```

#### 编译Go程序
```sh
go build 
```

#### 注意事项
+ /* ... */ cgo 注释部分和 import "C" 之间不能存在空行