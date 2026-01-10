import { NextRequest, NextResponse } from 'next/server';
import { Ratelimit } from '@upstash/ratelimit';
import { kv } from '@vercel/kv';

const ratelimit = new Ratelimit({
  redis: kv,
  // 5 requests from the same IP in 10 seconds
  limiter: Ratelimit.slidingWindow(2, '10 s'),
});

// Define which routes you want to rate limit
export const config = {
  matcher: '/',
};

export default async function middleware(request: NextRequest) {
  // Get IP address from headers (Next.js 16 compatible)
  const forwarded = request.headers.get('x-forwarded-for');
  const ip = forwarded ? forwarded.split(',')[0] : request.headers.get('x-real-ip') || '127.0.0.1';
  
  if(request.method === 'GET') {
    // You could alternatively limit based on user ID or similar
    await kv.incr(`ip_${ip}`);
    const request_num = await kv.get(`ip_${ip}`);
    console.log(`ip: ${ip} request_num: ${request_num}`);
    return NextResponse.next()
  } else {
    // You could alternatively limit based on user ID or similar
    await kv.incr(`ip_${ip}`);
    const request_num = await kv.get(`ip_${ip}`);
    const { success, pending, limit, reset, remaining } = await ratelimit.limit(
     ip
    );
    console.log(`ip: ${ip} request_num: ${request_num} success: ${success}`);
    return success
    ? NextResponse.next()
    : NextResponse.redirect(new URL('/blocked', request.url));
  }
}
