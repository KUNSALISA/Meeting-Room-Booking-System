import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { finalize } from 'rxjs/operators';
import {
  FormControl,
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
} from '@angular/forms';
import { LoginService } from '../login.service';
import { SignupIn } from '../login';
import { first } from 'rxjs';

@Component({
  selector: 'app-signup',
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './signup.component.html',
  styleUrl: './signup.component.css',
})
export class SignupComponent {
  signupForm!: FormGroup;
  uploadingImage: boolean = false;
  imageData: string = '';

  constructor(private fb: FormBuilder, private signupService: LoginService) {}

  ngOnInit() {
    this.signupForm = this.fb.group({
      codename: [''],
      password: [''],
      firstname: [''],
      lastname: [''],
      image: [''],
      email: [''],
      phoneNumber: [''],
      roleID: [2],
    });
  }

  onFileChange(event: any) {
    const file = event.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = () => {
        this.imageData = reader.result as string; // base64 string
        this.signupForm.patchValue({ image: this.imageData });
      };
      reader.readAsDataURL(file);
    }
  }

  submitSignup() {
    if (this.signupForm.valid) {
      const formData = this.signupForm.value;

      const signupData: SignupIn = {
        CodeName: formData.codename,
        Password: formData.password,
        Firstname: formData.firstname,
        Lastname: formData.lastname,
        Image: this.imageData || '',
        Email: formData.email,
        PhoneNumber: formData.phoneNumber,
        RoleID: formData.roleID,
      };

      console.log('Signup data:', signupData);

      this.signupService
        .signUpUser(signupData)
        .pipe(finalize(() => console.log('Signup request completed')))
        .subscribe({
          next: (res) => {
            console.log('Signup success', res);
            // ทำอะไรต่อ เช่น แสดงข้อความ แจ้งเตือน หรือ redirect
          },
          error: (err) => {
            console.error('Signup failed', err);
          },
        });
    } else {
      console.log('Signup form invalid');
    }
  }
}
