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
 
// todo 锁 事务 外键

 
 
 

 
 