
import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import { CommonButton } from "./CommonButton";

describe("CommonButton", () => {
  it("renders a link with the correct href and children", () => {
    render(<CommonButton href="/test">Test Button</CommonButton>);
    const link = screen.getByRole("link", { name: /Test Button/i });
    expect(link).toBeInTheDocument();
    expect(link).toHaveAttribute("href", "/test");
  });

  it("applies the provided className", () => {
    render(
      <CommonButton href="/test" className="custom-class">
        Test Button
      </CommonButton>
    );
    const link = screen.getByRole("link");
    expect(link).toHaveClass("button");
    expect(link).toHaveClass("custom-class");
  });
});
