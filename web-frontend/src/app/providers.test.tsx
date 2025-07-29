
import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ThemeProvider } from "next-themes";
import { Toaster } from "react-hot-toast";
import { describe, it, expect, vi, beforeEach } from "vitest";
import { Providers } from "./providers";
import { useRefreshToken } from "./libs/hooks/api/useAuth";
import { ReduxProvider } from "./libs/store/ReduxProvider";

// Mock child components
vi.mock("next-themes", () => ({
  ThemeProvider: vi.fn(({ children }) => <div data-testid="theme-provider">{children}</div>),
}));
vi.mock("react-hot-toast", () => ({
  Toaster: vi.fn(() => <div data-testid="toaster"></div>),
}));
vi.mock("./libs/hooks/api/useAuth", () => ({
  useRefreshToken: vi.fn(),
}));
vi.mock("./libs/store/ReduxProvider", () => ({
  ReduxProvider: vi.fn(({ children }) => <div data-testid="redux-provider">{children}</div>),
}));

const queryClient = new QueryClient();

describe("Providers", () => {
  const mockMutate = vi.fn();

  beforeEach(() => {
    vi.clearAllMocks();
    (useRefreshToken as vi.Mock).mockReturnValue({ mutate: mockMutate });
  });

  it("renders all providers and children", () => {
    render(
      <QueryClientProvider client={queryClient}>
        <Providers>
          <div data-testid="test-child">Test Child</div>
        </Providers>
      </QueryClientProvider>
    );

    expect(screen.getByTestId("theme-provider")).toBeInTheDocument();
    expect(screen.getByTestId("redux-provider")).toBeInTheDocument();
    expect(screen.getByTestId("toaster")).toBeInTheDocument();
    expect(screen.getByTestId("test-child")).toBeInTheDocument();
  });

  it("calls refreshTokenMutation on mount", () => {
    render(
      <QueryClientProvider client={queryClient}>
        <Providers>
          <div data-testid="test-child">Test Child</div>
        </Providers>
      </QueryClientProvider>
    );
    expect(mockMutate).toHaveBeenCalledTimes(1);
  });
});
