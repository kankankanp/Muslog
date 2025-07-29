
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { useRouter } from "next/navigation";
import toast from "react-hot-toast";
import { describe, it, expect, vi, beforeEach } from "vitest";
import NewBlogForm from "./NewBlogForm";
import { useCreateBlog } from "@/app/libs/hooks/api/useBlogs";
import { useAddTagsToPost } from "@/app/libs/hooks/api/useTags";

vi.mock("next/navigation", () => ({
  useRouter: vi.fn(),
}));

vi.mock("@/app/libs/hooks/api/useBlogs", () => ({
  useCreateBlog: vi.fn(),
}));

vi.mock("@/app/libs/hooks/api/useTags", () => ({
  useAddTagsToPost: vi.fn(),
}));

vi.mock("react-hot-toast", () => ({
  success: vi.fn(),
  error: vi.fn(),
}));

const queryClient = new QueryClient();

describe("NewBlogForm", () => {
  const mockPush = vi.fn();
  const mockCreateBlogMutation = vi.fn();
  const mockAddTagsToPostMutation = vi.fn();

  beforeEach(() => {
    vi.clearAllMocks();
    (useRouter as vi.Mock).mockReturnValue({ push: mockPush });
    (useCreateBlog as vi.Mock).mockReturnValue({ mutate: mockCreateBlogMutation });
    (useAddTagsToPost as vi.Mock).mockReturnValue({ mutate: mockAddTagsToPostMutation });
  });

  const renderNewBlogForm = (selectedTrack = null) => {
    render(
      <QueryClientProvider client={queryClient}>
        <NewBlogForm selectedTrack={selectedTrack} />
      </QueryClientProvider>
    );
  };

  it("renders the form correctly", () => {
    renderNewBlogForm();
    expect(screen.getByLabelText(/タイトル/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/内容/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/タグ/i)).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /Post/i })).toBeInTheDocument();
  });

  it("displays validation errors for empty fields", async () => {
    renderNewBlogForm();
    await userEvent.click(screen.getByRole("button", { name: /Post/i }));

    await waitFor(() => {
      expect(screen.getByText("タイトルを入力してください")).toBeInTheDocument();
      expect(screen.getByText("内容を入力してください")).toBeInTheDocument();
    });
  });

  it("calls createBlogMutation on successful submission", async () => {
    renderNewBlogForm();
    await userEvent.type(screen.getByLabelText(/タイトル/i), "Test Title");
    await userEvent.type(screen.getByLabelText(/内容/i), "Test Content");
    await userEvent.click(screen.getByRole("button", { name: /Post/i }));

    await waitFor(() => {
      expect(mockCreateBlogMutation).toHaveBeenCalled();
    });
  });

  it("calls addTagsToPostMutation when tags are provided", async () => {
    mockCreateBlogMutation.mockImplementation((_, { onSuccess }) => onSuccess({ id: "1" }));
    renderNewBlogForm();
    await userEvent.type(screen.getByLabelText(/タイトル/i), "Test Title");
    await userEvent.type(screen.getByLabelText(/内容/i), "Test Content");
    await userEvent.type(screen.getByLabelText(/タグ/i), "tag1, tag2");
    await userEvent.click(screen.getByRole("button", { name: /Post/i }));

    await waitFor(() => {
      expect(mockAddTagsToPostMutation).toHaveBeenCalledWith({
        postID: "1",
        requestBody: { tag_names: ["tag1", "tag2"] },
      });
    });
  });

  it("displays the selected track", () => {
    const selectedTrack = {
      id: 1,
      spotifyId: "1",
      name: "Test Track",
      artistName: "Test Artist",
      albumImageUrl: "/test.jpg",
    };
    renderNewBlogForm(selectedTrack);
    expect(screen.getByText("Test Track")).toBeInTheDocument();
    expect(screen.getByText("Test Artist")).toBeInTheDocument();
  });
});
