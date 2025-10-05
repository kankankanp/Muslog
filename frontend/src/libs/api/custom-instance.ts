import Axios, { AxiosRequestConfig } from "axios";
import { handleApiError } from "./interceptor";

const instance = Axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api/v1",
  withCredentials: true,
});

// レスポンスインターセプターを設定
instance.interceptors.response.use((response) => response, handleApiError);

export const customInstance = <T>(config: AxiosRequestConfig): Promise<T> => {
  const source = Axios.CancelToken.source();
  const promise = instance
    .request({ ...config, cancelToken: source.token })
    .then(({ data }) => data);

  // @ts-ignore
  promise.cancel = () => {
    source.cancel("Query was cancelled");
  };

  return promise;
};

export default customInstance;
