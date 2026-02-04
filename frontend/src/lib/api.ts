const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8081/api';

async function request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
  const token = typeof window !== 'undefined' ? localStorage.getItem('token') : null;
  
  const headers = {
    'Content-Type': 'application/json',
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
    ...options.headers,
  };

  const response = await fetch(`${API_BASE_URL}${endpoint}`, {
    ...options,
    headers,
  });

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}));
    throw new Error(errorData.error || 'Something went wrong');
  }

  return response.json();
}

export const api = {
  auth: {
    login: (data: any) => request<{ token: string }>('/auth/login', { method: 'POST', body: JSON.stringify(data) }),
    register: (data: any) => request<any>('/auth/register', { method: 'POST', body: JSON.stringify(data) }),
  },
  events: {
    getAll: () => request<any[]>('/events'),
    getById: (id: string | number) => request<any>(`/events/${id}`),
    create: (data: any) => request<any>('/events/create', { method: 'POST', body: JSON.stringify(data) }),
  },
  bookings: {
    create: (data: { event_id: number; seat_ids: number[] }) => 
      request<any>('/bookings', { method: 'POST', body: JSON.stringify(data) }),
    confirm: (bookingId: number) => 
      request<any>('/bookings/confirm', { method: 'POST', body: JSON.stringify({ booking_id: bookingId }) }),
    getHistory: () => request<any>('/bookings/history'),
  },
};
