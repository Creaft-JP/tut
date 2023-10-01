// Code generated by ent, DO NOT EDIT.

package remote

import (
	"entgo.io/ent/dialect/sql"
	"github.com/Creaft-JP/tit/db/local/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Remote {
	return predicate.Remote(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Remote {
	return predicate.Remote(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Remote {
	return predicate.Remote(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Remote {
	return predicate.Remote(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Remote {
	return predicate.Remote(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Remote {
	return predicate.Remote(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Remote {
	return predicate.Remote(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Remote {
	return predicate.Remote(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Remote {
	return predicate.Remote(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Remote {
	return predicate.Remote(sql.FieldEQ(FieldName, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Remote {
	return predicate.Remote(sql.FieldEQ(FieldURL, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Remote {
	return predicate.Remote(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Remote {
	return predicate.Remote(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Remote {
	return predicate.Remote(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Remote {
	return predicate.Remote(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Remote {
	return predicate.Remote(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Remote {
	return predicate.Remote(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Remote {
	return predicate.Remote(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Remote {
	return predicate.Remote(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Remote {
	return predicate.Remote(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Remote {
	return predicate.Remote(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Remote {
	return predicate.Remote(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Remote {
	return predicate.Remote(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Remote {
	return predicate.Remote(sql.FieldContainsFold(FieldName, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Remote {
	return predicate.Remote(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Remote {
	return predicate.Remote(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Remote {
	return predicate.Remote(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Remote {
	return predicate.Remote(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Remote {
	return predicate.Remote(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Remote {
	return predicate.Remote(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Remote {
	return predicate.Remote(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Remote {
	return predicate.Remote(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Remote {
	return predicate.Remote(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Remote {
	return predicate.Remote(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Remote {
	return predicate.Remote(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Remote {
	return predicate.Remote(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Remote {
	return predicate.Remote(sql.FieldContainsFold(FieldURL, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Remote) predicate.Remote {
	return predicate.Remote(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Remote) predicate.Remote {
	return predicate.Remote(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Remote) predicate.Remote {
	return predicate.Remote(func(s *sql.Selector) {
		p(s.Not())
	})
}
