//
// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

#ifndef ZETASQL_PARSER_PARSE_TREE_H_
#define ZETASQL_PARSER_PARSE_TREE_H_

// The complete set of headers comprises 3 parts:
// - ast_node.h is the header for ASTNode, the parent
// - parse_tree_generated.h contains classes generated by gen_parse_tree.py
// - additional miscellany defined in this file.
//
#include <stack>

#include "zetasql/parser/ast_node.h"  
#include "zetasql/parser/parse_tree_generated.h"  
#include "zetasql/parser/visit_result.h"

namespace zetasql {

// This is a fake ASTNode implementation that exists only for tests,
// which may need to pass an ASTNode* to some methods.
class FakeASTNode final : public ASTNode {
 public:
  static constexpr ASTNodeKind kConcreteNodeKind = AST_FAKE;

  FakeASTNode() : ASTNode(kConcreteNodeKind) {}
  void Accept(ParseTreeVisitor* visitor, void* data) const override {
    ZETASQL_LOG(FATAL) << "FakeASTNode does not support Accept";
  }
  absl::StatusOr<VisitResult> Accept(
      NonRecursiveParseTreeVisitor* visitor) const override {
    ZETASQL_LOG(FATAL) << "FakeASTNode does not support Accept";
  }

  void InitFields() final {
    {
      FieldLoader fl(this);  // Triggers check that there were no children.
    }
    set_start_location(ParseLocationPoint::FromByteOffset("fake_filename", 7));
    set_end_location(ParseLocationPoint::FromByteOffset("fake_filename", 10));
  }
};

inline IdString ASTAlias::GetAsIdString() const {
  return identifier()->GetAsIdString();
}

namespace parse_tree_internal {

// Concrete types (the 'leaves' of the hierarchy) must be constructible.
template <typename T>
using EnableIfConcrete =
    typename std::enable_if<std::is_constructible<T>::value, int>::type;

// Non Concrete types (internal nodes) of the hierarchy must _not_ be
// constructible.
template <typename T>
using EnableIfNotConcrete =
    typename std::enable_if<!std::is_constructible<T>::value, int>::type;

// GetAsOrNull implementation optimized for concrete types.  We assume that all
// concrete types define:
//   static constexpr Type kConcreteNodeKind;
//
// This allows us to avoid invoking dynamic_cast.
template <typename T, typename MaybeConstRoot, EnableIfConcrete<T> = 0>
inline T* GetAsOrNullImpl(MaybeConstRoot* n) {
  if (n->node_kind() == T::kConcreteNodeKind) {
    return static_cast<T*>(n);
  } else {
    return nullptr;
  }
}

// GetAsOrNull implemented simply with dynamic_cast.  This is used for
// intermediate nodes (such as ASTExpression).
template <typename T, typename MaybeConstRoot, EnableIfNotConcrete<T> = 0>
inline T* GetAsOrNullImpl(MaybeConstRoot* r) {
  // Note, if this proves too slow, it could be implemented with ancestor enums
  // sets.
  return dynamic_cast<T*>(r);
}

}  // namespace parse_tree_internal

template <typename NodeType>
inline const NodeType* ASTNode::GetAsOrNull() const {
  static_assert(std::is_base_of<ASTNode, NodeType>::value,
                "NodeType must be a member of the ASTNode class hierarchy");
  return parse_tree_internal::GetAsOrNullImpl<const NodeType, const ASTNode>(
      this);
}

template <typename NodeType>
inline NodeType* ASTNode::GetAsOrNull() {
  static_assert(std::is_base_of<ASTNode, NodeType>::value,
                "NodeType must be a member of the ASTNode class hierarchy");
  return parse_tree_internal::GetAsOrNullImpl<NodeType, ASTNode>(this);
}

}  // namespace zetasql
#endif  // ZETASQL_PARSER_PARSE_TREE_H_