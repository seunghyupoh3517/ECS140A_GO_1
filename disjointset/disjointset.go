package disjointset

// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
// import (
// 	"fmt"
// )
type DisjointSet interface {
	// UnionSet(s, t) merges (unions) the sets containing s and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int
	// FindSet(s) returns representative of the class that s belongs to.
	FindSet(int) int
}



type Set struct{
	// s int
	// t int
	DisSet map[int]int
	Rank int
	
}




// TODO: implement a type that satisfies the DisjointSet interface.

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
func NewDisjointSet() DisjointSet {
	DisSet := make(map[int]int)
	//rank := make(map[int]int)

	var d DisjointSet = Set{DisSet,0}

	
	//var DisSet = map[int]map[int]int{}  //may not need a nested map


	return d

}

func (ss Set) FindSet(x int) int{


	// if x != ss.DisSet[x]{
	// 	ss.DisSet[x] = ss.FindSet(ss.DisSet[x])
	// }
	// return ss.DisSet[x]
	//https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
	// check if a map contains a key
	

	if val, ok := ss.DisSet[x]; ok{ 
		return val
	}

	return x
}



func (ss Set) UnionSet(x, y int)int{
	s := ss.FindSet(x)
	t := ss.FindSet(y)
	if s == t{
		return ss.DisSet[s]
	}

	ss.DisSet[s] = t
	return ss.DisSet[s]
}