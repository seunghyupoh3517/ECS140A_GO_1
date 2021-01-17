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
	//size int
	parent *Set
	//DisjointSet interface{}

}

// function MakeSet(x) is
//     if x is not already in the forest then
//         x.parent := x
//         x.size := 1     // if nodes store size
//         x.rank := 0     // if nodes store rank
//     end if
// end function

func MakeSet(x *Set) *Set{
	s :=new(Set)
	s.parent = s
	//s.size = 1
	s.rank = 0
	//s.value = x
	return s
} 
// function Find(x) is
//     if x.parent ≠ x then
//         x.parent := Find(x.parent)
//         return x.parent
//     else
//         return x
//     end if
// end function
func FindSet(x *Set) *Set{
	if x.parent != x{
		x.parent = FindSet(x.parent)
	}
	
	return x.parent
	
}


//https://zh.wikipedia.org/wiki/%E5%B9%B6%E6%9F%A5%E9%9B%86
func UnionSet(x, y *Set){
	Node_X := FindSet(x)
	Node_Y := FindSet(y)
	if Node_X == Node_Y{
		return //auto return NodeX && NodeY
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



// TODO: implement a type that satisfies the DisjointSet interface.

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.
func NewDisjointSet() DisjointSet {
	// panic("TODO: implement this!")

	s :=new(Set)
	s.parent = s
	//s.size = 1
	s.rank = 0
	//s.value = x
	// a :=FindSet(s)
	// b :=FindSet(s)
	// NewSet := UnionSet(a,b)

	return s

}
