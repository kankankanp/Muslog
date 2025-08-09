import { RootState } from "./store";

export const selectIsAuthenticated = (state: RootState) => state.auth.isAuthenticated;
export const selectAuthUser = (state: RootState) => state.auth.user;
export const selectAccessToken = (state: RootState) => state.auth.accessToken;
export const selectIsAuthInitialized = (state: RootState) => state.auth.isInitialized;