export interface UserData {
  id: string;
  username: string;
  email: string;
}

export interface DecodedToken {
  user: UserData;
  exp: number;
  iat: number;
}



export interface Task {
  id: number;
  title: string;
  description: string;
  assignee_id: number;
  assigner_id: number;
  priority: 'HIGH' | 'MEDIUM' | 'LOW';
  status: 'TO_DO' | 'IN_PROGRESS' | 'DONE';
  created_at: string;
  updated_at: string;
}