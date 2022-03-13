package ast

import (
	"fmt"
	"strconv"
	"unsafe"

	internalparser "github.com/goccy/go-zetasql/internal/ccall/go-zetasql/parser/parser"
)

import "C"

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

type Node interface {
	setRaw(unsafe.Pointer)
	getRaw() unsafe.Pointer
	ID() int
	Kind() Kind
	NumChildren() int
	SingleNodeDebugString() string
	Child(int) Node
	MutableChild(int) Node
	SetParent(Node)
	Parent() Node
	AddChild(Node)
	AddChildFront(Node)
	FindChildIndex(Kind) int
	DebugString(int) string
	MoveStartLocation(int)
	MoveStartLocationBack(int)
	SetStartLocationToEndLocation()
	MoveEndLocationBack(int)
	SetStartLocation(*ParseLocationPoint)
	SetEndLocation(*ParseLocationPoint)
	IsTableExpression() bool
	IsQueryExpression() bool
	IsExpression() bool
	IsType() bool
	IsLeaf() bool
	IsStatement() bool
	IsScriptStatement() bool
	IsLoopStatement() bool
	IsSqlStatement() bool
	IsDdlStatement() bool
	IsCreateStatement() bool
	IsAlterStatement() bool
	ParseLocationRange() *ParseLocationRange
	LocationString() string
}

type ParseLocationPoint struct {
	raw unsafe.Pointer
}

func (p *ParseLocationPoint) getRaw() unsafe.Pointer {
	return p.raw
}

func (p *ParseLocationPoint) Filename() string {
	var s unsafe.Pointer
	internalparser.ParseLocationPoint_filename(p.getRaw(), &s)
	return C.GoString((*C.char)(s))
}

func (p *ParseLocationPoint) ByteOffset() int {
	var offset int
	internalparser.ParseLocationPoint_GetByteOffset(p.getRaw(), &offset)
	return offset
}

func (p *ParseLocationPoint) String() string {
	var s unsafe.Pointer
	internalparser.ParseLocationPoint_GetString(p.getRaw(), &s)
	return C.GoString((*C.char)(s))
}

func newParseLocationPoint(raw unsafe.Pointer) *ParseLocationPoint {
	return &ParseLocationPoint{raw: raw}
}

type ParseLocationRange struct {
	raw unsafe.Pointer
}

func (r *ParseLocationRange) getRaw() unsafe.Pointer {
	return r.raw
}

func (r *ParseLocationRange) Start() *ParseLocationPoint {
	var p unsafe.Pointer
	internalparser.ParseLocationRange_start(r.getRaw(), &p)
	return newParseLocationPoint(p)
}

func (r *ParseLocationRange) End() *ParseLocationPoint {
	var p unsafe.Pointer
	internalparser.ParseLocationRange_end(r.getRaw(), &p)
	return newParseLocationPoint(p)
}

func (r *ParseLocationRange) String() string {
	var s unsafe.Pointer
	internalparser.ParseLocationRange_GetString(r.getRaw(), &s)
	return C.GoString((*C.char)(s))
}

func newParseLocationRange(raw unsafe.Pointer) *ParseLocationRange {
	return &ParseLocationRange{raw: raw}
}

type StatementNode interface {
	Node
}

type ScriptNode interface {
	Node
}

type TypeNode interface {
	Node
}

type ExpressionNode interface {
	Node
}

type LeafNode interface {
	ExpressionNode
}

type QueryExpressionNode interface {
	Node
	SetParenthesized(bool)
	Parenthesized() bool
}

type GeneralizedPathExpressionNode interface {
	ExpressionNode
}

type TableExpressionNode interface {
	Node
}

type TransactionModeNode interface {
	Node
}

type DdlStatementNode interface {
	StatementNode
}

type ColumnAttributeNode interface {
	Node
}

type TableElementNode interface {
	Node
}

type TableConstraintNode interface {
	TableElementNode
}

type AlterActionNode interface {
	Node
}

type BaseNode struct {
	raw unsafe.Pointer
}

func (n *BaseNode) setRaw(p unsafe.Pointer) {
	n.raw = p
}

func (n *BaseNode) getRaw() unsafe.Pointer {
	return n.raw
}

func (n *BaseNode) ID() int {
	var id int
	internalparser.ASTNode_getId(n.getRaw(), &id)
	return id
}

func (n *BaseNode) Kind() Kind {
	var kind int
	internalparser.ASTNode_node_kind(n.getRaw(), &kind)
	return Kind(kind)
}

func (n *BaseNode) NumChildren() int {
	var children int
	internalparser.ASTNode_num_children(n.getRaw(), &children)
	return children
}

func (n *BaseNode) SingleNodeDebugString() string {
	var str unsafe.Pointer
	internalparser.ASTNode_SingleNodeDebugString(n.getRaw(), &str)
	return C.GoString((*C.char)(str))
}

func (n *BaseNode) Child(i int) Node {
	var child unsafe.Pointer
	internalparser.ASTNode_child(n.getRaw(), i, &child)
	return newNode(child)
}

func (n *BaseNode) MutableChild(i int) Node {
	var child unsafe.Pointer
	internalparser.ASTNode_mutable_child(n.getRaw(), i, &child)
	return newNode(child)
}

func (n *BaseNode) SetParent(parent Node) {
	internalparser.ASTNode_set_parent(n.getRaw(), parent.getRaw())
}

func (n *BaseNode) Parent() Node {
	var parent unsafe.Pointer
	internalparser.ASTNode_parent(n.getRaw(), &parent)
	return newNode(parent)
}

func (n *BaseNode) AddChild(child Node) {
	internalparser.ASTNode_AddChild(n.getRaw(), child.getRaw())
}

func (n *BaseNode) AddChildFront(child Node) {
	internalparser.ASTNode_AddChildFront(n.getRaw(), child.getRaw())
}

func (n *BaseNode) FindChildIndex(kind Kind) int {
	var index int
	internalparser.ASTNode_find_child_index(n.getRaw(), int(kind), &index)
	return index
}

func (n *BaseNode) DebugString(maxDepth int) string {
	var str unsafe.Pointer
	internalparser.ASTNode_DebugString(n.getRaw(), maxDepth, &str)
	return C.GoString((*C.char)(str))
}

func (n *BaseNode) MoveStartLocation(bytes int) {
	internalparser.ASTNode_MoveStartLocation(n.getRaw(), bytes)
}

func (n *BaseNode) MoveStartLocationBack(bytes int) {
	internalparser.ASTNode_MoveStartLocationBack(n.getRaw(), bytes)
}

func (n *BaseNode) SetStartLocationToEndLocation() {
	internalparser.ASTNode_SetStartLocationToEndLocation(n.getRaw())
}

func (n *BaseNode) MoveEndLocationBack(bytes int) {
	internalparser.ASTNode_MoveEndLocationBack(n.getRaw(), bytes)
}

func (n *BaseNode) SetStartLocation(point *ParseLocationPoint) {
	internalparser.ASTNode_set_start_location(n.getRaw(), point.getRaw())
}

func (n *BaseNode) SetEndLocation(point *ParseLocationPoint) {
	internalparser.ASTNode_set_end_location(n.getRaw(), point.getRaw())
}

func (n *BaseNode) IsTableExpression() bool {
	var ret bool
	internalparser.ASTNode_IsTableExpression(n.getRaw(), &ret)
	return ret
}

func (n *BaseNode) IsQueryExpression() bool {
	var ret bool
	internalparser.ASTNode_IsQueryExpression(n.getRaw(), &ret)
	return ret
}

func (n *BaseNode) IsExpression() bool {
	var ret bool
	internalparser.ASTNode_IsExpression(n.getRaw(), &ret)
	return ret
}

func (n *BaseNode) IsType() bool {
	var ret bool
	internalparser.ASTNode_IsType(n.getRaw(), &ret)
	return ret
}

func (n *BaseNode) IsLeaf() bool {
	var ret bool
	internalparser.ASTNode_IsLeaf(n.getRaw(), &ret)
	return ret
}

func (n *BaseNode) IsStatement() bool {
	var ret bool
	internalparser.ASTNode_IsStatement(n.getRaw(), &ret)
	return ret
}

func (n *BaseNode) IsScriptStatement() bool {
	var ret bool
	internalparser.ASTNode_IsScriptStatement(n.getRaw(), &ret)
	return ret
}

func (n *BaseNode) IsLoopStatement() bool {
	var ret bool
	internalparser.ASTNode_IsLoopStatement(n.getRaw(), &ret)
	return ret
}

func (n *BaseNode) IsSqlStatement() bool {
	var ret bool
	internalparser.ASTNode_IsSqlStatement(n.getRaw(), &ret)
	return ret
}

func (n *BaseNode) IsDdlStatement() bool {
	var ret bool
	internalparser.ASTNode_IsDdlStatement(n.getRaw(), &ret)
	return ret
}

func (n *BaseNode) IsCreateStatement() bool {
	var ret bool
	internalparser.ASTNode_IsCreateStatement(n.getRaw(), &ret)
	return ret
}

func (n *BaseNode) IsAlterStatement() bool {
	var ret bool
	internalparser.ASTNode_IsAlterStatement(n.getRaw(), &ret)
	return ret
}

func (n *BaseNode) ParseLocationRange() *ParseLocationRange {
	var ret unsafe.Pointer
	internalparser.ASTNode_GetParseLocationRange(n.getRaw(), &ret)
	return newParseLocationRange(ret)
}

func (n *BaseNode) LocationString() string {
	var str unsafe.Pointer
	internalparser.ASTNode_GetLocationString(n.getRaw(), &str)
	return C.GoString((*C.char)(str))
}

type StatementBaseNode struct {
	*BaseNode
}

type QueryStatementNode struct {
	*StatementBaseNode
}

func (n *QueryStatementNode) Query() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTQueryStatement_query(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

type QueryExpressionBaseNode struct {
	*BaseNode
}

func (n *QueryExpressionBaseNode) SetParenthesized(parenthesized bool) {
	internalparser.ASTQueryExpression_set_parenthesized(n.getRaw(), boolToInt(parenthesized))
}

func (n *QueryExpressionBaseNode) Parenthesized() bool {
	var v bool
	internalparser.ASTQueryExpression_parenthesized(n.getRaw(), &v)
	return v
}

type QueryNode struct {
	*QueryExpressionBaseNode
}

func (n *QueryNode) SetIsNested(isNested bool) {
	internalparser.ASTQuery_set_is_nested(n.getRaw(), boolToInt(isNested))
}

func (n *QueryNode) IsNested() bool {
	var v bool
	internalparser.ASTQuery_is_nested(n.getRaw(), &v)
	return v
}

func (n *QueryNode) SetIsPivotInput(isPivotInput bool) {
	internalparser.ASTQuery_set_is_pivot_input(n.getRaw(), boolToInt(isPivotInput))
}

func (n *QueryNode) IsPivotInput() bool {
	var v bool
	internalparser.ASTQuery_is_pivot_input(n.getRaw(), &v)
	return v
}

func (n *QueryNode) WithClause() *WithClauseNode {
	var v unsafe.Pointer
	internalparser.ASTQuery_with_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWithClauseNode(v)
}

func (n *QueryNode) QueryExpr() QueryExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTQuery_query_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(QueryExpressionNode)
}

func (n *QueryNode) OrderBy() *OrderByNode {
	var v unsafe.Pointer
	internalparser.ASTQuery_order_by(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOrderByNode(v)
}

func (n *QueryNode) LimitOffset() *LimitOffsetNode {
	var v unsafe.Pointer
	internalparser.ASTQuery_limit_offset(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newLimitOffsetNode(v)
}

type SelectNode struct {
	*QueryExpressionBaseNode
}

func (n *SelectNode) SetDistinct(distinct bool) {
	internalparser.ASTSelect_set_distinct(n.getRaw(), boolToInt(distinct))
}

func (n *SelectNode) Distinct() bool {
	var v bool
	internalparser.ASTSelect_distinct(n.getRaw(), &v)
	return v
}

func (n *SelectNode) Hint() *HintNode {
	var v unsafe.Pointer
	internalparser.ASTSelect_hint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHintNode(v)
}

func (n *SelectNode) AnonymizationOptions() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTSelect_anonymization_options(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

func (n *SelectNode) SelectAs() *SelectAsNode {
	var v unsafe.Pointer
	internalparser.ASTSelect_select_as(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newSelectAsNode(v)
}

func (n *SelectNode) SelectList() *SelectListNode {
	var v unsafe.Pointer
	internalparser.ASTSelect_select_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newSelectListNode(v)
}

func (n *SelectNode) FromClause() *FromClauseNode {
	var v unsafe.Pointer
	internalparser.ASTSelect_from_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newFromClauseNode(v)
}

func (n *SelectNode) WhereClause() *WhereClauseNode {
	var v unsafe.Pointer
	internalparser.ASTSelect_where_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWhereClauseNode(v)
}

func (n *SelectNode) GroupBy() *GroupByNode {
	var v unsafe.Pointer
	internalparser.ASTSelect_group_by(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newGroupByNode(v)
}

func (n *SelectNode) Having() *HavingNode {
	var v unsafe.Pointer
	internalparser.ASTSelect_having(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHavingNode(v)
}

func (n *SelectNode) Qualify() *QualifyNode {
	var v unsafe.Pointer
	internalparser.ASTSelect_qualify(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQualifyNode(v)
}

func (n *SelectNode) WindowClause() *WindowClauseNode {
	var v unsafe.Pointer
	internalparser.ASTSelect_window_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWindowClauseNode(v)
}

type SelectListNode struct {
	*BaseNode
}

func (n *SelectListNode) Columns() []*SelectColumnNode {
	var num int
	internalparser.ASTSelectList_column_num(n.getRaw(), &num)
	columns := make([]*SelectColumnNode, 0, num)
	for i := 0; i < num; i++ {
		var col unsafe.Pointer
		internalparser.ASTSelectList_column(n.getRaw(), i, &col)
		columns = append(columns, newSelectColumnNode(col))
	}
	return columns
}

type SelectColumnNode struct {
	*BaseNode
}

func (n *SelectColumnNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTSelectColumn_expression(n.getRaw(), &v)
	return newNode(v).(ExpressionNode)
}

func (n *SelectColumnNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTSelectColumn_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

type ExpressionBaseNode struct {
	*BaseNode
}

func (n *ExpressionBaseNode) SetParenthesized(parenthesized bool) {
	internalparser.ASTExpression_set_parenthesized(n.getRaw(), boolToInt(parenthesized))
}

func (n *ExpressionBaseNode) Parenthesized() bool {
	var v bool
	internalparser.ASTExpression_parenthesized(n.getRaw(), &v)
	return v
}

func (n *ExpressionBaseNode) IsAllowedInComparison() bool {
	var v bool
	internalparser.ASTExpression_IsAllowedInComparison(n.getRaw(), &v)
	return v
}

type LeafBaseNode struct {
	*ExpressionBaseNode
}

func (n *LeafBaseNode) Image() string {
	var v unsafe.Pointer
	internalparser.ASTLeaf_image(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

func (n *LeafBaseNode) SetImage(image string) {
	internalparser.ASTLeaf_set_image(n.getRaw(), unsafe.Pointer(C.CString(image)))
}

type IntLiteralNode struct {
	*LeafBaseNode
}

func (n *IntLiteralNode) IsHex() bool {
	var v bool
	internalparser.ASTIntLiteral_is_hex(n.getRaw(), &v)
	return v
}

func (n *IntLiteralNode) SetValue(v int64) {
	n.SetImage(fmt.Sprint(v))
}

func (n *IntLiteralNode) Value() (int64, error) {
	return strconv.ParseInt(n.Image(), 0, 64)
}

type IdentifierNode struct {
	*ExpressionBaseNode
}

func (n *IdentifierNode) SetName(name string) {
	internalparser.ASTIdentifier_SetIdentifier(n.getRaw(), unsafe.Pointer(C.CString(name)))
}

func (n *IdentifierNode) Name() string {
	var name unsafe.Pointer
	internalparser.ASTIdentifier_GetAsString(n.getRaw(), &name)
	return C.GoString((*C.char)(name))
}

type AliasNode struct {
	*BaseNode
}

func (n *AliasNode) Identifier() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTAlias_identifier(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *AliasNode) Name() string {
	var v unsafe.Pointer
	internalparser.ASTAlias_GetAsString(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type GeneralizedPathExpressionBaseNode struct {
	*ExpressionBaseNode
}

// PathExpressionNode is used for dotted identifier paths only, not dotting into
// arbitrary expressions (see DotIdentifierNode).
type PathExpressionNode struct {
	*GeneralizedPathExpressionBaseNode
}

func (n *PathExpressionNode) Names() []*IdentifierNode {
	var num int
	internalparser.ASTPathExpression_num_names(n.getRaw(), &num)
	names := make([]*IdentifierNode, 0, num)
	for i := 0; i < num; i++ {
		var name unsafe.Pointer
		internalparser.ASTPathExpression_name(n.getRaw(), i, &name)
		names = append(names, newIdentifierNode(name))
	}
	return names
}

// ToIdentifierPath String return this PathExpression as a dotted SQL identifier string, with
// quoting if necessary.  If maxPrefixSize is non-zero, include at most
// that many identifiers from the prefix of <path>.
func (n *PathExpressionNode) ToIdentifierPathString(maxPrefixSize uint32) string {
	var v unsafe.Pointer
	internalparser.ASTPathExpression_ToIdentifierPathString(n.getRaw(), maxPrefixSize, &v)
	return C.GoString((*C.char)(v))
}

type TableExpressionBaseNode struct {
	*BaseNode
}

type TablePathExpressionNode struct {
	*TableExpressionBaseNode
}

func (n *TablePathExpressionNode) PathExpr() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTTablePathExpression_path_expr(n.getRaw(), &v)
	return newPathExpressionNode(v)
}

func (n *TablePathExpressionNode) UnnestExpr() *UnnestExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTTablePathExpression_unnest_expr(n.getRaw(), &v)
	return newUnnestExpressionNode(v)
}

func (n *TablePathExpressionNode) Hint() *HintNode {
	var v unsafe.Pointer
	internalparser.ASTTablePathExpression_hint(n.getRaw(), &v)
	return newHintNode(v)
}

func (n *TablePathExpressionNode) WithOffset() *WithOffsetNode {
	var v unsafe.Pointer
	internalparser.ASTTablePathExpression_with_offset(n.getRaw(), &v)
	return newWithOffsetNode(v)
}

func (n *TablePathExpressionNode) PivotClause() *PivotClauseNode {
	var v unsafe.Pointer
	internalparser.ASTTablePathExpression_pivot_clause(n.getRaw(), &v)
	return newPivotClauseNode(v)
}

func (n *TablePathExpressionNode) UnpivotClause() *UnpivotClauseNode {
	var v unsafe.Pointer
	internalparser.ASTTablePathExpression_unpivot_clause(n.getRaw(), &v)
	return newUnpivotClauseNode(v)
}

func (n *TablePathExpressionNode) ForSystemTime() *ForSystemTimeNode {
	var v unsafe.Pointer
	internalparser.ASTTablePathExpression_for_system_time(n.getRaw(), &v)
	return newForSystemTimeNode(v)
}

func (n *TablePathExpressionNode) SampleClause() *SampleClauseNode {
	var v unsafe.Pointer
	internalparser.ASTTablePathExpression_sample_clause(n.getRaw(), &v)
	return newSampleClauseNode(v)
}

func (n *TablePathExpressionNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTTablePathExpression_alias(n.getRaw(), &v)
	return newAliasNode(v)
}

type FromClauseNode struct {
	*BaseNode
}

func (n *FromClauseNode) TableExpression() TableExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTFromClause_table_expression(n.getRaw(), &v)
	return newNode(v).(TableExpressionNode)
}

type WhereClauseNode struct {
	*BaseNode
}

func (n *WhereClauseNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTWhereClause_expression(n.getRaw(), &v)
	return newNode(v).(ExpressionNode)
}

type BooleanLiteralNode struct {
	*LeafBaseNode
}

func (n *BooleanLiteralNode) SetValue(value bool) {
	internalparser.ASTBooleanLiteral_set_value(n.getRaw(), boolToInt(value))
}

func (n *BooleanLiteralNode) Value() bool {
	var v bool
	internalparser.ASTBooleanLiteral_value(n.getRaw(), &v)
	return v
}

type AndExprNode struct {
	*ExpressionBaseNode
}

func (n *AndExprNode) Conjuncts() []ExpressionNode {
	var num int
	internalparser.ASTAndExpr_conjuncts_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTAndExpr_conjunct(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

type BinaryOp int

const (
	NotSetOp BinaryOp = iota
	LikeOp
	IsOp
	EqOp
	NeOp
	Ne2Op
	GtOp
	LtOp
	GeOp
	LeOp
	BitwiseOrOp
	BitwiseXorOp
	BitwiseAndOp
	PlusOp
	MinusOp
	MultiplyOp
	DivideOp
	ConcatOP
	DistinctOp
)

func (o BinaryOp) String() string {
	switch o {
	case NotSetOp:
		return "NOT_SET"
	case LikeOp:
		return "LIKE"
	case IsOp:
		return "IS"
	case EqOp:
		return "EQ"
	case NeOp:
		return "NE"
	case Ne2Op:
		return "NE2"
	case GtOp:
		return "GT"
	case LtOp:
		return "LT"
	case GeOp:
		return "GE"
	case LeOp:
		return "LE"
	case BitwiseOrOp:
		return "BITWISE_OR"
	case BitwiseXorOp:
		return "BITWISE_XOR"
	case BitwiseAndOp:
		return "BITWISE_AND"
	case PlusOp:
		return "PLUS"
	case MinusOp:
		return "MINUS"
	case MultiplyOp:
		return "MULTIPLY"
	case DivideOp:
		return "DIVIDE"
	case ConcatOP:
		return "CONCAT"
	case DistinctOp:
		return "DISTINCT"
	}
	return ""
}

type BinaryExpressionNode struct {
	*ExpressionBaseNode
}

func (n *BinaryExpressionNode) SetOp(op BinaryOp) {
	internalparser.ASTBinaryExpression_set_op(n.getRaw(), int(op))
}

func (n *BinaryExpressionNode) Op() BinaryOp {
	var v int
	internalparser.ASTBinaryExpression_op(n.getRaw(), &v)
	return BinaryOp(v)
}

func (n *BinaryExpressionNode) SetIsNot(isNot bool) {
	internalparser.ASTBinaryExpression_set_is_not(n.getRaw(), boolToInt(isNot))
}

func (n *BinaryExpressionNode) IsNot() bool {
	var v bool
	internalparser.ASTBinaryExpression_is_not(n.getRaw(), &v)
	return v
}

func (n *BinaryExpressionNode) Lhs() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTBinaryExpression_lhs(n.getRaw(), &v)
	return newNode(v).(ExpressionNode)
}

func (n *BinaryExpressionNode) Rhs() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTBinaryExpression_rhs(n.getRaw(), &v)
	return newNode(v).(ExpressionNode)
}

func (n *BinaryExpressionNode) SQLForOperator() string {
	var v unsafe.Pointer
	internalparser.ASTBinaryExpression_GetSQLForOperator(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type StringLiteralNode struct {
	*LeafBaseNode
}

func (n *StringLiteralNode) Value() string {
	var v unsafe.Pointer
	internalparser.ASTStringLiteral_string_value(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

func (n *StringLiteralNode) SetValue(value string) {
	internalparser.ASTStringLiteral_set_string_value(n.getRaw(), unsafe.Pointer(C.CString(value)))
}

type StarNode struct {
	*LeafBaseNode
}

type OrExprNode struct {
	*ExpressionBaseNode
}

func (n *OrExprNode) Disjuncts() []ExpressionNode {
	var num int
	internalparser.ASTOrExpr_disjuncts_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTOrExpr_disjunct(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

type GroupingItemNode struct {
	*BaseNode
}

func (n *GroupingItemNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTGroupingItem_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *GroupingItemNode) Rollup() *RollupNode {
	var v unsafe.Pointer
	internalparser.ASTGroupingItem_rollup(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newRollupNode(v)
}

type GroupByNode struct {
	*BaseNode
}

func (n *GroupByNode) Hint() *HintNode {
	var v unsafe.Pointer
	internalparser.ASTGroupBy_hint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHintNode(v)
}

func (n *GroupByNode) GroupingItems() []*GroupingItemNode {
	var num int
	internalparser.ASTGroupBy_grouping_items_num(n.getRaw(), &num)
	ret := make([]*GroupingItemNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTGroupBy_grouping_item(n.getRaw(), i, &v)
		ret = append(ret, newGroupingItemNode(v))
	}
	return ret
}

type OrderingSpec int

const (
	NotSetSpec OrderingSpec = iota
	AscSpec
	DescSpec
	UnspecifiedSpec
)

func (s OrderingSpec) String() string {
	switch s {
	case NotSetSpec:
		return "NOT_SET"
	case AscSpec:
		return "ASC"
	case DescSpec:
		return "DESC"
	case UnspecifiedSpec:
		return "UNSPECIFIED"
	}
	return ""
}

type OrderingExpressionNode struct {
	*BaseNode
}

func (n *OrderingExpressionNode) SetOrderingSpec(spec OrderingSpec) {
	internalparser.ASTOrderingExpression_set_ordering_spec(n.getRaw(), int(spec))
}

func (n *OrderingExpressionNode) OrderingSpec() OrderingSpec {
	var v int
	internalparser.ASTOrderingExpression_ordering_spec(n.getRaw(), &v)
	return OrderingSpec(v)
}

func (n *OrderingExpressionNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTOrderingExpression_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *OrderingExpressionNode) Collate() *CollateNode {
	var v unsafe.Pointer
	internalparser.ASTOrderingExpression_collate(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newCollateNode(v)
}

func (n *OrderingExpressionNode) NullOrder() *NullOrderNode {
	var v unsafe.Pointer
	internalparser.ASTOrderingExpression_null_order(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNullOrderNode(v)
}

type OrderByNode struct {
	*BaseNode
}

func (n *OrderByNode) Hint() *HintNode {
	var v unsafe.Pointer
	internalparser.ASTOrderBy_hint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHintNode(v)
}

func (n *OrderByNode) OrderingExpressions() []*OrderingExpressionNode {
	var num int
	internalparser.ASTOrderBy_ordering_expressions_num(n.getRaw(), &num)
	ret := make([]*OrderingExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTOrderBy_ordering_expression(n.getRaw(), i, &v)
		ret = append(ret, newOrderingExpressionNode(v))
	}
	return ret
}

type LimitOffsetNode struct {
	*BaseNode
}

func (n *LimitOffsetNode) Limit() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTLimitOffset_limit(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *LimitOffsetNode) Offset() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTLimitOffset_offset(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type FloatLiteralNode struct {
	*LeafBaseNode
}

func (n *FloatLiteralNode) Value() (float64, error) {
	return strconv.ParseFloat(n.Image(), 64)
}

func (n *FloatLiteralNode) SetValue(v float64) {
	n.SetImage(fmt.Sprint(v))
}

type NullLiteralNode struct {
	*LeafBaseNode
}

type OnClauseNode struct {
	*BaseNode
}

func (n *OnClauseNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTOnClause_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type WithClauseEntryNode struct {
	*BaseNode
}

func (n *WithClauseEntryNode) Alias() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTWithClauseEntry_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *WithClauseEntryNode) Query() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTWithClauseEntry_query(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

type JoinType int

const (
	DefaultJoinType JoinType = iota
	CommaJoinType
	CrossJoinType
	FullJoinType
	InnerJoinType
	LeftJoinType
	RightJoinType
)

func (t JoinType) String() string {
	switch t {
	case DefaultJoinType:
		return "DEFAULT_JOIN_TYPE"
	case CommaJoinType:
		return "COMMA"
	case CrossJoinType:
		return "CROSS"
	case FullJoinType:
		return "FULL"
	case InnerJoinType:
		return "INNER"
	case LeftJoinType:
		return "LEFT"
	case RightJoinType:
		return "RIGHT"
	}
	return ""
}

type JoinHint int

const (
	NoJoinHint JoinHint = iota
	HashJoinHint
	LookupJoinHint
)

func (h JoinHint) String() string {
	switch h {
	case NoJoinHint:
		return "NO_JOIN_HINT"
	case HashJoinHint:
		return "HASH"
	case LookupJoinHint:
		return "LOOKUP"
	}
	return ""
}

type JoinNode struct {
	*TableExpressionBaseNode
}

func (n *JoinNode) SetJoinType(typ JoinType) {
	internalparser.ASTJoin_set_join_type(n.getRaw(), int(typ))
}

func (n *JoinNode) JoinType() JoinType {
	var v int
	internalparser.ASTJoin_join_type(n.getRaw(), &v)
	return JoinType(v)
}

func (n *JoinNode) SetJoinHint(hint JoinHint) {
	internalparser.ASTJoin_set_join_hint(n.getRaw(), int(hint))
}

func (n *JoinNode) JoinHint() JoinHint {
	var v int
	internalparser.ASTJoin_join_hint(n.getRaw(), &v)
	return JoinHint(v)
}

func (n *JoinNode) SetNatural(natural bool) {
	internalparser.ASTJoin_set_natural(n.getRaw(), boolToInt(natural))
}

func (n *JoinNode) Natural() bool {
	var v bool
	internalparser.ASTJoin_natural(n.getRaw(), &v)
	return v
}

func (n *JoinNode) SetUnmatchedJoinCount(count int) {
	internalparser.ASTJoin_set_unmatched_join_count(n.getRaw(), count)
}

func (n *JoinNode) UnmatchedJoinCount() int {
	var v int
	internalparser.ASTJoin_unmatched_join_count(n.getRaw(), &v)
	return v
}

func (n *JoinNode) SetTransformationNeeded(needed bool) {
	internalparser.ASTJoin_set_transformation_needed(n.getRaw(), boolToInt(needed))
}

func (n *JoinNode) TransformationNeeded() bool {
	var v bool
	internalparser.ASTJoin_transformation_needed(n.getRaw(), &v)
	return v
}

func (n *JoinNode) SetContainsCommaJoin(commaJoin bool) {
	internalparser.ASTJoin_set_contains_comma_join(n.getRaw(), boolToInt(commaJoin))
}

func (n *JoinNode) ContainsCommaJoin() bool {
	var v bool
	internalparser.ASTJoin_contains_comma_join(n.getRaw(), &v)
	return v
}

func (n *JoinNode) Lhs() TableExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTJoin_lhs(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(TableExpressionNode)
}

func (n *JoinNode) Rhs() TableExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTJoin_rhs(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(TableExpressionNode)
}

func (n *JoinNode) Hint() *HintNode {
	var v unsafe.Pointer
	internalparser.ASTJoin_hint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHintNode(v)
}

func (n *JoinNode) OnClause() *OnClauseNode {
	var v unsafe.Pointer
	internalparser.ASTJoin_on_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOnClauseNode(v)
}

func (n *JoinNode) UsingClause() *UsingClauseNode {
	var v unsafe.Pointer
	internalparser.ASTJoin_using_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newUsingClauseNode(v)
}

type JoinParseError struct {
	raw unsafe.Pointer
}

func newJoinParseError(raw unsafe.Pointer) *JoinParseError {
	return &JoinParseError{raw: raw}
}

func (e *JoinParseError) ErrorNode() Node {
	var v unsafe.Pointer
	internalparser.JoinParseError_error_node(e.raw, &v)
	if v == nil {
		return nil
	}
	return newNode(v)
}

func (e *JoinParseError) Message() string {
	var v unsafe.Pointer
	internalparser.JoinParseError_message(e.raw, &v)
	return C.GoString((*C.char)(v))
}

func (e *JoinParseError) Error() string {
	return e.Message()
}

func (n *JoinNode) ParseError() *JoinParseError {
	var v unsafe.Pointer
	internalparser.ASTJoin_parse_error(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newJoinParseError(v)
}

func (n *JoinNode) SQLForJoinType() string {
	var v unsafe.Pointer
	internalparser.ASTJoin_GetSQLForJoinType(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

func (n *JoinNode) SQLForJoinHint() string {
	var v unsafe.Pointer
	internalparser.ASTJoin_GetSQLForJoinHint(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type WithClauseNode struct {
	*BaseNode
}

func (n *WithClauseNode) SetRecursive(recursive bool) {
	internalparser.ASTWithClause_set_recursive(n.getRaw(), boolToInt(recursive))
}

func (n *WithClauseNode) Recursive() bool {
	var v bool
	internalparser.ASTWithClause_recursive(n.getRaw(), &v)
	return v
}

func (n *WithClauseNode) With() []*WithClauseEntryNode {
	var num int
	internalparser.ASTWithClause_with_num(n.getRaw(), &num)
	ret := make([]*WithClauseEntryNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTWithClause_with(n.getRaw(), i, &v)
		ret = append(ret, newWithClauseEntryNode(v))
	}
	return ret
}

type HavingNode struct {
	*BaseNode
}

func (n *HavingNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTHaving_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type TypeBaseNode struct {
	*BaseNode
}

func (n *TypeBaseNode) TypeParameters() *TypeParameterListNode {
	var v unsafe.Pointer
	internalparser.ASTType_type_parameters(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newTypeParameterListNode(v)
}

func (n *TypeBaseNode) Collate() *CollateNode {
	var v unsafe.Pointer
	internalparser.ASTType_collate(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newCollateNode(v)
}

type SimpleTypeNode struct {
	*TypeBaseNode
}

func (n *SimpleTypeNode) TypeName() string {
	var v unsafe.Pointer
	internalparser.ASTSimpleType_type_name(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type ArrayTypeNode struct {
	*TypeBaseNode
}

func (n *ArrayTypeNode) ElementType() TypeNode {
	var v unsafe.Pointer
	internalparser.ASTArrayType_element_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(TypeNode)
}

type StructFieldNode struct {
	*BaseNode
}

func (n *StructFieldNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTStructField_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *StructFieldNode) Type() TypeNode {
	var v unsafe.Pointer
	internalparser.ASTStructField_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(TypeNode)
}

type StructTypeNode struct {
	*TypeBaseNode
}

func (n *StructTypeNode) StructFields() []*StructFieldNode {
	var num int
	internalparser.ASTStructType_struct_fields_num(n.getRaw(), &num)
	ret := make([]*StructFieldNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTStructType_struct_field(n.getRaw(), i, &v)
		ret = append(ret, newStructFieldNode(v))
	}
	return ret
}

type CastExpressionNode struct {
	*ExpressionBaseNode
}

func (n *CastExpressionNode) SetIsSafeCast(isSafe bool) {
	internalparser.ASTCastExpression_set_is_safe_cast(n.getRaw(), boolToInt(isSafe))
}

func (n *CastExpressionNode) IsSafeCast() bool {
	var v bool
	internalparser.ASTCastExpression_is_safe_cast(n.getRaw(), &v)
	return v
}

func (n *CastExpressionNode) Expr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCastExpression_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *CastExpressionNode) Type() TypeNode {
	var v unsafe.Pointer
	internalparser.ASTCastExpression_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(TypeNode)
}

func (n *CastExpressionNode) Format() *FormatClauseNode {
	var v unsafe.Pointer
	internalparser.ASTCastExpression_format(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newFormatClauseNode(v)
}

type AsMode int

const (
	NotSetMode AsMode = iota
	StructMode
	ValueMode
	TypeNameMode
)

func (m AsMode) String() string {
	switch m {
	case NotSetMode:
		return "NOT_SET"
	case StructMode:
		return "STRUCT"
	case ValueMode:
		return "VALUE"
	case TypeNameMode:
		return "TYPE_NAME"
	}
	return ""
}

type SelectAsNode struct {
	*BaseNode
}

func (n *SelectAsNode) SetAsMode(mode AsMode) {
	internalparser.ASTSelectAs_set_as_mode(n.getRaw(), int(mode))
}

func (n *SelectAsNode) AsMode() AsMode {
	var v int
	internalparser.ASTSelectAs_as_mode(n.getRaw(), &v)
	return AsMode(v)
}

func (n *SelectAsNode) TypeName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTSelectAs_type_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *SelectAsNode) IsSelectAsStruct() bool {
	var v bool
	internalparser.ASTSelectAs_is_select_as_struct(n.getRaw(), &v)
	return v
}

func (n *SelectAsNode) IsSelectAsValue() bool {
	var v bool
	internalparser.ASTSelectAs_is_select_as_value(n.getRaw(), &v)
	return v
}

type RollupNode struct {
	*BaseNode
}

func (n *RollupNode) Expressions() []ExpressionNode {
	var num int
	internalparser.ASTRollup_expressions_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTRollup_expression(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

type NullHandlingModifier int

const (
	DefaultNullHandling NullHandlingModifier = iota
	IgnoreNulls
	RespectNulls
)

func (m NullHandlingModifier) String() string {
	switch m {
	case DefaultNullHandling:
		return "DEFAULT_NULL_HANDLING"
	case IgnoreNulls:
		return "IGNORE_NULLS"
	case RespectNulls:
		return "RESPECT_NULLS"
	}
	return ""
}

type FunctionCallNode struct {
	*ExpressionBaseNode
}

func (n *FunctionCallNode) SetNullHandlingModifier(mod NullHandlingModifier) {
	internalparser.ASTFunctionCall_set_null_handling_modifier(n.getRaw(), int(mod))
}

func (n *FunctionCallNode) NullHandlingModifier() NullHandlingModifier {
	var v int
	internalparser.ASTFunctionCall_null_handling_modifier(n.getRaw(), &v)
	return NullHandlingModifier(v)
}

func (n *FunctionCallNode) SetDistinct(distinct bool) {
	internalparser.ASTFunctionCall_set_distinct(n.getRaw(), boolToInt(distinct))
}

func (n *FunctionCallNode) Distinct() bool {
	var v bool
	internalparser.ASTFunctionCall_distinct(n.getRaw(), &v)
	return v
}

func (n *FunctionCallNode) SetIsCurrentDateTimeWithoutParentheses(v bool) {
	internalparser.ASTFunctionCall_set_is_current_date_time_without_parentheses(n.getRaw(), boolToInt(v))
}

func (n *FunctionCallNode) IsCurrentDateTimeWithoutParentheses() bool {
	var v bool
	internalparser.ASTFunctionCall_is_current_date_time_without_parentheses(n.getRaw(), &v)
	return v
}

func (n *FunctionCallNode) Function() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionCall_function(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *FunctionCallNode) HavingModifier() *HavingModifierNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionCall_having_modifier(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHavingModifierNode(v)
}

func (n *FunctionCallNode) ClampedBetweenModifier() *ClampedBetweenModifierNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionCall_clamped_between_modifier(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newClampedBetweenModifierNode(v)
}

func (n *FunctionCallNode) OrderBy() *OrderByNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionCall_order_by(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOrderByNode(v)
}

func (n *FunctionCallNode) LimitOffset() *LimitOffsetNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionCall_limit_offset(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newLimitOffsetNode(v)
}

func (n *FunctionCallNode) Hint() *HintNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionCall_hint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHintNode(v)
}

func (n *FunctionCallNode) WithGroupRows() *WithGroupRowsNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionCall_with_group_rows(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWithGroupRowsNode(v)
}

func (n *FunctionCallNode) Arguments() []ExpressionNode {
	var num int
	internalparser.ASTFunctionCall_arguments_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTFunctionCall_argument(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

func (n *FunctionCallNode) HasModifiers() bool {
	var v bool
	internalparser.ASTFunctionCall_HasModifiers(n.getRaw(), &v)
	return v
}

type ArrayConstructorNode struct {
	*ExpressionBaseNode
}

func (n *ArrayConstructorNode) Type() *ArrayTypeNode {
	var v unsafe.Pointer
	internalparser.ASTArrayConstructor_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newArrayTypeNode(v)
}

func (n *ArrayConstructorNode) Elements() []ExpressionNode {
	var num int
	internalparser.ASTArrayConstructor_elements_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTArrayConstructor_element(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

type StructConstructorArgNode struct {
	*BaseNode
}

func (n *StructConstructorArgNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTStructConstructorArg_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *StructConstructorArgNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTStructConstructorArg_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

type StructConstructorWithParensNode struct {
	*ExpressionBaseNode
}

func (n *StructConstructorWithParensNode) FieldExpressions() []ExpressionNode {
	var num int
	internalparser.ASTStructConstructorWithParens_field_expressions_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTStructConstructorWithParens_field_expression(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

type StructConstructorWithKeywordNode struct {
	*ExpressionBaseNode
}

func (n *StructConstructorWithKeywordNode) StructType() *StructTypeNode {
	var v unsafe.Pointer
	internalparser.ASTStructConstructorWithKeyword_struct_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStructTypeNode(v)
}

func (n *StructConstructorWithKeywordNode) Fields() []*StructConstructorArgNode {
	var num int
	internalparser.ASTStructConstructorWithKeyword_fields_num(n.getRaw(), &num)
	ret := make([]*StructConstructorArgNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTStructConstructorWithKeyword_field(n.getRaw(), i, &v)
		ret = append(ret, newStructConstructorArgNode(v))
	}
	return ret
}

type InExpressionNode struct {
	*ExpressionBaseNode
}

func (n *InExpressionNode) SetIsNot(isNot bool) {
	internalparser.ASTInExpression_set_is_not(n.getRaw(), boolToInt(isNot))
}

func (n *InExpressionNode) IsNot() bool {
	var v bool
	internalparser.ASTInExpression_is_not(n.getRaw(), &v)
	return v
}

func (n *InExpressionNode) Lhs() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTInExpression_lhs(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *InExpressionNode) Hint() *HintNode {
	var v unsafe.Pointer
	internalparser.ASTInExpression_hint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHintNode(v)
}

func (n *InExpressionNode) InList() *InListNode {
	var v unsafe.Pointer
	internalparser.ASTInExpression_in_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newInListNode(v)
}

func (n *InExpressionNode) Query() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTInExpression_query(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

func (n *InExpressionNode) UnnestExpr() *UnnestExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTInExpression_unnest_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newUnnestExpressionNode(v)
}

type InListNode struct {
	*BaseNode
}

func (n *InListNode) List() []ExpressionNode {
	var num int
	internalparser.ASTInList_list_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTInList_list(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

type BetweenExpressionNode struct {
	*ExpressionBaseNode
}

func (n *BetweenExpressionNode) SetIsNot(isNot bool) {
	internalparser.ASTBetweenExpression_set_is_not(n.getRaw(), boolToInt(isNot))
}

func (n *BetweenExpressionNode) IsNot() bool {
	var v bool
	internalparser.ASTBetweenExpression_is_not(n.getRaw(), &v)
	return v
}

func (n *BetweenExpressionNode) Lhs() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTBetweenExpression_lhs(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *BetweenExpressionNode) Low() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTBetweenExpression_low(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *BetweenExpressionNode) High() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTBetweenExpression_high(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type NumericLiteralNode struct {
	*LeafBaseNode
}

func (n *NumericLiteralNode) Value() string {
	return n.Image()
}

func (n *NumericLiteralNode) SetValue(v string) {
	n.SetImage(v)
}

type BigNumericLiteralNode struct {
	*LeafBaseNode
}

func (n *BigNumericLiteralNode) Value() string {
	return n.Image()
}

func (n *BigNumericLiteralNode) SetValue(v string) {
	n.SetImage(v)
}

type BytesLiteralNode struct {
	*LeafBaseNode
}

func (n *BytesLiteralNode) Value() []byte {
	return []byte(n.Image())
}

func (n *BytesLiteralNode) SetValue(v []byte) {
	n.SetImage(string(v))
}

type DateOrTimeLiteralNode struct {
	*ExpressionBaseNode
}

func (n *DateOrTimeLiteralNode) SetTypeKind(kind TypeKind) {
	internalparser.ASTDateOrTimeLiteral_set_type_kind(n.getRaw(), int(kind))
}

func (n *DateOrTimeLiteralNode) TypeKind() TypeKind {
	var v int
	internalparser.ASTDateOrTimeLiteral_type_kind(n.getRaw(), &v)
	return TypeKind(v)
}

func (n *DateOrTimeLiteralNode) StringLiteral() *StringLiteralNode {
	var v unsafe.Pointer
	internalparser.ASTDateOrTimeLiteral_string_literal(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStringLiteralNode(v)
}

type MaxLiteralNode struct {
	*LeafBaseNode
}

func (n *MaxLiteralNode) Value() string {
	return n.Image()
}

func (n *MaxLiteralNode) SetValue(v string) {
	n.SetImage(v)
}

type JSONLiteralNode struct {
	*LeafBaseNode
}

func (n *JSONLiteralNode) Value() string {
	return n.Image()
}

func (n *JSONLiteralNode) SetValue(v string) {
	n.SetImage(v)
}

type CaseValueExpressionNode struct {
	*ExpressionBaseNode
}

func (n *CaseValueExpressionNode) Arguments() []ExpressionNode {
	var num int
	internalparser.ASTCaseValueExpression_arguments_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTCaseValueExpression_argument(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

type CaseNoValueExpressionNode struct {
	*ExpressionBaseNode
}

func (n *CaseNoValueExpressionNode) Arguments() []ExpressionNode {
	var num int
	internalparser.ASTCaseNoValueExpression_arguments_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTCaseNoValueExpression_argument(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

type ArrayElementNode struct {
	*GeneralizedPathExpressionBaseNode
}

func (n *ArrayElementNode) Array() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTArrayElement_array(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *ArrayElementNode) Position() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTArrayElement_position(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type BitwiseShiftExpressionNode struct {
	*ExpressionBaseNode
}

func (n *BitwiseShiftExpressionNode) SetIsLeftShift(isLeftShift bool) {
	internalparser.ASTBitwiseShiftExpression_set_is_left_shift(n.getRaw(), boolToInt(isLeftShift))
}

func (n *BitwiseShiftExpressionNode) IsLeftShift() bool {
	var v bool
	internalparser.ASTBitwiseShiftExpression_is_left_shift(n.getRaw(), &v)
	return v
}

func (n *BitwiseShiftExpressionNode) Lhs() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTBitwiseShiftExpression_lhs(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *BitwiseShiftExpressionNode) Rhs() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTBitwiseShiftExpression_rhs(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type CollateNode struct {
	*BaseNode
}

func (n *CollateNode) Name() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCollate_collation_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type DotGeneralizedFieldNode struct {
	*GeneralizedPathExpressionBaseNode
}

func (n *DotGeneralizedFieldNode) Expr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDotGeneralizedField_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *DotGeneralizedFieldNode) Path() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDotGeneralizedField_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type DotIdentifierNode struct {
	*GeneralizedPathExpressionBaseNode
}

func (n *DotIdentifierNode) Expr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDotIdentifier_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *DotIdentifierNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTDotIdentifier_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type DotStarNode struct {
	*ExpressionBaseNode
}

func (n *DotStarNode) Expr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDotStar_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type DotStarWithModifiersNode struct {
	*ExpressionBaseNode
}

func (n *DotStarWithModifiersNode) Expr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDotStarWithModifiers_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *DotStarWithModifiersNode) Modifiers() *StarModifiersNode {
	var v unsafe.Pointer
	internalparser.ASTDotStarWithModifiers_modifiers(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStarModifiersNode(v)
}

type ExpressionSubqueryModifier int

const (
	ExpressionSubqueryNone ExpressionSubqueryModifier = iota
	ExpressionSubqueryArray
	ExpressionSubqueryExists
)

func (m ExpressionSubqueryModifier) String() string {
	switch m {
	case ExpressionSubqueryNone:
		return "NONE"
	case ExpressionSubqueryArray:
		return "ARRAY"
	case ExpressionSubqueryExists:
		return "EXISTS"
	}
	return ""
}

type ExpressionSubqueryNode struct {
	*ExpressionBaseNode
}

func (n *ExpressionSubqueryNode) SetModifier(modifier ExpressionSubqueryModifier) {
	internalparser.ASTExpressionSubquery_set_modifier(n.getRaw(), int(modifier))
}

func (n *ExpressionSubqueryNode) Modifier() ExpressionSubqueryModifier {
	var v int
	internalparser.ASTExpressionSubquery_modifier(n.getRaw(), &v)
	return ExpressionSubqueryModifier(v)
}

func (n *ExpressionSubqueryNode) Hint() *HintNode {
	var v unsafe.Pointer
	internalparser.ASTExpressionSubquery_hint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHintNode(v)
}

func (n *ExpressionSubqueryNode) Query() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTExpressionSubquery_query(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

type ExtractExpressionNode struct {
	*ExpressionBaseNode
}

func (n *ExtractExpressionNode) LhsExpr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTExtractExpression_lhs_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *ExtractExpressionNode) RhsExpr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTExtractExpression_rhs_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *ExtractExpressionNode) TimeZoneExpr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTExtractExpression_time_zone_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type HavingModifierKind int

const (
	HavingModifierNotSet HavingModifierKind = iota
	HavingModifierMin
	HavingModifierMax
)

type HavingModifierNode struct {
	*BaseNode
}

func (n *HavingModifierNode) SetModifierKind(kind HavingModifierKind) {
	internalparser.ASTHavingModifier_set_modifier_kind(n.getRaw(), int(kind))
}

func (n *HavingModifierNode) ModifierKind() HavingModifierKind {
	var v int
	internalparser.ASTHavingModifier_modifier_kind(n.getRaw(), &v)
	return HavingModifierKind(v)
}

func (n *HavingModifierNode) Expr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTHavingModifier_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type IntervalExprNode struct {
	*ExpressionBaseNode
}

func (n *IntervalExprNode) InternalValue() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTIntervalExpr_interval_value(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *IntervalExprNode) DatePartName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTIntervalExpr_date_part_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *IntervalExprNode) DatePartNameTo() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTIntervalExpr_date_part_name_to(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type NamedArgumentNode struct {
	*ExpressionBaseNode
}

func (n *NamedArgumentNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTNamedArgument_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *NamedArgumentNode) Expr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTNamedArgument_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type NullOrderNode struct {
	*BaseNode
}

func (n *NullOrderNode) SetNullsFirst(nullsFirst bool) {
	internalparser.ASTNullOrder_set_nulls_first(n.getRaw(), boolToInt(nullsFirst))
}

func (n *NullOrderNode) NullsFirst() bool {
	var v bool
	internalparser.ASTNullOrder_nulls_first(n.getRaw(), &v)
	return v
}

type OnOrUsingClauseListNode struct {
	*BaseNode
}

func (n *OnOrUsingClauseListNode) OnOrUsingClauseList() []Node {
	var num int
	internalparser.ASTOnOrUsingClauseList_on_or_using_clause_list_num(n.getRaw(), &num)
	ret := make([]Node, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTOnUsingClauseList_on_or_using_clause_list(n.getRaw(), i, &v)
		ret = append(ret, newNode(v))
	}
	return ret
}

type ParenthesizedJoinNode struct {
	*TableExpressionBaseNode
}

func (n *ParenthesizedJoinNode) Join() *JoinNode {
	var v unsafe.Pointer
	internalparser.ASTParenthesizedJoin_join(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newJoinNode(v)
}

func (n *ParenthesizedJoinNode) SampleClause() *SampleClauseNode {
	var v unsafe.Pointer
	internalparser.ASTParenthesizedJoin_sample_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newSampleClauseNode(v)
}

type PartitionByNode struct {
	*BaseNode
}

func (n *PartitionByNode) Hint() *HintNode {
	var v unsafe.Pointer
	internalparser.ASTPartitionBy_hint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHintNode(v)
}

func (n *PartitionByNode) PartitioningExpressions() []ExpressionNode {
	var num int
	internalparser.ASTPartitionBy_partitioning_expressions_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTPartitionBy_partitioning_expression(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

type SetOperationType int

const (
	NotSetOperation SetOperationType = iota
	UnionSetOperation
	ExceptSetOperation
	IntersectSetOperation
)

type SetOperationNode struct {
	*QueryExpressionBaseNode
}

func (n *SetOperationNode) SetOpType(opType SetOperationType) {
	internalparser.ASTSetOperation_set_op_type(n.getRaw(), int(opType))
}

func (n *SetOperationNode) OpType() SetOperationType {
	var v int
	internalparser.ASTSetOperation_op_type(n.getRaw(), &v)
	return SetOperationType(v)
}

func (n *SetOperationNode) SetDistinct(distinct bool) {
	internalparser.ASTSetOperation_set_distinct(n.getRaw(), boolToInt(distinct))
}

func (n *SetOperationNode) Distinct() bool {
	var v bool
	internalparser.ASTSetOperation_distinct(n.getRaw(), &v)
	return v
}

func (n *SetOperationNode) Hint() *HintNode {
	var v unsafe.Pointer
	internalparser.ASTSetOperation_hint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHintNode(v)
}

func (n *SetOperationNode) Inputs() []QueryExpressionNode {
	var num int
	internalparser.ASTSetOperation_inputs_num(n.getRaw(), &num)
	ret := make([]QueryExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTSetOperation_input(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(QueryExpressionNode))
	}
	return ret
}

func (n *SetOperationNode) SQLForOperation() string {
	var v unsafe.Pointer
	internalparser.ASTSetOperation_GetSQLForOperation(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type StarExceptListNode struct {
	*BaseNode
}

func (n *StarExceptListNode) Identifiers() []*IdentifierNode {
	var num int
	internalparser.ASTStarExceptList_identifiers_num(n.getRaw(), &num)
	ret := make([]*IdentifierNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTStarExpcetList_identifier(n.getRaw(), i, &v)
		ret = append(ret, newIdentifierNode(v))
	}
	return ret
}

type StarModifiersNode struct {
	*BaseNode
}

func (n *StarModifiersNode) ExceptList() *StarExceptListNode {
	var v unsafe.Pointer
	internalparser.ASTStarModifiers_except_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStarExceptListNode(v)
}

func (n *StarModifiersNode) ReplaceItems() []*StarReplaceItemNode {
	var num int
	internalparser.ASTStarModifiers_replace_items_num(n.getRaw(), &num)
	ret := make([]*StarReplaceItemNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTStarModifiers_replace_item(n.getRaw(), i, &v)
		ret = append(ret, newStarReplaceItemNode(v))
	}
	return ret
}

type StarReplaceItemNode struct {
	*BaseNode
}

func (n *StarReplaceItemNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTStarReplaceItem_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *StarReplaceItemNode) Alias() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTStarReplaceItem_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type StarWithModifiersNode struct {
	*ExpressionBaseNode
}

func (n *StarWithModifiersNode) Modifiers() *StarModifiersNode {
	var v unsafe.Pointer
	internalparser.ASTStarWithModifiers_modifiers(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStarModifiersNode(v)
}

type TableSubqueryNode struct {
	*TableExpressionBaseNode
}

func (n *TableSubqueryNode) Subquery() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTTableSubquery_subquery(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

func (n *TableSubqueryNode) PivotClause() *PivotClauseNode {
	var v unsafe.Pointer
	internalparser.ASTTableSubquery_pivot_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPivotClauseNode(v)
}

func (n *TableSubqueryNode) UnpivotClause() *UnpivotClauseNode {
	var v unsafe.Pointer
	internalparser.ASTTableSubquery_unpivot_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newUnpivotClauseNode(v)
}

func (n *TableSubqueryNode) SampleClause() *SampleClauseNode {
	var v unsafe.Pointer
	internalparser.ASTTableSubquery_sample_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newSampleClauseNode(v)
}

func (n *TableSubqueryNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTTableSubquery_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

type UnaryExpressionOp int

const (
	NotSetUnaryOp UnaryExpressionOp = iota
	NotUnaryOp
	BitwiseNotUnaryOp
	MinusUnaryOp
	PlusUnaryOp
	IsUnknownUnaryOp
	IsNotUnknownUnaryOp
)

func (o UnaryExpressionOp) String() string {
	switch o {
	case NotSetUnaryOp:
		return "NOT_SET"
	case NotUnaryOp:
		return "NOT"
	case BitwiseNotUnaryOp:
		return "BITWISE_NOT"
	case MinusUnaryOp:
		return "MINUS"
	case PlusUnaryOp:
		return "PLUS"
	case IsUnknownUnaryOp:
		return "IS_UNKNOWN"
	case IsNotUnknownUnaryOp:
		return "IS_NOT_UNKNOWN"
	}
	return ""
}

type UnaryExpressionNode struct {
	*ExpressionBaseNode
}

func (n *UnaryExpressionNode) SetOp(op UnaryExpressionOp) {
	internalparser.ASTUnaryExpression_set_op(n.getRaw(), int(op))
}

func (n *UnaryExpressionNode) Op() UnaryExpressionOp {
	var v int
	internalparser.ASTUnaryExpression_op(n.getRaw(), &v)
	return UnaryExpressionOp(v)
}

func (n *UnaryExpressionNode) Operand() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTUnaryExpression_operand(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *UnaryExpressionNode) SQLForOperator() string {
	var v unsafe.Pointer
	internalparser.ASTUnaryExpression_GetSQLForOperator(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type UnnestExpressionNode struct {
	*BaseNode
}

func (n *UnnestExpressionNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTUnnestExpression_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type WindowClauseNode struct {
	*BaseNode
}

func (n *WindowClauseNode) Windows() []*WindowDefinitionNode {
	var num int
	internalparser.ASTWindowClause_windows_num(n.getRaw(), &num)
	ret := make([]*WindowDefinitionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTWindowClause_window(n.getRaw(), i, &v)
		ret = append(ret, newWindowDefinitionNode(v))
	}
	return ret
}

type WindowDefinitionNode struct {
	*BaseNode
}

func (n *WindowDefinitionNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTWindowDefinition_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *WindowDefinitionNode) WindowSpec() *WindowSpecificationNode {
	var v unsafe.Pointer
	internalparser.ASTWindowDefinition_window_spec(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWindowSpecificationNode(v)
}

type WindowFrameUnit int

const (
	WindowFrameRows WindowFrameUnit = iota
	WindowFrameRange
)

func (u WindowFrameUnit) String() string {
	switch u {
	case WindowFrameRows:
		return "ROWS"
	case WindowFrameRange:
		return "RANGE"
	}
	return ""
}

type WindowFrameNode struct {
	*BaseNode
}

func (n *WindowFrameNode) StartExpr() *WindowFrameExprNode {
	var v unsafe.Pointer
	internalparser.ASTWindowFrame_start_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWindowFrameExprNode(v)
}

func (n *WindowFrameNode) EndExpr() *WindowFrameExprNode {
	var v unsafe.Pointer
	internalparser.ASTWindowFrame_end_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWindowFrameExprNode(v)
}

func (n *WindowFrameNode) SetUnit(frameUnit WindowFrameUnit) {
	internalparser.ASTWindowFrame_set_unit(n.getRaw(), int(frameUnit))
}

func (n *WindowFrameNode) FrameUnit() WindowFrameUnit {
	var v int
	internalparser.ASTWindowFrame_frame_unit(n.getRaw(), &v)
	return WindowFrameUnit(v)
}

func (n *WindowFrameNode) FrameUnitString() string {
	var v unsafe.Pointer
	internalparser.ASTWindowFrame_GetFrameUnitString(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type WindowFrameBoundaryType int

const (
	UnboundedPrecedingType WindowFrameBoundaryType = iota
	OffsetPrecedingType
	CurrentRowType
	OffsetFollowingType
	UnboundedFollowingType
)

func (t WindowFrameBoundaryType) String() string {
	switch t {
	case UnboundedPrecedingType:
		return "UNBOUNDED_PRECEDING"
	case OffsetPrecedingType:
		return "OFFSET_PRECEDING"
	case CurrentRowType:
		return "CURRENT_ROW"
	case OffsetFollowingType:
		return "OFFSET_FOLLOWING"
	case UnboundedFollowingType:
		return "UNBOUNDED_FOLLOWING"
	}
	return ""
}

type WindowFrameExprNode struct {
	*BaseNode
}

func (n *WindowFrameExprNode) SetBoundaryType(boundaryType WindowFrameBoundaryType) {
	internalparser.ASTWindowFrameExpr_set_boundary_type(n.getRaw(), int(boundaryType))
}

func (n *WindowFrameExprNode) BoundaryType() WindowFrameBoundaryType {
	var v int
	internalparser.ASTWindowFrameExpr_boundary_type(n.getRaw(), &v)
	return WindowFrameBoundaryType(v)
}

func (n *WindowFrameExprNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTWindowFrameExpr_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type LikeExpressionNode struct {
	*ExpressionBaseNode
}

func (n *LikeExpressionNode) SetIsNot(isNot bool) {
	internalparser.ASTLikeExpression_set_is_not(n.getRaw(), boolToInt(isNot))
}

func (n *LikeExpressionNode) IsNot() bool {
	var v bool
	internalparser.ASTLikeExpression_is_not(n.getRaw(), &v)
	return v
}

func (n *LikeExpressionNode) Lhs() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTLikeExpression_lhs(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *LikeExpressionNode) Op() *AnySomeAllOpNode {
	var v unsafe.Pointer
	internalparser.ASTLikeExpression_op(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAnySomeAllOpNode(v)
}

func (n *LikeExpressionNode) Hint() *HintNode {
	var v unsafe.Pointer
	internalparser.ASTLikeExpression_hint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHintNode(v)
}

func (n *LikeExpressionNode) InList() *InListNode {
	var v unsafe.Pointer
	internalparser.ASTLikeExpression_in_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newInListNode(v)
}

func (n *LikeExpressionNode) Query() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTLikeExpression_query(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

func (n *LikeExpressionNode) UnnestExpr() *UnnestExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTLikeExpression_unnest_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newUnnestExpressionNode(v)
}

type WindowSpecificationNode struct {
	*BaseNode
}

func (n *WindowSpecificationNode) BaseWindowName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTWindowSpecification_base_window_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *WindowSpecificationNode) PartitionBy() *PartitionByNode {
	var v unsafe.Pointer
	internalparser.ASTWindowSpecification_partition_by(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPartitionByNode(v)
}

func (n *WindowSpecificationNode) OrderBy() *OrderByNode {
	var v unsafe.Pointer
	internalparser.ASTWindowSpecification_order_by(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOrderByNode(v)
}

func (n *WindowSpecificationNode) WindowFrame() *WindowFrameNode {
	var v unsafe.Pointer
	internalparser.ASTWindowSpecification_window_frame(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWindowFrameNode(v)
}

type WithOffsetNode struct {
	*BaseNode
}

func (n *WithOffsetNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTWithOffset_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

type AnySomeAllOpType int

const (
	UninitializedAnySomeAllOp AnySomeAllOpType = iota
	AnyOp
	SomeOp
	AllOp
)

func (o AnySomeAllOpType) String() string {
	switch o {
	case UninitializedAnySomeAllOp:
		return "UNINITIALIZED"
	case AnyOp:
		return "ANY"
	case SomeOp:
		return "SOME"
	case AllOp:
		return "ALL"
	}
	return ""
}

type AnySomeAllOpNode struct {
	*BaseNode
}

func (n *AnySomeAllOpNode) SetOp(op AnySomeAllOpType) {
	internalparser.ASTAnySomeAllOp_set_op(n.getRaw(), int(op))
}

func (n *AnySomeAllOpNode) Op() AnySomeAllOpType {
	var v int
	internalparser.ASTAnySomeAllOp_op(n.getRaw(), &v)
	return AnySomeAllOpType(v)
}

func (n *AnySomeAllOpNode) SQLForOperator() string {
	var v unsafe.Pointer
	internalparser.ASTAnySomeAllOp_GetSQLForOperator(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type ParameterExprBaseNode struct {
	*ExpressionBaseNode
}

type StatementListNode struct {
	*BaseNode
}

func (n *StatementListNode) SetVariableDeclarationsAllowed(allowed bool) {
	internalparser.ASTStatementList_set_variable_declarations_allowed(n.getRaw(), boolToInt(allowed))
}

func (n *StatementListNode) VariableDeclarationsAllowed() bool {
	var v bool
	internalparser.ASTStatementList_variable_declarations_allowed(n.getRaw(), &v)
	return v
}

func (n *StatementListNode) StatementList() []StatementNode {
	var num int
	internalparser.ASTStatementList_statement_list_num(n.getRaw(), &num)
	ret := make([]StatementNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTStatementList_statement_list(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(StatementNode))
	}
	return ret
}

type ScriptStatementNode struct {
	*StatementBaseNode
}

type HintedStatementNode struct {
	*StatementBaseNode
}

func (n *HintedStatementNode) Hint() *HintNode {
	var v unsafe.Pointer
	internalparser.ASTHintedStatement_hint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHintNode(v)
}

func (n *HintedStatementNode) Statement() StatementNode {
	var v unsafe.Pointer
	internalparser.ASTHintedStatement_statement(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(StatementNode)
}

type ExplainStatementNode struct {
	*StatementBaseNode
}

func (n *ExplainStatementNode) Statement() StatementNode {
	var v unsafe.Pointer
	internalparser.ASTExplainStatement_statement(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(StatementNode)
}

type DescribeStatementNode struct {
	*StatementBaseNode
}

func (n *DescribeStatementNode) OptionalIdentifier() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTDescribeStatement_optional_identifier(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *DescribeStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDescribeStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *DescribeStatementNode) OptionalFromName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDescribeStatement_optional_from_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type ShowStatementNode struct {
	*StatementBaseNode
}

func (n *ShowStatementNode) Identifier() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTShowStatement_identifier(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *ShowStatementNode) OptionalName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTShowStatement_optional_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *ShowStatementNode) OptionalLikeString() *StringLiteralNode {
	var v unsafe.Pointer
	internalparser.ASTShowStatement_optional_like_string(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStringLiteralNode(v)
}

type TransactionModeBaseNode struct {
	*BaseNode
}

type TransactionIsolationLevelNode struct {
	*TransactionModeBaseNode
}

func (n *TransactionIsolationLevelNode) Identifier1() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTTransactionIsolationLevel_identifier1(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

// Identifier2 second identifier can be non-null only if first identifier is non-null.
func (n *TransactionIsolationLevelNode) Identifier2() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTTransactionIsolationLevel_identifier2(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type TransactionReadWriteModeType int

const (
	InvalidTransactionMode TransactionReadWriteModeType = iota
	ReadOnlyTransactionMode
	ReadWriteTransactionMode
)

func (m TransactionReadWriteModeType) String() string {
	switch m {
	case InvalidTransactionMode:
		return "INVALID"
	case ReadOnlyTransactionMode:
		return "READ_ONLY"
	case ReadWriteTransactionMode:
		return "READ_WRITE"
	}
	return ""
}

type TransactionReadWriteModeNode struct {
	*TransactionModeBaseNode
}

func (n *TransactionReadWriteModeNode) SetMode(modeType TransactionReadWriteModeType) {
	internalparser.ASTTransactionReadWriteMode_set_mode(n.getRaw(), int(modeType))
}

func (n *TransactionReadWriteModeNode) Mode() TransactionReadWriteModeType {
	var v int
	internalparser.ASTTransactionReadWriteMode_mode(n.getRaw(), &v)
	return TransactionReadWriteModeType(v)
}

type TransactionModeListNode struct {
	*BaseNode
}

func (n *TransactionModeListNode) Elements() []TransactionModeNode {
	var num int
	internalparser.ASTTransactionModeList_elements_num(n.getRaw(), &num)
	ret := make([]TransactionModeNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTTransactionModeList_element(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(TransactionModeNode))
	}
	return ret
}

type BeginStatementNode struct {
	*StatementBaseNode
}

func (n *BeginStatementNode) ModeList() *TransactionModeListNode {
	var v unsafe.Pointer
	internalparser.ASTBeginStatement_mode_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newTransactionModeListNode(v)
}

type SetTransactionStatementNode struct {
	*StatementBaseNode
}

func (n *SetTransactionStatementNode) ModeList() *TransactionModeListNode {
	var v unsafe.Pointer
	internalparser.ASTSetTransactionStatement_mode_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newTransactionModeListNode(v)
}

type CommitStatementNode struct {
	*StatementBaseNode
}

type RollbackStatementNode struct {
	*StatementBaseNode
}

type StartBatchStatementNode struct {
	*StatementBaseNode
}

func (n *StartBatchStatementNode) BatchType() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTStartBatchStatement_batch_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type RunBatchStatementNode struct {
	*StatementBaseNode
}

type AbortBatchStatementNode struct {
	*StatementBaseNode
}

type DdlStatementBaseNode struct {
	*StatementBaseNode
}

func (n *DdlStatementBaseNode) DdlTarget() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDdlStatement_GetDdlTarget(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type DropEntityStatementNode struct {
	*DdlStatementBaseNode
}

func (n *DropEntityStatementNode) SetIsIfExists(isIfExists bool) {
	internalparser.ASTDropEntityStatement_set_is_if_exists(n.getRaw(), boolToInt(isIfExists))
}

func (n *DropEntityStatementNode) IsIfExists() bool {
	var v bool
	internalparser.ASTDropEntityStatement_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *DropEntityStatementNode) EntityType() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTDropEntityStatement_entity_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *DropEntityStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDropEntityStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type DropFunctionStatementNode struct {
	*DdlStatementBaseNode
}

func (n *DropFunctionStatementNode) SetIsIfExists(isIfExists bool) {
	internalparser.ASTDropFunctionStatement_set_is_if_exists(n.getRaw(), boolToInt(isIfExists))
}

func (n *DropFunctionStatementNode) IsIfExists() bool {
	var v bool
	internalparser.ASTDropFunctionStatement_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *DropFunctionStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDropFunctionStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *DropFunctionStatementNode) Parameters() *FunctionParametersNode {
	var v unsafe.Pointer
	internalparser.ASTDropFunctionStatement_parameters(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newFunctionParametersNode(v)
}

type DropTableFunctionStatementNode struct {
	*DdlStatementBaseNode
}

func (n *DropTableFunctionStatementNode) SetIsIfExists(isIfExists bool) {
	internalparser.ASTDropTableFunctionStatement_set_is_if_exists(n.getRaw(), boolToInt(isIfExists))
}

func (n *DropTableFunctionStatementNode) IsIfExists() bool {
	var v bool
	internalparser.ASTDropTableFunctionStatement_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *DropTableFunctionStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDropTableFunctionStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type DropAllRowAccessPoliciesStatementNode struct {
	*StatementBaseNode
}

func (n *DropAllRowAccessPoliciesStatementNode) SetHasAccessKeyword(keyword bool) {
	internalparser.ASTDropAllRowAccessPoliciesStatement_set_has_access_keyword(n.getRaw(), boolToInt(keyword))
}

func (n *DropAllRowAccessPoliciesStatementNode) HasAccessKeyword() bool {
	var v bool
	internalparser.ASTDropAllRowAccessPoliciesStatement_has_access_keyword(n.getRaw(), &v)
	return v
}

func (n *DropAllRowAccessPoliciesStatementNode) TableName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDropAllRowAccessPoliciesStatement_table_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type DropMaterializedViewStatementNode struct {
	*DdlStatementBaseNode
}

func (n *DropMaterializedViewStatementNode) SetIsIfExists(isIfExists bool) {
	internalparser.ASTDropMaterializedViewStatement_set_is_if_exists(n.getRaw(), boolToInt(isIfExists))
}

func (n *DropMaterializedViewStatementNode) IsIfExists() bool {
	var v bool
	internalparser.ASTDropMaterializedViewStatement_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *DropMaterializedViewStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDropMaterializedViewStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type DropSnapshotTableStatementNode struct {
	*DdlStatementBaseNode
}

func (n *DropSnapshotTableStatementNode) SetIsIfExists(isIfExists bool) {
	internalparser.ASTDropSnapshotTableStatement_set_is_if_exists(n.getRaw(), boolToInt(isIfExists))
}

func (n *DropSnapshotTableStatementNode) IsIfExists() bool {
	var v bool
	internalparser.ASTDropSnapshotTableStatement_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *DropSnapshotTableStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDropSnapshotTableStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type DropSearchIndexStatementNode struct {
	*DdlStatementBaseNode
}

func (n *DropSearchIndexStatementNode) SetIsIfExists(isIfExists bool) {
	internalparser.ASTDropSearchIndexStatement_set_is_if_exists(n.getRaw(), boolToInt(isIfExists))
}

func (n *DropSearchIndexStatementNode) IsIfExists() bool {
	var v bool
	internalparser.ASTDropSearchIndexStatement_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *DropSearchIndexStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDropSearchIndexStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *DropSearchIndexStatementNode) TableName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDropSearchIndexStatement_table_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type RenameStatementNode struct {
	*StatementBaseNode
}

func (n *RenameStatementNode) Identifier() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTRenameStatement_identifier(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *RenameStatementNode) OldName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTRenameStatement_old_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *RenameStatementNode) NewName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTRenameStatement_new_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type ImportStatementKind int

const (
	ImportModule ImportStatementKind = iota
	ImportProto
)

type ImportStatementNode struct {
	*StatementBaseNode
}

func (n *ImportStatementNode) SetImportKind(kind ImportStatementKind) {
	internalparser.ASTImportStatement_set_import_kind(n.getRaw(), int(kind))
}

func (n *ImportStatementNode) ImportKind() ImportStatementKind {
	var v int
	internalparser.ASTImportStatement_import_kind(n.getRaw(), &v)
	return ImportStatementKind(v)
}

func (n *ImportStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTImportStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *ImportStatementNode) StringValue() *StringLiteralNode {
	var v unsafe.Pointer
	internalparser.ASTImportStatement_string_value(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStringLiteralNode(v)
}

func (n *ImportStatementNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTImportStatement_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

func (n *ImportStatementNode) IntoAlias() *IntoAliasNode {
	var v unsafe.Pointer
	internalparser.ASTImportStatement_into_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIntoAliasNode(v)
}

func (n *ImportStatementNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTImportStatement_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type ModuleStatementNode struct {
	*StatementBaseNode
}

func (n *ModuleStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTModuleStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *ModuleStatementNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTModuleStatement_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type WithConnectionClauseNode struct {
	*BaseNode
}

func (n *WithConnectionClauseNode) ConnectionClause() *ConnectionClauseNode {
	var v unsafe.Pointer
	internalparser.ASTWithConnectionClause_connection_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newConnectionClauseNode(v)
}

type IntoAliasNode struct {
	*BaseNode
}

func (n *IntoAliasNode) Identifier() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTIntoAlias_identifier(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *IntoAliasNode) Name() string {
	var v unsafe.Pointer
	internalparser.ASTIntoAlias_GetAsString(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type UnnestExpressionWithOptAliasAndOffsetNode struct {
	*BaseNode
}

func (n *UnnestExpressionWithOptAliasAndOffsetNode) UnnestExpression() *UnnestExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTUnnestExpressionWithOptAliasAndOffset_unnest_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newUnnestExpressionNode(v)
}

func (n *UnnestExpressionWithOptAliasAndOffsetNode) OptionalAlias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTUnnestExpressionWithOptAliasAndOffset_optional_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

func (n *UnnestExpressionWithOptAliasAndOffsetNode) OptionalWithOffset() *WithOffsetNode {
	var v unsafe.Pointer
	internalparser.ASTUnnestExpressionWithOptAliasAndOffset_optional_with_offset(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWithOffsetNode(v)
}

type PivotExpressionNode struct {
	*BaseNode
}

func (n *PivotExpressionNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTPivotExpression_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *PivotExpressionNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTPivotExpression_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

type PivotValueNode struct {
	*BaseNode
}

func (n *PivotValueNode) Value() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTPivotValue_value(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *PivotValueNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTPivotValue_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

type PivotExpressionListNode struct {
	*BaseNode
}

func (n *PivotExpressionListNode) Expressions() []*PivotExpressionNode {
	var num int
	internalparser.ASTPivotExpressionList_expressions_num(n.getRaw(), &num)
	ret := make([]*PivotExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTPivotExpressionList_expression(n.getRaw(), i, &v)
		ret = append(ret, newPivotExpressionNode(v))
	}
	return ret
}

type PivotValueListNode struct {
	*BaseNode
}

func (n *PivotValueListNode) Values() []*PivotValueNode {
	var num int
	internalparser.ASTPivotValueList_values_num(n.getRaw(), &num)
	ret := make([]*PivotValueNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTPivotValueList_value(n.getRaw(), i, &v)
		ret = append(ret, newPivotValueNode(v))
	}
	return ret
}

type PivotClauseNode struct {
	*BaseNode
}

func (n *PivotClauseNode) PivotExpressions() *PivotExpressionListNode {
	var v unsafe.Pointer
	internalparser.ASTPivotClause_pivot_expressions(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPivotExpressionListNode(v)
}

func (n *PivotClauseNode) ForExpression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTPivotClause_for_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *PivotClauseNode) PivotValues() *PivotValueListNode {
	var v unsafe.Pointer
	internalparser.ASTPivotClause_pivot_values(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPivotValueListNode(v)
}

func (n *PivotClauseNode) OutputAlias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTPivotClause_output_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

type UnpivotInItemNode struct {
	*BaseNode
}

func (n *UnpivotInItemNode) UnpivotColumns() *PathExpressionListNode {
	var v unsafe.Pointer
	internalparser.ASTUnpivotInItem_unpivot_columns(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionListNode(v)
}

func (n *UnpivotInItemNode) Alias() *UnpivotInItemLabelNode {
	var v unsafe.Pointer
	internalparser.ASTUnpivotInItem_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newUnpivotInItemLabelNode(v)
}

type UnpivotInItemListNode struct {
	*BaseNode
}

func (n *UnpivotInItemListNode) InItems() []*UnpivotInItemNode {
	var num int
	internalparser.ASTUnpivotInItemList_in_items_num(n.getRaw(), &num)
	ret := make([]*UnpivotInItemNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTUnpivotInItemList_in_item(n.getRaw(), i, &v)
		ret = append(ret, newUnpivotInItemNode(v))
	}
	return ret
}

type UnpivotNullFilter int

const (
	UnpivotUnspecified UnpivotNullFilter = iota
	UnpivotInclude
	UnpivotExclude
)

func (f UnpivotNullFilter) String() string {
	switch f {
	case UnpivotUnspecified:
		return "UNSPECIFIED"
	case UnpivotInclude:
		return "INCLUDE"
	case UnpivotExclude:
		return "EXCLUDE"
	}
	return ""
}

type UnpivotClauseNode struct {
	*BaseNode
}

func (n *UnpivotClauseNode) SetNullFilter(filter UnpivotNullFilter) {
	internalparser.ASTUnpivotClause_set_null_filter(n.getRaw(), int(filter))
}

func (n *UnpivotClauseNode) NullFilter() UnpivotNullFilter {
	var v int
	internalparser.ASTUnpivotClause_null_filter(n.getRaw(), &v)
	return UnpivotNullFilter(v)
}

func (n *UnpivotClauseNode) UnpivotOutputValueColumns() *PathExpressionListNode {
	var v unsafe.Pointer
	internalparser.ASTUnpivotClause_unpivot_output_value_columns(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionListNode(v)
}

func (n *UnpivotClauseNode) UnpivotOutputNameColumn() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTUnpivotClause_unpivot_output_name_column(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *UnpivotClauseNode) UnpivotInItems() *UnpivotInItemListNode {
	var v unsafe.Pointer
	internalparser.ASTUnpivotClause_unpivot_in_items(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newUnpivotInItemListNode(v)
}

func (n *UnpivotClauseNode) OutputAlias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTUnpivotClause_output_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

type UsingClauseNode struct {
	*BaseNode
}

func (n *UsingClauseNode) Keys() []*IdentifierNode {
	var num int
	internalparser.ASTUsingClause_keys_num(n.getRaw(), &num)
	ret := make([]*IdentifierNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTUsingClause_key(n.getRaw(), i, &v)
		ret = append(ret, newIdentifierNode(v))
	}
	return ret
}

type ForSystemTimeNode struct {
	*BaseNode
}

func (n *ForSystemTimeNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTForSystemTime_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type QualifyNode struct {
	*BaseNode
}

func (n *QualifyNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTQualify_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type ClampedBetweenModifierNode struct {
	*BaseNode
}

func (n *ClampedBetweenModifierNode) Low() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTClampedBetweenModifier_low(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *ClampedBetweenModifierNode) High() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTClampedBetweenModifier_high(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type FormatClauseNode struct {
	*BaseNode
}

func (n *FormatClauseNode) Format() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTFormatClause_format(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *FormatClauseNode) TimeZoneExpr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTFormatClause_time_zone_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type PathExpressionListNode struct {
	*BaseNode
}

func (n *PathExpressionListNode) PathExpressionList() []*PathExpressionNode {
	var num int
	internalparser.ASTPathExpressionList_path_expression_list_num(n.getRaw(), &num)
	ret := make([]*PathExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTPathExpressionList_path_expression_list(n.getRaw(), i, &v)
		ret = append(ret, newPathExpressionNode(v))
	}
	return ret
}

type ParameterExprNode struct {
	*ParameterExprBaseNode
}

func (n *ParameterExprNode) SetPosition(pos int) {
	internalparser.ASTParameterExpr_set_position(n.getRaw(), pos)
}

func (n *ParameterExprNode) Position() int {
	var v int
	internalparser.ASTParameterExpr_position(n.getRaw(), &v)
	return v
}

func (n *ParameterExprNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTParameterExpr_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type SystemVariableExprNode struct {
	*ParameterExprBaseNode
}

func (n *SystemVariableExprNode) Path() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTSystemVariableExpr_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type WithGroupRowsNode struct {
	*BaseNode
}

func (n *WithGroupRowsNode) Subquery() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTWithGroupRows_subquery(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

type LambdaNode struct {
	*ExpressionBaseNode
}

func (n *LambdaNode) ArgumentList() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTLambda_argument_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *LambdaNode) Body() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTLambda_body(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type AnalyticFunctionCallNode struct {
	*ExpressionBaseNode
}

func (n *AnalyticFunctionCallNode) WindowSpec() *WindowSpecificationNode {
	var v unsafe.Pointer
	internalparser.ASTAnalyticFunctionCall_window_spec(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWindowSpecificationNode(v)
}

func (n *AnalyticFunctionCallNode) Function() *FunctionCallNode {
	var v unsafe.Pointer
	internalparser.ASTAnalyticFunctionCall_function(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newFunctionCallNode(v)
}

func (n *AnalyticFunctionCallNode) FunctionWithGroupRows() *FunctionCallWithGroupRowsNode {
	var v unsafe.Pointer
	internalparser.ASTAnalyticFunctionCall_function_with_group_rows(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newFunctionCallWithGroupRowsNode(v)
}

type FunctionCallWithGroupRowsNode struct {
	*ExpressionBaseNode
}

func (n *FunctionCallWithGroupRowsNode) Function() *FunctionCallNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionCallWithGroupRows_function(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newFunctionCallNode(v)
}

func (n *FunctionCallWithGroupRowsNode) Subquery() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionCallWithGroupRows_subquery(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

type ClusterByNode struct {
	*BaseNode
}

func (n *ClusterByNode) ClusteringExpressions() []ExpressionNode {
	var num int
	internalparser.ASTClusterBy_clustering_expressions_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTClusterBy_clustering_expression(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

type NewConstructorArgNode struct {
	*BaseNode
}

func (n *NewConstructorArgNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTNewConstructorArg_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *NewConstructorArgNode) OptionalIdentifier() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTNewConstructorArg_optional_identifier(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *NewConstructorArgNode) OptionalPathExpression() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTNewConstructorArg_optional_path_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type NewConstructorNode struct {
	*ExpressionBaseNode
}

func (n *NewConstructorNode) TypeName() *SimpleTypeNode {
	var v unsafe.Pointer
	internalparser.ASTNewConstructor_type_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newSimpleTypeNode(v)
}

func (n *NewConstructorNode) Arguments() []*NewConstructorArgNode {
	var num int
	internalparser.ASTNewConstructor_arguments_num(n.getRaw(), &num)
	ret := make([]*NewConstructorArgNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTNewConstructor_argument(n.getRaw(), i, &v)
		ret = append(ret, newNewConstructorArgNode(v))
	}
	return ret
}

type OptionsListNode struct {
	*BaseNode
}

func (n *OptionsListNode) OptionsEntries() []*OptionsEntryNode {
	var num int
	internalparser.ASTOptionsList_options_entries_num(n.getRaw(), &num)
	ret := make([]*OptionsEntryNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTOptionsList_options_entry(n.getRaw(), i, &v)
		ret = append(ret, newOptionsEntryNode(v))
	}
	return ret
}

type OptionsEntryNode struct {
	*BaseNode
}

func (n *OptionsEntryNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTOptionsEntry_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *OptionsEntryNode) Value() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTOptionsEntry_value(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type CreateStatementScope int

const (
	CreateStatementDefaultScope CreateStatementScope = iota
	CreateStatementPrivate
	CreateStatementPublic
	CreateStatementTemporary
)

func (s CreateStatementScope) String() string {
	switch s {
	case CreateStatementDefaultScope:
		return "DEFAULT_SCOPE"
	case CreateStatementPrivate:
		return "PRIVATE"
	case CreateStatementPublic:
		return "PUBLIC"
	case CreateStatementTemporary:
		return "TEMPORARY"
	}
	return ""
}

type SqlSecurity int

const (
	SqlSecurityUnspecified SqlSecurity = iota
	SqlSecurityDefiner
	SqlSecurityInvoker
)

func (s SqlSecurity) String() string {
	switch s {
	case SqlSecurityUnspecified:
		return "SQL_SECURITY_UNSPECIFIED"
	case SqlSecurityDefiner:
		return "SQL_SECURITY_DEFINER"
	case SqlSecurityInvoker:
		return "SQL_SECURITY_INVOKER"
	}
	return ""
}

type CreateStatementNode struct {
	*DdlStatementBaseNode
}

func (n *CreateStatementNode) SetScope(scope CreateStatementScope) {
	internalparser.ASTCreateStatement_set_scope(n.getRaw(), int(scope))
}

func (n *CreateStatementNode) Scope() CreateStatementScope {
	var v int
	internalparser.ASTCreateStatement_scope(n.getRaw(), &v)
	return CreateStatementScope(v)
}

func (n *CreateStatementNode) SetIsOrReplace(isOrReplace bool) {
	internalparser.ASTCreateStatement_set_is_or_replace(n.getRaw(), boolToInt(isOrReplace))
}

func (n *CreateStatementNode) IsOrReplace() bool {
	var v bool
	internalparser.ASTCreateStatement_is_or_replace(n.getRaw(), &v)
	return v
}

func (n *CreateStatementNode) SetIsIfNotExists(isIfNotExists bool) {
	internalparser.ASTCreateStatement_set_is_if_not_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *CreateStatementNode) IsIfNotExists() bool {
	var v bool
	internalparser.ASTCreateStatement_is_if_not_exists(n.getRaw(), &v)
	return v
}

func (n *CreateStatementNode) IsDefaultScope() bool {
	var v bool
	internalparser.ASTCreateStatement_is_default_scope(n.getRaw(), &v)
	return v
}

func (n *CreateStatementNode) IsPrivate() bool {
	var v bool
	internalparser.ASTCreateStatement_is_private(n.getRaw(), &v)
	return v
}

func (n *CreateStatementNode) IsPublic() bool {
	var v bool
	internalparser.ASTCreateStatement_is_public(n.getRaw(), &v)
	return v
}

func (n *CreateStatementNode) IsTemp() bool {
	var v bool
	internalparser.ASTCreateStatement_is_temp(n.getRaw(), &v)
	return v
}

type ProcedureParameterMode int

const (
	NotSetProcedureParameter ProcedureParameterMode = iota
	InProcedureParameter
	OutProcedureParameter
	InOutProcedureParameter
)

func (m ProcedureParameterMode) String() string {
	switch m {
	case NotSetProcedureParameter:
		return "NOT_SET"
	case InProcedureParameter:
		return "IN"
	case OutProcedureParameter:
		return "OUT"
	case InOutProcedureParameter:
		return "INOUT"
	}
	return ""
}

type FunctionParameterNode struct {
	*BaseNode
}

func (n *FunctionParameterNode) SetProcedureParameterMode(mode ProcedureParameterMode) {
	internalparser.ASTFunctionParameter_set_procedure_parameter_mode(n.getRaw(), int(mode))
}

func (n *FunctionParameterNode) ProcedureParameterMode() ProcedureParameterMode {
	var v int
	internalparser.ASTFunctionParameter_procedure_parameter_mode(n.getRaw(), &v)
	return ProcedureParameterMode(v)
}

func (n *FunctionParameterNode) SetIsNotAggregate(isNotAggregate bool) {
	internalparser.ASTFunctionParameter_set_is_not_aggregate(n.getRaw(), boolToInt(isNotAggregate))
}

func (n *FunctionParameterNode) IsNotAggregate() bool {
	var v bool
	internalparser.ASTFunctionParameter_is_not_aggregate(n.getRaw(), &v)
	return v
}

func (n *FunctionParameterNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionParameter_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *FunctionParameterNode) Type() TypeNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionParameter_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(TypeNode)
}

func (n *FunctionParameterNode) TemplatedParameterType() *TemplatedParameterTypeNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionParameter_templated_parameter_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newTemplatedParameterTypeNode(v)
}

func (n *FunctionParameterNode) TVFSchema() *TVFSchemaNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionParameter_tvf_schema(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newTVFSchemaNode(v)
}

func (n *FunctionParameterNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionParameter_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

func (n *FunctionParameterNode) DefaultValue() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionParameter_default_value(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *FunctionParameterNode) IsTableParameter() bool {
	var v bool
	internalparser.ASTFunctionParameter_IsTableParameter(n.getRaw(), &v)
	return v
}

func (n *FunctionParameterNode) IsTemplated() bool {
	var v bool
	internalparser.ASTFunctionParameter_IsTemplated(n.getRaw(), &v)
	return v
}

type FunctionParametersNode struct {
	*BaseNode
}

func (n *FunctionParametersNode) ParameterEntries() []*FunctionParameterNode {
	var num int
	internalparser.ASTFunctionParameters_parameter_entries_num(n.getRaw(), &num)
	ret := make([]*FunctionParameterNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTFunctionParameters_parameter_entry(n.getRaw(), i, &v)
		ret = append(ret, newFunctionParameterNode(v))
	}
	return ret
}

type FunctionDeclarationNode struct {
	*BaseNode
}

func (n *FunctionDeclarationNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionDeclaration_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *FunctionDeclarationNode) Parameters() *FunctionParametersNode {
	var v unsafe.Pointer
	internalparser.ASTFunctionDeclaration_parameters(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newFunctionParametersNode(v)
}

func (n *FunctionDeclarationNode) IsTemplated() bool {
	var v bool
	internalparser.ASTFunctionDeclaration_IsTemplated(n.getRaw(), &v)
	return v
}

type SqlFunctionBodyNode struct {
	*BaseNode
}

func (n *SqlFunctionBodyNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTSqlFunctionBody_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type TVFArgumentNode struct {
	*BaseNode
}

func (n *TVFArgumentNode) Expr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTTVFArgument_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *TVFArgumentNode) TableClause() *TableClauseNode {
	var v unsafe.Pointer
	internalparser.ASTTVFArgument_table_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newTableClauseNode(v)
}

func (n *TVFArgumentNode) ModelClause() *ModelClauseNode {
	var v unsafe.Pointer
	internalparser.ASTTVFArgument_model_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newModelClauseNode(v)
}

func (n *TVFArgumentNode) ConnectionClause() *ConnectionClauseNode {
	var v unsafe.Pointer
	internalparser.ASTTVFArgument_connection_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newConnectionClauseNode(v)
}

func (n *TVFArgumentNode) Descriptor() *DescriptorNode {
	var v unsafe.Pointer
	internalparser.ASTTVFArgument_descriptor(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newDescriptorNode(v)
}

type TVFNode struct {
	*TableExpressionBaseNode
}

func (n *TVFNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTTVF_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *TVFNode) Hint() *HintNode {
	var v unsafe.Pointer
	internalparser.ASTTVF_hint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newHintNode(v)
}

func (n *TVFNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTTVF_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

func (n *TVFNode) PivotClause() *PivotClauseNode {
	var v unsafe.Pointer
	internalparser.ASTTVF_pivot_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPivotClauseNode(v)
}

func (n *TVFNode) UnpivotClause() *UnpivotClauseNode {
	var v unsafe.Pointer
	internalparser.ASTTVF_unpivot_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newUnpivotClauseNode(v)
}

func (n *TVFNode) SampleClause() *SampleClauseNode {
	var v unsafe.Pointer
	internalparser.ASTTVF_sample(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newSampleClauseNode(v)
}

func (n *TVFNode) ArgumentEntries() []*TVFArgumentNode {
	var num int
	internalparser.ASTTVF_argument_entries_num(n.getRaw(), &num)
	ret := make([]*TVFArgumentNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTTVF_argument_entry(n.getRaw(), i, &v)
		ret = append(ret, newTVFArgumentNode(v))
	}
	return ret
}

type TableClauseNode struct {
	*BaseNode
}

func (n *TableClauseNode) TablePath() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTTableClause_table_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *TableClauseNode) TVF() *TVFNode {
	var v unsafe.Pointer
	internalparser.ASTTableClause_tvf(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newTVFNode(v)
}

type ModelClauseNode struct {
	*BaseNode
}

func (n *ModelClauseNode) ModelPath() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTModelClause_model_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type ConnectionClauseNode struct {
	*BaseNode
}

func (n *ConnectionClauseNode) ConnectionPath() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTConnectionClause_connection_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type TableDataSourceNode struct {
	*TableExpressionBaseNode
}

func (n *TableDataSourceNode) PathExpr() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTTableDataSource_path_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *TableDataSourceNode) ForSystemTime() *ForSystemTimeNode {
	var v unsafe.Pointer
	internalparser.ASTTableDataSource_for_system_time(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newForSystemTimeNode(v)
}

func (n *TableDataSourceNode) WhereClause() *WhereClauseNode {
	var v unsafe.Pointer
	internalparser.ASTTableDataSource_where_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWhereClauseNode(v)
}

type CloneDataSourceNode struct {
	*TableDataSourceNode
}

type CopyDataSourceNode struct {
	*TableDataSourceNode
}

type CloneDataSourceListNode struct {
	*BaseNode
}

func (n *CloneDataSourceListNode) DataSources() []*CloneDataSourceNode {
	var num int
	internalparser.ASTCloneDataSourceList_data_sources_num(n.getRaw(), &num)
	ret := make([]*CloneDataSourceNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTCloneDataSourceList_data_source(n.getRaw(), i, &v)
		ret = append(ret, newCloneDataSourceNode(v))
	}
	return ret
}

type CloneDataStatementNode struct {
	*StatementBaseNode
}

func (n *CloneDataStatementNode) TargetPath() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCloneDataStatement_target_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CloneDataStatementNode) DataSourceList() *CloneDataSourceListNode {
	var v unsafe.Pointer
	internalparser.ASTCloneDataStatement_data_source_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newCloneDataSourceListNode(v)
}

type CreateConstantStatementNode struct {
	*CreateStatementNode
}

func (n *CreateConstantStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateConstantStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreateConstantStatementNode) Expr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateConstantStatement_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type CreateDatabaseStatementNode struct {
	*StatementBaseNode
}

func (n *CreateDatabaseStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateDatabaseStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreateDatabaseStatementNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateDatabaseStatement_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type CreateProcedureStatementNode struct {
	*CreateStatementNode
}

func (n *CreateProcedureStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateProcedureStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreateProcedureStatementNode) Parameters() *FunctionParametersNode {
	var v unsafe.Pointer
	internalparser.ASTCreateProcedureStatement_parameters(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newFunctionParametersNode(v)
}

func (n *CreateProcedureStatementNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateProcedureStatement_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

func (n *CreateProcedureStatementNode) Body() ScriptNode {
	var v unsafe.Pointer
	internalparser.ASTCreateProcedureStatement_body(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ScriptNode)
}

type CreateSchemaStatementNode struct {
	*CreateStatementNode
}

func (n *CreateSchemaStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateSchemaStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreateSchemaStatementNode) Collate() *CollateNode {
	var v unsafe.Pointer
	internalparser.ASTCreateSchemaStatement_collate(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newCollateNode(v)
}

func (n *CreateSchemaStatementNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateSchemaStatement_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type TransformClauseNode struct {
	*BaseNode
}

func (n *TransformClauseNode) SelectList() *SelectListNode {
	var v unsafe.Pointer
	internalparser.ASTTransformClause_select_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newSelectListNode(v)
}

type CreateModelStatementNode struct {
	*CreateStatementNode
}

func (n *CreateModelStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateModelStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreateModelStatementNode) TransformClause() *TransformClauseNode {
	var v unsafe.Pointer
	internalparser.ASTCreateModelStatement_transform_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newTransformClauseNode(v)
}

func (n *CreateModelStatementNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateModelStatement_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

func (n *CreateModelStatementNode) Query() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTCreateModelStatement_query(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

type IndexAllColumnsNode struct {
	*LeafBaseNode
}

type IndexItemListNode struct {
	*BaseNode
}

func (n *IndexItemListNode) OrderingExpressions() []*OrderingExpressionNode {
	var num int
	internalparser.ASTIndexItemList_ordering_expressions_num(n.getRaw(), &num)
	ret := make([]*OrderingExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTIndexItemList_ordering_expression(n.getRaw(), i, &v)
		ret = append(ret, newOrderingExpressionNode(v))
	}
	return ret
}

type IndexStoringExpressionListNode struct {
	*BaseNode
}

func (n *IndexStoringExpressionListNode) Expressions() []ExpressionNode {
	var num int
	internalparser.ASTIndexStoringExpressionList_expressions_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTIndexStoringExpressionList_expression(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

type IndexUnnestExpressionListNode struct {
	*BaseNode
}

func (n *IndexUnnestExpressionListNode) UnnestExpressions() []*UnnestExpressionWithOptAliasAndOffsetNode {
	var num int
	internalparser.ASTIndexUnnestExpressionList_unnest_expressions_num(n.getRaw(), &num)
	ret := make([]*UnnestExpressionWithOptAliasAndOffsetNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTIndexUnnestExpressionList_unnest_expression(n.getRaw(), i, &v)
		ret = append(ret, newUnnestExpressionWithOptAliasAndOffsetNode(v))
	}
	return ret
}

type CreateIndexStatementNode struct {
	*CreateStatementNode
}

func (n *CreateIndexStatementNode) SetIsUnique(isUnique bool) {
	internalparser.ASTCreateIndexStatement_set_is_unique(n.getRaw(), boolToInt(isUnique))
}

func (n *CreateIndexStatementNode) IsUnique() bool {
	var v bool
	internalparser.ASTCreateIndexStatement_is_unique(n.getRaw(), &v)
	return v
}

func (n *CreateIndexStatementNode) SetIsSearch(isSearch bool) {
	internalparser.ASTCreateIndexStatement_set_is_search(n.getRaw(), boolToInt(isSearch))
}

func (n *CreateIndexStatementNode) IsSearch() bool {
	var v bool
	internalparser.ASTCreateIndexStatement_is_search(n.getRaw(), &v)
	return v
}

func (n *CreateIndexStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateIndexStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreateIndexStatementNode) TableName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateIndexStatement_table_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreateIndexStatementNode) OptionalTableAlias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTCreateIndexStatement_optional_table_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

func (n *CreateIndexStatementNode) OptionalIndexUnnestExpressionList() *IndexUnnestExpressionListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateIndexStatement_optional_index_unnest_expression_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIndexUnnestExpressionListNode(v)
}

func (n *CreateIndexStatementNode) IndexItemList() *IndexItemListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateIndexStatement_index_item_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIndexItemListNode(v)
}

func (n *CreateIndexStatementNode) OptionalIndexSotringExpressions() *IndexStoringExpressionListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateIndexStatement_optional_index_storing_expressions(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIndexStoringExpressionListNode(v)
}

func (n *CreateIndexStatementNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateIndexStatement_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type ExportDataStatementNode struct {
	*StatementBaseNode
}

func (n *ExportDataStatementNode) WithConnectionClause() *WithConnectionClauseNode {
	var v unsafe.Pointer
	internalparser.ASTExportDataStatement_with_connection_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWithConnectionClauseNode(v)
}

func (n *ExportDataStatementNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTExportDataStatement_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

func (n *ExportDataStatementNode) Query() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTExportDataStatement_query(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

type ExportModelStatementNode struct {
	*StatementBaseNode
}

func (n *ExportModelStatementNode) ModelNamePath() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTExportModelStatement_model_name_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *ExportModelStatementNode) WithConnectionClause() *WithConnectionClauseNode {
	var v unsafe.Pointer
	internalparser.ASTExportModelStatement_with_connection_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWithConnectionClauseNode(v)
}

func (n *ExportModelStatementNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTExportModelStatement_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type CallStatementNode struct {
	*StatementBaseNode
}

func (n *CallStatementNode) ProcedureName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCallStatement_procedure_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CallStatementNode) Arguments() []*TVFArgumentNode {
	var num int
	internalparser.ASTCallStatement_arguments_num(n.getRaw(), &num)
	ret := make([]*TVFArgumentNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTCallStatement_argument(n.getRaw(), i, &v)
		ret = append(ret, newTVFArgumentNode(v))
	}
	return ret
}

type DefineTableStatementNode struct {
	*StatementBaseNode
}

func (n *DefineTableStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDefineTableStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *DefineTableStatementNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTDefineTableStatement_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type WithPartitionColumnsClauseNode struct {
	*BaseNode
}

func (n *WithPartitionColumnsClauseNode) TableElementList() *TableElementListNode {
	var v unsafe.Pointer
	internalparser.ASTWithPartitionColumnsClause_table_element_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newTableElementListNode(v)
}

type CreateSnapshotTableStatementNode struct {
	*CreateStatementNode
}

func (n *CreateSnapshotTableStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateSnapshotTableStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreateSnapshotTableStatementNode) CloneDataSource() *CloneDataSourceNode {
	var v unsafe.Pointer
	internalparser.ASTCreateSnapshotTableStatement_clone_data_source(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newCloneDataSourceNode(v)
}

func (n *CreateSnapshotTableStatementNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateSnapshotTableStatement_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type TypeParameterListNode struct {
	*BaseNode
}

func (n *TypeParameterListNode) Parameters() []LeafNode {
	var num int
	internalparser.ASTTypeParameterList_parameters_num(n.getRaw(), &num)
	ret := make([]LeafNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTTypeParameterList_parameter(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(LeafNode))
	}
	return ret
}

type TVFSchemaNode struct {
	*BaseNode
}

func (n *TVFSchemaNode) Columns() []*TVFSchemaColumnNode {
	var num int
	internalparser.ASTTVFSchema_columns_num(n.getRaw(), &num)
	ret := make([]*TVFSchemaColumnNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTTVFSchema_column(n.getRaw(), i, &v)
		ret = append(ret, newTVFSchemaColumnNode(v))
	}
	return ret
}

type TVFSchemaColumnNode struct {
	*BaseNode
}

func (n *TVFSchemaColumnNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTTVFSchemaColumn_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *TVFSchemaColumnNode) Type() TypeNode {
	var v unsafe.Pointer
	internalparser.ASTTVFSchemaColumn_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(TypeNode)
}

type TableAndColumnInfoNode struct {
	*BaseNode
}

func (n *TableAndColumnInfoNode) TableName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTTableAndColumnInfo_table_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *TableAndColumnInfoNode) ColumnList() *ColumnListNode {
	var v unsafe.Pointer
	internalparser.ASTTableAndColumnInfo_column_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnListNode(v)
}

type TableAndColumnInfoListNode struct {
	*BaseNode
}

func (n *TableAndColumnInfoListNode) TableAndColumnInfoEntries() []*TableAndColumnInfoNode {
	var num int
	internalparser.ASTTableAndColumnInfoList_table_and_column_info_entries_num(n.getRaw(), &num)
	ret := make([]*TableAndColumnInfoNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTTableAndColumnInfoList_table_and_column_info_entry(n.getRaw(), i, &v)
		ret = append(ret, newTableAndColumnInfoNode(v))
	}
	return ret
}

type TemplatedTypeKind int

const (
	TemplatedUninitialized TemplatedTypeKind = iota
	TemplatedAnyType
	TemplatedAnyProto
	TemplatedAnyEnum
	TemplatedAnyStruct
	TemplatedAnyArray
	TemplatedAnyTable
)

func (k TemplatedTypeKind) String() string {
	switch k {
	case TemplatedUninitialized:
		return "UNINITIALIZED"
	case TemplatedAnyType:
		return "ANY_TYPE"
	case TemplatedAnyProto:
		return "ANY_PROTO"
	case TemplatedAnyEnum:
		return "ANY_ENUM"
	case TemplatedAnyStruct:
		return "ANY_STRUCT"
	case TemplatedAnyArray:
		return "ANY_ARRAY"
	case TemplatedAnyTable:
		return "ANY_TABLE"
	}
	return ""
}

type TemplatedParameterTypeNode struct {
	*BaseNode
}

func (n *TemplatedParameterTypeNode) SetKind(kind TemplatedTypeKind) {
	internalparser.ASTTemplatedParameterType_set_kind(n.getRaw(), int(kind))
}

func (n *TemplatedParameterTypeNode) TemplatedKind() TemplatedTypeKind {
	var v int
	internalparser.ASTTemplatedParameterType_kind(n.getRaw(), &v)
	return TemplatedTypeKind(v)
}

type DefaultLiteralNode struct {
	*ExpressionBaseNode
}

type AnalyzeStatementNode struct {
	*StatementBaseNode
}

func (n *AnalyzeStatementNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTAnalyzeStatement_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

func (n *AnalyzeStatementNode) TableAndColumnInfoList() *TableAndColumnInfoListNode {
	var v unsafe.Pointer
	internalparser.ASTAnalyzeStatement_table_and_column_info_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newTableAndColumnInfoListNode(v)
}

type AssertStatementNode struct {
	*StatementBaseNode
}

func (n *AssertStatementNode) Expr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTAssertStatement_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *AssertStatementNode) Description() *StringLiteralNode {
	var v unsafe.Pointer
	internalparser.ASTAssertStatement_description(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStringLiteralNode(v)
}

type AssertRowsModifiedNode struct {
	*BaseNode
}

func (n *AssertRowsModifiedNode) NumRows() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTAssertRowsModified_num_rows(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type ReturningClauseNode struct {
	*BaseNode
}

func (n *ReturningClauseNode) SelectList() *SelectListNode {
	var v unsafe.Pointer
	internalparser.ASTReturningClause_select_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newSelectListNode(v)
}

func (n *ReturningClauseNode) ActionAlias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTReturningClause_action_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

type DeleteStatementNode struct {
	*StatementBaseNode
}

func (n *DeleteStatementNode) TargetPath() GeneralizedPathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDeleteStatement_target_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(GeneralizedPathExpressionNode)
}

func (n *DeleteStatementNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTDeleteStatement_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

func (n *DeleteStatementNode) Offset() *WithOffsetNode {
	var v unsafe.Pointer
	internalparser.ASTDeleteStatement_offset(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWithOffsetNode(v)
}

func (n *DeleteStatementNode) Where() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDeleteStatement_where(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *DeleteStatementNode) AssertRowsModified() *AssertRowsModifiedNode {
	var v unsafe.Pointer
	internalparser.ASTDeleteStatement_assert_rows_modified(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAssertRowsModifiedNode(v)
}

func (n *DeleteStatementNode) Returning() *ReturningClauseNode {
	var v unsafe.Pointer
	internalparser.ASTDeleteStatement_returning(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newReturningClauseNode(v)
}

type ColumnAttributeBaseNode struct {
	*BaseNode
}

type NotNullColumnAttributeNode struct {
	*ColumnAttributeBaseNode
}

type HiddenColumnAttributeNode struct {
	*ColumnAttributeBaseNode
}

type PrimaryKeyColumnAttributeNode struct {
	*ColumnAttributeBaseNode
}

func (n *PrimaryKeyColumnAttributeNode) SetEnforced(enforced bool) {
	internalparser.ASTPrimaryKeyColumnAttribute_set_enforced(n.getRaw(), boolToInt(enforced))
}

func (n *PrimaryKeyColumnAttributeNode) Enforced() bool {
	var v bool
	internalparser.ASTPrimaryKeyColumnAttribute_enforced(n.getRaw(), &v)
	return v
}

type ForeignKeyColumnAttributeNode struct {
	*ColumnAttributeBaseNode
}

func (n *ForeignKeyColumnAttributeNode) ConstraintName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTForeignKeyColumnAttribute_constraint_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *ForeignKeyColumnAttributeNode) Reference() *ForeignKeyReferenceNode {
	var v unsafe.Pointer
	internalparser.ASTForeignKeyColumnAttribute_reference(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newForeignKeyReferenceNode(v)
}

type ColumnAttributeListNode struct {
	*BaseNode
}

func (n *ColumnAttributeListNode) Values() []ColumnAttributeNode {
	var num int
	internalparser.ASTColumnAttributeList_values_num(n.getRaw(), &num)
	ret := make([]ColumnAttributeNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTColumnAttributeList_value(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ColumnAttributeNode))
	}
	return ret
}

type StructColumnFieldNode struct {
	*BaseNode
}

func (n *StructColumnFieldNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTStructColumnField_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *StructColumnFieldNode) Schema() *ColumnSchemaNode {
	var v unsafe.Pointer
	internalparser.ASTStructColumnField_schema(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnSchemaNode(v)
}

type GeneratedColumnInfoStoredMode int

const (
	GeneratedColumnInfoNonStored GeneratedColumnInfoStoredMode = iota
	GeneratedColumnInfoStored
	GeneratedColumnInfoStoredVolatile
)

func (m GeneratedColumnInfoStoredMode) String() string {
	switch m {
	case GeneratedColumnInfoNonStored:
		return "NON_STORED"
	case GeneratedColumnInfoStored:
		return "STORED"
	case GeneratedColumnInfoStoredVolatile:
		return "STORED_VOLATILE"
	}
	return ""
}

type GeneratedColumnInfoNode struct {
	*BaseNode
}

func (n *GeneratedColumnInfoNode) SetStoredMode(mode GeneratedColumnInfoStoredMode) {
	internalparser.ASTGeneratedColumnInfo_set_stored_mode(n.getRaw(), int(mode))
}

func (n *GeneratedColumnInfoNode) StoredMode() GeneratedColumnInfoStoredMode {
	var v int
	internalparser.ASTGeneratedColumnInfo_stored_mode(n.getRaw(), &v)
	return GeneratedColumnInfoStoredMode(v)
}

func (n *GeneratedColumnInfoNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTGeneratedColumnInfo_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *GeneratedColumnInfoNode) SqlForStoredMode() string {
	var v unsafe.Pointer
	internalparser.ASTGeneratedColumnInfo_GetSqlForStoredMode(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type TableElementBaseNode struct {
	*BaseNode
}

type ColumnDefinitionNode struct {
	*TableElementBaseNode
}

func (n *ColumnDefinitionNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTColumnDefinition_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *ColumnDefinitionNode) Schema() *ColumnSchemaNode {
	var v unsafe.Pointer
	internalparser.ASTColumnDefinition_schema(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnSchemaNode(v)
}

type TableElementListNode struct {
	*BaseNode
}

func (n *TableElementListNode) Elements() []TableElementNode {
	var num int
	internalparser.ASTTableElementList_elements_num(n.getRaw(), &num)
	ret := make([]TableElementNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTTableElementList_element(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(TableElementNode))
	}
	return ret
}

type ColumnListNode struct {
	*BaseNode
}

func (n *ColumnListNode) Identifiers() []*IdentifierNode {
	var num int
	internalparser.ASTColumnList_identifiers_num(n.getRaw(), &num)
	ret := make([]*IdentifierNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTColumnList_identifier(n.getRaw(), i, &v)
		ret = append(ret, newIdentifierNode(v))
	}
	return ret
}

type ColumnRelativePositionType int

const (
	ColumnRelativePositionPreceding ColumnRelativePositionType = iota
	ColumnRelativePositionFollowing
)

func (t ColumnRelativePositionType) String() string {
	switch t {
	case ColumnRelativePositionPreceding:
		return "PRECEDING"
	case ColumnRelativePositionFollowing:
		return "FOLLOWING"
	}
	return ""
}

type ColumnPositionNode struct {
	*BaseNode
}

func (n *ColumnPositionNode) SetType(typ ColumnRelativePositionType) {
	internalparser.ASTColumnPosition_set_type(n.getRaw(), int(typ))
}

func (n *ColumnPositionNode) Type() ColumnRelativePositionType {
	var v int
	internalparser.ASTColumnPosition_type(n.getRaw(), &v)
	return ColumnRelativePositionType(v)
}

func (n *ColumnPositionNode) Identifier() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTColumnPosition_identifier(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type InsertValuesRowNode struct {
	*BaseNode
}

func (n *InsertValuesRowNode) Values() []ExpressionNode {
	var num int
	internalparser.ASTInsertValuesRow_values_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTInsertValuesRow_value(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

type InsertValuesRowListNode struct {
	*BaseNode
}

func (n *InsertValuesRowListNode) Rows() []*InsertValuesRowNode {
	var num int
	internalparser.ASTInsertValuesRowList_rows_num(n.getRaw(), &num)
	ret := make([]*InsertValuesRowNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTInsertValuesRowList_row(n.getRaw(), i, &v)
		ret = append(ret, newInsertValuesRowNode(v))
	}
	return ret
}

type InsertMode int

const (
	InsertDefaultMode InsertMode = iota
	InsertReplaceMode
	InsertUpdateMode
	InsertIgnoreMode
)

func (m InsertMode) String() string {
	switch m {
	case InsertDefaultMode:
		return "DEFAULT_MODE"
	case InsertReplaceMode:
		return "REPLACE"
	case InsertUpdateMode:
		return "UPDATE"
	case InsertIgnoreMode:
		return "IGNORE"
	}
	return ""
}

type ParseProgress int

const (
	ParseProgressInitial ParseProgress = iota
	ParseProgressSeenOrIgnoreReplaceUpdate
	ParseProgressSeenTargetPath
	ParseProgressSeenColumnList
	ParseProgressSeenValuesList
)

type InsertStatementNode struct {
	*StatementBaseNode
}

func (n *InsertStatementNode) SetParseProgress(progress ParseProgress) {
	internalparser.ASTInsertStatement_set_parse_progress(n.getRaw(), int(progress))
}

func (n *InsertStatementNode) ParseProgress() ParseProgress {
	var v int
	internalparser.ASTInsertStatement_parse_progress(n.getRaw(), &v)
	return ParseProgress(v)
}

func (n *InsertStatementNode) SetInsertMode(mode InsertMode) {
	internalparser.ASTInsertStatement_set_insert_mode(n.getRaw(), int(mode))
}

func (n *InsertStatementNode) InsertMode() InsertMode {
	var v int
	internalparser.ASTInsertStatement_insert_mode(n.getRaw(), &v)
	return InsertMode(v)
}

func (n *InsertStatementNode) TargetPath() GeneralizedPathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTInsertStatement_target_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(GeneralizedPathExpressionNode)
}

func (n *InsertStatementNode) ColumnList() *ColumnListNode {
	var v unsafe.Pointer
	internalparser.ASTInsertStatement_column_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnListNode(v)
}

func (n *InsertStatementNode) Rows() *InsertValuesRowListNode {
	var v unsafe.Pointer
	internalparser.ASTInsertStatement_rows(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newInsertValuesRowListNode(v)
}

func (n *InsertStatementNode) Query() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTInsertStatement_query(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

func (n *InsertStatementNode) AssertRowsModified() *AssertRowsModifiedNode {
	var v unsafe.Pointer
	internalparser.ASTInsertStatement_assert_rows_modified(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAssertRowsModifiedNode(v)
}

func (n *InsertStatementNode) Returning() *ReturningClauseNode {
	var v unsafe.Pointer
	internalparser.ASTInsertStatement_returning(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newReturningClauseNode(v)
}

func (n *InsertStatementNode) SQLForInsertMode() string {
	var v unsafe.Pointer
	internalparser.ASTInsertStatement_GetSQLForInsertMode(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type UpdateSetValueNode struct {
	*BaseNode
}

func (n *UpdateSetValueNode) Path() GeneralizedPathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateSetValue_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(GeneralizedPathExpressionNode)
}

func (n *UpdateSetValueNode) Value() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateSetValue_value(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type UpdateItemNode struct {
	*BaseNode
}

func (n *UpdateItemNode) SetValue() *UpdateSetValueNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateItem_set_value(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newUpdateSetValueNode(v)
}

func (n *UpdateItemNode) InsertStatement() *InsertStatementNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateItem_insert_statement(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newInsertStatementNode(v)
}

func (n *UpdateItemNode) DeleteStatement() *DeleteStatementNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateItem_delete_statement(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newDeleteStatementNode(v)
}

func (n *UpdateItemNode) UpdateStatement() *UpdateStatementNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateItem_update_statement(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newUpdateStatementNode(v)
}

type UpdateItemListNode struct {
	*BaseNode
}

func (n *UpdateItemListNode) UpdateItems() []*UpdateItemNode {
	var num int
	internalparser.ASTUpdateItemList_update_items_num(n.getRaw(), &num)
	ret := make([]*UpdateItemNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTUpdateItemList_update_item(n.getRaw(), i, &v)
		ret = append(ret, newUpdateItemNode(v))
	}
	return ret
}

type UpdateStatementNode struct {
	*StatementBaseNode
}

func (n *UpdateStatementNode) TargetPath() GeneralizedPathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateStatement_target_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(GeneralizedPathExpressionNode)
}

func (n *UpdateStatementNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateStatement_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

func (n *UpdateStatementNode) Offset() *WithOffsetNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateStatement_offset(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWithOffsetNode(v)
}

func (n *UpdateStatementNode) UpdateItemList() *UpdateItemListNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateStatement_update_item_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newUpdateItemListNode(v)
}

func (n *UpdateStatementNode) FromClause() *FromClauseNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateStatement_from_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newFromClauseNode(v)
}

func (n *UpdateStatementNode) Where() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateStatement_where(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *UpdateStatementNode) AssertRowsModified() *AssertRowsModifiedNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateStatement_assert_rows_modified(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAssertRowsModifiedNode(v)
}

func (n *UpdateStatementNode) Returning() *ReturningClauseNode {
	var v unsafe.Pointer
	internalparser.ASTUpdateStatement_returning(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newReturningClauseNode(v)
}

type TrucateStatementNode struct {
	*StatementBaseNode
}

func (n *TrucateStatementNode) TargetPath() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTTruncateStatement_target_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *TrucateStatementNode) Where() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTTruncateStatement_where(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type MergeActionType int

const (
	MergeActionNotSet MergeActionType = iota
	MergeActionInsert
	MergeActionUpdate
	MergeActionDelete
)

type MergeActionNode struct {
	*BaseNode
}

func (n *MergeActionNode) SetActionType(typ MergeActionType) {
	internalparser.ASTMergeAction_set_action_type(n.getRaw(), int(typ))
}

func (n *MergeActionNode) ActionType() MergeActionType {
	var v int
	internalparser.ASTMergeAction_action_type(n.getRaw(), &v)
	return MergeActionType(v)
}

func (n *MergeActionNode) InsertColumnList() *ColumnListNode {
	var v unsafe.Pointer
	internalparser.ASTMergeAction_insert_column_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnListNode(v)
}

func (n *MergeActionNode) InsertRow() *InsertValuesRowNode {
	var v unsafe.Pointer
	internalparser.ASTMergeAction_insert_row(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newInsertValuesRowNode(v)
}

func (n *MergeActionNode) UpdateItemList() *UpdateItemListNode {
	var v unsafe.Pointer
	internalparser.ASTMergeAction_update_item_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newUpdateItemListNode(v)
}

type MergeMatchType int

const (
	MergeMatchNotSet MergeMatchType = iota
	MergeMatched
	MergeNotMatchedBySource
	MergeNotMatchedByTarget
)

type MergeWhenClauseNode struct {
	*BaseNode
}

func (n *MergeWhenClauseNode) SetMatchType(matchType MergeMatchType) {
	internalparser.ASTMergeWhenClause_set_match_type(n.getRaw(), int(matchType))
}

func (n *MergeWhenClauseNode) MatchType() MergeMatchType {
	var v int
	internalparser.ASTMergeWhenClause_match_type(n.getRaw(), &v)
	return MergeMatchType(v)
}

func (n *MergeWhenClauseNode) SearchCondition() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTMergeWhenClause_search_condition(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *MergeWhenClauseNode) Action() *MergeActionNode {
	var v unsafe.Pointer
	internalparser.ASTMergeWhenClause_action(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newMergeActionNode(v)
}

func (n *MergeWhenClauseNode) SQLForMatchType() string {
	var v unsafe.Pointer
	internalparser.ASTMergeWhenClause_GetSQLForMatchType(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type MergeWhenClauseListNode struct {
	*BaseNode
}

func (n *MergeWhenClauseListNode) ClauseList() []*MergeWhenClauseNode {
	var num int
	internalparser.ASTMergeWhenClauseList_clause_list_num(n.getRaw(), &num)
	ret := make([]*MergeWhenClauseNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTMergeWhenClauseList_clause_list(n.getRaw(), i, &v)
		ret = append(ret, newMergeWhenClauseNode(v))
	}
	return ret
}

type MergeStatementNode struct {
	*StatementBaseNode
}

func (n *MergeStatementNode) TargetPath() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTMergeStatement_target_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *MergeStatementNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTMergeStatement_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

func (n *MergeStatementNode) TableExpression() TableExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTMergeStatement_table_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(TableExpressionNode)
}

func (n *MergeStatementNode) MergeCondition() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTMergeStatement_merge_condition(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *MergeStatementNode) WhenClauses() *MergeWhenClauseListNode {
	var v unsafe.Pointer
	internalparser.ASTMergeStatement_when_clauses(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newMergeWhenClauseListNode(v)
}

type PrivilegeNode struct {
	*BaseNode
}

func (n *PrivilegeNode) PrivilegeAction() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTPrivilege_privilege_action(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *PrivilegeNode) Paths() *PathExpressionListNode {
	var v unsafe.Pointer
	internalparser.ASTPrivilege_paths(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionListNode(v)
}

type PrivilegesNode struct {
	*BaseNode
}

func (n *PrivilegesNode) Privileges() []*PrivilegeNode {
	var num int
	internalparser.ASTPrivileges_privileges_num(n.getRaw(), &num)
	ret := make([]*PrivilegeNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTPrivileges_privilege(n.getRaw(), i, &v)
		ret = append(ret, newPrivilegeNode(v))
	}
	return ret
}

func (n *PrivilegesNode) IsAllPrivileges() bool {
	var v bool
	internalparser.ASTPrivileges_is_all_privileges(n.getRaw(), &v)
	return v
}

type GranteeListNode struct {
	*BaseNode
}

func (n *GranteeListNode) GranteeList() []ExpressionNode {
	var num int
	internalparser.ASTGranteeList_grantee_list_num(n.getRaw(), &num)
	ret := make([]ExpressionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTGranteeList_grantee_list(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(ExpressionNode))
	}
	return ret
}

type GrantStatementNode struct {
	*StatementBaseNode
}

func (n *GrantStatementNode) Privileges() *PrivilegesNode {
	var v unsafe.Pointer
	internalparser.ASTGrantStatement_privileges(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPrivilegesNode(v)
}

func (n *GrantStatementNode) TargetType() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTGrantStatement_target_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *GrantStatementNode) TargetPath() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTGrantStatement_target_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *GrantStatementNode) GranteeList() *GranteeListNode {
	var v unsafe.Pointer
	internalparser.ASTGrantStatement_grantee_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newGranteeListNode(v)
}

type RevokeStatementNode struct {
	*StatementBaseNode
}

func (n *RevokeStatementNode) Privileges() *PrivilegesNode {
	var v unsafe.Pointer
	internalparser.ASTRevokeStatement_privileges(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPrivilegesNode(v)
}

func (n *RevokeStatementNode) TargetType() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTRevokeStatement_target_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *RevokeStatementNode) TargetPath() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTRevokeStatement_target_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *RevokeStatementNode) GranteeList() *GranteeListNode {
	var v unsafe.Pointer
	internalparser.ASTRevokeStatement_grantee_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newGranteeListNode(v)
}

type RepeatableClauseNode struct {
	*BaseNode
}

func (n *RepeatableClauseNode) Argument() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTRepeatableClause_argument(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type FilterType int

const (
	FilterNotSet FilterType = iota
	FilterInclude
	FilterExclude
)

func (t FilterType) String() string {
	switch t {
	case FilterNotSet:
		return "NOT_SET"
	case FilterInclude:
		return "INCLUDE"
	case FilterExclude:
		return "EXCLUDE"
	}
	return ""
}

type FilterFieldsArgNode struct {
	*BaseNode
}

func (n *FilterFieldsArgNode) SetFilterType(typ FilterType) {
	internalparser.ASTFilterFieldsArg_set_filter_type(n.getRaw(), int(typ))
}

func (n *FilterFieldsArgNode) FilterType() FilterType {
	var v int
	internalparser.ASTFilterFieldsArg_filter_type(n.getRaw(), &v)
	return FilterType(v)
}

func (n *FilterFieldsArgNode) PathExpression() GeneralizedPathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTFilterFieldsArg_path_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(GeneralizedPathExpressionNode)
}

func (n *FilterFieldsArgNode) SQLForOperator() string {
	var v unsafe.Pointer
	internalparser.ASTFilterFieldsArg_GetSQLForOperator(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type ReplaceFieldsArgNode struct {
	*BaseNode
}

func (n *ReplaceFieldsArgNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTReplaceFieldsArg_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *ReplaceFieldsArgNode) PathExpression() GeneralizedPathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTReplaceFieldsArg_path_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(GeneralizedPathExpressionNode)
}

type ReplaceFieldsExpressionNode struct {
	*ExpressionBaseNode
}

func (n *ReplaceFieldsExpressionNode) Expr() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTReplaceFieldsExpression_expr(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *ReplaceFieldsExpressionNode) Arguments() []*ReplaceFieldsArgNode {
	var num int
	internalparser.ASTReplaceFieldsExpression_arguments_num(n.getRaw(), &num)
	ret := make([]*ReplaceFieldsArgNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTReplaceFieldsExpression_argument(n.getRaw(), i, &v)
		ret = append(ret, newReplaceFieldsArgNode(v))
	}
	return ret
}

type SampleSizeUnit int

const (
	SampleSizeNotSet SampleSizeUnit = iota
	SampleSizeRows
	SampleSizePercent
)

func (u SampleSizeUnit) String() string {
	switch u {
	case SampleSizeNotSet:
		return "NOT_SET"
	case SampleSizeRows:
		return "ROWS"
	case SampleSizePercent:
		return "PERCENT"
	}
	return ""
}

type SampleSizeNode struct {
	*BaseNode
}

func (n *SampleSizeNode) SetUnit(unit SampleSizeUnit) {
	internalparser.ASTSampleSize_set_unit(n.getRaw(), int(unit))
}

func (n *SampleSizeNode) Unit() SampleSizeUnit {
	var v int
	internalparser.ASTSampleSize_unit(n.getRaw(), &v)
	return SampleSizeUnit(v)
}

func (n *SampleSizeNode) Size() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTSampleSize_size(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *SampleSizeNode) PartitionBy() *PartitionByNode {
	var v unsafe.Pointer
	internalparser.ASTSampleSize_partition_by(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPartitionByNode(v)
}

func (n *SampleSizeNode) SQLForUnit() string {
	var v unsafe.Pointer
	internalparser.ASTSampleSize_GetSQLForUnit(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type WithWeightNode struct {
	*BaseNode
}

func (n *WithWeightNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTWithWeight_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

type SampleSuffixNode struct {
	*BaseNode
}

func (n *SampleSuffixNode) Weight() *WithWeightNode {
	var v unsafe.Pointer
	internalparser.ASTSampleSuffix_weight(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWithWeightNode(v)
}

func (n *SampleSuffixNode) Repeat() *RepeatableClauseNode {
	var v unsafe.Pointer
	internalparser.ASTSampleSuffix_repeat(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newRepeatableClauseNode(v)
}

type SampleClauseNode struct {
	*BaseNode
}

func (n *SampleClauseNode) SampleMethod() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTSampleClause_sample_method(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *SampleClauseNode) SampleSize() *SampleSizeNode {
	var v unsafe.Pointer
	internalparser.ASTSampleClause_sample_size(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newSampleSizeNode(v)
}

func (n *SampleClauseNode) SampleSuffix() *SampleSuffixNode {
	var v unsafe.Pointer
	internalparser.ASTSampleClause_sample_suffix(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newSampleSuffixNode(v)
}

type AlterActionBaseNode struct {
	*BaseNode
}

func (n *AlterActionBaseNode) SQLForAlterAction() string {
	var v unsafe.Pointer
	internalparser.ASTAlterAction_GetSQLForAlterAction(n.getRaw(), &v)
	return C.GoString((*C.char)(v))
}

type SetOptionsActionNode struct {
	*AlterActionBaseNode
}

func (n *SetOptionsActionNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTSetOptionsAction_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type SetAsActionNode struct {
	*AlterActionBaseNode
}

func (n *SetAsActionNode) JSONBody() *JSONLiteralNode {
	var v unsafe.Pointer
	internalparser.ASTSetAsAction_json_body(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newJSONLiteralNode(v)
}

func (n *SetAsActionNode) TextBody() *StringLiteralNode {
	var v unsafe.Pointer
	internalparser.ASTSetAsAction_text_body(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStringLiteralNode(v)
}

type AddConstraintActionNode struct {
	*AlterActionBaseNode
}

func (n *AddConstraintActionNode) SetIsIfNotExists(isIfNotExists bool) {
	internalparser.ASTAddConstraintAction_set_is_if_not_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *AddConstraintActionNode) IsIfNotExists() bool {
	var v bool
	internalparser.ASTAddConstraintAction_is_if_not_exists(n.getRaw(), &v)
	return v
}

func (n *AddConstraintActionNode) Constraint() TableConstraintNode {
	var v unsafe.Pointer
	internalparser.ASTAddConstraintAction_constraint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(TableConstraintNode)
}

type DropPrimaryKeyActionNode struct {
	*AlterActionBaseNode
}

func (n *DropPrimaryKeyActionNode) SetIsIfExists(isIfNotExists bool) {
	internalparser.ASTDropPrimaryKeyAction_set_is_if_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *DropPrimaryKeyActionNode) IsIfExists() bool {
	var v bool
	internalparser.ASTDropPrimaryKeyAction_is_if_exists(n.getRaw(), &v)
	return v
}

type DropConstraintActionNode struct {
	*AlterActionBaseNode
}

func (n *DropConstraintActionNode) SetIsIfExists(isIfNotExists bool) {
	internalparser.ASTDropConstraintAction_set_is_if_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *DropConstraintActionNode) IsIfExists() bool {
	var v bool
	internalparser.ASTDropConstraintAction_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *DropConstraintActionNode) ConstraintName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTDropConstraintAction_constraint_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type AlterConstraintEnforcementActionNode struct {
	*AlterActionBaseNode
}

func (n *AlterConstraintEnforcementActionNode) SetIsIfExists(isIfNotExists bool) {
	internalparser.ASTAlterConstraintEnforcementAction_set_is_if_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *AlterConstraintEnforcementActionNode) IsIfExists() bool {
	var v bool
	internalparser.ASTAlterConstraintEnforcementAction_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *AlterConstraintEnforcementActionNode) SetIsEnforced(enforced bool) {
	internalparser.ASTAlterConstraintEnforcementAction_set_is_enforced(n.getRaw(), boolToInt(enforced))
}

func (n *AlterConstraintEnforcementActionNode) IsEnforced() bool {
	var v bool
	internalparser.ASTAlterConstraintEnforcementAction_is_enforced(n.getRaw(), &v)
	return v
}

func (n *AlterConstraintEnforcementActionNode) ConstraintName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTAlterConstraintEnforcementAction_constraint_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type AlterConstraintSetOptionsActionNode struct {
	*AlterActionBaseNode
}

func (n *AlterConstraintSetOptionsActionNode) SetIsIfExists(isIfNotExists bool) {
	internalparser.ASTAlterConstraintSetOptionsAction_set_is_if_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *AlterConstraintSetOptionsActionNode) IsIfExists() bool {
	var v bool
	internalparser.ASTAlterConstraintSetOptionsAction_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *AlterConstraintSetOptionsActionNode) ConstraintName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTAlterConstraintSetOptionsAction_constraint_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *AlterConstraintSetOptionsActionNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTAlterConstraintSetOptionsAction_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type AddColumnActionNode struct {
	*AlterActionBaseNode
}

func (n *AddColumnActionNode) SetIsIfNotExists(isIfNotExists bool) {
	internalparser.ASTAddColumnAction_set_is_if_not_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *AddColumnActionNode) IsIfNotExists() bool {
	var v bool
	internalparser.ASTAddColumnAction_is_if_not_exists(n.getRaw(), &v)
	return v
}

func (n *AddColumnActionNode) ColumnDefinition() *ColumnDefinitionNode {
	var v unsafe.Pointer
	internalparser.ASTAddColumnAction_column_definition(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnDefinitionNode(v)
}

func (n *AddColumnActionNode) ColumnPosition() *ColumnPositionNode {
	var v unsafe.Pointer
	internalparser.ASTAddColumnAction_column_position(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnPositionNode(v)
}

func (n *AddColumnActionNode) FillExpression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTAddColumnAction_fill_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type DropColumnActionNode struct {
	*AlterActionBaseNode
}

func (n *DropColumnActionNode) SetIsIfExists(isIfNotExists bool) {
	internalparser.ASTDropColumnAction_set_is_if_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *DropColumnActionNode) IsIfExists() bool {
	var v bool
	internalparser.ASTDropColumnAction_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *DropColumnActionNode) ColumnName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTDropColumnAction_column_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type RenameColumnActionNode struct {
	*AlterActionBaseNode
}

func (n *RenameColumnActionNode) SetIsIfExists(isIfNotExists bool) {
	internalparser.ASTRenameColumnAction_set_is_if_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *RenameColumnActionNode) IsIfExists() bool {
	var v bool
	internalparser.ASTRenameColumnAction_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *RenameColumnActionNode) ColumnName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTRenameColumnAction_column_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *RenameColumnActionNode) NewColumnName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTRenameColumnAction_new_column_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type AlterColumnTypeActionNode struct {
	*AlterActionBaseNode
}

func (n *AlterColumnTypeActionNode) SetIsIfExists(isIfNotExists bool) {
	internalparser.ASTAlterColumnTypeAction_set_is_if_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *AlterColumnTypeActionNode) IsIfExists() bool {
	var v bool
	internalparser.ASTAlterColumnTypeAction_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *AlterColumnTypeActionNode) ColumnName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTAlterColumnTypeAction_column_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *AlterColumnTypeActionNode) Schema() *ColumnSchemaNode {
	var v unsafe.Pointer
	internalparser.ASTAlterColumnTypeAction_schema(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnSchemaNode(v)
}

func (n *AlterColumnTypeActionNode) Collate() *CollateNode {
	var v unsafe.Pointer
	internalparser.ASTAlterColumnTypeAction_collate(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newCollateNode(v)
}

type AlterColumnOptionsActionNode struct {
	*AlterActionBaseNode
}

func (n *AlterColumnOptionsActionNode) SetIsIfExists(isIfNotExists bool) {
	internalparser.ASTAlterColumnOptionsAction_set_is_if_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *AlterColumnOptionsActionNode) IsIfExists() bool {
	var v bool
	internalparser.ASTAlterColumnOptionsAction_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *AlterColumnOptionsActionNode) ColumnName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTAlterColumnOptionsAction_column_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *AlterColumnOptionsActionNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTAlterColumnOptionsAction_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type AlterColumnSetDefaultActionNode struct {
	*AlterActionBaseNode
}

func (n *AlterColumnSetDefaultActionNode) SetIsIfExists(isIfNotExists bool) {
	internalparser.ASTAlterColumnSetDefaultAction_set_is_if_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *AlterColumnSetDefaultActionNode) IsIfExists() bool {
	var v bool
	internalparser.ASTAlterColumnSetDefaultAction_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *AlterColumnSetDefaultActionNode) ColumnName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTAlterColumnSetDefaultAction_column_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *AlterColumnSetDefaultActionNode) DefaultExpression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTAlterColumnSetDefaultAction_default_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type AlterColumnDropDefaultActionNode struct {
	*AlterActionBaseNode
}

func (n *AlterColumnDropDefaultActionNode) SetIsIfExists(isIfNotExists bool) {
	internalparser.ASTAlterColumnDropDefaultAction_set_is_if_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *AlterColumnDropDefaultActionNode) IsIfExists() bool {
	var v bool
	internalparser.ASTAlterColumnDropDefaultAction_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *AlterColumnDropDefaultActionNode) ColumnName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTAlterColumnDropDefaultAction_column_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type AlterColumnDropNotNullActionNode struct {
	*AlterActionBaseNode
}

func (n *AlterColumnDropNotNullActionNode) SetIsIfExists(isIfNotExists bool) {
	internalparser.ASTAlterColumnDropNotNullAction_set_is_if_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *AlterColumnDropNotNullActionNode) IsIfExists() bool {
	var v bool
	internalparser.ASTAlterColumnDropNotNullAction_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *AlterColumnDropNotNullActionNode) ColumnName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTAlterColumnDropNotNullAction_column_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type GrantToClauseNode struct {
	*AlterActionBaseNode
}

func (n *GrantToClauseNode) SetHasGrantKeywordAndParens(v bool) {
	internalparser.ASTGrantToClause_set_has_grant_keyword_and_parens(n.getRaw(), boolToInt(v))
}

func (n *GrantToClauseNode) HasGrantKeywordAndParens() bool {
	var v bool
	internalparser.ASTGrantToClause_has_grant_keyword_and_parens(n.getRaw(), &v)
	return v
}

func (n *GrantToClauseNode) GranteeList() *GranteeListNode {
	var v unsafe.Pointer
	internalparser.ASTGrantToClause_grantee_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newGranteeListNode(v)
}

type RestrictToClauseNode struct {
	*AlterActionBaseNode
}

func (n *RestrictToClauseNode) RestricteeList() *GranteeListNode {
	var v unsafe.Pointer
	internalparser.ASTRestrictToClause_restrictee_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newGranteeListNode(v)
}

type AddToRestricteeListClauseNode struct {
	*AlterActionBaseNode
}

func (n *AddToRestricteeListClauseNode) SetIsIfNotExists(isIfNotExists bool) {
	internalparser.ASTAddToRestricteeListClause_set_is_if_not_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *AddToRestricteeListClauseNode) IsIfNotExists() bool {
	var v bool
	internalparser.ASTAddToRestricteeListClause_is_if_not_exists(n.getRaw(), &v)
	return v
}

func (n *AddToRestricteeListClauseNode) RestricteeList() *GranteeListNode {
	var v unsafe.Pointer
	internalparser.ASTAddToRestricteeListClause_restrictee_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newGranteeListNode(v)
}

type RemoveFromRestricteeListClauseNode struct {
	*AlterActionBaseNode
}

func (n *RemoveFromRestricteeListClauseNode) SetIsIfExists(isIfNotExists bool) {
	internalparser.ASTRemoveFromRestricteeListClause_set_is_if_exists(n.getRaw(), boolToInt(isIfNotExists))
}

func (n *RemoveFromRestricteeListClauseNode) IsIfExists() bool {
	var v bool
	internalparser.ASTRemoveFromRestricteeListClause_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *RemoveFromRestricteeListClauseNode) RestricteeList() *GranteeListNode {
	var v unsafe.Pointer
	internalparser.ASTRemoveFromRestricteeListClause_restrictee_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newGranteeListNode(v)
}

type FilterUsingClauseNode struct {
	*AlterActionBaseNode
}

func (n *FilterUsingClauseNode) SetHasFilterKeyword(keyword bool) {
	internalparser.ASTFilterUsingClause_set_has_filter_keyword(n.getRaw(), boolToInt(keyword))
}

func (n *FilterUsingClauseNode) HasFilterKeyword() bool {
	var v bool
	internalparser.ASTFilterUsingClause_has_filter_keyword(n.getRaw(), &v)
	return v
}

func (n *FilterUsingClauseNode) Predicate() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTFilterUsingClause_predicate(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type RevokeFromClauseNode struct {
	*AlterActionBaseNode
}

func (n *RevokeFromClauseNode) SetIsRevokeFromAll(v bool) {
	internalparser.ASTRevokeFromClause_set_is_revoke_from_all(n.getRaw(), boolToInt(v))
}

func (n *RevokeFromClauseNode) IsRevokeFromAll() bool {
	var v bool
	internalparser.ASTRevokeFromClause_is_revoke_from_all(n.getRaw(), &v)
	return v
}

func (n *RevokeFromClauseNode) RevokeFromList() *GranteeListNode {
	var v unsafe.Pointer
	internalparser.ASTRevokeFromClause_revoke_from_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newGranteeListNode(v)
}

type RenameToClauseNode struct {
	*AlterActionBaseNode
}

func (n *RenameToClauseNode) NewName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTRenameToClause_new_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type SetCollateClauseNode struct {
	*AlterActionBaseNode
}

func (n *SetCollateClauseNode) Collate() *CollateNode {
	var v unsafe.Pointer
	internalparser.ASTSetCollateClause_collate(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newCollateNode(v)
}

type AlterActionListNode struct {
	*BaseNode
}

func (n *AlterActionListNode) Actions() []AlterActionNode {
	var num int
	internalparser.ASTAlterActionList_actions_num(n.getRaw(), &num)
	ret := make([]AlterActionNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTAlterActionList_action(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(AlterActionNode))
	}
	return ret
}

type AlterAllRowAccessPoliciesStatementNode struct {
	*StatementBaseNode
}

func (n *AlterAllRowAccessPoliciesStatementNode) TableNamePath() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTAlterAllRowAccessPoliciesStatement_table_name_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *AlterAllRowAccessPoliciesStatementNode) AlterAction() AlterActionNode {
	var v unsafe.Pointer
	internalparser.ASTAlterAllRowAccessPoliciesStatement_alter_action(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(AlterActionNode)
}

type ForeignKeyAction int

const (
	ForeignKeyNoAction ForeignKeyAction = iota
	ForeignKeyRestrictAction
	ForeignKeyCascadeAction
	ForeignKeySetNullAction
)

func (a ForeignKeyAction) String() string {
	switch a {
	case ForeignKeyNoAction:
		return "NO_ACTION"
	case ForeignKeyRestrictAction:
		return "RESTRICT"
	case ForeignKeyCascadeAction:
		return "CASCADE"
	case ForeignKeySetNullAction:
		return "SET_NULL"
	}
	return ""
}

type ForeignKeyActionsNode struct {
	*BaseNode
}

func (n *ForeignKeyActionsNode) SetUpdateAction(action ForeignKeyAction) {
	internalparser.ASTForeignKeyActions_set_udpate_action(n.getRaw(), int(action))
}

func (n *ForeignKeyActionsNode) UpdateAction() ForeignKeyAction {
	var v int
	internalparser.ASTForeignKeyActions_udpate_action(n.getRaw(), &v)
	return ForeignKeyAction(v)
}

func (n *ForeignKeyActionsNode) SetDeleteAction(action ForeignKeyAction) {
	internalparser.ASTForeignKeyActions_set_delete_action(n.getRaw(), int(action))
}

func (n *ForeignKeyActionsNode) DeleteAction() ForeignKeyAction {
	var v int
	internalparser.ASTForeignKeyActions_delete_action(n.getRaw(), &v)
	return ForeignKeyAction(v)
}

type ForeignKeyReferenceMatch int

const (
	ForeignKeyReferenceSimple ForeignKeyReferenceMatch = iota
	ForeignKeyReferenceFull
	ForeignKeyReferenceNotDistinct
)

func (m ForeignKeyReferenceMatch) String() string {
	switch m {
	case ForeignKeyReferenceSimple:
		return "SIMPLE"
	case ForeignKeyReferenceFull:
		return "FULL"
	case ForeignKeyReferenceNotDistinct:
		return "NOT_DISTINCT"
	}
	return ""
}

type ForeignKeyReferenceNode struct {
	*BaseNode
}

func (n *ForeignKeyReferenceNode) SetMatch(match ForeignKeyReferenceMatch) {
	internalparser.ASTForeignKeyReference_set_match(n.getRaw(), int(match))
}

func (n *ForeignKeyReferenceNode) Match() ForeignKeyReferenceMatch {
	var v int
	internalparser.ASTForeignKeyReference_match(n.getRaw(), &v)
	return ForeignKeyReferenceMatch(v)
}

func (n *ForeignKeyReferenceNode) SetEnforced(enforced bool) {
	internalparser.ASTForeignKeyReference_set_enforced(n.getRaw(), boolToInt(enforced))
}

func (n *ForeignKeyReferenceNode) Enforced() bool {
	var v bool
	internalparser.ASTForeignKeyReference_enforced(n.getRaw(), &v)
	return v
}

func (n *ForeignKeyReferenceNode) TableName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTForeignKeyReference_table_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *ForeignKeyReferenceNode) ColumnList() *ColumnListNode {
	var v unsafe.Pointer
	internalparser.ASTForeignKeyReference_column_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnListNode(v)
}

func (n *ForeignKeyReferenceNode) Actions() *ForeignKeyActionsNode {
	var v unsafe.Pointer
	internalparser.ASTForeignKeyReference_actions(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newForeignKeyActionsNode(v)
}

type ScriptBaseNode struct {
	*BaseNode
}

func (n *ScriptBaseNode) StatementListNode() *StatementListNode {
	var v unsafe.Pointer
	internalparser.ASTScript_statement_list_node(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStatementListNode(v)
}

func (n *ScriptBaseNode) StatementList() []StatementNode {
	var num int
	internalparser.ASTScript_statement_list_num(n.getRaw(), &num)
	ret := make([]StatementNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTScript_statement_list(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(StatementNode))
	}
	return ret
}

type ElseifClauseNode struct {
	*BaseNode
}

func (n *ElseifClauseNode) Condition() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTElseifClause_condition(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *ElseifClauseNode) Body() *StatementListNode {
	var v unsafe.Pointer
	internalparser.ASTElseifClause_body(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStatementListNode(v)
}

func (n *ElseifClauseNode) IfStmt() *IfStatementNode {
	var v unsafe.Pointer
	internalparser.ASTElseifClause_if_stmt(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIfStatementNode(v)
}

type ElseifClauseListNode struct {
	*BaseNode
}

func (n *ElseifClauseListNode) ElseifClauses() []*ElseifClauseNode {
	var num int
	internalparser.ASTElseifClauseList_elseif_clauses_num(n.getRaw(), &num)
	ret := make([]*ElseifClauseNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTElseifClauseList_elseif_clause(n.getRaw(), i, &v)
		ret = append(ret, newElseifClauseNode(v))
	}
	return ret
}

type IfStatementNode struct {
	*ScriptStatementNode
}

func (n *IfStatementNode) Condition() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTIfStatement_condition(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *IfStatementNode) ThenList() *StatementListNode {
	var v unsafe.Pointer
	internalparser.ASTIfStatement_then_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStatementListNode(v)
}

func (n *IfStatementNode) ElseifClauses() *ElseifClauseListNode {
	var v unsafe.Pointer
	internalparser.ASTIfStatement_elseif_clauses(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newElseifClauseListNode(v)
}

func (n *IfStatementNode) ElseList() *StatementListNode {
	var v unsafe.Pointer
	internalparser.ASTIfStatement_else_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStatementListNode(v)
}

type WhenThenClauseNode struct {
	*BaseNode
}

func (n *WhenThenClauseNode) Condition() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTWhenThenClause_condition(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *WhenThenClauseNode) Body() *StatementListNode {
	var v unsafe.Pointer
	internalparser.ASTWhenThenClause_body(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStatementListNode(v)
}

func (n *WhenThenClauseNode) CaseStmt() *CaseStatementNode {
	var v unsafe.Pointer
	internalparser.ASTWhenThenClause_case_stmt(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newCaseStatementNode(v)
}

type WhenThenClauseListNode struct {
	*BaseNode
}

func (n *WhenThenClauseListNode) WhenThenClauses() []*WhenThenClauseNode {
	var num int
	internalparser.ASTWhenThenClauseList_when_then_clauses_num(n.getRaw(), &num)
	ret := make([]*WhenThenClauseNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTWhenThenClauseList_when_then_clause(n.getRaw(), i, &v)
		ret = append(ret, newWhenThenClauseNode(v))
	}
	return ret
}

type CaseStatementNode struct {
	*ScriptStatementNode
}

func (n *CaseStatementNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCaseStatement_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *CaseStatementNode) WhenThenClauses() *WhenThenClauseListNode {
	var v unsafe.Pointer
	internalparser.ASTCaseStatement_when_then_clauses(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWhenThenClauseListNode(v)
}

func (n *CaseStatementNode) ElseList() *StatementListNode {
	var v unsafe.Pointer
	internalparser.ASTCaseStatement_else_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStatementListNode(v)
}

type HintNode struct {
	*BaseNode
}

func (n *HintNode) NumShardsHint() *IntLiteralNode {
	var v unsafe.Pointer
	internalparser.ASTHint_num_shards_hint(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIntLiteralNode(v)
}

func (n *HintNode) HintEntries() []*HintEntryNode {
	var num int
	internalparser.ASTHint_hint_entries_num(n.getRaw(), &num)
	ret := make([]*HintEntryNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTHint_hint_entry(n.getRaw(), i, &v)
		ret = append(ret, newHintEntryNode(v))
	}
	return ret
}

type HintEntryNode struct {
	*BaseNode
}

func (n *HintEntryNode) Qualifier() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTHintEntry_qualifier(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *HintEntryNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTHintEntry_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *HintEntryNode) Value() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTHintEntry_value(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type UnpivotInItemLabelNode struct {
	*BaseNode
}

func (n *UnpivotInItemLabelNode) Label() LeafNode {
	var v unsafe.Pointer
	internalparser.ASTUnpivotInItemLabel_label(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(LeafNode)
}

type DescriptorNode struct {
	*BaseNode
}

func (n *DescriptorNode) Columns() *DescriptorColumnListNode {
	var v unsafe.Pointer
	internalparser.ASTDescriptor_columns(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newDescriptorColumnListNode(v)
}

type ColumnSchemaNode struct {
	*BaseNode
}

func (n *ColumnSchemaNode) TypeParameters() *TypeParameterListNode {
	var v unsafe.Pointer
	internalparser.ASTColumnSchema_type_parameters(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newTypeParameterListNode(v)
}

func (n *ColumnSchemaNode) GeneratedColumnInfo() *GeneratedColumnInfoNode {
	var v unsafe.Pointer
	internalparser.ASTColumnSchema_generated_column_info(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newGeneratedColumnInfoNode(v)
}

func (n *ColumnSchemaNode) DefaultExpression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTColumnSchema_default_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *ColumnSchemaNode) Collate() *CollateNode {
	var v unsafe.Pointer
	internalparser.ASTColumnSchema_collate(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newCollateNode(v)
}

func (n *ColumnSchemaNode) Attributes() *ColumnAttributeListNode {
	var v unsafe.Pointer
	internalparser.ASTColumnSchema_attributes(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnAttributeListNode(v)
}

func (n *ColumnSchemaNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTColumnSchema_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

func (n *ColumnSchemaNode) ContainsAttribute(kind Kind) bool {
	var v bool
	internalparser.ASTColumnSchema_ContainsAttribute(n.getRaw(), int(kind), &v)
	return v
}

type SimpleColumnSchemaNode struct {
	*ColumnSchemaNode
}

func (n *SimpleColumnSchemaNode) TypeName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTSimpleColumnSchema_type_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type ArrayColumnSchemaNode struct {
	*ColumnSchemaNode
}

func (n *ArrayColumnSchemaNode) ElementSchema() *ColumnSchemaNode {
	var v unsafe.Pointer
	internalparser.ASTArrayColumnSchema_element_schema(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnSchemaNode(v)
}

type TableConstraintBaseNode struct {
	*TableElementBaseNode
}

func (n *TableConstraintBaseNode) ConstraintName() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTTableConstraint_constraint_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type PrimaryKeyNode struct {
	*TableConstraintBaseNode
}

func (n *PrimaryKeyNode) SetEnforced(enforced bool) {
	internalparser.ASTPrimaryKey_set_enforced(n.getRaw(), boolToInt(enforced))
}

func (n *PrimaryKeyNode) Enforced() bool {
	var v bool
	internalparser.ASTPrimaryKey_enforced(n.getRaw(), &v)
	return v
}

func (n *PrimaryKeyNode) ColumnList() *ColumnListNode {
	var v unsafe.Pointer
	internalparser.ASTPrimaryKey_column_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnListNode(v)
}

func (n *PrimaryKeyNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTPrimaryKey_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type ForeignKeyNode struct {
	*TableConstraintBaseNode
}

func (n *ForeignKeyNode) ColumnList() *ColumnListNode {
	var v unsafe.Pointer
	internalparser.ASTForeignKey_column_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnListNode(v)
}

func (n *ForeignKeyNode) Reference() *ForeignKeyReferenceNode {
	var v unsafe.Pointer
	internalparser.ASTForeignKey_reference(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newForeignKeyReferenceNode(v)
}

func (n *ForeignKeyNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTForeignKey_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type CheckConstraintNode struct {
	*TableConstraintBaseNode
}

func (n *CheckConstraintNode) SetIsEnforced(enforced bool) {
	internalparser.ASTCheckConstraint_set_is_enforced(n.getRaw(), boolToInt(enforced))
}

func (n *CheckConstraintNode) IsEnforced() bool {
	var v bool
	internalparser.ASTCheckConstraint_is_enforced(n.getRaw(), &v)
	return v
}

func (n *CheckConstraintNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCheckConstraint_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *CheckConstraintNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTCheckConstraint_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type DescriptorColumnNode struct {
	*BaseNode
}

func (n *DescriptorColumnNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTDescriptorColumn_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type DescriptorColumnListNode struct {
	*BaseNode
}

func (n *DescriptorColumnListNode) DescriptorColumnList() []*DescriptorColumnNode {
	var num int
	internalparser.ASTDescriptorColumnList_descriptor_column_list_num(n.getRaw(), &num)
	ret := make([]*DescriptorColumnNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTDescriptorColumnList_descriptor_column_list(n.getRaw(), i, &v)
		ret = append(ret, newDescriptorColumnNode(v))
	}
	return ret
}

type CreateEntityStatementNode struct {
	*CreateStatementNode
}

func (n *CreateEntityStatementNode) Type() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTCreateEntityStatement_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *CreateEntityStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateEntityStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreateEntityStatementNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateEntityStatement_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

func (n *CreateEntityStatementNode) JSONBody() *JSONLiteralNode {
	var v unsafe.Pointer
	internalparser.ASTCreateEntityStatement_json_body(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newJSONLiteralNode(v)
}

func (n *CreateEntityStatementNode) TextBody() *StringLiteralNode {
	var v unsafe.Pointer
	internalparser.ASTCreateEntityStatement_text_body(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStringLiteralNode(v)
}

type RaiseStatementNode struct {
	*ScriptStatementNode
}

func (n *RaiseStatementNode) Message() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTRaiseStatement_message(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *RaiseStatementNode) IsRethrow() bool {
	var v bool
	internalparser.ASTRaiseStatement_is_rethrow(n.getRaw(), &v)
	return v
}

type ExceptionHandlerNode struct {
	*BaseNode
}

func (n *ExceptionHandlerNode) StatementList() *StatementListNode {
	var v unsafe.Pointer
	internalparser.ASTExceptionHandler_statement_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStatementListNode(v)
}

type ExceptionHandlerListNode struct {
	*BaseNode
}

func (n *ExceptionHandlerListNode) ExceptionHandlerList() []*ExceptionHandlerNode {
	var num int
	internalparser.ASTExceptionHandlerList_exception_handler_list_num(n.getRaw(), &num)
	ret := make([]*ExceptionHandlerNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTExceptionHandlerList_exception_handler_list(n.getRaw(), i, &v)
		ret = append(ret, newExceptionHandlerNode(v))
	}
	return ret
}

type BeginEndBlockNode struct {
	*ScriptStatementNode
}

func (n *BeginEndBlockNode) Label() *LabelNode {
	var v unsafe.Pointer
	internalparser.ASTBeginEndBlock_label(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newLabelNode(v)
}

func (n *BeginEndBlockNode) StatementListNode() *StatementListNode {
	var v unsafe.Pointer
	internalparser.ASTBeginEndBlock_statement_list_node(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStatementListNode(v)
}

func (n *BeginEndBlockNode) HandlerList() *ExceptionHandlerListNode {
	var v unsafe.Pointer
	internalparser.ASTBeginEndBlock_handler_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newExceptionHandlerListNode(v)
}

func (n *BeginEndBlockNode) StatementList() []StatementNode {
	var num int
	internalparser.ASTBeginEndBlock_statement_list_num(n.getRaw(), &num)
	ret := make([]StatementNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTBeginEndBlock_statement_list(n.getRaw(), i, &v)
		ret = append(ret, newNode(v).(StatementNode))
	}
	return ret
}

func (n *BeginEndBlockNode) HasExceptionHandler() bool {
	var v bool
	internalparser.ASTBeginEndBlock_has_exception_handler(n.getRaw(), &v)
	return v
}

type IdentifierListNode struct {
	*BaseNode
}

func (n *IdentifierListNode) IdentifierList() []*IdentifierNode {
	var num int
	internalparser.ASTIdentifierList_identifier_list_num(n.getRaw(), &num)
	ret := make([]*IdentifierNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTIdentifierList_identifier_list(n.getRaw(), i, &v)
		ret = append(ret, newIdentifierNode(v))
	}
	return ret
}

type VariableDeclarationNode struct {
	*ScriptStatementNode
}

func (n *VariableDeclarationNode) VariableList() *IdentifierListNode {
	var v unsafe.Pointer
	internalparser.ASTVariableDeclaration_variable_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierListNode(v)
}

func (n *VariableDeclarationNode) Type() TypeNode {
	var v unsafe.Pointer
	internalparser.ASTVariableDeclaration_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(TypeNode)
}

func (n *VariableDeclarationNode) DefaultValue() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTVariableDeclaration_default_value(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type UntilClauseNode struct {
	*BaseNode
}

func (n *UntilClauseNode) Condition() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTUntilClause_condition(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *UntilClauseNode) RepeatStmt() *RepeatStatementNode {
	var v unsafe.Pointer
	internalparser.ASTUntilClause_repeat_stmt(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newRepeatStatementNode(v)
}

type BreakContinueKeyword int

const (
	BreakKeyword BreakContinueKeyword = iota
	LeaveKeyword
	ContinueKeyword
	IterateKeyword
)

func (k BreakContinueKeyword) String() string {
	switch k {
	case BreakKeyword:
		return "BREAK"
	case LeaveKeyword:
		return "LEAVE"
	case ContinueKeyword:
		return "CONTINUE"
	case IterateKeyword:
		return "ITERATE"
	}
	return ""
}

type BreakContinueStatementNode struct {
	*ScriptStatementNode
}

func (n *BreakContinueStatementNode) Label() *LabelNode {
	var v unsafe.Pointer
	internalparser.ASTBreakContinueStatement_label(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newLabelNode(v)
}

func (n *BreakContinueStatementNode) SetKeyword(keyword BreakContinueKeyword) {
	internalparser.ASTBreakContinueStatement_set_keyword(n.getRaw(), int(keyword))
}

func (n *BreakContinueStatementNode) Keyword() BreakContinueKeyword {
	var v int
	internalparser.ASTBreakContinueStatement_keyword(n.getRaw(), &v)
	return BreakContinueKeyword(v)
}

type BreakStatementNode struct {
	*BreakContinueStatementNode
}

func (n *BreakStatementNode) SetKeyword(keyword BreakContinueKeyword) {
	internalparser.ASTBreakStatement_set_keyword(n.getRaw(), int(keyword))
}

func (n *BreakStatementNode) Keyword() BreakContinueKeyword {
	var v int
	internalparser.ASTBreakStatement_keyword(n.getRaw(), &v)
	return BreakContinueKeyword(v)
}

type ContinueStatementNode struct {
	*BreakContinueStatementNode
}

func (n *ContinueStatementNode) SetKeyword(keyword BreakContinueKeyword) {
	internalparser.ASTContinueStatement_set_keyword(n.getRaw(), int(keyword))
}

func (n *ContinueStatementNode) Keyword() BreakContinueKeyword {
	var v int
	internalparser.ASTContinueStatement_keyword(n.getRaw(), &v)
	return BreakContinueKeyword(v)
}

type DropPrivilegeRestrictionStatementNode struct {
	*DdlStatementBaseNode
}

func (n *DropPrivilegeRestrictionStatementNode) SetIsIfExists(isIfExists bool) {
	internalparser.ASTDropPrivilegeRestrictionStatement_set_is_if_exists(n.getRaw(), boolToInt(isIfExists))
}

func (n *DropPrivilegeRestrictionStatementNode) IsIfExists() bool {
	var v bool
	internalparser.ASTDropPrivilegeRestrictionStatement_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *DropPrivilegeRestrictionStatementNode) Privileges() *PrivilegesNode {
	var v unsafe.Pointer
	internalparser.ASTDropPrivilegeRestrictionStatement_privileges(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPrivilegesNode(v)
}

func (n *DropPrivilegeRestrictionStatementNode) ObjectType() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTDropPrivilegeRestrictionStatement_object_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *DropPrivilegeRestrictionStatementNode) NamePath() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDropPrivilegeRestrictionStatement_name_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type DropRowAccessPolicyStatementNode struct {
	*DdlStatementBaseNode
}

func (n *DropRowAccessPolicyStatementNode) SetIsIfExists(isIfExists bool) {
	internalparser.ASTDropRowAccessPolicyStatement_set_is_if_exists(n.getRaw(), boolToInt(isIfExists))
}

func (n *DropRowAccessPolicyStatementNode) IsIfExists() bool {
	var v bool
	internalparser.ASTDropRowAccessPolicyStatement_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *DropRowAccessPolicyStatementNode) TableName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDropRowAccessPolicyStatement_table_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *DropRowAccessPolicyStatementNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTDropRowAccessPolicyStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type CreatePrivilegeRestrictionStatementNode struct {
	*CreateStatementNode
}

func (n *CreatePrivilegeRestrictionStatementNode) Privileges() *PrivilegesNode {
	var v unsafe.Pointer
	internalparser.ASTCreatePrivilegeRestrictionStatement_privileges(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPrivilegesNode(v)
}

func (n *CreatePrivilegeRestrictionStatementNode) ObjectType() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTCreatePrivilegeRestrictionStatement_object_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *CreatePrivilegeRestrictionStatementNode) NamePath() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreatePrivilegeRestrictionStatement_name_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreatePrivilegeRestrictionStatementNode) RestrictTo() *RestrictToClauseNode {
	var v unsafe.Pointer
	internalparser.ASTCreatePrivilegeRestrictionStatement_restrict_to(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newRestrictToClauseNode(v)
}

type CreateRowAccessPolicyStatementNode struct {
	*CreateStatementNode
}

func (n *CreateRowAccessPolicyStatementNode) SetHasAccessKeyword(v bool) {
	internalparser.ASTCreateRowAccessPolicyStatement_set_has_access_keyword(n.getRaw(), boolToInt(v))
}

func (n *CreateRowAccessPolicyStatementNode) HasAccessKeyword() bool {
	var v bool
	internalparser.ASTCreateRowAccessPolicyStatement_has_access_keyword(n.getRaw(), &v)
	return v
}

func (n *CreateRowAccessPolicyStatementNode) TargetPath() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateRowAccessPolicyStatement_target_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreateRowAccessPolicyStatementNode) GrantTo() *GrantToClauseNode {
	var v unsafe.Pointer
	internalparser.ASTCreateRowAccessPolicyStatement_grant_to(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newGrantToClauseNode(v)
}

func (n *CreateRowAccessPolicyStatementNode) FilterUsing() *FilterUsingClauseNode {
	var v unsafe.Pointer
	internalparser.ASTCreateRowAccessPolicyStatement_filter_using(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newFilterUsingClauseNode(v)
}

func (n *CreateRowAccessPolicyStatementNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTCreateRowAccessPolicyStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type DropMode int

const (
	DropModeUnspecified DropMode = iota
	DropModeRestrict
	DropModeCascade
)

func (m DropMode) String() string {
	switch m {
	case DropModeUnspecified:
		return "UNSPECIFIED"
	case DropModeRestrict:
		return "RESTRICT"
	case DropModeCascade:
		return "CASCADE"
	}
	return ""
}

type DropStatementNode struct {
	*DdlStatementBaseNode
}

func (n *DropStatementNode) SetDropMode(mode DropMode) {
	internalparser.ASTDropStatement_set_drop_mode(n.getRaw(), int(mode))
}

func (n *DropStatementNode) DropMode() DropMode {
	var v int
	internalparser.ASTDropStatement_drop_mode(n.getRaw(), &v)
	return DropMode(v)
}

func (n *DropStatementNode) SetIsIfExists(isIfExists bool) {
	internalparser.ASTDropStatement_set_is_if_exists(n.getRaw(), boolToInt(isIfExists))
}

func (n *DropStatementNode) IsIfExists() bool {
	var v bool
	internalparser.ASTDropStatement_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *DropStatementNode) SetSchemaObjectKind(kind SchemaObjectKind) {
	internalparser.ASTDropStatement_set_schema_object_kind(n.getRaw(), int(kind))
}

func (n *DropStatementNode) SchemaObjectKind() SchemaObjectKind {
	var v int
	internalparser.ASTDropStatement_schema_object_kind(n.getRaw(), &v)
	return SchemaObjectKind(v)
}

func (n *DropStatementNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTDropStatemnt_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

type ReturnStatementNode struct {
	*ScriptStatementNode
}

type SingleAssignmentNode struct {
	*ScriptStatementNode
}

func (n *SingleAssignmentNode) Variable() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTSingleAssignment_variable(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *SingleAssignmentNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTSingleAssignment_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type ParameterAssignmentNode struct {
	*StatementBaseNode
}

func (n *ParameterAssignmentNode) Parameter() *ParameterExprNode {
	var v unsafe.Pointer
	internalparser.ASTParameterAssignment_parameter(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newParameterExprNode(v)
}

func (n *ParameterAssignmentNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTParameterAssignment_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type SystemVariableAssignmentNode struct {
	*StatementBaseNode
}

func (n *SystemVariableAssignmentNode) SystemVariable() *SystemVariableExprNode {
	var v unsafe.Pointer
	internalparser.ASTSystemVariableAssignment_system_variable(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newSystemVariableExprNode(v)
}

func (n *SystemVariableAssignmentNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTSystemVariableAssignment_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type AssignmentFromStructNode struct {
	*ScriptStatementNode
}

func (n *AssignmentFromStructNode) Variables() *IdentifierListNode {
	var v unsafe.Pointer
	internalparser.ASTAssignmentFromStruct_variables(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierListNode(v)
}

func (n *AssignmentFromStructNode) StructExpression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTAssignmentFromStruct_struct_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type CreateTableStmtBaseNode struct {
	*CreateStatementNode
}

func (n *CreateTableStmtBaseNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateTableStmtBase_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreateTableStmtBaseNode) TableElementList() *TableElementListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateTableStmtBase_table_element_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newTableElementListNode(v)
}

func (n *CreateTableStmtBaseNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateTableStmtBase_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

func (n *CreateTableStmtBaseNode) LikeTableName() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateTableStmtBase_like_table_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreateTableStmtBaseNode) Collate() *CollateNode {
	var v unsafe.Pointer
	internalparser.ASTCreateTableStmtBase_collate(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newCollateNode(v)
}

type CreateTableStatementNode struct {
	*CreateTableStmtBaseNode
}

func (n *CreateTableStatementNode) CloneDataSource() *CloneDataSourceNode {
	var v unsafe.Pointer
	internalparser.ASTCreateTableStatement_clone_data_source(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newCloneDataSourceNode(v)
}

func (n *CreateTableStatementNode) CopyDataSource() *CopyDataSourceNode {
	var v unsafe.Pointer
	internalparser.ASTCreateTableStatement_copy_data_source(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newCopyDataSourceNode(v)
}

func (n *CreateTableStatementNode) PartitionBy() *PartitionByNode {
	var v unsafe.Pointer
	internalparser.ASTCreateTableStatement_partition_by(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPartitionByNode(v)
}

func (n *CreateTableStatementNode) ClusterBy() *ClusterByNode {
	var v unsafe.Pointer
	internalparser.ASTCreateTableStatement_cluster_by(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newClusterByNode(v)
}

func (n *CreateTableStatementNode) Query() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTCreateTableStatement_query(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

type CreateExternalTableStatementNode struct {
	*CreateTableStmtBaseNode
}

func (n *CreateExternalTableStatementNode) WithPartitionColumnsClause() *WithPartitionColumnsClauseNode {
	var v unsafe.Pointer
	internalparser.ASTCreateExternalTableStatement_with_partition_columns_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWithPartitionColumnsClauseNode(v)
}

func (n *CreateExternalTableStatementNode) WithConnectionClause() *WithConnectionClauseNode {
	var v unsafe.Pointer
	internalparser.ASTCreateExternalTableStatement_with_connection_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWithConnectionClauseNode(v)
}

type CreateViewStatementBaseNode struct {
	*CreateStatementNode
}

func (n *CreateViewStatementBaseNode) Name() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTCreateViewStatementBase_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *CreateViewStatementBaseNode) ColumnList() *ColumnListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateViewStatementBase_column_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newColumnListNode(v)
}

func (n *CreateViewStatementBaseNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateViewStatementBase_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

func (n *CreateViewStatementBaseNode) Query() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTCreateViewStatementBase_query(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

type CreateViewStatementNode struct {
	*CreateViewStatementBaseNode
}

type CreateMaterializedViewStatementNode struct {
	*CreateViewStatementBaseNode
}

func (n *CreateMaterializedViewStatementNode) PartitionBy() *PartitionByNode {
	var v unsafe.Pointer
	internalparser.ASTCreateMaterializedViewStatement_partition_by(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPartitionByNode(v)
}

func (n *CreateMaterializedViewStatementNode) ClusterBy() *ClusterByNode {
	var v unsafe.Pointer
	internalparser.ASTCreateMaterializedViewStatement_cluster_by(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newClusterByNode(v)
}

type LoopStatementNode struct {
	*ScriptStatementNode
}

func (n *LoopStatementNode) Label() *LabelNode {
	var v unsafe.Pointer
	internalparser.ASTLoopStatement_label(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newLabelNode(v)
}

func (n *LoopStatementNode) Body() *StatementListNode {
	var v unsafe.Pointer
	internalparser.ASTLoopStatement_body(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStatementListNode(v)
}

func (n *LoopStatementNode) IsLoopStatement() bool {
	var v bool
	internalparser.ASTLoopStatement_IsLoopStatement(n.getRaw(), &v)
	return v
}

type WhileStatementNode struct {
	*LoopStatementNode
}

func (n *WhileStatementNode) Condition() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTWhileStatement_condition(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

type RepeatStatementNode struct {
	*LoopStatementNode
}

func (n *RepeatStatementNode) UntilClause() *UntilClauseNode {
	var v unsafe.Pointer
	internalparser.ASTRepeatStatement_until_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newUntilClauseNode(v)
}

type ForInStatementNode struct {
	*LoopStatementNode
}

func (n *ForInStatementNode) Variable() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTForInStatement_variable(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *ForInStatementNode) Query() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTForInStatement_query(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

type AlterStatementBaseNode struct {
	*DdlStatementBaseNode
}

func (n *AlterStatementBaseNode) SetIsIfExists(isIfExists bool) {
	internalparser.ASTAlterStatementBase_set_is_if_exists(n.getRaw(), boolToInt(isIfExists))
}

func (n *AlterStatementBaseNode) IsIfExists() bool {
	var v bool
	internalparser.ASTAlterStatementBase_is_if_exists(n.getRaw(), &v)
	return v
}

func (n *AlterStatementBaseNode) Path() *PathExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTAlterStatementBase_path(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPathExpressionNode(v)
}

func (n *AlterStatementBaseNode) ActionList() *AlterActionListNode {
	var v unsafe.Pointer
	internalparser.ASTAlterStatementBase_action_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAlterActionListNode(v)
}

type AlterDatabaseStatementNode struct {
	*AlterStatementBaseNode
}

type AlterSchemaStatementNode struct {
	*AlterStatementBaseNode
}

type AlterTableStatementNode struct {
	*AlterStatementBaseNode
}

type AlterViewStatementNode struct {
	*AlterStatementBaseNode
}

type AlterMaterializedViewStatementNode struct {
	*AlterStatementBaseNode
}

type AlterPrivilegeRestrictionStatementNode struct {
	*AlterStatementBaseNode
}

func (n *AlterPrivilegeRestrictionStatementNode) Privileges() *PrivilegesNode {
	var v unsafe.Pointer
	internalparser.ASTAlterPrivilegeRestrictionStatement_privileges(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPrivilegesNode(v)
}

func (n *AlterPrivilegeRestrictionStatementNode) ObjectType() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTAlterPrivilegeRestrictionStatement_object_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type AlterRowAccessPolicyStatementNode struct {
	*AlterStatementBaseNode
}

func (n *AlterRowAccessPolicyStatementNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTAlterRowAccessPolicyStatement_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type AlterEntityStatementNode struct {
	*AlterStatementBaseNode
}

func (n *AlterEntityStatementNode) Type() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTAlterEntityStatement_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

type DeterminismLevel int

const (
	DeterminismUnspecified DeterminismLevel = iota
	DeterministicLevel
	NotDeterministicLevel
	ImmutableLevel
	StableLevel
	VolatileLevel
)

func (l DeterminismLevel) String() string {
	switch l {
	case DeterminismUnspecified:
		return "DETERMINISM_UNSPECIFIED"
	case DeterministicLevel:
		return "DETERMINISTIC"
	case NotDeterministicLevel:
		return "NOT_DETERMINISTIC"
	case ImmutableLevel:
		return "IMMUTABLE"
	case StableLevel:
		return "STABLE"
	case VolatileLevel:
		return "VOLATILE"
	}
	return ""
}

type CreateFunctionStmtBaseNode struct {
	*CreateStatementNode
}

func (n *CreateFunctionStmtBaseNode) SetDeterminismLevel(level DeterminismLevel) {
	internalparser.ASTCreateFunctionStmtBase_set_determinism_level(n.getRaw(), int(level))
}

func (n *CreateFunctionStmtBaseNode) DeterminismLevel() DeterminismLevel {
	var v int
	internalparser.ASTCreateFunctionStmtBase_determinism_level(n.getRaw(), &v)
	return DeterminismLevel(v)
}

func (n *CreateFunctionStmtBaseNode) SetSqlSecurity(security SqlSecurity) {
	internalparser.ASTCreateFunctionStmtBase_set_sql_security(n.getRaw(), int(security))
}

func (n *CreateFunctionStmtBaseNode) SqlSecurity() SqlSecurity {
	var v int
	internalparser.ASTCreateFunctionStmtBase_sql_security(n.getRaw(), &v)
	return SqlSecurity(v)
}

func (n *CreateFunctionStmtBaseNode) FunctionDeclaration() *FunctionDeclarationNode {
	var v unsafe.Pointer
	internalparser.ASTCreateFunctionStmtBase_function_declaration(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newFunctionDeclarationNode(v)
}

func (n *CreateFunctionStmtBaseNode) Language() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTCreateFunctionStmtBase_language(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func (n *CreateFunctionStmtBaseNode) Code() *StringLiteralNode {
	var v unsafe.Pointer
	internalparser.ASTCreateFunctionStmtBase_code(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newStringLiteralNode(v)
}

func (n *CreateFunctionStmtBaseNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTCreateFunctionStmtBase_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type CreateFunctionStatementNode struct {
	*CreateFunctionStmtBaseNode
}

func (n *CreateFunctionStatementNode) SetIsAggregate(isAggregate bool) {
	internalparser.ASTCreateFunctionStatement_set_is_aggregate(n.getRaw(), boolToInt(isAggregate))
}

func (n *CreateFunctionStatementNode) IsAggregate() bool {
	var v bool
	internalparser.ASTCreateFunctionStatement_is_aggregate(n.getRaw(), &v)
	return v
}

func (n *CreateFunctionStatementNode) SetIsRemote(isRemote bool) {
	internalparser.ASTCreateFunctionStatement_set_is_remote(n.getRaw(), boolToInt(isRemote))
}

func (n *CreateFunctionStatementNode) IsRemote() bool {
	var v bool
	internalparser.ASTCreateFunctionStatement_is_remote(n.getRaw(), &v)
	return v
}

func (n *CreateFunctionStatementNode) ReturnType() TypeNode {
	var v unsafe.Pointer
	internalparser.ASTCreateFunctionStatement_return_type(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(TypeNode)
}

func (n *CreateFunctionStatementNode) SqlFunctionBody() *SqlFunctionBodyNode {
	var v unsafe.Pointer
	internalparser.ASTCreateFunctionStatement_sql_function_body(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newSqlFunctionBodyNode(v)
}

func (n *CreateFunctionStatementNode) WithConnectionClause() *WithConnectionClauseNode {
	var v unsafe.Pointer
	internalparser.ASTCreateFunctionStatement_with_connection_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWithConnectionClauseNode(v)
}

type CreateTableFunctionStatementNode struct {
	*CreateFunctionStmtBaseNode
}

func (n *CreateTableFunctionStatementNode) ReturnTVFSchema() *TVFSchemaNode {
	var v unsafe.Pointer
	internalparser.ASTCreateTableFunctionStatement_return_tvf_schema(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newTVFSchemaNode(v)
}

func (n *CreateTableFunctionStatementNode) Query() *QueryNode {
	var v unsafe.Pointer
	internalparser.ASTCreateTableFunctionStatement_query(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newQueryNode(v)
}

type StructColumnSchemaNode struct {
	*ColumnSchemaNode
}

func (n *StructColumnSchemaNode) StructFields() []*StructColumnFieldNode {
	var num int
	internalparser.ASTStructColumnSchema_struct_fields_num(n.getRaw(), &num)
	ret := make([]*StructColumnFieldNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTStructColumnSchema_struct_field(n.getRaw(), i, &v)
		ret = append(ret, newStructColumnFieldNode(v))
	}
	return ret
}

type InferredTypeColumnSchemaNode struct {
	*ColumnSchemaNode
}

type ExecuteIntoClauseNode struct {
	*BaseNode
}

func (n *ExecuteIntoClauseNode) Identifiers() *IdentifierListNode {
	var v unsafe.Pointer
	internalparser.ASTExecuteIntoClause_identifiers(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierListNode(v)
}

type ExecuteUsingArgumentNode struct {
	*BaseNode
}

func (n *ExecuteUsingArgumentNode) Expression() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTExecuteUsingArgument_expression(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *ExecuteUsingArgumentNode) Alias() *AliasNode {
	var v unsafe.Pointer
	internalparser.ASTExecuteUsingArgument_alias(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAliasNode(v)
}

type ExecuteUsingClauseNode struct {
	*BaseNode
}

func (n *ExecuteUsingClauseNode) Arguments() []*ExecuteUsingArgumentNode {
	var num int
	internalparser.ASTExecuteUsingClause_arguments_num(n.getRaw(), &num)
	ret := make([]*ExecuteUsingArgumentNode, 0, num)
	for i := 0; i < num; i++ {
		var v unsafe.Pointer
		internalparser.ASTExecuteUsingClause_argument(n.getRaw(), i, &v)
		ret = append(ret, newExecuteUsingArgumentNode(v))
	}
	return ret
}

type ExecuteImmediateStatementNode struct {
	*StatementBaseNode
}

func (n *ExecuteImmediateStatementNode) SQL() ExpressionNode {
	var v unsafe.Pointer
	internalparser.ASTExecuteImmediateStatement_sql(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newNode(v).(ExpressionNode)
}

func (n *ExecuteImmediateStatementNode) IntoClause() *ExecuteIntoClauseNode {
	var v unsafe.Pointer
	internalparser.ASTExecuteImmediateStatement_into_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newExecuteIntoClauseNode(v)
}

func (n *ExecuteImmediateStatementNode) UsingClause() *ExecuteUsingClauseNode {
	var v unsafe.Pointer
	internalparser.ASTExecuteImmediateStatement_using_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newExecuteUsingClauseNode(v)
}

type AuxLoadDataFromFilesOptionsListNode struct {
	*BaseNode
}

func (n *AuxLoadDataFromFilesOptionsListNode) OptionsList() *OptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTAuxLoadDataFromFilesOptionsList_options_list(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newOptionsListNode(v)
}

type InsertionMode int

const (
	InsertionModeNotSet InsertionMode = iota
	InsertionModeAppend
	InsertionModeOverwrite
)

func (m InsertionMode) String() string {
	switch m {
	case InsertionModeNotSet:
		return "NOT_SET"
	case InsertionModeAppend:
		return "APPEND"
	case InsertionModeOverwrite:
		return "OVERWRITE"
	}
	return ""
}

type AuxLoadDataStatementNode struct {
	*CreateTableStmtBaseNode
}

func (n *AuxLoadDataStatementNode) SetInsertionMode(mode InsertionMode) {
	internalparser.ASTAuxLoadDataStatement_set_insertion_mode(n.getRaw(), int(mode))
}

func (n *AuxLoadDataStatementNode) InsertionMode() InsertionMode {
	var v int
	internalparser.ASTAuxLoadDataStatement_insertion_mode(n.getRaw(), &v)
	return InsertionMode(v)
}

func (n *AuxLoadDataStatementNode) PartitionBy() *PartitionByNode {
	var v unsafe.Pointer
	internalparser.ASTAuxLoadDataStatement_partition_by(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newPartitionByNode(v)
}

func (n *AuxLoadDataStatementNode) ClusterBy() *ClusterByNode {
	var v unsafe.Pointer
	internalparser.ASTAuxLoadDataStatement_cluster_by(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newClusterByNode(v)
}

func (n *AuxLoadDataStatementNode) FromFiles() *AuxLoadDataFromFilesOptionsListNode {
	var v unsafe.Pointer
	internalparser.ASTAuxLoadDataStatement_from_files(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newAuxLoadDataFromFilesOptionsListNode(v)
}

func (n *AuxLoadDataStatementNode) WithPartitionColumnsClause() *WithPartitionColumnsClauseNode {
	var v unsafe.Pointer
	internalparser.ASTAuxLoadDataStatement_with_partition_columns_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWithPartitionColumnsClauseNode(v)
}

func (n *AuxLoadDataStatementNode) WithConnectionClause() *WithConnectionClauseNode {
	var v unsafe.Pointer
	internalparser.ASTAuxLoadDataStatement_with_connection_clause(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newWithConnectionClauseNode(v)
}

type LabelNode struct {
	*BaseNode
}

func (n *LabelNode) Name() *IdentifierNode {
	var v unsafe.Pointer
	internalparser.ASTLabel_name(n.getRaw(), &v)
	if v == nil {
		return nil
	}
	return newIdentifierNode(v)
}

func newBaseNode(raw unsafe.Pointer) *BaseNode {
	return &BaseNode{raw: raw}
}

func newStatementBaseNode(n unsafe.Pointer) *StatementBaseNode {
	return &StatementBaseNode{BaseNode: newBaseNode(n)}
}

func newScriptBaseNode(n unsafe.Pointer) *ScriptBaseNode {
	return &ScriptBaseNode{BaseNode: newBaseNode(n)}
}

func newTypeBaseNode(n unsafe.Pointer) *TypeBaseNode {
	return &TypeBaseNode{BaseNode: newBaseNode(n)}
}

func newExpressionBaseNode(n unsafe.Pointer) *ExpressionBaseNode {
	return &ExpressionBaseNode{BaseNode: newBaseNode(n)}
}

func newQueryStatementNode(n unsafe.Pointer) *QueryStatementNode {
	return &QueryStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newQueryExpressionBaseNode(n unsafe.Pointer) *QueryExpressionBaseNode {
	return &QueryExpressionBaseNode{BaseNode: newBaseNode(n)}
}

func newQueryNode(n unsafe.Pointer) *QueryNode {
	return &QueryNode{QueryExpressionBaseNode: newQueryExpressionBaseNode(n)}
}

func newSelectNode(n unsafe.Pointer) *SelectNode {
	return &SelectNode{QueryExpressionBaseNode: newQueryExpressionBaseNode(n)}
}

func newSelectListNode(n unsafe.Pointer) *SelectListNode {
	return &SelectListNode{BaseNode: newBaseNode(n)}
}

func newSelectColumnNode(n unsafe.Pointer) *SelectColumnNode {
	return &SelectColumnNode{BaseNode: newBaseNode(n)}
}

func newLeafBaseNode(n unsafe.Pointer) *LeafBaseNode {
	return &LeafBaseNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newIntLiteralNode(n unsafe.Pointer) *IntLiteralNode {
	return &IntLiteralNode{LeafBaseNode: newLeafBaseNode(n)}
}

func newIdentifierNode(n unsafe.Pointer) *IdentifierNode {
	return &IdentifierNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newAliasNode(n unsafe.Pointer) *AliasNode {
	return &AliasNode{BaseNode: newBaseNode(n)}
}

func newGeneralizedPathExpressionBaseNode(n unsafe.Pointer) *GeneralizedPathExpressionBaseNode {
	return &GeneralizedPathExpressionBaseNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newPathExpressionNode(n unsafe.Pointer) *PathExpressionNode {
	return &PathExpressionNode{GeneralizedPathExpressionBaseNode: newGeneralizedPathExpressionBaseNode(n)}
}

func newTableExpressionBaseNode(n unsafe.Pointer) *TableExpressionBaseNode {
	return &TableExpressionBaseNode{BaseNode: newBaseNode(n)}
}

func newTablePathExpressionNode(n unsafe.Pointer) *TablePathExpressionNode {
	return &TablePathExpressionNode{TableExpressionBaseNode: newTableExpressionBaseNode(n)}
}

func newFromClauseNode(n unsafe.Pointer) *FromClauseNode {
	return &FromClauseNode{BaseNode: newBaseNode(n)}
}

func newWhereClauseNode(n unsafe.Pointer) *WhereClauseNode {
	return &WhereClauseNode{BaseNode: newBaseNode(n)}
}

func newBooleanLiteralNode(n unsafe.Pointer) *BooleanLiteralNode {
	return &BooleanLiteralNode{LeafBaseNode: newLeafBaseNode(n)}
}

func newAndExprNode(n unsafe.Pointer) *AndExprNode {
	return &AndExprNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newBinaryExpressionNode(n unsafe.Pointer) *BinaryExpressionNode {
	return &BinaryExpressionNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newStringLiteralNode(n unsafe.Pointer) *StringLiteralNode {
	return &StringLiteralNode{LeafBaseNode: newLeafBaseNode(n)}
}

func newStarNode(n unsafe.Pointer) *StarNode {
	return &StarNode{LeafBaseNode: newLeafBaseNode(n)}
}

func newOrExprNode(n unsafe.Pointer) *OrExprNode {
	return &OrExprNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newGroupingItemNode(n unsafe.Pointer) *GroupingItemNode {
	return &GroupingItemNode{BaseNode: newBaseNode(n)}
}

func newGroupByNode(n unsafe.Pointer) *GroupByNode {
	return &GroupByNode{BaseNode: newBaseNode(n)}
}

func newOrderingExpressionNode(n unsafe.Pointer) *OrderingExpressionNode {
	return &OrderingExpressionNode{BaseNode: newBaseNode(n)}
}

func newOrderByNode(n unsafe.Pointer) *OrderByNode {
	return &OrderByNode{BaseNode: newBaseNode(n)}
}

func newLimitOffsetNode(n unsafe.Pointer) *LimitOffsetNode {
	return &LimitOffsetNode{BaseNode: newBaseNode(n)}
}

func newFloatLiteralNode(n unsafe.Pointer) *FloatLiteralNode {
	return &FloatLiteralNode{LeafBaseNode: newLeafBaseNode(n)}
}

func newNullLiteralNode(n unsafe.Pointer) *NullLiteralNode {
	return &NullLiteralNode{LeafBaseNode: newLeafBaseNode(n)}
}

func newOnClauseNode(n unsafe.Pointer) *OnClauseNode {
	return &OnClauseNode{BaseNode: newBaseNode(n)}
}

func newWithClauseEntryNode(n unsafe.Pointer) *WithClauseEntryNode {
	return &WithClauseEntryNode{BaseNode: newBaseNode(n)}
}

func newJoinNode(n unsafe.Pointer) *JoinNode {
	return &JoinNode{TableExpressionBaseNode: newTableExpressionBaseNode(n)}
}

func newWithClauseNode(n unsafe.Pointer) *WithClauseNode {
	return &WithClauseNode{BaseNode: newBaseNode(n)}
}

func newHavingNode(n unsafe.Pointer) *HavingNode {
	return &HavingNode{BaseNode: newBaseNode(n)}
}

func newSimpleTypeNode(n unsafe.Pointer) *SimpleTypeNode {
	return &SimpleTypeNode{TypeBaseNode: newTypeBaseNode(n)}
}

func newArrayTypeNode(n unsafe.Pointer) *ArrayTypeNode {
	return &ArrayTypeNode{TypeBaseNode: newTypeBaseNode(n)}
}

func newStructFieldNode(n unsafe.Pointer) *StructFieldNode {
	return &StructFieldNode{BaseNode: newBaseNode(n)}
}

func newStructTypeNode(n unsafe.Pointer) *StructTypeNode {
	return &StructTypeNode{TypeBaseNode: newTypeBaseNode(n)}
}

func newCastExpressionNode(n unsafe.Pointer) *CastExpressionNode {
	return &CastExpressionNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newSelectAsNode(n unsafe.Pointer) *SelectAsNode {
	return &SelectAsNode{BaseNode: newBaseNode(n)}
}

func newRollupNode(n unsafe.Pointer) *RollupNode {
	return &RollupNode{BaseNode: newBaseNode(n)}
}

func newFunctionCallNode(n unsafe.Pointer) *FunctionCallNode {
	return &FunctionCallNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newArrayConstructorNode(n unsafe.Pointer) *ArrayConstructorNode {
	return &ArrayConstructorNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newStructConstructorArgNode(n unsafe.Pointer) *StructConstructorArgNode {
	return &StructConstructorArgNode{BaseNode: newBaseNode(n)}
}

func newStructConstructorWithParensNode(n unsafe.Pointer) *StructConstructorWithParensNode {
	return &StructConstructorWithParensNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newStructConstructorWithKeywordNode(n unsafe.Pointer) *StructConstructorWithKeywordNode {
	return &StructConstructorWithKeywordNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newInExpressionNode(n unsafe.Pointer) *InExpressionNode {
	return &InExpressionNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newInListNode(n unsafe.Pointer) *InListNode {
	return &InListNode{BaseNode: newBaseNode(n)}
}

func newBetweenExpressionNode(n unsafe.Pointer) *BetweenExpressionNode {
	return &BetweenExpressionNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newNumericLiteralNode(n unsafe.Pointer) *NumericLiteralNode {
	return &NumericLiteralNode{LeafBaseNode: newLeafBaseNode(n)}
}

func newBigNumericLiteralNode(n unsafe.Pointer) *BigNumericLiteralNode {
	return &BigNumericLiteralNode{LeafBaseNode: newLeafBaseNode(n)}
}

func newBytesLiteralNode(n unsafe.Pointer) *BytesLiteralNode {
	return &BytesLiteralNode{LeafBaseNode: newLeafBaseNode(n)}
}

func newDateOrTimeLiteralNode(n unsafe.Pointer) *DateOrTimeLiteralNode {
	return &DateOrTimeLiteralNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newMaxLiteralNode(n unsafe.Pointer) *MaxLiteralNode {
	return &MaxLiteralNode{LeafBaseNode: newLeafBaseNode(n)}
}

func newJSONLiteralNode(n unsafe.Pointer) *JSONLiteralNode {
	return &JSONLiteralNode{LeafBaseNode: newLeafBaseNode(n)}
}

func newCaseValueExpressionNode(n unsafe.Pointer) *CaseValueExpressionNode {
	return &CaseValueExpressionNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newCaseNoValueExpressionNode(n unsafe.Pointer) *CaseNoValueExpressionNode {
	return &CaseNoValueExpressionNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newArrayElementNode(n unsafe.Pointer) *ArrayElementNode {
	return &ArrayElementNode{GeneralizedPathExpressionBaseNode: newGeneralizedPathExpressionBaseNode(n)}
}

func newBitwiseShiftExpressionNode(n unsafe.Pointer) *BitwiseShiftExpressionNode {
	return &BitwiseShiftExpressionNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newCollateNode(n unsafe.Pointer) *CollateNode {
	return &CollateNode{BaseNode: newBaseNode(n)}
}

func newDotGeneralizedFieldNode(n unsafe.Pointer) *DotGeneralizedFieldNode {
	return &DotGeneralizedFieldNode{GeneralizedPathExpressionBaseNode: newGeneralizedPathExpressionBaseNode(n)}
}

func newDotIdentifierNode(n unsafe.Pointer) *DotIdentifierNode {
	return &DotIdentifierNode{GeneralizedPathExpressionBaseNode: newGeneralizedPathExpressionBaseNode(n)}
}

func newDotStarNode(n unsafe.Pointer) *DotStarNode {
	return &DotStarNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newDotStarWithModifiersNode(n unsafe.Pointer) *DotStarWithModifiersNode {
	return &DotStarWithModifiersNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newExpressionSubqueryNode(n unsafe.Pointer) *ExpressionSubqueryNode {
	return &ExpressionSubqueryNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newExtractExpressionNode(n unsafe.Pointer) *ExtractExpressionNode {
	return &ExtractExpressionNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newHavingModifierNode(n unsafe.Pointer) *HavingModifierNode {
	return &HavingModifierNode{BaseNode: newBaseNode(n)}
}

func newIntervalExprNode(n unsafe.Pointer) *IntervalExprNode {
	return &IntervalExprNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newNamedArgumentNode(n unsafe.Pointer) *NamedArgumentNode {
	return &NamedArgumentNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newNullOrderNode(n unsafe.Pointer) *NullOrderNode {
	return &NullOrderNode{BaseNode: newBaseNode(n)}
}

func newOnOrUsingClauseListNode(n unsafe.Pointer) *OnOrUsingClauseListNode {
	return &OnOrUsingClauseListNode{BaseNode: newBaseNode(n)}
}

func newParenthesizedJoinNode(n unsafe.Pointer) *ParenthesizedJoinNode {
	return &ParenthesizedJoinNode{TableExpressionBaseNode: newTableExpressionBaseNode(n)}
}

func newPartitionByNode(n unsafe.Pointer) *PartitionByNode {
	return &PartitionByNode{BaseNode: newBaseNode(n)}
}

func newSetOperationNode(n unsafe.Pointer) *SetOperationNode {
	return &SetOperationNode{QueryExpressionBaseNode: newQueryExpressionBaseNode(n)}
}

func newStarExceptListNode(n unsafe.Pointer) *StarExceptListNode {
	return &StarExceptListNode{BaseNode: newBaseNode(n)}
}

func newStarModifiersNode(n unsafe.Pointer) *StarModifiersNode {
	return &StarModifiersNode{BaseNode: newBaseNode(n)}
}

func newStarReplaceItemNode(n unsafe.Pointer) *StarReplaceItemNode {
	return &StarReplaceItemNode{BaseNode: newBaseNode(n)}
}

func newStarWithModifiersNode(n unsafe.Pointer) *StarWithModifiersNode {
	return &StarWithModifiersNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newTableSubqueryNode(n unsafe.Pointer) *TableSubqueryNode {
	return &TableSubqueryNode{TableExpressionBaseNode: newTableExpressionBaseNode(n)}
}

func newUnaryExpressionNode(n unsafe.Pointer) *UnaryExpressionNode {
	return &UnaryExpressionNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newUnnestExpressionNode(n unsafe.Pointer) *UnnestExpressionNode {
	return &UnnestExpressionNode{BaseNode: newBaseNode(n)}
}

func newWindowClauseNode(n unsafe.Pointer) *WindowClauseNode {
	return &WindowClauseNode{BaseNode: newBaseNode(n)}
}

func newWindowDefinitionNode(n unsafe.Pointer) *WindowDefinitionNode {
	return &WindowDefinitionNode{BaseNode: newBaseNode(n)}
}

func newWindowFrameNode(n unsafe.Pointer) *WindowFrameNode {
	return &WindowFrameNode{BaseNode: newBaseNode(n)}
}

func newWindowFrameExprNode(n unsafe.Pointer) *WindowFrameExprNode {
	return &WindowFrameExprNode{BaseNode: newBaseNode(n)}
}

func newLikeExpressionNode(n unsafe.Pointer) *LikeExpressionNode {
	return &LikeExpressionNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newWindowSpecificationNode(n unsafe.Pointer) *WindowSpecificationNode {
	return &WindowSpecificationNode{BaseNode: newBaseNode(n)}
}

func newWithOffsetNode(n unsafe.Pointer) *WithOffsetNode {
	return &WithOffsetNode{BaseNode: newBaseNode(n)}
}

func newAnySomeAllOpNode(n unsafe.Pointer) *AnySomeAllOpNode {
	return &AnySomeAllOpNode{BaseNode: newBaseNode(n)}
}

func newParameterExprBaseNode(n unsafe.Pointer) *ParameterExprBaseNode {
	return &ParameterExprBaseNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newStatementListNode(n unsafe.Pointer) *StatementListNode {
	return &StatementListNode{BaseNode: newBaseNode(n)}
}

func newScriptStatementNode(n unsafe.Pointer) *ScriptStatementNode {
	return &ScriptStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newHintedStatementNode(n unsafe.Pointer) *HintedStatementNode {
	return &HintedStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newExplainStatementNode(n unsafe.Pointer) *ExplainStatementNode {
	return &ExplainStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newDescribeStatementNode(n unsafe.Pointer) *DescribeStatementNode {
	return &DescribeStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newShowStatementNode(n unsafe.Pointer) *ShowStatementNode {
	return &ShowStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newTransactionModeBaseNode(n unsafe.Pointer) *TransactionModeBaseNode {
	return &TransactionModeBaseNode{BaseNode: newBaseNode(n)}
}

func newTransactionIsolationLevelNode(n unsafe.Pointer) *TransactionIsolationLevelNode {
	return &TransactionIsolationLevelNode{TransactionModeBaseNode: newTransactionModeBaseNode(n)}
}

func newTransactionReadWriteModeNode(n unsafe.Pointer) *TransactionReadWriteModeNode {
	return &TransactionReadWriteModeNode{TransactionModeBaseNode: newTransactionModeBaseNode(n)}
}

func newTransactionModeListNode(n unsafe.Pointer) *TransactionModeListNode {
	return &TransactionModeListNode{BaseNode: newBaseNode(n)}
}

func newBeginStatementNode(n unsafe.Pointer) *BeginStatementNode {
	return &BeginStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newSetTransactionStatementNode(n unsafe.Pointer) *SetTransactionStatementNode {
	return &SetTransactionStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newCommitStatementNode(n unsafe.Pointer) *CommitStatementNode {
	return &CommitStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newRollbackStatementNode(n unsafe.Pointer) *RollbackStatementNode {
	return &RollbackStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newStartBatchStatementNode(n unsafe.Pointer) *StartBatchStatementNode {
	return &StartBatchStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newRunBatchStatementNode(n unsafe.Pointer) *RunBatchStatementNode {
	return &RunBatchStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newAbortBatchStatementNode(n unsafe.Pointer) *AbortBatchStatementNode {
	return &AbortBatchStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newDdlStatementBaseNode(n unsafe.Pointer) *DdlStatementBaseNode {
	return &DdlStatementBaseNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newDropEntityStatementNode(n unsafe.Pointer) *DropEntityStatementNode {
	return &DropEntityStatementNode{DdlStatementBaseNode: newDdlStatementBaseNode(n)}
}

func newDropFunctionStatementNode(n unsafe.Pointer) *DropFunctionStatementNode {
	return &DropFunctionStatementNode{DdlStatementBaseNode: newDdlStatementBaseNode(n)}
}

func newDropTableFunctionStatementNode(n unsafe.Pointer) *DropTableFunctionStatementNode {
	return &DropTableFunctionStatementNode{DdlStatementBaseNode: newDdlStatementBaseNode(n)}
}

func newDropAllRowAccessPoliciesStatementNode(n unsafe.Pointer) *DropAllRowAccessPoliciesStatementNode {
	return &DropAllRowAccessPoliciesStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newDropMaterializedViewStatementNode(n unsafe.Pointer) *DropMaterializedViewStatementNode {
	return &DropMaterializedViewStatementNode{DdlStatementBaseNode: newDdlStatementBaseNode(n)}
}

func newDropSnapshotTableStatementNode(n unsafe.Pointer) *DropSnapshotTableStatementNode {
	return &DropSnapshotTableStatementNode{DdlStatementBaseNode: newDdlStatementBaseNode(n)}
}

func newDropSearchIndexStatementNode(n unsafe.Pointer) *DropSearchIndexStatementNode {
	return &DropSearchIndexStatementNode{DdlStatementBaseNode: newDdlStatementBaseNode(n)}
}

func newRenameStatementNode(n unsafe.Pointer) *RenameStatementNode {
	return &RenameStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newImportStatementNode(n unsafe.Pointer) *ImportStatementNode {
	return &ImportStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newModuleStatementNode(n unsafe.Pointer) *ModuleStatementNode {
	return &ModuleStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newWithConnectionClauseNode(n unsafe.Pointer) *WithConnectionClauseNode {
	return &WithConnectionClauseNode{BaseNode: newBaseNode(n)}
}

func newIntoAliasNode(n unsafe.Pointer) *IntoAliasNode {
	return &IntoAliasNode{BaseNode: newBaseNode(n)}
}

func newUnnestExpressionWithOptAliasAndOffsetNode(n unsafe.Pointer) *UnnestExpressionWithOptAliasAndOffsetNode {
	return &UnnestExpressionWithOptAliasAndOffsetNode{BaseNode: newBaseNode(n)}
}

func newPivotExpressionNode(n unsafe.Pointer) *PivotExpressionNode {
	return &PivotExpressionNode{BaseNode: newBaseNode(n)}
}

func newPivotValueNode(n unsafe.Pointer) *PivotValueNode {
	return &PivotValueNode{BaseNode: newBaseNode(n)}
}

func newPivotExpressionListNode(n unsafe.Pointer) *PivotExpressionListNode {
	return &PivotExpressionListNode{BaseNode: newBaseNode(n)}
}

func newPivotValueListNode(n unsafe.Pointer) *PivotValueListNode {
	return &PivotValueListNode{BaseNode: newBaseNode(n)}
}

func newPivotClauseNode(n unsafe.Pointer) *PivotClauseNode {
	return &PivotClauseNode{BaseNode: newBaseNode(n)}
}

func newUnpivotInItemNode(n unsafe.Pointer) *UnpivotInItemNode {
	return &UnpivotInItemNode{BaseNode: newBaseNode(n)}
}

func newUnpivotInItemListNode(n unsafe.Pointer) *UnpivotInItemListNode {
	return &UnpivotInItemListNode{BaseNode: newBaseNode(n)}
}

func newUnpivotClauseNode(n unsafe.Pointer) *UnpivotClauseNode {
	return &UnpivotClauseNode{BaseNode: newBaseNode(n)}
}

func newUsingClauseNode(n unsafe.Pointer) *UsingClauseNode {
	return &UsingClauseNode{BaseNode: newBaseNode(n)}
}

func newForSystemTimeNode(n unsafe.Pointer) *ForSystemTimeNode {
	return &ForSystemTimeNode{BaseNode: newBaseNode(n)}
}

func newQualifyNode(n unsafe.Pointer) *QualifyNode {
	return &QualifyNode{BaseNode: newBaseNode(n)}
}

func newClampedBetweenModifierNode(n unsafe.Pointer) *ClampedBetweenModifierNode {
	return &ClampedBetweenModifierNode{BaseNode: newBaseNode(n)}
}

func newFormatClauseNode(n unsafe.Pointer) *FormatClauseNode {
	return &FormatClauseNode{BaseNode: newBaseNode(n)}
}

func newPathExpressionListNode(n unsafe.Pointer) *PathExpressionListNode {
	return &PathExpressionListNode{BaseNode: newBaseNode(n)}
}

func newParameterExprNode(n unsafe.Pointer) *ParameterExprNode {
	return &ParameterExprNode{ParameterExprBaseNode: newParameterExprBaseNode(n)}
}

func newSystemVariableExprNode(n unsafe.Pointer) *SystemVariableExprNode {
	return &SystemVariableExprNode{ParameterExprBaseNode: newParameterExprBaseNode(n)}
}

func newWithGroupRowsNode(n unsafe.Pointer) *WithGroupRowsNode {
	return &WithGroupRowsNode{BaseNode: newBaseNode(n)}
}

func newLambdaNode(n unsafe.Pointer) *LambdaNode {
	return &LambdaNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newAnalyticFunctionCallNode(n unsafe.Pointer) *AnalyticFunctionCallNode {
	return &AnalyticFunctionCallNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newFunctionCallWithGroupRowsNode(n unsafe.Pointer) *FunctionCallWithGroupRowsNode {
	return &FunctionCallWithGroupRowsNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newClusterByNode(n unsafe.Pointer) *ClusterByNode {
	return &ClusterByNode{BaseNode: newBaseNode(n)}
}

func newNewConstructorArgNode(n unsafe.Pointer) *NewConstructorArgNode {
	return &NewConstructorArgNode{BaseNode: newBaseNode(n)}
}

func newNewConstructorNode(n unsafe.Pointer) *NewConstructorNode {
	return &NewConstructorNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newOptionsListNode(n unsafe.Pointer) *OptionsListNode {
	return &OptionsListNode{BaseNode: newBaseNode(n)}
}

func newOptionsEntryNode(n unsafe.Pointer) *OptionsEntryNode {
	return &OptionsEntryNode{BaseNode: newBaseNode(n)}
}

func newCreateStatementNode(n unsafe.Pointer) *CreateStatementNode {
	return &CreateStatementNode{DdlStatementBaseNode: newDdlStatementBaseNode(n)}
}

func newFunctionParameterNode(n unsafe.Pointer) *FunctionParameterNode {
	return &FunctionParameterNode{BaseNode: newBaseNode(n)}
}

func newFunctionParametersNode(n unsafe.Pointer) *FunctionParametersNode {
	return &FunctionParametersNode{BaseNode: newBaseNode(n)}
}

func newFunctionDeclarationNode(n unsafe.Pointer) *FunctionDeclarationNode {
	return &FunctionDeclarationNode{BaseNode: newBaseNode(n)}
}

func newSqlFunctionBodyNode(n unsafe.Pointer) *SqlFunctionBodyNode {
	return &SqlFunctionBodyNode{BaseNode: newBaseNode(n)}
}

func newTVFArgumentNode(n unsafe.Pointer) *TVFArgumentNode {
	return &TVFArgumentNode{BaseNode: newBaseNode(n)}
}

func newTVFNode(n unsafe.Pointer) *TVFNode {
	return &TVFNode{TableExpressionBaseNode: newTableExpressionBaseNode(n)}
}

func newTableClauseNode(n unsafe.Pointer) *TableClauseNode {
	return &TableClauseNode{BaseNode: newBaseNode(n)}
}

func newModelClauseNode(n unsafe.Pointer) *ModelClauseNode {
	return &ModelClauseNode{BaseNode: newBaseNode(n)}
}

func newConnectionClauseNode(n unsafe.Pointer) *ConnectionClauseNode {
	return &ConnectionClauseNode{BaseNode: newBaseNode(n)}
}

func newTableDataSourceNode(n unsafe.Pointer) *TableDataSourceNode {
	return &TableDataSourceNode{TableExpressionBaseNode: newTableExpressionBaseNode(n)}
}

func newCloneDataSourceNode(n unsafe.Pointer) *CloneDataSourceNode {
	return &CloneDataSourceNode{TableDataSourceNode: newTableDataSourceNode(n)}
}

func newCopyDataSourceNode(n unsafe.Pointer) *CopyDataSourceNode {
	return &CopyDataSourceNode{TableDataSourceNode: newTableDataSourceNode(n)}
}

func newCloneDataSourceListNode(n unsafe.Pointer) *CloneDataSourceListNode {
	return &CloneDataSourceListNode{BaseNode: newBaseNode(n)}
}

func newCloneDataStatementNode(n unsafe.Pointer) *CloneDataStatementNode {
	return &CloneDataStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newCreateConstantStatementNode(n unsafe.Pointer) *CreateConstantStatementNode {
	return &CreateConstantStatementNode{CreateStatementNode: newCreateStatementNode(n)}
}

func newCreateDatabaseStatementNode(n unsafe.Pointer) *CreateDatabaseStatementNode {
	return &CreateDatabaseStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newCreateProcedureStatementNode(n unsafe.Pointer) *CreateProcedureStatementNode {
	return &CreateProcedureStatementNode{CreateStatementNode: newCreateStatementNode(n)}
}

func newCreateSchemaStatementNode(n unsafe.Pointer) *CreateSchemaStatementNode {
	return &CreateSchemaStatementNode{CreateStatementNode: newCreateStatementNode(n)}
}

func newTransformClauseNode(n unsafe.Pointer) *TransformClauseNode {
	return &TransformClauseNode{BaseNode: newBaseNode(n)}
}

func newCreateModelStatementNode(n unsafe.Pointer) *CreateModelStatementNode {
	return &CreateModelStatementNode{CreateStatementNode: newCreateStatementNode(n)}
}

func newIndexAllColumnsNode(n unsafe.Pointer) *IndexAllColumnsNode {
	return &IndexAllColumnsNode{LeafBaseNode: newLeafBaseNode(n)}
}

func newIndexItemListNode(n unsafe.Pointer) *IndexItemListNode {
	return &IndexItemListNode{BaseNode: newBaseNode(n)}
}

func newIndexStoringExpressionListNode(n unsafe.Pointer) *IndexStoringExpressionListNode {
	return &IndexStoringExpressionListNode{BaseNode: newBaseNode(n)}
}

func newIndexUnnestExpressionListNode(n unsafe.Pointer) *IndexUnnestExpressionListNode {
	return &IndexUnnestExpressionListNode{BaseNode: newBaseNode(n)}
}

func newCreateIndexStatementNode(n unsafe.Pointer) *CreateIndexStatementNode {
	return &CreateIndexStatementNode{CreateStatementNode: newCreateStatementNode(n)}
}

func newExportDataStatementNode(n unsafe.Pointer) *ExportDataStatementNode {
	return &ExportDataStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newExportModelStatementNode(n unsafe.Pointer) *ExportModelStatementNode {
	return &ExportModelStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newCallStatementNode(n unsafe.Pointer) *CallStatementNode {
	return &CallStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newDefineTableStatementNode(n unsafe.Pointer) *DefineTableStatementNode {
	return &DefineTableStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newWithPartitionColumnsClauseNode(n unsafe.Pointer) *WithPartitionColumnsClauseNode {
	return &WithPartitionColumnsClauseNode{BaseNode: newBaseNode(n)}
}

func newCreateSnapshotTableStatementNode(n unsafe.Pointer) *CreateSnapshotTableStatementNode {
	return &CreateSnapshotTableStatementNode{CreateStatementNode: newCreateStatementNode(n)}
}

func newTypeParameterListNode(n unsafe.Pointer) *TypeParameterListNode {
	return &TypeParameterListNode{BaseNode: newBaseNode(n)}
}

func newTVFSchemaNode(n unsafe.Pointer) *TVFSchemaNode {
	return &TVFSchemaNode{BaseNode: newBaseNode(n)}
}

func newTVFSchemaColumnNode(n unsafe.Pointer) *TVFSchemaColumnNode {
	return &TVFSchemaColumnNode{BaseNode: newBaseNode(n)}
}

func newTableAndColumnInfoNode(n unsafe.Pointer) *TableAndColumnInfoNode {
	return &TableAndColumnInfoNode{BaseNode: newBaseNode(n)}
}

func newTableAndColumnInfoListNode(n unsafe.Pointer) *TableAndColumnInfoListNode {
	return &TableAndColumnInfoListNode{BaseNode: newBaseNode(n)}
}

func newTemplatedParameterTypeNode(n unsafe.Pointer) *TemplatedParameterTypeNode {
	return &TemplatedParameterTypeNode{BaseNode: newBaseNode(n)}
}

func newDefaultLiteralNode(n unsafe.Pointer) *DefaultLiteralNode {
	return &DefaultLiteralNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newAnalyzeStatementNode(n unsafe.Pointer) *AnalyzeStatementNode {
	return &AnalyzeStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newAssertStatementNode(n unsafe.Pointer) *AssertStatementNode {
	return &AssertStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newAssertRowsModifiedNode(n unsafe.Pointer) *AssertRowsModifiedNode {
	return &AssertRowsModifiedNode{BaseNode: newBaseNode(n)}
}

func newReturningClauseNode(n unsafe.Pointer) *ReturningClauseNode {
	return &ReturningClauseNode{BaseNode: newBaseNode(n)}
}

func newDeleteStatementNode(n unsafe.Pointer) *DeleteStatementNode {
	return &DeleteStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newColumnAttributeBaseNode(n unsafe.Pointer) *ColumnAttributeBaseNode {
	return &ColumnAttributeBaseNode{BaseNode: newBaseNode(n)}
}

func newNotNullColumnAttributeNode(n unsafe.Pointer) *NotNullColumnAttributeNode {
	return &NotNullColumnAttributeNode{ColumnAttributeBaseNode: newColumnAttributeBaseNode(n)}
}

func newHiddenColumnAttributeNode(n unsafe.Pointer) *HiddenColumnAttributeNode {
	return &HiddenColumnAttributeNode{ColumnAttributeBaseNode: newColumnAttributeBaseNode(n)}
}

func newPrimaryKeyColumnAttributeNode(n unsafe.Pointer) *PrimaryKeyColumnAttributeNode {
	return &PrimaryKeyColumnAttributeNode{ColumnAttributeBaseNode: newColumnAttributeBaseNode(n)}
}

func newForeignKeyColumnAttributeNode(n unsafe.Pointer) *ForeignKeyColumnAttributeNode {
	return &ForeignKeyColumnAttributeNode{ColumnAttributeBaseNode: newColumnAttributeBaseNode(n)}
}

func newColumnAttributeListNode(n unsafe.Pointer) *ColumnAttributeListNode {
	return &ColumnAttributeListNode{BaseNode: newBaseNode(n)}
}

func newStructColumnFieldNode(n unsafe.Pointer) *StructColumnFieldNode {
	return &StructColumnFieldNode{BaseNode: newBaseNode(n)}
}

func newGeneratedColumnInfoNode(n unsafe.Pointer) *GeneratedColumnInfoNode {
	return &GeneratedColumnInfoNode{BaseNode: newBaseNode(n)}
}

func newTableElementBaseNode(n unsafe.Pointer) *TableElementBaseNode {
	return &TableElementBaseNode{BaseNode: newBaseNode(n)}
}

func newColumnDefinitionNode(n unsafe.Pointer) *ColumnDefinitionNode {
	return &ColumnDefinitionNode{TableElementBaseNode: newTableElementBaseNode(n)}
}

func newTableElementListNode(n unsafe.Pointer) *TableElementListNode {
	return &TableElementListNode{BaseNode: newBaseNode(n)}
}

func newColumnListNode(n unsafe.Pointer) *ColumnListNode {
	return &ColumnListNode{BaseNode: newBaseNode(n)}
}

func newColumnPositionNode(n unsafe.Pointer) *ColumnPositionNode {
	return &ColumnPositionNode{BaseNode: newBaseNode(n)}
}

func newInsertValuesRowNode(n unsafe.Pointer) *InsertValuesRowNode {
	return &InsertValuesRowNode{BaseNode: newBaseNode(n)}
}

func newInsertValuesRowListNode(n unsafe.Pointer) *InsertValuesRowListNode {
	return &InsertValuesRowListNode{BaseNode: newBaseNode(n)}
}

func newInsertStatementNode(n unsafe.Pointer) *InsertStatementNode {
	return &InsertStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newUpdateSetValueNode(n unsafe.Pointer) *UpdateSetValueNode {
	return &UpdateSetValueNode{BaseNode: newBaseNode(n)}
}

func newUpdateItemNode(n unsafe.Pointer) *UpdateItemNode {
	return &UpdateItemNode{BaseNode: newBaseNode(n)}
}

func newUpdateItemListNode(n unsafe.Pointer) *UpdateItemListNode {
	return &UpdateItemListNode{BaseNode: newBaseNode(n)}
}

func newUpdateStatementNode(n unsafe.Pointer) *UpdateStatementNode {
	return &UpdateStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newTrucateStatementNode(n unsafe.Pointer) *TrucateStatementNode {
	return &TrucateStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newMergeActionNode(n unsafe.Pointer) *MergeActionNode {
	return &MergeActionNode{BaseNode: newBaseNode(n)}
}

func newMergeWhenClauseNode(n unsafe.Pointer) *MergeWhenClauseNode {
	return &MergeWhenClauseNode{BaseNode: newBaseNode(n)}
}

func newMergeWhenClauseListNode(n unsafe.Pointer) *MergeWhenClauseListNode {
	return &MergeWhenClauseListNode{BaseNode: newBaseNode(n)}
}

func newMergeStatementNode(n unsafe.Pointer) *MergeStatementNode {
	return &MergeStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newPrivilegeNode(n unsafe.Pointer) *PrivilegeNode {
	return &PrivilegeNode{BaseNode: newBaseNode(n)}
}

func newPrivilegesNode(n unsafe.Pointer) *PrivilegesNode {
	return &PrivilegesNode{BaseNode: newBaseNode(n)}
}

func newGranteeListNode(n unsafe.Pointer) *GranteeListNode {
	return &GranteeListNode{BaseNode: newBaseNode(n)}
}

func newGrantStatementNode(n unsafe.Pointer) *GrantStatementNode {
	return &GrantStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newRevokeStatementNode(n unsafe.Pointer) *RevokeStatementNode {
	return &RevokeStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newRepeatableClauseNode(n unsafe.Pointer) *RepeatableClauseNode {
	return &RepeatableClauseNode{BaseNode: newBaseNode(n)}
}

func newFilterFieldsArgNode(n unsafe.Pointer) *FilterFieldsArgNode {
	return &FilterFieldsArgNode{BaseNode: newBaseNode(n)}
}

func newReplaceFieldsArgNode(n unsafe.Pointer) *ReplaceFieldsArgNode {
	return &ReplaceFieldsArgNode{BaseNode: newBaseNode(n)}
}

func newReplaceFieldsExpressionNode(n unsafe.Pointer) *ReplaceFieldsExpressionNode {
	return &ReplaceFieldsExpressionNode{ExpressionBaseNode: newExpressionBaseNode(n)}
}

func newSampleSizeNode(n unsafe.Pointer) *SampleSizeNode {
	return &SampleSizeNode{BaseNode: newBaseNode(n)}
}

func newWithWeightNode(n unsafe.Pointer) *WithWeightNode {
	return &WithWeightNode{BaseNode: newBaseNode(n)}
}

func newSampleSuffixNode(n unsafe.Pointer) *SampleSuffixNode {
	return &SampleSuffixNode{BaseNode: newBaseNode(n)}
}

func newSampleClauseNode(n unsafe.Pointer) *SampleClauseNode {
	return &SampleClauseNode{BaseNode: newBaseNode(n)}
}

func newAlterActionBaseNode(n unsafe.Pointer) *AlterActionBaseNode {
	return &AlterActionBaseNode{BaseNode: newBaseNode(n)}
}

func newSetOptionsActionNode(n unsafe.Pointer) *SetOptionsActionNode {
	return &SetOptionsActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newSetAsActionNode(n unsafe.Pointer) *SetAsActionNode {
	return &SetAsActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newAddConstraintActionNode(n unsafe.Pointer) *AddConstraintActionNode {
	return &AddConstraintActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newDropPrimaryKeyActionNode(n unsafe.Pointer) *DropPrimaryKeyActionNode {
	return &DropPrimaryKeyActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newDropConstraintActionNode(n unsafe.Pointer) *DropConstraintActionNode {
	return &DropConstraintActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newAlterConstraintEnforcementActionNode(n unsafe.Pointer) *AlterConstraintEnforcementActionNode {
	return &AlterConstraintEnforcementActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newAlterConstraintSetOptionsActionNode(n unsafe.Pointer) *AlterConstraintSetOptionsActionNode {
	return &AlterConstraintSetOptionsActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newAddColumnActionNode(n unsafe.Pointer) *AddColumnActionNode {
	return &AddColumnActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newDropColumnActionNode(n unsafe.Pointer) *DropColumnActionNode {
	return &DropColumnActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newRenameColumnActionNode(n unsafe.Pointer) *RenameColumnActionNode {
	return &RenameColumnActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newAlterColumnTypeActionNode(n unsafe.Pointer) *AlterColumnTypeActionNode {
	return &AlterColumnTypeActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newAlterColumnOptionsActionNode(n unsafe.Pointer) *AlterColumnOptionsActionNode {
	return &AlterColumnOptionsActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newAlterColumnSetDefaultActionNode(n unsafe.Pointer) *AlterColumnSetDefaultActionNode {
	return &AlterColumnSetDefaultActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newAlterColumnDropDefaultActionNode(n unsafe.Pointer) *AlterColumnDropDefaultActionNode {
	return &AlterColumnDropDefaultActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newAlterColumnDropNotNullActionNode(n unsafe.Pointer) *AlterColumnDropNotNullActionNode {
	return &AlterColumnDropNotNullActionNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newGrantToClauseNode(n unsafe.Pointer) *GrantToClauseNode {
	return &GrantToClauseNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newRestrictToClauseNode(n unsafe.Pointer) *RestrictToClauseNode {
	return &RestrictToClauseNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newAddToRestricteeListClauseNode(n unsafe.Pointer) *AddToRestricteeListClauseNode {
	return &AddToRestricteeListClauseNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newRemoveFromRestricteeListClauseNode(n unsafe.Pointer) *RemoveFromRestricteeListClauseNode {
	return &RemoveFromRestricteeListClauseNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newFilterUsingClauseNode(n unsafe.Pointer) *FilterUsingClauseNode {
	return &FilterUsingClauseNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newRevokeFromClauseNode(n unsafe.Pointer) *RevokeFromClauseNode {
	return &RevokeFromClauseNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newRenameToClauseNode(n unsafe.Pointer) *RenameToClauseNode {
	return &RenameToClauseNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newSetCollateClauseNode(n unsafe.Pointer) *SetCollateClauseNode {
	return &SetCollateClauseNode{AlterActionBaseNode: newAlterActionBaseNode(n)}
}

func newAlterActionListNode(n unsafe.Pointer) *AlterActionListNode {
	return &AlterActionListNode{BaseNode: newBaseNode(n)}
}

func newAlterAllRowAccessPoliciesStatementNode(n unsafe.Pointer) *AlterAllRowAccessPoliciesStatementNode {
	return &AlterAllRowAccessPoliciesStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newForeignKeyActionsNode(n unsafe.Pointer) *ForeignKeyActionsNode {
	return &ForeignKeyActionsNode{BaseNode: newBaseNode(n)}
}

func newForeignKeyReferenceNode(n unsafe.Pointer) *ForeignKeyReferenceNode {
	return &ForeignKeyReferenceNode{BaseNode: newBaseNode(n)}
}

func newElseifClauseNode(n unsafe.Pointer) *ElseifClauseNode {
	return &ElseifClauseNode{BaseNode: newBaseNode(n)}
}

func newElseifClauseListNode(n unsafe.Pointer) *ElseifClauseListNode {
	return &ElseifClauseListNode{BaseNode: newBaseNode(n)}
}

func newIfStatementNode(n unsafe.Pointer) *IfStatementNode {
	return &IfStatementNode{ScriptStatementNode: newScriptStatementNode(n)}
}

func newWhenThenClauseNode(n unsafe.Pointer) *WhenThenClauseNode {
	return &WhenThenClauseNode{BaseNode: newBaseNode(n)}
}

func newWhenThenClauseListNode(n unsafe.Pointer) *WhenThenClauseListNode {
	return &WhenThenClauseListNode{BaseNode: newBaseNode(n)}
}

func newCaseStatementNode(n unsafe.Pointer) *CaseStatementNode {
	return &CaseStatementNode{ScriptStatementNode: newScriptStatementNode(n)}
}

func newHintNode(n unsafe.Pointer) *HintNode {
	return &HintNode{BaseNode: newBaseNode(n)}
}

func newHintEntryNode(n unsafe.Pointer) *HintEntryNode {
	return &HintEntryNode{BaseNode: newBaseNode(n)}
}

func newUnpivotInItemLabelNode(n unsafe.Pointer) *UnpivotInItemLabelNode {
	return &UnpivotInItemLabelNode{BaseNode: newBaseNode(n)}
}

func newDescriptorNode(n unsafe.Pointer) *DescriptorNode {
	return &DescriptorNode{BaseNode: newBaseNode(n)}
}

func newColumnSchemaNode(n unsafe.Pointer) *ColumnSchemaNode {
	return &ColumnSchemaNode{BaseNode: newBaseNode(n)}
}

func newSimpleColumnSchemaNode(n unsafe.Pointer) *SimpleColumnSchemaNode {
	return &SimpleColumnSchemaNode{ColumnSchemaNode: newColumnSchemaNode(n)}
}

func newArrayColumnSchemaNode(n unsafe.Pointer) *ArrayColumnSchemaNode {
	return &ArrayColumnSchemaNode{ColumnSchemaNode: newColumnSchemaNode(n)}
}

func newTableConstraintBaseNode(n unsafe.Pointer) *TableConstraintBaseNode {
	return &TableConstraintBaseNode{TableElementBaseNode: newTableElementBaseNode(n)}
}

func newPrimaryKeyNode(n unsafe.Pointer) *PrimaryKeyNode {
	return &PrimaryKeyNode{TableConstraintBaseNode: newTableConstraintBaseNode(n)}
}

func newForeignKeyNode(n unsafe.Pointer) *ForeignKeyNode {
	return &ForeignKeyNode{TableConstraintBaseNode: newTableConstraintBaseNode(n)}
}

func newCheckConstraintNode(n unsafe.Pointer) *CheckConstraintNode {
	return &CheckConstraintNode{TableConstraintBaseNode: newTableConstraintBaseNode(n)}
}

func newDescriptorColumnNode(n unsafe.Pointer) *DescriptorColumnNode {
	return &DescriptorColumnNode{BaseNode: newBaseNode(n)}
}

func newDescriptorColumnListNode(n unsafe.Pointer) *DescriptorColumnListNode {
	return &DescriptorColumnListNode{BaseNode: newBaseNode(n)}
}

func newCreateEntityStatementNode(n unsafe.Pointer) *CreateEntityStatementNode {
	return &CreateEntityStatementNode{CreateStatementNode: newCreateStatementNode(n)}
}

func newRaiseStatementNode(n unsafe.Pointer) *RaiseStatementNode {
	return &RaiseStatementNode{ScriptStatementNode: newScriptStatementNode(n)}
}

func newExceptionHandlerNode(n unsafe.Pointer) *ExceptionHandlerNode {
	return &ExceptionHandlerNode{BaseNode: newBaseNode(n)}
}

func newExceptionHandlerListNode(n unsafe.Pointer) *ExceptionHandlerListNode {
	return &ExceptionHandlerListNode{BaseNode: newBaseNode(n)}
}

func newBeginEndBlockNode(n unsafe.Pointer) *BeginEndBlockNode {
	return &BeginEndBlockNode{ScriptStatementNode: newScriptStatementNode(n)}
}

func newIdentifierListNode(n unsafe.Pointer) *IdentifierListNode {
	return &IdentifierListNode{BaseNode: newBaseNode(n)}
}

func newVariableDeclarationNode(n unsafe.Pointer) *VariableDeclarationNode {
	return &VariableDeclarationNode{ScriptStatementNode: newScriptStatementNode(n)}
}

func newUntilClauseNode(n unsafe.Pointer) *UntilClauseNode {
	return &UntilClauseNode{BaseNode: newBaseNode(n)}
}

func newBreakContinueStatementNode(n unsafe.Pointer) *BreakContinueStatementNode {
	return &BreakContinueStatementNode{ScriptStatementNode: newScriptStatementNode(n)}
}

func newBreakStatementNode(n unsafe.Pointer) *BreakStatementNode {
	return &BreakStatementNode{BreakContinueStatementNode: newBreakContinueStatementNode(n)}
}

func newContinueStatementNode(n unsafe.Pointer) *ContinueStatementNode {
	return &ContinueStatementNode{BreakContinueStatementNode: newBreakContinueStatementNode(n)}
}

func newDropPrivilegeRestrictionStatementNode(n unsafe.Pointer) *DropPrivilegeRestrictionStatementNode {
	return &DropPrivilegeRestrictionStatementNode{DdlStatementBaseNode: newDdlStatementBaseNode(n)}
}

func newDropRowAccessPolicyStatementNode(n unsafe.Pointer) *DropRowAccessPolicyStatementNode {
	return &DropRowAccessPolicyStatementNode{DdlStatementBaseNode: newDdlStatementBaseNode(n)}
}

func newCreatePrivilegeRestrictionStatementNode(n unsafe.Pointer) *CreatePrivilegeRestrictionStatementNode {
	return &CreatePrivilegeRestrictionStatementNode{CreateStatementNode: newCreateStatementNode(n)}
}

func newCreateRowAccessPolicyStatementNode(n unsafe.Pointer) *CreateRowAccessPolicyStatementNode {
	return &CreateRowAccessPolicyStatementNode{CreateStatementNode: newCreateStatementNode(n)}
}

func newDropStatementNode(n unsafe.Pointer) *DropStatementNode {
	return &DropStatementNode{DdlStatementBaseNode: newDdlStatementBaseNode(n)}
}

func newReturnStatementNode(n unsafe.Pointer) *ReturnStatementNode {
	return &ReturnStatementNode{ScriptStatementNode: newScriptStatementNode(n)}
}

func newSingleAssignmentNode(n unsafe.Pointer) *SingleAssignmentNode {
	return &SingleAssignmentNode{ScriptStatementNode: newScriptStatementNode(n)}
}

func newParameterAssignmentNode(n unsafe.Pointer) *ParameterAssignmentNode {
	return &ParameterAssignmentNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newSystemVariableAssignmentNode(n unsafe.Pointer) *SystemVariableAssignmentNode {
	return &SystemVariableAssignmentNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newAssignmentFromStructNode(n unsafe.Pointer) *AssignmentFromStructNode {
	return &AssignmentFromStructNode{ScriptStatementNode: newScriptStatementNode(n)}
}

func newCreateTableStmtBaseNode(n unsafe.Pointer) *CreateTableStmtBaseNode {
	return &CreateTableStmtBaseNode{CreateStatementNode: newCreateStatementNode(n)}
}

func newCreateTableStatementNode(n unsafe.Pointer) *CreateTableStatementNode {
	return &CreateTableStatementNode{CreateTableStmtBaseNode: newCreateTableStmtBaseNode(n)}
}

func newCreateExternalTableStatementNode(n unsafe.Pointer) *CreateExternalTableStatementNode {
	return &CreateExternalTableStatementNode{CreateTableStmtBaseNode: newCreateTableStmtBaseNode(n)}
}

func newCreateViewStatementBaseNode(n unsafe.Pointer) *CreateViewStatementBaseNode {
	return &CreateViewStatementBaseNode{CreateStatementNode: newCreateStatementNode(n)}
}

func newCreateViewStatementNode(n unsafe.Pointer) *CreateViewStatementNode {
	return &CreateViewStatementNode{CreateViewStatementBaseNode: newCreateViewStatementBaseNode(n)}
}

func newCreateMaterializedViewStatementNode(n unsafe.Pointer) *CreateMaterializedViewStatementNode {
	return &CreateMaterializedViewStatementNode{CreateViewStatementBaseNode: newCreateViewStatementBaseNode(n)}
}

func newLoopStatementNode(n unsafe.Pointer) *LoopStatementNode {
	return &LoopStatementNode{ScriptStatementNode: newScriptStatementNode(n)}
}

func newWhileStatementNode(n unsafe.Pointer) *WhileStatementNode {
	return &WhileStatementNode{LoopStatementNode: newLoopStatementNode(n)}
}

func newRepeatStatementNode(n unsafe.Pointer) *RepeatStatementNode {
	return &RepeatStatementNode{LoopStatementNode: newLoopStatementNode(n)}
}

func newForInStatementNode(n unsafe.Pointer) *ForInStatementNode {
	return &ForInStatementNode{LoopStatementNode: newLoopStatementNode(n)}
}

func newAlterStatementBaseNode(n unsafe.Pointer) *AlterStatementBaseNode {
	return &AlterStatementBaseNode{DdlStatementBaseNode: newDdlStatementBaseNode(n)}
}

func newAlterDatabaseStatementNode(n unsafe.Pointer) *AlterDatabaseStatementNode {
	return &AlterDatabaseStatementNode{AlterStatementBaseNode: newAlterStatementBaseNode(n)}
}

func newAlterSchemaStatementNode(n unsafe.Pointer) *AlterSchemaStatementNode {
	return &AlterSchemaStatementNode{AlterStatementBaseNode: newAlterStatementBaseNode(n)}
}

func newAlterTableStatementNode(n unsafe.Pointer) *AlterTableStatementNode {
	return &AlterTableStatementNode{AlterStatementBaseNode: newAlterStatementBaseNode(n)}
}

func newAlterViewStatementNode(n unsafe.Pointer) *AlterViewStatementNode {
	return &AlterViewStatementNode{AlterStatementBaseNode: newAlterStatementBaseNode(n)}
}

func newAlterMaterializedViewStatementNode(n unsafe.Pointer) *AlterMaterializedViewStatementNode {
	return &AlterMaterializedViewStatementNode{AlterStatementBaseNode: newAlterStatementBaseNode(n)}
}

func newAlterPrivilegeRestrictionStatementNode(n unsafe.Pointer) *AlterPrivilegeRestrictionStatementNode {
	return &AlterPrivilegeRestrictionStatementNode{AlterStatementBaseNode: newAlterStatementBaseNode(n)}
}

func newAlterRowAccessPolicyStatementNode(n unsafe.Pointer) *AlterRowAccessPolicyStatementNode {
	return &AlterRowAccessPolicyStatementNode{AlterStatementBaseNode: newAlterStatementBaseNode(n)}
}

func newAlterEntityStatementNode(n unsafe.Pointer) *AlterEntityStatementNode {
	return &AlterEntityStatementNode{AlterStatementBaseNode: newAlterStatementBaseNode(n)}
}

func newCreateFunctionStmtBaseNode(n unsafe.Pointer) *CreateFunctionStmtBaseNode {
	return &CreateFunctionStmtBaseNode{CreateStatementNode: newCreateStatementNode(n)}
}

func newCreateFunctionStatementNode(n unsafe.Pointer) *CreateFunctionStatementNode {
	return &CreateFunctionStatementNode{CreateFunctionStmtBaseNode: newCreateFunctionStmtBaseNode(n)}
}

func newCreateTableFunctionStatementNode(n unsafe.Pointer) *CreateTableFunctionStatementNode {
	return &CreateTableFunctionStatementNode{CreateFunctionStmtBaseNode: newCreateFunctionStmtBaseNode(n)}
}

func newStructColumnSchemaNode(n unsafe.Pointer) *StructColumnSchemaNode {
	return &StructColumnSchemaNode{ColumnSchemaNode: newColumnSchemaNode(n)}
}

func newInferredTypeColumnSchemaNode(n unsafe.Pointer) *InferredTypeColumnSchemaNode {
	return &InferredTypeColumnSchemaNode{ColumnSchemaNode: newColumnSchemaNode(n)}
}

func newExecuteIntoClauseNode(n unsafe.Pointer) *ExecuteIntoClauseNode {
	return &ExecuteIntoClauseNode{BaseNode: newBaseNode(n)}
}

func newExecuteUsingArgumentNode(n unsafe.Pointer) *ExecuteUsingArgumentNode {
	return &ExecuteUsingArgumentNode{BaseNode: newBaseNode(n)}
}

func newExecuteUsingClauseNode(n unsafe.Pointer) *ExecuteUsingClauseNode {
	return &ExecuteUsingClauseNode{BaseNode: newBaseNode(n)}
}

func newExecuteImmediateStatementNode(n unsafe.Pointer) *ExecuteImmediateStatementNode {
	return &ExecuteImmediateStatementNode{StatementBaseNode: newStatementBaseNode(n)}
}

func newAuxLoadDataFromFilesOptionsListNode(n unsafe.Pointer) *AuxLoadDataFromFilesOptionsListNode {
	return &AuxLoadDataFromFilesOptionsListNode{BaseNode: newBaseNode(n)}
}

func newAuxLoadDataStatementNode(n unsafe.Pointer) *AuxLoadDataStatementNode {
	return &AuxLoadDataStatementNode{CreateTableStmtBaseNode: newCreateTableStmtBaseNode(n)}
}

func newLabelNode(n unsafe.Pointer) *LabelNode {
	return &LabelNode{BaseNode: newBaseNode(n)}
}

func newNode(n unsafe.Pointer) Node {
	var kind int
	internalparser.ASTNode_node_kind(n, &kind)
	switch Kind(kind) {
	case Unknown:
		return nil
	case Fake:
		return nil
	case AbortBatchStatement:
		return newAbortBatchStatementNode(n)
	case AddColumnAction:
		return newAddColumnActionNode(n)
	case AddConstraintAction:
		return newAddConstraintActionNode(n)
	case AddToRestricteeListClause:
		return newAddToRestricteeListClauseNode(n)
	case FunctionCallWithGroupRows:
		return newFunctionCallWithGroupRowsNode(n)
	case Alias:
		return newAliasNode(n)
	case AlterActionList:
		return newAlterActionListNode(n)
	case AlterAllRowAccessPoliciesStatement:
		return newAlterAllRowAccessPoliciesStatementNode(n)
	case AlterColumnOptionsAction:
		return newAlterColumnOptionsActionNode(n)
	case AlterColumnDropNotNullAction:
		return newAlterColumnDropNotNullActionNode(n)
	case AlterColumnTypeAction:
		return newAlterColumnTypeActionNode(n)
	case AlterColumnSetDefaultAction:
		return newAlterColumnSetDefaultActionNode(n)
	case AlterColumnDropDefaultAction:
		return newAlterColumnDropDefaultActionNode(n)
	case AlterConstraintEnforcementAction:
		return newAlterConstraintEnforcementActionNode(n)
	case AlterConstraintSetOptionsAction:
		return newAlterConstraintSetOptionsActionNode(n)
	case AlterDatabaseStatement:
		return newAlterDatabaseStatementNode(n)
	case AlterEntityStatement:
		return newAlterEntityStatementNode(n)
	case AlterMaterializedViewStatement:
		return newAlterMaterializedViewStatementNode(n)
	case AlterPrivilegeRestrictionStatement:
		return newAlterPrivilegeRestrictionStatementNode(n)
	case AlterRowAccessPolicyStatement:
		return newAlterRowAccessPolicyStatementNode(n)
	case AlterSchemaStatement:
		return newAlterSchemaStatementNode(n)
	case AlterTableStatement:
		return newAlterTableStatementNode(n)
	case AlterViewStatement:
		return newAlterViewStatementNode(n)
	case AnalyticFunctionCall:
		return newAnalyticFunctionCallNode(n)
	case AnalyzeStatement:
		return newAnalyzeStatementNode(n)
	case AndExpr:
		return newAndExprNode(n)
	case AnySomeAllOp:
		return newAnySomeAllOpNode(n)
	case ArrayColumnSchema:
		return newArrayColumnSchemaNode(n)
	case ArrayConstructor:
		return newArrayConstructorNode(n)
	case ArrayElement:
		return newArrayElementNode(n)
	case ArrayType:
		return newArrayTypeNode(n)
	case AssertRowsModified:
		return newAssertRowsModifiedNode(n)
	case AssertStatement:
		return newAssertStatementNode(n)
	case AssignmentFromStruct:
		return newAssignmentFromStructNode(n)
	case BeginStatement:
		return newBeginStatementNode(n)
	case BetweenExpression:
		return newBetweenExpressionNode(n)
	case AuxLoadDataFromFilesOptionsList:
		return newAuxLoadDataFromFilesOptionsListNode(n)
	case AuxLoadDataStatement:
		return newAuxLoadDataStatementNode(n)
	case BignumericLiteral:
		return newBigNumericLiteralNode(n)
	case BinaryExpression:
		return newBinaryExpressionNode(n)
	case BitwiseShiftExpression:
		return newBitwiseShiftExpressionNode(n)
	case BeginEndBlock:
		return newBeginEndBlockNode(n)
	case BooleanLiteral:
		return newBooleanLiteralNode(n)
	case BreakStatement:
		return newBreakStatementNode(n)
	case BytesLiteral:
		return newBytesLiteralNode(n)
	case CallStatement:
		return newCallStatementNode(n)
	case CaseStatement:
		return newCaseStatementNode(n)
	case CaseNoValueExpression:
		return newCaseNoValueExpressionNode(n)
	case CaseValueExpression:
		return newCaseValueExpressionNode(n)
	case CastExpression:
		return newCastExpressionNode(n)
	case CheckConstraint:
		return newCheckConstraintNode(n)
	case ClampedBetweenModifier:
		return newClampedBetweenModifierNode(n)
	case CloneDataSource:
		return newCloneDataSourceNode(n)
	case CloneDataSourceList:
		return newCloneDataSourceListNode(n)
	case CloneDataStatement:
		return newCloneDataStatementNode(n)
	case ClusterBy:
		return newClusterByNode(n)
	case Collate:
		return newCollateNode(n)
	case ColumnAttributeList:
		return newColumnAttributeListNode(n)
	case ColumnDefinition:
		return newColumnDefinitionNode(n)
	case ColumnList:
		return newColumnListNode(n)
	case ColumnPosition:
		return newColumnPositionNode(n)
	case CommitStatement:
		return newCommitStatementNode(n)
	case ConnectionClause:
		return newConnectionClauseNode(n)
	case ContinueStatement:
		return newContinueStatementNode(n)
	case CopyDataSource:
		return newCopyDataSourceNode(n)
	case CreateConstantStatement:
		return newCreateConstantStatementNode(n)
	case CreateDatabaseStatement:
		return newCreateDatabaseStatementNode(n)
	case CreateExternalTableStatement:
		return newCreateExternalTableStatementNode(n)
	case CreateFunctionStatement:
		return newCreateFunctionStatementNode(n)
	case CreateIndexStatement:
		return newCreateIndexStatementNode(n)
	case CreateModelStatement:
		return newCreateModelStatementNode(n)
	case CreateProcedureStatement:
		return newCreateProcedureStatementNode(n)
	case CreatePrivilegeRestrictionStatement:
		return newCreatePrivilegeRestrictionStatementNode(n)
	case CreateRowAccessPolicyStatement:
		return newCreateRowAccessPolicyStatementNode(n)
	case CreateSchemaStatement:
		return newCreateSchemaStatementNode(n)
	case CreateSnapshotTableStatement:
		return newCreateSnapshotTableStatementNode(n)
	case CreateTableFunctionStatement:
		return newCreateTableFunctionStatementNode(n)
	case CreateTableStatement:
		return newCreateTableStatementNode(n)
	case CreateEntityStatement:
		return newCreateEntityStatementNode(n)
	case CreateViewStatement:
		return newCreateViewStatementNode(n)
	case CreateMaterializedViewStatement:
		return newCreateMaterializedViewStatementNode(n)
	case DateOrTimeLiteral:
		return newDateOrTimeLiteralNode(n)
	case DefaultLiteral:
		return newDefaultLiteralNode(n)
	case DefineTableStatement:
		return newDefineTableStatementNode(n)
	case DeleteStatement:
		return newDeleteStatementNode(n)
	case DescribeStatement:
		return newDescribeStatementNode(n)
	case DescriptorColumn:
		return newDescriptorColumnNode(n)
	case DescriptorColumnList:
		return newDescriptorColumnListNode(n)
	case Descriptor:
		return newDescriptorNode(n)
	case DotGeneralizedField:
		return newDotGeneralizedFieldNode(n)
	case DotIdentifier:
		return newDotIdentifierNode(n)
	case DotStar:
		return newDotStarNode(n)
	case DotStarWithModifiers:
		return newDotStarWithModifiersNode(n)
	case DropAllRowAccessPoliciesStatement:
		return newDropAllRowAccessPoliciesStatementNode(n)
	case DropColumnAction:
		return newDropColumnActionNode(n)
	case DropConstraintAction:
		return newDropConstraintActionNode(n)
	case DropEntityStatement:
		return newDropEntityStatementNode(n)
	case DropFunctionStatement:
		return newDropFunctionStatementNode(n)
	case DropPrimaryKeyAction:
		return newDropPrimaryKeyActionNode(n)
	case DropPrivilegeRestrictionStatement:
		return newDropPrivilegeRestrictionStatementNode(n)
	case DropRowAccessPolicyStatement:
		return newDropRowAccessPolicyStatementNode(n)
	case DropSearchIndexStatement:
		return newDropSearchIndexStatementNode(n)
	case DropStatement:
		return newDropStatementNode(n)
	case DropTableFunctionStatement:
		return newDropTableFunctionStatementNode(n)
	case DropMaterializedViewStatement:
		return newDropMaterializedViewStatementNode(n)
	case DropSnapshotTableStatement:
		return newDropSnapshotTableStatementNode(n)
	case ElseifClause:
		return newElseifClauseNode(n)
	case ElseifClauseList:
		return newElseifClauseListNode(n)
	case ExceptionHandler:
		return newExceptionHandlerNode(n)
	case ExceptionHandlerList:
		return newExceptionHandlerListNode(n)
	case ExecuteImmediateStatement:
		return newExecuteImmediateStatementNode(n)
	case ExecuteIntoClause:
		return newExecuteIntoClauseNode(n)
	case ExecuteUsingArgument:
		return newExecuteUsingArgumentNode(n)
	case ExecuteUsingClause:
		return newExecuteUsingClauseNode(n)
	case ExplainStatement:
		return newExplainStatementNode(n)
	case ExportDataStatement:
		return newExportDataStatementNode(n)
	case ExportModelStatement:
		return newExportModelStatementNode(n)
	case ExpressionSubquery:
		return newExpressionSubqueryNode(n)
	case ExtractExpression:
		return newExtractExpressionNode(n)
	case FilterFieldsArg:
		return newFilterFieldsArgNode(n)
	case FilterFieldsExpression:
		return nil
	case FilterUsingClause:
		return newFilterUsingClauseNode(n)
	case FloatLiteral:
		return newFloatLiteralNode(n)
	case ForInStatement:
		return newForInStatementNode(n)
	case ForeignKey:
		return newForeignKeyNode(n)
	case ForeignKeyActions:
		return newForeignKeyActionsNode(n)
	case ForeignKeyColumnAttribute:
		return newForeignKeyColumnAttributeNode(n)
	case ForeignKeyReference:
		return newForeignKeyReferenceNode(n)
	case FormatClause:
		return newFormatClauseNode(n)
	case ForSystemTime:
		return newForSystemTimeNode(n)
	case FromClause:
		return newFromClauseNode(n)
	case FunctionCall:
		return newFunctionCallNode(n)
	case FunctionDeclaration:
		return newFunctionDeclarationNode(n)
	case FunctionParameter:
		return newFunctionParameterNode(n)
	case FunctionParameters:
		return newFunctionParametersNode(n)
	case GeneratedColumnInfo:
		return newGeneratedColumnInfoNode(n)
	case GranteeList:
		return newGranteeListNode(n)
	case GrantStatement:
		return newGrantStatementNode(n)
	case GrantToClause:
		return newGrantToClauseNode(n)
	case RestrictToClause:
		return newRestrictToClauseNode(n)
	case GroupBy:
		return newGroupByNode(n)
	case GroupingItem:
		return newGroupingItemNode(n)
	case Having:
		return newHavingNode(n)
	case HavingModifier:
		return newHavingModifierNode(n)
	case HiddenColumnAttribute:
		return newHiddenColumnAttributeNode(n)
	case Hint:
		return newHintNode(n)
	case HintedStatement:
		return newHintedStatementNode(n)
	case HintEntry:
		return newHintEntryNode(n)
	case Identifier:
		return newIdentifierNode(n)
	case IdentifierList:
		return newIdentifierListNode(n)
	case IfStatement:
		return newIfStatementNode(n)
	case ImportStatement:
		return newImportStatementNode(n)
	case InExpression:
		return newInExpressionNode(n)
	case InList:
		return newInListNode(n)
	case IndexAllColumns:
		return newIndexAllColumnsNode(n)
	case IndexItemList:
		return newIndexItemListNode(n)
	case IndexStoringExpressionList:
		return newIndexStoringExpressionListNode(n)
	case IndexUnnestExpressionList:
		return newIndexUnnestExpressionListNode(n)
	case InferredTypeColumnSchema:
		return newInferredTypeColumnSchemaNode(n)
	case InsertStatement:
		return newInsertStatementNode(n)
	case InsertValuesRow:
		return newInsertValuesRowNode(n)
	case InsertValuesRowList:
		return newInsertValuesRowListNode(n)
	case IntervalExpr:
		return newIntervalExprNode(n)
	case IntoAlias:
		return newIntoAliasNode(n)
	case IntLiteral:
		return newIntLiteralNode(n)
	case Join:
		return newJoinNode(n)
	case JoinLiteral:
		return nil
	case Label:
		return newLabelNode(n)
	case Lambda:
		return newLambdaNode(n)
	case LikeExpression:
		return newLikeExpressionNode(n)
	case LimitOffset:
		return newLimitOffsetNode(n)
	case MaxLiteral:
		return newMaxLiteralNode(n)
	case MergeAction:
		return newMergeActionNode(n)
	case MergeStatement:
		return newMergeStatementNode(n)
	case MergeWhenClause:
		return newMergeWhenClauseNode(n)
	case MergeWhenClauseList:
		return newMergeWhenClauseListNode(n)
	case ModelClause:
		return newModelClauseNode(n)
	case ModuleStatement:
		return newModuleStatementNode(n)
	case NamedArgument:
		return newNamedArgumentNode(n)
	case NewConstructor:
		return newNewConstructorNode(n)
	case NewConstructorArg:
		return newNewConstructorArgNode(n)
	case NotNullColumnAttribute:
		return newNotNullColumnAttributeNode(n)
	case NullLiteral:
		return newNullLiteralNode(n)
	case NullOrder:
		return newNullOrderNode(n)
	case NumericLiteral:
		return newNumericLiteralNode(n)
	case OnClause:
		return newOnClauseNode(n)
	case OnOrUsingClauseList:
		return newOnOrUsingClauseListNode(n)
	case OptionsEntry:
		return newOptionsEntryNode(n)
	case OptionsList:
		return newOptionsListNode(n)
	case OrderBy:
		return newOrderByNode(n)
	case OrderingExpression:
		return newOrderingExpressionNode(n)
	case OrExpr:
		return newOrExprNode(n)
	case ParameterAssignment:
		return newParameterAssignmentNode(n)
	case ParameterExpr:
		return newParameterExprNode(n)
	case ParenthesizedJoin:
		return newParenthesizedJoinNode(n)
	case PartitionBy:
		return newPartitionByNode(n)
	case PathExpression:
		return newPathExpressionNode(n)
	case PathExpressionList:
		return newPathExpressionListNode(n)
	case PivotClause:
		return newPivotClauseNode(n)
	case UnpivotClause:
		return newUnpivotClauseNode(n)
	case UnpivotInItemLabel:
		return newUnpivotInItemLabelNode(n)
	case UnpivotInItem:
		return newUnpivotInItemNode(n)
	case UnpivotInItemList:
		return newUnpivotInItemListNode(n)
	case PivotExpression:
		return newPivotExpressionNode(n)
	case PivotExpressionList:
		return newPivotExpressionListNode(n)
	case PivotValue:
		return newPivotValueNode(n)
	case PivotValueList:
		return newPivotValueListNode(n)
	case PrimaryKey:
		return newPrimaryKeyNode(n)
	case PrimaryKeyColumnAttribute:
		return newPrimaryKeyColumnAttributeNode(n)
	case Privilege:
		return newPrivilegeNode(n)
	case Privileges:
		return newPrivilegesNode(n)
	case Qualify:
		return newQualifyNode(n)
	case Query:
		return newQueryNode(n)
	case QueryStatement:
		return newQueryStatementNode(n)
	case RaiseStatement:
		return newRaiseStatementNode(n)
	case RemoveFromRestricteeListClause:
		return newRemoveFromRestricteeListClauseNode(n)
	case RenameColumnAction:
		return newRenameColumnActionNode(n)
	case RenameToClause:
		return newRenameToClauseNode(n)
	case RenameStatement:
		return newRenameStatementNode(n)
	case RepeatStatement:
		return newRepeatStatementNode(n)
	case RepeatableClause:
		return newRepeatableClauseNode(n)
	case ReplaceFieldsArg:
		return newReplaceFieldsArgNode(n)
	case ReplaceFieldsExpression:
		return newReplaceFieldsExpressionNode(n)
	case ReturnStatement:
		return newReturnStatementNode(n)
	case ReturningClause:
		return newReturningClauseNode(n)
	case RevokeFromClause:
		return newRevokeFromClauseNode(n)
	case RevokeStatement:
		return newRevokeStatementNode(n)
	case RollbackStatement:
		return newRollbackStatementNode(n)
	case Rollup:
		return newRollupNode(n)
	case RunBatchStatement:
		return newRunBatchStatementNode(n)
	case SampleClause:
		return newSampleClauseNode(n)
	case SampleSize:
		return newSampleSizeNode(n)
	case SampleSuffix:
		return newSampleSuffixNode(n)
	case Script:
		return newScriptBaseNode(n)
	case Select:
		return newSelectNode(n)
	case SelectAs:
		return newSelectAsNode(n)
	case SelectColumn:
		return newSelectColumnNode(n)
	case SelectList:
		return newSelectListNode(n)
	case SetOperation:
		return newSetOperationNode(n)
	case SetOptionsAction:
		return newSetOptionsActionNode(n)
	case SetAsAction:
		return newSetAsActionNode(n)
	case SetCollateClause:
		return newSetCollateClauseNode(n)
	case SetTransactionStatement:
		return newSetTransactionStatementNode(n)
	case SingleAssignment:
		return newSingleAssignmentNode(n)
	case ShowStatement:
		return newShowStatementNode(n)
	case SimpleColumnSchema:
		return newSimpleColumnSchemaNode(n)
	case SimpleType:
		return newSimpleTypeNode(n)
	case SqlFunctionBody:
		return newSqlFunctionBodyNode(n)
	case Star:
		return newStarNode(n)
	case StarExceptList:
		return newStarExceptListNode(n)
	case StarModifiers:
		return newStarModifiersNode(n)
	case StarReplaceItem:
		return newStarReplaceItemNode(n)
	case StarWithModifiers:
		return newStarWithModifiersNode(n)
	case StarBatchStatement:
		return nil
	case StatementList:
		return newStatementListNode(n)
	case StringLiteral:
		return newStringLiteralNode(n)
	case StructColumnField:
		return newStructColumnFieldNode(n)
	case StructColumnSchema:
		return newStructColumnSchemaNode(n)
	case StructConstructorArg:
		return newStructConstructorArgNode(n)
	case StructConstructorWithKeyword:
		return newStructConstructorWithKeywordNode(n)
	case StructConstructorWithParens:
		return newStructConstructorWithParensNode(n)
	case StructField:
		return newStructFieldNode(n)
	case StructType:
		return newStructTypeNode(n)
	case SystemVariableAssignment:
		return newSystemVariableAssignmentNode(n)
	case SystemVariableExpr:
		return newSystemVariableExprNode(n)
	case TableAndColumnInfo:
		return newTableAndColumnInfoNode(n)
	case TableAndColumnInfoList:
		return newTableAndColumnInfoListNode(n)
	case TableClause:
		return newTableClauseNode(n)
	case TableElementList:
		return newTableElementListNode(n)
	case TablePathExpression:
		return newTablePathExpressionNode(n)
	case TableSubquery:
		return newTableSubqueryNode(n)
	case TemplatedParameterType:
		return newTemplatedParameterTypeNode(n)
	case TransactionIsolationLevel:
		return newTransactionIsolationLevelNode(n)
	case TransactionModeList:
		return newTransactionModeListNode(n)
	case TransactionReadWriteMode:
		return newTransactionReadWriteModeNode(n)
	case TransformClause:
		return newTransformClauseNode(n)
	case TrucateStatement:
		return newTrucateStatementNode(n)
	case Tvf:
		return newTVFNode(n)
	case TvfArgument:
		return newTVFArgumentNode(n)
	case TvfSchema:
		return newTVFSchemaNode(n)
	case TvfSchemaColumn:
		return newTVFSchemaColumnNode(n)
	case TypeParameterList:
		return newTypeParameterListNode(n)
	case UnaryExpression:
		return newUnaryExpressionNode(n)
	case UnnestExpression:
		return newUnnestExpressionNode(n)
	case UnnestExpressionWithOptAliasAndOffset:
		return newUnnestExpressionWithOptAliasAndOffsetNode(n)
	case UntilClause:
		return newUntilClauseNode(n)
	case UpdateItem:
		return newUpdateItemNode(n)
	case UpdateItemList:
		return newUpdateItemListNode(n)
	case UpdateSetValue:
		return newUpdateSetValueNode(n)
	case UpdateStatement:
		return newUpdateStatementNode(n)
	case UsingClause:
		return newUsingClauseNode(n)
	case VariableDeclaration:
		return newVariableDeclarationNode(n)
	case WhenThenClause:
		return newWhenThenClauseNode(n)
	case WhenThenClauseList:
		return newWhenThenClauseListNode(n)
	case WhereClause:
		return newWhereClauseNode(n)
	case WhileStatement:
		return newWhileStatementNode(n)
	case WindowClause:
		return newWindowClauseNode(n)
	case WindowDefinition:
		return newWindowDefinitionNode(n)
	case WindowFrame:
		return newWindowFrameNode(n)
	case WindowFrameExpr:
		return newWindowFrameExprNode(n)
	case WindowSpecification:
		return newWindowSpecificationNode(n)
	case WithClause:
		return newWithClauseNode(n)
	case WithClauseEntry:
		return newWithClauseEntryNode(n)
	case WithConnectionClause:
		return newWithConnectionClauseNode(n)
	case WithGroupRows:
		return newWithGroupRowsNode(n)
	case WithOffset:
		return newWithOffsetNode(n)
	case WithPartitionColumnsClause:
		return newWithPartitionColumnsClauseNode(n)
	case WithWeight:
		return newWithWeightNode(n)
	}
	return newBaseNode(n)
}

func getNodeRaw(n Node) unsafe.Pointer {
	return n.getRaw()
}