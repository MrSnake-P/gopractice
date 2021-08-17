package atomic

import "sync/atomic"

type Boolean uint32

func (b *Boolean) Get() bool {
	/* 函数atomic.LoadInt32接受一个*int32类型的指针值，并会返回该指针值指向的那个值
	 有了“原子的”这个形容词就意味着，在这里读取value的值的同时，当前计算机中的任何CPU都不会进行其它的针对此值的读或写操作。
	这样的约束是受到底层硬件的支持的。 */
	return atomic.LoadUint32((*uint32)(b)) != 0
}

func (b *Boolean) Set(v bool) {
	/*   在原子的存储某个值的过程中，任何CPU都不会进行针对同一个值的读或写操作。
	如果我们把所有针对此值的写操作都改为原子操作，那么就不会出现针对此值的读操作因被并发的进行而读到修改了一半的值的情况了。
	原子的值存储操作总会成功，因为它并不会关心被操作值的旧值是什么。
	函数atomic.StoreInt32会接受两个参数。第一个参数的类型是*int 32类型的，
	其含义同样是指向被操作值的指针。而第二个参数则是int32类型的，它的值应该代表欲存储的新值。
	其它的同类函数也会有类似的参数声明列表。*/
	if v {
		atomic.StoreUint32((*uint32)(b), 1)
	} else {
		atomic.StoreUint32((*uint32)(b), 0)
	}
}