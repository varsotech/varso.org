package rssfeed

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/luminancetech/varso/src/services/app/internal/ent"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/organization"
	"github.com/luminancetech/varso/src/services/app/internal/ent/build/rssfeed"
)

func Upsert(ctx context.Context, orgUuid uuid.UUID, rssFeedUrl, contentWhitelistRegex, htmlPaywallRegex, titleTrimRight string, rssFeedRank float32, maxFetchIntervalMin int64) error {
	if rssFeedUrl == "" {
		return fmt.Errorf("rss feed url cannot be empty")
	}

	err := ent.Database.RSSFeed.Create().
		SetOrganizationID(orgUuid).
		SetRssFeedURL(rssFeedUrl).
		SetRssFeedRank(rssFeedRank).
		SetContentWhitelistRegex(contentWhitelistRegex).
		SetHTMLPaywallRegex(htmlPaywallRegex).
		SetTitleTrimRight(titleTrimRight).
		SetMaxFetchIntervalMin(maxFetchIntervalMin).
		OnConflictColumns(rssfeed.FieldRssFeedURL).
		UpdateNewValues().
		UpdateContentWhitelistRegex().
		UpdateHTMLPaywallRegex().
		Exec(ctx)
	if err != nil {
		return err
	}

	updater := ent.Database.RSSFeed.Update().Where(rssfeed.RssFeedURL(rssFeedUrl))

	if rssFeedRank == 0 {
		updater.ClearRssFeedRank()
	}

	if contentWhitelistRegex == "" {
		updater.ClearContentWhitelistRegex()
	}

	if htmlPaywallRegex == "" {
		updater.ClearHTMLPaywallRegex()
	}

	if titleTrimRight == "" {
		updater.ClearTitleTrimRight()
	}

	if maxFetchIntervalMin == 0 {
		updater.ClearMaxFetchIntervalMin()
	}

	return updater.Exec(ctx)
}

func UpdateUpdateTime(ctx context.Context, feed *build.RSSFeed) error {
	return ent.Database.RSSFeed.UpdateOne(feed).SetUpdateTime(time.Now()).Exec(ctx)
}

func GetByOrganization(ctx context.Context, orgs ...*build.Organization) ([]*build.RSSFeed, error) {
	var orgUUIDs []uuid.UUID
	for _, org := range orgs {
		orgUUIDs = append(orgUUIDs, org.ID)
	}

	feeds, err := ent.Database.RSSFeed.
		Query().
		Where(
			rssfeed.HasOrganizationWith(organization.IDIn(orgUUIDs...)),
		).All(ctx)
	if err != nil {
		return nil, err
	}

	return feeds, nil
}
