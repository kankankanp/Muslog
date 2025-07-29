import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { useRouter } from "next/navigation";
import toast from "react-hot-toast";
import { useDispatch } from "react-redux";
import { describe, it, expect, vi, beforeEach } from "vitest";
import LoginForm from "./LoginForm";
import { useLogin } from "@/app/libs/hooks/api/useAuth";

// Mock next/navigation
vi.mock("next/navigation", () => ({
  useRouter: vi.fn(),
}));

// Mock react-redux
vi.mock("react-redux", () => ({
  useDispatch: vi.fn(),
}));

// Mock @/app/libs/hooks/api/useAuth
vi.mock("@/app/libs/hooks/api/useAuth", () => ({
  useLogin: vi.fn(),
}));

// Mock react-hot-toast
vi.mock("react-hot-toast", () => ({
  success: vi.fn(),
  error: vi.fn(),
}));

const queryClient = new QueryClient();

describe("LoginForm", () => {
  const mockPush = vi.fn();
  const mockDispatch = vi.fn();
  const mockLoginMutation = vi.fn();

  beforeEach(() => {
    vi.clearAllMocks();
    (useRouter as vi.Mock).mockReturnValue({
      push: mockPush,
    });
    (useDispatch as vi.Mock).mockReturnValue(mockDispatch);
    (useLogin as vi.Mock).mockReturnValue({
      mutate: mockLoginMutation,
      isPending: false,
    });
  });

  const renderLoginForm = () => {
    render(
      <QueryClientProvider client={queryClient}>
        <LoginForm />
      </QueryClientProvider>
    );
  };

  it("renders the login form correctly", () => {
    renderLoginForm();
    expect(
      screen.getByRole("heading", { name: /ログイン/i })
    ).toBeInTheDocument();
    expect(screen.getByLabelText(/メールアドレス:/i)).toBeInTheDocument();
    expect(screen.getByLabelText(/パスワード:/i)).toBeInTheDocument();
    expect(
      screen.getByRole("button", { name: /ログイン/i })
    ).toBeInTheDocument();
  });

  it("displays validation errors for empty fields", async () => {
    renderLoginForm();
    await userEvent.click(screen.getByRole("button", { name: /ログイン/i }));

    await waitFor(() => {
      expect(
        screen.getByText("有効なメールアドレスを入力してください。")
      ).toBeInTheDocument();
      expect(
        screen.getByText("パスワードは6文字以上で入力してください。")
      ).toBeInTheDocument();
    });
  });

  it("displays validation error for invalid email format", async () => {
    renderLoginForm();
    await userEvent.type(
      screen.getByLabelText(/メールアドレス:/i),
      "invalid-email"
    );
    await userEvent.click(screen.getByRole("button", { name: /ログイン/i }));

    await waitFor(() => {
      expect(
        screen.getByText("有効なメールアドレスを入力してください。")
      ).toBeInTheDocument();
    });
  });

  it("calls login mutation with correct data on successful submission", async () => {
    renderLoginForm();

    await userEvent.type(
      screen.getByLabelText(/メールアドレス:/i),
      "test@example.com"
    );
    await userEvent.type(screen.getByLabelText(/パスワード:/i), "password123");
    await userEvent.click(screen.getByRole("button", { name: /ログイン/i }));

    await waitFor(() => {
      expect(mockLoginMutation).toHaveBeenCalledWith(
        {
          email: "test@example.com",
          password: "password123",
        },
        expect.any(Object)
      );
    });
  });

  it("shows success toast and redirects on successful login", async () => {
    mockLoginMutation.mockImplementation((_data, { onSuccess }) => {
      onSuccess({ id: "1", email: "test@example.com" });
    });

    renderLoginForm();

    await userEvent.type(
      screen.getByLabelText(/メールアドレス:/i),
      "test@example.com"
    );
    await userEvent.type(screen.getByLabelText(/パスワード:/i), "password123");
    await userEvent.click(screen.getByRole("button", { name: /ログイン/i }));

    await waitFor(() => {
      expect(toast.success).toHaveBeenCalledWith("ログインに成功しました");
      expect(mockPush).toHaveBeenCalledWith("/dashboard");
      expect(mockDispatch).toHaveBeenCalledWith({
        payload: { id: "1", email: "test@example.com" },
        type: "auth/login",
      });
    });
  });

  it("shows error toast on failed login", async () => {
    mockLoginMutation.mockImplementation((_data, { onError }) => {
      onError(new Error("Invalid credentials"));
    });

    renderLoginForm();

    await userEvent.type(
      screen.getByLabelText(/メールアドレス:/i),
      "test@example.com"
    );
    await userEvent.type(screen.getByLabelText(/パスワード:/i), "password123");
    await userEvent.click(screen.getByRole("button", { name: /ログイン/i }));

    await waitFor(() => {
      expect(toast.error).toHaveBeenCalledWith("Invalid credentials");
      expect(mockPush).not.toHaveBeenCalled();
      expect(mockDispatch).not.toHaveBeenCalled();
    });
  });

  it("handles copy to clipboard for example email", async () => {
    Object.defineProperty(navigator, "clipboard", {
      value: {
        writeText: vi.fn().mockResolvedValue(undefined),
      },
      writable: true,
    });

    renderLoginForm();

    const emailCopyButton = screen.getByRole("button", {
      name: "コピー",
      exact: false,
    });
    await userEvent.click(emailCopyButton);

    await waitFor(() => {
      expect(navigator.clipboard.writeText).toHaveBeenCalledWith(
        "EygQJpu@NillQOs.net"
      );
    });
  });

  it("handles copy to clipboard for example password", async () => {
    Object.defineProperty(navigator, "clipboard", {
      value: {
        writeText: vi.fn().mockResolvedValue(undefined),
      },
      writable: true,
    });

    renderLoginForm();

    const passwordCopyButton = screen.getAllByRole("button", {
      name: "コピー",
      exact: false,
    })[1]; // Get the second copy button
    await userEvent.click(passwordCopyButton);

    await waitFor(() => {
      expect(navigator.clipboard.writeText).toHaveBeenCalledWith("password");
    });
  });
});
