# Sequencer
## About it
- In a large distributed system, there are millions of events created by services. If developers want to distinguish these events from each other, a sequencer service is required to generate unique IDs.
## Functional requirements
1. Uniqueness: The service should generate a unique identifier.

## Non-functional requirements
1. Scalability: The ID generation system should generate at least a billion unique IDs per day.
2. Availability: Since multiple events happen even at the level of nanoseconds, our system should generate IDs for all the events that occur.

## Solutions
### UUID
- A simple solution for the system is using UUIDs. However, this approach also has some disadvantages.
  1. Using 128-bit numbers as primary keys makes the primary key indexing slower, which results in slow inserts.
  2. There is a chance of duplication because every service generates unique IDs by themselves.

### Using a database
- Use a central database to provide IDs that are incremented to ensure uniqueness. However, a single point of failure might be a problem for this central database.
- To solve the SPOF situation, we can deploy `n` databases as ID generators and each database increments `n` value to keep the ID unique. This approach is scalable. We can add more servers, and the value of m will be updated accordingly.
- This solution still has some cons.
  1. The scalability is not as good as we think because it has similar problems to the hash function `mod n`.

### Using a range handler
- Instead of having each database in turn increment the value and generate an ID, it would be better to give each database a range number values.
- This approach is scalable, available, and yields user IDs that have no duplicates.
- We might lose a significant range when a server dies and can only provide a new range once it’s live again.

### Unix timestamp
- This approach is simple, scalable, and easy to implement. It also enables multiple servers to handle concurrent requests.
- For two concurrent events, the same timestamp is returned and the same ID can be assigned to them. This way, the IDs are no longer unique.

### Twitter snowflake
- This approach improves Unix timestamp solutions to be unique. The identifier is 64 bits and combines several parts: Sign bit, Timestamp, Worker number, and Sequence number.
  1. Sign bit: A single bit is assigned as a sign bit, and its value will always be zero. It makes the overall number positive.
  2. Timestamp: the timestamp will consume 41 bits and assign these for milliseconds.
  3. Worker number: The worker number is 10 bits so it gives us 1024 worker IDs.
  4. Sequence number: The sequence number is 12 bits so it gives us 4096 unique sequence numbers.
- Twitter Snowflake uses the time stamp as the first component. Therefore, they’re time sortable. The ID generator is highly available as well.
- IDs generated in a dead period are a problem. The dead period is when no request for generating an ID is made to the server. These IDs will be wasted since they take up identifier space. The unique range possible will deplete earlier than expected and create gaps in our global set of user IDs.

### Summary

| solution          | unique | scalable | available | causality |
| ----------------- | ------ | -------- | --------- | --------- |
| UUID              | x      | v        | v         | x         |
| database          | x      | x        | v         | x         |
| range handler     | v      | v        | v         | x         |
| unix timestamp    | x      | weak     | v         | weak      |
| Twitter Snowflake | v      | v        | v         | weak      |
