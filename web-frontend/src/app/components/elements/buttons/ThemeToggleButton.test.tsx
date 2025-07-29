
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { useTheme } from "next-themes";
import { describe, it, expect, vi, beforeEach } from "vitest";
import ThemeToggleButton from "./ThemeToggleButton";

vi.mock("next-themes", () => ({
  useTheme: vi.fn(),
}));

describe("ThemeToggleButton", () => {
  const mockSetTheme = vi.fn();

  beforeEach(() => {
    vi.clearAllMocks();
  });

  it("renders correctly with light theme", () => {
    (useTheme as vi.Mock).mockReturnValue({ theme: "light", setTheme: mockSetTheme });
    render(<ThemeToggleButton />);
    const checkbox = screen.getByRole("checkbox");
    expect(checkbox).not.toBeChecked();
  });

  it("renders correctly with dark theme", () => {
    (useTheme as vi.Mock).mockReturnValue({ theme: "dark", setTheme: mockSetTheme });
    render(<ThemeToggleButton />);
    const checkbox = screen.getByRole("checkbox");
    expect(checkbox).toBeChecked();
  });

  it("calls setTheme with 'dark' when toggled from light theme", async () => {
    (useTheme as vi.Mock).mockReturnValue({ theme: "light", setTheme: mockSetTheme });
    render(<ThemeToggleButton />);
    const checkbox = screen.getByRole("checkbox");
    await userEvent.click(checkbox);
    expect(mockSetTheme).toHaveBeenCalledWith("dark");
  });

  it("calls setTheme with 'light' when toggled from dark theme", async () => {
    (useTheme as vi.Mock).mockReturnValue({ theme: "dark", setTheme: mockSetTheme });
    render(<ThemeToggleButton />);
    const checkbox = screen.getByRole("checkbox");
    await userEvent.click(checkbox);
    expect(mockSetTheme).toHaveBeenCalledWith("light");
  });
});
