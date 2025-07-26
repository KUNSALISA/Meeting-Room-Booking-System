import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute } from '@angular/router';
import { FormControl,FormBuilder, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { first } from 'rxjs';

@Component({
  selector: 'app-signup',
  imports: [CommonModule, ReactiveFormsModule],
  templateUrl: './signup.component.html',
  styleUrl: './signup.component.css',
})
export class SignupComponent {
  signupForm!: FormGroup;

  constructor(private fb: FormBuilder) {}

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

  submitSignup() {
    if (this.signupForm.valid) {
      const formData = this.signupForm.value;
      console.log('Signup data:', formData);
      formData.roleID = 2;

    } else {
      console.log('Signup form invalid');
    }
  }
}
