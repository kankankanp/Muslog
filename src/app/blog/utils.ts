import { PrismaClient } from "@prisma/client";

export const fetchAllBlogs = async () => {
  const res = await fetch("https://my-next-blog-iota-six.vercel.app/api/blog", {
    cache: "no-store",
  });

  const data = await res.json();

  return data.posts;
};

export const countAllBlogs = async () => {
  const prisma = new PrismaClient();
  const countData = await prisma.post.count();
  return countData;
};
