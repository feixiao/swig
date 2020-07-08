#### Basic 例子介绍

##### 指导go使用swig的例子
+ [simple](./simple)  C函数和全局变量的使用。
    ```
    cd ./simple/example
    swig -go -cgo -intgosize 64 example.i   
    go build
    ```
+ [constants](./constants) 常量的使用。
    ```
    cd ./constants/example
    swig -go -cgo -intgosize 64 example.i   
    go build 
    ```
+ [variables](./variables) 演示怎么从Go访问C的全局变量。。
    ```
    cd ./variables/example
    swig -go -cgo -intgosize 64 example.i   
    go build
    ```
+ [enum](./enum) 枚举类型封装。
    ```
    cd ./enum/example
    swig -go -cgo -c++ -intgosize 64 example.i  
    go build
    ```
+ [class](./class) C++类。
    ```
    cd ./go/class/example
    swig -go -cgo -c++ -intgosize 64 example.i   
    go build
    ```
+ [reference](./reference) C++引用。
    ```
    cd ./go/reference/example
    swig -go -cgo -c++ -intgosize 64 example.i   
    go build
    ```
+ [pointer](./pointer) 简单的指针操作。
    ```
    cd ./go/pointer/example
    swig -go -cgo -intgosize 64 example.i
    go build
    ```
+ [funcptr](./go/funcptr) 函数指针。
    ```
    cd ./go/funcptr/example
    swig -go -cgo -intgosize 64 example.i
    go build
    ```
+ [template](./template) C++模板类.
    ```
    cd ./go/template/example
    swig -go -cgo -c++ -intgosize 64 example.i
    go build
    ```
+ [callback](./go/callback)  利用directors使用C++回调函数。
    ```
    cd ./go/callback/example
    swig -go -cgo -c++ -intgosize 64 example.i
    go build
    ```
+ [extend](./extend) 利用directors使用C++多态。
    ```
    cd ./extend/example
    swig -go -cgo -c++ -intgosize 64 example.i
    go build
    ```
+ [director](./director)  如何利用utilize特性。
    ```
    cd ./director/example
    swig -go -cgo -c++ -intgosize 64 example.i
    go build
    ```