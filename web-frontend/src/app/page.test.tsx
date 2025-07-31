import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import Page from "./page";

// Mock framer-motion components to avoid actual animations in tests
vi.mock("framer-motion", () => ({
  motion: {
    div: vi.fn(({ children }) => (
      <div data-testid="motion-div">{children}</div>
    )),
    h1: vi.fn(({ children }) => <h1 data-testid="motion-h1">{children}</h1>),
    p: vi.fn(({ children }) => <p data-testid="motion-p">{children}</p>),
  },
}));

describe("Landing Page", () => {
  it("renders the main sections and content", () => {
    render(<Page />);

    // Hero section
    expect(
      screen.getByText(/あなたの音楽の物語をシェアしよう。/i)
    ).toBeInTheDocument();
    expect(
      screen.getByText(/好きな曲と共に、あなたの想いを綴る。/i)
    ).toBeInTheDocument();
    const startButton = screen.getByRole("link", { name: /はじめる/i });
    expect(startButton).toBeInTheDocument();
    expect(startButton).toHaveAttribute("href", "/dashboard");

    // Features section
    expect(screen.getByText("3つの特徴")).toBeInTheDocument();
    expect(screen.getByText("音楽との出会い")).toBeInTheDocument();
    expect(screen.getByText("ブログ機能")).toBeInTheDocument();
    expect(screen.getByText("コミュニティ")).toBeInTheDocument();

    // 3 Steps section
    expect(screen.getByText("簡単3ステップ")).toBeInTheDocument();
    expect(screen.getByText("アカウント作成")).toBeInTheDocument();
    expect(screen.getByText("音楽を選択")).toBeInTheDocument();
    expect(screen.getByText("ブログを投稿")).toBeInTheDocument();

    // Recommendations section
    expect(screen.getByText("こんな方におすすめ")).toBeInTheDocument();
    expect(screen.getByText("音楽が好きな方")).toBeInTheDocument();
    expect(screen.getByText("思い出を残したい方")).toBeInTheDocument();
  });
});
