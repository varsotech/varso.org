import axios, { AxiosError, AxiosInstance } from "axios";
import config from "../config/config";

const axiosInstance: AxiosInstance = fetchClient();

function fetchClient() {
  return axios.create({
    baseURL: config.SYSTEM_EXTERNAL_URL,
  });
}

export function setAxiosToken(token: string, deleteTokenFn: () => void) {
  // Add an interceptor to inject the token in the Authorization header
  axiosInstance.interceptors.request.use(function (config) {
    config.headers.Authorization = token ? `Bearer ${token}` : "";
    return config;
  });

  // Add an interceptor to handle 401 responses
  axiosInstance.interceptors.response.use(
    (response) => response,
    (error: AxiosError<{ code: string }>) => {
      if (error.response && error.response.status === 401) {
        deleteTokenFn();
      }
    }
  );
}

export default axiosInstance;
