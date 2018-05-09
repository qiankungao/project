redis连接池的一些相关结构和方法
type Pool struct {
    //创建一个tcp链接的匿名函数
    Dial func() (Conn, error)
    //可选的函数，用来对之前使用过的空闲链接进行安全检查
    TestOnBorrow func(c Conn, t time.Time) error
    //最大可用链接数
    MaxIdle int
    //在给定的时间，最大可分配的连接数。为0则不限制
    MaxActive int
    //空闲链接的关闭时间，如果为空，空闲链接不会被关闭
    IdleTimeout time.Duration
    //如果Wait为true并且进行MaxActive限制了，Get()将会等待链接被返回
    Wait bool
    //关闭老链接的时间区间。如果为空则不会在生命周期内关闭链接
    MaxConnLifetime time.Duration
}
//创建一个Pool结构体
func NewPool(newFn func() (Conn, error), maxIdle int) *Pool

// 获取pool的连接数，包含空闲链接和使用连接 
func (p *Pool) ActiveCount() int


//关闭并释放pool中的资源
func (p *Pool) Close() error

//从pool中获取一个连接
func (p *Pool) Get() Conn

// 获取空闲链接数
func (p *Pool) IdleCount() int

// 获取连接池的状态信息
func (p *Pool) Stats() PoolStats
type PoolStats struct {
    ActiveCount int
    IdleCount   int
}


