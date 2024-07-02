import {
  Organization,
  RSSItem as RSSItemModel,
} from "@luminancetech/varso/src/app/appbase";
import * as Styled from "./RSSItemPreview.style";
import RSSItemMetadata from "./RSSItemMetadata";
import RSSImage from "./RSSItemImage";
import { useTheme } from "components/Theme/Theme";

type RSSItemPreviewProps = {
  item: RSSItemModel;
  organization: Organization | undefined;
  index: number;
  featured?: boolean;
  style?: React.CSSProperties;
};

function RSSItemPreview({
  item,
  organization,
  index,
  featured,
  style = {},
}: RSSItemPreviewProps) {
  const [theme] = useTheme();

  return (
    <Styled.RSSItemPreviewWrapper theme={theme} style={style}>
      <Styled.RSSItemPreview theme={theme}>
        <div
          style={{
            display: "flex",
            flex: 2,
            flexDirection: "column",
            minWidth: 300,
          }}
        >
          <RSSItemMetadata item={item} organization={organization} />
          <Styled.RSSItemContentContainer href={item.link}>
            <Styled.RSSItemTitle
              style={{
                fontSize: featured ? 27 : 20,
                fontFamily: featured ? "Source Serif Pro" : "inherit",
                position: "relative",
              }}
            >
              <span>{item.title}</span>
            </Styled.RSSItemTitle>
            <Styled.RSSItemDescription theme={theme}>
              {item.description}
            </Styled.RSSItemDescription>
          </Styled.RSSItemContentContainer>
        </div>
        <RSSImage
          href={item.link}
          image={item.image}
          featured={featured}
          organizationLogoImageUrl={organization?.logoImageUrl}
        />
      </Styled.RSSItemPreview>
      {item.seeAlsoItem ? (
        <Styled.SeeAlso theme={theme}>
          <span style={{ fontWeight: 500 }}>See also:</span>{" "}
          <a
            href={item.seeAlsoItem.link}
            style={{
              color: "#06c",
            }}
          >
            {item.seeAlsoItem.title + " "}
            <span
              style={{ fontSize: 14, textOverflow: "ellipsis" }}
              className="material-symbols-outlined"
            >
              open_in_new
            </span>
          </a>
        </Styled.SeeAlso>
      ) : null}
    </Styled.RSSItemPreviewWrapper>
  );
}

export default RSSItemPreview;
