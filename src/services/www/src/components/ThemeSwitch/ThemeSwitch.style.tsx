import { Theme } from "components/Theme/Theme";
import styled from "styled-components";

export const ThemeSwitch = styled.button<{ theme: Theme }>`
  display: flex;
  width: 30px;
  text-align: center;
  align-items: center;
  justify-content: center;
  background-color: initial;
  color: ${({ theme }) => theme.colors.neutral.text};
  border-radius: 10px;
  padding: 7px 15px;
`;

export const ThemeSwitchMenu = styled.div<{ theme: Theme }>`
  background-color: ${({ theme }) => theme.colors.neutral.background};
  /* color: ${({ theme }) => theme.colors.neutral.secondaryText}; */
  border: 1px solid ${({ theme }) => theme.colors.neutral.border};
  border-radius: 10px;
  z-index: 5;
  width: 155px;
  display: flex;
  flex-direction: column;
  overflow: hidden;
`;
