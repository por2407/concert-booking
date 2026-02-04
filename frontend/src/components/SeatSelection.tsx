'use client';
import { useState } from 'react';
import { api } from '@/lib/api';
import { Event } from '@/lib/types';
import { Button, Card } from '@/components/ui';
import { useRouter } from 'next/navigation';

export default function SeatSelection({ event }: { event: Event }) {
  const router = useRouter();
  const [selectedSeats, setSelectedSeats] = useState<number[]>([]);
  const [bookingLoading, setBookingLoading] = useState(false);

  const toggleSeat = (seatId: number) => {
    setSelectedSeats(prev => 
      prev.includes(seatId) ? prev.filter(s => s !== seatId) : [...prev, seatId]
    );
  };

  const handleBooking = async () => {
    if (!localStorage.getItem('token')) {
      router.push('/login');
      return;
    }

    try {
      setBookingLoading(true);
      const booking = await api.bookings.create({
        event_id: Number(event.id),
        seat_ids: selectedSeats
      });
      alert('Reserved! Please confirm payment.');
      await api.bookings.confirm(booking.id);
      alert('Booking Confirmed!');
      window.location.reload();
    } catch (err: any) {
      alert(err.message);
    } finally {
      setBookingLoading(false);
    }
  };

  const totalPrice = selectedSeats.length * 1200; // Standard price for now

  return (
    <div className="space-y-12">
      <Card className="p-8 md:p-12 dark:bg-slate-900 dark:border-slate-800 shadow-2xl relative overflow-hidden">
        <div className="absolute top-0 left-0 w-full h-1 bg-indigo-600/20" />
        
        {/* Stage Visualization */}
        <div className="mb-20">
          <div className="w-full h-2 bg-gradient-to-r from-transparent via-indigo-500 to-transparent rounded-full mb-4 opacity-50" />
          <p className="text-center text-xs font-black uppercase tracking-[0.5em] text-indigo-400 opacity-60">STAGE / PERFORMANCE AREA</p>
        </div>
        
        {/* Seat Grid */}
        <div className="flex justify-center overflow-x-auto pb-6">
          <div className="inline-grid grid-cols-5 sm:grid-cols-10 gap-3 min-w-max px-4">
            {event.seats?.map((seat) => (
              <button
                key={seat.id}
                disabled={seat.status !== 'AVAILABLE'}
                onClick={() => toggleSeat(seat.id)}
                className={`
                  w-10 h-10 rounded-xl flex items-center justify-center text-xs font-bold transition-all duration-300 relative group
                  ${seat.status === 'AVAILABLE' ? (
                    selectedSeats.includes(seat.id) 
                      ? 'bg-indigo-600 text-white scale-110 shadow-lg shadow-indigo-500/40 ring-2 ring-indigo-400 ring-offset-2 dark:ring-offset-slate-900' 
                      : 'bg-indigo-50 dark:bg-slate-800 text-indigo-600 dark:text-indigo-400 hover:bg-indigo-600 hover:text-white dark:hover:bg-indigo-600'
                  ) : 'bg-slate-100 dark:bg-slate-800/50 text-slate-300 dark:text-slate-700 cursor-not-allowed'}
                `}
                title={`Seat ${seat.seat_number} - ${seat.status}`}
              >
                {seat.seat_number}
                {seat.status === 'SOLD' && <div className="absolute w-full h-[2px] bg-slate-300 dark:bg-slate-700 rotate-45" />}
              </button>
            ))}
          </div>
        </div>

        {/* Legend */}
        <div className="mt-12 flex flex-wrap justify-center gap-6 text-xs font-bold uppercase tracking-widest text-slate-400">
           <div className="flex items-center gap-2">
             <div className="w-4 h-4 rounded-md bg-indigo-50 dark:bg-slate-800 border border-indigo-100 dark:border-slate-700" />
             <span>Available</span>
           </div>
           <div className="flex items-center gap-2">
             <div className="w-4 h-4 rounded-md bg-indigo-600 shadow-lg shadow-indigo-500/20" />
             <span className="text-indigo-600 dark:text-indigo-400">Selected</span>
           </div>
           <div className="flex items-center gap-2">
             <div className="w-4 h-4 rounded-md bg-slate-100 dark:bg-slate-800/50 flex items-center justify-center relative overflow-hidden">
                <div className="absolute w-full h-[1px] bg-slate-300 dark:bg-slate-700 rotate-45" />
             </div>
             <span>Sold Out</span>
           </div>
        </div>
      </Card>

      {/* Floating Action Bar */}
      {selectedSeats.length > 0 && (
        <div className="fixed bottom-8 left-0 right-0 z-50 px-4 animate-in fade-in slide-in-from-bottom-4 duration-500">
          <Card className="max-w-3xl mx-auto p-5 bg-white/80 dark:bg-slate-900/80 backdrop-blur-2xl border-indigo-500/20 shadow-2xl flex items-center justify-between gap-6">
            <div className="flex items-center gap-4">
              <div className="w-12 h-12 rounded-2xl bg-indigo-600 flex items-center justify-center text-white shrink-0 shadow-lg shadow-indigo-500/20">
                <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 5v2m0 4v2m0 4v2M5 5a2 2 0 00-2 2v3a2 2 0 110 4v3a2 2 0 002 2h14a2 2 0 002-2v-3a2 2 0 110-4V7a2 2 0 00-2-2H5z" />
                </svg>
              </div>
              <div>
                <p className="text-[10px] font-black uppercase tracking-widest text-slate-400">Seats Reserved: {selectedSeats.length}</p>
                <p className="text-2xl font-black text-slate-900 dark:text-white">฿{totalPrice.toLocaleString()}</p>
              </div>
            </div>
            <Button 
              onClick={handleBooking} 
              isLoading={bookingLoading}
              className="px-10 py-4 h-auto text-lg font-bold rounded-2xl shadow-xl shadow-indigo-600/30"
            >
              Confirm Booking
            </Button>
          </Card>
        </div>
      )}
    </div>
  );
}
