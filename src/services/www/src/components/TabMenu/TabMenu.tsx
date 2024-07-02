import { useTheme } from "components/Theme/Theme";
import * as Styled from "./TabMenu.style";

export type TabMenuItemData = {
  id: string;
  text: string;
};

export type TabMenuProps = {
  options: TabMenuItemData[];
  selected: string;
  onChange: (newSelected: string) => void;
};

function TabMenu({ options, selected, onChange }: TabMenuProps) {
  const [theme] = useTheme();

  return (
    <Styled.TabMenu theme={theme}>
      {options.map((v: TabMenuItemData) => {
        return (
          <Styled.TabMenuItem
            key={v.id}
            theme={theme}
            selected={v.id === selected}
            onClick={() => onChange(v.id)}
          >
            {v.text}
          </Styled.TabMenuItem>
        );
      })}
    </Styled.TabMenu>
  );
}

export default TabMenu;
