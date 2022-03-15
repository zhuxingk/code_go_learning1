http 分为四部分 
1.请求行
 1.格式：方法+URL+协议版本号
 2.实例：post+/chapter17/user + HTTP/1.1 
 3.请求方法： 
   a.GET：获取数据 
   b.POST：上传数据（表单格式，JSON格式）
   c.PUT：修改数据
   d.DELETE:用于删除数据 
2.请球头 
 1.格式：key:value
 2.可以有很多键值对（包括协议自带，也包含用户自定义）
 3.字段：
   a.ACCEPT：接受数据的格式
   b.User-Agent:用于描述用户浏览器的信息
   c.Connection：Keep_Alive（长链接），close（短链接）
   d.ACCEPT—Encoding:gzip,xxx,描述可以接受的编码
   e.Cookie:由服务器设置的key=value数据，客户端下次请求的时候可以携带
Content——Type：application/-form(表示 上传的是表单格式)
               application/JSON（表示BODY的数据是JSON格式）
   f.用户自定义：
    1.name:BOB
    2.age
    3.gender
    4.
3.空行//预示请求结束,用于分隔
4.请求包体//可选项
一般在POST方法的时候，会提供BODY
在GET的时候也可以提供BODY，一般不如此使用，会出现混淆
上传两种数据格式：
表单：姓名，性别，年龄
JSON数据格式
