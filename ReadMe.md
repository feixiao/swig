## SWIG快速入门
#### 简介
SWIG是一个软件开发工具，它将用C和c++编写的程序与各种高级编程语言连接起来（为其他需要生成库）。准确地说，SWIG生成了两个文件，一个文件是*_wrapper.cpp文件，一个是*.go文件。*_wrapper.cpp文件将C++接口封装为C接口。go文件通过上一节说的import "C"来引用C接口，并把对这些C接口的调用，封装为不涉及任何C特性的Go函数或方法。
+ [项目地址](https://github.com/swig/swig)
+ [项目自带Examples](https://github.com/swig/swig/tree/master/Examples)
+ [文档](http://www.swig.org/Doc3.0/Contents.html#Contents)

#### SWIG安装
```shell
sudo apt-get install bison 

unzip swig-rel-3.0.12.zip
cd   swig-rel-3.0.12
./autogen.sh  && ./configure
make -j 4 && sudo make install
```

#### 使用入门
##### SWIG:Examples:Go
The following examples illustrate the use of SWIG with Go.
+ [simple](./go/simple/ReadMe.md)  A minimal example showing how SWIG can
be used to wrap a C function, a global variable, and a constant.
+ [constants](./go/constants/ReadMe.md) This shows how preprocessor macros and
certain C declarations are turned into constants.
+ [variables](./go/variables/ReadMe.md) An example showing how to access C global variables from Go.
+ [enum](./go/enum/ReadMe.md) Wrapping enumerations.
+ [class](./go/class/ReadMe.md) Wrapping a simple C++ class.
+ [reference](./go/reference/ReadMe.md) C++ references.
+ [pointer](./go/pointer/ReadMe.md) Simple pointer handling.
+ [funcptr](./go/funcptr/ReadMe.md)  Pointers to functions.
+ [template](./go/template/ReadMe.md) C++ templates.
+ [callback](./go/callback/ReadMe.md) C++ callbacks using directors.
+ [extend](./go/extend/ReadMe.md) Polymorphism using directors.
+ [director](./go/director/ReadMe.md)  Example how to utilize the director feature.

#### 参考资料
+ [《go通过swig封装、调用c++共享库的技术总结》](https://www.cnblogs.com/terencezhou/p/10059156.html)
+ [《使用swig工具为go语言与c++进行交互》](https://www.cnblogs.com/dongc/p/6896850.html)