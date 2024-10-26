import prisma from "@/app/lib/prisma";

const ENDPOINT = process.env.NEXT_PUBLIC_API_URL;

export const fetchAllBlogs = async () => {
  try {
    const res = await fetch(`${ENDPOINT}/api/blog`);

    if (!res.ok) {
      throw new Error(`Error fetching blogs: ${res.statusText}`);
    }

    const data = await res.json();
    return data.posts;
  } catch (error) {
    console.error("Error in fetchAllBlogs:", error);
    return null;
  }
};

export const fetchBlogsByPage = async (page: number) => {
  try {
    const res = await fetch(`${ENDPOINT}/api/blog/page/${page}`);

    if (!res.ok) {
      throw new Error(`Error fetching blogs by page: ${res.statusText}`);
    }

    const data = await res.json();
    return { posts: data.posts, totalCount: data.totalCount };
  } catch (error) {
    console.error("Error in fetchBlogsByPage:", error);
    return { posts: [], totalCount: 0 };
  }
};

export const countAllBlogs = async () => {
  try {
    const countData = await prisma.post.count();
    return countData;
  } catch (error) {
    console.error("Error in countAllBlogs:", error);
    return 0;
  }
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
    console.error("Error in editBlog:", error);
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
    console.error("Error in deleteBlog:", error);
    return null;
  }
};

export const getBlogById = async (id: number) => {
  try {
    const res = await fetch(`${ENDPOINT}/api/blog/${id}`);

    if (!res.ok) {
      throw new Error(`Error fetching blog by ID: ${res.statusText}`);
    }

    const data = await res.json();
    return data.post;
  } catch (error) {
    console.error("Error in getBlogById:", error);
    return null;
  }
};

export const getAllBlogIds = async () => {
  try {
    const res = await fetch(`${ENDPOINT}/api/blog`);

    if (!res.ok) {
      throw new Error(`Error fetching all blog IDs: ${res.statusText}`);
    }

    const data = await res.json();
    return data.posts.map((post: { id: number }) => post.id);
  } catch (error) {
    console.error("Error in getAllBlogIds:", error);
    return [];
  }
};
