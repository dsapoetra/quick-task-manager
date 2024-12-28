import { jwtDecode } from 'jwt-decode';
import { UserData } from '../types/auth.types';

interface JWTPayload {
  exp: number;
  user_id: number;
}

export const getToken = (): string | null => {
  if (typeof window !== 'undefined') {
    return localStorage.getItem('jwt_token');
  }
  return null;
};

export const setToken = (token: string): void => {
  if (typeof window !== 'undefined') {
    localStorage.setItem('jwt_token', token);
  }
};

export const removeToken = (): void => {
  if (typeof window !== 'undefined') {
    localStorage.removeItem('jwt_token');
  }
};

export const isTokenValid = (): boolean => {
  try {
    const token = getToken();
    if (!token) {
      return false;
    }

    const decoded = jwtDecode<JWTPayload>(token);
    const currentTime = Date.now() / 1000;
    return decoded.exp > currentTime;
  } catch (error) {
    console.error('Error validating token:', error);
    return false;
  }
};

export const getUserFromToken = async (): Promise<UserData | null> => {
  try {
    const token = getToken();
    if (!token) {
      return null;
    }

    // Fetch user profile from API
    const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/auth/profile`, {
      headers: {
        'Authorization': `Bearer ${token}`
      }
    });

    if (!response.ok) {
      throw new Error('Failed to fetch user profile');
    }

    const userData = await response.json();
    return userData;

  } catch (error) {
    console.error('Error getting user from token:', error);
    return null;
  }
};