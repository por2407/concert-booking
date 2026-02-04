'use client';
import Link from 'next/link';
import { Button } from './ui';
import { useEffect, useState } from 'react';

export const Navbar = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [scrolled, setScrolled] = useState(false);
  const [theme, setTheme] = useState<'light' | 'dark'>('light');

  useEffect(() => {
    setIsLoggedIn(!!localStorage.getItem('token'));
    
    // Initialize theme from document class
    if (document.documentElement.classList.contains('dark')) {
      setTheme('dark');
    }

    const handleScroll = () => {
      setScrolled(window.scrollY > 20);
    };
    window.addEventListener('scroll', handleScroll);
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  const toggleTheme = () => {
    const newTheme = theme === 'light' ? 'dark' : 'light';
    setTheme(newTheme);
    if (newTheme === 'dark') {
      document.documentElement.classList.add('dark');
      document.documentElement.style.colorScheme = 'dark';
      localStorage.setItem('theme', 'dark');
    } else {
      document.documentElement.classList.remove('dark');
      document.documentElement.style.colorScheme = 'light';
      localStorage.setItem('theme', 'light');
    }
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    window.location.reload();
  };

  return (
    <nav className={`fixed top-0 left-0 right-0 z-50 transition-all duration-500 ${
      scrolled 
        ? 'bg-white/80 dark:bg-slate-900/80 backdrop-blur-xl border-b border-slate-200 dark:border-slate-800 py-3 shadow-sm' 
        : 'bg-transparent py-6'
    }`}>
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center text-slate-900 dark:text-white">
          <Link href="/" className="group flex items-center gap-3">
            <div className="w-11 h-11 rounded-2xl bg-indigo-600 flex items-center justify-center text-white font-black text-2xl shadow-xl shadow-indigo-500/20 group-hover:scale-110 group-hover:rotate-3 transition-all duration-300">
              T
            </div>
            <div className="flex flex-col -gap-1">
              <span className={`text-2xl font-black tracking-tighter transition-colors ${scrolled ? 'text-slate-900 dark:text-white' : 'text-slate-900 dark:text-white md:text-white'}`}>
                Ticket<span className="text-indigo-500 italic">Ex</span>
              </span>
              <span className={`text-[10px] font-bold uppercase tracking-[0.2em] transition-opacity ${scrolled ? 'opacity-40' : 'opacity-40 md:opacity-60'}`}>
                Premium Experience
              </span>
            </div>
          </Link>
          
          <div className="flex items-center gap-1 md:gap-3">
            {/* Theme Toggle Button - Enhanced Visibility */}
            <button 
              onClick={toggleTheme}
              aria-label="Toggle Theme"
              className={`p-2.5 rounded-xl transition-all duration-300 border shadow-sm ${
                scrolled 
                  ? 'bg-white dark:bg-slate-800 border-slate-200 dark:border-slate-700 text-slate-600 dark:text-slate-300 hover:border-indigo-500' 
                  : 'bg-slate-100/80 dark:bg-white/10 border-slate-200/50 dark:border-white/20 text-slate-600 dark:text-white hover:bg-slate-200/80 dark:hover:bg-white/20'
              }`}
            >
              <div className="relative w-5 h-5 overflow-hidden">
                <div className={`absolute inset-0 transition-all duration-500 ${theme === 'light' ? 'translate-y-0 opacity-100 rotate-0' : '-translate-y-full opacity-0 rotate-90'}`}>
                  <svg className="w-5 h-5 fill-slate-500" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2}>
                    <path strokeLinecap="round" strokeLinejoin="round" d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z" />
                  </svg>
                </div>
                <div className={`absolute inset-0 transition-all duration-500 ${theme === 'dark' ? 'translate-y-0 opacity-100 rotate-0' : 'translate-y-full opacity-0 -rotate-90'}`}>
                  <svg className="w-5 h-5 fill-yellow-400 text-yellow-500" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={2}>
                    <path strokeLinecap="round" strokeLinejoin="round" d="M12 3v1m0 16v1m9-9h1M4 12H3m15.364-6.364l-.707.707M6.343 17.657l-.707.707m12.728 0l-.707-.707M6.343 6.343l-.707-.707M12 5a7 7 0 100 14 7 7 0 000-14z" />
                  </svg>
                </div>
              </div>
            </button>

            <div className={`h-6 w-[1px] mx-1 hidden sm:block transition-colors ${scrolled ? 'bg-slate-200 dark:bg-slate-700' : 'bg-slate-300 dark:bg-slate-700'}`} />

            {isLoggedIn ? (
              <div className="flex items-center gap-2">
                <Link href="/history">
                  <Button 
                    variant="secondary" 
                    className={`hidden sm:flex transition-all border-transparent shadow-none ${!scrolled ? 'bg-slate-100 dark:bg-white/10 text-slate-700 dark:text-white border-slate-200 dark:border-white/20 hover:bg-slate-200 dark:hover:bg-white/20' : ''}`}
                  >
                    My History
                  </Button>
                </Link>
                <Button 
                  variant="secondary" 
                  onClick={handleLogout}
                  className={`bg-rose-500 hover:bg-rose-600 text-white border-none shadow-lg shadow-rose-500/20 ${!scrolled ? '' : ''}`}
                >
                  Logout
                </Button>
              </div>
            ) : (
              <div className="flex gap-2">
                <Link href="/login" className="hidden sm:block">
                  <Button 
                    variant="secondary" 
                    className={`transition-all border-transparent shadow-none ${!scrolled ? 'text-slate-600 dark:text-slate-300 hover:bg-slate-100 dark:hover:bg-white/10' : 'text-slate-600 dark:text-slate-300'}`}
                  >
                    Login
                  </Button>
                </Link>
                <Link href="/register">
                  <Button className="shadow-lg shadow-indigo-600/20 px-5 sm:px-7 font-bold">
                    Start Now
                  </Button>
                </Link>
              </div>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
};
