# LoadBalancer
- There are different types of load balancing, static and dynamic.
- A typical static method is Round Robin, which distributes requests evenly.
- A typical dynamic method is Least Connection, which distributes requests to the server with the least number of unprocessed requests.

### Reverse Proxy 
A reverse proxy is a server that sits in front of web servers and forwards client (e.g. web browser) requests to those web servers. Reverse proxies are typically implemented to help increase security, performance, and reliability. In order to better understand how a reverse proxy works and the benefits it can provide, let’s first define what a proxy server is.

#### A forward proxy, often called a proxy, proxy server, or web proxy, is a server that sits in front of a group of client machines. When those computers make requests to sites and services on the Internet, the proxy server intercepts those requests and then communicates with web servers on behalf of those clients, like a middleman.
For example, let’s name 3 computers involved in a typical forward proxy communication:
A: This is a user’s home computer
B: This is a forward proxy server
C: This is a website’s origin server (where the website data is stored)
