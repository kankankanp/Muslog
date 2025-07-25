'use client';

import { configureStore } from '@reduxjs/toolkit';
import themeReducer from './themeSlice';
import authReducer from '../../store/authSlice';

export const store = configureStore({
  reducer: {
    theme: themeReducer,
    auth: authReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
export default store;
