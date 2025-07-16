package jsonnet

import (
	"encoding/json"
	"errors"
	"github.com/google/go-jsonnet/ast"
	"github.com/google/go-jsonnet/formatter"
)

func Parse(val string) (any, error) {
	node, _, err := formatter.SnippetToRawAST("main.jsonnet", val)
	if err != nil {
		return nil, err
	}
	b, err := MarshalNode(node)
	if err != nil {
		return nil, err
	}
	res := make(map[string]any)
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Manifest(elem any) (string, error) {
	b, err := json.Marshal(elem)
	if err != nil {
		return "", err
	}
	node, err := UnmarshalNode(b)
	if err != nil {
		return "", err
	}
	res, err := formatter.FormatNode(node, nil, formatter.DefaultOptions())
	if err != nil {
		return "", err
	}
	return res, nil
}

func MarshalNode(node ast.Node) ([]byte, error) {
	wrappedNode := NewNode(node)
	return json.Marshal(wrappedNode)
}

func UnmarshalNode(data []byte) (ast.Node, error) {
	var node Node
	err := json.Unmarshal(data, &node)
	if err != nil {
		return nil, err
	}
	return node.Node, nil
}

type Node struct {
	Node ast.Node
}

func NewNode(node ast.Node) Node {
	return Node{Node: node}
}

func (n Node) MarshalJSON() ([]byte, error) {
	if n.Node == nil {
		return []byte("null"), nil
	}
	var proxy any
	switch v := n.Node.(type) {
	case *ast.Apply:
		proxy = Apply(*v)
	case *ast.ApplyBrace:
		proxy = ApplyBrace(*v)
	case *ast.Array:
		proxy = Array(*v)
	case *ast.ArrayComp:
		proxy = ArrayComp(*v)
	case *ast.Assert:
		proxy = Assert(*v)
	case *ast.Binary:
		proxy = Binary(*v)
	case *ast.Conditional:
		proxy = Conditional(*v)
	case *ast.Dollar:
		proxy = Dollar(*v)
	case *ast.Error:
		proxy = Error(*v)
	case *ast.Function:
		proxy = Function(*v)
	case *ast.Import:
		proxy = Import(*v)
	case *ast.ImportBin:
		proxy = ImportBin(*v)
	case *ast.ImportStr:
		proxy = ImportStr(*v)
	case *ast.InSuper:
		proxy = InSuper(*v)
	case *ast.Index:
		proxy = Index(*v)
	case *ast.LiteralBoolean:
		proxy = LiteralBoolean(*v)
	case *ast.LiteralNull:
		proxy = LiteralNull(*v)
	case *ast.LiteralNumber:
		proxy = LiteralNumber(*v)
	case *ast.LiteralString:
		proxy = LiteralString(*v)
	case *ast.Local:
		proxy = Local(*v)
	case *ast.Object:
		proxy = Object(*v)
	case *ast.ObjectComp:
		proxy = ObjectComp(*v)
	case *ast.Parens:
		proxy = Parens(*v)
	case *ast.Self:
		proxy = Self(*v)
	case *ast.Slice:
		proxy = Slice(*v)
	case *ast.SuperIndex:
		proxy = SuperIndex(*v)
	case *ast.Unary:
		proxy = Unary(*v)
	case *ast.Var:
		proxy = Var(*v)
	default:
	}
	b, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (n *Node) UnmarshalJSON(data []byte) error {
	k := struct {
		Kind string `json:"__kind__"`
	}{}
	err := json.Unmarshal(data, &k)
	if err != nil {
		return err
	}
	if string(data) == "null" {
		return nil
	}
	if k.Kind == "" {
		return errors.New("unknown Node kind")
	}
	var node ast.Node
	switch k.Kind {
	case "Apply":
		node = &Apply{}
	case "ApplyBrace":
		node = &ApplyBrace{}
	case "Array":
		node = &Array{}
	case "ArrayComp":
		node = &ArrayComp{}
	case "Assert":
		node = &Assert{}
	case "Binary":
		node = &Binary{}
	case "Conditional":
		node = &Conditional{}
	case "Dollar":
		node = &Dollar{}
	case "Error":
		node = &Error{}
	case "Function":
		node = &Function{}
	case "Import":
		node = &Import{}
	case "ImportBin":
		node = &ImportBin{}
	case "ImportStr":
		node = &ImportStr{}
	case "InSuper":
		node = &InSuper{}
	case "Index":
		node = &Index{}
	case "LiteralBoolean":
		node = &LiteralBoolean{}
	case "LiteralNull":
		node = &LiteralNull{}
	case "LiteralNumber":
		node = &LiteralNumber{}
	case "LiteralString":
		node = &LiteralString{}
	case "Local":
		node = &Local{}
	case "Object":
		node = &Object{}
	case "ObjectComp":
		node = &ObjectComp{}
	case "Parens":
		node = &Parens{}
	case "Self":
		node = &Self{}
	case "Slice":
		node = &Slice{}
	case "SuperIndex":
		node = &SuperIndex{}
	case "Unary":
		node = &Unary{}
	case "Var":
		node = &Var{}
	default:
		// Handle unknown kind
	}
	err = json.Unmarshal(data, node)
	if err != nil {
		return err
	}

	var astNode ast.Node
	switch v := node.(type) {
	case *Apply:
		n := ast.Apply(*v)
		astNode = &n
	case *ApplyBrace:
		n := ast.ApplyBrace(*v)
		astNode = &n
	case *Array:
		n := ast.Array(*v)
		astNode = &n
	case *ArrayComp:
		n := ast.ArrayComp(*v)
		astNode = &n
	case *Assert:
		n := ast.Assert(*v)
		astNode = &n
	case *Binary:
		n := ast.Binary(*v)
		astNode = &n
	case *Conditional:
		n := ast.Conditional(*v)
		astNode = &n
	case *Dollar:
		n := ast.Dollar(*v)
		astNode = &n
	case *Error:
		n := ast.Error(*v)
		astNode = &n
	case *Function:
		n := ast.Function(*v)
		astNode = &n
	case *Import:
		n := ast.Import(*v)
		astNode = &n
	case *ImportBin:
		n := ast.ImportBin(*v)
		astNode = &n
	case *ImportStr:
		n := ast.ImportStr(*v)
		astNode = &n
	case *InSuper:
		n := ast.InSuper(*v)
		astNode = &n
	case *Index:
		n := ast.Index(*v)
		astNode = &n
	case *LiteralBoolean:
		n := ast.LiteralBoolean(*v)
		astNode = &n
	case *LiteralNull:
		n := ast.LiteralNull(*v)
		astNode = &n
	case *LiteralNumber:
		n := ast.LiteralNumber(*v)
		astNode = &n
	case *LiteralString:
		n := ast.LiteralString(*v)
		astNode = &n
	case *Local:
		n := ast.Local(*v)
		astNode = &n
	case *Object:
		n := ast.Object(*v)
		astNode = &n
	case *ObjectComp:
		n := ast.ObjectComp(*v)
		astNode = &n
	case *Parens:
		n := ast.Parens(*v)
		astNode = &n
	case *Self:
		n := ast.Self(*v)
		astNode = &n
	case *Slice:
		n := ast.Slice(*v)
		astNode = &n
	case *SuperIndex:
		n := ast.SuperIndex(*v)
		astNode = &n
	case *Unary:
		n := ast.Unary(*v)
		astNode = &n
	case *Var:
		n := ast.Var(*v)
		astNode = &n
	default:
	}
	n.Node = astNode
	return nil
}

type IfSpec ast.IfSpec

type ProxyIfSpec struct {
	Kind     string `json:"__kind__"`
	Expr     Node   `json:"expr"`
	IfFodder Fodder `json:"ifFodder"`
}

func (i IfSpec) MarshalJSON() ([]byte, error) {
	proxy := ProxyIfSpec{}
	proxy.Kind = "IfSpec"
	proxy.Expr = NewNode(i.Expr)
	proxy.IfFodder = NewFodder(i.IfFodder)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (i *IfSpec) UnmarshalJSON(data []byte) error {
	var proxy ProxyIfSpec
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	i.Expr = proxy.Expr.Node
	i.IfFodder = proxy.IfFodder.Fodder()
	return nil
}

type ForSpec ast.ForSpec

type ProxyForSpec struct {
	Kind       string         `json:"__kind__"`
	ForFodder  Fodder         `json:"forFodder"`
	VarFodder  Fodder         `json:"varFodder"`
	Conditions []IfSpec       `json:"conditions"`
	Outer      *ForSpec       `json:"outer"`
	Expr       Node           `json:"expr"`
	VarName    ast.Identifier `json:"varName"`
	InFodder   Fodder         `json:"inFodder"`
}

func (f ForSpec) MarshalJSON() ([]byte, error) {
	proxy := ProxyForSpec{}
	proxy.Kind = "ForSpec"
	proxy.ForFodder = NewFodder(f.ForFodder)
	proxy.VarFodder = NewFodder(f.VarFodder)
	proxy.Conditions = make([]IfSpec, len(f.Conditions))
	for i, condition := range f.Conditions {
		proxy.Conditions[i] = IfSpec(condition)
	}
	if f.Outer != nil {
		outer := ForSpec(*f.Outer)
		proxy.Outer = &outer
	}
	proxy.Expr = NewNode(f.Expr)
	proxy.VarName = f.VarName
	proxy.InFodder = NewFodder(f.InFodder)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (f *ForSpec) UnmarshalJSON(data []byte) error {
	var proxy ProxyForSpec
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	f.ForFodder = proxy.ForFodder.Fodder()
	f.VarFodder = proxy.VarFodder.Fodder()
	f.Conditions = make([]ast.IfSpec, len(proxy.Conditions))
	for i, condition := range proxy.Conditions {
		f.Conditions[i] = ast.IfSpec(condition)
	}
	if proxy.Outer != nil {
		outer := ast.ForSpec(*proxy.Outer)
		f.Outer = &outer
	}
	f.Expr = proxy.Expr.Node
	f.VarName = proxy.VarName
	f.InFodder = proxy.InFodder.Fodder()
	return nil
}

type Apply ast.Apply

type ProxyApply struct {
	Kind             string    `json:"__kind__"`
	Target           Node      `json:"target"`
	FodderLeft       Fodder    `json:"fodderLeft"`
	Arguments        Arguments `json:"arguments"`
	FodderRight      Fodder    `json:"fodderRight"`
	TailStrictFodder Fodder    `json:"tailStrictFodder"`
	ProxyNodeBase
	TrailingComma bool `json:"trailingComma"`
	TailStrict    bool `json:"tailStrict"`
}

func (a Apply) MarshalJSON() ([]byte, error) {
	proxy := ProxyApply{}
	proxy.Kind = "Apply"
	proxy.Target = NewNode(a.Target)
	proxy.FodderLeft = NewFodder(a.FodderLeft)
	proxy.Arguments = Arguments(a.Arguments)
	proxy.FodderRight = NewFodder(a.FodderRight)
	proxy.TailStrictFodder = NewFodder(a.TailStrictFodder)
	proxy.ProxyNodeBase = NewProxyNodeBase(a.NodeBase)
	proxy.TrailingComma = a.TrailingComma
	proxy.TailStrict = a.TailStrict
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (a *Apply) UnmarshalJSON(data []byte) error {
	var proxy ProxyApply
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	a.Target = proxy.Target.Node
	a.FodderLeft = proxy.FodderLeft.Fodder()
	a.Arguments = ast.Arguments(proxy.Arguments)
	a.FodderRight = proxy.FodderRight.Fodder()
	a.TailStrictFodder = proxy.TailStrictFodder.Fodder()
	a.NodeBase = proxy.NodeBase()
	a.TrailingComma = proxy.TrailingComma
	a.TailStrict = proxy.TailStrict
	return nil
}

type NamedArgument ast.NamedArgument

type ProxyNamedArgument struct {
	Kind        string         `json:"__kind__"`
	NameFodder  Fodder         `json:"nameFodder"`
	Name        ast.Identifier `json:"name"`
	EqFodder    Fodder         `json:"eqFodder"`
	Arg         Node           `json:"arg"`
	CommaFodder Fodder         `json:"commaFodder"`
}

func (n NamedArgument) MarshalJSON() ([]byte, error) {
	proxy := ProxyNamedArgument{}
	proxy.Kind = "NamedArgument"
	proxy.NameFodder = NewFodder(n.NameFodder)
	proxy.Name = n.Name
	proxy.EqFodder = NewFodder(n.EqFodder)
	proxy.Arg = NewNode(n.Arg)
	proxy.CommaFodder = NewFodder(n.CommaFodder)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (n *NamedArgument) UnmarshalJSON(data []byte) error {
	var proxy ProxyNamedArgument
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	n.NameFodder = proxy.NameFodder.Fodder()
	n.Name = proxy.Name
	n.EqFodder = proxy.EqFodder.Fodder()
	n.Arg = proxy.Arg.Node
	n.CommaFodder = proxy.CommaFodder.Fodder()
	return nil
}

type CommaSeparatedExpr ast.CommaSeparatedExpr

type ProxyCommaSeparatedExpr struct {
	Kind        string `json:"__kind__"`
	Expr        Node   `json:"expr"`
	CommaFodder Fodder `json:"commaFodder"`
}

func (c CommaSeparatedExpr) MarshalJSON() ([]byte, error) {
	proxy := ProxyCommaSeparatedExpr{}
	proxy.Kind = "CommaSeparatedExpr"
	proxy.Expr = NewNode(c.Expr)
	proxy.CommaFodder = NewFodder(c.CommaFodder)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (c *CommaSeparatedExpr) UnmarshalJSON(data []byte) error {
	var proxy ProxyCommaSeparatedExpr
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	c.Expr = proxy.Expr.Node
	c.CommaFodder = proxy.CommaFodder.Fodder()
	return nil
}

type Arguments ast.Arguments

type ProxyArguments struct {
	Positional []CommaSeparatedExpr `json:"positional"`
	Named      []NamedArgument      `json:"named"`
}

func (a Arguments) MarshalJSON() ([]byte, error) {
	proxy := ProxyArguments{}
	proxy.Positional = make([]CommaSeparatedExpr, len(a.Positional))
	for i, positional := range a.Positional {
		proxy.Positional[i] = CommaSeparatedExpr(positional)
	}
	proxy.Named = make([]NamedArgument, len(a.Named))
	for i, named := range a.Named {
		proxy.Named[i] = NamedArgument(named)
	}
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (a *Arguments) UnmarshalJSON(data []byte) error {
	var proxy ProxyArguments
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	a.Positional = make([]ast.CommaSeparatedExpr, len(proxy.Positional))
	for i, positional := range proxy.Positional {
		a.Positional[i] = ast.CommaSeparatedExpr(positional)
	}
	a.Named = make([]ast.NamedArgument, len(proxy.Named))
	for i, named := range proxy.Named {
		a.Named[i] = ast.NamedArgument(named)
	}
	return nil
}

type ApplyBrace ast.ApplyBrace

type ProxyApplyBrace struct {
	Kind  string `json:"__kind__"`
	Left  Node   `json:"left"`
	Right Node   `json:"right"`
	ProxyNodeBase
}

func (a ApplyBrace) MarshalJSON() ([]byte, error) {
	proxy := ProxyApplyBrace{}
	proxy.Kind = "ApplyBrace"
	proxy.Left = NewNode(a.Left)
	proxy.Right = NewNode(a.Right)
	proxy.ProxyNodeBase = NewProxyNodeBase(a.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (a *ApplyBrace) UnmarshalJSON(data []byte) error {
	var proxy ProxyApplyBrace
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	a.Left = proxy.Left.Node
	a.Right = proxy.Right.Node
	a.NodeBase = proxy.NodeBase()
	return nil
}

type Array ast.Array

type ProxyArray struct {
	Kind        string               `json:"__kind__"`
	Elements    []CommaSeparatedExpr `json:"elements"`
	CloseFodder Fodder               `json:"closeFodder"`
	ProxyNodeBase
	TrailingComma bool `json:"trailingComma"`
}

func (a Array) MarshalJSON() ([]byte, error) {
	proxy := ProxyArray{}
	proxy.Kind = "Array"
	proxy.CloseFodder = NewFodder(a.CloseFodder)
	proxy.ProxyNodeBase = NewProxyNodeBase(a.NodeBase)
	proxy.TrailingComma = a.TrailingComma
	proxy.Elements = make([]CommaSeparatedExpr, len(a.Elements))
	for i, element := range a.Elements {
		proxy.Elements[i] = CommaSeparatedExpr(element)
	}
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (a *Array) UnmarshalJSON(data []byte) error {
	var proxy ProxyArray
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	a.Elements = make([]ast.CommaSeparatedExpr, len(proxy.Elements))
	for i, element := range proxy.Elements {
		a.Elements[i] = ast.CommaSeparatedExpr(element)
	}
	a.CloseFodder = proxy.CloseFodder.Fodder()
	a.NodeBase = proxy.NodeBase()
	a.TrailingComma = proxy.TrailingComma
	return nil
}

type ArrayComp ast.ArrayComp

type ProxyArrayComp struct {
	Kind                string  `json:"__kind__"`
	Body                Node    `json:"body"`
	TrailingCommaFodder Fodder  `json:"trailingCommaFodder"`
	Spec                ForSpec `json:"spec"`
	CloseFodder         Fodder  `json:"closeFodder"`
	ProxyNodeBase
	TrailingComma bool `json:"trailingComma"`
}

func (a ArrayComp) MarshalJSON() ([]byte, error) {
	proxy := ProxyArrayComp{}
	proxy.Kind = "ArrayComp"
	proxy.Body = NewNode(a.Body)
	proxy.TrailingCommaFodder = NewFodder(a.TrailingCommaFodder)
	proxy.Spec = ForSpec(a.Spec)
	proxy.CloseFodder = NewFodder(a.CloseFodder)
	proxy.ProxyNodeBase = NewProxyNodeBase(a.NodeBase)
	proxy.TrailingComma = a.TrailingComma
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (a *ArrayComp) UnmarshalJSON(data []byte) error {
	var proxy ProxyArrayComp
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	a.Body = proxy.Body.Node
	a.TrailingCommaFodder = proxy.TrailingCommaFodder.Fodder()
	a.Spec = ast.ForSpec(proxy.Spec)
	a.CloseFodder = proxy.CloseFodder.Fodder()
	a.NodeBase = proxy.NodeBase()
	a.TrailingComma = proxy.TrailingComma
	return nil
}

type Assert ast.Assert

type ProxyAssert struct {
	Kind            string `json:"__kind__"`
	Cond            Node   `json:"cond"`
	Message         Node   `json:"message"`
	Rest            Node   `json:"rest"`
	ColonFodder     Fodder `json:"colonFodder"`
	SemicolonFodder Fodder `json:"semicolonFodder"`
	ProxyNodeBase
}

func (a Assert) MarshalJSON() ([]byte, error) {
	proxy := ProxyAssert{}
	proxy.Kind = "Assert"
	proxy.Cond = NewNode(a.Cond)
	proxy.Message = NewNode(a.Message)
	proxy.Rest = NewNode(a.Rest)
	proxy.ColonFodder = NewFodder(a.ColonFodder)
	proxy.SemicolonFodder = NewFodder(a.SemicolonFodder)
	proxy.ProxyNodeBase = NewProxyNodeBase(a.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (a *Assert) UnmarshalJSON(data []byte) error {
	var proxy ProxyAssert
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	a.Cond = proxy.Cond.Node
	a.Message = proxy.Message.Node
	a.Rest = proxy.Rest.Node
	a.ColonFodder = proxy.ColonFodder.Fodder()
	a.SemicolonFodder = proxy.SemicolonFodder.Fodder()
	a.NodeBase = proxy.NodeBase()
	return nil
}

type Binary ast.Binary

type ProxyBinary struct {
	Kind     string `json:"__kind__"`
	Right    Node   `json:"right"`
	Left     Node   `json:"left"`
	OpFodder Fodder `json:"opFodder"`
	ProxyNodeBase
	Op ast.BinaryOp `json:"op"`
}

func (b Binary) MarshalJSON() ([]byte, error) {
	proxy := ProxyBinary{}
	proxy.Kind = "Binary"
	proxy.Right = NewNode(b.Right)
	proxy.Left = NewNode(b.Left)
	proxy.OpFodder = NewFodder(b.OpFodder)
	proxy.ProxyNodeBase = NewProxyNodeBase(b.NodeBase)
	proxy.Op = b.Op
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (b *Binary) UnmarshalJSON(data []byte) error {
	var proxy ProxyBinary
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	b.Right = proxy.Right.Node
	b.Left = proxy.Left.Node
	b.OpFodder = proxy.OpFodder.Fodder()
	b.NodeBase = proxy.NodeBase()
	b.Op = proxy.Op
	return nil
}

type Conditional ast.Conditional

type ProxyConditional struct {
	Kind        string `json:"__kind__"`
	Cond        Node   `json:"cond"`
	BranchTrue  Node   `json:"branchTrue"`
	BranchFalse Node   `json:"branchFalse"`
	ThenFodder  Fodder `json:"thenFodder"`
	ElseFodder  Fodder `json:"elseFodder"`
	ProxyNodeBase
}

func (c Conditional) MarshalJSON() ([]byte, error) {
	proxy := ProxyConditional{}
	proxy.Kind = "Conditional"
	proxy.Cond = NewNode(c.Cond)
	proxy.BranchTrue = NewNode(c.BranchTrue)
	proxy.BranchFalse = NewNode(c.BranchFalse)
	proxy.ThenFodder = NewFodder(c.ThenFodder)
	proxy.ElseFodder = NewFodder(c.ElseFodder)
	proxy.ProxyNodeBase = NewProxyNodeBase(c.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (c *Conditional) UnmarshalJSON(data []byte) error {
	var proxy ProxyConditional
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	c.Cond = proxy.Cond.Node
	c.BranchTrue = proxy.BranchTrue.Node
	c.BranchFalse = proxy.BranchFalse.Node
	c.ThenFodder = proxy.ThenFodder.Fodder()
	c.ElseFodder = proxy.ElseFodder.Fodder()
	c.NodeBase = proxy.NodeBase()
	return nil
}

type Dollar ast.Dollar

type ProxyDollar struct {
	Kind string `json:"__kind__"`
	ProxyNodeBase
}

func (d Dollar) MarshalJSON() ([]byte, error) {
	proxy := ProxyDollar{}
	proxy.Kind = "Dollar"
	proxy.ProxyNodeBase = NewProxyNodeBase(d.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (d *Dollar) UnmarshalJSON(data []byte) error {
	var proxy ProxyDollar
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	d.NodeBase = proxy.NodeBase()
	return nil
}

type Error ast.Error

type ProxyError struct {
	Kind string `json:"__kind__"`
	Expr Node   `json:"expr"`
	ProxyNodeBase
}

func (e Error) MarshalJSON() ([]byte, error) {
	proxy := ProxyError{}
	proxy.Kind = "Error"
	proxy.Expr = NewNode(e.Expr)
	proxy.ProxyNodeBase = NewProxyNodeBase(e.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (e *Error) UnmarshalJSON(data []byte) error {
	var proxy ProxyError
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	e.Expr = proxy.Expr.Node
	e.NodeBase = proxy.NodeBase()
	return nil
}

type Function ast.Function

type ProxyFunction struct {
	Kind             string      `json:"__kind__"`
	ParenLeftFodder  Fodder      `json:"parenLeftFodder"`
	ParenRightFodder Fodder      `json:"parenRightFodder"`
	Body             Node        `json:"body"`
	Parameters       []Parameter `json:"parameters"`
	ProxyNodeBase
	TrailingComma bool `json:"trailingComma"`
}

func (f Function) MarshalJSON() ([]byte, error) {
	proxy := ProxyFunction{}
	proxy.Kind = "Function"
	proxy.ParenLeftFodder = NewFodder(f.ParenLeftFodder)
	proxy.ParenRightFodder = NewFodder(f.ParenRightFodder)
	proxy.Body = NewNode(f.Body)
	proxy.ProxyNodeBase = NewProxyNodeBase(f.NodeBase)
	proxy.TrailingComma = f.TrailingComma
	proxy.Parameters = make([]Parameter, len(f.Parameters))
	for i, parameter := range f.Parameters {
		proxy.Parameters[i] = Parameter(parameter)
	}
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (f *Function) UnmarshalJSON(data []byte) error {
	var proxy ProxyFunction
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	f.ParenLeftFodder = proxy.ParenLeftFodder.Fodder()
	f.ParenRightFodder = proxy.ParenRightFodder.Fodder()
	f.Body = proxy.Body.Node
	f.Parameters = make([]ast.Parameter, len(proxy.Parameters))
	for i, parameter := range proxy.Parameters {
		f.Parameters[i] = ast.Parameter(parameter)
	}
	f.NodeBase = proxy.NodeBase()
	f.TrailingComma = proxy.TrailingComma
	return nil
}

type Parameter ast.Parameter

type ProxyParameter struct {
	Kind        string            `json:"__kind__"`
	NameFodder  Fodder            `json:"nameFodder"`
	Name        ast.Identifier    `json:"name"`
	CommaFodder Fodder            `json:"commaFodder"`
	EqFodder    Fodder            `json:"eqFodder"`
	DefaultArg  Node              `json:"defaultArg"`
	LocRange    ast.LocationRange `json:"locRange"`
}

func (p Parameter) MarshalJSON() ([]byte, error) {
	proxy := ProxyParameter{}
	proxy.Kind = "Parameter"
	proxy.NameFodder = NewFodder(p.NameFodder)
	proxy.Name = p.Name
	proxy.CommaFodder = NewFodder(p.CommaFodder)
	proxy.EqFodder = NewFodder(p.EqFodder)
	proxy.DefaultArg = NewNode(p.DefaultArg)
	proxy.LocRange = p.LocRange
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (p *Parameter) UnmarshalJSON(data []byte) error {
	var proxy ProxyParameter
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	p.NameFodder = proxy.NameFodder.Fodder()
	p.Name = proxy.Name
	p.CommaFodder = proxy.CommaFodder.Fodder()
	p.EqFodder = proxy.EqFodder.Fodder()
	p.DefaultArg = proxy.DefaultArg.Node
	p.LocRange = proxy.LocRange
	return nil
}

type Import ast.Import

type ProxyImport struct {
	Kind string             `json:"__kind__"`
	File *ast.LiteralString `json:"file"`
	ProxyNodeBase
}

func (i Import) MarshalJSON() ([]byte, error) {
	proxy := ProxyImport{}
	proxy.Kind = "Import"
	proxy.File = i.File
	proxy.ProxyNodeBase = NewProxyNodeBase(i.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (i *Import) UnmarshalJSON(data []byte) error {
	var proxy ProxyImport
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	i.File = proxy.File
	i.NodeBase = proxy.NodeBase()
	return nil
}

type ImportBin ast.ImportBin

type ProxyImportBin struct {
	Kind string             `json:"__kind__"`
	File *ast.LiteralString `json:"file"`
	ProxyNodeBase
}

func (i ImportBin) MarshalJSON() ([]byte, error) {
	proxy := ProxyImportBin{}
	proxy.Kind = "ImportBin"
	proxy.File = i.File
	proxy.ProxyNodeBase = NewProxyNodeBase(i.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (i *ImportBin) UnmarshalJSON(data []byte) error {
	var proxy ProxyImportBin
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	i.File = proxy.File
	i.NodeBase = proxy.NodeBase()
	return nil
}

type ImportStr ast.ImportStr

type ProxyImportStr struct {
	Kind string             `json:"__kind__"`
	File *ast.LiteralString `json:"file"`
	ProxyNodeBase
}

func (i ImportStr) MarshalJSON() ([]byte, error) {
	proxy := ProxyImportStr{}
	proxy.Kind = "ImportStr"
	proxy.File = i.File
	proxy.ProxyNodeBase = NewProxyNodeBase(i.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (i *ImportStr) UnmarshalJSON(data []byte) error {
	var proxy ProxyImportStr
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	i.File = proxy.File
	i.NodeBase = proxy.NodeBase()
	return nil
}

type Index ast.Index

type ProxyIndex struct {
	Kind               string          `json:"__kind__"`
	Target             Node            `json:"target"`
	Index              Node            `json:"index"`
	RightBracketFodder Fodder          `json:"rightBracketFodder"`
	LeftBracketFodder  Fodder          `json:"leftBracketFodder"`
	Id                 *ast.Identifier `json:"id"`
	ProxyNodeBase
}

func (i Index) MarshalJSON() ([]byte, error) {
	proxy := ProxyIndex{}
	proxy.Kind = "Index"
	proxy.Target = NewNode(i.Target)
	proxy.Index = NewNode(i.Index)
	proxy.RightBracketFodder = NewFodder(i.RightBracketFodder)
	proxy.LeftBracketFodder = NewFodder(i.LeftBracketFodder)
	proxy.Id = i.Id
	proxy.ProxyNodeBase = NewProxyNodeBase(i.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (i *Index) UnmarshalJSON(data []byte) error {
	var proxy ProxyIndex
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	i.Target = proxy.Target.Node
	i.Index = proxy.Index.Node
	i.RightBracketFodder = proxy.RightBracketFodder.Fodder()
	i.LeftBracketFodder = proxy.LeftBracketFodder.Fodder()
	i.Id = proxy.Id
	i.NodeBase = proxy.NodeBase()
	return nil
}

type LiteralBoolean ast.LiteralBoolean

type ProxyLiteralBoolean struct {
	Kind  string `json:"__kind__"`
	Value bool   `json:"value"`
	ProxyNodeBase
}

func (l LiteralBoolean) MarshalJSON() ([]byte, error) {
	proxy := ProxyLiteralBoolean{}
	proxy.Kind = "LiteralBoolean"
	proxy.Value = l.Value
	proxy.ProxyNodeBase = NewProxyNodeBase(l.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (l *LiteralBoolean) UnmarshalJSON(data []byte) error {
	var proxy ProxyLiteralBoolean
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	l.Value = proxy.Value
	l.NodeBase = proxy.NodeBase()
	return nil
}

type LiteralNull ast.LiteralNull

type ProxyLiteralNull struct {
	Kind string `json:"__kind__"`
	ProxyNodeBase
}

func (l LiteralNull) MarshalJSON() ([]byte, error) {
	proxy := ProxyLiteralNull{}
	proxy.Kind = "LiteralNull"
	proxy.ProxyNodeBase = NewProxyNodeBase(l.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (l *LiteralNull) UnmarshalJSON(data []byte) error {
	var proxy ProxyLiteralNull
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	l.NodeBase = proxy.NodeBase()
	return nil
}

type LiteralString ast.LiteralString

type ProxyLiteralString struct {
	NodeKind        string `json:"__kind__"`
	Value           string `json:"value"`
	BlockIndent     string `json:"blockIndent"`
	BlockTermIndent string `json:"blockTermIndent"`
	ProxyNodeBase
	Kind ast.LiteralStringKind `json:"kind"`
}

func (l LiteralString) MarshalJSON() ([]byte, error) {
	proxy := ProxyLiteralString{}
	proxy.NodeKind = "LiteralString"
	proxy.Value = l.Value
	proxy.ProxyNodeBase = NewProxyNodeBase(l.NodeBase)
	proxy.Kind = l.Kind
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (l *LiteralString) UnmarshalJSON(data []byte) error {
	var proxy ProxyLiteralString
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	l.Value = proxy.Value
	l.NodeBase = proxy.NodeBase()
	return nil
}

type LiteralNumber ast.LiteralNumber

type ProxyLiteralNumber struct {
	Kind           string `json:"__kind__"`
	OriginalString string `json:"originalString"`
	ProxyNodeBase
}

func (l LiteralNumber) MarshalJSON() ([]byte, error) {
	proxy := ProxyLiteralNumber{}
	proxy.Kind = "LiteralNumber"
	proxy.OriginalString = l.OriginalString
	proxy.ProxyNodeBase = NewProxyNodeBase(l.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (l *LiteralNumber) UnmarshalJSON(data []byte) error {
	var proxy ProxyLiteralNumber
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	l.OriginalString = proxy.OriginalString
	l.NodeBase = proxy.NodeBase()
	return nil
}

type Slice ast.Slice

type ProxySlice struct {
	Kind               string `json:"__kind__"`
	Target             Node   `json:"target"`
	LeftBracketFodder  Fodder `json:"leftBracketFodder"`
	BeginIndex         Node   `json:"beginIndex"`
	EndColonFodder     Fodder `json:"endColonFodder"`
	EndIndex           Node   `json:"endIndex"`
	StepColonFodder    Fodder `json:"stepColonFodder"`
	Step               Node   `json:"step"`
	RightBracketFodder Fodder `json:"rightBracketFodder"`
	ProxyNodeBase
}

func (s Slice) MarshalJSON() ([]byte, error) {
	proxy := ProxySlice{}
	proxy.Kind = "Slice"
	proxy.Target = NewNode(s.Target)
	proxy.LeftBracketFodder = NewFodder(s.LeftBracketFodder)
	proxy.BeginIndex = NewNode(s.BeginIndex)
	proxy.EndColonFodder = NewFodder(s.EndColonFodder)
	proxy.EndIndex = NewNode(s.EndIndex)
	proxy.StepColonFodder = NewFodder(s.StepColonFodder)
	proxy.Step = NewNode(s.Step)
	proxy.RightBracketFodder = NewFodder(s.RightBracketFodder)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (s *Slice) UnmarshalJSON(data []byte) error {
	var proxy ProxySlice
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	s.Target = proxy.Target.Node
	s.LeftBracketFodder = proxy.LeftBracketFodder.Fodder()
	s.BeginIndex = proxy.BeginIndex.Node
	s.EndColonFodder = proxy.EndColonFodder.Fodder()
	s.EndIndex = proxy.EndIndex.Node
	s.StepColonFodder = proxy.StepColonFodder.Fodder()
	s.Step = proxy.Step.Node
	s.RightBracketFodder = proxy.RightBracketFodder.Fodder()
	s.NodeBase = proxy.NodeBase()
	return nil
}

type LocalBind ast.LocalBind

type ProxyLocalBind struct {
	Kind        string         `json:"__kind__"`
	VarFodder   Fodder         `json:"varFodder"`
	Body        Node           `json:"body"`
	EqFodder    Fodder         `json:"eqFodder"`
	Variable    ast.Identifier `json:"variable"`
	CloseFodder Fodder         `json:"closeFodder"`
	Fun         *Function      `json:"fun"`
}

func (l LocalBind) MarshalJSON() ([]byte, error) {
	proxy := ProxyLocalBind{}
	proxy.Kind = "LocalBind"
	proxy.VarFodder = NewFodder(l.VarFodder)
	proxy.Body = NewNode(l.Body)
	proxy.EqFodder = NewFodder(l.EqFodder)
	proxy.Variable = l.Variable
	proxy.CloseFodder = NewFodder(l.CloseFodder)
	if l.Fun != nil {
		fun := Function(*l.Fun)
		proxy.Fun = &fun
	}
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (l *LocalBind) UnmarshalJSON(data []byte) error {
	var proxy ProxyLocalBind
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	l.VarFodder = proxy.VarFodder.Fodder()
	l.Body = proxy.Body.Node
	l.EqFodder = proxy.EqFodder.Fodder()
	l.Variable = proxy.Variable
	l.CloseFodder = proxy.CloseFodder.Fodder()
	if proxy.Fun != nil {
		fun := ast.Function(*proxy.Fun)
		l.Fun = &fun
	}
	return nil
}

type ProxyLocal struct {
	Kind  string      `json:"__kind__"`
	Binds []LocalBind `json:"binds"`
	Body  Node        `json:"body"`
	ProxyNodeBase
}

type Local ast.Local

func (l Local) MarshalJSON() ([]byte, error) {
	proxy := ProxyLocal{}
	proxy.Kind = "Local"
	proxy.Binds = make([]LocalBind, len(l.Binds))
	for i, bind := range l.Binds {
		proxy.Binds[i] = LocalBind(bind)
	}
	proxy.Body = NewNode(l.Body)
	proxy.ProxyNodeBase = NewProxyNodeBase(l.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (l *Local) UnmarshalJSON(data []byte) error {
	var proxy ProxyLocal
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	l.Binds = make([]ast.LocalBind, len(proxy.Binds))
	for i, bind := range proxy.Binds {
		l.Binds[i] = ast.LocalBind(bind)
	}
	l.Body = proxy.Body.Node
	l.NodeBase = proxy.NodeBase()
	return nil
}

type ObjectField ast.ObjectField

type ProxyObjectField struct {
	NodeKind    string              `json:"__kind__"`
	Method      *Function           `json:"method"`
	Id          *ast.Identifier     `json:"id"`
	Fodder2     Fodder              `json:"fodder2"`
	Fodder1     Fodder              `json:"fodder1"`
	OpFodder    Fodder              `json:"opFodder"`
	CommaFodder Fodder              `json:"commaFodder"`
	Expr1       Node                `json:"expr1"`
	Expr2       Node                `json:"expr2"`
	Expr3       Node                `json:"expr3"`
	LocRange    ast.LocationRange   `json:"locRange"`
	Kind        ast.ObjectFieldKind `json:"kind"`
	Hide        ast.ObjectFieldHide
	SuperSugar  bool
}

func (o ObjectField) MarshalJSON() ([]byte, error) {
	proxy := ProxyObjectField{}
	proxy.NodeKind = "ObjectField"
	if o.Method != nil {
		method := Function(*o.Method)
		proxy.Method = &method
	}
	proxy.Id = o.Id
	proxy.Fodder2 = NewFodder(o.Fodder2)
	proxy.Fodder1 = NewFodder(o.Fodder1)
	proxy.OpFodder = NewFodder(o.OpFodder)
	proxy.CommaFodder = NewFodder(o.CommaFodder)
	proxy.Expr1 = NewNode(o.Expr1)
	proxy.Expr2 = NewNode(o.Expr2)
	proxy.Expr3 = NewNode(o.Expr3)
	proxy.LocRange = o.LocRange
	proxy.Kind = o.Kind
	proxy.Hide = o.Hide
	proxy.SuperSugar = o.SuperSugar
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (o *ObjectField) UnmarshalJSON(data []byte) error {
	var proxy ProxyObjectField
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	if proxy.Method != nil {
		method := ast.Function(*proxy.Method)
		o.Method = &method
	}
	o.Id = proxy.Id
	o.Fodder2 = proxy.Fodder2.Fodder()
	o.Fodder1 = proxy.Fodder1.Fodder()
	o.OpFodder = proxy.OpFodder.Fodder()
	o.CommaFodder = proxy.CommaFodder.Fodder()
	o.Expr1 = proxy.Expr1.Node
	o.Expr2 = proxy.Expr2.Node
	o.Expr3 = proxy.Expr3.Node
	o.LocRange = proxy.LocRange
	o.Kind = proxy.Kind
	o.Hide = proxy.Hide
	o.SuperSugar = proxy.SuperSugar
	return nil
}

type Object ast.Object

type ProxyObject struct {
	Kind        string        `json:"__kind__"`
	Fields      []ObjectField `json:"fields"`
	CloseFodder Fodder        `json:"closeFodder"`
	ProxyNodeBase
	TrailingComma bool `json:"trailingComma"`
}

func (o Object) MarshalJSON() ([]byte, error) {
	proxy := ProxyObject{}
	proxy.Kind = "Object"
	proxy.Fields = make([]ObjectField, len(o.Fields))
	for i, field := range o.Fields {
		proxy.Fields[i] = ObjectField(field)
	}
	proxy.CloseFodder = NewFodder(o.CloseFodder)
	proxy.ProxyNodeBase = NewProxyNodeBase(o.NodeBase)
	proxy.TrailingComma = o.TrailingComma
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (o *Object) UnmarshalJSON(data []byte) error {
	var proxy ProxyObject
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	o.Fields = make([]ast.ObjectField, len(proxy.Fields))
	for i, field := range proxy.Fields {
		o.Fields[i] = ast.ObjectField(field)
	}
	o.CloseFodder = proxy.CloseFodder.Fodder()
	o.NodeBase = proxy.NodeBase()
	o.TrailingComma = proxy.TrailingComma
	return nil
}

type ObjectComp ast.ObjectComp

type ProxyObjectComp struct {
	Kind                string        `json:"__kind__"`
	Fields              []ObjectField `json:"fields"`
	TrailingCommaFodder Fodder        `json:"trailingCommaFodder"`
	CloseFodder         Fodder        `json:"closeFodder"`
	Spec                ForSpec       `json:"spec"`
	ProxyNodeBase
	TrailingComma bool `json:"trailingComma"`
}

func (o ObjectComp) MarshalJSON() ([]byte, error) {
	proxy := ProxyObjectComp{}
	proxy.Kind = "ObjectComp"
	proxy.Fields = make([]ObjectField, len(o.Fields))
	for i, field := range o.Fields {
		proxy.Fields[i] = ObjectField(field)
	}
	proxy.TrailingCommaFodder = NewFodder(o.TrailingCommaFodder)
	proxy.CloseFodder = NewFodder(o.CloseFodder)
	proxy.Spec = ForSpec(o.Spec)
	proxy.ProxyNodeBase = NewProxyNodeBase(o.NodeBase)
	proxy.TrailingComma = o.TrailingComma
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (o *ObjectComp) UnmarshalJSON(data []byte) error {
	var proxy ProxyObjectComp
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	o.Fields = make([]ast.ObjectField, len(proxy.Fields))
	for i, field := range proxy.Fields {
		o.Fields[i] = ast.ObjectField(field)
	}
	o.TrailingCommaFodder = proxy.TrailingCommaFodder.Fodder()
	o.CloseFodder = proxy.CloseFodder.Fodder()
	o.Spec = ast.ForSpec(proxy.Spec)
	o.NodeBase = proxy.NodeBase()
	o.TrailingComma = proxy.TrailingComma
	return nil
}

type Parens ast.Parens

type ProxyParens struct {
	Kind        string `json:"__kind__"`
	Inner       Node   `json:"inner"`
	CloseFodder Fodder `json:"closeFodder"`
	ProxyNodeBase
}

func (p Parens) MarshalJSON() ([]byte, error) {
	proxy := ProxyParens{}
	proxy.Kind = "Parens"
	proxy.Inner = NewNode(p.Inner)
	proxy.CloseFodder = NewFodder(p.CloseFodder)
	proxy.ProxyNodeBase = NewProxyNodeBase(p.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (p *Parens) UnmarshalJSON(data []byte) error {
	var proxy ProxyParens
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	p.Inner = proxy.Inner.Node
	p.CloseFodder = proxy.CloseFodder.Fodder()
	p.NodeBase = proxy.NodeBase()
	return nil
}

type Self ast.Self

type ProxySelf struct {
	Kind string `json:"__kind__"`
	ProxyNodeBase
}

func (s Self) MarshalJSON() ([]byte, error) {
	proxy := ProxySelf{}
	proxy.Kind = "Self"
	proxy.ProxyNodeBase = NewProxyNodeBase(s.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (s *Self) UnmarshalJSON(data []byte) error {
	var proxy ProxySelf
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	s.NodeBase = proxy.NodeBase()
	return nil
}

type SuperIndex ast.SuperIndex

type ProxySuperIndex struct {
	Kind      string          `json:"__kind__"`
	IDFodder  Fodder          `json:"idFodder"`
	Index     Node            `json:"index"`
	DotFodder Fodder          `json:"dotFodder"`
	Id        *ast.Identifier `json:"id"`
	ProxyNodeBase
}

func (s SuperIndex) MarshalJSON() ([]byte, error) {
	proxy := ProxySuperIndex{}
	proxy.Kind = "SuperIndex"
	proxy.IDFodder = NewFodder(s.IDFodder)
	proxy.Index = NewNode(s.Index)
	proxy.DotFodder = NewFodder(s.DotFodder)
	proxy.Id = s.Id
	proxy.ProxyNodeBase = NewProxyNodeBase(s.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (s *SuperIndex) UnmarshalJSON(data []byte) error {
	var proxy ProxySuperIndex
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	s.IDFodder = proxy.IDFodder.Fodder()
	s.Index = proxy.Index.Node
	s.DotFodder = proxy.DotFodder.Fodder()
	s.Id = proxy.Id
	s.NodeBase = proxy.NodeBase()
	return nil
}

type InSuper ast.InSuper

type ProxyInSuper struct {
	Kind        string `json:"__kind__"`
	Index       Node   `json:"index"`
	InFodder    Fodder `json:"inFodder"`
	SuperFodder Fodder `json:"superFodder"`
	ProxyNodeBase
}

func (i InSuper) MarshalJSON() ([]byte, error) {
	proxy := ProxyInSuper{}
	proxy.Kind = "InSuper"
	proxy.Index = NewNode(i.Index)
	proxy.InFodder = NewFodder(i.InFodder)
	proxy.SuperFodder = NewFodder(i.SuperFodder)
	proxy.ProxyNodeBase = NewProxyNodeBase(i.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (i *InSuper) UnmarshalJSON(data []byte) error {
	var proxy ProxyInSuper
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	i.Index = proxy.Index.Node
	i.InFodder = proxy.InFodder.Fodder()
	i.SuperFodder = proxy.SuperFodder.Fodder()
	i.NodeBase = proxy.NodeBase()
	return nil
}

type Unary ast.Unary

type ProxyUnary struct {
	Kind string `json:"__kind__"`
	Expr Node   `json:"expr"`
	ProxyNodeBase
	Op ast.UnaryOp `json:"op"`
}

func (u Unary) MarshalJSON() ([]byte, error) {
	proxy := ProxyUnary{}
	proxy.Kind = "Unary"
	proxy.Expr = NewNode(u.Expr)
	proxy.ProxyNodeBase = NewProxyNodeBase(u.NodeBase)
	proxy.Op = u.Op
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (u *Unary) UnmarshalJSON(data []byte) error {
	var proxy ProxyUnary
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	u.Expr = proxy.Expr.Node
	u.NodeBase = proxy.NodeBase()
	u.Op = proxy.Op
	return nil
}

type Var ast.Var

type ProxyVar struct {
	Kind string         `json:"__kind__"`
	Id   ast.Identifier `json:"id"`
	ProxyNodeBase
}

func (v Var) MarshalJSON() ([]byte, error) {
	proxy := ProxyVar{}
	proxy.Kind = "Var"
	proxy.Id = v.Id
	proxy.ProxyNodeBase = NewProxyNodeBase(v.NodeBase)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (v *Var) UnmarshalJSON(data []byte) error {
	var proxy ProxyVar
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	v.Id = proxy.Id
	v.NodeBase = proxy.NodeBase()
	return nil
}

type Fodder []FodderElement

func NewFodder(fodder ast.Fodder) Fodder {
	res := make(Fodder, len(fodder))
	for i, f := range fodder {
		res[i] = FodderElement(f)
	}
	return res
}

func (f Fodder) Fodder() ast.Fodder {
	res := make(ast.Fodder, len(f))
	for i, fe := range f {
		res[i] = ast.FodderElement(fe)
	}
	return res
}

type FodderElement ast.FodderElement

type ProxyFodderElement struct {
	Comment []string       `json:"comment"`
	Kind    ast.FodderKind `json:"kind"`
	Blanks  int            `json:"blanks"`
	Indent  int            `json:"indent"`
}

func (f FodderElement) MarshalJSON() ([]byte, error) {
	proxy := ProxyFodderElement{}
	proxy.Comment = f.Comment
	proxy.Kind = f.Kind
	proxy.Blanks = f.Blanks
	proxy.Indent = f.Indent
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (f *FodderElement) UnmarshalJSON(data []byte) error {
	var proxy ProxyFodderElement
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	f.Comment = proxy.Comment
	f.Kind = proxy.Kind
	f.Blanks = proxy.Blanks
	f.Indent = proxy.Indent
	return nil
}

type ProxyNodeBase struct {
	Fodder   Fodder          `json:"fodder"`
	Ctx      ast.Context     `json:"context"`
	FreeVars ast.Identifiers `json:"freeVars"`
	LocRange LocationRange   `json:"locRange"`
}

func NewProxyNodeBase(n ast.NodeBase) ProxyNodeBase {
	proxy := ProxyNodeBase{}
	proxy.Fodder = NewFodder(n.Fodder)
	proxy.Ctx = n.Ctx
	proxy.FreeVars = n.FreeVars
	proxy.LocRange = NewLocationRange(n.LocRange)
	return proxy
}

func (p ProxyNodeBase) NodeBase() ast.NodeBase {
	n := ast.NodeBase{}
	n.Fodder = p.Fodder.Fodder()
	n.Ctx = p.Ctx
	n.FreeVars = p.FreeVars
	n.LocRange = p.LocRange.LocationRange()
	return n
}

type LocationRange ast.LocationRange

func NewLocationRange(locationRange ast.LocationRange) LocationRange {
	return LocationRange(locationRange)
}

func (l LocationRange) LocationRange() ast.LocationRange {
	return ast.LocationRange(l)
}

type ProxyLocationRange struct {
	File     *Source  `json:"file"`
	FileName string   `json:"fileName"`
	Begin    Location `json:"begin"`
	End      Location `json:"end"`
}

func (l LocationRange) MarshalJSON() ([]byte, error) {
	proxy := ProxyLocationRange{}
	if l.File != nil {
		file := NewSource(*l.File)
		proxy.File = &file
	}
	proxy.FileName = l.FileName
	proxy.Begin = NewLocation(l.Begin)
	proxy.End = NewLocation(l.End)
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (l *LocationRange) UnmarshalJSON(data []byte) error {
	var proxy ProxyLocationRange
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	if proxy.File != nil {
		file := proxy.File.Source()
		l.File = &file
	}
	l.FileName = proxy.FileName
	l.Begin = proxy.Begin.Location()
	l.End = proxy.End.Location()
	return nil
}

type Source ast.Source

func NewSource(source ast.Source) Source {
	return Source(source)
}

func (s Source) Source() ast.Source {
	return ast.Source(s)
}

type ProxySource struct {
	DiagnosticFileName ast.DiagnosticFileName `json:"diagnosticFileName"`
	Lines              []string               `json:"lines"`
}

func (s Source) MarshalJSON() ([]byte, error) {
	proxy := ProxySource{}
	proxy.DiagnosticFileName = s.DiagnosticFileName
	proxy.Lines = s.Lines
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (s *Source) UnmarshalJSON(data []byte) error {
	var proxy ProxySource
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	s.DiagnosticFileName = proxy.DiagnosticFileName
	s.Lines = proxy.Lines
	return nil
}

type Location ast.Location

func NewLocation(location ast.Location) Location {
	return Location(location)
}

func (s Location) Location() ast.Location {
	return ast.Location(s)
}

type ProxyLocation struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

func (s Location) MarshalJSON() ([]byte, error) {
	proxy := ProxyLocation{}
	proxy.Line = s.Line
	proxy.Column = s.Column
	j, err := json.Marshal(proxy)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (s *Location) UnmarshalJSON(data []byte) error {
	var proxy ProxyLocation
	err := json.Unmarshal(data, &proxy)
	if err != nil {
		return err
	}
	s.Line = proxy.Line
	s.Column = proxy.Column
	return nil
}
