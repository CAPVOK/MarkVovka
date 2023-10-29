import {
  createGlobalStyle,
  ThemeProvider as StyledThemeProvider,
  DefaultTheme,
} from "styled-components";
import React, { ReactNode, createContext, useContext } from "react";
import { useSelector } from "./core";
import { isDarkThemeSelector } from "./core/slices/appSlice";

interface ThemeProviderProps {
  children: ReactNode;
}

const GlobalStyle = createGlobalStyle`
:root {
  --background_color: ${(props) => props.theme.backgroundColor};
  --background_light_color: ${(props) => props.theme.backgroundLightColor};
  --primary_color: ${(props) => props.theme.primaryColor};
  --secondary_color: ${(props) => props.theme.secondaryColor};
  --accent_color: ${(props) => props.theme.accentColor};
  --error_color: ${(props) => props.theme.errorColor};
  --success_color: ${(props) => props.theme.successColor};
  --text_color: ${(props) => props.theme.textColor};
  --error_text_color: ${(props) => props.theme.errorTextColor};
  --success_text_color: ${(props) => props.theme.successTextColor};
  --border_color: ${(props) => props.theme.borderColor};
}
`;

interface Theme {
  backgroundColor: string;
  backgroundLightColor: string;
  primaryColor: string;
  borderColor: string;
  secondaryColor: string;
  accentColor: string;
  errorColor: string;
  successColor: string;
  errorTextColor: string;
  successTextColor: string;
  textColor: string;
}

const lightTheme: Theme = {
  backgroundColor: "#E0FFFF",
  backgroundLightColor: "#000080",
  primaryColor: "#191970",
  secondaryColor: "#40E0D0",
  borderColor: "#40E0D0",
  accentColor: "#FF6347",
  errorColor: "#FF6347",
  successColor: "#40ca6e",
  errorTextColor: "#fbfbfb",
  successTextColor: "#fbfbfb",
  textColor: "#070707",
};

const darkTheme: Theme = {
  backgroundColor: "#141414",
  backgroundLightColor: "#002518",
  secondaryColor: "#1e3d32",

  primaryColor: "#7DFFAF",
  accentColor: "#FF5303",/* #FF5303 */
  
  borderColor: "#07593A",
  
  errorColor: "#f30f11",
  successColor: "#00F55A",
  errorTextColor: "white",
  successTextColor: "black",
  textColor: "white",
};

const ThemeContext = createContext<DefaultTheme | undefined>(undefined);

// eslint-disable-next-line react-refresh/only-export-components
export const useTheme = (): DefaultTheme => {
  const theme = useContext(ThemeContext);
  if (!theme) {
    throw new Error("useTheme must be used within a ThemeProvider");
  }
  return theme;
};

export const ThemeProvider: React.FC<ThemeProviderProps> = ({ children }) => {
  const darkMode = useSelector(isDarkThemeSelector);

  const theme: DefaultTheme = darkMode ? darkTheme : lightTheme;

  return (
    <ThemeContext.Provider value={theme}>
      <StyledThemeProvider theme={theme}>
        <GlobalStyle />
        {children}
      </StyledThemeProvider>
    </ThemeContext.Provider>
  );
};
