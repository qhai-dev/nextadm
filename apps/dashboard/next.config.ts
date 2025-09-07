import type { NextConfig } from 'next';
import createNextIntlPlugin from 'next-intl/plugin';

const withNextIntl = createNextIntlPlugin({
	requestConfig: './i18n/server.ts',
});

const nextConfig: NextConfig = {
	async redirects() {
		return [
			{
				source: '/',
				destination: '/dashboard',
				permanent: false,
			},
		];
	},
	output: 'standalone',
};

export default withNextIntl(nextConfig);
