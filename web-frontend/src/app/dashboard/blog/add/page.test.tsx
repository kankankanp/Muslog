
import { render, screen } from "@testing-library/react";
import { describe, it, expect, vi } from "vitest";
import Page from "./page";
import NewBlogForm from "@/app/components/elements/forms/NewBlogForm";
import SelectMusicArea from "@/app/components/elements/others/SelectMusciArea";

vi.mock("@/app/components/elements/forms/NewBlogForm", () => ({
  __esModule: true,
  default: vi.fn(() => <div data-testid="new-blog-form">New Blog Form</div>),
}));

vi.mock("@/app/components/elements/others/SelectMusciArea", () => ({
  __esModule: true,
  default: vi.fn(() => <div data-testid="select-music-area">Select Music Area</div>),
}));

describe("Add Blog Page", () => {
  it("renders SelectMusicArea and NewBlogForm", () => {
    render(<Page />);
    expect(screen.getByTestId("select-music-area")).toBeInTheDocument();
    expect(screen.getByTestId("new-blog-form")).toBeInTheDocument();
  });
});
