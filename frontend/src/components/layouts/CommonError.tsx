import type { AppError, ValidationError, AuthError } from '@/types/error';

export function CommonError(props: Partial<AppError> & { className?: string }) {
  const { message, code, className } = props;

  // 型ガード関数（引数型をunknownにする）
  function hasFields(obj: unknown): obj is ValidationError {
    return (
      typeof obj === 'object' &&
      obj !== null &&
      'fields' in obj &&
      typeof (obj as any).fields === 'object'
    );
  }
  function hasReason(obj: unknown): obj is AuthError {
    return (
      typeof obj === 'object' &&
      obj !== null &&
      'reason' in obj &&
      typeof (obj as any).reason === 'string'
    );
  }

  return (
    <div className={className ?? 'text-red-600 dark:text-red-400'}>
      <strong>エラー:</strong> {message ?? '不明なエラー'}
      {code && <span> (コード: {code})</span>}
      {hasFields(props) && (
        <ul>
          {Object.entries(props.fields).map(([field, msg]) => (
            <li key={field}>
              {field}: {msg}
            </li>
          ))}
        </ul>
      )}
      {hasReason(props) && <div>理由: {props.reason}</div>}
    </div>
  );
}
