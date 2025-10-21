// src/components/elements/forms/LoginForm.tsx
'use client';

import { Suspense } from 'react';
import LoginForm from './LoginForm';
import Spinner from '@/components/layouts/Spinner';

export default function LoginFormWrapper() {
  return (
    <Suspense fallback={<Spinner />}>
      <LoginForm />
    </Suspense>
  );
}
