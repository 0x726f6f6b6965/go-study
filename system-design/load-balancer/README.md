# Load Balancer

## About it

- Nowadays, there are many applications that need to deal with millions of requests per second. This scenario might cause some problems, such as overloading. To solve this situation, people design the component, the `load balancer`. The main job of the load balancer is to fairly divide all the client's requests to the available servers on the pool. With this, systems can avoid overloading or crashing servers.

## Non-functional

- The load balancer provides three main non-functional features: scalability, availability, and performance.
  1. Scalability: With a load balancer, the capacity of the application can be increased seamlessly by adding servers. Load balancers make such upscaling or downscaling transparent to the end users.
  2. Availability: Even if some servers go down, the load balancers can still find the available ones and forward the requests to the healthy servers. This action also exhibits a characteristic of load balancers, hiding failed servers.
  3. Performance: Load balancers can also forward requests to less loaded servers so users can get faster response times. This not only improves performance but also improves resource utilization.

- There are some additional key services because of load balancers.
  - Health checking
    - Load balancers can use the heartbeat protocol to monitor the health and reliability of end-servers.
  - TSL termination
    - Load balancers can reduce the burden on the end servers by handling TSL termination with the client.
  - Predictive analytics
    - Load balancers can predict traffic patterns by analyzing the traffic passing through them or using traffic statistics obtained over time.
  - Reduced human intervention
    - Because load balancing is automated, the system administration effort required to handle failures is reduced.
  - Service discovery
  - Security
    - Load balancers can mitigate attacks like Dos at different layers of the OSI model.

## Placing load balancer

- Generally, the load balancers sit between clients and servers. Requests go through to servers and back to clients via the load-balancing layer. However, load balancers can be placed in other locations to improve the system's performance, such as between applications and servers or between databases and applications.

- To prevent a single point of failure, the load balancers are usually deployed in pairs as a means of disaster recovery. Generally, to maintain high availability, developers use clusters of load balancers that use heartbeat communication to check the health of load balancers at all times.

