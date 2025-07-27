import axios from "axios";
import { AuthService } from "@/app/libs/api/generated";
import { logout } from "@/app/libs/store/authSlice";
import { store } from "@/app/libs/store/store";

const NEXT_PUBLIC_API_URL: string =
  process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

const apiClient = axios.create({
  baseURL: NEXT_PUBLIC_API_URL,
  withCredentials: true,
});

apiClient.interceptors.request.use(
  (config) => {
    const state = store.getState();
    const accessToken = state.auth.accessToken;
    if (accessToken) {
      config.headers.Authorization = `Bearer ${accessToken}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;
    if (error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;
      try {
        const refreshResponse = await AuthService.postRefresh();
        const newAccessToken = refreshResponse.accessToken;

        // store.dispatch(login());
        originalRequest.headers.Authorization = `Bearer ${newAccessToken}`;
        return apiClient(originalRequest);
      } catch (refreshError) {
        store.dispatch(logout());
        window.location.href = "/registration/login";
        return Promise.reject(refreshError);
      }
    }
    return Promise.reject(error);
  }
);

export default apiClient;
