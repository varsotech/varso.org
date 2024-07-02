import { useTheme } from "components/Theme/Theme";
import * as Styled from "./Layout.style";
import ThemeSwitch from "components/ThemeSwitch/ThemeSwitch";

type LayoutProps = {
  children: React.ReactNode;
  style?: React.CSSProperties;
  footer?: boolean;
};

function Layout({ children, style, footer }: LayoutProps) {
  const [theme] = useTheme();

  return (
    <Styled.Layout theme={theme}>
      <meta name="theme-color" content={theme.colors.neutral.background} />
      <Styled.Navbar theme={theme}>
        <div
          style={{
            flex: 1,
            maxWidth: 1200,
            display: "flex",
            justifyContent: "space-between",
          }}
        >
          <a
            href="/"
            style={{
              display: "flex",
              alignItems: "center",
              fontSize: 24,
              fontFamily: "Source Serif Pro",
            }}
          >
            <b>Varso</b>
            {/* <span
              style={{
                fontWeight: 100,
              }}
            >
              {"_"}
            </span> */}
          </a>
          <ThemeSwitch />
        </div>
      </Styled.Navbar>
      <Styled.PageWrapper style={style}>{children}</Styled.PageWrapper>
      {footer && (
        <Styled.Footer theme={theme}>
          <a href="/login">Login</a>
        </Styled.Footer>
      )}
    </Styled.Layout>
  );
}

export default Layout;
