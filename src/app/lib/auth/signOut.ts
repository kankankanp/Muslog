"use server";

import { signOut } from "@/app/lib/auth/auth";

export async function handleSignOut() {
  await signOut();
}
