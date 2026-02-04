'use client';
import { useState } from 'react';
import { api } from '@/lib/api';
import { Button, Card } from '@/components/ui';
import { Navbar } from '@/components/Navbar';

export default function AdminPage() {
  const [name, setName] = useState('');
  const [location, setLocation] = useState('');
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState('');

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setMessage('');
    try {
      await api.events.create({ name, location });
      setMessage('✅ Event and 50 seats created successfully!');
      setName('');
      setLocation('');
    } catch (err: any) {
      setMessage('❌ Error: ' + err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-background transition-colors duration-300">
      <Navbar />
      <div className="max-w-4xl mx-auto px-4 pt-32 pb-20">
        <div className="mb-10">
          <h1 className="text-4xl font-black text-slate-900 dark:text-white tracking-tight">Admin <span className="text-indigo-600">Dashboard</span></h1>
          <p className="text-slate-500 dark:text-slate-400 mt-2">จัดการข้อมูลรายการแสดงและที่นั่งแบบระบบหลังบ้าน</p>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <div className="lg:col-span-2">
            <Card className="p-8 dark:bg-slate-900 dark:border-slate-800 shadow-xl relative overflow-hidden">
              <div className="absolute top-0 left-0 w-1 h-full bg-indigo-600" />
              <h2 className="text-2xl font-bold dark:text-white mb-8 flex items-center gap-2">
                <span className="p-2 rounded-lg bg-indigo-50 dark:bg-indigo-900/30 text-indigo-600 dark:text-indigo-400">
                  <svg className="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
                  </svg>
                </span>
                Create New Event
              </h2>
              
              <form onSubmit={handleCreate} className="space-y-6">
                <div className="grid grid-cols-1 gap-6">
                  <div>
                    <label className="block text-sm font-bold text-slate-700 dark:text-slate-300 mb-2">Event Display Name</label>
                    <input 
                      type="text" 
                      className="w-full p-3 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl outline-none focus:ring-2 focus:ring-indigo-500 transition-all dark:text-white"
                      value={name}
                      placeholder="e.g. Taylor Swift | The Eras Tour"
                      onChange={e => setName(e.target.value)}
                      required
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-bold text-slate-700 dark:text-slate-300 mb-2">Venue Location</label>
                    <input 
                      type="text" 
                      className="w-full p-3 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl outline-none focus:ring-2 focus:ring-indigo-500 transition-all dark:text-white"
                      value={location}
                      placeholder="e.g. IMPACT Arena, Muang Thong Thani"
                      onChange={e => setLocation(e.target.value)}
                      required
                    />
                  </div>
                </div>

                {message && (
                  <div className={`p-4 rounded-xl text-sm font-bold ${message.startsWith('✅') ? 'bg-green-50 text-green-700 dark:bg-green-900/20 dark:text-green-400 border border-green-100 dark:border-green-900/30' : 'bg-red-50 text-red-700 dark:bg-red-900/20 dark:text-red-400 border border-red-100 dark:border-red-900/30'}`}>
                    {message}
                  </div>
                )}

                <Button type="submit" className="w-full py-4 text-lg font-bold shadow-lg shadow-indigo-600/20" isLoading={loading}>
                  Generate Event & 500 Seats
                </Button>
              </form>
            </Card>
          </div>

          <div className="space-y-6">
            <Card className="p-6 dark:bg-slate-900 dark:border-slate-800">
              <h3 className="font-bold dark:text-white mb-4">Quick Stats</h3>
              <div className="space-y-4">
                <div className="flex justify-between items-center text-sm">
                  <span className="text-slate-500">Total Events</span>
                  <span className="font-mono font-bold dark:text-indigo-400">12</span>
                </div>
                <div className="flex justify-between items-center text-sm">
                  <span className="text-slate-500">Total Revenue</span>
                  <span className="font-mono font-bold dark:text-indigo-400">฿1.2M</span>
                </div>
              </div>
            </Card>
            
            <Card className="p-6 bg-indigo-600 text-white border-none shadow-xl shadow-indigo-500/20">
              <h3 className="font-bold mb-2 italic">Admin Tip</h3>
              <p className="text-sm text-indigo-100">การสร้าง Event ใหม่จะรันระบบอัตโนมัติเพื่อเจนที่นั่งมาตรฐานให้ทันที จำนวน 5 แถวๆ ละ 100 ที่นั่ง</p>
            </Card>
          </div>
        </div>
      </div>
    </div>
  );
}
