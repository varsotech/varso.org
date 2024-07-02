import { useStoredItem } from "components/LocalStorage/LocalStorage";
import React, {
  createContext,
  useContext,
  useState,
  ReactNode,
  useEffect,
  useMemo,
} from "react";

export interface Theme {
  colors: {
    neutral: {
      background: string;
      text: string;
      border: string;
      secondaryText: string;
      imageFallbackBackground: string;
      buttonBackground: string;
      dropdownSelected: string;
      orgUnderlineColor?: string;
      tabColor: string;
      selectedTabColor: string;
    };
  };
}

export interface ThemeContextProps {
  theme: Theme;
  setTheme: (themeName?: string) => void;
  storedThemeName: string | undefined;
}

const ThemeContext = createContext<ThemeContextProps | undefined>(undefined);

export interface ThemeProviderProps {
  lightTheme: string;
  darkTheme: string;
  themes: Record<string, Theme>;
  children: ReactNode;
}

export const ThemeProvider: React.FC<ThemeProviderProps> = ({
  lightTheme,
  darkTheme,
  themes,
  children,
}) => {
  const [storedTheme, setTheme] = useStoredItem("theme");

  const [preferredTheme, setPreferredTheme] = useState(
    window.matchMedia("(prefers-color-scheme: dark)").matches
      ? darkTheme
      : lightTheme
  );
  const currentTheme = useMemo(
    () => storedTheme || preferredTheme,
    [storedTheme, preferredTheme]
  );

  useEffect(() => {
    function handleDarkModePrefferedChange(e: MediaQueryListEvent) {
      setPreferredTheme(e.matches ? darkTheme : lightTheme);
    }

    window
      .matchMedia("(prefers-color-scheme: dark)")
      .addEventListener("change", handleDarkModePrefferedChange);

    // good house keeping to remove listener, good article here https://www.pluralsight.com/guides/how-to-cleanup-event-listeners-react
    return () => {
      window
        .matchMedia("(prefers-color-scheme: dark)")
        .removeEventListener("change", handleDarkModePrefferedChange);
    };
  }, []);

  return (
    <ThemeContext.Provider
      value={{
        theme: themes[currentTheme],
        setTheme,
        storedThemeName: storedTheme,
      }}
    >
      {children}
    </ThemeContext.Provider>
  );
};

export const useTheme = (): [
  Theme,
  (themeName?: string) => void,
  string | undefined
] => {
  const context = useContext(ThemeContext);

  if (!context) {
    throw new Error("useTheme must be used within a ThemeProvider");
  }

  return [context.theme, context.setTheme, context.storedThemeName];
};
