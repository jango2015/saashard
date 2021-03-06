// The MIT License (MIT)

// Copyright (c) 2016 Jerry Bai

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package route

import (
	"github.com/berkaroad/saashard/sqlparser"
	"github.com/berkaroad/saashard/utils"
)

func (r *Router) buildCreateTablePlan(statement *sqlparser.CreateTable) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	statement.Table.Qualifier = nil
	if schemaConfig.ShardEnabled() {
	}
	hint := ReadHint(&statement.Comments)

	plan := new(normalPlan)
	if len(hint.Nodes) > 0 {
		plan.nodeNames = utils.StringCollectionIntersection(hint.Nodes, schemaConfig.Nodes)
	} else {
		plan.nodeNames = schemaConfig.Nodes
	}
	plan.Statement = statement
	return plan, nil
}

func (r *Router) buildCreateIndexPlan(statement *sqlparser.CreateIndex) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	statement.Table.Qualifier = nil
	if schemaConfig.ShardEnabled() {
	}
	hint := ReadHint(&statement.Comments)

	plan := new(normalPlan)
	if len(hint.Nodes) > 0 {
		plan.nodeNames = utils.StringCollectionIntersection(hint.Nodes, schemaConfig.Nodes)
	} else {
		plan.nodeNames = schemaConfig.Nodes
	}
	plan.Statement = statement
	return plan, nil
}

func (r *Router) buildRenameTablePlan(statement *sqlparser.RenameTable) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	statement.OldName.Qualifier = nil
	statement.NewName.Qualifier = nil
	if schemaConfig.ShardEnabled() {
	}
	hint := ReadHint(&statement.Comments)

	plan := new(normalPlan)
	if len(hint.Nodes) > 0 {
		plan.nodeNames = utils.StringCollectionIntersection(hint.Nodes, schemaConfig.Nodes)
	} else {
		plan.nodeNames = schemaConfig.Nodes
	}
	plan.Statement = statement
	return plan, nil
}

func (r *Router) buildDropTablePlan(statement *sqlparser.DropTable) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	statement.Name.Qualifier = nil
	if schemaConfig.ShardEnabled() {
	}
	hint := ReadHint(&statement.Comments)

	plan := new(normalPlan)
	if len(hint.Nodes) > 0 {
		plan.nodeNames = utils.StringCollectionIntersection(hint.Nodes, schemaConfig.Nodes)
	} else {
		plan.nodeNames = schemaConfig.Nodes
	}
	plan.Statement = statement
	return plan, nil
}

func (r *Router) buildDropIndexPlan(statement *sqlparser.DropIndex) (*normalPlan, error) {
	schemaConfig := r.Schemas[r.SchemaName]
	statement.Table.Qualifier = nil
	if schemaConfig.ShardEnabled() {
	}
	hint := ReadHint(&statement.Comments)

	plan := new(normalPlan)
	if len(hint.Nodes) > 0 {
		plan.nodeNames = utils.StringCollectionIntersection(hint.Nodes, schemaConfig.Nodes)
	} else {
		plan.nodeNames = schemaConfig.Nodes
	}
	plan.Statement = statement
	return plan, nil
}
