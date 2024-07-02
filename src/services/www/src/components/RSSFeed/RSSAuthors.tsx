import {
  Organization,
  RSSAuthor as RSSAuthorModel,
} from "@luminancetech/varso/src/app/appbase";
import { useTheme } from "components/Theme/Theme";

type RSSAuthorsProps = {
  authors: RSSAuthorModel[];
  organization: Organization | undefined;
};

function RSSAuthors({ authors, organization }: RSSAuthorsProps) {
  const [theme] = useTheme();
  return (
    <span style={{ textTransform: "capitalize" }}>
      {authors
        ?.map((author: RSSAuthorModel) => {
          return author.name;
        })
        .join(", ")}
      {organization ? (
        <span>
          <span style={{ color: theme.colors.neutral.secondaryText }}>
            {" "}
            From{" "}
          </span>
          <a
            href={organization.websiteUrl}
            style={{
              fontSize: 17,
              fontFamily: "Source Serif Pro",
              fontWeight: 600,
              textDecoration: theme.colors.neutral.orgUnderlineColor
                ? "underline"
                : "inherit",
              textDecorationColor:
                theme.colors.neutral.orgUnderlineColor || "inherit",
            }}
          >
            {organization.name}
          </a>
        </span>
      ) : (
        ""
      )}
    </span>
  );
}

export default RSSAuthors;
