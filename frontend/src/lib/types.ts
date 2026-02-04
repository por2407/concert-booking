export interface User {
  id: number;
  email: string;
  role: string;
}

export interface Seat {
  id: number;
  event_id: number;
  row_label: string;
  seat_number: number;
  price: number;
  status: 'AVAILABLE' | 'LOCKED' | 'SOLD';
}

export interface Event {
  id: number;
  name: string;
  location: string;
  date_time: string;
  seats?: Seat[];
}

export interface Booking {
  id: number;
  user_id: number;
  event_id: number;
  total_amount: number;
  status: 'PENDING' | 'PAID' | 'CANCELLED';
  items: BookingItem[];
}

export interface BookingItem {
  id: number;
  booking_id: number;
  seat_id: number;
  seat?: Seat;
}
