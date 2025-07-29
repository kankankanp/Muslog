
import { render, screen } from "@testing-library/react";
import { describe, it, expect, vi } from "vitest";
import Page from "./page";
import SignupForm from "@/app/components/elements/forms/SignupForm";

vi.mock("@/app/components/elements/forms/SignupForm", () => ({
  __esModule: true,
  default: vi.fn(() => <div data-testid="signup-form">Signup Form</div>),
}));

describe("Signup Page", () => {
  it("renders SignupForm", () => {
    render(<Page />);
    expect(screen.getByTestId("signup-form")).toBeInTheDocument();
  });
});
