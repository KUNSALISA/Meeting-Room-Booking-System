import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { LoginIn, LoginResponse, PasswordIn, SignupIn } from './login';

@Injectable({
  providedIn: 'root',
})
export class LoginService {
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

  signinUser(loginData: LoginIn): Observable<LoginResponse> {
    const url = `${this.baseUrl}/signin`;
    return this.http
      .post<LoginResponse>(url, loginData, {
        headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      })
      .pipe(catchError(this.handleError<LoginResponse>('signinUser')));
  }

  signUpUser(signupData: SignupIn): Observable<any> {
    const url = `${this.baseUrl}/signup`;
    return this.http
      .post<any>(url, signupData, {
        headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      })
      .pipe(catchError(this.handleError<any>('signUpUser')));
  }

  changePassword(passwordData: PasswordIn): Observable<PasswordIn> {
    const url = `${this.baseUrl}/change-password`;
    return this.http
      .patch<PasswordIn>(url, passwordData, {
        headers: this.getAuthHeaders(),
      })
      .pipe(catchError(this.handleError<PasswordIn>('changePassword')));
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
