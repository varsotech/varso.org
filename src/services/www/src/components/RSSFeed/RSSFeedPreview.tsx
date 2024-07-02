import {
  Organization,
  RSSItem as RSSItemModel,
} from "@luminancetech/varso/src/app/appbase";
import RSSItemPreview from "./RSSItemPreview";
import * as Styled from "./RSSFeedPreview.style";
import TabMenu from "components/TabMenu/TabMenu";

type RSSFeedProps = {
  items: RSSItemModel[];
  organizations: { [key: string]: Organization };
  selectedFeedTab: string;
  setSelectedFeedTab: (v: string) => void;
  isFetching?: boolean;
};

function RSSFeedPreview({
  items,
  organizations,
  selectedFeedTab,
  setSelectedFeedTab,
  isFetching,
}: RSSFeedProps) {
  return (
    <Styled.RSSFeedPreview>
      <TabMenu
        options={[
          {
            id: "foryou",
            text: "For You",
          },
          {
            id: "latest",
            text: "Latest",
          },
        ]}
        selected={selectedFeedTab}
        onChange={setSelectedFeedTab}
      />
      {isFetching ? (
        <div style={{ height: 800 }} />
      ) : (
        <Styled.RSSFeedItems>
          {items.map((item: RSSItemModel, index: number) => {
            return (
              <RSSItemPreview
                key={item.guid}
                item={item}
                organization={organizations[item.organizationUuid]}
                index={index + 2}
              />
            );
          })}
        </Styled.RSSFeedItems>
      )}
    </Styled.RSSFeedPreview>
  );
}

export default RSSFeedPreview;
