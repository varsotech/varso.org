// Code generated by ent, DO NOT EDIT.

package newsitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldUpdateTime, v))
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v bool) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldDeleted, v))
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldDeleteTime, v))
}

// RssGUID applies equality check predicate on the "rss_guid" field. It's identical to RssGUIDEQ.
func RssGUID(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldRssGUID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldDescription, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldContent, v))
}

// Link applies equality check predicate on the "link" field. It's identical to LinkEQ.
func Link(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldLink, v))
}

// ItemPublishTime applies equality check predicate on the "item_publish_time" field. It's identical to ItemPublishTimeEQ.
func ItemPublishTime(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldItemPublishTime, v))
}

// ItemUpdateTime applies equality check predicate on the "item_update_time" field. It's identical to ItemUpdateTimeEQ.
func ItemUpdateTime(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldItemUpdateTime, v))
}

// ImageURL applies equality check predicate on the "image_url" field. It's identical to ImageURLEQ.
func ImageURL(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldImageURL, v))
}

// ImageTitle applies equality check predicate on the "image_title" field. It's identical to ImageTitleEQ.
func ImageTitle(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldImageTitle, v))
}

// Blur applies equality check predicate on the "blur" field. It's identical to BlurEQ.
func Blur(v bool) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldBlur, v))
}

// Paywalled applies equality check predicate on the "paywalled" field. It's identical to PaywalledEQ.
func Paywalled(v bool) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldPaywalled, v))
}

// Rank applies equality check predicate on the "rank" field. It's identical to RankEQ.
func Rank(v int64) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldRank, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldUpdateTime, v))
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v bool) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldDeleted, v))
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v bool) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldDeleted, v))
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldDeleteTime, v))
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldDeleteTime, v))
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldDeleteTime, vs...))
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldDeleteTime, vs...))
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldDeleteTime, v))
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldDeleteTime, v))
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldDeleteTime, v))
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldDeleteTime, v))
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIsNull(FieldDeleteTime))
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotNull(FieldDeleteTime))
}

// RssGUIDEQ applies the EQ predicate on the "rss_guid" field.
func RssGUIDEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldRssGUID, v))
}

// RssGUIDNEQ applies the NEQ predicate on the "rss_guid" field.
func RssGUIDNEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldRssGUID, v))
}

// RssGUIDIn applies the In predicate on the "rss_guid" field.
func RssGUIDIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldRssGUID, vs...))
}

// RssGUIDNotIn applies the NotIn predicate on the "rss_guid" field.
func RssGUIDNotIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldRssGUID, vs...))
}

// RssGUIDGT applies the GT predicate on the "rss_guid" field.
func RssGUIDGT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldRssGUID, v))
}

// RssGUIDGTE applies the GTE predicate on the "rss_guid" field.
func RssGUIDGTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldRssGUID, v))
}

// RssGUIDLT applies the LT predicate on the "rss_guid" field.
func RssGUIDLT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldRssGUID, v))
}

// RssGUIDLTE applies the LTE predicate on the "rss_guid" field.
func RssGUIDLTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldRssGUID, v))
}

// RssGUIDContains applies the Contains predicate on the "rss_guid" field.
func RssGUIDContains(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContains(FieldRssGUID, v))
}

// RssGUIDHasPrefix applies the HasPrefix predicate on the "rss_guid" field.
func RssGUIDHasPrefix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasPrefix(FieldRssGUID, v))
}

// RssGUIDHasSuffix applies the HasSuffix predicate on the "rss_guid" field.
func RssGUIDHasSuffix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasSuffix(FieldRssGUID, v))
}

// RssGUIDEqualFold applies the EqualFold predicate on the "rss_guid" field.
func RssGUIDEqualFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEqualFold(FieldRssGUID, v))
}

// RssGUIDContainsFold applies the ContainsFold predicate on the "rss_guid" field.
func RssGUIDContainsFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContainsFold(FieldRssGUID, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContainsFold(FieldDescription, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContainsFold(FieldContent, v))
}

// LinkEQ applies the EQ predicate on the "link" field.
func LinkEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldLink, v))
}

// LinkNEQ applies the NEQ predicate on the "link" field.
func LinkNEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldLink, v))
}

// LinkIn applies the In predicate on the "link" field.
func LinkIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldLink, vs...))
}

// LinkNotIn applies the NotIn predicate on the "link" field.
func LinkNotIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldLink, vs...))
}

// LinkGT applies the GT predicate on the "link" field.
func LinkGT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldLink, v))
}

// LinkGTE applies the GTE predicate on the "link" field.
func LinkGTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldLink, v))
}

// LinkLT applies the LT predicate on the "link" field.
func LinkLT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldLink, v))
}

// LinkLTE applies the LTE predicate on the "link" field.
func LinkLTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldLink, v))
}

// LinkContains applies the Contains predicate on the "link" field.
func LinkContains(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContains(FieldLink, v))
}

// LinkHasPrefix applies the HasPrefix predicate on the "link" field.
func LinkHasPrefix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasPrefix(FieldLink, v))
}

// LinkHasSuffix applies the HasSuffix predicate on the "link" field.
func LinkHasSuffix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasSuffix(FieldLink, v))
}

// LinkEqualFold applies the EqualFold predicate on the "link" field.
func LinkEqualFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEqualFold(FieldLink, v))
}

// LinkContainsFold applies the ContainsFold predicate on the "link" field.
func LinkContainsFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContainsFold(FieldLink, v))
}

// ItemPublishTimeEQ applies the EQ predicate on the "item_publish_time" field.
func ItemPublishTimeEQ(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldItemPublishTime, v))
}

// ItemPublishTimeNEQ applies the NEQ predicate on the "item_publish_time" field.
func ItemPublishTimeNEQ(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldItemPublishTime, v))
}

// ItemPublishTimeIn applies the In predicate on the "item_publish_time" field.
func ItemPublishTimeIn(vs ...time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldItemPublishTime, vs...))
}

// ItemPublishTimeNotIn applies the NotIn predicate on the "item_publish_time" field.
func ItemPublishTimeNotIn(vs ...time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldItemPublishTime, vs...))
}

// ItemPublishTimeGT applies the GT predicate on the "item_publish_time" field.
func ItemPublishTimeGT(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldItemPublishTime, v))
}

// ItemPublishTimeGTE applies the GTE predicate on the "item_publish_time" field.
func ItemPublishTimeGTE(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldItemPublishTime, v))
}

// ItemPublishTimeLT applies the LT predicate on the "item_publish_time" field.
func ItemPublishTimeLT(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldItemPublishTime, v))
}

// ItemPublishTimeLTE applies the LTE predicate on the "item_publish_time" field.
func ItemPublishTimeLTE(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldItemPublishTime, v))
}

// ItemPublishTimeIsNil applies the IsNil predicate on the "item_publish_time" field.
func ItemPublishTimeIsNil() predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIsNull(FieldItemPublishTime))
}

// ItemPublishTimeNotNil applies the NotNil predicate on the "item_publish_time" field.
func ItemPublishTimeNotNil() predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotNull(FieldItemPublishTime))
}

// ItemUpdateTimeEQ applies the EQ predicate on the "item_update_time" field.
func ItemUpdateTimeEQ(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldItemUpdateTime, v))
}

// ItemUpdateTimeNEQ applies the NEQ predicate on the "item_update_time" field.
func ItemUpdateTimeNEQ(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldItemUpdateTime, v))
}

// ItemUpdateTimeIn applies the In predicate on the "item_update_time" field.
func ItemUpdateTimeIn(vs ...time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldItemUpdateTime, vs...))
}

// ItemUpdateTimeNotIn applies the NotIn predicate on the "item_update_time" field.
func ItemUpdateTimeNotIn(vs ...time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldItemUpdateTime, vs...))
}

// ItemUpdateTimeGT applies the GT predicate on the "item_update_time" field.
func ItemUpdateTimeGT(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldItemUpdateTime, v))
}

// ItemUpdateTimeGTE applies the GTE predicate on the "item_update_time" field.
func ItemUpdateTimeGTE(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldItemUpdateTime, v))
}

// ItemUpdateTimeLT applies the LT predicate on the "item_update_time" field.
func ItemUpdateTimeLT(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldItemUpdateTime, v))
}

// ItemUpdateTimeLTE applies the LTE predicate on the "item_update_time" field.
func ItemUpdateTimeLTE(v time.Time) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldItemUpdateTime, v))
}

// ItemUpdateTimeIsNil applies the IsNil predicate on the "item_update_time" field.
func ItemUpdateTimeIsNil() predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIsNull(FieldItemUpdateTime))
}

// ItemUpdateTimeNotNil applies the NotNil predicate on the "item_update_time" field.
func ItemUpdateTimeNotNil() predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotNull(FieldItemUpdateTime))
}

// ImageURLEQ applies the EQ predicate on the "image_url" field.
func ImageURLEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldImageURL, v))
}

// ImageURLNEQ applies the NEQ predicate on the "image_url" field.
func ImageURLNEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldImageURL, v))
}

// ImageURLIn applies the In predicate on the "image_url" field.
func ImageURLIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldImageURL, vs...))
}

// ImageURLNotIn applies the NotIn predicate on the "image_url" field.
func ImageURLNotIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldImageURL, vs...))
}

// ImageURLGT applies the GT predicate on the "image_url" field.
func ImageURLGT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldImageURL, v))
}

// ImageURLGTE applies the GTE predicate on the "image_url" field.
func ImageURLGTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldImageURL, v))
}

// ImageURLLT applies the LT predicate on the "image_url" field.
func ImageURLLT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldImageURL, v))
}

// ImageURLLTE applies the LTE predicate on the "image_url" field.
func ImageURLLTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldImageURL, v))
}

// ImageURLContains applies the Contains predicate on the "image_url" field.
func ImageURLContains(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContains(FieldImageURL, v))
}

// ImageURLHasPrefix applies the HasPrefix predicate on the "image_url" field.
func ImageURLHasPrefix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasPrefix(FieldImageURL, v))
}

// ImageURLHasSuffix applies the HasSuffix predicate on the "image_url" field.
func ImageURLHasSuffix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasSuffix(FieldImageURL, v))
}

// ImageURLEqualFold applies the EqualFold predicate on the "image_url" field.
func ImageURLEqualFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEqualFold(FieldImageURL, v))
}

// ImageURLContainsFold applies the ContainsFold predicate on the "image_url" field.
func ImageURLContainsFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContainsFold(FieldImageURL, v))
}

// ImageTitleEQ applies the EQ predicate on the "image_title" field.
func ImageTitleEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldImageTitle, v))
}

// ImageTitleNEQ applies the NEQ predicate on the "image_title" field.
func ImageTitleNEQ(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldImageTitle, v))
}

// ImageTitleIn applies the In predicate on the "image_title" field.
func ImageTitleIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldImageTitle, vs...))
}

// ImageTitleNotIn applies the NotIn predicate on the "image_title" field.
func ImageTitleNotIn(vs ...string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldImageTitle, vs...))
}

// ImageTitleGT applies the GT predicate on the "image_title" field.
func ImageTitleGT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldImageTitle, v))
}

// ImageTitleGTE applies the GTE predicate on the "image_title" field.
func ImageTitleGTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldImageTitle, v))
}

// ImageTitleLT applies the LT predicate on the "image_title" field.
func ImageTitleLT(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldImageTitle, v))
}

// ImageTitleLTE applies the LTE predicate on the "image_title" field.
func ImageTitleLTE(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldImageTitle, v))
}

// ImageTitleContains applies the Contains predicate on the "image_title" field.
func ImageTitleContains(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContains(FieldImageTitle, v))
}

// ImageTitleHasPrefix applies the HasPrefix predicate on the "image_title" field.
func ImageTitleHasPrefix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasPrefix(FieldImageTitle, v))
}

// ImageTitleHasSuffix applies the HasSuffix predicate on the "image_title" field.
func ImageTitleHasSuffix(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldHasSuffix(FieldImageTitle, v))
}

// ImageTitleEqualFold applies the EqualFold predicate on the "image_title" field.
func ImageTitleEqualFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEqualFold(FieldImageTitle, v))
}

// ImageTitleContainsFold applies the ContainsFold predicate on the "image_title" field.
func ImageTitleContainsFold(v string) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldContainsFold(FieldImageTitle, v))
}

// BlurEQ applies the EQ predicate on the "blur" field.
func BlurEQ(v bool) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldBlur, v))
}

// BlurNEQ applies the NEQ predicate on the "blur" field.
func BlurNEQ(v bool) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldBlur, v))
}

// PaywalledEQ applies the EQ predicate on the "paywalled" field.
func PaywalledEQ(v bool) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldPaywalled, v))
}

// PaywalledNEQ applies the NEQ predicate on the "paywalled" field.
func PaywalledNEQ(v bool) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldPaywalled, v))
}

// RankEQ applies the EQ predicate on the "rank" field.
func RankEQ(v int64) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldEQ(FieldRank, v))
}

// RankNEQ applies the NEQ predicate on the "rank" field.
func RankNEQ(v int64) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNEQ(FieldRank, v))
}

// RankIn applies the In predicate on the "rank" field.
func RankIn(vs ...int64) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldIn(FieldRank, vs...))
}

// RankNotIn applies the NotIn predicate on the "rank" field.
func RankNotIn(vs ...int64) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldNotIn(FieldRank, vs...))
}

// RankGT applies the GT predicate on the "rank" field.
func RankGT(v int64) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGT(FieldRank, v))
}

// RankGTE applies the GTE predicate on the "rank" field.
func RankGTE(v int64) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldGTE(FieldRank, v))
}

// RankLT applies the LT predicate on the "rank" field.
func RankLT(v int64) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLT(FieldRank, v))
}

// RankLTE applies the LTE predicate on the "rank" field.
func RankLTE(v int64) predicate.NewsItem {
	return predicate.NewsItem(sql.FieldLTE(FieldRank, v))
}

// HasAuthors applies the HasEdge predicate on the "authors" edge.
func HasAuthors() predicate.NewsItem {
	return predicate.NewsItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, AuthorsTable, AuthorsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorsWith applies the HasEdge predicate on the "authors" edge with a given conditions (other predicates).
func HasAuthorsWith(preds ...predicate.RSSAuthor) predicate.NewsItem {
	return predicate.NewsItem(func(s *sql.Selector) {
		step := newAuthorsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFeed applies the HasEdge predicate on the "feed" edge.
func HasFeed() predicate.NewsItem {
	return predicate.NewsItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FeedTable, FeedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFeedWith applies the HasEdge predicate on the "feed" edge with a given conditions (other predicates).
func HasFeedWith(preds ...predicate.RSSFeed) predicate.NewsItem {
	return predicate.NewsItem(func(s *sql.Selector) {
		step := newFeedStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NewsItem) predicate.NewsItem {
	return predicate.NewsItem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NewsItem) predicate.NewsItem {
	return predicate.NewsItem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NewsItem) predicate.NewsItem {
	return predicate.NewsItem(sql.NotPredicates(p))
}
