
import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import Page from "./page";

describe("Dashboard Page", () => {
  it("renders the library link", () => {
    render(<Page />);
    const link = screen.getByRole("link", { name: /あなたのライブラリ/i });
    expect(link).toBeInTheDocument();
    expect(link).toHaveAttribute("href", "/library");
  });
});
