# Mini Graph

Let's get to the heart of the matter.
By finish this task you will considerably increase your street credibility and thus earn our highest respect.

At Pulsar AI, our main algorithm use graph theory to find the best price across thousand of exchange in the DEFI ecosystem.
That's why you need to master this kind of algorithm.

## Introduction

The principle is simple, you have to solve the problem of the shortest path in a graph.

Multiple map are provided, each map can have an arbitrary number of node and each node can have an arbitrary number of path to any other node.
Moreover, each path have a length (minimum 1).

A map is modeled by a txt file that follow specific format below. A parser is already provided in the package to avoid you this painful task.


## Map format

```txt
SHORTEST_WAY_LEN
AMOUNT_OF_NODE
startIndex-endIndex
nodeIndex-nodeIndex pathLength
nodeIndex-nodeIndex pathLength
nodeIndex-nodeIndex pathLength
[...]
```

Notes:
- `SHORTEST_WAY_LEN` : this is an indication to help you check that your result is correct, if you create maps yourself for your test, make sure it is correct (or put 0).
- `AMOUNT_OF_NODE` : The total number of node the map have.
- `startIndex-endIndex` : indexes of the two nodes between which the search for the shortest path must be done (note that the index start at 0),
  also you cannot set the same node for start & end.
- `nodeIndex0-nodeIndex1 pathLength` : Create a path between two node with a fixed length, this line can appear as many times as possible.
- Every path are bidirectionnal.
- The parser assume that the map isn't corrupted else it may have unexpected behaviors. 
- All maps provided under the `maps` folder have a solution.
- Your program didn't have to manage error in case a map is corrupted or has no solution (this may happen if there is no path available between START & END nodes).


## Project Architecture

Multiples file are already provided under the `graph` folder :
 - `graphmain/main.go`: entry point to run the project
 - `models.go`: Contains `Map`, `Node` & `Path` models
 - `parser.go`: Contains a function `ParseFile` that return a `*Map` for a given file path (used in main.go)
 - `graph.go`: Only contains an empty function `FindShortestPath(graphMap *Map) uint` that you have to fulfill, this function take the `*Map` returned by the parser
   and should travel the graph and return the length of the shortest path


## Parser & Models

As you can see in the `models.go` file, every `Node` & `Path` we use are pointers which is very convenient to work directly on
the same objects without the risk of ending up with an unwanted copy.
`Map` represent the entire Graph, `Node` represent a node of the graph and hold every `Path` to next nodes, `Path` knows from and to which node it comes and its length.
You can add fields on these structs if you want, but it's possible to solve the problem with the fields already present only.

## The algorithm

You have to fulfill the function `FindShortestPath(graphMap *Map) uint` in the `graph.go` file. This function take a `*Map` and should return
the shortest path length between the start & end nodes. As we say before, each path between two node have a length. So, the length of the full path
between start & end nodes is the sum of every intermediate path you use to rush the end node.

You can run the program by using `go run graph/graph.go <path_to_map_file>`

Notes:
- The algorithm speed is also evaluated, you should avoid useless calculation
- You can create as many files as you want in the `graph` package
- Don't hesitate to comment your code & explains your choices
- You can edit others files if you think it's relevant

It's one you now,
may the odds be ever in your favor!