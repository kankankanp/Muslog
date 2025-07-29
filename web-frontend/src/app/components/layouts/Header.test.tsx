
import { render, screen } from "@testing-library/react";
import { describe, it, expect, vi } from "vitest";
import Header from "./Header";
import ThemeToggleButton from "../elements/buttons/ThemeToggleButton";

// Mock ThemeToggleButton as it's an external component and already tested
vi.mock("../elements/buttons/ThemeToggleButton", () => ({
  __esModule: true,
  default: vi.fn(() => <div data-testid="theme-toggle-button">Theme Toggle Button</div>),
}));

describe("Header", () => {
  it("renders the header with navigation links and theme toggle button", () => {
    render(<Header />);

    // Check for logo link
    expect(screen.getByAltText("BLOG")).toBeInTheDocument();
    expect(screen.getByRole("link", { name: /BLOG/i })).toHaveAttribute("href", "/");

    // Check for navigation links
    expect(screen.getByRole("link", { name: /ホーム/i })).toHaveAttribute("href", "/");
    expect(screen.getByRole("link", { name: /管理/i })).toHaveAttribute("href", "/dashboard");
    expect(screen.getByRole("link", { name: /記事/i })).toHaveAttribute("href", "/dashboard/blog/page/1");

    // Check for ThemeToggleButton
    expect(screen.getByTestId("theme-toggle-button")).toBeInTheDocument();
  });
});
