package session

// session中间件开发
// 1、设计一个通用的session服务,支持内存存储和redis存储
// 2、session模块设计
// 2.1、本质k-v 通过k进行增删改查
// 2.2、session可以存储在内存或者redis
// 3、session接口设计
// 3.1、Set()
// 3.2、Get()
// 3.3、Del()
// 3.4、Save():session存储，redis实现延迟加载,用的时候再加载
// 4、sessionMgr接口设计：
// 4.1、CreateSession()：创建一个新的session
// 4.2、Init():初始化，加载redis地址
// 4.3、GetSession()：通过sessionId获取对应的session对象
// 5、MemorySession设计：
// 5.1、定义MemorySession对象(字段：sessionId、存kv的map，读写锁)
// 5.2、构造函数，为了获取对象
// 6、MemorySessionMgr设计：
// 6.1、定义MemorySessionMgr对象（字段：存放所有session的map，读写锁）
// 6.2、构造函数
// 6.3、Init()
// 6.4、CreateSession()
// 6.5、GetSession()
// 7、RedisSession设计
// 7.1、定义RedisSession对象（字段:sessionid，存kv的map，读写锁，redis连接池，记录内存中map是否被修改的标记）
// 7.2、构造函数
// 7.3、Set():将session存到内存中的map
// 7.4、Get():取数据，实现延迟加载
// 7.5、Del()
// 7.6、Save()：将session存到redis
// 8、RedisSessionMgr设计：
// 8.1、定义RedisSessionMgr对象(字段：redis地址、redis密码、连接池、读写锁、大map)
// 8.2、构造函数
// 8.3、Init()
// 8.4、CreateSession()
// 8.5、GetSession()

func main() {

}
