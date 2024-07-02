import RSSItemPreview from "../../components/RSSFeed/RSSItemPreview";
import { useNews } from "../../api/news";
import Layout from "../../components/Layout/Layout";
import RSSFeedPreview from "components/RSSFeed/RSSFeedPreview";
import { useEffect, useState } from "react";

function Home() {
  const [selectedFeedTab, setSelectedFeedTab] = useState("foryou");
  const { data, isLoading, isFetching, refetch } = useNews(selectedFeedTab);

  useEffect(() => {
    refetch();
  }, [refetch, selectedFeedTab]);

  if (isLoading) {
    return (
      <Layout style={{ flexDirection: "column", flex: 1 }}>
        <div style={{ flex: 1 }} />
      </Layout>
    );
  }

  return (
    <Layout footer>
      {data?.data.featured && (
        <RSSItemPreview
          style={{ border: "none" }}
          item={data?.data.featured}
          organization={
            data?.data.organizations[data?.data.featured.organizationUuid]
          }
          featured
          index={1}
        />
      )}
      <div style={{ display: "flex", width: "100%" }}>
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            // borderRight: "1px solid #eaeaea",
            // paddingRight: 20,
            flex: 4,
            maxWidth: 900,
          }}
        >
          <div style={{ display: "flex", flexDirection: "column" }}>
            <RSSFeedPreview
              items={data?.data?.latest?.items || []}
              organizations={data?.data?.organizations || {}}
              selectedFeedTab={selectedFeedTab}
              setSelectedFeedTab={setSelectedFeedTab}
              isFetching={isFetching}
            />
          </div>
        </div>
        {/* <div
          style={{
            display: "flex",
            flex: 1,
          }}
        >
          <h2>Organizations</h2>
        </div> */}
      </div>
    </Layout>
  );
}

export default Home;
