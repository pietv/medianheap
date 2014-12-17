MedianHeap [![GoDoc](https://godoc.org/github.com/pietv/medianheap?status.png)](https://godoc.org/github.com/pietv/medianheap) [![Build Status](https://drone.io/github.com/pietv/medianheap/status.png)](https://drone.io/github.com/pietv/medianheap/latest)[![Build status](https://ci.appveyor.com/api/projects/status/cbmoigwvylv8p1kk/branch/master?svg=true)](https://ci.appveyor.com/project/pietv/medianheap/branch/master)
==========

Implementation of the Running Median algorithm. The provided operations are for adding elements and
retrieving the median. The time complexity for updating the median is O(log N), retrieving it O(1).

Install
=======

```shell
$ go get github.com/pietv/medianheap
```

