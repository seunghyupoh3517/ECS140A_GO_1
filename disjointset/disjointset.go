package disjointset

// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
type DisjointSet interface {
	// UnionSet(s, t) merges (unions) the sets containing s and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int
	// FindSet(s) returns representative of the class that s belongs to.
	FindSet(int) int
}

type Set struct{
	rank int
	//num array
	parent *Set
}

// function MakeSet(x) is
//     if x is not already in the forest then
//         x.parent := x
//         x.size := 1     // if nodes store size
//         x.rank := 0     // if nodes store rank
//     end if
// end function

func MakeSet() *Set{
	s := new(Set)
	s.parent = s
	s.rank = 0
	return s
}



func Find(x *Set) *Set{
	if x.parent != x{
		x.parent = Find(x.parent)
	}
	
	return x.parent

}
func FindSet(x int) int{
	
	return x
}

//https://zh.wikipedia.org/wiki/%E5%B9%B6%E6%9F%A5%E9%9B%86
func Union(x , y *Set){
	Node_X := Find(x)
	Node_Y := Find(y)
	if Node_X == Node_Y{
		return 
	}


	if Node_X.rank < Node_Y.rank {
		Node_X.parent = Node_Y
	} else if Node_Y.rank < Node_X.rank {
		Node_Y.parent = Node_X
	} else {
		Node_Y.parent = Node_X 
		Node_X.rank = Node_X.rank + 1
	}

}

func UnionSet(x, y int)int{
	
	return x
}


// TODO: implement a type that satisfies the DisjointSet interface.

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
func NewDisjointSet() DisjointSet {
	// panic("TODO: implement this!")
	var d DisjointSet
	
	//var DisSet = map[int]map[int]int{}  //may not need a nested map
	DisSet := make(map[int]int) 

	//try to find element from each key, after Union change the key name 
	// for key, element := range DisSet{
		
	// }



	return d

}
