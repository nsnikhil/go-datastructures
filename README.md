## DataStructures in GO

### Motivation

This Project aims to port/build well known and commonly used data structures to GO.

---

#### This library uses reflection for every method call to verify if the type the element to add is same as other elements in the list.

---

### DataStructures this project wish to cover, this list might grow with time:  

#### List
- [x] ArrayList
- [x] Doubly LinkedList
- [ ] Skip List

#### Map
- [ ] Hash Map
- [ ] Linked Hash Map
- [ ] Tree Map

#### Set
- [ ] Set

#### Stack
- [x] Stack

#### Queue
- [x] Linked Queue
- [x] DeQue
- [x] Blocking Queue
- [x] Priority Queue

#### Tree
- [ ] Binary Tree
- [ ] Binary Search Tree
- [ ] AVL Tree
- [ ] Red Black Tree
- [ ] N-Ary Tree
- [ ] Segment Tree
- [x] Trie

#### B-trees
- [ ] B-Tree

#### Heaps
- [x] Heap

#### Graph
- [ ] Graphs
- [ ] Adjacency List
- [ ] Adjacency Matrix
- [ ] Directed Graph
- [ ] Directed Acyclic Graph

#### Hash Based
- [ ] Hash List
- [ ] Hash Table
- [ ] Hash Tree
- [ ] Bloom Filters
- [ ] Compressed Trie
- [ ] Hash Trie

#### Streams
- [ ] Stream

---

### Usage

1. `go get -u github.com/nsnikhil/go-datastructures`

---

### Contributing

1. Fork it (<https://github.com/nsnikhil/go-datastructures>)
2. `make setup`
3. Create your feature branch (`git checkout -b feature/fooBar`)
4. Commit your changes (`git commit -am 'Add some fooBar'`)
5. Push to the branch (`git push origin feature/fooBar`)
6. Create a new Pull Request

---

### Known Issues

1. Multiple implementation of sorting algorithm.
2. Concurrent search dose not provide enough performance benefits and also the benchmark test for the same sometimes gets into infinite loop.
3. In heap one can add two elements even though the comparator type and element type is not same.  

---

### License

 Copyright 2020 Nikhil Soni

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.