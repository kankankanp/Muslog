
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { describe, it, expect, vi } from "vitest";
import BlogCard from "./BlogCard";
import { Post } from "@/app/libs/api/generated";

const mockPosts: Post[] = [
  {
    id: "1",
    title: "Test Post",
    description: "This is a long description that should be truncated.",
    tags: [{ id: "1", name: "tag1" }],
    tracks: [
      {
        spotifyId: "1",
        name: "Track 1",
        artistName: "Artist 1",
        albumImageUrl: "/test.jpg",
      },
    ],
    likesCount: 10,
    isLiked: false,
    userId: "1",
    createdAt: new Date(),
    updatedAt: new Date(),
  },
];

describe("BlogCard", () => {
  it("renders blog cards correctly", () => {
    render(<BlogCard posts={mockPosts} />);
    expect(screen.getByText("Test Post")).toBeInTheDocument();
    expect(screen.getByText(/This is a long description/)).toBeInTheDocument();
  });

  it("truncates description when not on detail page", () => {
    render(<BlogCard posts={mockPosts} />);
    expect(screen.getByText("This is a long description that should...")).toBeInTheDocument();
  });

  it("shows full description on detail page", () => {
    render(<BlogCard posts={mockPosts} isDetailPage />);
    expect(screen.getByText("This is a long description that should be truncated.")).toBeInTheDocument();
  });

  it("calls onLikeClick when like icon is clicked", async () => {
    const onLikeClick = vi.fn();
    render(<BlogCard posts={mockPosts} onLikeClick={onLikeClick} />);
    await userEvent.click(screen.getByTestId("like-icon"));
    expect(onLikeClick).toHaveBeenCalled();
  });

  it("shows 'Show more' button when not on detail page", () => {
    render(<BlogCard posts={mockPosts} />);
    expect(screen.getByRole("link", { name: /Show more/i })).toBeInTheDocument();
  });

  it("shows 'Back' button on detail page", () => {
    render(<BlogCard posts={mockPosts} isDetailPage />);
    expect(screen.getByRole("link", { name: /Back/i })).toBeInTheDocument();
  });
});
