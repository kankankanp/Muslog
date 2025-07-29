"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ThemeProvider } from "next-themes";
import { useEffect } from "react";
import { Toaster } from "react-hot-toast";
import { useRefreshToken } from "./libs/hooks/api/useAuth";
import { ReduxProvider } from "./libs/store/ReduxProvider";

const queryClient = new QueryClient();

export function Providers({ children }: { children: React.ReactNode }) {
  const refreshTokenMutation = useRefreshToken();

  useEffect(() => {
    refreshTokenMutation.mutate();
  }, []);

  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider attribute="class" defaultTheme="light">
        <ReduxProvider>{children}</ReduxProvider>
        <Toaster />
      </ThemeProvider>
    </QueryClientProvider>
  );
}
