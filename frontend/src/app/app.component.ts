import { Component } from '@angular/core';
import { RouterModule, Router } from '@angular/router';
import { LoginResponse } from './login';
import { CommonModule } from '@angular/common';

@Component({
  standalone: true,
  selector: 'app-root',
  imports: [RouterModule, CommonModule],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css',
})
export class AppComponent {
  user: LoginResponse | null = null;

  constructor(private router: Router) {}

  ngOnInit(): void {
    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      this.user = JSON.parse(storedUser) as LoginResponse;
    }
  }

  logout(): void {
    localStorage.removeItem('user');
    this.user = null;
    this.router.navigate(['/login']);
  }
}
