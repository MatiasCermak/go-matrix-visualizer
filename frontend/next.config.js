/** @type {import('next').NextConfig} */
const nextConfig = {
    reactStrictMode: true,
    publicRuntimeConfig: {
        API_URL: "localhost:8080",
    },
};

module.exports = nextConfig;
