import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { AuthResponse } from "../api/generated";

interface AuthState {
  accessToken: string | null;
  user: AuthResponse | null;
}

const initialState: AuthState = {
  accessToken: null,
  user: null,
};

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    login(state, action: PayloadAction<AuthResponse>) {
      state.user = action.payload;
      state.accessToken = action.payload.accessToken;
    },
    logout(state) {
      state.user = null;
      state.accessToken = null;
    },
  },
});

export const { login, logout } = authSlice.actions;
export default authSlice.reducer;
