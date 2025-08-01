import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { RoomIn, BookingIn, RoomStatus } from './meeting-room';

@Injectable({
  providedIn: 'root',
})
export class MeetingRoomService {
  private baseUrl = 'http://localhost:8000';

  constructor(private http: HttpClient) {}

  private getAuthHeaders(): HttpHeaders {
    const token = localStorage.getItem('token');
    const tokenType = localStorage.getItem('token_type');
    return new HttpHeaders({
      'Content-Type': 'application/json',
      Authorization: `${tokenType} ${token}`,
    });
  }

  getAllRooms(): Observable<RoomIn[]> {
    const url = `${this.baseUrl}/all-room`;
    const headers = this.getAuthHeaders();
    return this.http.get<RoomIn[]>(url, { headers }).pipe(
      catchError((err) => {
        console.error('Error fetching rooms', err);
        return throwError(() => err);
      })
    );
  }

  getRoomById(id: string): Observable<RoomIn> {
    const url = `${this.baseUrl}/all-room/${id}`;
    const headers = this.getAuthHeaders();
    return this.http.get<RoomIn>(url, { headers }).pipe(
      catchError((err) => {
        console.error('Error fetching room', err);
        return throwError(() => err);
      })
    );
  }

  getAllStatus(): Observable<RoomStatus[]> {
    const url = `${this.baseUrl}/status`;
    const headers = this.getAuthHeaders();
    return this.http.get<RoomStatus[]>(url, { headers }).pipe(
      catchError((err) => {
        console.error('Error fetching rooms', err);
        return throwError(() => err);
      })
    );
  }

  PostBooking(bookingData: BookingIn): Observable<BookingIn> {
    const url = `${this.baseUrl}/booking`;
    return this.http
      .post<BookingIn>(url, bookingData, {
        headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      })
      .pipe(catchError(this.handleError<BookingIn>('PostBooking')));
  }

  private handleError<T>(operation = 'operation') {
    return (error: any): Observable<T> => {
      console.error(`${operation} failed: ${error.message}`);
      return throwError(
        () => new Error(`${operation} failed: ${error.message}`)
      );
    };
  }
}
