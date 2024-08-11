import { PrismaClient } from "@prisma/client";

export const fetchAllBlogs = async () => {
  const res = await fetch("http://localhost:3000/api/blog", {
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
