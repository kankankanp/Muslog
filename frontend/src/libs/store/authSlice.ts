import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { AuthResponse } from '../api/generated/orval/model/authResponse';

interface AuthState {
  accessToken: any;
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
  tokenExpiry: undefined,
};

// localStorage操作のヘルパー関数
const saveAuthToStorage = (state: AuthState) => {
  if (typeof window !== 'undefined') {
    const authData = {
      user: state.user,
      isAuthenticated: state.isAuthenticated,
      tokenExpiry: state.tokenExpiry,
      timestamp: Date.now(), // 保存時刻を記録
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
  name: 'auth',
  initialState,
  reducers: {
    login(
      state,
      action: PayloadAction<AuthResponse & { tokenExpiry?: number }>
    ) {
      state.user = action.payload;
      state.isAuthenticated = true;
      state.isInitialized = true;
      state.tokenExpiry =
        action.payload.tokenExpiry || Date.now() + 24 * 60 * 60 * 1000; // デフォルト24時間
      saveAuthToStorage(state);
    },
    logout(state) {
      state.user = null;
      state.isAuthenticated = false;
      state.isInitialized = true;
      state.tokenExpiry = undefined;
      clearAuthFromStorage();
    },
    initializeAuth(
      state,
      action: PayloadAction<{
        user: AuthResponse | null;
        tokenExpiry?: number;
      }>
    ) {
      state.user = action.payload.user;
      state.isAuthenticated = !!action.payload.user;
      state.isInitialized = true;
      state.tokenExpiry = action.payload.tokenExpiry;
      if (state.isAuthenticated) {
        saveAuthToStorage(state);
      }
    },
    // トークン期限切れ時の処理
    expireAuth(state) {
      state.user = null;
      state.isAuthenticated = false;
      state.tokenExpiry = undefined;
      clearAuthFromStorage();
    },
  },
});

export const { login, logout, initializeAuth, expireAuth } = authSlice.actions;
export default authSlice.reducer;
