
import { render, screen } from "@testing-library/react";
import { describe, it, expect, vi } from "vitest";
import RootLayout from "./layout";
import { Providers } from "./providers";

// Mock the Providers component
vi.mock("./providers", () => ({
  Providers: vi.fn(({ children }) => <div data-testid="providers-mock">{children}</div>),
}));

describe("RootLayout", () => {
  it("renders children within Providers component", () => {
    render(
      <RootLayout>
        <div data-testid="test-child">Test Child</div>
      </RootLayout>
    );

    expect(screen.getByTestId("providers-mock")).toBeInTheDocument();
    expect(screen.getByTestId("test-child")).toBeInTheDocument();
    expect(screen.getByTestId("providers-mock")).toContainElement(
      screen.getByTestId("test-child")
    );
  });
});
