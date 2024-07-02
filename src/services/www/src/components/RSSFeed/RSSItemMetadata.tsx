import { Organization, RSSItem } from "@luminancetech/varso/src/app/appbase";
import RSSAuthors from "./RSSAuthors";
import { timeAgo } from "../../utils/timeago";
import { useToken } from "components/Auth/AuthProvider";
import { useDelete, useToggleBlur } from "api/news";
import ButtonToggle from "components/ButtonToggle/ButtonToggle";
import { useTheme } from "components/Theme/Theme";

type RSSItemMetadataProps = {
  item: RSSItem;
  organization: Organization | undefined;
};

function RSSItemMetadata({ item, organization }: RSSItemMetadataProps) {
  const [theme] = useTheme();
  const { claims } = useToken();
  const { mutate: toggleBlur, isPending: isBlurPending } = useToggleBlur(
    item.id
  );
  const { mutate: deleteItem, isPending: isDeletePending } = useDelete(item.id);

  const hasSetBlurPermission =
    claims?.["permissions"]?.["news.item.image.blur"] === true;

  const hasDeletePermission =
    claims?.["permissions"]?.["news.item.delete"] === true;

  return (
    <div style={{ fontSize: 15, marginBottom: 8 }}>
      <RSSAuthors authors={item.authors} organization={organization} />
      <span
        style={{
          marginLeft: 5,
          color: theme.colors.neutral.secondaryText,
          display: "inline",
        }}
      >
        {item.publishDate ? "· " + timeAgo(item.publishDate) : ""}
      </span>
      {hasSetBlurPermission ? (
        <ButtonToggle
          state={item.image?.blur}
          isLoading={isBlurPending}
          onClick={() => toggleBlur(undefined)}
          onText={"🫣"}
          offText={"👁️"}
        />
      ) : null}
      {hasDeletePermission ? (
        <ButtonToggle
          state={false}
          isLoading={isDeletePending}
          onText={"✖️"}
          offText={"✖️"}
          onClick={() => deleteItem(undefined)}
        />
      ) : null}
    </div>
  );
}

export default RSSItemMetadata;
