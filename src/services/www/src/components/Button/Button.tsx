import * as Styled from "./Button.style";

export type ButtonProps = {
  style: React.CSSProperties;
  children: React.ReactNode;
  onClick: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
  disabled?: boolean;
};

function Button({ style, children, onClick, disabled }: ButtonProps) {
  return (
    <Styled.ButtonBase style={style} onClick={onClick} disabled={disabled}>
      {children}
    </Styled.ButtonBase>
  );
}

export default Button;
