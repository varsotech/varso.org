import { Theme } from "components/Theme/Theme";
import styled from "styled-components";

export const RSSItemPreviewWrapper = styled.div<{ theme: Theme }>`
  border-bottom: 1px solid ${({ theme }) => theme.colors.neutral.border};
  padding-bottom: 25px;
`;

export const RSSItemPreview = styled.div<{ theme: Theme }>`
  display: flex;
  flex-wrap: wrap-reverse;
  gap: 20px;
  align-items: center;
  justify-content: center;
`;

export const RSSItemContentContainer = styled.a`
  display: flex;
  flex: 1;
  flex-direction: column;
  /* flex-wrap: wrap; */
`;

export const RSSItemTitle = styled.h3``;

export const RSSItemDescription = styled.span<{ theme: Theme }>`
  font-size: 16;
  color: ${({ theme }) => theme.colors.neutral.secondaryText};
`;

export const SeeAlso = styled.div<{ theme: Theme }>`
  padding-top: 25px;
  font-size: 15px;
  color: ${({ theme }) => theme.colors.neutral.secondaryText};
`;
