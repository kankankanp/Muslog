import { prisma } from "../api/blog/route";

const ENDPOINT = process.env.NEXT_PUBLIC_API_URL;

export const fetchAllBlogs = async () => {
  const res = await fetch(`${ENDPOINT}/api/blog`, {
    cache: "no-store",
  });

  const data = await res.json();

  return data.posts;
};

export const fetchBlogsByPage = async (page: number) => {
  const res = await fetch(`${ENDPOINT}/api/blog/page/${page}`, {
    cache: "no-store",
  });

  const data = await res.json();

  return { posts: data.posts, totalCount: data.totalCount };
};

export const countAllBlogs = async () => {
  const countData = await prisma.post.count();
  return countData;
};

export const editBlog = async (
  title: string | undefined,
  description: string | undefined,
  id: number
) => {
  const res = await fetch(`${ENDPOINT}/api/blog/${id}`, {
    method: "PUT",
    body: JSON.stringify({ title, description, id }),
    headers: {
      "Content-Type": "application/json",
    },
  });

  return res.json();
};

export const deleteBlog = async (id: number) => {
  const res = await fetch(`${ENDPOINT}/api/blog/${id}`, {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  });

  return res.json();
};

export const getBlogById = async (id: number) => {
  const res = await fetch(`${ENDPOINT}/api/blog/${id}`);
  const data = await res.json();
  return data.post;
};
