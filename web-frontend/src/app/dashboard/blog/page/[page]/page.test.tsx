
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, screen } from "@testing-library/react";
import { useParams } from "next/navigation";
import { describe, it, expect, vi, beforeEach } from "vitest";
import Page from "./page";
import AddButton from "@/app/components/elements/buttons/AddButton";
import BlogCard from "@/app/components/elements/cards/BlogCard";
import Pagination from "@/app/components/elements/others/Pagination";
import { useGetBlogsByPage } from "@/app/libs/hooks/api/useBlogs";

vi.mock("next/navigation", () => ({
  useParams: vi.fn(),
}));

vi.mock("@/app/components/elements/buttons/AddButton", () => ({
  __esModule: true,
  default: vi.fn(() => <div data-testid="add-button">Add Button</div>),
}));

vi.mock("@/app/components/elements/cards/BlogCard", () => ({
  __esModule: true,
  default: vi.fn(() => <div data-testid="blog-card">Blog Card</div>),
}));

vi.mock("@/app/components/elements/others/Pagination", () => ({
  __esModule: true,
  default: vi.fn(() => <div data-testid="pagination">Pagination</div>),
}));

vi.mock("@/app/libs/hooks/api/useBlogs", () => ({
  useGetBlogsByPage: vi.fn(),
}));

const queryClient = new QueryClient();

describe("Blog Page", () => {
  beforeEach(() => {
    vi.clearAllMocks();
    (useParams as vi.Mock).mockReturnValue({ page: "1" });
    (useGetBlogsByPage as vi.Mock).mockReturnValue({
      data: { posts: [], totalCount: 0 },
      isPending: false,
      error: null,
    });
  });

  const renderBlogPage = () => {
    render(
      <QueryClientProvider client={queryClient}>
        <Page />
      </QueryClientProvider>
    );
  };

  it("renders AddButton, BlogCard, and Pagination", () => {
    renderBlogPage();
    expect(screen.getByTestId("add-button")).toBeInTheDocument();
    expect(screen.getByTestId("blog-card")).toBeInTheDocument();
    expect(screen.getByTestId("pagination")).toBeInTheDocument();
  });

  it("displays loading state", () => {
    (useGetBlogsByPage as vi.Mock).mockReturnValue({ isPending: true });
    renderBlogPage();
    expect(screen.getByText("Loading...")).toBeInTheDocument();
  });

  it("displays error state", () => {
    (useGetBlogsByPage as vi.Mock).mockReturnValue({ error: new Error("Test Error"), isPending: false });
    renderBlogPage();
    expect(screen.getByText("Error: Test Error")).toBeInTheDocument();
  });
});
