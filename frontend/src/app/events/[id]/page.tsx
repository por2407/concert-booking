import { Event } from '@/lib/types';
import { Navbar } from '@/components/Navbar';
import SeatSelection from '@/components/SeatSelection';

async function getEvent(id: string): Promise<Event | null> {
  const baseUrl = process.env.INTERNAL_API_URL || 'http://ticket-api:8080/api';
  try {
    const res = await fetch(`${baseUrl}/events/${id}`, { 
      cache: 'no-store',
      headers: { 'Content-Type': 'application/json' },
      next: { revalidate: 0 }
    });
    if (!res.ok) return null;
    return await res.json();
  } catch (err) {
    console.error('[SSR] GetEvent failed:', err);
    return null;
  }
}

export default async function EventDetail({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params;
  const event = await getEvent(id);

  if (!event) {
    return (
      <main className="min-h-screen bg-background">
        <Navbar />
        <div className="max-w-5xl mx-auto px-4 pt-40 pb-20 text-center">
           <h1 className="text-2xl font-bold dark:text-white">Event Not Found</h1>
        </div>
      </main>
    );
  }

  return (
    <main className="min-h-screen bg-background transition-colors duration-300">
      <Navbar />
      
      {/* Hero Section */}
      <div className="relative pt-32 pb-20 overflow-hidden">
        <div className="absolute inset-0 bg-indigo-600/5 dark:bg-indigo-600/10 -skew-y-3 origin-top-left" />
        <div className="max-w-5xl mx-auto px-4 relative">
          <div className="flex flex-col md:flex-row md:items-end justify-between gap-8">
            <div>
              <div className="flex items-center gap-3 mb-4">
                <span className="px-3 py-1 bg-indigo-600 text-white text-[10px] font-black uppercase tracking-widest rounded-full">Official Event</span>
                <span className="text-xs font-bold text-slate-400">ID: #{event.id}</span>
              </div>
              <h1 className="text-4xl md:text-6xl font-black text-slate-900 dark:text-white tracking-tight leading-none mb-4">
                {event.name}
              </h1>
              <div className="flex flex-wrap gap-6 text-slate-500 dark:text-slate-400 font-medium">
                <div className="flex items-center gap-2">
                  <span className="text-indigo-600 dark:text-indigo-400">📍</span> {event.location}
                </div>
                <div className="flex items-center gap-2">
                  <span className="text-indigo-600 dark:text-indigo-400">📅</span> {new Date(event.date_time).toLocaleDateString('th-TH', { dateStyle: 'full' })}
                </div>
              </div>
            </div>
            
            <div className="flex items-center gap-4 p-4 bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 shadow-xl">
               <div className="text-right">
                 <p className="text-[10px] font-black uppercase text-slate-400 tracking-widest">Time Remaining</p>
                 <p className="text-xl font-black text-indigo-600 dark:text-indigo-400">Booking Open</p>
               </div>
               <div className="w-12 h-12 rounded-xl bg-indigo-50 dark:bg-indigo-900/30 flex items-center justify-center text-indigo-600 animate-pulse">
                  <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2.5} d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
               </div>
            </div>
          </div>
        </div>
      </div>

      <div className="max-w-5xl mx-auto px-4 pb-32">
        <header className="mb-10 flex items-center justify-between">
           <h2 className="text-2xl font-black text-slate-900 dark:text-white uppercase">Choose your <span className="text-indigo-600">Seats</span></h2>
           <div className="text-sm font-bold text-slate-400 italic">Seat Map Version 2.4.0</div>
        </header>

        <SeatSelection event={event} />
        
        <div className="mt-20 border-t border-slate-200 dark:border-slate-800 pt-12 grid grid-cols-1 md:grid-cols-3 gap-12 text-center md:text-left">
           <div>
             <h4 className="font-bold dark:text-white mb-2 uppercase tracking-widest text-sm">Security</h4>
             <p className="text-sm text-slate-500">ระบบการจองของเราได้รับมาตรฐานสากล ข้อมูลของคุณจะถูกเข้ารหัสระดับสูงสุด</p>
           </div>
           <div>
             <h4 className="font-bold dark:text-white mb-2 uppercase tracking-widest text-sm">Refund Policy</h4>
             <p className="text-sm text-slate-500">สามารถแจ้งยกเลิกและขอคืนเงินได้ล่วงหน้า 7 วันก่อนวันงานแสดงเท่านั้น</p>
           </div>
           <div>
             <h4 className="font-bold dark:text-white mb-2 uppercase tracking-widest text-sm">Support</h4>
             <p className="text-sm text-slate-500">ต้องการความช่วยเหลือ? ติดต่อได้ที่ support@ticketex.com ตลอด 24 ชม.</p>
           </div>
        </div>
      </div>
    </main>
  );
}
