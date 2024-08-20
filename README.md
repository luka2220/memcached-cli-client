## Description

This project aims to create a cli client that can interact with the Memcached server. The client will send commands based on the memcached TCP communication protocol.

Memcached is a free, open-source, high-performance, distributed memory object caching system. It is intended to speed up dynamic web applications by reducing database load. 
It is also an in-memory key-value store for small chunks of arbitrary data retrieved from back-end systems with higher latency. It is simple yet powerful.
Its simple design promotes quick deployment, and ease of development, and solves many problems facing large data caches. Its relatively simple API is available for most popular languages.
It uses a simple text-based network protocol, making it a great platform to learn how to build network clients and servers.

## Commands

- host flag (-o): specifies the host where the server is running, default is localhost
- port flags (-p): specifies te port number the server is running, default is 11211

**Set** _store the data associated with the key_ <br>

- example ccmc -o lukapiplica.net -p 4206969 set testkey testvalue
- the first argument is the key associated with the value
- the second argument is the value to store

**Get** _retrieve the data associated with the key if it exists_ <br>

- example ccmc -o lukapiplica.net -p 4206969 get testkey
- the first argument is the key to retrieve data from
  

