package prototype

type INode interface {
	Clone() INode
	Print(s string)
}
