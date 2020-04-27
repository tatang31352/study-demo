# 一 索引分类
 普通索引：仅仅加速查询
 
 唯一索引：加速查询 + 列值唯一（可以有null）
 
 主键索引：加速查询 + 列值唯一（不可以有null）
 
 组合索引：多列值组成一个索引,专门用于组合搜索
 
 全文索引：对文本的内容进行分词,进行搜索
 
 # 二 explain详解
 id:选择标识符  
 select_type:表示查询的类型。  
 table:输出结果集的表  
 partitions:匹配的分区  
 type:表示表的连接类型  
 possible_keys:表示查询时,可能使用索引  
 key:表示实际使用的索引  
 key_len:索引字段的长度  
 ref:列与索引的比较  
 rows:扫描出的行数(估算的行数)  
 filtered:按表条件过滤的行百分比
 Extra:执行情况的描述和说明
 
 
 # 三 索引类型
 FULLTEXT 全文索引   
 HASH    
 BTREE   
 PTREE
 
 #  四 查看sql执行时间
 1 show profiles;  
 2 show variables;查看profiling 是否是on状态；  
 3 如果是off，则 set profiling = 1；  
 4 执行自己的sql语句；  
 5 show profiles；就可以查到sql语句的执行时间；
 
**锁 事务 外键**  

        外键:外键是指引用另外一个表中的一列或多列数据，被引用的列应该具有主键约束或者唯一性约束。
    外键用来建立和加强两个表数据之间的连接。
     
    事务和锁之间的关系:锁是实现事务其中一个特性的机制。
    
     悲观锁定义
    ​   每次获取数据的时候担心数据被修改, 所以每次获取数据的时候都会进行加锁, 确保自己使用过程中数
    据不会被别人修改, 使用完成后对数据进行解锁. 由于数据进行加锁, 期间对改数据进行读写的其他线程
    都会进行等待
    
    乐观锁定义
    ​  每次获取数据的时候都不会担心数据被修改, 所以每次获取数据的时候都不会进行加锁, 但是在更新数据
    的时候需要判断该数据是否被别人修改过, 如果数据被其他线程修改则不进行数据更新, 如果数据没有被其
    他线程修改, 则进行数据更新, 由于数据没有进行加锁, 期间该数据可以被其他线程进行读写操作
    
    悲观锁使用场景
    ​   比较适合写入操作比较频繁的场景, 如果出现大量的读取操作, 每次读取的时候都会进行加锁, 这样会增
    加大量的锁的开销, 降低了系统的吞吐量
 
    乐观锁使用场景
        比较适合读取操作比较频繁的场景, 如果出现大量的写入操作, 数据发生冲突的可能性就会增大, 为了保
     证数据的一致性, 应用层需要不断的重新获取数据, 这样会增加大量的查询操作, 降低了系统的吞吐量
    
    小结: 两者各有优缺点, 读取频繁使用乐观锁, 写入频繁使用悲观锁
    ​ 悲观锁比较适合强一致性的场景, 但效率比较低, 特别是读的并发低, 乐观锁则适用于读多写少, 并发冲突少的场景!!
    
    乐观锁实现
        比如我们查询到goods表中version 为1 那么在更新这个表的时候Sql将是
        select * from goods where id=1;
        update goods set status=2,version=version+1 where id=1 and version=1;
          
    悲观锁实现
        首先设置MySQL非autocommit模式:
        set autocommit=0;
        设置完autocommit后,我们就可以执行我们正常的业务了。具体如下:
        
        0.开始事务
        begin;/begin work;/start transaction;(三者选一就可以)
        
        1.查询出商品信息
        select status from t_goods where id=1 for update;
        
        2.根据商品信息生成订单
        insert into t_orders(id,goods_id) values(null,1);
        
        3.修改商品status为2
        update t_goods set status=2;
        
        4.提交事务
        commit;/commit work;
        
    死锁怎么解决
        等待，直到超时（innodb_lock_wait_timeout=50s）。
        发起死锁检测，主动回滚一条事务，让其他事务继续执行(innodb_deadlock_detect=on)。
               
 

 
 