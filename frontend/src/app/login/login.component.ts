import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute } from '@angular/router';
import { FormControl,FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { first } from 'rxjs';

@Component({
  selector: 'app-login',
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {

  loginForm!: FormGroup;
  resetForm!: FormGroup;
  showResetModal = false;

  constructor(private fb: FormBuilder) {}

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
      console.log('Codename:', this.loginForm.value.codename);
      console.log('Password:', this.loginForm.value.password);
    } else {
      console.log('Login form invalid');
    }
  }

  openResetModal() {
    this.showResetModal = true;
  }

  closeResetModal() {
    this.showResetModal = false;
  }

  submitReset() {
    if (this.resetForm.valid) {
      const { codename, newPassword } = this.resetForm.value;
      console.log('Reset Password for:', codename);
      console.log('New Password:', newPassword);
      // TODO: Send to backend service
      this.closeResetModal();
    } else {
      console.log('Reset form invalid');
    }
  }

}
