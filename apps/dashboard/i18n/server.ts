import { getRequestConfig } from 'next-intl/server';

import { getUserLocale } from './index';

export default getRequestConfig(async () => {
	const locale = await getUserLocale();
	const message = await import(`./${locale}/index.json`);

	return {
		locale,
		messages: message.default,
	};
});
