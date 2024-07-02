import { useTheme } from "components/Theme/Theme";
import * as Styled from "./ThemeSwitchMenuItem.style";

type ThemeSwitchMenuItemProps = {
  themeName: string | undefined;
  displayName: string;
  closeMenu: () => void;
};

export default function ThemeSwitchMenuItem({
  themeName,
  displayName,
  closeMenu,
}: ThemeSwitchMenuItemProps) {
  const [theme, setTheme, storedTheme] = useTheme();

  const selected = storedTheme === themeName;

  return (
    <Styled.ThemeSwitchMenuItem
      theme={theme}
      selected={selected}
      onClick={() => {
        setTheme(themeName);
        closeMenu();
      }}
    >
      <span>{displayName}</span>
      {selected && (
        <span style={{ fontSize: 14 }} className="material-symbols-outlined">
          done
        </span>
      )}
    </Styled.ThemeSwitchMenuItem>
  );
}
