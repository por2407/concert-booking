'use client';
import { useEffect, useState } from 'react';
import { api } from '@/lib/api';
import { Navbar } from '@/components/Navbar';
import { Card } from '@/components/ui';
import Link from 'next/link';

export default function BookingHistory() {
  const [bookings, setBookings] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchHistory = async () => {
      try {
        const res = await api.bookings.getHistory();
        // ใน Go Handler คุณส่งกลับมาแบบ {"booking items": bookings} 
        // หรือส่งมาเป็น array ตรงๆ? 
        // อ้างอิงจากโค้ด handler ลล่าสุด: return c.Status(fiber.StatusOK).JSON(fiber.Map{"booking items": bookingItems})
        setBookings(res["booking_items"] || []);
      } catch (err: any) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchHistory();
  }, []);

  if (loading) {
    return (
      <div className="min-h-screen bg-background">
        <Navbar />
        <div className="max-w-5xl mx-auto px-4 pt-32 text-center">
          <div className="inline-block w-8 h-8 border-4 border-indigo-600 border-t-transparent rounded-full animate-spin mb-4" />
          <p className="text-slate-500 font-bold uppercase tracking-widest text-xs">Loading your history...</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-background transition-colors duration-300">
      <Navbar />
      <div className="max-w-5xl mx-auto px-4 pt-32 pb-20">
        <div className="mb-12">
          <h1 className="text-4xl font-black text-slate-900 dark:text-white tracking-tight">Booking <span className="text-indigo-600">History</span></h1>
          <p className="text-slate-500 dark:text-slate-400 mt-2">ประวัติการจองและตั๋วทั้งหมดของคุณ</p>
        </div>

        {error && (
          <div className="p-4 rounded-xl bg-red-50 dark:bg-red-900/10 border border-red-100 dark:border-red-900/20 text-red-600 mb-8 font-bold">
            ⚠️ {error}
          </div>
        )}

        {bookings.length === 0 ? (
          <Card className="p-16 text-center dark:bg-slate-900/50 border-dashed border-2">
            <div className="w-16 h-16 bg-slate-100 dark:bg-slate-800 rounded-full flex items-center justify-center mx-auto mb-6 text-slate-400">
              <svg className="w-8 h-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 5v2m0 4v2m0 4v2M5 5a2 2 0 00-2 2v3a2 2 0 110 4v3a2 2 0 002 2h14a2 2 0 002-2v-3a2 2 0 110-4V7a2 2 0 00-2-2H5z" />
              </svg>
            </div>
            <h3 className="text-xl font-bold dark:text-white mb-2">ยังไม่มีประวัติการจอง</h3>
            <p className="text-slate-500 dark:text-slate-400 mb-8">คุณยังไม่ได้ทำการจองที่นั่งสำหรับงานคอนเสิร์ตใดๆ</p>
            <Link href="/" className="inline-flex items-center gap-2 px-6 py-3 bg-indigo-600 text-white font-bold rounded-xl hover:bg-indigo-700 transition-all shadow-lg shadow-indigo-600/20">
              ค้นหาคอนเสิร์ตที่น่าสนใจ
            </Link>
          </Card>
        ) : (
          <div className="grid grid-cols-1 gap-6">
            {bookings.map((booking) => (
              <Card key={booking.id} className="p-6 md:p-8 dark:bg-slate-900 dark:border-slate-800 shadow-xl hover:border-indigo-500/50 transition-all group overflow-hidden relative">
                <div className={`absolute top-0 right-0 px-8 py-1 font-black text-[10px] uppercase tracking-widest rotate-45 translate-x-10 translate-y-3 ${
                  booking.status === 'PAID' ? 'bg-green-500 text-white' : 'bg-amber-500 text-white'
                }`}>
                  {booking.status}
                </div>

                <div className="flex flex-col md:flex-row md:items-center justify-between gap-8">
                  <div className="flex-1">
                    <div className="flex items-center gap-3 mb-3 text-[10px] font-black uppercase tracking-[0.2em] text-indigo-500">
                      <span>Order #{booking.id}</span>
                      <span className="w-1 h-1 rounded-full bg-slate-300" />
                      <span>{new Date(booking.created_at).toLocaleDateString('th-TH')}</span>
                    </div>
                    <h2 className="text-2xl font-black text-slate-900 dark:text-white mb-4 group-hover:text-indigo-600 transition-colors">
                      {booking.event?.name || 'Unknown Event'}
                    </h2>
                    
                    <div className="flex flex-wrap gap-4">
                      {booking.items?.map((item: any) => (
                        <div key={item.id} className="flex flex-col p-3 bg-slate-50 dark:bg-slate-800 rounded-xl border border-slate-100 dark:border-slate-700 min-w-[100px]">
                           <span className="text-[10px] text-slate-400 font-bold uppercase tracking-widest">Seat</span>
                           <span className="text-lg font-black text-indigo-600 dark:text-indigo-400">
                             {item.seat?.row_label}{item.seat?.seat_number}
                           </span>
                        </div>
                      ))}
                    </div>
                  </div>

                  <div className="md:text-right flex flex-col justify-between h-full min-w-[150px]">
                    <div>
                      <p className="text-[10px] font-black text-slate-400 uppercase tracking-widest mb-1">Total Price</p>
                      <p className="text-3xl font-black text-slate-900 dark:text-white">฿{booking.total_amount?.toLocaleString()}</p>
                    </div>
                    
                    {booking.status === 'PENDING' && (
                       <button className="mt-4 w-full md:w-auto px-6 py-2 bg-indigo-600 text-white text-xs font-black uppercase tracking-widest rounded-lg shadow-lg shadow-indigo-600/20 hover:bg-white hover:text-indigo-600 border-2 border-transparent hover:border-indigo-600 transition-all">
                         Pay Now
                       </button>
                    )}
                  </div>
                </div>
              </Card>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
