import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginService } from '../login.service';
import {
  FormControl,
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
} from '@angular/forms';
import { Router } from '@angular/router';
import { MatIconModule } from '@angular/material/icon';
import { first } from 'rxjs';

@Component({
  selector: 'app-login',
  imports: [CommonModule, ReactiveFormsModule,MatIconModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css',
})
export class LoginComponent {
  loginForm!: FormGroup;
  resetForm!: FormGroup;
  showResetBox = false;

  openResetBox() {
    this.showResetBox = true;
  }

  closeResetBox() {
    this.showResetBox = false;
  }

  constructor(
    private fb: FormBuilder,
    private authService: LoginService,
    private router: Router
  ) {}

  ngOnInit(): void {
    this.loginForm = this.fb.group({
      codename: [''],
      password: [''],
    });

    this.resetForm = this.fb.group({
      codename: [''],
      newPassword: [''],
    });
  }

  submitLogin() {
    if (this.loginForm.valid) {
      const loginData = this.loginForm.value;

      this.authService
        .signinUser(loginData)
        .pipe(first())
        .subscribe({
          next: (res) => {
            console.log('เข้าสู่ระบบสำเร็จ:', res);
            const userData = {
              token_type: res.token_type,
              token: res.token,
              user_id: res.user_id,
              codename: res.codename,
              first_name: res.first_name,
              last_name: res.last_name,
              image: res.image,
              email: res.email,
              phone_number: res.phone_number,
              role: res.role,
            };
            localStorage.setItem('user', JSON.stringify(userData));
            this.router.navigate(['']);
          },
          error: (err) => {
            console.error('Login error:', err.message);
          },
        });
    } else {
      console.log('ไม่สามารถเข้าใช้ได้');
    }
  }

  submitReset() {
    if (this.resetForm.valid) {
      const passwordData = this.resetForm.value;

      this.authService
        .changePassword(passwordData)
        .pipe(first())
        .subscribe({
          next: (res) => {
            console.log('เปลี่ยนรหัสผ่านสำเร็จ:', res);
            this.resetForm.reset();
            this.closeResetBox();
          },
          error: (err) => {
            console.error('เกิดข้อผิดพลาดในการเปลี่ยนรหัสผ่าน:', err.message);
          },
        });
    } else {
      console.log('รีเซ็ตแบบฟอร์มไม่ถูกต้อง');
    }
  }
}
