
# devourKV

A **Key-Value** store to explore on database and storage internals in **Golang**.


## What is a Key:Value store ?

A key–value database, or key–value store, is a **data storage paradigm** designed for `storing`, `retrieving`, and `managing` associative arrays, and a data structure more commonly known today as a **dictionary** or **hash table**. 
Dictionaries contain a collection of **objects**, or **records**, which in turn have different fields within them, each containing data. 
These records are stored and retrieved using a key that uniquely identifies the record, and is used to find the data within the database.


## Overview

This key-value store is a small version of a database. The inspiration for this project has been taken from the book
**[Introduction to Algorithms](http://staff.ustc.edu.cn/~csli/graduate/algorithms/book6/toc.htm)** popularly
know as `CLRS`. The objective is to design a configurable `key-value` store having different implementations for each of the data 
structure listed below:

1. **B Tree**
2. **B+ Tree**
3. **Skip List**

Database tuples will be stored in `main memory` as well as `disk` depending upon the implementation of the data structure.
Tuples are being stored in main memory for `quick access` and in the disk to add some sort of `persistance` to the database.

For the user to consume this, he needs to pass a `config object` based on which the data structure for the store
will be chosen.


## Features

- **Insert** a `key-value` pair into the store
- **Retrieve** `value` for a particular `key` from the store
- **Remove** a `key-value` pair from the store

## API Reference


#### Get `value` for a given `key`

```go
  Get(key interface{}) (interface{}, error)
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `key` | `interface{}` | **Required**. Unique key associated with the value |

<br>

#### Insert a `key-value` pair

```http
  Store(key interface{}, value interface{}) error
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `key`      | `interface{}` | **Required**. Key to be inserted into the store |
| `value`   |   `interface{}`| **Required**. Value associated with the given key |

<br>

#### Delete a `key-value` pair

```go
  Delete(key interface{}) error
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `key` | `interface{}` | **Required**. Unique key associated with the value |

## Installing

Using `devourKV` is easy. First, use **go get** to install the latest version of the library. 
```go
go get -u github.com/DevourTech/devourKV
```
This command will install `devourKV`  along with the library and its dependencies.


<br>

Next, include **devourKV** in your application:

```go
import "github.com/DevourTech/devourKV/apis"
```

## Important !!

Consider versions released only on the `master` branch as **stable**. The one's released on develop branch are mere snapshots and they do not guarantee stability. 
Snapshots are released using a **GitHub action** called [Snapshotter](https://github.com/DevourTech/snapshotter).

## Contributing

Contributions are always welcome!

Refer [CONTRIBUTING.md](./CONTRIBUTING.md) for ways to get started.


## Authors

- [@rohan23chhabra](https://github.com/rohan23chhabra)
- [@archit-1997](https://github.com/archit-1997)

  