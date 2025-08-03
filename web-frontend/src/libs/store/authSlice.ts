import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { AuthResponse } from "../api/generated/orval/model/authResponse";

interface AuthState {
  accessToken: string | undefined; // string型に変更
  user: AuthResponse | null;
  isAuthenticated: boolean;
  isInitialized: boolean;
  tokenExpiry?: number; // トークンの有効期限を管理
}

const initialState: AuthState = {
  user: null,
  accessToken: undefined,
  isAuthenticated: false,
  isInitialized: false,
  tokenExpiry: undefined
};

// localStorage操作のヘルパー関数
const saveAuthToStorage = (state: AuthState) => {
  if (typeof window !== 'undefined') {
    const authData = {
      user: state.user,
      isAuthenticated: state.isAuthenticated,
      tokenExpiry: state.tokenExpiry,
      accessToken: state.accessToken, // accessTokenも保存
      timestamp: Date.now()
    };
    localStorage.setItem('auth', JSON.stringify(authData));
  }
};

const clearAuthFromStorage = () => {
  if (typeof window !== 'undefined') {
    localStorage.removeItem('auth');
  }
};

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    login(state, action: PayloadAction<AuthResponse & { tokenExpiry?: number; accessToken: string }>) { // accessTokenを追加
      state.user = action.payload;
      state.isAuthenticated = true;
      state.isInitialized = true;
      state.tokenExpiry = action.payload.tokenExpiry || Date.now() + (24 * 60 * 60 * 1000);
      state.accessToken = action.payload.accessToken; // accessTokenを保存
      saveAuthToStorage(state);
    },
    logout(state) {
      state.user = null;
      state.accessToken = undefined; // accessTokenをクリア
      state.isAuthenticated = false;
      state.isInitialized = true;
      state.tokenExpiry = undefined;
      clearAuthFromStorage();
    },
    initializeAuth(state, action: PayloadAction<{ user: AuthResponse | null; tokenExpiry?: number; accessToken?: string }>) { // accessTokenを追加
      state.user = action.payload.user;
      state.isAuthenticated = !!action.payload.user;
      state.isInitialized = true;
      state.tokenExpiry = action.payload.tokenExpiry;
      state.accessToken = action.payload.accessToken; // accessTokenを保存
      if (state.isAuthenticated) {
        saveAuthToStorage(state);
      }
    },
    // トークン期限切れ時の処理
    expireAuth(state) {
      state.user = null;
      state.accessToken = undefined; // accessTokenをクリア
      state.isAuthenticated = false;
      state.tokenExpiry = undefined;
      clearAuthFromStorage();
    },
  },
});

export const { login, logout, initializeAuth, expireAuth } = authSlice.actions;
export default authSlice.reducer;
