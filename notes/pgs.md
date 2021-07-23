# 索引

索引的扫描方式有3种

 1 Indexscan

先查索引找到匹配记录的ctid，再通过ctid查堆表

 2 bitmapscan

先查索引找到匹配记录的ctid集合，把ctid通过bitmap做集合运算和排序后再查堆表

 3 Indexonlyscan

如果索引字段中包含了所有返回字段，对可见性映射 (vm)中全为可见的数据块，不查堆表直接返回索引中的值。

> 在文本型数据列上创建索引，要加上**TEXT_PATTERN_OPS**，否则索引不能用于诸如like这样的操作符。

