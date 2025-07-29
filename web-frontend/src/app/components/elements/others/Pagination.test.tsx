
import { render, screen } from "@testing-library/react";
import { useParams } from "next/navigation";
import { describe, it, expect, vi } from "vitest";
import Pagination from "./Pagination";

vi.mock("next/navigation", () => ({
  useParams: vi.fn(),
}));

describe("Pagination", () => {
  it("renders the correct number of pages", () => {
    (useParams as vi.Mock).mockReturnValue({ page: "1" });
    render(<Pagination totalCount={20} />);
    const pageItems = screen.getAllByRole("link");
    expect(pageItems).toHaveLength(5);
  });

  it("highlights the current page", () => {
    (useParams as vi.Mock).mockReturnValue({ page: "3" });
    render(<Pagination totalCount={20} />);
    const currentPageItem = screen.getByText("3").closest("li");
    expect(currentPageItem).toHaveClass("is-active");
  });
});
