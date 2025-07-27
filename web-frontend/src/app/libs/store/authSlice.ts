import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import { AuthService } from "@/app/libs/api/generated/services/AuthService";

interface AuthState {
  user: any | null;
  loading: boolean;
  error: string | null;
}

const initialState: AuthState = {
  user: null,
  loading: false,
  error: null,
};

export const fetchUser = createAsyncThunk(
  "auth/fetchUser",
  async (_, { rejectWithValue }) => {
    try {
      const response = await AuthService.getMe();
      return response;
    } catch (error: any) {
      return rejectWithValue(error.body);
    }
  }
);

export const logout = createAsyncThunk(
  "auth/logout",
  async (_, { rejectWithValue }) => {
    try {
      await AuthService.postLogout();
      return true;
    } catch (error: any) {
      return rejectWithValue(error.body);
    }
  }
);

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchUser.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(fetchUser.fulfilled, (state, action) => {
        state.loading = false;
        state.user = action.payload;
      })
      .addCase(fetchUser.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
        state.user = null;
      })
      .addCase(logout.fulfilled, (state) => {
        state.user = null;
        state.error = null;
        state.loading = false;
      });
  },
});

export default authSlice.reducer;
