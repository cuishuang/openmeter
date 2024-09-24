// Code generated by ent, DO NOT EDIT.

package billingworkflowconfig

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/openmeterio/openmeter/openmeter/billing"
	"github.com/openmeterio/openmeter/openmeter/ent/db/predicate"
	"github.com/openmeterio/openmeter/pkg/timezone"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldContainsFold(FieldID, id))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldNamespace, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldDeletedAt, v))
}

// Timezone applies equality check predicate on the "timezone" field. It's identical to TimezoneEQ.
func Timezone(v timezone.Timezone) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldTimezone, vc))
}

// ItemCollectionPeriodSeconds applies equality check predicate on the "item_collection_period_seconds" field. It's identical to ItemCollectionPeriodSecondsEQ.
func ItemCollectionPeriodSeconds(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldItemCollectionPeriodSeconds, v))
}

// InvoiceAutoAdvance applies equality check predicate on the "invoice_auto_advance" field. It's identical to InvoiceAutoAdvanceEQ.
func InvoiceAutoAdvance(v bool) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceAutoAdvance, v))
}

// InvoiceDraftPeriodSeconds applies equality check predicate on the "invoice_draft_period_seconds" field. It's identical to InvoiceDraftPeriodSecondsEQ.
func InvoiceDraftPeriodSeconds(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceDraftPeriodSeconds, v))
}

// InvoiceDueAfterSeconds applies equality check predicate on the "invoice_due_after_seconds" field. It's identical to InvoiceDueAfterSecondsEQ.
func InvoiceDueAfterSeconds(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceDueAfterSeconds, v))
}

// InvoiceItemPerSubject applies equality check predicate on the "invoice_item_per_subject" field. It's identical to InvoiceItemPerSubjectEQ.
func InvoiceItemPerSubject(v bool) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceItemPerSubject, v))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldContainsFold(FieldNamespace, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotNull(FieldDeletedAt))
}

// TimezoneEQ applies the EQ predicate on the "timezone" field.
func TimezoneEQ(v timezone.Timezone) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldTimezone, vc))
}

// TimezoneNEQ applies the NEQ predicate on the "timezone" field.
func TimezoneNEQ(v timezone.Timezone) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldTimezone, vc))
}

// TimezoneIn applies the In predicate on the "timezone" field.
func TimezoneIn(vs ...timezone.Timezone) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldTimezone, v...))
}

// TimezoneNotIn applies the NotIn predicate on the "timezone" field.
func TimezoneNotIn(vs ...timezone.Timezone) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldTimezone, v...))
}

// TimezoneGT applies the GT predicate on the "timezone" field.
func TimezoneGT(v timezone.Timezone) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldTimezone, vc))
}

// TimezoneGTE applies the GTE predicate on the "timezone" field.
func TimezoneGTE(v timezone.Timezone) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldTimezone, vc))
}

// TimezoneLT applies the LT predicate on the "timezone" field.
func TimezoneLT(v timezone.Timezone) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldTimezone, vc))
}

// TimezoneLTE applies the LTE predicate on the "timezone" field.
func TimezoneLTE(v timezone.Timezone) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldTimezone, vc))
}

// TimezoneContains applies the Contains predicate on the "timezone" field.
func TimezoneContains(v timezone.Timezone) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldContains(FieldTimezone, vc))
}

// TimezoneHasPrefix applies the HasPrefix predicate on the "timezone" field.
func TimezoneHasPrefix(v timezone.Timezone) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldHasPrefix(FieldTimezone, vc))
}

// TimezoneHasSuffix applies the HasSuffix predicate on the "timezone" field.
func TimezoneHasSuffix(v timezone.Timezone) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldHasSuffix(FieldTimezone, vc))
}

// TimezoneIsNil applies the IsNil predicate on the "timezone" field.
func TimezoneIsNil() predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIsNull(FieldTimezone))
}

// TimezoneNotNil applies the NotNil predicate on the "timezone" field.
func TimezoneNotNil() predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotNull(FieldTimezone))
}

// TimezoneEqualFold applies the EqualFold predicate on the "timezone" field.
func TimezoneEqualFold(v timezone.Timezone) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldEqualFold(FieldTimezone, vc))
}

// TimezoneContainsFold applies the ContainsFold predicate on the "timezone" field.
func TimezoneContainsFold(v timezone.Timezone) predicate.BillingWorkflowConfig {
	vc := string(v)
	return predicate.BillingWorkflowConfig(sql.FieldContainsFold(FieldTimezone, vc))
}

// CollectionAlignmentEQ applies the EQ predicate on the "collection_alignment" field.
func CollectionAlignmentEQ(v billing.AlignmentKind) predicate.BillingWorkflowConfig {
	vc := v
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldCollectionAlignment, vc))
}

// CollectionAlignmentNEQ applies the NEQ predicate on the "collection_alignment" field.
func CollectionAlignmentNEQ(v billing.AlignmentKind) predicate.BillingWorkflowConfig {
	vc := v
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldCollectionAlignment, vc))
}

// CollectionAlignmentIn applies the In predicate on the "collection_alignment" field.
func CollectionAlignmentIn(vs ...billing.AlignmentKind) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldCollectionAlignment, v...))
}

// CollectionAlignmentNotIn applies the NotIn predicate on the "collection_alignment" field.
func CollectionAlignmentNotIn(vs ...billing.AlignmentKind) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldCollectionAlignment, v...))
}

// ItemCollectionPeriodSecondsEQ applies the EQ predicate on the "item_collection_period_seconds" field.
func ItemCollectionPeriodSecondsEQ(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldItemCollectionPeriodSeconds, v))
}

// ItemCollectionPeriodSecondsNEQ applies the NEQ predicate on the "item_collection_period_seconds" field.
func ItemCollectionPeriodSecondsNEQ(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldItemCollectionPeriodSeconds, v))
}

// ItemCollectionPeriodSecondsIn applies the In predicate on the "item_collection_period_seconds" field.
func ItemCollectionPeriodSecondsIn(vs ...int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldItemCollectionPeriodSeconds, vs...))
}

// ItemCollectionPeriodSecondsNotIn applies the NotIn predicate on the "item_collection_period_seconds" field.
func ItemCollectionPeriodSecondsNotIn(vs ...int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldItemCollectionPeriodSeconds, vs...))
}

// ItemCollectionPeriodSecondsGT applies the GT predicate on the "item_collection_period_seconds" field.
func ItemCollectionPeriodSecondsGT(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldItemCollectionPeriodSeconds, v))
}

// ItemCollectionPeriodSecondsGTE applies the GTE predicate on the "item_collection_period_seconds" field.
func ItemCollectionPeriodSecondsGTE(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldItemCollectionPeriodSeconds, v))
}

// ItemCollectionPeriodSecondsLT applies the LT predicate on the "item_collection_period_seconds" field.
func ItemCollectionPeriodSecondsLT(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldItemCollectionPeriodSeconds, v))
}

// ItemCollectionPeriodSecondsLTE applies the LTE predicate on the "item_collection_period_seconds" field.
func ItemCollectionPeriodSecondsLTE(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldItemCollectionPeriodSeconds, v))
}

// InvoiceAutoAdvanceEQ applies the EQ predicate on the "invoice_auto_advance" field.
func InvoiceAutoAdvanceEQ(v bool) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceAutoAdvance, v))
}

// InvoiceAutoAdvanceNEQ applies the NEQ predicate on the "invoice_auto_advance" field.
func InvoiceAutoAdvanceNEQ(v bool) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldInvoiceAutoAdvance, v))
}

// InvoiceDraftPeriodSecondsEQ applies the EQ predicate on the "invoice_draft_period_seconds" field.
func InvoiceDraftPeriodSecondsEQ(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceDraftPeriodSeconds, v))
}

// InvoiceDraftPeriodSecondsNEQ applies the NEQ predicate on the "invoice_draft_period_seconds" field.
func InvoiceDraftPeriodSecondsNEQ(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldInvoiceDraftPeriodSeconds, v))
}

// InvoiceDraftPeriodSecondsIn applies the In predicate on the "invoice_draft_period_seconds" field.
func InvoiceDraftPeriodSecondsIn(vs ...int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldInvoiceDraftPeriodSeconds, vs...))
}

// InvoiceDraftPeriodSecondsNotIn applies the NotIn predicate on the "invoice_draft_period_seconds" field.
func InvoiceDraftPeriodSecondsNotIn(vs ...int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldInvoiceDraftPeriodSeconds, vs...))
}

// InvoiceDraftPeriodSecondsGT applies the GT predicate on the "invoice_draft_period_seconds" field.
func InvoiceDraftPeriodSecondsGT(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldInvoiceDraftPeriodSeconds, v))
}

// InvoiceDraftPeriodSecondsGTE applies the GTE predicate on the "invoice_draft_period_seconds" field.
func InvoiceDraftPeriodSecondsGTE(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldInvoiceDraftPeriodSeconds, v))
}

// InvoiceDraftPeriodSecondsLT applies the LT predicate on the "invoice_draft_period_seconds" field.
func InvoiceDraftPeriodSecondsLT(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldInvoiceDraftPeriodSeconds, v))
}

// InvoiceDraftPeriodSecondsLTE applies the LTE predicate on the "invoice_draft_period_seconds" field.
func InvoiceDraftPeriodSecondsLTE(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldInvoiceDraftPeriodSeconds, v))
}

// InvoiceDueAfterSecondsEQ applies the EQ predicate on the "invoice_due_after_seconds" field.
func InvoiceDueAfterSecondsEQ(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceDueAfterSeconds, v))
}

// InvoiceDueAfterSecondsNEQ applies the NEQ predicate on the "invoice_due_after_seconds" field.
func InvoiceDueAfterSecondsNEQ(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldInvoiceDueAfterSeconds, v))
}

// InvoiceDueAfterSecondsIn applies the In predicate on the "invoice_due_after_seconds" field.
func InvoiceDueAfterSecondsIn(vs ...int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldInvoiceDueAfterSeconds, vs...))
}

// InvoiceDueAfterSecondsNotIn applies the NotIn predicate on the "invoice_due_after_seconds" field.
func InvoiceDueAfterSecondsNotIn(vs ...int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldInvoiceDueAfterSeconds, vs...))
}

// InvoiceDueAfterSecondsGT applies the GT predicate on the "invoice_due_after_seconds" field.
func InvoiceDueAfterSecondsGT(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGT(FieldInvoiceDueAfterSeconds, v))
}

// InvoiceDueAfterSecondsGTE applies the GTE predicate on the "invoice_due_after_seconds" field.
func InvoiceDueAfterSecondsGTE(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldGTE(FieldInvoiceDueAfterSeconds, v))
}

// InvoiceDueAfterSecondsLT applies the LT predicate on the "invoice_due_after_seconds" field.
func InvoiceDueAfterSecondsLT(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLT(FieldInvoiceDueAfterSeconds, v))
}

// InvoiceDueAfterSecondsLTE applies the LTE predicate on the "invoice_due_after_seconds" field.
func InvoiceDueAfterSecondsLTE(v int64) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldLTE(FieldInvoiceDueAfterSeconds, v))
}

// InvoiceCollectionMethodEQ applies the EQ predicate on the "invoice_collection_method" field.
func InvoiceCollectionMethodEQ(v billing.CollectionMethod) predicate.BillingWorkflowConfig {
	vc := v
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceCollectionMethod, vc))
}

// InvoiceCollectionMethodNEQ applies the NEQ predicate on the "invoice_collection_method" field.
func InvoiceCollectionMethodNEQ(v billing.CollectionMethod) predicate.BillingWorkflowConfig {
	vc := v
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldInvoiceCollectionMethod, vc))
}

// InvoiceCollectionMethodIn applies the In predicate on the "invoice_collection_method" field.
func InvoiceCollectionMethodIn(vs ...billing.CollectionMethod) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldInvoiceCollectionMethod, v...))
}

// InvoiceCollectionMethodNotIn applies the NotIn predicate on the "invoice_collection_method" field.
func InvoiceCollectionMethodNotIn(vs ...billing.CollectionMethod) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldInvoiceCollectionMethod, v...))
}

// InvoiceItemResolutionEQ applies the EQ predicate on the "invoice_item_resolution" field.
func InvoiceItemResolutionEQ(v billing.GranularityResolution) predicate.BillingWorkflowConfig {
	vc := v
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceItemResolution, vc))
}

// InvoiceItemResolutionNEQ applies the NEQ predicate on the "invoice_item_resolution" field.
func InvoiceItemResolutionNEQ(v billing.GranularityResolution) predicate.BillingWorkflowConfig {
	vc := v
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldInvoiceItemResolution, vc))
}

// InvoiceItemResolutionIn applies the In predicate on the "invoice_item_resolution" field.
func InvoiceItemResolutionIn(vs ...billing.GranularityResolution) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingWorkflowConfig(sql.FieldIn(FieldInvoiceItemResolution, v...))
}

// InvoiceItemResolutionNotIn applies the NotIn predicate on the "invoice_item_resolution" field.
func InvoiceItemResolutionNotIn(vs ...billing.GranularityResolution) predicate.BillingWorkflowConfig {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BillingWorkflowConfig(sql.FieldNotIn(FieldInvoiceItemResolution, v...))
}

// InvoiceItemPerSubjectEQ applies the EQ predicate on the "invoice_item_per_subject" field.
func InvoiceItemPerSubjectEQ(v bool) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldEQ(FieldInvoiceItemPerSubject, v))
}

// InvoiceItemPerSubjectNEQ applies the NEQ predicate on the "invoice_item_per_subject" field.
func InvoiceItemPerSubjectNEQ(v bool) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.FieldNEQ(FieldInvoiceItemPerSubject, v))
}

// HasBillingInvoices applies the HasEdge predicate on the "billing_invoices" edge.
func HasBillingInvoices() predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BillingInvoicesTable, BillingInvoicesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingInvoicesWith applies the HasEdge predicate on the "billing_invoices" edge with a given conditions (other predicates).
func HasBillingInvoicesWith(preds ...predicate.BillingInvoice) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(func(s *sql.Selector) {
		step := newBillingInvoicesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBillingProfile applies the HasEdge predicate on the "billing_profile" edge.
func HasBillingProfile() predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BillingProfileTable, BillingProfileColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillingProfileWith applies the HasEdge predicate on the "billing_profile" edge with a given conditions (other predicates).
func HasBillingProfileWith(preds ...predicate.BillingProfile) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(func(s *sql.Selector) {
		step := newBillingProfileStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BillingWorkflowConfig) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BillingWorkflowConfig) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BillingWorkflowConfig) predicate.BillingWorkflowConfig {
	return predicate.BillingWorkflowConfig(sql.NotPredicates(p))
}
