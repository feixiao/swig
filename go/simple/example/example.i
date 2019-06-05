/* File : example.i */
%module example   // 定义模块

%inline %{
extern int    gcd(int x, int y);    // 导出方法
extern double Foo;                  // 导出全局变量
%}
