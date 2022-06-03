package graph

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readln(r *bufio.Reader) (string, error) {
	var err error
	var line, ln []byte
	isPrefix := true

	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func safeParseUInt(str string) uint {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("ParseInt Error (%s): %s\n", str, err)
		os.Exit(0)
	} else if num < 0 {
		fmt.Printf("ParseInt Error (%s): number less than 0\n", str)
		os.Exit(0)
	}
	return uint(num)
}

func ParseFile(filePath string) *Map {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Can't read file", filePath)
	}

	graphMap := new(Map)

	r := bufio.NewReader(file)
	i := 0
	var amountOfNodes uint = 0
	for line, e := readln(r); e == nil; line, e = readln(r) {
		if i == 0 {
			graphMap.ShortestPathLength = safeParseUInt(line)
		} else if i == 1 {
			amountOfNodes = safeParseUInt(line)
		} else if i == 2 {
			array := strings.Split(line, "-")
			startIndex := safeParseUInt(array[0])
			endIndex := safeParseUInt(array[1])
			graphMap.AllNodes = make([]*Node, amountOfNodes)

			for j := uint(0); j < amountOfNodes; j++ {
				node := new(Node)
				node.Index = j
				node.Type = MIDDLE
				node.Paths = make([]*Path, 0)
				node.MinLengthFound = ^uint(0) // max uint
				if node.Index == startIndex {
					node.Type = START
					graphMap.StartNode = node
				} else if node.Index == endIndex {
					node.Type = END
					graphMap.EndNode = node
				}
				graphMap.AllNodes[j] = node
			}
		} else {
			pathAndLength := strings.Split(line, " ")
			pathNodes := strings.Split(pathAndLength[0], "-")

			nodeIndex0 := safeParseUInt(pathNodes[0])
			nodeIndex1 := safeParseUInt(pathNodes[1])

			path0 := new(Path)
			path0.FromNode = graphMap.AllNodes[nodeIndex0]
			path0.ToNode = graphMap.AllNodes[nodeIndex1]
			path0.Length = safeParseUInt(pathAndLength[1])
			graphMap.AllNodes[nodeIndex0].Paths = append(graphMap.AllNodes[nodeIndex0].Paths, path0)

			path1 := new(Path)
			path1.FromNode = graphMap.AllNodes[nodeIndex1]
			path1.ToNode = graphMap.AllNodes[nodeIndex0]
			path1.Length = safeParseUInt(pathAndLength[1])
			graphMap.AllNodes[nodeIndex1].Paths = append(graphMap.AllNodes[nodeIndex1].Paths, path1)
		}
		i++
	}

	file.Close()
	return graphMap
}
