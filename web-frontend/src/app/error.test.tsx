
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { describe, it, expect, vi } from "vitest";
import Error from "./error";

describe("Error Component", () => {
  it("renders the error message and a retry button", () => {
    const mockError = new Error("Test error message");
    const mockReset = vi.fn();

    render(<Error error={mockError} reset={mockReset} />);

    expect(screen.getByText("エラーが発生しました！")).toBeInTheDocument();
    expect(screen.getByText("Test error message")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /再試行/i })).toBeInTheDocument();
  });

  it("calls reset function when retry button is clicked", async () => {
    const mockError = new Error("Test error message");
    const mockReset = vi.fn();

    render(<Error error={mockError} reset={mockReset} />);

    await userEvent.click(screen.getByRole("button", { name: /再試行/i }));

    expect(mockReset).toHaveBeenCalledTimes(1);
  });
});
