"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ThemeProvider } from "next-themes";
import { Toaster } from "react-hot-toast";
import { ReduxProvider } from "../libs/store/ReduxProvider";
import AuthInitializer from "@/components/guards/AuthInitializer";

const queryClient = new QueryClient();

export function Providers({ children }: { children: React.ReactNode }) {
  return (
    <QueryClientProvider client={queryClient}>
      <ThemeProvider attribute="class" defaultTheme="light">
        <ReduxProvider>
          <AuthInitializer>{children}</AuthInitializer>
        </ReduxProvider>
        <Toaster />
      </ThemeProvider>
    </QueryClientProvider>
  );
}
