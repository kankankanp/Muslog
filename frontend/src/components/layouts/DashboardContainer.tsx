"use client";

import React, { useState, ReactNode } from "react";
import Header from "@/components/layouts/Header";
import Sidebar from "@/components/layouts/Sidebar";
import { SidebarProvider } from "@/contexts/SidebarContext";

export default function DashboardContainer({
  children,
}: {
  children: ReactNode;
}) {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  return (
    <SidebarProvider>
      <Header />
      <div className="flex h-screen">
        <Sidebar />
        {isSidebarOpen && (
          <div
            className="fixed inset-0 bg-black bg-opacity-50 z-40 md:hidden"
            onClick={() => setIsSidebarOpen(false)}
          ></div>
        )}
        <main className="flex-1 p-8 overflow-y-auto w-full">{children}</main>
      </div>
    </SidebarProvider>
  );
}
