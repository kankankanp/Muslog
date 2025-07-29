
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { useParams, useRouter } from "next/navigation";
import toast from "react-hot-toast";
import { describe, it, expect, vi, beforeEach } from "vitest";
import Page from "./page";
import {
  useDeleteBlog,
  useGetBlogById,
  useUpdateBlog,
} from "@/app/libs/hooks/api/useBlogs";
import {
  useAddTagsToPost,
  useGetTagsByPostID,
  useRemoveTagsFromPost,
} from "@/app/libs/hooks/api/useTags";

vi.mock("next/navigation", () => ({
  useParams: vi.fn(),
  useRouter: vi.fn(),
}));

vi.mock("@/app/libs/hooks/api/useBlogs", () => ({
  useDeleteBlog: vi.fn(),
  useGetBlogById: vi.fn(),
  useUpdateBlog: vi.fn(),
}));

vi.mock("@/app/libs/hooks/api/useTags", () => ({
  useAddTagsToPost: vi.fn(),
  useGetTagsByPostID: vi.fn(),
  useRemoveTagsFromPost: vi.fn(),
}));

vi.mock("react-hot-toast", () => ({
  success: vi.fn(),
  error: vi.fn(),
  Toaster: vi.fn(() => null),
}));

const queryClient = new QueryClient();

describe("Edit Blog Page", () => {
  const mockPush = vi.fn();
  const mockRefresh = vi.fn();
  const mockUpdateBlog = vi.fn();
  const mockDeleteBlog = vi.fn();
  const mockAddTagsToPost = vi.fn();
  const mockRemoveTagsFromPost = vi.fn();

  beforeEach(() => {
    vi.clearAllMocks();
    (useParams as vi.Mock).mockReturnValue({ id: "1" });
    (useRouter as vi.Mock).mockReturnValue({ push: mockPush, refresh: mockRefresh });
    (useGetBlogById as vi.Mock).mockReturnValue({
      data: { title: "Existing Title", description: "Existing Description" },
      isPending: false,
      error: null,
    });
    (useGetTagsByPostID as vi.Mock).mockReturnValue({
      data: { tags: [{ id: 1, name: "tag1" }] },
    });
    (useUpdateBlog as vi.Mock).mockReturnValue({ mutate: mockUpdateBlog });
    (useDeleteBlog as vi.Mock).mockReturnValue({ mutate: mockDeleteBlog });
    (useAddTagsToPost as vi.Mock).mockReturnValue({ mutate: mockAddTagsToPost });
    (useRemoveTagsFromPost as vi.Mock).mockReturnValue({ mutate: mockRemoveTagsFromPost });
  });

  const renderEditPage = () => {
    render(
      <QueryClientProvider client={queryClient}>
        <Page />
      </QueryClientProvider>
    );
  };

  it("renders the form with existing data", async () => {
    renderEditPage();
    await waitFor(() => {
      expect(screen.getByLabelText(/タイトル/i)).toHaveValue("Existing Title");
      expect(screen.getByLabelText(/内容/i)).toHaveValue("Existing Description");
      expect(screen.getByLabelText(/タグ/i)).toHaveValue("tag1");
    });
  });

  it("calls updateBlog on form submission", async () => {
    renderEditPage();
    await userEvent.clear(screen.getByLabelText(/タイトル/i));
    await userEvent.type(screen.getByLabelText(/タイトル/i), "Updated Title");
    await userEvent.click(screen.getByRole("button", { name: /更新/i }));

    await waitFor(() => {
      expect(mockUpdateBlog).toHaveBeenCalledWith({
        id: 1,
        title: "Updated Title",
        description: "Existing Description",
      });
      expect(toast.success).toHaveBeenCalledWith("更新しました！", { duration: 1500 });
    });
  });

  it("calls deleteBlog on delete button click", async () => {
    renderEditPage();
    await userEvent.click(screen.getByRole("button", { name: /削除/i }));

    await waitFor(() => {
      expect(mockDeleteBlog).toHaveBeenCalledWith(1);
      expect(toast.success).toHaveBeenCalledWith("削除しました！");
    });
  });

  it("displays loading state", () => {
    (useGetBlogById as vi.Mock).mockReturnValue({ isPending: true });
    renderEditPage();
    expect(screen.getByText("Loading...")).toBeInTheDocument();
  });

  it("displays error state", () => {
    (useGetBlogById as vi.Mock).mockReturnValue({ error: new Error("Test Error"), isPending: false });
    renderEditPage();
    expect(screen.getByText("Error loading post.")).toBeInTheDocument();
  });
});
