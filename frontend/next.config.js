/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    appDir: true,
  },
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        destination: 'http://gin_hackbook:8080/:path*',
      },
    ]
  },
}

module.exports = nextConfig
