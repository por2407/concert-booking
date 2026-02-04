'use client';
import { useState } from 'react';
import { api } from '@/lib/api';
import { useRouter } from 'next/navigation';
import { Button, Card } from '@/components/ui';
import Link from 'next/link';
import { Navbar } from '@/components/Navbar';

export default function Register() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);
  const router = useRouter();

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError('');
    try {
      await api.auth.register({ email, password });
      alert('Registration successful! Please login.');
      router.push('/login');
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex flex-col bg-background transition-colors duration-300">
      <Navbar />
      <div className="flex-1 flex items-center justify-center p-4 relative">
        <div className="absolute inset-0 bg-gradient-to-tr from-indigo-500/10 via-transparent to-pink-500/10 pointer-events-none" />
      
      <Card className="w-full max-w-md p-10 dark:bg-slate-900 dark:border-slate-800 shadow-2xl relative overflow-hidden">
        <div className="absolute top-0 left-0 w-full h-1 bg-indigo-600" />
        
        <div className="mb-10 text-center">
          <Link href="/" className="inline-flex items-center gap-2 mb-6">
            <div className="w-10 h-10 rounded-xl bg-indigo-600 flex items-center justify-center text-white font-black text-xl shadow-lg shadow-indigo-500/30">
              T
            </div>
            <span className="text-2xl font-black tracking-tight dark:text-white">
              Ticket<span className="text-indigo-500 italic">Ex</span>
            </span>
          </Link>
          <h1 className="text-3xl font-bold dark:text-white">Join Experience</h1>
          <p className="text-slate-500 dark:text-slate-400 mt-2">เริ่มต้นการเดินทางสายดนตรีไปกับเรา</p>
        </div>

        <form onSubmit={handleRegister} className="space-y-5">
          <div>
            <label className="block text-sm font-bold text-slate-700 dark:text-slate-300 mb-2 px-1">Email Address</label>
            <input 
              type="email" 
              placeholder="name@example.com"
              className="w-full p-3 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl focus:ring-2 focus:ring-indigo-500 dark:focus:ring-indigo-400 outline-none transition-all dark:text-white"
              value={email}
              onChange={e => setEmail(e.target.value)}
              required
            />
          </div>
          <div>
            <label className="block text-sm font-bold text-slate-700 dark:text-slate-300 mb-2 px-1">Set Password</label>
            <input 
              type="password" 
              placeholder="Minimum 8 characters"
              className="w-full p-3 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl focus:ring-2 focus:ring-indigo-500 dark:focus:ring-indigo-400 outline-none transition-all dark:text-white"
              value={password}
              onChange={e => setPassword(e.target.value)}
              required
            />
          </div>
          
          {error && (
            <div className="p-3 rounded-lg bg-red-50 dark:bg-red-900/20 border border-red-100 dark:border-red-900/30 text-red-600 dark:text-red-400 text-sm font-medium">
              ⚠️ {error}
            </div>
          )}

          <Button type="submit" className="w-full py-4 text-base font-bold shadow-lg shadow-indigo-600/20 mt-4" isLoading={loading}>
            Create My Account
          </Button>
        </form>

        <div className="mt-8 pt-8 border-t border-slate-100 dark:border-slate-800 text-center">
          <p className="text-sm text-slate-500 dark:text-slate-400">
            มีบัญชีแล้วใช่ไหม? <Link href="/login" className="text-indigo-600 dark:text-indigo-400 font-bold hover:underline">เข้าสู่ระบบ</Link>
          </p>
        </div>
      </Card>
      </div>
    </div>
  );
}
