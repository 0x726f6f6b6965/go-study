# Content delivery network
## About it
- A CDN is a group of geographically distributed proxy servers that help quickly deliver content to the clients by reducing latency and saving bandwidth via the nearest proxy server. CDN mainly stores two types of data: static and dynamic. 
- With CDN, the service can solve some problems.
  1. High latency: CDN reduces the physical distance and latency.
  2. Data-intensive applications: Since the path to the data includes only the ISP and the nearby CDN components, there’s no issue in serving a large number of users through a few CDN components in a specific area.
  3. Scarcity of data center resources: Most of the traffic is handled at the CDN instead of the origin servers.

## Requirements
### Functional requirements
1. Retrieve: CDN should be able to retrieve content from the origin servers.
2. Request: CDN proxy servers should be able to respond to each user’s request in this regard.
3. Delivery: In the case of the push model, the origin servers should be able to send the content to the CDN proxy servers.
4. Search: CDN should be able to execute a search according to the user's query to get cached or find the content stored within CDN infrastructure.
5. Update: In most cases, content comes from the origin server, but it should support updating data from other peers of CDN.
6. Delete: CDN should be possible to delete cached entries from the CDN servers after a certain period.

### Non-functional requirements
1. Performance: Minimizing latency is the most important feature of a CDN.
2. Availability: CDNs are expected to be available at all times because of their effectiveness. 
3. Scalability: CDN should be able to scale horizontally as the requirements increase.
4. Reliability and security: CDN should ensure no single point of failure and must reliably handle massive traffic loads. Furthermore, it should protect hosted content from various attacks.

## Components
1. Routing system: The routing system directs clients to the nearest CDN facility.
2. Scrubber servers: Scrubber servers are used to separate good traffic from malicious traffic and protect against well-known attacks, like DDoS. Scrubber servers are generally used only when an attack is detected.
3. Proxy servers: The proxy or edge proxy servers serve the content from RAM to the users. Proxy servers store hot data in RAM, though they can store cold data in SSD or hard drives as well.
4. Distribution system: The distribution system is responsible for distributing content to all the edge proxy servers to different CDN facilities. This system uses the Internet and intelligent broadcast-like approaches to distribute content across the active edge proxy servers. 
5. Origin servers: The origin servers serve any unavailable data at the CDN to the clients and allow the CDN to retrieve data.
6. Management system: The management systems can observe resource usage and statistics constantly.

## Workflow
1. The origin servers provide the URI namespace delegation of all objects cached in the CDN to the request routing system.
2. The origin servers publish the content to the distribution system responsible for data across the available proxy servers.
3. The distribution system distributes the content among the proxy servers and provides feedback to the request routing system. The feedback helps optimize the selection of the nearest proxy server for a client.
4. The client requests the routing system for a suitable proxy server from the request routing system.
5. The routing system returns the IP address of an appropriate proxy server.
6. The client requests routes through the scrubber servers for security reasons.
7. The scrubber server forwards good traffic to the edge proxy server.
8. The proxy server serves the client's request and periodically forwards accounting information to the management system. However, the request is routed to the origin servers if the content isn’t available in the proxy servers.
9. The management system updates the origin servers and sends feedback to the routing system about the statistics and details of the content. 

## API Design
1. `func retrieve_content(proxy_id string, content_type int, content_version int, description interface{}) interface{}` 
   - The main purpose of this function is to retrieve the content and this is for proxy servers to origin servers. The `proxy_id` is a unique ID of the requesting proxy server, the `content_type` represents the structure of the content, the `content_version` represents the version number of the content, and the `description` specifies the content detail. The function will return a JSON file, which contains the text, content types, links to the images or videos in the content, and so on.
2. `func deliver_content(origin_id string, server_list []string, content_type int, content_version int, description interface{})`
   - The main function is to deliver the specified content from the origin servers to proxy servers. The `origin_id` recognizes each origin server uniquely, the `server_list` identifies the list of servers the content will be pushed to by the distribution system, and the `content_version` represents the updated version of the content at the origin server. The proxy server receiving the content will discard the previous version.
3. `func get_content(user_id string, content_type int, description interface{}) interface{}`
   - The user uses this to request the content from the proxy servers. The `user_id` is the unique ID of the user.
4. `func search_content(proxyID string, content_type int, description interface{}) interface{}`
   - The function allows the proxy server to search the content from other available proxy servers.
5. `func update_content(proxy_id string, content_type int, description interface{})`
   - The function allows the proxy server to update the content to other available proxy servers.

## Content caching 
- Cache content is important in delivering up-to-date and popular web content. There are two types of CDNs are used to get the content from the origin servers: Push CDN, and Pull CDN.
### Push CDN
- Content is sent automatically to the CDN proxy servers from the origin servers in the push CDN. The content delivery to the proxy servers is the responsibility of its provider. The push CDN is suitable for static content. However, if the content is rapidly changing, the push model might struggle to keep up and will do redundant content pushes.
### Pull CDN
- The content will be pulled from origin servers when a user requests unavailable data. The proxy servers will give content time to live, and then remove them from the cache if they're no longer requested to balance capacity and cost. This type of CDN is more suited for serving dynamic content.
### Dynamic content caching optimization
- Create dynamic content by running scripts on proxy servers rather than the origin server.
- Using compression techniques to reduce the communication between the origin server and the proxy servers and storage requirements at proxy servers.
## Multi-tier CDN architecture
- CDNs follow a tree-like structure to ease the data distribution process for the origin server.
- If a child or parent proxy server fails, DNS can route clients to a different child-level proxy server. Each child proxy server knows many upper-layer parent servers, and if one fails, it can go to the other one.
## Nearest proxy server
### DNS redirection
- The DNS can return a URI instead of an IP to the client. This action is called DNS redirect. Content providers can use DNS redirects to send a client to a specific CDN. 
- There are two steps in DNS redirection: 
  1. It maps the clients to the appropriate network location.
  2. It distributes the load over the proxy servers in that location to balance the load between proxy servers.
- To shift a client from one machine in a cluster to another, the DNS replies at the second step are given with short TTLs so that the client repeats the resolution after a short while.
### Anycast
- It is a routing method that provides multiple routing paths for each group of endpoints assigned the same IP address, so a CDN provider can use the anycast mechanism so that clients are directed to the nearest proxy servers for content.
### HTTP redirection
- It's the simplest of all approaches. With this scheme, the client requests content from the origin server. The origin server responds with an HTTP protocol to redirect the user via a URL of the content.
## Content consistency
### Periodic polling
- Using the pull model, proxy servers request the origin server periodically for updated data and change the content in the cache accordingly. 
- Periodic polling uses time-to-refresh (TTR) to adjust the period for requesting updated data from the origin servers.
### Time-to-live
- The TTL defines the expiration time of the content. Each object has a TTL attribute assigned to it by the origin server, and when the TTL expires, the proxy server checks for an update with the origin server. 
### Leases
- The lease denotes the time interval for which the origin server agrees to notify the proxy server if there’s any change in the data. The proxy server must send a message requesting a lease renewal after the expiration of the lease. 
- The lease duration can be optimized dynamically according to the observed load on the proxy servers. This technique is referred to as an adaptive lease.

