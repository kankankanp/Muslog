import { useMutation } from "@tanstack/react-query";
import {
  AuthResponse,
  AuthService,
  LoginRequest,
  RegisterRequest,
} from "@/app/libs/api/generated";

export const useLogin = () => {
  const { mutate, isPending, error } = useMutation<
    AuthResponse,
    Error,
    LoginRequest
  >({
    mutationFn: async (credentials) => {
      const loginResponse = await AuthService.postLogin(credentials);
      return loginResponse.user!;
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
  const { mutate, isPending, error } = useMutation<
    AuthResponse,
    Error,
    RegisterRequest
  >({
    mutationFn: async (credentials) => {
      const response = await AuthService.postRegister(credentials);
      return response.user!;
    },
  });
  return { mutate, isPending, error };
};
