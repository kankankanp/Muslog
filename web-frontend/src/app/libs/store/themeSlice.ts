"use client";

import { createSlice, PayloadAction } from "@reduxjs/toolkit";

type ThemeMode = "light" | "dark";

interface ThemeState {
  theme: ThemeMode;
}

const getInitialTheme = (): ThemeMode => {
  if (typeof window !== "undefined") {
    return (localStorage.getItem("theme") as ThemeMode) || "light";
  }
  return "light";
};

const initialState: ThemeState = {
  theme: getInitialTheme(),
};

const themeSlice = createSlice({
  name: "theme",
  initialState,
  reducers: {
    setTheme: (state, action: PayloadAction<ThemeMode>) => {
      state.theme = action.payload;
      if (typeof window !== "undefined") {
        localStorage.setItem("theme", action.payload);
      }
    },
  },
});

export const { setTheme } = themeSlice.actions;
export default themeSlice.reducer;
