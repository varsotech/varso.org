import { useTheme } from "components/Theme/Theme";
import * as Styled from "./ThemeSwitch.style";
import { useState } from "react";
import ThemeSwitchMenuItem from "./ThemeSwitchMenuItem";

type ThemeSwitchProps = {};

function ThemeSwitch({}: ThemeSwitchProps) {
  const [theme] = useTheme();
  const [menuOpen, setMenuOpen] = useState(false);

  return (
    <div style={{ position: "relative" }}>
      <Styled.ThemeSwitch theme={theme} onClick={() => setMenuOpen(!menuOpen)}>
        <span style={{ fontSize: 18 }} className="material-symbols-outlined">
          contrast
        </span>
      </Styled.ThemeSwitch>
      <div
        style={{
          display: "flex",
          maxHeight: menuOpen ? "300px" : 0,
          overflow: "hidden",
          transition: "max-height 0.15s ease-out",
          right: 0,
          position: "absolute",
        }}
      >
        <Styled.ThemeSwitchMenu theme={theme} onBlur={() => setMenuOpen(false)}>
          <ThemeSwitchMenuItem
            displayName={"System"}
            themeName={undefined}
            closeMenu={() => setMenuOpen(false)}
          />
          <ThemeSwitchMenuItem
            displayName={"Light"}
            themeName={"light"}
            closeMenu={() => setMenuOpen(false)}
          />
          <ThemeSwitchMenuItem
            displayName={"Dark"}
            themeName={"dark"}
            closeMenu={() => setMenuOpen(false)}
          />
          <ThemeSwitchMenuItem
            displayName={"Light High-Contrast"}
            themeName={"light_highcontrast"}
            closeMenu={() => setMenuOpen(false)}
          />
        </Styled.ThemeSwitchMenu>
      </div>
    </div>
  );
}

export default ThemeSwitch;
