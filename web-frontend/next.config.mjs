/** @type {import('next').NextConfig} */
const nextConfig = {
  eslint: {
    ignoreDuringBuilds: true,
  },
  webpack: (config) => {
    config.module.rules.push({
      test: /\.html$/,
      use: "raw-loader",
    });

    return config;
  },
  experimental: {
    serverComponentsExternalPackages: ["bcrypt"],
    reactCompiler: true,
  },
  images: {
    domains: ["i.scdn.co"],
  },
};

export default nextConfig;
