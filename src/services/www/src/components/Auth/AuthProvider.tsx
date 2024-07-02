import { setAxiosToken } from "api/axios";
import React, { createContext, useContext, useEffect, useReducer } from "react";
import { JWTClaims, parseJwt } from "utils/jwt";

type AuthState = {
  token: string | undefined;
};

const initialState: AuthState = {
  token: undefined, // undefined means it's not loaded. empty string means it's missing.
};

const AuthContext = createContext<
  { state: AuthState; dispatch: Dispatch } | undefined
>(undefined);

type Action = {
  type: "setToken";
  token: string;
};
type Dispatch = (action: Action) => void;

function authReducer(state: AuthState, action: Action): AuthState {
  switch (action.type) {
    case "setToken": {
      return { ...state, token: action.token };
    }
    default: {
      throw new Error(`Unhandled action: ${action}`);
    }
  }
}

type AuthProviderProps = {
  children: React.ReactNode;
};

export function AuthProvider({ children }: AuthProviderProps) {
  const [state, dispatch] = useReducer(authReducer, initialState);
  const value = { state, dispatch };

  useEffect(() => {
    const fetchToken = async () => {
      const fetchedToken = localStorage.getItem("token");
      dispatch({ type: "setToken", token: fetchedToken || "" });
    };
    fetchToken();
  }, []);

  return (
    <AuthContext.Provider value={value}>
      <AuthRedirection>{children}</AuthRedirection>
    </AuthContext.Provider>
  );
}

type AuthRedirectionProps = {
  children: React.ReactNode;
};

function AuthRedirection({ children }: AuthRedirectionProps) {
  const { token, setToken } = useToken();

  useEffect(() => {
    if (token) {
      setAxiosToken(token, () => {
        setToken("");
      });
    }
  }, [token, setToken]);

  return <div>{children}</div>;
}

function useAuth(): [AuthState, Dispatch] {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within a AuthContext");
  }
  return [context.state, context.dispatch];
}

export function useToken(): {
  token: string | undefined;
  setToken: (token: string) => void;
  claims: JWTClaims | undefined;
} {
  const [state, dispatch] = useAuth();

  function setToken(token: string) {
    if (token == null || token === "") {
      localStorage.removeItem("token");
    } else {
      localStorage.setItem("token", token);
    }
    dispatch({ type: "setToken", token: token });
  }

  var claims: JWTClaims | undefined = undefined;
  if (state.token != null && state.token !== "") {
    claims = parseJwt(state.token);
  }

  return { token: state.token, setToken, claims };
}

export default AuthContext;
