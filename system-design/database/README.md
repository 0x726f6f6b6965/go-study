# Database

## About it
- A database is an organized collection of data that can be managed and accessed easily. Databases are created to make it easier to store, modify, and delete data in connection with different data-processing procedures. There are two basic types of databases: SQL and NoSQL.

## Advantages

1. Managing large data
   - A large amount of data can be easily handled with a database.
2. Retrieving accurate data
   - Due to different constraints in databases, we can retrieve accurate data whenever we want.
3. Easy updation
   - It's easy to update data in databases using DML.
4. Security
   - Databases ensure the security of the data. A database only allows authorized users to access data.
5. Data integrity
   - Databases ensure data integrity by using different constraints for data.
6. Availability
   - Databases can be replicated on different servers.
7. Scalability

## Relational databases

- Relational databases need to create schemas before storing the data so the data stored in databases is structured. All the data has at least one unique key, called the primary key, on the table, and the data in one table can be linked to another table by the foreign key.

- Relational databases provide the ACID properties to maintain the integrity of the databases. The ACID will be explained below. 
  1. Atomicity: A transaction is considered an atomic unit. Therefore, either all the statements in a transaction will execute successfully, or none of them will execute. If a statement fails in a transaction, it should be aborted and rolled back.
  2. Consistency: The databases should be in a consistent state, and they should remain consistent after every transaction.
  3. Isolation: In the case of multiple transactions running concurrently, they shouldn't be affected by each other. The final state of the database should be the same as the transactions that were executed sequentially.
  4. Durability: The system should guarantee that completed transactions will survive permanently in databases.

## Non-relational databases

- NoSQL databases are designed with a variety of data models to access and manage data. These databases are used in applications that require a large volume of semi-structured and unstructured data, low latency, and flexible data models. Some different types of NoSQL databases will be discussed below: Key-value databases, Document databases, Graph databases, and Columnar databases.

1. Key-value databases
   - It uses key-value methods like hash tables to store data in pairs. The key serves as a unique key and the values can be anything ranging from simple scalar values to complex objects.
   - It is suitable for session-oriented applications, such as web applications that store users' data in main memory or a database during a session.

2. Document databases
   - It is designed to store and retrieve documents in formats like XML, JSON, BSON and so on. Documents in this type of database may have varying structures and data.
   - It is suitable for unstructured catalog data. For example, in e-commerce applications, a product has thousands of attributes, which is unfeasible to store in relational databases due to its impact on reading performance. However, the document databases can store each attribute in a single file for easy management and faster reading speed. Moreover, it's also a good option for content management applications, such as blogs and video platforms.

3. Graph databases
   - It uses the graph structure to store data, where nodes represent entities, and edges show the relationships between entities. This database allows us to store data once and then interpret it differently based on relationships.
   - It can be used in social applications and provide interesting facts and data about different types of users and their activities.

4. Columnar databases
   - It stores data in columns instead of rows. This way can enable access to all entries in the database column quickly and efficiently.

## Replication

- Databases usually keep multiple copies of the data at various nodes to achieve availability, scalability and performance. There are two ways to propagate changes to replica nodes: Synchronous replication and Asynchronous replication.
  1. In synchronous replication, the primary node waits for acknowledgments from secondary nodes about updating the data. After the primary node receives data from all the secondary nodes, the primary node reports success to clients.
     - Advantage: All the secondary nodes are completely up to date with the primary node.
     - Disadvantage: High latency in response from the primary node to clients.
  2. In asynchronous replication, the primary node doesn't need to wait for acknowledgments from secondary nodes and reports success to the client after updating itself.
     - Advantage: It tends to cost significantly less than synchronous replication.
     - Disadvantage: weak consistency.

- There are some data replication models I will discuss below: Primary-secondary replication, Multi-leader replication, and Leaderless replication.
  1. Primary-secondary(Single leader):
     - In primary-secondary replication, data is replicated across multiple nodes. The primary node is responsible for processing any writes to data stored on the clusters.
     - This model is appropriate when the application is read-heavy. To better scale with increasing readers, developers can add more followers and distribute the read load across the available nodes. However, if the workload is write-heavy, there is a primary bottleneck in the replicating data to all the secondary nodes.
     - If the primary node fails, any missed updates not passed on to the secondary nodes can be lost.
  2. Multi-leader:
     - This kind of replication is quite useful in applications in which we can continue to work even if we’re offline—for example, a calendar application in which we can set our meetings even if we don’t have access to the internet. Once we’re online, it replicates its changes from our local database to other nodes. However, this action creates a problem called data conflict. There are some approaches to handle the problem: Conflict avoidance, Last-write-win, and Custom logic.
       - Conflict avoidance: Conflicts can be avoided if the application can verify that all writes for a given record go via the same leader.
       - Last-write-win: Using their local clock or Google's Spanner, all nodes assign a timestamp to each update. When a conflict occurs, the update with the latest timestamp is selected.
       - Custom logic: In this approach, we can write our logic to handle conflicts according to the needs of our application. 
  3. Leaderless(Peer-to-peer):
     - Unlike primary-secondary replication, in the leaderless approach, all the nodes have equal weightage and can accept read and write requests. However, this approach also yields inconsistency. A helpful method used for solving write-write inconsistency is called quorums.

## Partitioning
- Nowadays, a single node-based database isn't enough to tackle the load because of increasing data. We might need to distribute the data over many nodes but still export all the nice properties of relational databases. Data partitioning enables us to use multiple nodes where each node manages some part of the whole data. There are different ways to partition data: Vertical sharding, and Horizontal sharding.

### Vertical sharding
- Data will be divided and stored by attributes, and all data will have a unique key, just like a foreign key in a relational database, used to link other data. When data is divided into multiple pieces, they can be separated on different servers. In this way, the speed of retrieving data stored in the binary large object will be increased. However, developers should be careful if there are joint operations between multiple tables because its cost is expensive.

### Horizontal sharding
- Data will be divided and stored by rows. Each partition of the original table distributed over database servers is called a shard. Usually, there are two strategies: Key-range-based sharding, and Hash-based sharding.

#### Key-range-based sharding
- In the key-range-based sharding, each partition is assigned a continuous range of keys.
- Advantages
  - The range-query-based schema is easy to implement.
  - Range queries can be performed using the partitioning keys, and those can be kept in partitions in sorted order. 
- Disadvantages
  - Range queries can't be performed using keys other than the partitioning key.
  - If keys aren’t selected properly, some nodes may have to store more data due to an uneven distribution of the traffic.

#### Hash-based sharding
- Hash-based sharding uses a hash-like function on an attribute and separates all the data by the function result.
- Advantages
  - Keys are uniformly distributed across the nodes.
- Disadvantages
  - We can’t perform range queries with this technique. Keys will be spread over all partitions.

### Consistent hashing
- Consistent hashing assigns each server or item in a distributed hash table a place on an abstract circle, called a ring, irrespective of the number of servers in the table. This permits servers and objects to scale without compromising the system’s overall performance.
- Advantages
  - It’s easy to scale horizontally.
  - It increases the throughput and improves the latency of the application.
- Disadvantages
  - Randomly assigning nodes in the ring may cause non-uniform distribution.

### Rebalance the partition
- When data increases, developers may need to add new servers to handle the new data. At this time the partitions will need to be rebalanced. We can apply the following strategies to rebalance partitions.
  - Avoid hash mod `n`
    - The problem with the addition or removal of nodes in the case of `hash mod n` is that every node’s partition number changes and a lot of data moves.
  - Fixed number of partitions
    - We can fix the higher number of partitions than the nodes when we set our databases up, so when a new node is added to the system, it can take a few partitions from the existing nodes until the partitions are equally divided. There is a downside to this approach. The size of each partition grows with the total amount of data in the cluster since all the partitions contain a small part of the total data. If a partition is very small, it will result in too much overhead because we may have to make a large number of small-sized partitions, each costing us some overhead. If the partition is very large, rebalancing the nodes and recovering from node failures will be expensive.
  - Dynamic partitioning
    - When the size of a partition reaches the threshold, it’s split equally into two partitions. In this way, the load is divided equally. The number of partitions adapts to the overall data amount, which is an advantage of dynamic partitioning. However, there's a downside to this approach. It’s difficult to apply dynamic rebalancing while serving the reads and writes which means the system needs to go down for rebalanceing.
  - Partition proportionally to nodes
    - The number of partitions is proportionate to the number of nodes, which means every node has fixed partitions. However, as the number of nodes increases, the partitions shrink. When a new node enters the network, it splits a certain number of current partitions at random, then takes one-half of the split and leaves the other half alone. This approach can result in an unfair split.

### Partitioning and secondary indexes
- Some approaches can partition with secondary indexes: By document and By term. 
#### By document
- Each partition is fully independent in this indexing approach. Each partition has its secondary indexes covering just the documents in that partition. If we want to write anything to our database, we need to handle that partition only containing the document ID we’re writing. It’s also known as the local index.
- However, this type of querying on secondary indexes can be expensive. As a result of being restricted by the latency of a poor-performing partition, read query latencies may increase.
#### By term
- This is also known as the global index which means that the index encompasses data from all partitions and this approach is more read-efficient than partitioning secondary indexes by the document.
- However, a single write in this approach affects multiple partitions, making the method write-intensive and complex.

## Advantages and disadvantages of a centralized database
### Advantages
- Data maintenance, such as updating and taking backups of a centralized database, is easy.
- Centralized databases provide stronger consistency and ACID transactions than distributed databases.
- Centralized databases provide a much simpler programming model for the end programmers as compared to distributed databases.
- It’s more efficient for businesses that have a small amount of data to store that can reside on a single node.
### Disadvantages
- A centralized database can slow down, causing high latency for end users, when the number of queries per second accessing the centralized database is approaching single-node limits.
- A centralized database has a single point of failure. Because of this, its probability of not being accessible is much higher.

## Advantages and disadvantages of a distributed database
### Advantages
- It’s fast and easy to access data in a distributed database because data is retrieved from the nearest database shard or the one frequently used.
- Data with different levels of distribution transparency can be stored in separate places.
- Intensive transactions consisting of queries can be divided into multiple optimized subqueries, which can be processed in a parallel fashion.
### Disadvantages
- Sometimes, data is required from multiple sites, which takes more time than expected.
- Relations are partitioned vertically or horizontally among different nodes. Therefore, operations such as joins need to reconstruct complete relations by carefully fetching data. These operations can become much more expensive and complex.
- It’s difficult to maintain consistency of data across sites in the distributed database, and it requires extra measures.
- Updations and backups in distributed databases take time to synchronize data.