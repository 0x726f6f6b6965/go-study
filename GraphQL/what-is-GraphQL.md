# What is GraphQL

## About GraphQL

- GraphQL is a query Language for APIs and a runtime for executing those queries with existing data. Client can get the object data by their query statement.

## Features

1. Declarative data fetching
   - Clients can request exactly the data they want without any redundant field.
2. Single entrypoint
   - The GraphQL api typically expose a single endpoint for all operations, making it more straightforward for clients to interact with the api.
3. Hierarchical Structure
   - The structure of the GraphQL query closely mirrors the structure of the response, allowing clients to specify nested queries and get nested responses.
4. Strong Typing
   - GraphQL are strongly typed, meaning that the schema defines the types of data that can be queried and the relationships between them.
5. Real-time Data with Subscriptions
   - GraphQL supports real-time updates using subscriptions, allowing clients to receive updates when data changes.
6. Documentation
   - The GraphQL schema is self-documenting, and clients can introspect the schema to discover what types and operations are available.

## Different between RESTful and GraphQL

- GraphQL
  - Cons
    1. Caching Complexity
        - Caching can be more complex in GraphQL compared to REST because of the flexibility in the structure of queries.
    2. Potential for Over-fetching
        - While GraphQL allows clients to request only the data they need, there's a risk of developers inadvertently requesting too much data.
- RESTful
  - Cons
    1. Multiple Endpoints
        - RESTful APIs often have multiple endpoints for different resources, leading to a more complex API surface.
    2. Limited Flexibility
        - Adding new features may require changes to the API, impacting existing clients.
    3. Over-fetching and Under-fetching
        - Clients may receive more or less data than they actually need, leading to over-fetching or under-fetching.

## The difference between a query and a muataion

- In GraphQL, a query is similar to a GET request in a RESTful API. It's used to retrieve data from the server, allowing clients to specify the exact structure and fields they need. This helps in avoiding over-fetching of data. On the other hand, a mutation is analogous to a POST, PUT, or DELETE request in RESTful API. Mutations are used when you want to modify data on the server, whether it's creating, updating, or deleting. They enable clients to make changes to the data according to their needs.

## What is a GraphQL schema

- The schema plays a crucial role as it defines the structure of the data that clients can query and the operations they can perform. The schema servers as a contract between the client and the server, outling the types of data available and the relationships between them. By defining a clear schema, the server communicates to the clent what data is accessible and how it can be requested. This not only provides a standardized way to interact with the API but also enables powerful introspection, allowing clients to discover the available tpyes and operations dynamically.

## Summary

- GraphQL is powerful for scenarios where flexibility in data retrieval is crucial, while RESTful APIs may be preferred for simpler, more traditional use cases.
