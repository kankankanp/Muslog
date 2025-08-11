/** @type {import('next').NextConfig} */
const nextConfig = {
  output: 'export',
  async redirects() {
    return [
      {
        source: '/',
        destination: '/registration/login',
        permanent: true,
      },
    ]
  },
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
  images: {
    unoptimized: true,
    domains: ["i.scdn.co", "picsum.photos"],
  },
};

export default nextConfig;
