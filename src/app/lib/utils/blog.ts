import { Track } from "@prisma/client";
import { auth } from "../auth/auth";
import prisma from "../db/prisma";

const ENDPOINT = process.env.NEXT_PUBLIC_API_URL;

export const fetchAllBlogs = async () => {
  try {
    const res = await fetch("/api/blog", { cache: "no-store" });

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
  const session = await auth();

  if (!session?.user?.id) {
    return { posts: [], totalCount: 0 };
  }

  const PER_PAGE = 4;
  const skip = (page - 1) * PER_PAGE;

  const [posts, totalCount] = await Promise.all([
    prisma.post.findMany({
      where: { userId: session.user.id },
      skip,
      take: PER_PAGE,
      orderBy: { createdAt: "desc" },
      include: {
        tracks: true,
      },
    }),
    prisma.post.count({
      where: { userId: session.user.id },
    }),
  ]);

  return { posts, totalCount };
};

export const countAllBlogs = async () => {
  try {
    const countData = await prisma.post.count();
    return countData;
  } catch (error) {
    return 0;
  }
};

export const postBlog = async (
  title: string,
  description: string,
  track: Track | null
) => {
  const res = await fetch(`${ENDPOINT}/api/blog`, {
    method: "POST",
    body: JSON.stringify({
      title,
      description,
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
    const res = await fetch(`${ENDPOINT}/api/blog/${id}`, {
      method: "PUT",
      body: JSON.stringify({ title, description, id }),
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
    const res = await fetch(`${ENDPOINT}/api/blog/${id}`, {
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
    const res = await fetch(`${ENDPOINT}/api/blog/${id}`, {
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
    const posts = await prisma.post.findMany({
      select: { id: true },
    });
    return posts.map((post) => post.id);
  } catch (error) {
    return [];
  }
};
