import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { AuthResponse } from "../api/generated";

interface AuthState {
  accessToken: any;
  user: AuthResponse | null;
}

const initialState: AuthState = {
  user: null,
};

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    login(state, action: PayloadAction<AuthResponse>) {
      state.user = action.payload;
    },
    logout(state) {
      state.user = null;
    },
  },
});

export const { login, logout } = authSlice.actions;
export default authSlice.reducer;
