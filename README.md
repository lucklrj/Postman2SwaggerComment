# Postman2SwaggerComment
将postman导出的**json**文件，转换为**Swagger**的注释文档，一般应用在php开发中，postman里调试接口，采用Swagger产生api文档的场景。

-------

注意：

1.需要安装go包管理工具：[glide](https://github.com/Masterminds/glide)

2.部分包需要翻墙，可以设置包的github上的镜像，参考：[go](https://github.com/golang/go)

3.postman参数没有必填，选填选项，产生的swagger注释文档都设置为required=true，需要手动补充

4.个别字段需要手动补充说明

5.postman数据里有data:[]这样的空数据形式，已经转换为data为空的字符串