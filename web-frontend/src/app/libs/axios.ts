import axios from "axios";
import { AuthService } from "@/app/libs/api/generated";
import { fetchUser, logout } from "@/app/libs/store/authSlice";
import { store } from "@/app/libs/store/store";

const apiClient = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL,
  withCredentials: true, // Cookieを送信するために必要
});

// Request interceptor to add access token to headers
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

// Response interceptor to handle token refresh
apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;
    // If the error is 401 Unauthorized and it's not a retry attempt
    if (error.response.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;
      try {
        // Attempt to refresh the token using AuthService
        const refreshResponse = await AuthService.postRefresh();
        // Assuming refreshResponse contains the new access token
        // You might need to adjust this based on your actual API response structure
        const newAccessToken = refreshResponse.accessToken; // Adjust this line

        // Update Redux store with new access token
        store.dispatch(fetchUser() as any); // Adjust this line

        // Update the original request with the new token and retry
        originalRequest.headers.Authorization = `Bearer ${newAccessToken}`;
        return apiClient(originalRequest);
      } catch (refreshError) {
        // If refresh fails, log out the user
        store.dispatch(logout());
        window.location.href = "/registration/login";
        return Promise.reject(refreshError);
      }
    }
    return Promise.reject(error);
  }
);

export default apiClient;
