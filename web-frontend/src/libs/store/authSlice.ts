import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { AuthResponse } from "../api/generated/orval/model/authResponse";

interface AuthState {
  accessToken: any;
  user: AuthResponse | null;
  isAuthenticated: boolean;
}

const initialState: AuthState = {
  user: null,
  accessToken: undefined,
  isAuthenticated: false
};

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    login(state, action: PayloadAction<AuthResponse>) {
      state.user = action.payload;
      state.isAuthenticated = true;
    },
    logout(state) {
      state.user = null;
      state.isAuthenticated = false;
    },
  },
});

export const { login, logout } = authSlice.actions;
export default authSlice.reducer;
