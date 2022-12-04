package main


type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

type Traversal interface {
	Initialize() []int
	PrintTraversal(stack []*Tree, res []int, visited map[int]bool) []int
	util(stack []*Tree, res []int, visited map[int]bool) ([]int, []*Tree)
}

type defImplement struct {
}

func (*defImplement) Initialize() []int {
	return []int{}
}

func (*defImplement) PrintTraversal(stack []*Tree, res []int, visited map[int]bool) []int {
	return []int{}
}

func (*defImplement) util(stack []*Tree, res []int, visited map[int]bool) ([]int, []*Tree) {
	visited[stack[len(stack)-1].Value] = true
	res = append(res, stack[len(stack)-1].Value)
	stack = stack[:len(stack)-1]
	return res, stack
}

// Inorder DFS with InOrder, embedding Traversal interface
// we will intialize this struct with defImplement as Traversal
// type to  inherit util implementation
type Inorder struct {
	tree *Tree
	Traversal
}

func (ord *Inorder) Initialize() []int {
	stack := []*Tree{}
	stack = append(stack, ord.tree)
	result := []int{}
	visited := map[int]bool{}
	return ord.PrintTraversal(stack, result, visited)
}

func (ord *Inorder) PrintTraversal(stack []*Tree, res []int, visited map[int]bool) []int {
	if len(stack) == 0 {
		return res
	}

	subTree := stack[len(stack)-1] // accessing last element
	switch {
	case subTree.Left != nil && visited[subTree.Left.Value] != true:
		stack = append(stack, subTree.Left)
		return ord.PrintTraversal(stack, res, visited)
	case subTree.Right != nil && visited[subTree.Right.Value] != true:
		res, stack = ord.util(stack, res, visited)
		stack = append(stack, subTree.Right)
		return ord.PrintTraversal(stack, res, visited)
	default:
		res, stack = ord.util(stack, res, visited)
		return ord.PrintTraversal(stack, res, visited)
	}
}

type Preorder struct {
	tree *Tree
	Traversal
}

func (ord *Preorder) Initialize() []int {
	stack := []*Tree{}
	stack = append(stack, ord.tree)
	result := []int{ord.tree.Value}
	visited := map[int]bool{
		ord.tree.Value: true,
	}

	return ord.PrintTraversal(stack, result, visited)

}

func (ord *Preorder) PrintTraversal(stack []*Tree, res []int, visited map[int]bool) []int {
	if len(stack) == 0 {
		return res
	}

	subTree := stack[len(stack)-1] //acessing last element
	switch {
	case subTree.Left != nil && visited[subTree.Left.Value] != true:

		res = append(res, subTree.Left.Value)
		visited[subTree.Left.Value] = true
		stack = append(stack, subTree.Left)
		return ord.PrintTraversal(stack, res, visited)

	case subTree.Right != nil && visited[subTree.Right.Value] != true:
		res = append(res, subTree.Right.Value)
		visited[subTree.Right.Value] = true
		stack = append(stack, subTree.Right)
		return ord.PrintTraversal(stack, res, visited)

	default:
		stack = stack[:len(stack)-1]
		return ord.PrintTraversal(stack, res, visited)

	}

}

func main() {

	// numbers := []int{53, 23, 19, 5776, 170, 223, 45, 75, 90, 802, 63, 29, 3, 887, 456, 24, 2, 21, 34, 49, 6555}

	
}