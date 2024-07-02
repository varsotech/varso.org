import Button from "components/Button/Button";
import { useTheme } from "components/Theme/Theme";

type ButtonToggleProps = {
  state: boolean | undefined;
  isLoading: boolean;
  onClick: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;
  onText: string;
  offText: string;
  style?: React.CSSProperties;
};

function ButtonToggle({
  state,
  isLoading,
  onClick,
  onText,
  offText,
  style = {},
}: ButtonToggleProps) {
  const [theme] = useTheme();

  return (
    <Button
      style={{
        display: "inline-block",
        padding: 3,
        marginLeft: 6,
        border: `1px solid ${theme.colors.neutral.border}`,
        height: 26,
        width: 26,
        ...style,
      }}
      onClick={onClick}
    >
      {isLoading ? ".." : state ? onText : offText}
    </Button>
  );
}

export default ButtonToggle;
