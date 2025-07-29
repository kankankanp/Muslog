
import { render, screen } from "@testing-library/react";
import { describe, it, expect, vi } from "vitest";
import Page from "./page";
import LoginForm from "@/app/components/elements/forms/LoginForm";

vi.mock("@/app/components/elements/forms/LoginForm", () => ({
  __esModule: true,
  default: vi.fn(() => <div data-testid="login-form">Login Form</div>),
}));

describe("Login Page", () => {
  it("renders LoginForm", () => {
    render(<Page />);
    expect(screen.getByTestId("login-form")).toBeInTheDocument();
  });
});
