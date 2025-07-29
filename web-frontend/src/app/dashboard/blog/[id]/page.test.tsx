
import { render, screen } from "@testing-library/react";
import { useParams } from "next/navigation";
import { describe, it, expect, vi } from "vitest";
import Page from "./page";
import BlogCard from "@/app/components/elements/cards/BlogCard";

vi.mock("next/navigation", () => ({
  useParams: vi.fn(),
}));

vi.mock("@/app/components/elements/cards/BlogCard", () => ({
  __esModule: true,
  default: vi.fn(() => <div data-testid="blog-card">Mocked BlogCard</div>),
}));

describe("Blog Detail Page", () => {
  beforeEach(() => {
    vi.clearAllMocks();
    (useParams as vi.Mock).mockReturnValue({ id: "123" });
  });

  it("renders without crashing", () => {
    render(<Page />);
    expect(screen.getByTestId("blog-card")).toBeInTheDocument();
  });

  it("renders BlogCard with isDetailPage true", () => {
    render(<Page />);
    expect(BlogCard).toHaveBeenCalledWith(expect.objectContaining({ isDetailPage: true }), {});
  });
});
