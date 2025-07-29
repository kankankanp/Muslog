
import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import toast from "react-hot-toast";
import { describe, it, expect, vi, beforeEach } from "vitest";
import SelectMusicArea, { Track } from "./SelectMusciArea";

vi.mock("react-hot-toast", () => ({
  error: vi.fn(),
}));

global.fetch = vi.fn();

describe("SelectMusicArea", () => {
  const mockOnSelect = vi.fn();

  beforeEach(() => {
    vi.clearAllMocks();
  });

  const renderSelectMusicArea = () => {
    render(<SelectMusicArea onSelect={mockOnSelect} />);
  };

  it("renders the search area correctly", () => {
    renderSelectMusicArea();
    expect(screen.getByPlaceholderText("曲名を入力")).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /検索/i })).toBeInTheDocument();
  });

  it("shows an error toast if the search query is empty", async () => {
    renderSelectMusicArea();
    await userEvent.click(screen.getByRole("button", { name: /検索/i }));
    expect(toast.error).toHaveBeenCalledWith("検索ワードを入力してください");
  });

  it("fetches and displays tracks on search", async () => {
    const mockTracks: Track[] = [
      {
        id: 1,
        spotifyId: "1",
        name: "Test Track",
        artistName: "Test Artist",
        albumImageUrl: "/test.jpg",
      },
    ];
    (fetch as vi.Mock).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ tracks: mockTracks }),
    });

    renderSelectMusicArea();
    await userEvent.type(screen.getByPlaceholderText("曲名を入力"), "test");
    await userEvent.click(screen.getByRole("button", { name: /検索/i }));

    await waitFor(() => {
      expect(screen.getByText("Test Track")).toBeInTheDocument();
    });
  });

  it("calls onSelect when a track is clicked", async () => {
    const mockTracks: Track[] = [
      {
        id: 1,
        spotifyId: "1",
        name: "Test Track",
        artistName: "Test Artist",
        albumImageUrl: "/test.jpg",
      },
    ];
    (fetch as vi.Mock).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ tracks: mockTracks }),
    });

    renderSelectMusicArea();
    await userEvent.type(screen.getByPlaceholderText("曲名を入力"), "test");
    await userEvent.click(screen.getByRole("button", { name: /検索/i }));

    await waitFor(async () => {
      await userEvent.click(screen.getByText("Test Track"));
      expect(mockOnSelect).toHaveBeenCalledWith(mockTracks[0]);
    });
  });
});
