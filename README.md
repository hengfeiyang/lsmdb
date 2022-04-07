# LSM Tree demo project


#### set flow

1. command -> memory SSTable -> WAL
2. memory SSTable -> swich SSTable to immutable
3. flush immutable to sdb
4. delete WAL

#### query flow

1. query -> memory SSTable -> active SSTable -> immutable SSTable -> sparse Index -> sdb file

#### sdb file (SSTable)

|           N            |     |           N            |      N       |    40    |
|------------------------|-----|------------------------|--------------|----------|
| blockLength, blockData | ... | blockLength, blockData | sparse index | metainfo |

#### spare index data

|      4      |     4     |      N                   |
|-------------|-----------|--------------------------|
| index.start | keyLength | index.key                |


#### meata info data

|     8     |      8     |      8     |     8       |     2       |     2         |    4    |
|-----------|------------|------------|-------------|-------------|---------------|---------|
| dataStart | dataLength | indexStart | indexLength | blockKeyNum | tableBlockNum | version |


#### block data && wal file

|     4         |    N    |      4        |    N    |     |     4         |    N    |
|---------------|---------|---------------|---------|-----|---------------|---------|
| commandLength | command | commandLength | command | ... | commandLength | command |


#### command data

|      1      |     4     |  N    |    4        |  N       |
|-------------|-----------|-------|-------------|----------|
| commandType | keyLength | key   | valueLength | value    |



1. block data use LZ4 compressed
2. wal log is not compressed
