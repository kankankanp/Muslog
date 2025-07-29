
import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import AddButton from "./AddButton";

describe("AddButton", () => {
  it("renders a link to the add blog page", () => {
    render(<AddButton />);
    const link = screen.getByRole("link");
    expect(link).toBeInTheDocument();
    expect(link).toHaveAttribute("href", "/dashboard/blog/add");
  });
});
