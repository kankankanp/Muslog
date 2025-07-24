import { Track } from "@prisma/client";

export const ENDPOINT = "http://localhost:8080";

export const fetchAllBlogs = async () => {
  try {
    const res = await fetch(`${ENDPOINT}/blogs`, { cache: "no-store" });

    if (!res.ok) {
      throw new Error(`Error fetching blogs: ${res.statusText}`);
    }

    const data = await res.json();
    return data.posts;
  } catch (error) {
    return [];
  }
};

export const fetchBlogsByPage = async (page: number) => {
  try {
    const res = await fetch(`${ENDPOINT}/blogs/page/${page}`, { cache: "no-store" });

    if (!res.ok) {
      throw new Error(`Error fetching blogs: ${res.statusText}`);
    }

    const data = await res.json();
    return data;
  } catch (error) {
    return { posts: [], totalCount: 0 };
  }
};

export const postBlog = async (
  title: string,
  description: string,
  track: Track | null,
  userId: string
) => {
  const res = await fetch(`${ENDPOINT}/blogs`, {
    method: "POST",
    body: JSON.stringify({
      title,
      description,
      userId,
      tracks: track ? [track] : [],
    }),
    headers: {
      "Content-Type": "application/json",
    },
  });

  return res.json();
};

export const editBlog = async (
  title: string | undefined,
  description: string | undefined,
  id: number
) => {
  try {
    const res = await fetch(`${ENDPOINT}/blogs/${id}`, {
      method: "PUT",
      body: JSON.stringify({ title, description }),
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      throw new Error(`Error editing blog: ${res.statusText}`);
    }

    return await res.json();
  } catch (error) {
    return null;
  }
};

export const deleteBlog = async (id: number) => {
  try {
    const res = await fetch(`${ENDPOINT}/blogs/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      throw new Error(`Error deleting blog: ${res.statusText}`);
    }

    return await res.json();
  } catch (error) {
    return null;
  }
};

export const getBlogById = async (id: number) => {
  try {
    const res = await fetch(`${ENDPOINT}/blogs/${id}`, {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    });
    if (!res.ok) {
      throw new Error(`Error fetching blog by ID: ${res.statusText}`);
    }

    const data = await res.json();

    return data.post;
  } catch (error) {
    return null;
  }
};

export const getAllBlogIds = async () => {
  try {
    const res = await fetch(`${ENDPOINT}/blogs`, { cache: "no-store" });
    if (!res.ok) {
      throw new Error(`Error fetching blogs: ${res.statusText}`);
    }
    const data = await res.json();
    return data.posts.map((post: { id: number }) => post.id);
  } catch (error) {
    return [];
  }
};
