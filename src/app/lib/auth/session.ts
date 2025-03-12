export const fetchSession = async () => {
  try {
    const res = await fetch("/api/auth/session");
    if (!res.ok) throw new Error("Failed to fetch session");
    const data = await res.json();
    return data;
  } catch (error) {
    console.error("Error fetching session:", error);
    return null;
  }
};
