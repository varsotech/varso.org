import styled from "styled-components";

export const BlurStretch = styled.div<{ blur?: boolean; unblur?: boolean }>`
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;

  background-color: ${({ blur, unblur }) =>
    blur && !unblur ? "rgba(59, 59, 59, 0.85)" : "none"};

  backdrop-filter: ${({ blur, unblur }) =>
    blur && !unblur ? "blur(15px)" : "none"};
  -webkit-backdrop-filter: ${({ blur, unblur }) =>
    blur && !unblur ? "blur(15px)" : "none"};

  transition: backdrop-filter 0.3s;
`;

export const BlurWarning = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  color: white;
  gap: 10;
  font-size: 14;
`;
