package rank

type PublishDateRanker struct {
}

func NewPublishDateRanker() *PublishDateRanker {
	return &PublishDateRanker{}
}

func (r *PublishDateRanker) Rank(item *RankedItem) int64 {
	return item.PublishDate.Unix()
}
