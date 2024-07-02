import { Theme } from "components/Theme/Theme";
import styled from "styled-components";

export const Layout = styled.div<{ theme: Theme }>`
  display: flex;
  align-items: center;
  flex-direction: column;
  background-color: ${({ theme }) => theme.colors.neutral.background};
  color: ${({ theme }) => theme.colors.neutral.text};
  min-height: 100%;
`;

export const Navbar = styled.div<{ theme: Theme }>`
  display: flex;
  outline: 1px solid ${({ theme }) => theme.colors.neutral.border};
  height: 55px;
  width: 100%;
  align-items: center;
  justify-content: center;
  font-size: 18px;
  padding-left: 20px;
  padding-right: 20px;
  box-sizing: border-box;
`;

export const PageWrapper = styled.div`
  display: flex;
  width: 100%;
  box-sizing: border-box;
  max-width: 1200px;
  flex-wrap: wrap;
  padding: 20px;
  gap: 20px;
`;

export const Footer = styled.div<{ theme: Theme }>`
  width: 100%;
  text-align: center;
  padding-top: 40px;
  padding-bottom: 40px;
  font-size: 14px;
  color: ${({ theme }) => theme.colors.neutral.secondaryText};
  border-top: 1px solid ${({ theme }) => theme.colors.neutral.border};
`;
