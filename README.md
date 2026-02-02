## work 3 - ToDoList

### 项目框架： GORM + HERTZ

### Bonus1 : 自动生成接口文档
使用swagger自动生成了文档

### Bonus2 : 数据库交互安全性
使用了 GORM 的参数化查询（? 占位符），而不是字符串拼接，防止了SQL注入攻击

### Bonus3 : 三层架构设计
项目结构非常清晰对应三层：
Controller (api/v1)：接收 HTTP 请求，参数校验<br>Service (service)：核心业务逻辑（分页计算、逻辑判断）<br>DAO (dao)：直接与数据库交互（GORM 操作）<br>

#### bonus4, 5 还没想好。



