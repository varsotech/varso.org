import { BrowserRouter, Route, Routes } from "react-router-dom";
import Home from "./routes/home/Home";
import { PersistQueryClientProvider } from "./api/reactquery";
import Login from "./routes/login/Login";
import { AuthProvider } from "components/Auth/AuthProvider";
import Boycott from "routes/boycott/Boycott";
import { ThemeProvider } from "components/Theme/Theme";
import { lightTheme } from "components/Theme/light";
import { darkTheme } from "components/Theme/dark";
import { lightHighContrastTheme } from "components/Theme/light_highcontrast";

function App() {
  return (
    <div className="App">
      <PersistQueryClientProvider>
        <ThemeProvider
          lightTheme={"light"}
          darkTheme={"dark"}
          themes={{
            light: lightTheme,
            dark: darkTheme,
            light_highcontrast: lightHighContrastTheme,
          }}
        >
          <AuthProvider>
            <BrowserRouter>
              <Routes>
                <Route path="/" element={<Home />} />
                <Route path="/login" element={<Login />} />
                <Route path="/boycott" element={<Boycott />} />
              </Routes>
            </BrowserRouter>
          </AuthProvider>
        </ThemeProvider>
      </PersistQueryClientProvider>
    </div>
  );
}

export default App;
