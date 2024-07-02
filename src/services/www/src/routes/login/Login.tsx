import { useEffect, useState } from "react";
import Layout from "../../components/Layout/Layout";
import { useLogin } from "../../api/auth";
import { LoginRequest } from "../../proto/src/auth/authrequests"; // TODO: Fix import to work with @
import { useNavigate } from "react-router-dom";
import { useToken } from "components/Auth/AuthProvider";
import { AxiosError } from "axios";
import Button from "components/Button/Button";

function Login() {
  const { token, setToken } = useToken();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { mutate: login, data, isPending, error } = useLogin();
  const navigate = useNavigate();

  const axiosError = error as AxiosError;

  useEffect(() => {
    if (data?.data?.token != null) {
      setToken(data?.data?.token);
      navigate("/");
    }
  }, [data?.data?.token, navigate, setToken]);

  const doLogin = () => {
    login(
      LoginRequest.create({
        usernameOrEmail: email,
        password: password,
      })
    );
  };

  return (
    <Layout>
      <div
        style={{
          flex: 1,
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          marginTop: 150,
        }}
      >
        <form
          onSubmit={(e) => {
            e.preventDefault();
          }}
          style={{
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            justifyContent: "center",
            gap: 10,
          }}
        >
          <h1>Login</h1>
          <span>
            {axiosError && "😞"}
            {data && "😌"}
          </span>
          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e) => {
              setEmail(e.target.value);
            }}
          />
          <input
            type="password"
            placeholder="Password"
            value={password}
            onChange={(e) => {
              setPassword(e.target.value);
            }}
          />
          <Button
            style={{ width: "100%" }}
            disabled={isPending}
            onClick={doLogin}
          >
            Login
          </Button>
        </form>
      </div>
    </Layout>
  );
}

export default Login;
