'use server';

import { cookies } from 'next/headers';

import { defaultLocale, Locale } from './conf';

export async function getUserLocale() {
	const cookie = await cookies();
	return cookie.get('NEXT_LOCALE')?.value || defaultLocale;
}

export async function setUserLocale(locale: Locale) {
	const cookie = await cookies();
	cookie.set('NEXT_LOCALE', locale);
}
