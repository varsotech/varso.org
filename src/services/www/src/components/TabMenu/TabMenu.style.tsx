import { Theme } from "components/Theme/Theme";
import styled from "styled-components";

export const TabMenu = styled.div<{ theme: Theme }>`
  display: flex;
  font-weight: 500;
  /* border-bottom: 1px solid ${({ theme }) => theme.colors.neutral.border}; */
  box-shadow: ${({ theme }) => theme.colors.neutral.border} 0px -1px 0px inset;
  color: ${({ theme }) => theme.colors.neutral.selectedTabColor};
  flex: 1;
  margin-bottom: 25px;
`;

export const TabMenuItem = styled.button<{ theme: Theme; selected: boolean }>`
  font-size: 14px;
  padding: 20px 30px;
  border-bottom: 1px solid
    ${({ theme, selected }) =>
      selected ? theme.colors.neutral.selectedTabColor : "transparent"};

  color: ${({ theme, selected }) =>
    selected
      ? theme.colors.neutral.selectedTabColor
      : theme.colors.neutral.tabColor};
  transition: color 80ms ease-out;
  transition: border 80ms ease-out;
`;
