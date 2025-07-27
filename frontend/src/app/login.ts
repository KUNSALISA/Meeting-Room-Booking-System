export interface LoginIn {
  Codename: string;
  Password: string;
}

export interface LoginResponse {
  token_type: string;
  token: string;
  user_id: string;
  codename: string;
  first_name: string;
  last_name: string;
  image: string;
  email: string;
  phone_number: string;
  role: string;
}

export interface PasswordIn {
  CodeName: string;
  NewPassword: string;
}

export interface SignupIn {
  ID?: number;
  CodeName: string;
  Password: string;
  Firstname: string;
  Lastname: string;
  Image: string;
  Email: string;
  PhoneNumber: string;
  RoleID: number;
}

interface RoleIn {
  ID: number;
  RoleName: string;
}