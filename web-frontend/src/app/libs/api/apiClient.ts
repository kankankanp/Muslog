// libs/apiClient.ts

import { Configuration, BlogsApi, UsersApi } from "./generated";

const config = new Configuration({
  basePath: "http://localhost:8080", // APIエンドポイント
});

export const blogsApi = new BlogsApi(config);
export const usersApi = new UsersApi(config);
export * from "./generated";
