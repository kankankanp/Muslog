
import { render, screen } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import ProfileCard from "./ProfileCard";

describe("ProfileCard", () => {
  it("renders the profile card with the correct name and email", () => {
    render(<ProfileCard name="Test User" email="test@example.com" />);
    expect(screen.getByText("PROFILE")).toBeInTheDocument();
    expect(screen.getByText("Test User")).toBeInTheDocument();
    expect(screen.getByText("test@example.com")).toBeInTheDocument();
  });
});
