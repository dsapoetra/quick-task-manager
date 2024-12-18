import { jwtDecode } from "jwt-decode";  // Change this line
import { DecodedToken } from './types/auth.types';
import { UserData } from './types/auth.types';

export const getToken = (): string | null => {
  return localStorage.getItem('jwt_token');
};

export const setToken = (token: string): void => {
  localStorage.setItem('jwt_token', token);
};

export const removeToken = (): void => {
  localStorage.removeItem('jwt_token');
};

export const isTokenValid = (): boolean => {
  const token = getToken();
  if (!token) return false;

  try {
    const decoded: DecodedToken = jwtDecode(token);
    const currentTime = Date.now() / 1000;
    return decoded.exp > currentTime;
  } catch {
    return false;
  }
};

export const getUserFromToken = (): UserData | null => {
  const token = getToken();
  if (!token) return null;

  try {
    const decoded: DecodedToken = jwtDecode(token);
    return decoded.user;
  } catch {
    return null;
  }
};