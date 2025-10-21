'use client';

export default function Skeleton() {
  return (
    <>
      <header className="bg-gray-800 text-white px-4 flex items-center justify-between h-16">
        <div className="flex items-center gap-4">
          <div className="w-[60px] h-[60px] bg-gray-700 rounded-full animate-pulse"></div>
          <div className="h-8 w-32 bg-gray-700 rounded animate-pulse"></div>
        </div>
      </header>
      <div className="flex h-screen">
        <aside className="h-screen w-64 bg-white border-r border-gray-200 p-4 hidden md:block">
          <nav>
            <ul className="flex flex-col gap-6 mt-8">
              {[1, 2, 3, 4, 5, 6].map((item) => (
                <li key={item} className="mb-2">
                  <div className="flex items-center p-2 rounded-lg hover:bg-gray-100">
                    <div className="mr-3 h-5 w-5 bg-gray-200 rounded animate-pulse"></div>
                    <div className="h-4 w-20 bg-gray-200 rounded animate-pulse"></div>
                  </div>
                </li>
              ))}
            </ul>
          </nav>
        </aside>
        <main className="flex-1 p-8 overflow-y-auto w-full">
          <div className="text-3xl font-bold border-gray-100 border-b-2 bg-gray-200 p-6 mb-6 h-10 w-1/3 rounded animate-pulse"></div>
          <div className="space-y-4">
            {[1, 2, 3].map((item) => (
              <div
                key={item}
                className="h-6 bg-gray-200 rounded animate-pulse w-full"
              ></div>
            ))}
          </div>
        </main>
      </div>
    </>
  );
}
