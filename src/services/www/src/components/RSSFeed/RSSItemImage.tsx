import { RSSImage as RSSImageModel } from "@luminancetech/varso/src/app/appbase";
import React, { useState } from "react";
import * as Styled from "./RSSItemImage.style";
import config from "config/config";
import { useTheme } from "components/Theme/Theme";
import Button from "components/Button/Button";

type RSSImageProps = {
  href: string;
  image: RSSImageModel | undefined;
  organizationLogoImageUrl: string | undefined;
  featured?: boolean;
};

function RSSImage({
  href,
  image,
  featured,
  organizationLogoImageUrl,
}: RSSImageProps) {
  const [theme] = useTheme();
  const [unblur, setUnblur] = useState(false);

  var Wrapper = ({
    children,
    href,
    ...props
  }: {
    children: React.ReactNode;
    href: string;
    [key: string]: any;
  }) => (
    <a href={href} {...props}>
      {children}
    </a>
  );

  if (image?.blur && !unblur) {
    Wrapper = ({
      children,
      href,
      ...props
    }: {
      children: React.ReactNode;
      href: string;
      [key: string]: any;
    }) => <div {...props}>{children}</div>;
  }

  return (
    <Wrapper
      href={image?.blur && !unblur ? "#" : href}
      style={{
        position: "relative",
        flex: featured ? 4 : 1,
        maxWidth: 700,
        borderRadius: 5,
        overflow: "hidden",
        minWidth: 250,
        maxHeight: 400,
        minHeight: 150,
      }}
    >
      {image?.url ? (
        <img
          src={image?.url}
          alt={image?.title}
          width={"100%"} // 100% ?
          style={{
            display: "block",
            objectFit: "cover",
            objectPosition: "center",
            minWidth: 250,
            maxHeight: 400,
          }}
        />
      ) : (
        <div
          style={{
            height: 280,
            backgroundColor: theme.colors.neutral.imageFallbackBackground,
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
          }}
        >
          {organizationLogoImageUrl && (
            <img
              src={config.SYSTEM_EXTERNAL_URL + organizationLogoImageUrl}
              alt={"Organization Logo"}
              style={{
                display: "block",
                width: "100%",
                maxWidth: 400,
                padding: 60,
                boxSizing: "border-box",
              }}
            />
          )}
        </div>
      )}
      <Styled.BlurStretch
        blur={image?.blur || undefined}
        unblur={unblur ? unblur : undefined}
      />
      <div
        style={{
          position: "absolute",
          height: "100%",
          width: "100%",
          top: 0,
          left: 0,
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          textAlign: "center",
          boxShadow: featured ? "0px 0px 30px 3px #dadada" : "none",
        }}
      >
        {image?.blur && !unblur && (
          <Styled.BlurWarning>
            <span
              style={{
                cursor: "default",
                padding: 5,
                marginBottom: 5,
                color: "#f0f0f0",
              }}
            >
              This image was marked as sensitive.
            </span>
            <Button
              style={{
                backgroundColor: "#343434",
                color: "#f0f0f0",
                border: "transparent",
                fontWeight: 600,
                paddingLeft: 20,
                paddingRight: 20,
                paddingTop: 8,
                paddingBottom: 8,
              }}
              onClick={() => setUnblur(true)}
            >
              View Anyway
            </Button>
          </Styled.BlurWarning>
        )}
      </div>
    </Wrapper>
  );
}

export default RSSImage;
