
import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { useRouter } from "next/navigation";
import toast from "react-hot-toast";
import { useDispatch } from "react-redux";
import { describe, it, expect, vi, beforeEach } from "vitest";
import LogoutButton from "./LogoutButton";
import { useLogout } from "@/app/libs/hooks/api/useAuth";

vi.mock("next/navigation", () => ({
  useRouter: vi.fn(),
}));

vi.mock("react-redux", () => ({
  useDispatch: vi.fn(),
}));

vi.mock("@/app/libs/hooks/api/useAuth", () => ({
  useLogout: vi.fn(),
}));

vi.mock("react-hot-toast", () => ({
  success: vi.fn(),
  error: vi.fn(),
}));

describe("LogoutButton", () => {
  const mockPush = vi.fn();
  const mockDispatch = vi.fn();
  const mockLogoutMutation = vi.fn();

  beforeEach(() => {
    vi.clearAllMocks();
    (useRouter as vi.Mock).mockReturnValue({ push: mockPush });
    (useDispatch as vi.Mock).mockReturnValue(mockDispatch);
    (useLogout as vi.Mock).mockReturnValue({ mutate: mockLogoutMutation });
  });

  it("renders the logout button", () => {
    render(<LogoutButton />);
    expect(screen.getByRole("button", { name: /ログアウト/i })).toBeInTheDocument();
  });

  it("calls logout mutation on click", async () => {
    render(<LogoutButton />);
    await userEvent.click(screen.getByRole("button", { name: /ログアウト/i }));
    expect(mockLogoutMutation).toHaveBeenCalled();
  });

  it("shows success toast and redirects on successful logout", async () => {
    mockLogoutMutation.mockImplementation((_, { onSuccess }) => onSuccess());
    render(<LogoutButton />);
    await userEvent.click(screen.getByRole("button", { name: /ログアウト/i }));

    await waitFor(() => {
      expect(toast.success).toHaveBeenCalledWith("ログアウトしました。");
      expect(mockPush).toHaveBeenCalledWith("/login");
      expect(mockDispatch).toHaveBeenCalled();
    });
  });

  it("shows error toast on failed logout", async () => {
    mockLogoutMutation.mockImplementation((_, { onError }) => onError(new Error()));
    render(<LogoutButton />);
    await userEvent.click(screen.getByRole("button", { name: /ログアウト/i }));

    await waitFor(() => {
      expect(toast.error).toHaveBeenCalledWith("ログアウトに失敗しました。");
    });
  });
});
