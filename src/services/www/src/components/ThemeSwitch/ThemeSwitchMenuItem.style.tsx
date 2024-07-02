import { Theme } from "components/Theme/Theme";
import styled from "styled-components";

export const ThemeSwitchMenuItem = styled.button<{
  theme: Theme;
  selected: boolean;
}>`
  padding: 5px 3px;
  color: ${({ theme, selected }) =>
    selected
      ? theme.colors.neutral.dropdownSelected
      : theme.colors.neutral.text};
  text-align: start;
  font-size: 13px;
  flex: 1;
  border-radius: 0;
  padding: 9px 10px;
  display: flex;
  justify-content: space-between;
  align-items: center;

  &:hover {
    opacity: 0.6;
  }
`;
