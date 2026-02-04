import { Event } from '@/lib/types';
import { Navbar } from '@/components/Navbar';
import { Card } from '@/components/ui';
import Link from 'next/link';

async function getEvents(): Promise<Event[]> {
  // ใช้ INTERNAL_API_URL สำหรับ SSR (ฝั่ง Server)
  const baseUrl = process.env.INTERNAL_API_URL || 'http://ticket-api:8080/api';

  try {
    console.log(`[SSR] Fetching from: ${baseUrl}/events`);
    const res = await fetch(`${baseUrl}/events`, { 
      cache: 'no-store',
      headers: { 'Content-Type': 'application/json' },
      next: { revalidate: 0 }
    });

    if (!res.ok) {
      console.error(`[SSR] Failed to fetch: ${res.status}`);
      return [];
    }

    return await res.json();
  } catch (error) {
    console.error('[SSR] Error:', error);
    return [];
  }
}

export default async function Home() {
  const events = await getEvents();

  return (
    <main className="min-h-screen bg-background transition-colors duration-300">
      <Navbar />
      
      <div className="max-w-5xl mx-auto px-4 pt-32 pb-20">
        <header className="mb-12 border-l-4 border-indigo-600 pl-6">
          <h1 className="text-3xl font-black text-slate-900 dark:text-white uppercase tracking-tight">
            Available <span className="text-indigo-600">Events</span>
          </h1>
          <p className="mt-2 text-slate-500 dark:text-slate-400 font-medium">
            รายการคอนเสิร์ตทั้งหมดที่ดึงข้อมูลจาก Backend API โดยตรง
          </p>
        </header>

        {events.length === 0 ? (
          <Card className="p-12 text-center border-dashed dark:bg-slate-900/50">
            <div className="text-4xl mb-4 text-slate-300">empty_inbox</div>
            <h3 className="text-lg font-bold text-slate-900 dark:text-white">ไม่พบรายการอีเวนต์</h3>
            <p className="text-slate-500 dark:text-slate-400">กรุณาตรวจสอบการเชื่อมต่อกับ Internal API ({process.env.INTERNAL_API_URL || 'http://ticket-api:8080/api'})</p>
          </Card>
        ) : (
          <div className="grid grid-cols-1 gap-6">
            {events.map((event) => (
              <Link key={event.id} href={`/events/${event.id}`}>
                <Card className="group p-6 hover:shadow-md transition-all dark:bg-slate-900 dark:border-slate-800 flex flex-col md:flex-row md:items-center justify-between gap-6 cursor-pointer">
                  <div className="flex-1">
                    <div className="flex items-center gap-3 mb-2">
                       <span className="text-xs font-mono px-2 py-0.5 bg-slate-100 dark:bg-slate-800 rounded text-slate-400">ID: {event.id}</span>
                       <span className="text-xs font-bold text-indigo-600 dark:text-indigo-400 uppercase tracking-widest">{event.location}</span>
                    </div>
                    <h2 className="text-2xl font-bold text-slate-900 dark:text-white group-hover:text-indigo-600 transition-colors">
                      {event.name}
                    </h2>
                    <div className="mt-2 flex flex-wrap gap-4 text-sm text-slate-500 dark:text-slate-400">
                      <div className="flex items-center gap-1 italic">
                        📅 {new Date(event.date_time).toLocaleDateString('th-TH', { dateStyle: 'long' })}
                      </div>
                      <div className="flex items-center gap-1 font-bold text-slate-700 dark:text-slate-300">
                        ⏰ {new Date(event.date_time).toLocaleTimeString('th-TH', { hour: '2-digit', minute: '2-digit' })} น.
                      </div>
                    </div>
                  </div>
                  
                  <div className="flex items-center gap-4">
                    <div className="text-right hidden md:block">
                      <p className="text-xs text-slate-400 uppercase font-bold">Status</p>
                      <p className="text-sm font-bold text-green-500">OPEN</p>
                    </div>
                    <div className="h-10 w-10 rounded-full bg-indigo-50 dark:bg-indigo-900/30 flex items-center justify-center text-indigo-600 dark:text-indigo-400 group-hover:bg-indigo-600 group-hover:text-white transition-all">
                      <svg className="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2.5} d="M9 5l7 7-7 7" />
                      </svg>
                    </div>
                  </div>
                </Card>
              </Link>
            ))}
          </div>
        )}
      </div>

      <footer className="mt-auto py-8 border-t border-slate-200 dark:border-slate-800 text-center bg-background">
        <p className="text-xs text-slate-400 font-mono">
          SSR_MODE: ENABLED | SOURCE: {process.env.INTERNAL_API_URL || 'DEFAULT_INTERNAL'}
        </p>
      </footer>
    </main>
  );
}
