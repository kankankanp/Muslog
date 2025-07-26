import { useMutation } from "@tanstack/react-query";
import { AuthService } from "@/app/libs/api/generated";
import { User } from "@/app/libs/api/generated/models/User";
import { UserLogin } from "@/app/libs/api/generated/models/UserLogin";

export const useLogin = () => {
  const { mutate, isPending, error } = useMutation<User, Error, UserLogin>({
    mutationFn: async (credentials) => {
      const response = await AuthService.postLogin(credentials);
      return response.user!;
    },
  });
  return { mutate, isPending, error };
};

export const useLogout = () => {
  const { mutate, isPending, error } = useMutation<void, Error>({
    mutationFn: async () => {
      await AuthService.postLogout();
    },
  });
  return { mutate, isPending, error };
};

export const useRefreshToken = () => {
  const { mutate, isPending, error } = useMutation<void, Error>({
    mutationFn: async () => {
      await AuthService.postRefresh();
    },
  });
  return { mutate, isPending, error };
};

export const useSignup = () => {
  const { mutate, data, isPending, error } = useMutation<User, Error, { name: string; email: string; password: string }>({
    mutationFn: async (credentials) => {
      const response = await AuthService.postRegister(credentials);
      return response.user!;
    },
  });
  return { mutate, data, isPending, error };
};
