![code coverage badge](https://github.com/luka2220/memcached-cli-client/actions/workflows/ci.yaml/badge.svg)

## Description

This project aims to create a cli client to interact with the Memcached server. The client will send commands using the Memcached TCP communication protocol.

Memcached is a free, open-source, high-performance, distributed memory object caching system. It is intended to speed up dynamic web applications by reducing database load. 
It is also an in-memory key-value store for small chunks of arbitrary data retrieved from back-end systems with higher latency. It is simple yet powerful.
Its simple design promotes quick deployment, and ease of development, and solves many problems facing large data caches. Its relatively simple API is available for most popular languages.
It uses a simple text-based network protocol, making it a great platform to learn how to build network clients and servers.

## Specification

Check out the [Memcached protocol specification](https://github.com/memcached/memcached/blob/master/doc/protocol.txt) for more in-depth detail and explanation for each command and response

## Commands

- host flag (-o): specifies the host where the server is running, default is localhost
- port flags (-p): specifies the port number the server is running, default is 11211

**Set** store the data associated with the key <br>

- _ccmc -o lukapiplica.net -p 4206969 set testkey testvalue_
- the first flag is the host name
- the second flags is the address
- the first argument is the key associated with the value
- the second argument is the value to store

**Get** retrieve the data associated with the key <br>

- _ccmc get testkey_

**Add** stores the data, but only if the server doesn't already hold data for this key <br>

- _ccmc add testkey "some value, this one is a long string"_

**Replace** stores the data, but only if the server does already hold data for this key <br>

- _ccmc replace testkey "some value, this one is a long string to replace the current value with"_

**Append** stores the data to an existing key after existing data <br>

- _ccmc append testkey "value to append"_

**Prepend** stores the data to an existing key before existing data <br>

- _ccmc append testkey "value to prepend"_

**Cas** check and set operation which means "store this data but only if no one else has updated since I last fetched it." <br>

- _ccmc cas testkey 250 1_

**Gets** gets the value associated with a key, also retrieves a CAS token (Check-And-Set token) <br>

- _ccmc gets testkey_

**Delete** deletes the given key and any data associated with it <br>

- _ccmc delete testkey_

**Increment** increments a key (must have an integer value or will not work) by the amount passed <br>

- _ccmc incr testkey 10_

**Decrement** decrements a key (must have an integer value or will not work) by the amount passed <br>

- _ccmc decr testkey 100_
  

