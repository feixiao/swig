## Simple Go Example
这个Sample会指导您如何将Go挂钩到一个非常简单的C程序，它只包含一个函数和一个全局变量。

#### C代码
+ [simple.c](./simple.c)


#### SWIG接口
+ [example.i](./example.i)


#### 编译使用
+ 生成example.go, example_gc.c和example_wrap.c文件
    ```
    cd ./example
    swig -c++ -go -intgosize 64 -cgo example.i  // C++ 
    swig -go -cgo -intgosize 64 example.i       // C

    go build 
    ```
+ 编译运行sample
    ```
    cd ../
    go build
    ./sample
    ```

