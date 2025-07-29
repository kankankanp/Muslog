
import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import Footer from "./Footer";

describe("Footer", () => {
  it("renders the footer with navigation links and copyright", () => {
    render(<Footer />);
    expect(screen.getByAltText("Muslog")).toBeInTheDocument();
    expect(screen.getByRole("link", { name: /ホーム/i })).toHaveAttribute("href", "/");
    expect(screen.getByRole("link", { name: /管理/i })).toHaveAttribute("href", "/dashboard");
    expect(screen.getByRole("link", { name: /記事/i })).toHaveAttribute("href", "/dashboard/blog/page/1");
    expect(screen.getByText(/© \d{4} BLOG. All rights reserved./i)).toBeInTheDocument();
  });
});
