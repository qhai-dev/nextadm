import { NextRequest, NextResponse } from 'next/server';

// TODO 权限中间件
function withAuth() {}

export function middleware(request: NextRequest) {
	// let cookie = request.cookies.get('local')?.value || 'zH-Hant';
	// console.log(cookie, ''); // => { name: 'nextjs', value: 'fast', Path: '/' }
	// Setting cookies on the response using the `ResponseCookies` API
	// const response = NextResponse.next();
	// response.cookies.set('local', cookie);
	// return response;
}
