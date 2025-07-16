import { HttpClient } from '@angular/common/http';
import { Component, inject } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-login',
  imports: [ReactiveFormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css'
})
export class LoginComponent {
  loginForm: FormGroup = new FormGroup({
    Codename: new FormControl(""),
    Password: new FormControl("")
  })

  http = inject(HttpClient);

  onLogin() {
    debugger;
    const formValue = this.loginForm.value;
    this.http.post("",formValue).subscribe({
      next:(result) => {

      },
      error:(error) => {

      }
    })
  }
}
