
import { render, screen } from "@testing-library/react";
import { describe, it, expect, vi } from "vitest";
import LoadingButton from "./LoadingButton";

vi.mock("react-dom", () => ({
  useFormStatus: () => ({ pending: false }),
}));

describe("LoadingButton", () => {
  it("renders the button with the correct label", () => {
    render(<LoadingButton label="Submit" />);
    expect(screen.getByRole("button", { name: /Submit/i })).toBeInTheDocument();
  });

  it("is disabled and shows a spinner when pending", () => {
    vi.spyOn(require("react-dom"), "useFormStatus").mockReturnValue({ pending: true });
    render(<LoadingButton label="Submit" />);
    const button = screen.getByRole("button");
    expect(button).toBeDisabled();
    expect(button.querySelector(".animate-spin")).toBeInTheDocument();
  });

  it("applies the correct color class", () => {
    render(<LoadingButton label="Submit" color="green" />);
    const button = screen.getByRole("button");
    expect(button).toHaveClass("bg-green-500");
  });
});
