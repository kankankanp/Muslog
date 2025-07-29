
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { useRouter } from "next/navigation";
import toast from "react-hot-toast";
import { describe, it, expect, vi, beforeEach } from "vitest";
import SignupForm from "./SignupForm";
import { useSignup } from "@/app/libs/hooks/api/useAuth";

vi.mock("next/navigation", () => ({
  useRouter: vi.fn(),
}));

vi.mock("@/app/libs/hooks/api/useAuth", () => ({
  useSignup: vi.fn(),
}));

vi.mock("react-hot-toast", () => ({
  success: vi.fn(),
  error: vi.fn(),
}));

const queryClient = new QueryClient();

describe("SignupForm", () => {
  const mockPush = vi.fn();
  const mockSignupMutation = vi.fn();

  beforeEach(() => {
    vi.clearAllMocks();
    (useRouter as vi.Mock).mockReturnValue({ push: mockPush });
    (useSignup as vi.Mock).mockReturnValue({ mutate: mockSignupMutation, isPending: false });
  });

  const renderSignupForm = () => {
    render(
      <QueryClientProvider client={queryClient}>
        <SignupForm />
      </QueryClientProvider>
    );
  };

  it("renders the signup form correctly", () => {
    renderSignupForm();
    expect(screen.getByRole("heading", { name: /新規登録/i })).toBeInTheDocument();
    expect(screen.getByLabelText(/名前/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/メールアドレス/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/パスワード/i)).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /登録する/i })).toBeInTheDocument();
  });

  it("displays validation errors for empty fields", async () => {
    renderSignupForm();
    await userEvent.click(screen.getByRole("button", { name: /登録する/i }));

    await waitFor(() => {
      expect(screen.getByText("名前を入力してください")).toBeInTheDocument();
      expect(screen.getByText("有効なメールアドレスを入力してください。")).toBeInTheDocument();
      expect(screen.getByText("パスワードは6文字以上で入力してください。")).toBeInTheDocument();
    });
  });

  it("calls signup mutation with correct data on successful submission", async () => {
    renderSignupForm();

    await userEvent.type(screen.getByLabelText(/名前/i), "Test User");
    await userEvent.type(screen.getByLabelText(/メールアドレス/i), "test@example.com");
    await userEvent.type(screen.getByLabelText(/パスワード/i), "password123");
    await userEvent.click(screen.getByRole("button", { name: /登録する/i }));

    await waitFor(() => {
      expect(mockSignupMutation).toHaveBeenCalledWith(
        {
          name: "Test User",
          email: "test@example.com",
          password: "password123",
        },
        expect.any(Object)
      );
    });
  });

  it("shows success toast and redirects on successful signup", async () => {
    mockSignupMutation.mockImplementation((_, { onSuccess }) => onSuccess());
    renderSignupForm();

    await userEvent.type(screen.getByLabelText(/名前/i), "Test User");
    await userEvent.type(screen.getByLabelText(/メールアドレス/i), "test@example.com");
    await userEvent.type(screen.getByLabelText(/パスワード/i), "password123");
    await userEvent.click(screen.getByRole("button", { name: /登録する/i }));

    await waitFor(() => {
      expect(toast.success).toHaveBeenCalledWith("登録しました。");
      expect(mockPush).toHaveBeenCalledWith("/dashboard");
    });
  });

  it("shows error toast on failed signup", async () => {
    mockSignupMutation.mockImplementation((_, { onError }) => onError(new Error()));
    renderSignupForm();

    await userEvent.type(screen.getByLabelText(/名前/i), "Test User");
    await userEvent.type(screen.getByLabelText(/メールアドレス/i), "test@example.com");
    await userEvent.type(screen.getByLabelText(/パスワード/i), "password123");
    await userEvent.click(screen.getByRole("button", { name: /登録する/i }));

    await waitFor(() => {
      expect(toast.error).toHaveBeenCalledWith("登録に失敗しました。");
    });
  });
});
