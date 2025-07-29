
import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { describe, it, expect, vi } from "vitest";
import { Book } from "./Book";
import { Post } from "@/app/libs/api/generated";

const mockPosts: Post[] = [
  {
    id: "1",
    title: "Test Post 1",
    description: "Test Description 1",
    createdAt: new Date(),
    updatedAt: new Date(),
    userId: "1",
    isLiked: false,
    likesCount: 0,
    tracks: [],
  },
  {
    id: "2",
    title: "Test Post 2",
    description: "Test Description 2",
    createdAt: new Date(),
    updatedAt: new Date(),
    userId: "1",
    isLiked: false,
    likesCount: 0,
    tracks: [],
  },
];

describe("Book", () => {
  it("renders the book with posts", () => {
    render(<Book posts={mockPosts} />);
    expect(screen.getByText("Subject: Test Post 1")).toBeInTheDocument();
    expect(screen.getByText("Test Description 1")).toBeInTheDocument();
  });

  it("plays audio on page flip", async () => {
    const playSpy = vi.spyOn(window.HTMLMediaElement.prototype, 'play').mockImplementation(() => {});
    render(<Book posts={mockPosts} />);
    const checkboxes = screen.getAllByRole("checkbox");
    await userEvent.click(checkboxes[1]);
    expect(playSpy).toHaveBeenCalled();
    playSpy.mockRestore();
  });
});
