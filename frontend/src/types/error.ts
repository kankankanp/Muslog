export type ErrorBase = {
  message: string;
  code?: number;
  cause?: unknown;
  description?: string;
};

export type ValidationError = ErrorBase & {
  fields: Record<string, string>;
};

export type AuthError = ErrorBase & {
  reason: "expired" | "invalid" | "forbidden";
};

export type NetworkError = ErrorBase & {
  network: boolean;
  status?: number;
  retryable?: boolean;
};

export type PermissionError = ErrorBase & {
  permission: "denied" | "restricted";
};

export type AppError =
  | ErrorBase
  | ValidationError
  | AuthError
  | NetworkError
  | PermissionError;

export type ErrorMessageProps = Partial<Omit<ErrorBase, "cause">> & {
  className?: string;
};
